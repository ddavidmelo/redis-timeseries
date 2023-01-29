// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timeseries.proto

package timeseries

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type DataPoint struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f432f906d7ab9cb, []int{0}
}

func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataPoint.Unmarshal(m, b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
}
func (m *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(m, src)
}
func (m *DataPoint) XXX_Size() int {
	return xxx_messageInfo_DataPoint.Size(m)
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DataPoint) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TimeSeries struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 []*DataPoint `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f432f906d7ab9cb, []int{1}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TimeSeries) GetData() []*DataPoint {
	if m != nil {
		return m.Data
	}
	return nil
}

type TimeSerieRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	FromTimestamp        int64    `protobuf:"varint,2,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp          int64    `protobuf:"varint,3,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSerieRequest) Reset()         { *m = TimeSerieRequest{} }
func (m *TimeSerieRequest) String() string { return proto.CompactTextString(m) }
func (*TimeSerieRequest) ProtoMessage()    {}
func (*TimeSerieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f432f906d7ab9cb, []int{2}
}

func (m *TimeSerieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSerieRequest.Unmarshal(m, b)
}
func (m *TimeSerieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSerieRequest.Marshal(b, m, deterministic)
}
func (m *TimeSerieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSerieRequest.Merge(m, src)
}
func (m *TimeSerieRequest) XXX_Size() int {
	return xxx_messageInfo_TimeSerieRequest.Size(m)
}
func (m *TimeSerieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSerieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSerieRequest proto.InternalMessageInfo

func (m *TimeSerieRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TimeSerieRequest) GetFromTimestamp() int64 {
	if m != nil {
		return m.FromTimestamp
	}
	return 0
}

func (m *TimeSerieRequest) GetToTimestamp() int64 {
	if m != nil {
		return m.ToTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*DataPoint)(nil), "timeseries.DataPoint")
	proto.RegisterType((*TimeSeries)(nil), "timeseries.TimeSeries")
	proto.RegisterType((*TimeSerieRequest)(nil), "timeseries.TimeSerieRequest")
}

func init() {
	proto.RegisterFile("timeseries.proto", fileDescriptor_8f432f906d7ab9cb)
}

var fileDescriptor_8f432f906d7ab9cb = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0xa3, 0x42, 0xdf, 0x9c, 0xd4, 0xa0, 0x52, 0x64, 0x87, 0x5a, 0x10, 0xea, 0x65,
	0x87, 0xf9, 0x01, 0xbc, 0x28, 0xe2, 0x4d, 0xb2, 0x9d, 0xbc, 0x48, 0xd4, 0x27, 0x04, 0xcd, 0x32,
	0x93, 0xd7, 0x82, 0xdf, 0x5e, 0x1a, 0x21, 0x09, 0xa3, 0xb7, 0xe4, 0xff, 0x7e, 0xf9, 0xe7, 0x97,
	0x40, 0x45, 0xda, 0xa0, 0x47, 0xa7, 0xd1, 0x2f, 0x77, 0xce, 0x92, 0x15, 0x90, 0x92, 0xf6, 0x0e,
	0xca, 0x7b, 0x45, 0xea, 0xd9, 0xea, 0x2d, 0x89, 0x05, 0x94, 0x61, 0x44, 0xca, 0xec, 0x6a, 0xd6,
	0xb0, 0x8e, 0xcb, 0x14, 0x88, 0x33, 0x38, 0x1c, 0xd4, 0x77, 0x8f, 0x75, 0xd1, 0xb0, 0x8e, 0xc9,
	0xff, 0x4d, 0xfb, 0x04, 0xb0, 0xd1, 0x06, 0xd7, 0xa1, 0x4e, 0x54, 0xc0, 0xbf, 0xf0, 0x37, 0x9c,
	0x2d, 0xe5, 0xb8, 0x14, 0x37, 0x70, 0xf0, 0xa1, 0x48, 0xd5, 0x45, 0xc3, 0xbb, 0xd9, 0xea, 0x7c,
	0x99, 0xd9, 0xc4, 0x8b, 0x65, 0x40, 0xda, 0x2d, 0x54, 0xb1, 0x4a, 0xe2, 0x4f, 0x8f, 0x9e, 0x26,
	0x0a, 0xaf, 0xe1, 0xe4, 0xd3, 0x59, 0xf3, 0x9a, 0x4c, 0x8b, 0x60, 0x3a, 0x1f, 0xd3, 0x4d, 0xb4,
	0xbd, 0x82, 0x63, 0xb2, 0x19, 0xc4, 0x03, 0x34, 0x23, 0x1b, 0x91, 0xd5, 0x0b, 0x9c, 0x26, 0xf5,
	0x35, 0xba, 0x41, 0xbf, 0xa3, 0x78, 0x80, 0xf9, 0x23, 0x52, 0xf6, 0xa4, 0x45, 0xae, 0xbc, 0xef,
	0x77, 0x79, 0x31, 0x39, 0xf5, 0x6f, 0x47, 0xe1, 0xab, 0x6f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0xcf, 0x45, 0x69, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TimeSeriesServiceClient is the client API for TimeSeriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeSeriesServiceClient interface {
	GetTimeSeries(ctx context.Context, in *TimeSerieRequest, opts ...grpc.CallOption) (*TimeSeries, error)
}

type timeSeriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSeriesServiceClient(cc grpc.ClientConnInterface) TimeSeriesServiceClient {
	return &timeSeriesServiceClient{cc}
}

func (c *timeSeriesServiceClient) GetTimeSeries(ctx context.Context, in *TimeSerieRequest, opts ...grpc.CallOption) (*TimeSeries, error) {
	out := new(TimeSeries)
	err := c.cc.Invoke(ctx, "/timeseries.TimeSeriesService/GetTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSeriesServiceServer is the server API for TimeSeriesService service.
type TimeSeriesServiceServer interface {
	GetTimeSeries(context.Context, *TimeSerieRequest) (*TimeSeries, error)
}

// UnimplementedTimeSeriesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTimeSeriesServiceServer struct {
}

func (*UnimplementedTimeSeriesServiceServer) GetTimeSeries(ctx context.Context, req *TimeSerieRequest) (*TimeSeries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeries not implemented")
}

func RegisterTimeSeriesServiceServer(s *grpc.Server, srv TimeSeriesServiceServer) {
	s.RegisterService(&_TimeSeriesService_serviceDesc, srv)
}

func _TimeSeriesService_GetTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeSerieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeseries.TimeSeriesService/GetTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeries(ctx, req.(*TimeSerieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeSeriesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timeseries.TimeSeriesService",
	HandlerType: (*TimeSeriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeSeries",
			Handler:    _TimeSeriesService_GetTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timeseries.proto",
}
