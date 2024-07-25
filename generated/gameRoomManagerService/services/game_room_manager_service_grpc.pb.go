// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: services/game_room_manager_service.proto

package gameRoomManagerService

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

// GameRoomManagerClient is the client API for GameRoomManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameRoomManagerClient interface {
	CreateInvite(ctx context.Context, in *InvitationMessage, opts ...grpc.CallOption) (*GameRoomMessage, error)
	AcceptInvite(ctx context.Context, in *InvitationMessage, opts ...grpc.CallOption) (*GameRoomMessage, error)
}

type gameRoomManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGameRoomManagerClient(cc grpc.ClientConnInterface) GameRoomManagerClient {
	return &gameRoomManagerClient{cc}
}

func (c *gameRoomManagerClient) CreateInvite(ctx context.Context, in *InvitationMessage, opts ...grpc.CallOption) (*GameRoomMessage, error) {
	out := new(GameRoomMessage)
	err := c.cc.Invoke(ctx, "/GameRoomManager/CreateInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRoomManagerClient) AcceptInvite(ctx context.Context, in *InvitationMessage, opts ...grpc.CallOption) (*GameRoomMessage, error) {
	out := new(GameRoomMessage)
	err := c.cc.Invoke(ctx, "/GameRoomManager/AcceptInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameRoomManagerServer is the server API for GameRoomManager service.
// All implementations must embed UnimplementedGameRoomManagerServer
// for forward compatibility
type GameRoomManagerServer interface {
	CreateInvite(context.Context, *InvitationMessage) (*GameRoomMessage, error)
	AcceptInvite(context.Context, *InvitationMessage) (*GameRoomMessage, error)
	mustEmbedUnimplementedGameRoomManagerServer()
}

// UnimplementedGameRoomManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGameRoomManagerServer struct {
}

func (UnimplementedGameRoomManagerServer) CreateInvite(context.Context, *InvitationMessage) (*GameRoomMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvite not implemented")
}
func (UnimplementedGameRoomManagerServer) AcceptInvite(context.Context, *InvitationMessage) (*GameRoomMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvite not implemented")
}
func (UnimplementedGameRoomManagerServer) mustEmbedUnimplementedGameRoomManagerServer() {}

// UnsafeGameRoomManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameRoomManagerServer will
// result in compilation errors.
type UnsafeGameRoomManagerServer interface {
	mustEmbedUnimplementedGameRoomManagerServer()
}

func RegisterGameRoomManagerServer(s grpc.ServiceRegistrar, srv GameRoomManagerServer) {
	s.RegisterService(&GameRoomManager_ServiceDesc, srv)
}

func _GameRoomManager_CreateInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRoomManagerServer).CreateInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameRoomManager/CreateInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRoomManagerServer).CreateInvite(ctx, req.(*InvitationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRoomManager_AcceptInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRoomManagerServer).AcceptInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameRoomManager/AcceptInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRoomManagerServer).AcceptInvite(ctx, req.(*InvitationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GameRoomManager_ServiceDesc is the grpc.ServiceDesc for GameRoomManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameRoomManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameRoomManager",
	HandlerType: (*GameRoomManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvite",
			Handler:    _GameRoomManager_CreateInvite_Handler,
		},
		{
			MethodName: "AcceptInvite",
			Handler:    _GameRoomManager_AcceptInvite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/game_room_manager_service.proto",
}
