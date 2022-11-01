// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: api/api.proto

package api

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

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	GetQuantity(ctx context.Context, in *GetQuantityRequest, opts ...grpc.CallOption) (*GetQuantityResponse, error)
	UpdateQuantity(ctx context.Context, in *UpdateQuantityRequest, opts ...grpc.CallOption) (*UpdateQuantityResponse, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetQuantity(ctx context.Context, in *GetQuantityRequest, opts ...grpc.CallOption) (*GetQuantityResponse, error) {
	out := new(GetQuantityResponse)
	err := c.cc.Invoke(ctx, "/api.Inventory/GetQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) UpdateQuantity(ctx context.Context, in *UpdateQuantityRequest, opts ...grpc.CallOption) (*UpdateQuantityResponse, error) {
	out := new(UpdateQuantityResponse)
	err := c.cc.Invoke(ctx, "/api.Inventory/UpdateQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	GetQuantity(context.Context, *GetQuantityRequest) (*GetQuantityResponse, error)
	UpdateQuantity(context.Context, *UpdateQuantityRequest) (*UpdateQuantityResponse, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) GetQuantity(context.Context, *GetQuantityRequest) (*GetQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuantity not implemented")
}
func (UnimplementedInventoryServer) UpdateQuantity(context.Context, *UpdateQuantityRequest) (*UpdateQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuantity not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_GetQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Inventory/GetQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetQuantity(ctx, req.(*GetQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_UpdateQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).UpdateQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Inventory/UpdateQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).UpdateQuantity(ctx, req.(*UpdateQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuantity",
			Handler:    _Inventory_GetQuantity_Handler,
		},
		{
			MethodName: "UpdateQuantity",
			Handler:    _Inventory_UpdateQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
