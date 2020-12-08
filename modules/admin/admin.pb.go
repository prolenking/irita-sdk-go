// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin/admin.proto

package admin

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Role is a type alias that represents a proposal status as a byte
type Role int32

const (
	// ROOT_ADMIN defines the root admin role index.
	RoleRootAdmin Role = 0
	// PERM_ADMIN defines the permission admin role index.
	RolePermAdmin Role = 1
	// BLACKLIST_ADMIN defines the blacklist admin role index.
	RoleBlacklistAdmin Role = 2
	// NODE_ADMIN defines the validator node admin role index.
	RoleNodeAdmin Role = 3
	// PARAM_ADMIN defines the param admin role index.
	RoleParamAdmin Role = 4
	// POWER_USER defines the power user role index.
	RolePowerUser Role = 5
	// RELAYER_USER defines the relayer role index.
	RoleRelayerUser Role = 6
	// IDAdmin defines the identity role index.
	RoleIDAdmin Role = 7
)

var Role_name = map[int32]string{
	0: "ROOT_ADMIN",
	1: "PERM_ADMIN",
	2: "BLACKLIST_ADMIN",
	3: "NODE_ADMIN",
	4: "PARAM_ADMIN",
	5: "POWER_USER",
	6: "RELAYER_USER",
	7: "ID_ADMIN",
}

var Role_value = map[string]int32{
	"ROOT_ADMIN":      0,
	"PERM_ADMIN":      1,
	"BLACKLIST_ADMIN": 2,
	"NODE_ADMIN":      3,
	"PARAM_ADMIN":     4,
	"POWER_USER":      5,
	"RELAYER_USER":    6,
	"ID_ADMIN":        7,
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{0}
}

// MsgAddRoles defines an SDK message for adding roles to a address.
type MsgAddRoles struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Roles    []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=iritamod.admin.Role" json:"roles,omitempty"`
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgAddRoles) Reset()         { *m = MsgAddRoles{} }
func (m *MsgAddRoles) String() string { return proto.CompactTextString(m) }
func (*MsgAddRoles) ProtoMessage()    {}
func (*MsgAddRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{0}
}
func (m *MsgAddRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddRoles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddRoles.Merge(m, src)
}
func (m *MsgAddRoles) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddRoles.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddRoles proto.InternalMessageInfo

// MsgRemoveRoles defines an SDK message for removing roles from an existing address.
type MsgRemoveRoles struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Roles    []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=iritamod.admin.Role" json:"roles,omitempty"`
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgRemoveRoles) Reset()         { *m = MsgRemoveRoles{} }
func (m *MsgRemoveRoles) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveRoles) ProtoMessage()    {}
func (*MsgRemoveRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{1}
}
func (m *MsgRemoveRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveRoles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveRoles.Merge(m, src)
}
func (m *MsgRemoveRoles) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveRoles.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveRoles proto.InternalMessageInfo

// MsgBlockAccount defines an SDK message for blocking an account.
type MsgBlockAccount struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgBlockAccount) Reset()         { *m = MsgBlockAccount{} }
func (m *MsgBlockAccount) String() string { return proto.CompactTextString(m) }
func (*MsgBlockAccount) ProtoMessage()    {}
func (*MsgBlockAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{2}
}
func (m *MsgBlockAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBlockAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBlockAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBlockAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlockAccount.Merge(m, src)
}
func (m *MsgBlockAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgBlockAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlockAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlockAccount proto.InternalMessageInfo

// MsgUnblockAccount defines an SDK message for unblocking an account.
type MsgUnblockAccount struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgUnblockAccount) Reset()         { *m = MsgUnblockAccount{} }
func (m *MsgUnblockAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUnblockAccount) ProtoMessage()    {}
func (*MsgUnblockAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{3}
}
func (m *MsgUnblockAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnblockAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnblockAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnblockAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnblockAccount.Merge(m, src)
}
func (m *MsgUnblockAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnblockAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnblockAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnblockAccount proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("iritamod.admin.Role", Role_name, Role_value)
	proto.RegisterType((*MsgAddRoles)(nil), "iritamod.admin.MsgAddRoles")
	proto.RegisterType((*MsgRemoveRoles)(nil), "iritamod.admin.MsgRemoveRoles")
	proto.RegisterType((*MsgBlockAccount)(nil), "iritamod.admin.MsgBlockAccount")
	proto.RegisterType((*MsgUnblockAccount)(nil), "iritamod.admin.MsgUnblockAccount")
}

