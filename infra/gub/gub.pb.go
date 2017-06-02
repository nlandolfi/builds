// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/nlandolfi/builds/infra/gub/gub.proto

/*
Package gub is a generated protocol buffer package.

It is generated from these files:
	github.com/nlandolfi/builds/infra/gub/gub.proto

It has these top-level messages:
	JobStatus
*/
package gub

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

type JobStatus struct {
	StartedMillis uint64 `protobuf:"varint,2,opt,name=started_millis,json=startedMillis" json:"started_millis,omitempty"`
	StoppedMillis uint64 `protobuf:"varint,3,opt,name=stopped_millis,json=stoppedMillis" json:"stopped_millis,omitempty"`
}

func (m *JobStatus) Reset()                    { *m = JobStatus{} }
func (m *JobStatus) String() string            { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()               {}
func (*JobStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JobStatus) GetStartedMillis() uint64 {
	if m != nil {
		return m.StartedMillis
	}
	return 0
}

func (m *JobStatus) GetStoppedMillis() uint64 {
	if m != nil {
		return m.StoppedMillis
	}
	return 0
}

func init() {
	proto.RegisterType((*JobStatus)(nil), "gub.JobStatus")
}

func init() { proto.RegisterFile("github.com/nlandolfi/builds/infra/gub/gub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcb, 0x49, 0xcc, 0x4b, 0xc9, 0xcf, 0x49, 0xcb,
	0xd4, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0x29, 0xd6, 0xcf, 0xcc, 0x4b, 0x2b, 0x4a, 0xd4, 0x4f, 0x2f,
	0x4d, 0x02, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xe6, 0xf4, 0xd2, 0x24, 0xa5, 0x48,
	0x2e, 0x4e, 0xaf, 0xfc, 0xa4, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x21, 0x55, 0x2e, 0xbe, 0xe2,
	0x92, 0xc4, 0xa2, 0x92, 0xd4, 0x94, 0xf8, 0xdc, 0xcc, 0x9c, 0x9c, 0xcc, 0x62, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x5e, 0xa8, 0xa8, 0x2f, 0x58, 0x10, 0xa2, 0x2c, 0xbf, 0xa0, 0x00, 0xa1,
	0x8c, 0x19, 0xa6, 0x0c, 0x2c, 0x0a, 0x51, 0xe6, 0xa4, 0x1e, 0xa5, 0x4a, 0x94, 0x93, 0x92, 0xd8,
	0xc0, 0xee, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x78, 0x03, 0x84, 0xcb, 0xc2, 0x00, 0x00,
	0x00,
}
