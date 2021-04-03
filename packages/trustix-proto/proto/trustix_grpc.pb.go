// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	api "github.com/tweag/trustix/packages/trustix-proto/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrustixCombinedRPCClient is the client API for TrustixCombinedRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrustixCombinedRPCClient interface {
	// Get map[LogID]Log
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error)
	// TODO: I'm not sure if this belongs here in it's current shape...
	GetLogEntries(ctx context.Context, in *GetLogEntriesRequestNamed, opts ...grpc.CallOption) (*api.LogEntriesResponse, error)
	// Get map[LogID]OutputHash
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EntriesResponse, error)
	GetStream(ctx context.Context, opts ...grpc.CallOption) (TrustixCombinedRPC_GetStreamClient, error)
	// Compare(inputHash)
	Decide(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DecisionResponse, error)
	DecideStream(ctx context.Context, opts ...grpc.CallOption) (TrustixCombinedRPC_DecideStreamClient, error)
	// Get stored value by digest (TODO: Remove, it's a duplicate from api.proto
	GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type trustixCombinedRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustixCombinedRPCClient(cc grpc.ClientConnInterface) TrustixCombinedRPCClient {
	return &trustixCombinedRPCClient{cc}
}

func (c *trustixCombinedRPCClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) GetLogEntries(ctx context.Context, in *GetLogEntriesRequestNamed, opts ...grpc.CallOption) (*api.LogEntriesResponse, error) {
	out := new(api.LogEntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/GetLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EntriesResponse, error) {
	out := new(EntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) GetStream(ctx context.Context, opts ...grpc.CallOption) (TrustixCombinedRPC_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrustixCombinedRPC_ServiceDesc.Streams[0], "/trustix.TrustixCombinedRPC/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &trustixCombinedRPCGetStreamClient{stream}
	return x, nil
}

type TrustixCombinedRPC_GetStreamClient interface {
	Send(*KeyRequest) error
	Recv() (*EntriesResponse, error)
	grpc.ClientStream
}

type trustixCombinedRPCGetStreamClient struct {
	grpc.ClientStream
}

func (x *trustixCombinedRPCGetStreamClient) Send(m *KeyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trustixCombinedRPCGetStreamClient) Recv() (*EntriesResponse, error) {
	m := new(EntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trustixCombinedRPCClient) Decide(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*DecisionResponse, error) {
	out := new(DecisionResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/Decide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) DecideStream(ctx context.Context, opts ...grpc.CallOption) (TrustixCombinedRPC_DecideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TrustixCombinedRPC_ServiceDesc.Streams[1], "/trustix.TrustixCombinedRPC/DecideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &trustixCombinedRPCDecideStreamClient{stream}
	return x, nil
}

type TrustixCombinedRPC_DecideStreamClient interface {
	Send(*KeyRequest) error
	Recv() (*DecisionResponse, error)
	grpc.ClientStream
}

type trustixCombinedRPCDecideStreamClient struct {
	grpc.ClientStream
}

func (x *trustixCombinedRPCDecideStreamClient) Send(m *KeyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *trustixCombinedRPCDecideStreamClient) Recv() (*DecisionResponse, error) {
	m := new(DecisionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trustixCombinedRPCClient) GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error) {
	out := new(api.ValueResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustixCombinedRPCClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/trustix.TrustixCombinedRPC/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustixCombinedRPCServer is the server API for TrustixCombinedRPC service.
// All implementations must embed UnimplementedTrustixCombinedRPCServer
// for forward compatibility
type TrustixCombinedRPCServer interface {
	// Get map[LogID]Log
	Logs(context.Context, *LogsRequest) (*LogsResponse, error)
	// TODO: I'm not sure if this belongs here in it's current shape...
	GetLogEntries(context.Context, *GetLogEntriesRequestNamed) (*api.LogEntriesResponse, error)
	// Get map[LogID]OutputHash
	Get(context.Context, *KeyRequest) (*EntriesResponse, error)
	GetStream(TrustixCombinedRPC_GetStreamServer) error
	// Compare(inputHash)
	Decide(context.Context, *KeyRequest) (*DecisionResponse, error)
	DecideStream(TrustixCombinedRPC_DecideStreamServer) error
	// Get stored value by digest (TODO: Remove, it's a duplicate from api.proto
	GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	mustEmbedUnimplementedTrustixCombinedRPCServer()
}

// UnimplementedTrustixCombinedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTrustixCombinedRPCServer struct {
}

func (UnimplementedTrustixCombinedRPCServer) Logs(context.Context, *LogsRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) GetLogEntries(context.Context, *GetLogEntriesRequestNamed) (*api.LogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogEntries not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) Get(context.Context, *KeyRequest) (*EntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) GetStream(TrustixCombinedRPC_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) Decide(context.Context, *KeyRequest) (*DecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decide not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) DecideStream(TrustixCombinedRPC_DecideStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DecideStream not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedTrustixCombinedRPCServer) mustEmbedUnimplementedTrustixCombinedRPCServer() {}

// UnsafeTrustixCombinedRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustixCombinedRPCServer will
// result in compilation errors.
type UnsafeTrustixCombinedRPCServer interface {
	mustEmbedUnimplementedTrustixCombinedRPCServer()
}

func RegisterTrustixCombinedRPCServer(s grpc.ServiceRegistrar, srv TrustixCombinedRPCServer) {
	s.RegisterService(&TrustixCombinedRPC_ServiceDesc, srv)
}

func _TrustixCombinedRPC_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).Logs(ctx, req.(*LogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_GetLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogEntriesRequestNamed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).GetLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/GetLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).GetLogEntries(ctx, req.(*GetLogEntriesRequestNamed))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrustixCombinedRPCServer).GetStream(&trustixCombinedRPCGetStreamServer{stream})
}

type TrustixCombinedRPC_GetStreamServer interface {
	Send(*EntriesResponse) error
	Recv() (*KeyRequest, error)
	grpc.ServerStream
}

type trustixCombinedRPCGetStreamServer struct {
	grpc.ServerStream
}

func (x *trustixCombinedRPCGetStreamServer) Send(m *EntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trustixCombinedRPCGetStreamServer) Recv() (*KeyRequest, error) {
	m := new(KeyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TrustixCombinedRPC_Decide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).Decide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/Decide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).Decide(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_DecideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TrustixCombinedRPCServer).DecideStream(&trustixCombinedRPCDecideStreamServer{stream})
}

type TrustixCombinedRPC_DecideStreamServer interface {
	Send(*DecisionResponse) error
	Recv() (*KeyRequest, error)
	grpc.ServerStream
}

type trustixCombinedRPCDecideStreamServer struct {
	grpc.ServerStream
}

func (x *trustixCombinedRPCDecideStreamServer) Send(m *DecisionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *trustixCombinedRPCDecideStreamServer) Recv() (*KeyRequest, error) {
	m := new(KeyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TrustixCombinedRPC_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).GetValue(ctx, req.(*api.ValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustixCombinedRPC_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustixCombinedRPCServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.TrustixCombinedRPC/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustixCombinedRPCServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrustixCombinedRPC_ServiceDesc is the grpc.ServiceDesc for TrustixCombinedRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrustixCombinedRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.TrustixCombinedRPC",
	HandlerType: (*TrustixCombinedRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _TrustixCombinedRPC_Logs_Handler,
		},
		{
			MethodName: "GetLogEntries",
			Handler:    _TrustixCombinedRPC_GetLogEntries_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TrustixCombinedRPC_Get_Handler,
		},
		{
			MethodName: "Decide",
			Handler:    _TrustixCombinedRPC_Decide_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _TrustixCombinedRPC_GetValue_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _TrustixCombinedRPC_Submit_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _TrustixCombinedRPC_Flush_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _TrustixCombinedRPC_GetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DecideStream",
			Handler:       _TrustixCombinedRPC_DecideStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "trustix.proto",
}