func init() { proto.RegisterFile("admin/admin.proto", fileDescriptor_8595c8dce2486799) }

var fileDescriptor_8595c8dce2486799 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x93, 0xb6, 0xfb, 0xe2, 0x74, 0x6d, 0xb7, 0x71, 0x91, 0x12, 0x70, 0x8c, 0x2b, 0xc2,
	0xb2, 0xb2, 0x09, 0xe8, 0xcd, 0x5b, 0x6a, 0x73, 0x28, 0x36, 0x4d, 0x99, 0xb5, 0x88, 0x5e, 0x96,
	0x69, 0x32, 0xc4, 0xd8, 0xa4, 0xcf, 0x32, 0x93, 0xfa, 0xf2, 0x0d, 0x24, 0x27, 0xbf, 0x40, 0x40,
	0x70, 0x0f, 0x1e, 0x3d, 0xfa, 0x11, 0xf6, 0xb8, 0x47, 0x8f, 0xda, 0x5e, 0xfc, 0x18, 0x92, 0x49,
	0x6b, 0x57, 0x10, 0x4f, 0xb2, 0x97, 0x90, 0x49, 0x7e, 0xf9, 0xfd, 0x9f, 0x3f, 0xe4, 0x41, 0x2d,
	0x1a, 0x24, 0xd1, 0xd4, 0x92, 0x57, 0xf3, 0x94, 0x43, 0x0a, 0x5a, 0x23, 0xe2, 0x51, 0x4a, 0x13,
	0x08, 0x4c, 0xf9, 0x54, 0xdf, 0x0b, 0x21, 0x04, 0xf9, 0xca, 0x2a, 0xee, 0x4a, 0x6a, 0x7f, 0x86,
	0xea, 0xae, 0x08, 0xed, 0x20, 0x20, 0x10, 0x33, 0xa1, 0xb5, 0xd1, 0x16, 0x0d, 0x02, 0xce, 0x84,
	0x68, 0xab, 0x86, 0x7a, 0x70, 0x8d, 0xac, 0x8e, 0xda, 0x21, 0xda, 0xe0, 0x05, 0xd2, 0xae, 0x18,
	0xd5, 0x83, 0xc6, 0x83, 0x3d, 0xf3, 0x4f, 0xbd, 0x59, 0x7c, 0x4f, 0x4a, 0x44, 0xd3, 0xd1, 0x36,
	0x9c, 0x32, 0x4e, 0x53, 0xe0, 0xed, 0xaa, 0xd4, 0xfc, 0x3e, 0x3f, 0xaa, 0xfd, 0xfc, 0x78, 0x5b,
	0xdd, 0x7f, 0x8b, 0x1a, 0xae, 0x08, 0x09, 0x4b, 0xe0, 0x35, 0xbb, 0xda, 0x64, 0x17, 0x35, 0x5d,
	0x11, 0x76, 0x62, 0xf0, 0x27, 0xb6, 0xef, 0xc3, 0x6c, 0x9a, 0xfe, 0x23, 0xfa, 0xb2, 0xae, 0xf2,
	0x57, 0x9d, 0x87, 0x5a, 0xae, 0x08, 0x47, 0xd3, 0xf1, 0x7f, 0x12, 0x1e, 0x7e, 0xad, 0xa0, 0x5a,
	0xd1, 0x48, 0xbb, 0x83, 0x10, 0xf1, 0xbc, 0xa7, 0x27, 0x76, 0xd7, 0xed, 0x0d, 0x76, 0x15, 0xbd,
	0x95, 0xe5, 0xc6, 0x75, 0xd9, 0x15, 0x20, 0xb5, 0x8b, 0xe6, 0x05, 0x32, 0x74, 0x88, 0xbb, 0x44,
	0xd4, 0x35, 0x32, 0x64, 0x3c, 0x29, 0x91, 0xfb, 0xa8, 0xd9, 0xe9, 0xdb, 0x8f, 0x9f, 0xf4, 0x7b,
	0xc7, 0x2b, 0x55, 0x45, 0xbf, 0x99, 0xe5, 0x86, 0x56, 0x70, 0x9d, 0x98, 0xfa, 0x93, 0x38, 0x12,
	0x6b, 0xdf, 0xc0, 0xeb, 0x3a, 0x4b, 0xae, 0xba, 0xf6, 0x0d, 0x20, 0x60, 0x25, 0x72, 0x17, 0xd5,
	0x87, 0x36, 0xb1, 0x57, 0x99, 0x35, 0x5d, 0xcb, 0x72, 0xa3, 0x21, 0x33, 0x29, 0xa7, 0xc9, 0x7a,
	0x2e, 0xef, 0x99, 0x43, 0x4e, 0x46, 0xc7, 0x0e, 0xd9, 0xdd, 0xb8, 0x34, 0x17, 0xbc, 0x61, 0x7c,
	0x24, 0x18, 0xd7, 0xee, 0xa1, 0x1d, 0xe2, 0xf4, 0xed, 0xe7, 0x2b, 0x68, 0x53, 0xbf, 0x91, 0xe5,
	0x46, 0x53, 0xf6, 0x63, 0x31, 0x7d, 0xb7, 0xc4, 0x6e, 0xa1, 0xed, 0x5e, 0x77, 0x99, 0xb5, 0xa5,
	0x37, 0xb3, 0xdc, 0xa8, 0x17, 0x48, 0xaf, 0x2b, 0x83, 0xf4, 0x9d, 0xf7, 0x9f, 0xb0, 0xf2, 0xf9,
	0x0c, 0x2b, 0x5f, 0xce, 0xb0, 0xda, 0x71, 0xcf, 0x7f, 0x60, 0xe5, 0x7c, 0x8e, 0xd5, 0x8b, 0x39,
	0x56, 0xbf, 0xcf, 0xb1, 0xfa, 0x61, 0x81, 0x95, 0x8b, 0x05, 0x56, 0xbe, 0x2d, 0xb0, 0xf2, 0xc2,
	0x0a, 0xa3, 0xf4, 0xe5, 0x6c, 0x6c, 0xfa, 0x90, 0x58, 0xe3, 0x88, 0x4e, 0x5f, 0x45, 0x8c, 0x46,
	0x96, 0xfc, 0x97, 0x8e, 0x44, 0x30, 0x39, 0x0a, 0xc1, 0x4a, 0x20, 0x98, 0xc5, 0x4c, 0x94, 0x6b,
	0x34, 0xde, 0x94, 0x1b, 0xf2, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x19, 0xc4, 0xe8,
	0x5c, 0x03, 0x00, 0x00,
}

