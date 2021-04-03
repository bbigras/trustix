// Copyright (C) 2021 Tweag IO
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"context"

	"github.com/tweag/trustix/packages/trustix-proto/api"
	"github.com/tweag/trustix/packages/trustix-proto/schema"
)

type TrustixLogAPI interface {
	// GetSTH - Get a signed tree head
	GetSTH(context.Context, *api.STHRequest) (*schema.STH, error)

	// Get log consistency proof
	GetLogConsistencyProof(context.Context, *api.GetLogConsistencyProofRequest) (*api.ProofResponse, error)

	// Get leaf audit proof
	GetLogAuditProof(context.Context, *api.GetLogAuditProofRequest) (*api.ProofResponse, error)

	// Get log entries (batched)
	GetLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error)

	// Get value from the map
	GetMapValue(context.Context, *api.GetMapValueRequest) (*api.MapValueResponse, error)

	// Get content-addressed computation outputs
	GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error)

	// Get map head log consistency proof
	GetMHLogConsistencyProof(context.Context, *api.GetLogConsistencyProofRequest) (*api.ProofResponse, error)

	// Get map head log leaf audit proof
	GetMHLogAuditProof(context.Context, *api.GetLogAuditProofRequest) (*api.ProofResponse, error)

	// Get map head log entries (batched)
	GetMHLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error)
}
