// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tariff.proto

package tariff

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
	TariffService_Get_FullMethodName    = "/tariff.TariffService/Get"
	TariffService_GetAll_FullMethodName = "/tariff.TariffService/GetAll"
	TariffService_Create_FullMethodName = "/tariff.TariffService/Create"
	TariffService_Update_FullMethodName = "/tariff.TariffService/Update"
	TariffService_Delete_FullMethodName = "/tariff.TariffService/Delete"
)

// TariffServiceClient is the client API for TariffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TariffServiceClient interface {
	Get(ctx context.Context, in *GetTariffReq, opts ...grpc.CallOption) (*GetTariffRes, error)
	GetAll(ctx context.Context, in *GetAllTariffReq, opts ...grpc.CallOption) (*GetAllTariffRes, error)
	Create(ctx context.Context, in *CreateTariffReq, opts ...grpc.CallOption) (*CreateTariffRes, error)
	Update(ctx context.Context, in *UpdateTariffReq, opts ...grpc.CallOption) (*UpdateTariffRes, error)
	Delete(ctx context.Context, in *DeleteTariffReq, opts ...grpc.CallOption) (*DeleteTariffRes, error)
}

type tariffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTariffServiceClient(cc grpc.ClientConnInterface) TariffServiceClient {
	return &tariffServiceClient{cc}
}

func (c *tariffServiceClient) Get(ctx context.Context, in *GetTariffReq, opts ...grpc.CallOption) (*GetTariffRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTariffRes)
	err := c.cc.Invoke(ctx, TariffService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) GetAll(ctx context.Context, in *GetAllTariffReq, opts ...grpc.CallOption) (*GetAllTariffRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTariffRes)
	err := c.cc.Invoke(ctx, TariffService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) Create(ctx context.Context, in *CreateTariffReq, opts ...grpc.CallOption) (*CreateTariffRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTariffRes)
	err := c.cc.Invoke(ctx, TariffService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) Update(ctx context.Context, in *UpdateTariffReq, opts ...grpc.CallOption) (*UpdateTariffRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTariffRes)
	err := c.cc.Invoke(ctx, TariffService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) Delete(ctx context.Context, in *DeleteTariffReq, opts ...grpc.CallOption) (*DeleteTariffRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTariffRes)
	err := c.cc.Invoke(ctx, TariffService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TariffServiceServer is the server API for TariffService service.
// All implementations must embed UnimplementedTariffServiceServer
// for forward compatibility.
type TariffServiceServer interface {
	Get(context.Context, *GetTariffReq) (*GetTariffRes, error)
	GetAll(context.Context, *GetAllTariffReq) (*GetAllTariffRes, error)
	Create(context.Context, *CreateTariffReq) (*CreateTariffRes, error)
	Update(context.Context, *UpdateTariffReq) (*UpdateTariffRes, error)
	Delete(context.Context, *DeleteTariffReq) (*DeleteTariffRes, error)
	mustEmbedUnimplementedTariffServiceServer()
}

// UnimplementedTariffServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTariffServiceServer struct{}

func (UnimplementedTariffServiceServer) Get(context.Context, *GetTariffReq) (*GetTariffRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTariffServiceServer) GetAll(context.Context, *GetAllTariffReq) (*GetAllTariffRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTariffServiceServer) Create(context.Context, *CreateTariffReq) (*CreateTariffRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTariffServiceServer) Update(context.Context, *UpdateTariffReq) (*UpdateTariffRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTariffServiceServer) Delete(context.Context, *DeleteTariffReq) (*DeleteTariffRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTariffServiceServer) mustEmbedUnimplementedTariffServiceServer() {}
func (UnimplementedTariffServiceServer) testEmbeddedByValue()                       {}

// UnsafeTariffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TariffServiceServer will
// result in compilation errors.
type UnsafeTariffServiceServer interface {
	mustEmbedUnimplementedTariffServiceServer()
}

func RegisterTariffServiceServer(s grpc.ServiceRegistrar, srv TariffServiceServer) {
	// If the following call pancis, it indicates UnimplementedTariffServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TariffService_ServiceDesc, srv)
}

func _TariffService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTariffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Get(ctx, req.(*GetTariffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTariffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).GetAll(ctx, req.(*GetAllTariffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTariffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Create(ctx, req.(*CreateTariffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTariffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Update(ctx, req.(*UpdateTariffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTariffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Delete(ctx, req.(*DeleteTariffReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TariffService_ServiceDesc is the grpc.ServiceDesc for TariffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TariffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tariff.TariffService",
	HandlerType: (*TariffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TariffService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TariffService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TariffService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TariffService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TariffService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tariff.proto",
}
