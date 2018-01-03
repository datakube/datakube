// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dumper.proto

/*
Package dumper is a generated protocol buffer package.

It is generated from these files:
	dumper.proto
	file.proto

It has these top-level messages:
	RegisterRequest
	Target
	RegisterResponse
	SaveDumpFileRequest
	SaveDumpFileResponse
*/
package dumper

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
	Token   string    `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Targets []*Target `protobuf:"bytes,2,rep,name=targets" json:"targets,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterRequest) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

type Target struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Schedule string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

type RegisterResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "dumper.RegisterRequest")
	proto.RegisterType((*Target)(nil), "dumper.Target")
	proto.RegisterType((*RegisterResponse)(nil), "dumper.RegisterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DumperService service

type DumperServiceClient interface {
	RegisterDumper(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type dumperServiceClient struct {
	cc *grpc.ClientConn
}

func NewDumperServiceClient(cc *grpc.ClientConn) DumperServiceClient {
	return &dumperServiceClient{cc}
}

func (c *dumperServiceClient) RegisterDumper(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/dumper.DumperService/RegisterDumper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DumperService service

type DumperServiceServer interface {
	RegisterDumper(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterDumperServiceServer(s *grpc.Server, srv DumperServiceServer) {
	s.RegisterService(&_DumperService_serviceDesc, srv)
}

func _DumperService_RegisterDumper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DumperServiceServer).RegisterDumper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dumper.DumperService/RegisterDumper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DumperServiceServer).RegisterDumper(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DumperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dumper.DumperService",
	HandlerType: (*DumperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDumper",
			Handler:    _DumperService_RegisterDumper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dumper.proto",
}

func init() { proto.RegisterFile("dumper.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4e, 0x86, 0x30,
	0x10, 0x84, 0xfd, 0x7f, 0x15, 0x70, 0x55, 0x34, 0x1b, 0x13, 0x1b, 0x4e, 0xa4, 0x27, 0x0e, 0x86,
	0x03, 0x5e, 0x7c, 0x00, 0x7d, 0x00, 0xab, 0xf1, 0x8e, 0x65, 0x83, 0x44, 0xa1, 0xd8, 0x6d, 0x7d,
	0x7e, 0x93, 0x96, 0x6a, 0xa2, 0xb7, 0x7e, 0x33, 0xe9, 0xec, 0xec, 0xc2, 0xd9, 0xe0, 0xe7, 0x95,
	0x6c, 0xbb, 0x5a, 0xe3, 0x0c, 0x66, 0x91, 0xe4, 0x23, 0x5c, 0x28, 0x1a, 0x27, 0x76, 0x64, 0x15,
	0x7d, 0x7a, 0x62, 0x87, 0x57, 0x70, 0xec, 0xcc, 0x3b, 0x2d, 0x62, 0x57, 0xef, 0x9a, 0x13, 0x15,
	0x01, 0x1b, 0xc8, 0x5d, 0x6f, 0x47, 0x72, 0x2c, 0xf6, 0xf5, 0x61, 0x73, 0xda, 0x95, 0xed, 0x16,
	0xf8, 0x1c, 0x64, 0x95, 0x6c, 0x79, 0x07, 0x59, 0x94, 0x10, 0xe1, 0x68, 0xe9, 0x67, 0xda, 0x82,
	0xc2, 0x1b, 0x2b, 0x28, 0x58, 0xbf, 0xd1, 0xe0, 0x3f, 0x48, 0xec, 0x83, 0xfe, 0xc3, 0xf2, 0x06,
	0x2e, 0x7f, 0xcb, 0xf0, 0x6a, 0x16, 0x26, 0x14, 0x90, 0xb3, 0xd7, 0x9a, 0x98, 0x43, 0x4c, 0xa1,
	0x12, 0x76, 0x2f, 0x70, 0x7e, 0x1f, 0x1a, 0x3c, 0x91, 0xfd, 0x9a, 0x34, 0xe1, 0x03, 0x94, 0xe9,
	0x7b, 0x34, 0xf0, 0x3a, 0x75, 0xfc, 0xb3, 0x63, 0x25, 0xfe, 0x1b, 0x71, 0x9e, 0x3c, 0x78, 0xcd,
	0xc2, 0x85, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xc0, 0x83, 0x54, 0x31, 0x01, 0x00,
	0x00,
}