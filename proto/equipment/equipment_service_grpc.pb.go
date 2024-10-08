// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: equipment_service.proto

package equipment

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
	EquipmentService_FindAllEquipments_FullMethodName = "/EquipmentService/FindAllEquipments"
)

// EquipmentServiceClient is the client API for EquipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EquipmentServiceClient interface {
	FindAllEquipments(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
}

type equipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquipmentServiceClient(cc grpc.ClientConnInterface) EquipmentServiceClient {
	return &equipmentServiceClient{cc}
}

func (c *equipmentServiceClient) FindAllEquipments(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, EquipmentService_FindAllEquipments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquipmentServiceServer is the server API for EquipmentService service.
// All implementations must embed UnimplementedEquipmentServiceServer
// for forward compatibility
type EquipmentServiceServer interface {
	FindAllEquipments(context.Context, *FindAllRequest) (*FindAllResponse, error)
	mustEmbedUnimplementedEquipmentServiceServer()
}

// UnimplementedEquipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEquipmentServiceServer struct {
}

func (UnimplementedEquipmentServiceServer) FindAllEquipments(context.Context, *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEquipments not implemented")
}
func (UnimplementedEquipmentServiceServer) mustEmbedUnimplementedEquipmentServiceServer() {}

// UnsafeEquipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EquipmentServiceServer will
// result in compilation errors.
type UnsafeEquipmentServiceServer interface {
	mustEmbedUnimplementedEquipmentServiceServer()
}

func RegisterEquipmentServiceServer(s grpc.ServiceRegistrar, srv EquipmentServiceServer) {
	s.RegisterService(&EquipmentService_ServiceDesc, srv)
}

func _EquipmentService_FindAllEquipments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).FindAllEquipments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EquipmentService_FindAllEquipments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).FindAllEquipments(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EquipmentService_ServiceDesc is the grpc.ServiceDesc for EquipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EquipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EquipmentService",
	HandlerType: (*EquipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllEquipments",
			Handler:    _EquipmentService_FindAllEquipments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "equipment_service.proto",
}
