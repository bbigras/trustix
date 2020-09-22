package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/lazyledger/smt"
	"github.com/tweag/trustix/config"
	"github.com/tweag/trustix/signer"
	"github.com/tweag/trustix/sth"
	"github.com/tweag/trustix/storage"
	"github.com/tweag/trustix/transport"
	"hash"
	"time"
)

type FlagConfig struct {
	StateDirectory string
}

type TrustixCore struct {
	store      storage.TrustixStorage
	tree       *smt.SparseMerkleTree
	sthManager *sth.STHManager
	mapStore   *smtMapStore
	hasher     hash.Hash
}

func (s *TrustixCore) Query(key []byte) ([]byte, error) {
	var buf []byte

	err := s.store.View(func(txn storage.Transaction) error {
		s.mapStore.setTxn(txn)
		defer s.mapStore.unsetTxn()

		v, err := s.tree.Get(key)
		if err != nil {
			return err
		}
		buf = v

		// Check inclusion proof
		proof, err := s.tree.Prove(key)
		if err != nil {
			return err
		}
		root := s.tree.Root()

		if !smt.VerifyProof(proof, root, key, v, s.hasher) {
			return fmt.Errorf("Proof verification failed")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TrustixCore) Get(key []byte) ([]byte, error) {
	var buf []byte

	err := s.store.View(func(txn storage.Transaction) error {
		v, err := txn.Get(key)
		if err != nil {
			return err
		}
		buf = v

		return nil
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TrustixCore) Submit(key []byte, value []byte) error {
	return s.store.Update(func(txn storage.Transaction) error {
		s.mapStore.setTxn(txn)
		defer s.mapStore.unsetTxn()

		s.tree.Update(key, value)

		sth, err := s.sthManager.Sign()
		if err != nil {
			return err
		}

		return s.mapStore.Set([]byte("HEAD"), sth)
	})
}

func (s *TrustixCore) updateRoot() error {
	return s.store.View(func(txn storage.Transaction) error {
		s.mapStore.setTxn(txn)
		defer s.mapStore.unsetTxn()

		oldHead, err := txn.Get([]byte("HEAD"))
		if err != nil {
			return err
		} else {
			oldSTH := &sth.STH{}
			err = oldSTH.FromJSON(oldHead)
			if err != nil {
				return err
			}

			rootBytes, err := oldSTH.UnmarshalRoot()
			if err != nil {
				return err
			}

			if bytes.Compare(rootBytes, s.tree.Root()) != 0 {
				s.tree.SetRoot(rootBytes)
				fmt.Println("Updated root")
			}
		}

		return nil
	})
}

func CoreFromConfig(conf *config.LogConfig, flags *FlagConfig) (*TrustixCore, error) {

	hasher := sha256.New()

	sig, err := signer.FromConfig(conf.Signer)
	if err != nil {
		return nil, err
	}

	if conf.Mode == "trustix-log" {
		if !sig.CanSign() {
			return nil, fmt.Errorf("Cannot sign using the current configuration, aborting.")
		}
	}

	var store storage.TrustixStorage
	switch conf.Mode {
	case "trustix-log":
		store, err = storage.FromConfig(conf.Name, flags.StateDirectory, conf.Storage)
		if err != nil {
			return nil, err
		}
	case "trustix-follower":
		// // Followers have a local store for storing proofs
		// backingStore, err := storage.NativeStorageFromConfig(conf.Name, flags.StateDirectory, nil)
		// if err != nil {
		// 	return nil, err
		// }
		store, err = transport.NewGRPCTransport(conf.Transport.GRPC)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Mode '%s' unhandled", conf.Mode)
	}

	var tree *smt.SparseMerkleTree

	mapStore := newMapStore()

	err = store.View(func(txn storage.Transaction) error {
		mapStore.setTxn(txn)
		defer mapStore.unsetTxn()

		oldHead, err := txn.Get([]byte("HEAD"))
		if err != nil {
			// No STH yet, new tree
			if err == storage.ObjectNotFoundError {
				tree = smt.NewSparseMerkleTree(mapStore, hasher)
			} else {
				return err
			}
		} else {
			oldSTH := &sth.STH{}
			err = oldSTH.FromJSON(oldHead)
			if err != nil {
				return err
			}

			rootBytes, err := oldSTH.UnmarshalRoot()
			if err != nil {
				return err
			}

			tree = smt.ImportSparseMerkleTree(mapStore, hasher, rootBytes)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	sthManager := sth.NewSTHManager(tree, sig)

	core := &TrustixCore{
		store:      store,
		tree:       tree,
		sthManager: sthManager,
		mapStore:   mapStore,
		hasher:     hasher,
	}

	switch conf.Mode {
	case "trustix-follower":
		go func() {
			for {
				// TODO: This is just an arbitrary interval for testing, not particularly intelligent
				time.Sleep(10 * time.Second)
				err := core.updateRoot()
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	default:
	}

	return core, nil
}
