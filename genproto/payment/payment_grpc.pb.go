// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
<<<<<<< HEAD
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: payment.proto
=======
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.18.0
// source: protos/payment/payment.proto
>>>>>>> ea8d6ed22d08fcf7e6340d33733c9cf041a2b1b8

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
<<<<<<< HEAD
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7
=======
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PaymentService_Get_FullMethodName    = "/payment.PaymentService/Get"
	PaymentService_GetAll_FullMethodName = "/payment.PaymentService/GetAll"
	PaymentService_Create_FullMethodName = "/payment.PaymentService/Create"
	PaymentService_Delete_FullMethodName = "/payment.PaymentService/Delete"
)
>>>>>>> ea8d6ed22d08fcf7e6340d33733c9cf041a2b1b8

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Get(ctx context.Context, in *GetPaymentReq, opts ...grpc.CallOption) (*GetPaymentRes, error)
	GetAll(ctx context.Context, in *GetAllPaymentReq, opts ...grpc.CallOption) (*GetAllPaymentRes, error)
	Create(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentRes, error)
	Delete(ctx context.Context, in *DeletePaymentReq, opts ...grpc.CallOption) (*DeletePaymentRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Get(ctx context.Context, in *GetPaymentReq, opts ...grpc.CallOption) (*GetPaymentRes, error) {
	out := new(GetPaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetAll(ctx context.Context, in *GetAllPaymentReq, opts ...grpc.CallOption) (*GetAllPaymentRes, error) {
	out := new(GetAllPaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) Create(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentRes, error) {
	out := new(CreatePaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) Delete(ctx context.Context, in *DeletePaymentReq, opts ...grpc.CallOption) (*DeletePaymentRes, error) {
	out := new(DeletePaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	Get(context.Context, *GetPaymentReq) (*GetPaymentRes, error)
	GetAll(context.Context, *GetAllPaymentReq) (*GetAllPaymentRes, error)
	Create(context.Context, *CreatePaymentReq) (*CreatePaymentRes, error)
	Delete(context.Context, *DeletePaymentReq) (*DeletePaymentRes, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) Get(context.Context, *GetPaymentReq) (*GetPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPaymentServiceServer) GetAll(context.Context, *GetAllPaymentReq) (*GetAllPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPaymentServiceServer) Create(context.Context, *CreatePaymentReq) (*CreatePaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentServiceServer) Delete(context.Context, *DeletePaymentReq) (*DeletePaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Get(ctx, req.(*GetPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetAll(ctx, req.(*GetAllPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Create(ctx, req.(*CreatePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Delete(ctx, req.(*DeletePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PaymentService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PaymentService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PaymentService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PaymentService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/payment/payment.proto",
}
