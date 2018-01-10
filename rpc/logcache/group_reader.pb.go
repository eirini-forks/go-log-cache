// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group_reader.proto

package logcache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AddToGroupRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SourceId string `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
}

func (m *AddToGroupRequest) Reset()                    { *m = AddToGroupRequest{} }
func (m *AddToGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*AddToGroupRequest) ProtoMessage()               {}
func (*AddToGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AddToGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddToGroupRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

type AddToGroupResponse struct {
}

func (m *AddToGroupResponse) Reset()                    { *m = AddToGroupResponse{} }
func (m *AddToGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*AddToGroupResponse) ProtoMessage()               {}
func (*AddToGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type RemoveFromGroupRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SourceId string `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
}

func (m *RemoveFromGroupRequest) Reset()                    { *m = RemoveFromGroupRequest{} }
func (m *RemoveFromGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveFromGroupRequest) ProtoMessage()               {}
func (*RemoveFromGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RemoveFromGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoveFromGroupRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

type RemoveFromGroupResponse struct {
}

func (m *RemoveFromGroupResponse) Reset()                    { *m = RemoveFromGroupResponse{} }
func (m *RemoveFromGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveFromGroupResponse) ProtoMessage()               {}
func (*RemoveFromGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type GroupReadRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RequesterId  uint64        `protobuf:"varint,2,opt,name=requester_id,json=requesterId" json:"requester_id,omitempty"`
	StartTime    int64         `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime      int64         `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Limit        int64         `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	EnvelopeType EnvelopeTypes `protobuf:"varint,6,opt,name=envelope_type,json=envelopeType,enum=logcache.EnvelopeTypes" json:"envelope_type,omitempty"`
}

func (m *GroupReadRequest) Reset()                    { *m = GroupReadRequest{} }
func (m *GroupReadRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupReadRequest) ProtoMessage()               {}
func (*GroupReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GroupReadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupReadRequest) GetRequesterId() uint64 {
	if m != nil {
		return m.RequesterId
	}
	return 0
}

func (m *GroupReadRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *GroupReadRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *GroupReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GroupReadRequest) GetEnvelopeType() EnvelopeTypes {
	if m != nil {
		return m.EnvelopeType
	}
	return EnvelopeTypes_ANY
}

type GroupReadResponse struct {
	Envelopes *loggregator_v2.EnvelopeBatch `protobuf:"bytes,1,opt,name=envelopes" json:"envelopes,omitempty"`
}

func (m *GroupReadResponse) Reset()                    { *m = GroupReadResponse{} }
func (m *GroupReadResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupReadResponse) ProtoMessage()               {}
func (*GroupReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GroupReadResponse) GetEnvelopes() *loggregator_v2.EnvelopeBatch {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

type GroupRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GroupRequest) Reset()                    { *m = GroupRequest{} }
func (m *GroupRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupRequest) ProtoMessage()               {}
func (*GroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GroupResponse struct {
	SourceIds    []string `protobuf:"bytes,1,rep,name=source_ids,json=sourceIds" json:"source_ids,omitempty"`
	RequesterIds []uint64 `protobuf:"varint,2,rep,packed,name=requester_ids,json=requesterIds" json:"requester_ids,omitempty"`
}

func (m *GroupResponse) Reset()                    { *m = GroupResponse{} }
func (m *GroupResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupResponse) ProtoMessage()               {}
func (*GroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *GroupResponse) GetSourceIds() []string {
	if m != nil {
		return m.SourceIds
	}
	return nil
}

func (m *GroupResponse) GetRequesterIds() []uint64 {
	if m != nil {
		return m.RequesterIds
	}
	return nil
}

func init() {
	proto.RegisterType((*AddToGroupRequest)(nil), "logcache.AddToGroupRequest")
	proto.RegisterType((*AddToGroupResponse)(nil), "logcache.AddToGroupResponse")
	proto.RegisterType((*RemoveFromGroupRequest)(nil), "logcache.RemoveFromGroupRequest")
	proto.RegisterType((*RemoveFromGroupResponse)(nil), "logcache.RemoveFromGroupResponse")
	proto.RegisterType((*GroupReadRequest)(nil), "logcache.GroupReadRequest")
	proto.RegisterType((*GroupReadResponse)(nil), "logcache.GroupReadResponse")
	proto.RegisterType((*GroupRequest)(nil), "logcache.GroupRequest")
	proto.RegisterType((*GroupResponse)(nil), "logcache.GroupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GroupReader service

type GroupReaderClient interface {
	AddToGroup(ctx context.Context, in *AddToGroupRequest, opts ...grpc.CallOption) (*AddToGroupResponse, error)
	RemoveFromGroup(ctx context.Context, in *RemoveFromGroupRequest, opts ...grpc.CallOption) (*RemoveFromGroupResponse, error)
	Read(ctx context.Context, in *GroupReadRequest, opts ...grpc.CallOption) (*GroupReadResponse, error)
	Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
}

type groupReaderClient struct {
	cc *grpc.ClientConn
}

func NewGroupReaderClient(cc *grpc.ClientConn) GroupReaderClient {
	return &groupReaderClient{cc}
}

func (c *groupReaderClient) AddToGroup(ctx context.Context, in *AddToGroupRequest, opts ...grpc.CallOption) (*AddToGroupResponse, error) {
	out := new(AddToGroupResponse)
	err := grpc.Invoke(ctx, "/logcache.GroupReader/AddToGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupReaderClient) RemoveFromGroup(ctx context.Context, in *RemoveFromGroupRequest, opts ...grpc.CallOption) (*RemoveFromGroupResponse, error) {
	out := new(RemoveFromGroupResponse)
	err := grpc.Invoke(ctx, "/logcache.GroupReader/RemoveFromGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupReaderClient) Read(ctx context.Context, in *GroupReadRequest, opts ...grpc.CallOption) (*GroupReadResponse, error) {
	out := new(GroupReadResponse)
	err := grpc.Invoke(ctx, "/logcache.GroupReader/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupReaderClient) Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := grpc.Invoke(ctx, "/logcache.GroupReader/Group", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupReader service

type GroupReaderServer interface {
	AddToGroup(context.Context, *AddToGroupRequest) (*AddToGroupResponse, error)
	RemoveFromGroup(context.Context, *RemoveFromGroupRequest) (*RemoveFromGroupResponse, error)
	Read(context.Context, *GroupReadRequest) (*GroupReadResponse, error)
	Group(context.Context, *GroupRequest) (*GroupResponse, error)
}

func RegisterGroupReaderServer(s *grpc.Server, srv GroupReaderServer) {
	s.RegisterService(&_GroupReader_serviceDesc, srv)
}

func _GroupReader_AddToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupReaderServer).AddToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.GroupReader/AddToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupReaderServer).AddToGroup(ctx, req.(*AddToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupReader_RemoveFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupReaderServer).RemoveFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.GroupReader/RemoveFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupReaderServer).RemoveFromGroup(ctx, req.(*RemoveFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupReader_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupReaderServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.GroupReader/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupReaderServer).Read(ctx, req.(*GroupReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupReader_Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupReaderServer).Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.GroupReader/Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupReaderServer).Group(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupReader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.GroupReader",
	HandlerType: (*GroupReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToGroup",
			Handler:    _GroupReader_AddToGroup_Handler,
		},
		{
			MethodName: "RemoveFromGroup",
			Handler:    _GroupReader_RemoveFromGroup_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _GroupReader_Read_Handler,
		},
		{
			MethodName: "Group",
			Handler:    _GroupReader_Group_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group_reader.proto",
}

func init() { proto.RegisterFile("group_reader.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0x55, 0x7e, 0x43, 0x5c, 0x49, 0x60, 0x52, 0x1a, 0x26, 0x1e, 0x27, 0x41, 0x19, 0xc3, 0x22,
	0x62, 0x11, 0x8b, 0xb0, 0x84, 0x0d, 0x88, 0x8f, 0xb2, 0x43, 0x26, 0x12, 0x62, 0x15, 0x35, 0x71,
	0xc9, 0x63, 0x29, 0x76, 0x9b, 0xee, 0x4e, 0xa4, 0xd1, 0x30, 0x1b, 0xae, 0xc0, 0x3d, 0xb8, 0x0c,
	0x0b, 0x2e, 0xc0, 0x41, 0x50, 0xda, 0x76, 0xec, 0x89, 0xa3, 0xb0, 0x60, 0xe7, 0x7e, 0xaf, 0xfc,
	0x5e, 0xbf, 0xaa, 0xb2, 0x01, 0x7d, 0xc1, 0xd7, 0xf1, 0x42, 0x10, 0xf3, 0x48, 0x4c, 0x62, 0xc1,
	0x15, 0xc7, 0xe6, 0x8a, 0xfb, 0x4b, 0xb6, 0xbc, 0x22, 0x6b, 0xe0, 0x73, 0xee, 0xaf, 0xc8, 0x61,
	0x71, 0xe0, 0xb0, 0x28, 0xe2, 0x8a, 0xa9, 0x80, 0x47, 0x32, 0xa9, 0xb3, 0xba, 0x9b, 0xa9, 0x43,
	0xd1, 0x86, 0x56, 0x3c, 0xa6, 0x14, 0x6a, 0x93, 0x2f, 0x48, 0xa6, 0x05, 0xf6, 0x1b, 0xe8, 0xbe,
	0xf2, 0xbc, 0x39, 0x7f, 0xbf, 0xf5, 0x70, 0xe9, 0xeb, 0x9a, 0xa4, 0x42, 0x84, 0x7a, 0xc4, 0x42,
	0x32, 0x2b, 0xa3, 0xca, 0xd8, 0x70, 0xf5, 0x33, 0xf6, 0xc1, 0x90, 0x7c, 0x2d, 0x96, 0xb4, 0x08,
	0x3c, 0xb3, 0xaa, 0x89, 0x66, 0x02, 0xcc, 0x3c, 0xfb, 0x0c, 0xb0, 0xa8, 0x22, 0x63, 0x1e, 0x49,
	0xb2, 0x67, 0x70, 0xee, 0x52, 0xc8, 0x37, 0xf4, 0x4e, 0xf0, 0xf0, 0xff, 0x0c, 0x2e, 0xa0, 0x57,
	0x92, 0x4a, 0x5d, 0x7e, 0x57, 0xe0, 0x34, 0x45, 0x98, 0x77, 0xcc, 0xe0, 0x12, 0xda, 0x22, 0xa1,
	0x49, 0x64, 0x1e, 0x75, 0xb7, 0xb5, 0xc3, 0x66, 0x1e, 0x0e, 0x01, 0xa4, 0x62, 0x42, 0x2d, 0x54,
	0x10, 0x92, 0x59, 0x1b, 0x55, 0xc6, 0x35, 0xd7, 0xd0, 0xc8, 0x3c, 0x08, 0x09, 0x2f, 0xa0, 0x49,
	0x91, 0x97, 0x90, 0x75, 0x4d, 0xde, 0xa3, 0xc8, 0xd3, 0xd4, 0x19, 0x34, 0x56, 0x41, 0x18, 0x28,
	0xb3, 0xa1, 0xf1, 0xe4, 0x80, 0x2f, 0xa1, 0x93, 0x75, 0x7f, 0xa1, 0xae, 0x63, 0x32, 0x4f, 0x46,
	0x95, 0xf1, 0xfd, 0x69, 0x6f, 0x92, 0x8d, 0x6f, 0xf2, 0x36, 0xa5, 0xe7, 0xd7, 0x31, 0x49, 0xb7,
	0x4d, 0x85, 0xa3, 0xfd, 0x01, 0xba, 0x85, 0x60, 0x49, 0x5c, 0x7c, 0x01, 0x46, 0x56, 0x24, 0x75,
	0xbc, 0xd6, 0x74, 0xb8, 0x95, 0xf3, 0x05, 0xf9, 0x4c, 0x71, 0x31, 0xd9, 0x4c, 0x77, 0xa2, 0xaf,
	0x99, 0x5a, 0x5e, 0xb9, 0x79, 0xbd, 0x6d, 0x43, 0xfb, 0x5f, 0x73, 0xb0, 0x3f, 0x42, 0xe7, 0x4e,
	0x83, 0x75, 0x53, 0xb2, 0xc1, 0x6c, 0x2d, 0x6b, 0x63, 0xc3, 0x35, 0xb2, 0xc9, 0x48, 0x7c, 0x0c,
	0x9d, 0x62, 0x5b, 0xa5, 0x59, 0x1d, 0xd5, 0xc6, 0x75, 0xb7, 0x5d, 0xe8, 0xab, 0x9c, 0xfe, 0xac,
	0x41, 0x6b, 0x97, 0x85, 0x04, 0x86, 0x00, 0xf9, 0xc2, 0x60, 0x3f, 0xef, 0x47, 0x69, 0x19, 0xad,
	0xc1, 0x61, 0x32, 0x9d, 0xfe, 0x93, 0xef, 0xbf, 0xfe, 0xfc, 0xa8, 0x3e, 0xb2, 0x06, 0xce, 0xe6,
	0x99, 0xa3, 0x3f, 0x14, 0xe7, 0x66, 0x1b, 0xe3, 0xd6, 0xb9, 0xd9, 0x5d, 0xfa, 0x16, 0xbf, 0xc1,
	0x83, 0xbd, 0xf5, 0xc1, 0x51, 0x2e, 0x7b, 0x78, 0x49, 0xad, 0xcb, 0x23, 0x15, 0x77, 0xdd, 0x9f,
	0x1e, 0x77, 0xff, 0x0c, 0xf5, 0x6d, 0x6c, 0xb4, 0x72, 0xc1, 0xfd, 0x85, 0xb5, 0xfa, 0x07, 0xb9,
	0xd4, 0xc6, 0xd4, 0x36, 0x88, 0xa7, 0xfb, 0x36, 0xf8, 0x09, 0x1a, 0x49, 0x9c, 0xf3, 0xd2, 0xfb,
	0x89, 0x6e, 0xaf, 0x84, 0xa7, 0x9a, 0x43, 0xad, 0xd9, 0xc3, 0x87, 0xa5, 0xab, 0x87, 0xa4, 0xd8,
	0x97, 0x13, 0xfd, 0x7b, 0x78, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x95, 0xc5, 0x66, 0x7d,
	0x04, 0x00, 0x00,
}
