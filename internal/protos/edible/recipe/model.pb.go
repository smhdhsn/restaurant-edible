// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/edible/recipe/model.proto

package erpb

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

// The schema of the Recipe message.
type Recipe struct {
	FoodTitle            string   `protobuf:"bytes,1,opt,name=food_title,json=foodTitle,proto3" json:"food_title,omitempty"`
	ComponentTitles      []string `protobuf:"bytes,2,rep,name=component_titles,json=componentTitles,proto3" json:"component_titles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_faec0a1e79b90475, []int{0}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetFoodTitle() string {
	if m != nil {
		return m.FoodTitle
	}
	return ""
}

func (m *Recipe) GetComponentTitles() []string {
	if m != nil {
		return m.ComponentTitles
	}
	return nil
}

func init() {
	proto.RegisterType((*Recipe)(nil), "edible.recipe.model.Recipe")
}

func init() { proto.RegisterFile("protos/edible/recipe/model.proto", fileDescriptor_faec0a1e79b90475) }

var fileDescriptor_faec0a1e79b90475 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4d, 0xc9, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0x4a, 0x4d, 0xce, 0x2c, 0x48,
	0xd5, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x03, 0x4b, 0x09, 0x09, 0x43, 0xa4, 0xf4, 0x20, 0x52,
	0x7a, 0x60, 0x29, 0xa5, 0x20, 0x2e, 0xb6, 0x20, 0x30, 0x5f, 0x48, 0x96, 0x8b, 0x2b, 0x2d, 0x3f,
	0x3f, 0x25, 0xbe, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x13,
	0x24, 0x12, 0x02, 0x12, 0x10, 0xd2, 0xe4, 0x12, 0x48, 0xce, 0xcf, 0x2d, 0xc8, 0xcf, 0x4b, 0xcd,
	0x2b, 0x81, 0xa8, 0x29, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x0c, 0xe2, 0x87, 0x8b, 0x83, 0x55,
	0x16, 0x3b, 0xb1, 0x45, 0xb1, 0xa4, 0x16, 0x15, 0x24, 0x25, 0xb1, 0x81, 0xed, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xec, 0x4b, 0x28, 0x0a, 0x9b, 0x00, 0x00, 0x00,
}