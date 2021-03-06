// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	Filter               *Filter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopRequest) Reset()         { *m = SearchLaptopRequest{} }
func (m *SearchLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopRequest) ProtoMessage()    {}
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{2}
}

func (m *SearchLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopRequest.Unmarshal(m, b)
}
func (m *SearchLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopRequest.Marshal(b, m, deterministic)
}
func (m *SearchLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopRequest.Merge(m, src)
}
func (m *SearchLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopRequest.Size(m)
}
func (m *SearchLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopRequest proto.InternalMessageInfo

func (m *SearchLaptopRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopResponse) Reset()         { *m = SearchLaptopResponse{} }
func (m *SearchLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopResponse) ProtoMessage()    {}
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{3}
}

func (m *SearchLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopResponse.Unmarshal(m, b)
}
func (m *SearchLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopResponse.Marshal(b, m, deterministic)
}
func (m *SearchLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopResponse.Merge(m, src)
}
func (m *SearchLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopResponse.Size(m)
}
func (m *SearchLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopResponse proto.InternalMessageInfo

func (m *SearchLaptopResponse) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type UploadImageRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data                 isUploadImageRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{4}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadImageRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadImageRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
}

type ImageInfo struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType            string   `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{5}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *ImageInfo) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageResponse) Reset()         { *m = UploadImageResponse{} }
func (m *UploadImageResponse) String() string { return proto.CompactTextString(m) }
func (*UploadImageResponse) ProtoMessage()    {}
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{6}
}

func (m *UploadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageResponse.Unmarshal(m, b)
}
func (m *UploadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageResponse.Marshal(b, m, deterministic)
}
func (m *UploadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageResponse.Merge(m, src)
}
func (m *UploadImageResponse) XXX_Size() int {
	return xxx_messageInfo_UploadImageResponse.Size(m)
}
func (m *UploadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageResponse proto.InternalMessageInfo

func (m *UploadImageResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadImageResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "leogsouza.pcbook.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "leogsouza.pcbook.CreateLaptopResponse")
	proto.RegisterType((*SearchLaptopRequest)(nil), "leogsouza.pcbook.SearchLaptopRequest")
	proto.RegisterType((*SearchLaptopResponse)(nil), "leogsouza.pcbook.SearchLaptopResponse")
	proto.RegisterType((*UploadImageRequest)(nil), "leogsouza.pcbook.UploadImageRequest")
	proto.RegisterType((*ImageInfo)(nil), "leogsouza.pcbook.ImageInfo")
	proto.RegisterType((*UploadImageResponse)(nil), "leogsouza.pcbook.UploadImageResponse")
}

func init() { proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71) }

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0x51, 0x55, 0x91, 0xd9, 0x2e, 0x42, 0xde, 0x1c, 0xaa, 0xac, 0x80, 0x25, 0x62,
	0x57, 0xbd, 0x90, 0x2e, 0xe5, 0xc4, 0xb5, 0xfc, 0x69, 0x2b, 0x71, 0x40, 0x29, 0x5c, 0x10, 0x28,
	0x72, 0x12, 0x37, 0xb5, 0x9a, 0xc4, 0x26, 0x76, 0x90, 0xda, 0x77, 0xe2, 0x1d, 0x51, 0x6d, 0xb7,
	0x6a, 0x48, 0x50, 0xc5, 0x2d, 0x9a, 0xf9, 0x7d, 0xdf, 0x4c, 0x66, 0xc6, 0xe0, 0xe6, 0x98, 0x4b,
	0xc6, 0x23, 0x41, 0xaa, 0x5f, 0x34, 0x21, 0x01, 0xaf, 0x98, 0x64, 0xe8, 0x49, 0x4e, 0x58, 0x26,
	0x58, 0xbd, 0xc3, 0x01, 0x4f, 0x62, 0xc6, 0x36, 0xde, 0x81, 0x2b, 0x88, 0x10, 0x38, 0x33, 0x9c,
	0xe7, 0xae, 0x68, 0x2e, 0x49, 0xd5, 0x8c, 0xfa, 0x33, 0xb8, 0x7e, 0x57, 0x11, 0x2c, 0xc9, 0x27,
	0xa5, 0x09, 0xc9, 0xcf, 0x9a, 0x08, 0x89, 0x1e, 0xa0, 0xaf, 0x4d, 0x86, 0xd6, 0xad, 0x35, 0xba,
	0x9c, 0x0c, 0x83, 0xbf, 0xab, 0x04, 0x46, 0x60, 0x38, 0xff, 0x1e, 0xdc, 0xa6, 0x91, 0xe0, 0xac,
	0x14, 0x04, 0x3d, 0x06, 0x9b, 0xa6, 0xca, 0xc5, 0x09, 0x6d, 0x9a, 0xee, 0x0b, 0x2e, 0x09, 0xae,
	0x92, 0x75, 0xab, 0xa0, 0xee, 0xef, 0xdf, 0x05, 0x3f, 0xaa, 0x7c, 0x68, 0x38, 0x7f, 0x0e, 0x6e,
	0xd3, 0xc8, 0x14, 0xfc, 0xff, 0xd6, 0x39, 0xa0, 0xaf, 0x3c, 0x67, 0x38, 0x5d, 0x14, 0x38, 0x23,
	0x87, 0x8e, 0x5e, 0x43, 0x8f, 0x96, 0x2b, 0x66, 0x5c, 0x6e, 0xda, 0x2e, 0x8a, 0x5e, 0x94, 0x2b,
	0x36, 0xbf, 0x08, 0x15, 0x8a, 0x9e, 0x03, 0x24, 0xeb, 0xba, 0xdc, 0x44, 0x29, 0x96, 0x78, 0x68,
	0xdf, 0x5a, 0xa3, 0xc1, 0xfc, 0x22, 0x74, 0x54, 0xec, 0x3d, 0x96, 0x78, 0xda, 0x87, 0xde, 0x3e,
	0xe5, 0xcf, 0xc0, 0x39, 0xaa, 0xd1, 0x0d, 0x38, 0x66, 0x61, 0xc7, 0x41, 0x3d, 0xd2, 0x81, 0x45,
	0x8a, 0x9e, 0x02, 0xd0, 0x3d, 0x19, 0xc9, 0x2d, 0x27, 0xca, 0xd2, 0x09, 0x1d, 0x15, 0xf9, 0xb2,
	0xe5, 0xc4, 0x7f, 0x0b, 0xd7, 0x8d, 0xd6, 0xbb, 0x87, 0x8e, 0x10, 0xf4, 0x04, 0xdd, 0x69, 0xfd,
	0x55, 0xa8, 0xbe, 0x27, 0xbf, 0x6d, 0xb8, 0xd2, 0x83, 0x58, 0xea, 0x7b, 0x42, 0x3f, 0x60, 0x70,
	0xba, 0x42, 0x74, 0xd7, 0xfe, 0xe7, 0x8e, 0x5b, 0xf1, 0xee, 0xcf, 0x61, 0xa6, 0xa9, 0x08, 0x06,
	0xa7, 0x0b, 0xeb, 0xb2, 0xef, 0xb8, 0x8c, 0x2e, 0xfb, 0xae, 0xbd, 0x3f, 0x58, 0xe8, 0x3b, 0x5c,
	0x9e, 0x0c, 0x03, 0xbd, 0x6c, 0x0b, 0xdb, 0x6b, 0xf6, 0xee, 0xce, 0x50, 0xda, 0x7d, 0x64, 0x4d,
	0x3f, 0xc0, 0xb3, 0x84, 0x15, 0x41, 0x46, 0xe5, 0xba, 0x8e, 0xdb, 0x22, 0x1e, 0x7f, 0xb6, 0xbe,
	0xbd, 0x30, 0xd9, 0x84, 0x15, 0xe3, 0x23, 0x31, 0xd6, 0xc4, 0xab, 0x8c, 0x8d, 0x79, 0x1c, 0xf7,
	0xd5, 0xbb, 0x7b, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xd4, 0x70, 0xc1, 0xcd, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error)
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadImageClient, error)
}

type laptopServiceClient struct {
	cc *grpc.ClientConn
}

func NewLaptopServiceClient(cc *grpc.ClientConn) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/leogsouza.pcbook.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[0], "/leogsouza.pcbook.LaptopService/SearchLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceSearchLaptopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LaptopService_SearchLaptopClient interface {
	Recv() (*SearchLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceSearchLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceSearchLaptopClient) Recv() (*SearchLaptopResponse, error) {
	m := new(SearchLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[1], "/leogsouza.pcbook.LaptopService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceUploadImageClient{stream}
	return x, nil
}

type LaptopService_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type laptopServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *laptopServiceUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceUploadImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error
	UploadImage(LaptopService_UploadImageServer) error
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leogsouza.pcbook.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_SearchLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchLaptopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).SearchLaptop(m, &laptopServiceSearchLaptopServer{stream})
}

type LaptopService_SearchLaptopServer interface {
	Send(*SearchLaptopResponse) error
	grpc.ServerStream
}

type laptopServiceSearchLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceSearchLaptopServer) Send(m *SearchLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LaptopService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).UploadImage(&laptopServiceUploadImageServer{stream})
}

type LaptopService_UploadImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type laptopServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *laptopServiceUploadImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "leogsouza.pcbook.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchLaptop",
			Handler:       _LaptopService_SearchLaptop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadImage",
			Handler:       _LaptopService_UploadImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "laptop_service.proto",
}
