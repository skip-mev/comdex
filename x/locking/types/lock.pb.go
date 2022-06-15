// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/locking/v1beta1/lock.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Locked is a single unit of lock by period. It's a record of locked coin
// at a specific time. It stores owner, duration, and the amount of
// coin locked.
type Lock struct {
	Id        uint64                                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner     string                                  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	CreatedAt time.Time                               `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	Duration  time.Duration                           `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	Coin      github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=coin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_538a493a7ad160f6, []int{0}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Lock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Lock) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Lock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Lock) GetCoin() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Coin
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

// Once unlocking is triggered the lock gets deleted and its being
// added to the Unlocking and it will stay untill the duration of end time.enum
// Once entime is reached lock  is being deleted from unlocking.
type Unlocking struct {
	Id       uint64                                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string                                  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Duration time.Duration                           `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	EndTime  time.Time                               `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	Coin     github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=coin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin"`
}

func (m *Unlocking) Reset()         { *m = Unlocking{} }
func (m *Unlocking) String() string { return proto.CompactTextString(m) }
func (*Unlocking) ProtoMessage()    {}
func (*Unlocking) Descriptor() ([]byte, []int) {
	return fileDescriptor_538a493a7ad160f6, []int{1}
}
func (m *Unlocking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unlocking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unlocking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unlocking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unlocking.Merge(m, src)
}
func (m *Unlocking) XXX_Size() int {
	return m.Size()
}
func (m *Unlocking) XXX_DiscardUnknown() {
	xxx_messageInfo_Unlocking.DiscardUnknown(m)
}

var xxx_messageInfo_Unlocking proto.InternalMessageInfo

func (m *Unlocking) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Unlocking) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Unlocking) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Unlocking) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *Unlocking) GetCoin() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Coin
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

type LockByOwner struct {
	Owner   string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	LockIds []uint64 `protobuf:"varint,2,rep,packed,name=lock_ids,json=lockIds,proto3" json:"lock_ids,omitempty"`
}

func (m *LockByOwner) Reset()         { *m = LockByOwner{} }
func (m *LockByOwner) String() string { return proto.CompactTextString(m) }
func (*LockByOwner) ProtoMessage()    {}
func (*LockByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_538a493a7ad160f6, []int{2}
}
func (m *LockByOwner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockByOwner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockByOwner.Merge(m, src)
}
func (m *LockByOwner) XXX_Size() int {
	return m.Size()
}
func (m *LockByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_LockByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_LockByOwner proto.InternalMessageInfo

func (m *LockByOwner) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *LockByOwner) GetLockIds() []uint64 {
	if m != nil {
		return m.LockIds
	}
	return nil
}

type UnlockingByOwner struct {
	Owner        string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	UnlockingIds []uint64 `protobuf:"varint,2,rep,packed,name=unlocking_ids,json=unlockingIds,proto3" json:"unlocking_ids,omitempty"`
}

func (m *UnlockingByOwner) Reset()         { *m = UnlockingByOwner{} }
func (m *UnlockingByOwner) String() string { return proto.CompactTextString(m) }
func (*UnlockingByOwner) ProtoMessage()    {}
func (*UnlockingByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_538a493a7ad160f6, []int{3}
}
func (m *UnlockingByOwner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnlockingByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnlockingByOwner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnlockingByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockingByOwner.Merge(m, src)
}
func (m *UnlockingByOwner) XXX_Size() int {
	return m.Size()
}
func (m *UnlockingByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockingByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockingByOwner proto.InternalMessageInfo

func (m *UnlockingByOwner) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UnlockingByOwner) GetUnlockingIds() []uint64 {
	if m != nil {
		return m.UnlockingIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Lock)(nil), "comdex.locking.v1beta1.Lock")
	proto.RegisterType((*Unlocking)(nil), "comdex.locking.v1beta1.Unlocking")
	proto.RegisterType((*LockByOwner)(nil), "comdex.locking.v1beta1.LockByOwner")
	proto.RegisterType((*UnlockingByOwner)(nil), "comdex.locking.v1beta1.UnlockingByOwner")
}

func init() { proto.RegisterFile("comdex/locking/v1beta1/lock.proto", fileDescriptor_538a493a7ad160f6) }

