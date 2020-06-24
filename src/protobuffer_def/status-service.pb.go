// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status-service.proto

package protobuffer_def // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type CMD int32

const (
	CMD_REGISTER_STATUS CMD = 0
	CMD_QUERY_STATUS    CMD = 1
)

var CMD_name = map[int32]string{
	0: "REGISTER_STATUS",
	1: "QUERY_STATUS",
}
var CMD_value = map[string]int32{
	"REGISTER_STATUS": 0,
	"QUERY_STATUS":    1,
}

func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}
func (CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{0}
}

type ReturnCode int32

const (
	ReturnCode_SUCCESS               ReturnCode = 0
	ReturnCode_UNKOWN_CMD            ReturnCode = 1
	ReturnCode_BODY_IS_NULL          ReturnCode = 2
	ReturnCode_DESERIALIZATION_ERROR ReturnCode = 3
	ReturnCode_SERIALIZATION_ERROR   ReturnCode = 4
	ReturnCode_UNKOWN_ERROR          ReturnCode = 5
)

var ReturnCode_name = map[int32]string{
	0: "SUCCESS",
	1: "UNKOWN_CMD",
	2: "BODY_IS_NULL",
	3: "DESERIALIZATION_ERROR",
	4: "SERIALIZATION_ERROR",
	5: "UNKOWN_ERROR",
}
var ReturnCode_value = map[string]int32{
	"SUCCESS":               0,
	"UNKOWN_CMD":            1,
	"BODY_IS_NULL":          2,
	"DESERIALIZATION_ERROR": 3,
	"SERIALIZATION_ERROR":   4,
	"UNKOWN_ERROR":          5,
}

func (x ReturnCode) String() string {
	return proto.EnumName(ReturnCode_name, int32(x))
}
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{1}
}

type BaseRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	C                    CMD      `protobuf:"varint,2,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{0}
}
func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (dst *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(dst, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseRequest) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_REGISTER_STATUS
}

func (m *BaseRequest) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

type BaseResponse struct {
	RequestId            string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Code                 ReturnCode `protobuf:"varint,2,opt,name=code,proto3,enum=ReturnCode" json:"code,omitempty"`
	Desc                 string     `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	C                    CMD        `protobuf:"varint,4,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{1}
}
func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (dst *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(dst, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseResponse) GetCode() ReturnCode {
	if m != nil {
		return m.Code
	}
	return ReturnCode_SUCCESS
}

func (m *BaseResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *BaseResponse) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_REGISTER_STATUS
}

func (m *BaseResponse) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

// 注册状态       请求
type RegisterStatusRequest struct {
	Identity              string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	DeviceType            []byte   `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	NextHeartbeatInterval int32    `protobuf:"varint,3,opt,name=next_heartbeat_interval,json=nextHeartbeatInterval,proto3" json:"next_heartbeat_interval,omitempty"`
	RegisterInfo          string   `protobuf:"bytes,4,opt,name=register_info,json=registerInfo,proto3" json:"register_info,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *RegisterStatusRequest) Reset()         { *m = RegisterStatusRequest{} }
func (m *RegisterStatusRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterStatusRequest) ProtoMessage()    {}
func (*RegisterStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{2}
}
func (m *RegisterStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterStatusRequest.Unmarshal(m, b)
}
func (m *RegisterStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterStatusRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterStatusRequest.Merge(dst, src)
}
func (m *RegisterStatusRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterStatusRequest.Size(m)
}
func (m *RegisterStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterStatusRequest proto.InternalMessageInfo

func (m *RegisterStatusRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *RegisterStatusRequest) GetDeviceType() []byte {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *RegisterStatusRequest) GetNextHeartbeatInterval() int32 {
	if m != nil {
		return m.NextHeartbeatInterval
	}
	return 0
}

func (m *RegisterStatusRequest) GetRegisterInfo() string {
	if m != nil {
		return m.RegisterInfo
	}
	return ""
}

// 注册状态        响应
type RegisterStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterStatusResponse) Reset()         { *m = RegisterStatusResponse{} }
func (m *RegisterStatusResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterStatusResponse) ProtoMessage()    {}
func (*RegisterStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{3}
}
func (m *RegisterStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterStatusResponse.Unmarshal(m, b)
}
func (m *RegisterStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterStatusResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterStatusResponse.Merge(dst, src)
}
func (m *RegisterStatusResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterStatusResponse.Size(m)
}
func (m *RegisterStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterStatusResponse proto.InternalMessageInfo

// 查询设备状态    请求
type QueryStatusRequest struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatusRequest) Reset()         { *m = QueryStatusRequest{} }
func (m *QueryStatusRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStatusRequest) ProtoMessage()    {}
func (*QueryStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{4}
}
func (m *QueryStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatusRequest.Unmarshal(m, b)
}
func (m *QueryStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatusRequest.Marshal(b, m, deterministic)
}
func (dst *QueryStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatusRequest.Merge(dst, src)
}
func (m *QueryStatusRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStatusRequest.Size(m)
}
func (m *QueryStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatusRequest proto.InternalMessageInfo

func (m *QueryStatusRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

// 查询设备状态    响应
type QueryStatusResponse struct {
	Status               []*QueryStatusResponse_StatusInfo `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *QueryStatusResponse) Reset()         { *m = QueryStatusResponse{} }
