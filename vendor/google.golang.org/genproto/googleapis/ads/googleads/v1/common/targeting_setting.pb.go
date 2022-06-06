// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/targeting_setting.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Settings for the
// <a href="https://support.google.com/google-ads/answer/7365594">
// targeting related features</a>, at Campaign and AdGroup level.
type TargetingSetting struct {
	// The per-targeting-dimension setting to restrict the reach of your campaign
	// or ad group.
	TargetRestrictions   []*TargetRestriction `protobuf:"bytes,1,rep,name=target_restrictions,json=targetRestrictions,proto3" json:"target_restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetingSetting) Reset()         { *m = TargetingSetting{} }
func (m *TargetingSetting) String() string { return proto.CompactTextString(m) }
func (*TargetingSetting) ProtoMessage()    {}
func (*TargetingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd746ab34204e313, []int{0}
}

func (m *TargetingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetingSetting.Unmarshal(m, b)
}
func (m *TargetingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetingSetting.Marshal(b, m, deterministic)
}
func (m *TargetingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetingSetting.Merge(m, src)
}
func (m *TargetingSetting) XXX_Size() int {
	return xxx_messageInfo_TargetingSetting.Size(m)
}
func (m *TargetingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_TargetingSetting proto.InternalMessageInfo

func (m *TargetingSetting) GetTargetRestrictions() []*TargetRestriction {
	if m != nil {
		return m.TargetRestrictions
	}
	return nil
}

// The list of per-targeting-dimension targeting settings.
type TargetRestriction struct {
	// The targeting dimension that these settings apply to.
	TargetingDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,1,opt,name=targeting_dimension,json=targetingDimension,proto3,enum=google.ads.googleads.v1.enums.TargetingDimensionEnum_TargetingDimension" json:"targeting_dimension,omitempty"`
	// Indicates whether to restrict your ads to show only for the criteria you
	// have selected for this targeting_dimension, or to target all values for
	// this targeting_dimension and show ads based on your targeting in other
	// TargetingDimensions. A value of 'true' means that these criteria will only
	// apply bid modifiers, and not affect targeting. A value of 'false' means
	// that these criteria will restrict targeting as well as applying bid
	// modifiers.
	BidOnly              *wrappers.BoolValue `protobuf:"bytes,2,opt,name=bid_only,json=bidOnly,proto3" json:"bid_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TargetRestriction) Reset()         { *m = TargetRestriction{} }
func (m *TargetRestriction) String() string { return proto.CompactTextString(m) }
func (*TargetRestriction) ProtoMessage()    {}
func (*TargetRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd746ab34204e313, []int{1}
}

func (m *TargetRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRestriction.Unmarshal(m, b)
}
func (m *TargetRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRestriction.Marshal(b, m, deterministic)
}
func (m *TargetRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRestriction.Merge(m, src)
}
func (m *TargetRestriction) XXX_Size() int {
	return xxx_messageInfo_TargetRestriction.Size(m)
}
func (m *TargetRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRestriction proto.InternalMessageInfo

func (m *TargetRestriction) GetTargetingDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if m != nil {
		return m.TargetingDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (m *TargetRestriction) GetBidOnly() *wrappers.BoolValue {
	if m != nil {
		return m.BidOnly
	}
	return nil
}

func init() {
	proto.RegisterType((*TargetingSetting)(nil), "google.ads.googleads.v1.common.TargetingSetting")
	proto.RegisterType((*TargetRestriction)(nil), "google.ads.googleads.v1.common.TargetRestriction")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/targeting_setting.proto", fileDescriptor_fd746ab34204e313)
}

var fileDescriptor_fd746ab34204e313 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xb1, 0x03, 0xdb, 0x70, 0x60, 0x6c, 0x1e, 0x83, 0x10, 0x46, 0x08, 0x3e, 0xe5, 0x24,
	0xe1, 0x8c, 0x6d, 0xa0, 0x9d, 0x9c, 0x65, 0xa4, 0xb7, 0x86, 0xb4, 0xf8, 0x50, 0x0c, 0x41, 0x8e,
	0x54, 0x21, 0xb0, 0x25, 0x63, 0xc9, 0x29, 0xf9, 0x3a, 0x3d, 0xf6, 0x7b, 0xf4, 0xd2, 0x8f, 0xd2,
	0x53, 0x3f, 0x42, 0xb1, 0x65, 0xb9, 0x69, 0xda, 0xb4, 0x27, 0xbd, 0x7a, 0xf5, 0x3c, 0xbf, 0xf7,
	0x0f, 0xf2, 0x7e, 0x33, 0x29, 0x59, 0x46, 0x21, 0x26, 0x0a, 0x9a, 0xb0, 0x8e, 0xb6, 0x21, 0xdc,
	0xc8, 0x3c, 0x97, 0x02, 0x6a, 0x5c, 0x32, 0xaa, 0xb9, 0x60, 0x6b, 0x45, 0x75, 0x7d, 0x82, 0xa2,
	0x94, 0x5a, 0xfa, 0x23, 0x23, 0x06, 0x98, 0x28, 0xd0, 0xf9, 0xc0, 0x36, 0x04, 0xc6, 0x37, 0xfc,
	0x73, 0x8c, 0x4b, 0x45, 0x95, 0xab, 0x3d, 0x2c, 0xe1, 0x39, 0x15, 0x8a, 0x4b, 0x61, 0xc0, 0xc3,
	0x16, 0x0c, 0x9b, 0x5b, 0x5a, 0x5d, 0xc2, 0xab, 0x12, 0x17, 0x05, 0x2d, 0x55, 0xfb, 0xfe, 0xc3,
	0x82, 0x0b, 0x0e, 0xb1, 0x10, 0x52, 0x63, 0xcd, 0xa5, 0x68, 0x5f, 0x83, 0xad, 0xf7, 0xe5, 0xdc,
	0xa2, 0xcf, 0x4c, 0xc3, 0x7e, 0xea, 0x7d, 0x33, 0xe5, 0xd6, 0x25, 0x55, 0xba, 0xe4, 0x9b, 0xc6,
	0x30, 0x70, 0xc6, 0xbd, 0x49, 0x7f, 0x1a, 0x82, 0xb7, 0x07, 0x01, 0x06, 0xb7, 0x7a, 0x72, 0xae,
	0x7c, 0x7d, 0x98, 0x52, 0xc1, 0xad, 0xe3, 0x7d, 0x7d, 0xa1, 0xf4, 0x77, 0xb6, 0xf2, 0xb3, 0x41,
	0x07, 0xce, 0xd8, 0x99, 0x7c, 0x9e, 0x9e, 0x1c, 0xad, 0xdc, 0xac, 0x08, 0x74, 0x73, 0xcc, 0xad,
	0xf1, 0xbf, 0xa8, 0xf2, 0x57, 0xd2, 0xb6, 0xa1, 0xfd, 0x9c, 0xff, 0xcb, 0xfb, 0x94, 0x72, 0xb2,
	0x96, 0x22, 0xdb, 0x0d, 0xdc, 0xb1, 0x33, 0xe9, 0x4f, 0x87, 0xb6, 0x9e, 0xdd, 0x2c, 0x98, 0x49,
	0x99, 0xc5, 0x38, 0xab, 0xe8, 0xea, 0x63, 0xca, 0xc9, 0xa9, 0xc8, 0x76, 0xb3, 0x07, 0xc7, 0x0b,
	0x36, 0x32, 0x7f, 0x67, 0x29, 0xb3, 0xef, 0x87, 0x4b, 0x5e, 0xd6, 0xcc, 0xa5, 0x73, 0x31, 0x6f,
	0x8d, 0x4c, 0x66, 0x58, 0x30, 0x20, 0x4b, 0x06, 0x19, 0x15, 0x4d, 0x45, 0xfb, 0x0d, 0x0a, 0xae,
	0x8e, 0xfd, 0xb6, 0xbf, 0xe6, 0xb8, 0x76, 0x7b, 0x8b, 0x28, 0xba, 0x71, 0x47, 0x0b, 0x03, 0x8b,
	0x88, 0x02, 0x26, 0xac, 0xa3, 0x38, 0x04, 0xff, 0x1a, 0xd9, 0x9d, 0x15, 0x24, 0x11, 0x51, 0x49,
	0x27, 0x48, 0xe2, 0x30, 0x31, 0x82, 0x7b, 0x37, 0x30, 0x59, 0x84, 0x22, 0xa2, 0x10, 0xea, 0x24,
	0x08, 0xc5, 0x21, 0x42, 0x46, 0x94, 0x7e, 0x68, 0xba, 0xfb, 0xf9, 0x18, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0xf2, 0x7a, 0x09, 0x0a, 0x03, 0x00, 0x00,
}
