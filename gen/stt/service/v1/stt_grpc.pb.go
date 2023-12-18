// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: stt/service/v1/stt.proto

// Additional imports go here

package sttservice

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
	SttService_Register_FullMethodName        = "/stt.service.v1.SttService/Register"
	SttService_Login_FullMethodName           = "/stt.service.v1.SttService/Login"
	SttService_TranscribeAudio_FullMethodName = "/stt.service.v1.SttService/TranscribeAudio"
	SttService_GetResult_FullMethodName       = "/stt.service.v1.SttService/GetResult"
	SttService_GetHistory_FullMethodName      = "/stt.service.v1.SttService/GetHistory"
)

// SttServiceClient is the client API for SttService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SttServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	TranscribeAudio(ctx context.Context, in *Audio, opts ...grpc.CallOption) (SttService_TranscribeAudioClient, error)
	GetResult(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	GetHistory(ctx context.Context, in *User, opts ...grpc.CallOption) (*History, error)
}

type sttServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSttServiceClient(cc grpc.ClientConnInterface) SttServiceClient {
	return &sttServiceClient{cc}
}

func (c *sttServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, SttService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sttServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SttService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sttServiceClient) TranscribeAudio(ctx context.Context, in *Audio, opts ...grpc.CallOption) (SttService_TranscribeAudioClient, error) {
	stream, err := c.cc.NewStream(ctx, &SttService_ServiceDesc.Streams[0], SttService_TranscribeAudio_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sttServiceTranscribeAudioClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SttService_TranscribeAudioClient interface {
	Recv() (*Status, error)
	grpc.ClientStream
}

type sttServiceTranscribeAudioClient struct {
	grpc.ClientStream
}

func (x *sttServiceTranscribeAudioClient) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sttServiceClient) GetResult(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, SttService_GetResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sttServiceClient) GetHistory(ctx context.Context, in *User, opts ...grpc.CallOption) (*History, error) {
	out := new(History)
	err := c.cc.Invoke(ctx, SttService_GetHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SttServiceServer is the server API for SttService service.
// All implementations must embed UnimplementedSttServiceServer
// for forward compatibility
type SttServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	TranscribeAudio(*Audio, SttService_TranscribeAudioServer) error
	GetResult(context.Context, *Text) (*Text, error)
	GetHistory(context.Context, *User) (*History, error)
	mustEmbedUnimplementedSttServiceServer()
}

// UnimplementedSttServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSttServiceServer struct {
}

func (UnimplementedSttServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSttServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSttServiceServer) TranscribeAudio(*Audio, SttService_TranscribeAudioServer) error {
	return status.Errorf(codes.Unimplemented, "method TranscribeAudio not implemented")
}
func (UnimplementedSttServiceServer) GetResult(context.Context, *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedSttServiceServer) GetHistory(context.Context, *User) (*History, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedSttServiceServer) mustEmbedUnimplementedSttServiceServer() {}

// UnsafeSttServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SttServiceServer will
// result in compilation errors.
type UnsafeSttServiceServer interface {
	mustEmbedUnimplementedSttServiceServer()
}

func RegisterSttServiceServer(s grpc.ServiceRegistrar, srv SttServiceServer) {
	s.RegisterService(&SttService_ServiceDesc, srv)
}

func _SttService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SttService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SttService_TranscribeAudio_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Audio)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SttServiceServer).TranscribeAudio(m, &sttServiceTranscribeAudioServer{stream})
}

type SttService_TranscribeAudioServer interface {
	Send(*Status) error
	grpc.ServerStream
}

type sttServiceTranscribeAudioServer struct {
	grpc.ServerStream
}

func (x *sttServiceTranscribeAudioServer) Send(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func _SttService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_GetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).GetResult(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _SttService_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_GetHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).GetHistory(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// SttService_ServiceDesc is the grpc.ServiceDesc for SttService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SttService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stt.service.v1.SttService",
	HandlerType: (*SttServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _SttService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SttService_Login_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _SttService_GetResult_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _SttService_GetHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TranscribeAudio",
			Handler:       _SttService_TranscribeAudio_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stt/service/v1/stt.proto",
}
