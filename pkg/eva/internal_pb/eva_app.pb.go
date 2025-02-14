// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eva_app.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IPApplicationLookupInfo struct {
	IpAddress            string   `protobuf:"bytes,1,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPApplicationLookupInfo) Reset()         { *m = IPApplicationLookupInfo{} }
func (m *IPApplicationLookupInfo) String() string { return proto.CompactTextString(m) }
func (*IPApplicationLookupInfo) ProtoMessage()    {}
func (*IPApplicationLookupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0306c89ff354b939, []int{0}
}

func (m *IPApplicationLookupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPApplicationLookupInfo.Unmarshal(m, b)
}
func (m *IPApplicationLookupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPApplicationLookupInfo.Marshal(b, m, deterministic)
}
func (m *IPApplicationLookupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPApplicationLookupInfo.Merge(m, src)
}
func (m *IPApplicationLookupInfo) XXX_Size() int {
	return xxx_messageInfo_IPApplicationLookupInfo.Size(m)
}
func (m *IPApplicationLookupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IPApplicationLookupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IPApplicationLookupInfo proto.InternalMessageInfo

func (m *IPApplicationLookupInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type IPApplicationLookupResult struct {
	AppID                string   `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPApplicationLookupResult) Reset()         { *m = IPApplicationLookupResult{} }
func (m *IPApplicationLookupResult) String() string { return proto.CompactTextString(m) }
func (*IPApplicationLookupResult) ProtoMessage()    {}
func (*IPApplicationLookupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0306c89ff354b939, []int{1}
}

func (m *IPApplicationLookupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPApplicationLookupResult.Unmarshal(m, b)
}
func (m *IPApplicationLookupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPApplicationLookupResult.Marshal(b, m, deterministic)
}
func (m *IPApplicationLookupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPApplicationLookupResult.Merge(m, src)
}
func (m *IPApplicationLookupResult) XXX_Size() int {
	return xxx_messageInfo_IPApplicationLookupResult.Size(m)
}
func (m *IPApplicationLookupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_IPApplicationLookupResult.DiscardUnknown(m)
}

var xxx_messageInfo_IPApplicationLookupResult proto.InternalMessageInfo

func (m *IPApplicationLookupResult) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func init() {
	proto.RegisterType((*IPApplicationLookupInfo)(nil), "pb.IPApplicationLookupInfo")
	proto.RegisterType((*IPApplicationLookupResult)(nil), "pb.IPApplicationLookupResult")
}

func init() { proto.RegisterFile("eva_app.proto", fileDescriptor_0306c89ff354b939) }

var fileDescriptor_0306c89ff354b939 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2d, 0x4b, 0x8c,
	0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe7,
	0x12, 0xf7, 0x0c, 0x70, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xf3, 0xc9,
	0xcf, 0xcf, 0x2e, 0x2d, 0xf0, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xe1, 0xe2, 0xcc, 0x2c, 0x70, 0x4c,
	0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x28, 0x19,
	0x72, 0x49, 0x62, 0xd1, 0x18, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x58,
	0x50, 0xe0, 0xe9, 0x02, 0xd5, 0x06, 0xe1, 0x18, 0x15, 0x70, 0x49, 0x61, 0xd1, 0x12, 0x9c, 0x5a,
	0x54, 0x96, 0x99, 0x9c, 0x2a, 0x14, 0xc4, 0x25, 0xe4, 0x9e, 0x5a, 0x82, 0x24, 0xed, 0x54, 0xe9,
	0x19, 0x20, 0x24, 0xad, 0x57, 0x90, 0xa4, 0x87, 0xc3, 0x85, 0x52, 0xb2, 0x38, 0x24, 0x21, 0xae,
	0x50, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x7b, 0x47,
	0x54, 0xf9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IPApplicationLookupServiceClient is the client API for IPApplicationLookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IPApplicationLookupServiceClient interface {
	GetApplicationByIP(ctx context.Context, in *IPApplicationLookupInfo, opts ...grpc.CallOption) (*IPApplicationLookupResult, error)
}

type iPApplicationLookupServiceClient struct {
	cc *grpc.ClientConn
}

func NewIPApplicationLookupServiceClient(cc *grpc.ClientConn) IPApplicationLookupServiceClient {
	return &iPApplicationLookupServiceClient{cc}
}

func (c *iPApplicationLookupServiceClient) GetApplicationByIP(ctx context.Context, in *IPApplicationLookupInfo, opts ...grpc.CallOption) (*IPApplicationLookupResult, error) {
	out := new(IPApplicationLookupResult)
	err := c.cc.Invoke(ctx, "/pb.IPApplicationLookupService/GetApplicationByIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPApplicationLookupServiceServer is the server API for IPApplicationLookupService service.
type IPApplicationLookupServiceServer interface {
	GetApplicationByIP(context.Context, *IPApplicationLookupInfo) (*IPApplicationLookupResult, error)
}

// UnimplementedIPApplicationLookupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIPApplicationLookupServiceServer struct {
}

func (*UnimplementedIPApplicationLookupServiceServer) GetApplicationByIP(ctx context.Context, req *IPApplicationLookupInfo) (*IPApplicationLookupResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationByIP not implemented")
}

func RegisterIPApplicationLookupServiceServer(s *grpc.Server, srv IPApplicationLookupServiceServer) {
	s.RegisterService(&_IPApplicationLookupService_serviceDesc, srv)
}

func _IPApplicationLookupService_GetApplicationByIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPApplicationLookupInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPApplicationLookupServiceServer).GetApplicationByIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPApplicationLookupService/GetApplicationByIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPApplicationLookupServiceServer).GetApplicationByIP(ctx, req.(*IPApplicationLookupInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPApplicationLookupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPApplicationLookupService",
	HandlerType: (*IPApplicationLookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApplicationByIP",
			Handler:    _IPApplicationLookupService_GetApplicationByIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eva_app.proto",
}
