// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jobs "github.com/pydio/cells/v4/common/proto/jobs"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserJobRequest struct {
	// Name of the job to create in the user space
	JobName string `protobuf:"bytes,1,opt,name=JobName" json:"JobName,omitempty"`
	// Json-encoded parameters for this job
	JsonParameters string `protobuf:"bytes,2,opt,name=JsonParameters" json:"JsonParameters,omitempty"`
}

func (m *UserJobRequest) Reset()                    { *m = UserJobRequest{} }
func (m *UserJobRequest) String() string            { return proto.CompactTextString(m) }
func (*UserJobRequest) ProtoMessage()               {}
func (*UserJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *UserJobRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *UserJobRequest) GetJsonParameters() string {
	if m != nil {
		return m.JsonParameters
	}
	return ""
}

type UserJobResponse struct {
	JobUuid string `protobuf:"bytes,1,opt,name=JobUuid" json:"JobUuid,omitempty"`
}

func (m *UserJobResponse) Reset()                    { *m = UserJobResponse{} }
func (m *UserJobResponse) String() string            { return proto.CompactTextString(m) }
func (*UserJobResponse) ProtoMessage()               {}
func (*UserJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *UserJobResponse) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

type UserJobsCollection struct {
	Jobs []*jobs.Job `protobuf:"bytes,1,rep,name=Jobs" json:"Jobs,omitempty"`
}

func (m *UserJobsCollection) Reset()                    { *m = UserJobsCollection{} }
func (m *UserJobsCollection) String() string            { return proto.CompactTextString(m) }
func (*UserJobsCollection) ProtoMessage()               {}
func (*UserJobsCollection) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *UserJobsCollection) GetJobs() []*jobs.Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func init() {
	proto.RegisterType((*UserJobRequest)(nil), "rest.UserJobRequest")
	proto.RegisterType((*UserJobResponse)(nil), "rest.UserJobResponse")
	proto.RegisterType((*UserJobsCollection)(nil), "rest.UserJobsCollection")
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xc1, 0x4a, 0xc5, 0x30,
	0x10, 0x45, 0xa9, 0x3e, 0x94, 0x17, 0xa1, 0x85, 0xac, 0x82, 0x20, 0x94, 0x2e, 0xa4, 0x20, 0x34,
	0x60, 0x3f, 0xc1, 0x5d, 0x17, 0x22, 0x85, 0x7e, 0x40, 0x93, 0x0e, 0x36, 0x92, 0x64, 0x6a, 0x26,
	0x59, 0xf8, 0xf7, 0xd2, 0x58, 0x45, 0xde, 0x66, 0xe0, 0x5e, 0xce, 0x1c, 0x2e, 0xab, 0x48, 0xaf,
	0xb0, 0x24, 0x0b, 0xa1, 0xdb, 0x02, 0x46, 0xe4, 0xa7, 0x00, 0x14, 0xef, 0xfb, 0x77, 0x13, 0xd7,
	0xa4, 0x3a, 0x8d, 0x4e, 0x6e, 0x5f, 0x8b, 0x41, 0xa9, 0xc1, 0x5a, 0x92, 0x1a, 0x9d, 0x43, 0x2f,
	0x33, 0x2a, 0x3f, 0x50, 0x51, 0x3e, 0x3f, 0xaf, 0xcd, 0xc8, 0xca, 0x89, 0x20, 0x0c, 0xa8, 0x46,
	0xf8, 0x4c, 0x40, 0x91, 0x0b, 0x76, 0x3b, 0xa0, 0x7a, 0x9d, 0x1d, 0x88, 0xa2, 0x2e, 0xda, 0xf3,
	0xf8, 0x1b, 0xf9, 0x23, 0x2b, 0x07, 0x42, 0xff, 0x36, 0x87, 0xd9, 0x41, 0x84, 0x40, 0xe2, 0x2a,
	0x03, 0x17, 0x6d, 0xf3, 0xc4, 0xaa, 0x3f, 0x27, 0x6d, 0xe8, 0x09, 0x0e, 0xe9, 0x94, 0xcc, 0xf2,
	0x4f, 0xba, 0xc7, 0xa6, 0x67, 0xfc, 0x80, 0xe9, 0x05, 0xad, 0x05, 0x1d, 0x0d, 0x7a, 0xfe, 0xc0,
	0x4e, 0x7b, 0x23, 0x8a, 0xfa, 0xba, 0xbd, 0x7b, 0x3e, 0x77, 0x79, 0xf1, 0x2e, 0xcc, 0xb5, 0xba,
	0xc9, 0xe3, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xb2, 0xcf, 0x09, 0x0a, 0x01, 0x00,
	0x00,
}