func (m *QueryStatusResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStatusResponse) ProtoMessage()    {}
func (*QueryStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{5}
}
func (m *QueryStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatusResponse.Unmarshal(m, b)
}
func (m *QueryStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatusResponse.Marshal(b, m, deterministic)
}
func (dst *QueryStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatusResponse.Merge(dst, src)
}
func (m *QueryStatusResponse) XXX_Size() int {
	return xxx_messageInfo_QueryStatusResponse.Size(m)
}
func (m *QueryStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatusResponse proto.InternalMessageInfo

func (m *QueryStatusResponse) GetStatus() []*QueryStatusResponse_StatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

type QueryStatusResponse_StatusInfo struct {
	DeviceType           []byte   `protobuf:"bytes,1,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	RegisterInfo         string   `protobuf:"bytes,2,opt,name=register_info,json=registerInfo,proto3" json:"register_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatusResponse_StatusInfo) Reset()         { *m = QueryStatusResponse_StatusInfo{} }
func (m *QueryStatusResponse_StatusInfo) String() string { return proto.CompactTextString(m) }
func (*QueryStatusResponse_StatusInfo) ProtoMessage()    {}
func (*QueryStatusResponse_StatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_service_f8aa4236551cad14, []int{5, 0}
}
func (m *QueryStatusResponse_StatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatusResponse_StatusInfo.Unmarshal(m, b)
}
func (m *QueryStatusResponse_StatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatusResponse_StatusInfo.Marshal(b, m, deterministic)
}
func (dst *QueryStatusResponse_StatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatusResponse_StatusInfo.Merge(dst, src)
}
func (m *QueryStatusResponse_StatusInfo) XXX_Size() int {
	return xxx_messageInfo_QueryStatusResponse_StatusInfo.Size(m)
}
func (m *QueryStatusResponse_StatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatusResponse_StatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatusResponse_StatusInfo proto.InternalMessageInfo

func (m *QueryStatusResponse_StatusInfo) GetDeviceType() []byte {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *QueryStatusResponse_StatusInfo) GetRegisterInfo() string {
	if m != nil {
		return m.RegisterInfo
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "BaseResponse")
	proto.RegisterType((*RegisterStatusRequest)(nil), "RegisterStatusRequest")
	proto.RegisterType((*RegisterStatusResponse)(nil), "RegisterStatusResponse")
	proto.RegisterType((*QueryStatusRequest)(nil), "QueryStatusRequest")
	proto.RegisterType((*QueryStatusResponse)(nil), "QueryStatusResponse")
	proto.RegisterType((*QueryStatusResponse_StatusInfo)(nil), "QueryStatusResponse.StatusInfo")
	proto.RegisterEnum("CMD", CMD_name, CMD_value)
	proto.RegisterEnum("ReturnCode", ReturnCode_name, ReturnCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatusServerClient is the client API for StatusServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusServerClient interface {
	BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type statusServerClient struct {
	cc *grpc.ClientConn
}

func NewStatusServerClient(cc *grpc.ClientConn) StatusServerClient {
	return &statusServerClient{cc}
}

func (c *statusServerClient) BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/StatusServer/BaseInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServerServer is the server API for StatusServer service.
type StatusServerServer interface {
	BaseInterface(context.Context, *BaseRequest) (*BaseResponse, error)
}

func RegisterStatusServerServer(s *grpc.Server, srv StatusServerServer) {
	s.RegisterService(&_StatusServer_serviceDesc, srv)
}

func _StatusServer_BaseInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServerServer).BaseInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StatusServer/BaseInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServerServer).BaseInterface(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StatusServer",
	HandlerType: (*StatusServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BaseInterface",
			Handler:    _StatusServer_BaseInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "status-service.proto",
}

func init() {
	proto.RegisterFile("status-service.proto", fileDescriptor_status_service_f8aa4236551cad14)
}

var fileDescriptor_status_service_f8aa4236551cad14 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x9d, 0xd7, 0x6c, 0xd0, 0x9b, 0x6c, 0x14, 0x77, 0x63, 0x5d, 0x85, 0xb4, 0x2a, 0xbc, 0x54,
	0x13, 0x78, 0xa8, 0x48, 0xf0, 0x80, 0x84, 0xd4, 0xa5, 0x11, 0x44, 0x6c, 0xad, 0xe6, 0xb4, 0x42,
	0xdb, 0x4b, 0x94, 0x26, 0x37, 0x23, 0xd2, 0x94, 0x14, 0xc7, 0x9d, 0xc8, 0x23, 0x5f, 0x82, 0xc4,
	0x1f, 0xf0, 0x87, 0x28, 0x71, 0xc2, 0xc6, 0x56, 0x69, 0xbc, 0xd9, 0xe7, 0x5c, 0xfb, 0x9c, 0x7b,
	0xec, 0x0b, 0x3b, 0x99, 0xf4, 0xe5, 0x32, 0x7b, 0x95, 0xa1, 0xb8, 0x8e, 0x03, 0x64, 0x0b, 0x91,
	0xca, 0xb4, 0xbb, 0x7f, 0x99, 0xa6, 0x97, 0x57, 0x78, 0x54, 0xee, 0xe6, 0xcb, 0xe8, 0xc8, 0x4f,
	0x72, 0x45, 0x99, 0x31, 0xe8, 0xc7, 0x7e, 0x86, 0x1c, 0xbf, 0x2d, 0x31, 0x93, 0xf4, 0x39, 0x34,
	0x85, 0x5a, 0x3a, 0x61, 0x87, 0xf4, 0x48, 0xbf, 0xc9, 0x6f, 0x00, 0x4a, 0x81, 0x04, 0x9d, 0xf5,
	0x1e, 0xe9, 0x6f, 0x0f, 0x34, 0x66, 0x9d, 0x8e, 0x38, 0x09, 0x68, 0x1f, 0xb4, 0x79, 0x1a, 0xe6,
	0x9d, 0x46, 0x8f, 0xf4, 0xf5, 0xc1, 0x0e, 0x53, 0x52, 0xac, 0x96, 0x62, 0xc3, 0x24, 0xe7, 0x65,
	0x85, 0xf9, 0x93, 0x80, 0xa1, 0xb4, 0xb2, 0x45, 0x9a, 0x64, 0xf8, 0x80, 0xd8, 0x01, 0x68, 0x41,
	0x1a, 0x62, 0xa5, 0xa7, 0x33, 0x8e, 0x72, 0x29, 0x12, 0x2b, 0x0d, 0x91, 0x97, 0x04, 0xa5, 0xa0,
	0x85, 0x98, 0x05, 0xa5, 0x72, 0x93, 0x97, 0x6b, 0xe5, 0x50, 0x5b, 0xed, 0x70, 0xe3, 0x41, 0x87,
	0xbf, 0x09, 0xec, 0x72, 0xbc, 0x8c, 0x33, 0x89, 0xc2, 0x2d, 0x83, 0xac, 0x73, 0xe9, 0xc2, 0xe3,
	0x38, 0xc4, 0x44, 0xc6, 0x32, 0xaf, 0x9c, 0xfe, 0xdd, 0xd3, 0x03, 0xd0, 0x43, 0x2c, 0xd2, 0xf6,
	0x64, 0xbe, 0x50, 0x7e, 0x0d, 0x0e, 0x0a, 0x9a, 0xe6, 0x0b, 0xa4, 0x6f, 0x61, 0x2f, 0xc1, 0xef,
	0xd2, 0xfb, 0x8a, 0xbe, 0x90, 0x73, 0xf4, 0xa5, 0x17, 0x27, 0x12, 0xc5, 0xb5, 0x7f, 0x55, 0x7a,
	0xdf, 0xe0, 0xbb, 0x05, 0xfd, 0xa9, 0x66, 0x9d, 0x8a, 0xa4, 0x2f, 0x60, 0x4b, 0x54, 0x6e, 0xbc,
	0x38, 0x89, 0xd2, 0xb2, 0xb1, 0x26, 0x37, 0x6a, 0xd0, 0x49, 0xa2, 0xd4, 0xec, 0xc0, 0xb3, 0xbb,
	0x96, 0x55, 0xbc, 0xe6, 0x6b, 0xa0, 0x67, 0x4b, 0x14, 0xf9, 0x7f, 0x77, 0x62, 0xfe, 0x22, 0xd0,
	0xfe, 0xe7, 0x48, 0xf5, 0x50, 0xef, 0x60, 0x53, 0xfd, 0xab, 0x0e, 0xe9, 0x35, 0xfa, 0xfa, 0xe0,
	0x80, 0xad, 0xa8, 0x62, 0x6a, 0x5b, 0x98, 0xe2, 0x55, 0x79, 0x97, 0x03, 0xdc, 0xa0, 0x77, 0x83,
	0x22, 0xf7, 0x82, 0xba, 0xd7, 0xf0, 0xfa, 0xfd, 0x86, 0x0f, 0x5f, 0x42, 0xc3, 0x3a, 0x1d, 0xd1,
	0x36, 0x3c, 0xe1, 0xf6, 0x47, 0xc7, 0x9d, 0xda, 0xdc, 0x73, 0xa7, 0xc3, 0xe9, 0xcc, 0x6d, 0xad,
	0xd1, 0x16, 0x18, 0x67, 0x33, 0x9b, 0x9f, 0xd7, 0x08, 0x39, 0xfc, 0x41, 0x00, 0x6e, 0x7e, 0x0e,
	0xd5, 0xe1, 0x91, 0x3b, 0xb3, 0x2c, 0xdb, 0x2d, 0xaa, 0xb7, 0x01, 0x66, 0xe3, 0xcf, 0x93, 0x2f,
	0x63, 0xcf, 0x3a, 0x1d, 0xb5, 0x48, 0x71, 0xfa, 0x78, 0x32, 0x3a, 0xf7, 0x1c, 0xd7, 0x1b, 0xcf,
	0x4e, 0x4e, 0x5a, 0xeb, 0x74, 0x1f, 0x76, 0x47, 0xb6, 0x6b, 0x73, 0x67, 0x78, 0xe2, 0x5c, 0x0c,
	0xa7, 0xce, 0x64, 0xec, 0xd9, 0x9c, 0x4f, 0x78, 0xab, 0x41, 0xf7, 0xa0, 0xbd, 0x8a, 0xd0, 0x8a,
	0x5b, 0xaa, 0x5b, 0x15, 0xb2, 0x31, 0xf8, 0x00, 0x86, 0x4a, 0xc1, 0x45, 0x71, 0x8d, 0x82, 0x32,
	0xd8, 0x2a, 0xe6, 0xa0, 0x7c, 0xe7, 0xc8, 0x0f, 0x90, 0x1a, 0xec, 0xd6, 0x0c, 0x76, 0xb7, 0xd8,
	0xed, 0x29, 0x31, 0xd7, 0x8e, 0xdb, 0x17, 0x4f, 0xd9, 0xfb, 0xfa, 0xbb, 0x46, 0x28, 0xbc, 0x10,
	0xa3, 0xf9, 0x66, 0x09, 0xbc, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x2e, 0x5d, 0x0c, 0xf2,
	0x03, 0x00, 0x00,
}