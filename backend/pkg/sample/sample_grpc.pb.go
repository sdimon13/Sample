// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: sample.proto

package sample

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

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleClient interface {
	Hello(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error)
	ServiceList(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error)
	AppointmentList(ctx context.Context, in *AppointmentServiceListRequest, opts ...grpc.CallOption) (*AppointmentServiceListResponse, error)
	AdminAppointmentList(ctx context.Context, in *AppointmentServiceListRequest, opts ...grpc.CallOption) (*AdminAppointmentServiceListResponse, error)
	AppointmentCreate(ctx context.Context, in *AppointmentServiceCreateRequest, opts ...grpc.CallOption) (*AppointmentServiceGetResponse, error)
	// Обновление критерия доступа оператора по ID
	AppointmentUpdate(ctx context.Context, in *AppointmentServiceUpdateRequest, opts ...grpc.CallOption) (*AppointmentServiceGetResponse, error)
}

type sampleClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleClient(cc grpc.ClientConnInterface) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) Hello(ctx context.Context, in *SampleRequest, opts ...grpc.CallOption) (*SampleResponse, error) {
	out := new(SampleResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) ServiceList(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListResponse, error) {
	out := new(ServiceListResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/ServiceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) AppointmentList(ctx context.Context, in *AppointmentServiceListRequest, opts ...grpc.CallOption) (*AppointmentServiceListResponse, error) {
	out := new(AppointmentServiceListResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/AppointmentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) AdminAppointmentList(ctx context.Context, in *AppointmentServiceListRequest, opts ...grpc.CallOption) (*AdminAppointmentServiceListResponse, error) {
	out := new(AdminAppointmentServiceListResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/AdminAppointmentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) AppointmentCreate(ctx context.Context, in *AppointmentServiceCreateRequest, opts ...grpc.CallOption) (*AppointmentServiceGetResponse, error) {
	out := new(AppointmentServiceGetResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/AppointmentCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) AppointmentUpdate(ctx context.Context, in *AppointmentServiceUpdateRequest, opts ...grpc.CallOption) (*AppointmentServiceGetResponse, error) {
	out := new(AppointmentServiceGetResponse)
	err := c.cc.Invoke(ctx, "/sample.Sample/AppointmentUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
// All implementations must embed UnimplementedSampleServer
// for forward compatibility
type SampleServer interface {
	Hello(context.Context, *SampleRequest) (*SampleResponse, error)
	ServiceList(context.Context, *ServiceListRequest) (*ServiceListResponse, error)
	AppointmentList(context.Context, *AppointmentServiceListRequest) (*AppointmentServiceListResponse, error)
	AdminAppointmentList(context.Context, *AppointmentServiceListRequest) (*AdminAppointmentServiceListResponse, error)
	AppointmentCreate(context.Context, *AppointmentServiceCreateRequest) (*AppointmentServiceGetResponse, error)
	// Обновление критерия доступа оператора по ID
	AppointmentUpdate(context.Context, *AppointmentServiceUpdateRequest) (*AppointmentServiceGetResponse, error)
	mustEmbedUnimplementedSampleServer()
}

// UnimplementedSampleServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (UnimplementedSampleServer) Hello(context.Context, *SampleRequest) (*SampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedSampleServer) ServiceList(context.Context, *ServiceListRequest) (*ServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceList not implemented")
}
func (UnimplementedSampleServer) AppointmentList(context.Context, *AppointmentServiceListRequest) (*AppointmentServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointmentList not implemented")
}
func (UnimplementedSampleServer) AdminAppointmentList(context.Context, *AppointmentServiceListRequest) (*AdminAppointmentServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAppointmentList not implemented")
}
func (UnimplementedSampleServer) AppointmentCreate(context.Context, *AppointmentServiceCreateRequest) (*AppointmentServiceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointmentCreate not implemented")
}
func (UnimplementedSampleServer) AppointmentUpdate(context.Context, *AppointmentServiceUpdateRequest) (*AppointmentServiceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppointmentUpdate not implemented")
}
func (UnimplementedSampleServer) mustEmbedUnimplementedSampleServer() {}

// UnsafeSampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServer will
// result in compilation errors.
type UnsafeSampleServer interface {
	mustEmbedUnimplementedSampleServer()
}

func RegisterSampleServer(s grpc.ServiceRegistrar, srv SampleServer) {
	s.RegisterService(&Sample_ServiceDesc, srv)
}

func _Sample_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).Hello(ctx, req.(*SampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_ServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).ServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/ServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).ServiceList(ctx, req.(*ServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_AppointmentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).AppointmentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/AppointmentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).AppointmentList(ctx, req.(*AppointmentServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_AdminAppointmentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).AdminAppointmentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/AdminAppointmentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).AdminAppointmentList(ctx, req.(*AppointmentServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_AppointmentCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).AppointmentCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/AppointmentCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).AppointmentCreate(ctx, req.(*AppointmentServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_AppointmentUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).AppointmentUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/AppointmentUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).AppointmentUpdate(ctx, req.(*AppointmentServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sample_ServiceDesc is the grpc.ServiceDesc for Sample service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sample_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Sample_Hello_Handler,
		},
		{
			MethodName: "ServiceList",
			Handler:    _Sample_ServiceList_Handler,
		},
		{
			MethodName: "AppointmentList",
			Handler:    _Sample_AppointmentList_Handler,
		},
		{
			MethodName: "AdminAppointmentList",
			Handler:    _Sample_AdminAppointmentList_Handler,
		},
		{
			MethodName: "AppointmentCreate",
			Handler:    _Sample_AppointmentCreate_Handler,
		},
		{
			MethodName: "AppointmentUpdate",
			Handler:    _Sample_AppointmentUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}
