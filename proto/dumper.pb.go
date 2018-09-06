// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dumper.proto

/*
Package dumper is a generated protocol buffer package.

It is generated from these files:
	dumper.proto

It has these top-level messages:
	Authorization
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Authorization struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Authorization) Reset()                    { *m = Authorization{} }
func (m *Authorization) String() string            { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()               {}
func (*Authorization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Authorization) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterRequest struct {
	Auth    *Authorization `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Targets []*Target      `protobuf:"bytes,2,rep,name=targets" json:"targets,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetAuth() *Authorization {
	if m != nil {
		return m.Auth
	}
	return nil
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
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SaveDumpFileRequest struct {
	Auth       *Authorization `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Targetname string         `protobuf:"bytes,2,opt,name=targetname" json:"targetname,omitempty"`
	Filename   string         `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Data       []byte         `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SaveDumpFileRequest) Reset()                    { *m = SaveDumpFileRequest{} }
func (m *SaveDumpFileRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveDumpFileRequest) ProtoMessage()               {}
func (*SaveDumpFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SaveDumpFileRequest) GetAuth() *Authorization {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *SaveDumpFileRequest) GetTargetname() string {
	if m != nil {
		return m.Targetname
	}
	return ""
}

func (m *SaveDumpFileRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *SaveDumpFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SaveDumpFileResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *SaveDumpFileResponse) Reset()                    { *m = SaveDumpFileResponse{} }
func (m *SaveDumpFileResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveDumpFileResponse) ProtoMessage()               {}
func (*SaveDumpFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SaveDumpFileResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Authorization)(nil), "dumper.Authorization")
	proto.RegisterType((*RegisterRequest)(nil), "dumper.RegisterRequest")
	proto.RegisterType((*Target)(nil), "dumper.Target")
	proto.RegisterType((*RegisterResponse)(nil), "dumper.RegisterResponse")
	proto.RegisterType((*SaveDumpFileRequest)(nil), "dumper.SaveDumpFileRequest")
	proto.RegisterType((*SaveDumpFileResponse)(nil), "dumper.SaveDumpFileResponse")
}

func init() { proto.RegisterFile("dumper.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x1d, 0x40, 0xc0, 0xcb, 0x8f, 0xa6, 0x62, 0x9c, 0xa0, 0x31, 0xa4, 0x89, 0xc9, 0x98,
	0x18, 0x62, 0xc6, 0x8d, 0x5b, 0x13, 0x74, 0xe3, 0xae, 0x18, 0x17, 0xee, 0xea, 0x70, 0x81, 0x46,
	0x98, 0x8e, 0xfd, 0x61, 0xe1, 0x53, 0xf8, 0xc8, 0x86, 0x96, 0x2a, 0x28, 0x89, 0x89, 0xbb, 0x9e,
	0x7b, 0xee, 0x9c, 0xf9, 0x4e, 0x67, 0xa0, 0x39, 0xb2, 0xf3, 0x02, 0x55, 0xbf, 0x50, 0xd2, 0x48,
	0x52, 0xf5, 0x8a, 0x9e, 0x43, 0xeb, 0xd6, 0x9a, 0xa9, 0x54, 0xe2, 0x9d, 0x1b, 0x21, 0x73, 0xd2,
	0x81, 0x5d, 0x23, 0x5f, 0x31, 0x8f, 0xa3, 0x5e, 0x94, 0xec, 0x31, 0x2f, 0xe8, 0x18, 0xf6, 0x19,
	0x4e, 0x84, 0x36, 0xa8, 0x18, 0xbe, 0x59, 0xd4, 0x86, 0x5c, 0x40, 0x85, 0x5b, 0x33, 0x75, 0x7b,
	0x8d, 0xf4, 0xa8, 0xbf, 0x8a, 0xdf, 0x48, 0x63, 0x6e, 0x85, 0x24, 0x50, 0x33, 0x5c, 0x4d, 0xd0,
	0xe8, 0xb8, 0xd4, 0x2b, 0x27, 0x8d, 0xb4, 0x1d, 0xb6, 0x1f, 0xdd, 0x98, 0x05, 0x9b, 0xde, 0x40,
	0xd5, 0x8f, 0x08, 0x81, 0x4a, 0xce, 0xe7, 0xb8, 0xc2, 0x70, 0x67, 0xd2, 0x85, 0xba, 0xce, 0xa6,
	0x38, 0xb2, 0x33, 0x8c, 0x4b, 0x6e, 0xfe, 0xa5, 0xe9, 0x25, 0x1c, 0x7c, 0x13, 0xea, 0x42, 0xe6,
	0x1a, 0x49, 0x0c, 0x35, 0x6d, 0xb3, 0x0c, 0xb5, 0x76, 0x31, 0x75, 0x16, 0x24, 0xfd, 0x88, 0xe0,
	0x70, 0xc8, 0x17, 0x38, 0xb0, 0xf3, 0xe2, 0x5e, 0xcc, 0xf0, 0x1f, 0xa5, 0xce, 0x00, 0x3c, 0xb5,
	0xc3, 0xf4, 0x38, 0x6b, 0x93, 0x25, 0xec, 0x58, 0xcc, 0xd0, 0xb9, 0x65, 0x0f, 0x1b, 0xf4, 0xb2,
	0xdc, 0x88, 0x1b, 0x1e, 0x57, 0x7a, 0x51, 0xd2, 0x64, 0xee, 0x4c, 0xaf, 0xa0, 0xb3, 0x49, 0xf4,
	0x57, 0x89, 0xf4, 0x09, 0x5a, 0x03, 0xc7, 0x37, 0x44, 0xb5, 0x10, 0x19, 0x92, 0x3b, 0x68, 0x87,
	0x3b, 0xf0, 0x06, 0x39, 0x0e, 0x0d, 0x7e, 0x7c, 0xbd, 0x6e, 0xfc, 0xdb, 0xf0, 0xef, 0xa3, 0x3b,
	0xe9, 0x33, 0x34, 0x96, 0x04, 0x21, 0xf5, 0x01, 0x9a, 0xeb, 0x60, 0xe4, 0x24, 0x3c, 0xba, 0xe5,
	0x02, 0xbb, 0xa7, 0xdb, 0xcd, 0x90, 0xfd, 0x52, 0x75, 0xbf, 0xdf, 0xf5, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x4f, 0x3c, 0x6a, 0x8e, 0x02, 0x00, 0x00,
}