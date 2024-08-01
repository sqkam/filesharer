// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: file/v1/file.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	File_ListByAddr_FullMethodName            = "/api.file.v1.File/ListByAddr"
	File_GetDetailByAddr_FullMethodName       = "/api.file.v1.File/GetDetailByAddr"
	File_DownloadByAddr_FullMethodName        = "/api.file.v1.File/DownloadByAddr"
	File_DownloadDirByAddr_FullMethodName     = "/api.file.v1.File/DownloadDirByAddr"
	File_DownloadByAddrHttp_FullMethodName    = "/api.file.v1.File/DownloadByAddrHttp"
	File_DownloadDirByAddrHttp_FullMethodName = "/api.file.v1.File/DownloadDirByAddrHttp"
	File_ListNode_FullMethodName              = "/api.file.v1.File/ListNode"
)

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	ListByAddr(ctx context.Context, in *ListByAddrRequest, opts ...grpc.CallOption) (*ListByAddrReply, error)
	GetDetailByAddr(ctx context.Context, in *GetDetailByAddrRequest, opts ...grpc.CallOption) (*GetDetailByAddrReply, error)
	DownloadByAddr(ctx context.Context, in *DownloadByAddrRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadByAddrReply], error)
	DownloadDirByAddr(ctx context.Context, in *DownloadDirByAddrRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadDirByAddrReply], error)
	DownloadByAddrHttp(ctx context.Context, in *DownloadByAddrRequest, opts ...grpc.CallOption) (*DownloadByAddrReply, error)
	DownloadDirByAddrHttp(ctx context.Context, in *DownloadDirByAddrRequest, opts ...grpc.CallOption) (*DownloadDirByAddrReply, error)
	ListNode(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeReply, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) ListByAddr(ctx context.Context, in *ListByAddrRequest, opts ...grpc.CallOption) (*ListByAddrReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListByAddrReply)
	err := c.cc.Invoke(ctx, File_ListByAddr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) GetDetailByAddr(ctx context.Context, in *GetDetailByAddrRequest, opts ...grpc.CallOption) (*GetDetailByAddrReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDetailByAddrReply)
	err := c.cc.Invoke(ctx, File_GetDetailByAddr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) DownloadByAddr(ctx context.Context, in *DownloadByAddrRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadByAddrReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &File_ServiceDesc.Streams[0], File_DownloadByAddr_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadByAddrRequest, DownloadByAddrReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type File_DownloadByAddrClient = grpc.ServerStreamingClient[DownloadByAddrReply]

func (c *fileClient) DownloadDirByAddr(ctx context.Context, in *DownloadDirByAddrRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadDirByAddrReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &File_ServiceDesc.Streams[1], File_DownloadDirByAddr_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DownloadDirByAddrRequest, DownloadDirByAddrReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type File_DownloadDirByAddrClient = grpc.ServerStreamingClient[DownloadDirByAddrReply]

func (c *fileClient) DownloadByAddrHttp(ctx context.Context, in *DownloadByAddrRequest, opts ...grpc.CallOption) (*DownloadByAddrReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadByAddrReply)
	err := c.cc.Invoke(ctx, File_DownloadByAddrHttp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) DownloadDirByAddrHttp(ctx context.Context, in *DownloadDirByAddrRequest, opts ...grpc.CallOption) (*DownloadDirByAddrReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadDirByAddrReply)
	err := c.cc.Invoke(ctx, File_DownloadDirByAddrHttp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) ListNode(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNodeReply)
	err := c.cc.Invoke(ctx, File_ListNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility.
type FileServer interface {
	ListByAddr(context.Context, *ListByAddrRequest) (*ListByAddrReply, error)
	GetDetailByAddr(context.Context, *GetDetailByAddrRequest) (*GetDetailByAddrReply, error)
	DownloadByAddr(*DownloadByAddrRequest, grpc.ServerStreamingServer[DownloadByAddrReply]) error
	DownloadDirByAddr(*DownloadDirByAddrRequest, grpc.ServerStreamingServer[DownloadDirByAddrReply]) error
	DownloadByAddrHttp(context.Context, *DownloadByAddrRequest) (*DownloadByAddrReply, error)
	DownloadDirByAddrHttp(context.Context, *DownloadDirByAddrRequest) (*DownloadDirByAddrReply, error)
	ListNode(context.Context, *ListNodeRequest) (*ListNodeReply, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServer struct{}

func (UnimplementedFileServer) ListByAddr(context.Context, *ListByAddrRequest) (*ListByAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByAddr not implemented")
}
func (UnimplementedFileServer) GetDetailByAddr(context.Context, *GetDetailByAddrRequest) (*GetDetailByAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailByAddr not implemented")
}
func (UnimplementedFileServer) DownloadByAddr(*DownloadByAddrRequest, grpc.ServerStreamingServer[DownloadByAddrReply]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadByAddr not implemented")
}
func (UnimplementedFileServer) DownloadDirByAddr(*DownloadDirByAddrRequest, grpc.ServerStreamingServer[DownloadDirByAddrReply]) error {
	return status.Errorf(codes.Unimplemented, "method DownloadDirByAddr not implemented")
}
func (UnimplementedFileServer) DownloadByAddrHttp(context.Context, *DownloadByAddrRequest) (*DownloadByAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadByAddrHttp not implemented")
}
func (UnimplementedFileServer) DownloadDirByAddrHttp(context.Context, *DownloadDirByAddrRequest) (*DownloadDirByAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadDirByAddrHttp not implemented")
}
func (UnimplementedFileServer) ListNode(context.Context, *ListNodeRequest) (*ListNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNode not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}
func (UnimplementedFileServer) testEmbeddedByValue()              {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	// If the following call pancis, it indicates UnimplementedFileServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_ListByAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).ListByAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_ListByAddr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).ListByAddr(ctx, req.(*ListByAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_GetDetailByAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailByAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetDetailByAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_GetDetailByAddr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetDetailByAddr(ctx, req.(*GetDetailByAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_DownloadByAddr_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadByAddrRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServer).DownloadByAddr(m, &grpc.GenericServerStream[DownloadByAddrRequest, DownloadByAddrReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type File_DownloadByAddrServer = grpc.ServerStreamingServer[DownloadByAddrReply]

func _File_DownloadDirByAddr_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadDirByAddrRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServer).DownloadDirByAddr(m, &grpc.GenericServerStream[DownloadDirByAddrRequest, DownloadDirByAddrReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type File_DownloadDirByAddrServer = grpc.ServerStreamingServer[DownloadDirByAddrReply]

func _File_DownloadByAddrHttp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadByAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).DownloadByAddrHttp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_DownloadByAddrHttp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).DownloadByAddrHttp(ctx, req.(*DownloadByAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_DownloadDirByAddrHttp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadDirByAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).DownloadDirByAddrHttp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_DownloadDirByAddrHttp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).DownloadDirByAddrHttp(ctx, req.(*DownloadDirByAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_ListNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).ListNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_ListNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).ListNode(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.file.v1.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListByAddr",
			Handler:    _File_ListByAddr_Handler,
		},
		{
			MethodName: "GetDetailByAddr",
			Handler:    _File_GetDetailByAddr_Handler,
		},
		{
			MethodName: "DownloadByAddrHttp",
			Handler:    _File_DownloadByAddrHttp_Handler,
		},
		{
			MethodName: "DownloadDirByAddrHttp",
			Handler:    _File_DownloadDirByAddrHttp_Handler,
		},
		{
			MethodName: "ListNode",
			Handler:    _File_ListNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadByAddr",
			Handler:       _File_DownloadByAddr_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadDirByAddr",
			Handler:       _File_DownloadDirByAddr_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file/v1/file.proto",
}
