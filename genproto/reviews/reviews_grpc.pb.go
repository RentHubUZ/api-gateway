// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: reviews.proto

package reviews

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReviewsClient is the client API for Reviews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewsClient interface {
	CreateReview(ctx context.Context, in *CreateReviewReq, opts ...grpc.CallOption) (*CreateReviewRes, error)
	GetAllReview(ctx context.Context, in *GetAllReviewReq, opts ...grpc.CallOption) (*GetAllReviewRes, error)
	GetByIdReview(ctx context.Context, in *GetByIdReviewReq, opts ...grpc.CallOption) (*GetByIdReviewRes, error)
	DeleteReview(ctx context.Context, in *DeleteReviewReq, opts ...grpc.CallOption) (*DeleteReviewRes, error)
}

type reviewsClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewsClient(cc grpc.ClientConnInterface) ReviewsClient {
	return &reviewsClient{cc}
}

func (c *reviewsClient) CreateReview(ctx context.Context, in *CreateReviewReq, opts ...grpc.CallOption) (*CreateReviewRes, error) {
	out := new(CreateReviewRes)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) GetAllReview(ctx context.Context, in *GetAllReviewReq, opts ...grpc.CallOption) (*GetAllReviewRes, error) {
	out := new(GetAllReviewRes)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/GetAllReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) GetByIdReview(ctx context.Context, in *GetByIdReviewReq, opts ...grpc.CallOption) (*GetByIdReviewRes, error) {
	out := new(GetByIdReviewRes)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/GetByIdReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) DeleteReview(ctx context.Context, in *DeleteReviewReq, opts ...grpc.CallOption) (*DeleteReviewRes, error) {
	out := new(DeleteReviewRes)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewsServer is the server API for Reviews service.
// All implementations must embed UnimplementedReviewsServer
// for forward compatibility
type ReviewsServer interface {
	CreateReview(context.Context, *CreateReviewReq) (*CreateReviewRes, error)
	GetAllReview(context.Context, *GetAllReviewReq) (*GetAllReviewRes, error)
	GetByIdReview(context.Context, *GetByIdReviewReq) (*GetByIdReviewRes, error)
	DeleteReview(context.Context, *DeleteReviewReq) (*DeleteReviewRes, error)
	mustEmbedUnimplementedReviewsServer()
}

// UnimplementedReviewsServer must be embedded to have forward compatible implementations.
type UnimplementedReviewsServer struct {
}

func (UnimplementedReviewsServer) CreateReview(context.Context, *CreateReviewReq) (*CreateReviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewsServer) GetAllReview(context.Context, *GetAllReviewReq) (*GetAllReviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReview not implemented")
}
func (UnimplementedReviewsServer) GetByIdReview(context.Context, *GetByIdReviewReq) (*GetByIdReviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdReview not implemented")
}
func (UnimplementedReviewsServer) DeleteReview(context.Context, *DeleteReviewReq) (*DeleteReviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedReviewsServer) mustEmbedUnimplementedReviewsServer() {}

// UnsafeReviewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewsServer will
// result in compilation errors.
type UnsafeReviewsServer interface {
	mustEmbedUnimplementedReviewsServer()
}

func RegisterReviewsServer(s grpc.ServiceRegistrar, srv ReviewsServer) {
	s.RegisterService(&Reviews_ServiceDesc, srv)
}

func _Reviews_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).CreateReview(ctx, req.(*CreateReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_GetAllReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).GetAllReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/GetAllReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).GetAllReview(ctx, req.(*GetAllReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_GetByIdReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).GetByIdReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/GetByIdReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).GetByIdReview(ctx, req.(*GetByIdReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).DeleteReview(ctx, req.(*DeleteReviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Reviews_ServiceDesc is the grpc.ServiceDesc for Reviews service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reviews_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reviews.Reviews",
	HandlerType: (*ReviewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _Reviews_CreateReview_Handler,
		},
		{
			MethodName: "GetAllReview",
			Handler:    _Reviews_GetAllReview_Handler,
		},
		{
			MethodName: "GetByIdReview",
			Handler:    _Reviews_GetByIdReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _Reviews_DeleteReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}
