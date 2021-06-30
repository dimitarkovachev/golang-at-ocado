// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

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

// SortingRobotClient is the client API for SortingRobot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SortingRobotClient interface {
	LoadItems(ctx context.Context, in *LoadItemsRequest, opts ...grpc.CallOption) (*LoadItemsResponse, error)
	MoveItem(ctx context.Context, in *MoveItemRequest, opts ...grpc.CallOption) (*MoveItemResponse, error)
	SelectItem(ctx context.Context, in *SelectItemRequest, opts ...grpc.CallOption) (*SelectItemResponse, error)
	RemoveItemsByCode(ctx context.Context, in *RemoveItemsRequest, opts ...grpc.CallOption) (*RemoveItemsResponse, error)
	AuditState(ctx context.Context, in *AuditStateRequest, opts ...grpc.CallOption) (*AuditStateResponse, error)
}

type sortingRobotClient struct {
	cc grpc.ClientConnInterface
}

func NewSortingRobotClient(cc grpc.ClientConnInterface) SortingRobotClient {
	return &sortingRobotClient{cc}
}

func (c *sortingRobotClient) LoadItems(ctx context.Context, in *LoadItemsRequest, opts ...grpc.CallOption) (*LoadItemsResponse, error) {
	out := new(LoadItemsResponse)
	err := c.cc.Invoke(ctx, "/SortingRobot/LoadItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sortingRobotClient) MoveItem(ctx context.Context, in *MoveItemRequest, opts ...grpc.CallOption) (*MoveItemResponse, error) {
	out := new(MoveItemResponse)
	err := c.cc.Invoke(ctx, "/SortingRobot/MoveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sortingRobotClient) SelectItem(ctx context.Context, in *SelectItemRequest, opts ...grpc.CallOption) (*SelectItemResponse, error) {
	out := new(SelectItemResponse)
	err := c.cc.Invoke(ctx, "/SortingRobot/SelectItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sortingRobotClient) RemoveItemsByCode(ctx context.Context, in *RemoveItemsRequest, opts ...grpc.CallOption) (*RemoveItemsResponse, error) {
	out := new(RemoveItemsResponse)
	err := c.cc.Invoke(ctx, "/SortingRobot/RemoveItemsByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sortingRobotClient) AuditState(ctx context.Context, in *AuditStateRequest, opts ...grpc.CallOption) (*AuditStateResponse, error) {
	out := new(AuditStateResponse)
	err := c.cc.Invoke(ctx, "/SortingRobot/AuditState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SortingRobotServer is the server API for SortingRobot service.
// All implementations should embed UnimplementedSortingRobotServer
// for forward compatibility
type SortingRobotServer interface {
	LoadItems(context.Context, *LoadItemsRequest) (*LoadItemsResponse, error)
	MoveItem(context.Context, *MoveItemRequest) (*MoveItemResponse, error)
	SelectItem(context.Context, *SelectItemRequest) (*SelectItemResponse, error)
	RemoveItemsByCode(context.Context, *RemoveItemsRequest) (*RemoveItemsResponse, error)
	AuditState(context.Context, *AuditStateRequest) (*AuditStateResponse, error)
}

// UnimplementedSortingRobotServer should be embedded to have forward compatible implementations.
type UnimplementedSortingRobotServer struct {
}

func (UnimplementedSortingRobotServer) LoadItems(context.Context, *LoadItemsRequest) (*LoadItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadItems not implemented")
}
func (UnimplementedSortingRobotServer) MoveItem(context.Context, *MoveItemRequest) (*MoveItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveItem not implemented")
}
func (UnimplementedSortingRobotServer) SelectItem(context.Context, *SelectItemRequest) (*SelectItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectItem not implemented")
}
func (UnimplementedSortingRobotServer) RemoveItemsByCode(context.Context, *RemoveItemsRequest) (*RemoveItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItemsByCode not implemented")
}
func (UnimplementedSortingRobotServer) AuditState(context.Context, *AuditStateRequest) (*AuditStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditState not implemented")
}

// UnsafeSortingRobotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SortingRobotServer will
// result in compilation errors.
type UnsafeSortingRobotServer interface {
	mustEmbedUnimplementedSortingRobotServer()
}

func RegisterSortingRobotServer(s grpc.ServiceRegistrar, srv SortingRobotServer) {
	s.RegisterService(&SortingRobot_ServiceDesc, srv)
}

func _SortingRobot_LoadItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingRobotServer).LoadItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SortingRobot/LoadItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingRobotServer).LoadItems(ctx, req.(*LoadItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SortingRobot_MoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingRobotServer).MoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SortingRobot/MoveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingRobotServer).MoveItem(ctx, req.(*MoveItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SortingRobot_SelectItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingRobotServer).SelectItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SortingRobot/SelectItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingRobotServer).SelectItem(ctx, req.(*SelectItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SortingRobot_RemoveItemsByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingRobotServer).RemoveItemsByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SortingRobot/RemoveItemsByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingRobotServer).RemoveItemsByCode(ctx, req.(*RemoveItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SortingRobot_AuditState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingRobotServer).AuditState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SortingRobot/AuditState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingRobotServer).AuditState(ctx, req.(*AuditStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SortingRobot_ServiceDesc is the grpc.ServiceDesc for SortingRobot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SortingRobot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SortingRobot",
	HandlerType: (*SortingRobotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadItems",
			Handler:    _SortingRobot_LoadItems_Handler,
		},
		{
			MethodName: "MoveItem",
			Handler:    _SortingRobot_MoveItem_Handler,
		},
		{
			MethodName: "SelectItem",
			Handler:    _SortingRobot_SelectItem_Handler,
		},
		{
			MethodName: "RemoveItemsByCode",
			Handler:    _SortingRobot_RemoveItemsByCode_Handler,
		},
		{
			MethodName: "AuditState",
			Handler:    _SortingRobot_AuditState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorting.proto",
}
