// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: reservation.proto

package reservation

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

const (
	ReservationService_Create_FullMethodName = "/reservation.ReservationService/Create"
	ReservationService_Get_FullMethodName    = "/reservation.ReservationService/Get"
	ReservationService_GetAll_FullMethodName = "/reservation.ReservationService/GetAll"
	ReservationService_Update_FullMethodName = "/reservation.ReservationService/Update"
	ReservationService_Delete_FullMethodName = "/reservation.ReservationService/Delete"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Create(ctx context.Context, in *ReservationReq, opts ...grpc.CallOption) (*Reservation, error)
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*ReservationRes, error)
	GetAll(ctx context.Context, in *GetAllReservationReq, opts ...grpc.CallOption) (*GetAllReservationRes, error)
	Update(ctx context.Context, in *ReservationUpdate, opts ...grpc.CallOption) (*Reservation, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Create(ctx context.Context, in *ReservationReq, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*ReservationRes, error) {
	out := new(ReservationRes)
	err := c.cc.Invoke(ctx, ReservationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllReservationReq, opts ...grpc.CallOption) (*GetAllReservationRes, error) {
	out := new(GetAllReservationRes)
	err := c.cc.Invoke(ctx, ReservationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Update(ctx context.Context, in *ReservationUpdate, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, ReservationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ReservationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Create(context.Context, *ReservationReq) (*Reservation, error)
	Get(context.Context, *GetByIdReq) (*ReservationRes, error)
	GetAll(context.Context, *GetAllReservationReq) (*GetAllReservationRes, error)
	Update(context.Context, *ReservationUpdate) (*Reservation, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Create(context.Context, *ReservationReq) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReservationServiceServer) Get(context.Context, *GetByIdReq) (*ReservationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllReservationReq) (*GetAllReservationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationServiceServer) Update(context.Context, *ReservationUpdate) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReservationServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Create(ctx, req.(*ReservationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReservationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllReservationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Update(ctx, req.(*ReservationUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Delete(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReservationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReservationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReservationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation.proto",
}