var fileDescriptor_538a493a7ad160f6 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x8d, 0x1d, 0xf7, 0x6b, 0xb2, 0xed, 0x07, 0xc5, 0x42, 0xc8, 0x0d, 0xc2, 0x0e, 0x46, 0x82,
	0x1c, 0xa8, 0x57, 0x81, 0x1b, 0x37, 0x0c, 0x97, 0x0a, 0x24, 0x90, 0x05, 0x12, 0xe2, 0x40, 0x64,
	0x7b, 0x37, 0xee, 0x2a, 0xb1, 0x27, 0x8a, 0x37, 0x40, 0xfe, 0x45, 0x8f, 0xfc, 0xa4, 0x1e, 0x7b,
	0xe4, 0x64, 0x20, 0xb9, 0x71, 0xcc, 0x91, 0x13, 0xda, 0xf5, 0xae, 0x5b, 0x40, 0xaa, 0x04, 0x52,
	0x4f, 0xf6, 0xcc, 0xbc, 0x79, 0x33, 0xfb, 0xde, 0x6a, 0xd1, 0xed, 0x14, 0x72, 0x42, 0x3f, 0xe2,
	0x29, 0xa4, 0x13, 0x56, 0x64, 0xf8, 0xfd, 0x30, 0xa1, 0x3c, 0x1e, 0xca, 0x38, 0x98, 0xcd, 0x81,
	0x83, 0x7d, 0xa3, 0x86, 0x04, 0x0a, 0x12, 0x28, 0x48, 0xef, 0x7a, 0x06, 0x19, 0x48, 0x08, 0x16,
	0x7f, 0x35, 0xba, 0xe7, 0x66, 0x00, 0xd9, 0x94, 0x62, 0x19, 0x25, 0x8b, 0x31, 0x26, 0x8b, 0x79,
	0xcc, 0x19, 0x14, 0xaa, 0xee, 0xfd, 0x5e, 0xe7, 0x2c, 0xa7, 0x25, 0x8f, 0xf3, 0x99, 0x26, 0x48,
	0xa1, 0xcc, 0xa1, 0xc4, 0x49, 0x5c, 0xd2, 0x66, 0x9d, 0x14, 0x98, 0x22, 0xf0, 0x57, 0x26, 0xb2,
	0x9e, 0x43, 0x3a, 0xb1, 0xaf, 0x20, 0x93, 0x11, 0xc7, 0xe8, 0x1b, 0x03, 0x2b, 0x32, 0x19, 0xb1,
	0xef, 0xa2, 0x2d, 0xf8, 0x50, 0xd0, 0xb9, 0x63, 0xf6, 0x8d, 0x41, 0x37, 0xdc, 0xdb, 0x54, 0xde,
	0xee, 0x32, 0xce, 0xa7, 0x8f, 0x7c, 0x99, 0xf6, 0xa3, 0xba, 0x6c, 0xbf, 0x41, 0x28, 0x9d, 0xd3,
	0x98, 0x53, 0x32, 0x8a, 0xb9, 0xd3, 0xee, 0x1b, 0x83, 0x9d, 0x07, 0xbd, 0xa0, 0x5e, 0x2b, 0xd0,
	0x6b, 0x05, 0xaf, 0xf4, 0x5a, 0xe1, 0xad, 0x93, 0xca, 0x6b, 0x6d, 0x2a, 0xef, 0x5a, 0x4d, 0x76,
	0xd6, 0xeb, 0x1f, 0x7f, 0xf1, 0x8c, 0xa8, 0xab, 0x12, 0x8f, 0xb9, 0x7d, 0x84, 0x3a, 0xfa, 0xb4,
	0x8e, 0x25, 0x79, 0xf7, 0xff, 0xe0, 0x7d, 0xaa, 0x00, 0xe1, 0x50, 0xd0, 0x7e, 0xaf, 0x3c, 0x5b,
	0xb7, 0xdc, 0x87, 0x9c, 0x71, 0x9a, 0xcf, 0xf8, 0x72, 0x53, 0x79, 0x57, 0xeb, 0x61, 0xba, 0xe6,
	0x7f, 0x12, 0xa3, 0x1a, 0x76, 0xfb, 0x1d, 0xb2, 0x84, 0x24, 0xce, 0x96, 0x9a, 0x52, 0x6b, 0x16,
	0x08, 0xcd, 0xb4, 0x3f, 0xc1, 0x13, 0x60, 0x45, 0x88, 0xc5, 0x94, 0x1f, 0x95, 0x77, 0x2f, 0x63,
	0xfc, 0x68, 0x91, 0x04, 0x29, 0xe4, 0x58, 0x09, 0x5c, 0x7f, 0x0e, 0x4a, 0x32, 0xc1, 0x7c, 0x39,
	0xa3, 0xa5, 0x6c, 0x88, 0x24, 0xaf, 0xff, 0xcd, 0x44, 0xdd, 0xd7, 0x85, 0x72, 0xfc, 0x9f, 0x95,
	0x3e, 0xaf, 0x47, 0xfb, 0x52, 0xf5, 0x88, 0x50, 0x87, 0x16, 0x64, 0x24, 0xee, 0x92, 0x52, 0xfe,
	0x22, 0x47, 0x6f, 0x2a, 0x47, 0x15, 0xa9, 0xee, 0xac, 0xfd, 0xdc, 0xa6, 0x05, 0x11, 0xd0, 0x4b,
	0xd7, 0xf8, 0x25, 0xda, 0x11, 0xf7, 0x38, 0x5c, 0xbe, 0x90, 0x62, 0x35, 0xa2, 0x1a, 0x17, 0x8b,
	0xba, 0x8f, 0x3a, 0xc2, 0x97, 0x11, 0x23, 0xa5, 0x63, 0xf6, 0xdb, 0x03, 0x2b, 0xda, 0x16, 0xf1,
	0x21, 0x29, 0xfd, 0x11, 0xda, 0x6b, 0x4c, 0xfb, 0x5b, 0xda, 0x3b, 0xe8, 0xff, 0x85, 0xee, 0x3d,
	0xc7, 0xbd, 0xdb, 0x24, 0x0f, 0x49, 0x19, 0x3e, 0x3b, 0x59, 0xb9, 0xc6, 0xe9, 0xca, 0x35, 0xbe,
	0xae, 0x5c, 0xe3, 0x78, 0xed, 0xb6, 0x4e, 0xd7, 0x6e, 0xeb, 0xf3, 0xda, 0x6d, 0xbd, 0x1d, 0xfe,
	0x72, 0x76, 0xf1, 0x5e, 0x1c, 0xc0, 0x78, 0xcc, 0x52, 0x16, 0x4f, 0x55, 0x8c, 0xcf, 0x1e, 0x19,
	0x29, 0x45, 0xf2, 0x9f, 0x74, 0xe6, 0xe1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xa1, 0xf3,
	0xb0, 0x83, 0x04, 0x00, 0x00,
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLock(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintLock(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Unlocking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unlocking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unlocking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintLock(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x22
	n6, err6 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintLock(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LockByOwner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockByOwner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockByOwner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockIds) > 0 {
		dAtA8 := make([]byte, len(m.LockIds)*10)
		var j7 int
		for _, num := range m.LockIds {
			for num >= 1<<7 {
				dAtA8[j7] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j7++
			}
			dAtA8[j7] = uint8(num)
			j7++
		}
		i -= j7
		copy(dAtA[i:], dAtA8[:j7])
		i = encodeVarintLock(dAtA, i, uint64(j7))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnlockingByOwner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnlockingByOwner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnlockingByOwner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnlockingIds) > 0 {
		dAtA10 := make([]byte, len(m.UnlockingIds)*10)
		var j9 int
		for _, num := range m.UnlockingIds {
			for num >= 1<<7 {
				dAtA10[j9] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j9++
			}
			dAtA10[j9] = uint8(num)
			j9++
		}
		i -= j9
		copy(dAtA[i:], dAtA10[:j9])
		i = encodeVarintLock(dAtA, i, uint64(j9))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLock(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = m.Coin.Size()
	n += 1 + l + sovLock(uint64(l))
	return n
}

func (m *Unlocking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLock(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	l = m.Coin.Size()
	n += 1 + l + sovLock(uint64(l))
	return n
}

func (m *LockByOwner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	if len(m.LockIds) > 0 {
		l = 0
		for _, e := range m.LockIds {
			l += sovLock(uint64(e))
		}
		n += 1 + sovLock(uint64(l)) + l
	}
	return n
}

func (m *UnlockingByOwner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	if len(m.UnlockingIds) > 0 {
		l = 0
		for _, e := range m.UnlockingIds {
			l += sovLock(uint64(e))
		}
		n += 1 + sovLock(uint64(l)) + l
	}
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Unlocking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unlocking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unlocking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LockByOwner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LockByOwner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockByOwner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLock
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.LockIds = append(m.LockIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLock
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLock
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLock
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.LockIds) == 0 {
					m.LockIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLock
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.LockIds = append(m.LockIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field LockIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnlockingByOwner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnlockingByOwner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnlockingByOwner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLock
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UnlockingIds = append(m.UnlockingIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLock
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLock
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLock
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.UnlockingIds) == 0 {
					m.UnlockingIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLock
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.UnlockingIds = append(m.UnlockingIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockingIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)