func (x Role) String() string {
	s, ok := Role_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *MsgAddRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAddRoles)
	if !ok {
		that2, ok := that.(MsgAddRoles)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}
func (this *MsgRemoveRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgRemoveRoles)
	if !ok {
		that2, ok := that.(MsgRemoveRoles)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}
func (this *MsgBlockAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBlockAccount)
	if !ok {
		that2, ok := that.(MsgBlockAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}
func (this *MsgUnblockAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnblockAccount)
	if !ok {
		that2, ok := that.(MsgUnblockAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
}
func (m *MsgAddRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Roles) > 0 {
		dAtA2 := make([]byte, len(m.Roles)*10)
		var j1 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAdmin(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Roles) > 0 {
		dAtA4 := make([]byte, len(m.Roles)*10)
		var j3 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintAdmin(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBlockAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBlockAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBlockAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnblockAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnblockAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnblockAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovAdmin(uint64(e))
		}
		n += 1 + sovAdmin(uint64(l)) + l
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *MsgRemoveRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovAdmin(uint64(e))
		}
		n += 1 + sovAdmin(uint64(l)) + l
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *MsgBlockAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func (m *MsgUnblockAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: MsgAddRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAdmin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Role(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAdmin
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
					return ErrInvalidLengthAdmin
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAdmin
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]Role, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAdmin
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Role(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *MsgRemoveRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: MsgRemoveRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAdmin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Role(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAdmin
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
					return ErrInvalidLengthAdmin
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAdmin
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]Role, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAdmin
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Role(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *MsgBlockAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: MsgBlockAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBlockAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *MsgUnblockAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: MsgUnblockAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnblockAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)