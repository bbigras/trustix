// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: sth.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type STH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogRoot   []byte `protobuf:"bytes,1,opt,name=LogRoot,proto3" json:"LogRoot,omitempty"`
	TreeSize  uint64 `protobuf:"varint,2,opt,name=TreeSize,proto3" json:"TreeSize,omitempty"`
	MapRoot   []byte `protobuf:"bytes,3,opt,name=MapRoot,proto3" json:"MapRoot,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *STH) Reset() {
	*x = STH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STH) ProtoMessage() {}

func (x *STH) ProtoReflect() protoreflect.Message {
	mi := &file_sth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STH.ProtoReflect.Descriptor instead.
func (*STH) Descriptor() ([]byte, []int) {
	return file_sth_proto_rawDescGZIP(), []int{0}
}

func (x *STH) GetLogRoot() []byte {
	if x != nil {
		return x.LogRoot
	}
	return nil
}

func (x *STH) GetTreeSize() uint64 {
	if x != nil {
		return x.TreeSize
	}
	return 0
}

func (x *STH) GetMapRoot() []byte {
	if x != nil {
		return x.MapRoot
	}
	return nil
}

func (x *STH) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_sth_proto protoreflect.FileDescriptor

var file_sth_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x03, 0x53,
	0x54, 0x48, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x52,
	0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x61, 0x70, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x77, 0x65, 0x61, 0x67, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sth_proto_rawDescOnce sync.Once
	file_sth_proto_rawDescData = file_sth_proto_rawDesc
)

func file_sth_proto_rawDescGZIP() []byte {
	file_sth_proto_rawDescOnce.Do(func() {
		file_sth_proto_rawDescData = protoimpl.X.CompressGZIP(file_sth_proto_rawDescData)
	})
	return file_sth_proto_rawDescData
}

var file_sth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sth_proto_goTypes = []interface{}{
	(*STH)(nil), // 0: STH
}
var file_sth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sth_proto_init() }
func file_sth_proto_init() {
	if File_sth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STH); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sth_proto_goTypes,
		DependencyIndexes: file_sth_proto_depIdxs,
		MessageInfos:      file_sth_proto_msgTypes,
	}.Build()
	File_sth_proto = out.File
	file_sth_proto_rawDesc = nil
	file_sth_proto_goTypes = nil
	file_sth_proto_depIdxs = nil
}
