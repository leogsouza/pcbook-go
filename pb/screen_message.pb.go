// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=leogsouza.pcbook.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("leogsouza.pcbook.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "leogsouza.pcbook.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "leogsouza.pcbook.Screen.Resolution")
}

func init() { proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7) }

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x4d, 0xf7, 0xd3, 0x37, 0x26, 0x25, 0x0c, 0x29, 0x0a, 0xa3, 0xce, 0x83, 0xbd, 0x98,
	0xc1, 0xf4, 0xe4, 0x51, 0xdd, 0x61, 0x28, 0xdb, 0xc8, 0x10, 0xc1, 0xcb, 0x68, 0x63, 0x68, 0x82,
	0x6d, 0x5f, 0x69, 0x52, 0x84, 0xfd, 0x3f, 0xfe, 0x9f, 0xb2, 0x76, 0xcc, 0xa1, 0x78, 0x7c, 0x2f,
	0x9f, 0xef, 0x8f, 0xf0, 0x60, 0x60, 0x44, 0x21, 0x65, 0xb6, 0x4e, 0xa5, 0x31, 0x61, 0x2c, 0x59,
	0x5e, 0xa0, 0x45, 0xea, 0x26, 0x12, 0x63, 0x83, 0xe5, 0x26, 0x64, 0xb9, 0x88, 0x10, 0x3f, 0x46,
	0x5f, 0x0e, 0xb4, 0x57, 0x15, 0x4a, 0xcf, 0xe1, 0xd8, 0xe8, 0x8d, 0x5c, 0xeb, 0x4c, 0x28, 0x8f,
	0xf8, 0x24, 0x70, 0x78, 0x77, 0xbb, 0x98, 0x65, 0x42, 0xd1, 0x07, 0x80, 0x42, 0x1a, 0x4c, 0x4a,
	0xab, 0x31, 0xf3, 0x1c, 0x9f, 0x04, 0xbd, 0xc9, 0x25, 0xfb, 0x6d, 0xc7, 0x6a, 0x2b, 0xc6, 0xf7,
	0x28, 0x3f, 0x90, 0xd1, 0x5b, 0x68, 0xe5, 0x61, 0x26, 0x13, 0xaf, 0xe1, 0x93, 0xe0, 0x64, 0x32,
	0xfc, 0x57, 0xbf, 0xdc, 0x52, 0xbc, 0x86, 0xe9, 0x10, 0x20, 0x2d, 0x13, 0xab, 0x2d, 0x96, 0x42,
	0x79, 0x4d, 0x9f, 0x04, 0x5d, 0x7e, 0xb0, 0x39, 0xbb, 0x03, 0xf8, 0xc9, 0xa3, 0x03, 0x68, 0x7d,
	0xea, 0x77, 0x5b, 0xff, 0xa0, 0xcf, 0xeb, 0x81, 0x9e, 0x42, 0x5b, 0x49, 0x1d, 0x2b, 0x5b, 0x55,
	0xef, 0xf3, 0xdd, 0x34, 0xba, 0x82, 0x56, 0x95, 0x45, 0x7b, 0xd0, 0x79, 0x99, 0x3f, 0xcd, 0x17,
	0xaf, 0x73, 0xf7, 0x88, 0x76, 0xa0, 0x31, 0x5b, 0xae, 0x5c, 0x42, 0xbb, 0xd0, 0x5c, 0x3c, 0x4f,
	0x1f, 0x5d, 0xe7, 0x7e, 0x0a, 0x43, 0x81, 0x29, 0x8b, 0xb5, 0x55, 0x65, 0xf4, 0xb7, 0x77, 0x1e,
	0x2d, 0xc9, 0xdb, 0xc5, 0xee, 0x55, 0x60, 0x3a, 0xde, 0x13, 0xe3, 0x9a, 0xb8, 0x8e, 0x71, 0x9c,
	0x47, 0x51, 0xbb, 0xba, 0xc3, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x0a, 0xe4, 0x7f,
	0x9f, 0x01, 0x00, 0x00,
}
