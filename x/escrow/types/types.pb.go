// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/escrow/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// EscrowState defines the state of an escrow
type EscrowState int32

const (
	// ESCROW_STATE_OPEN_UNSPECIFIED defines an open state. TODO:: review the
	// _UNSPECIFIED sufix
	EscrowState_Open EscrowState = 0
	// ESCROW_STATE_COMPLETED defines a completed state.
	EscrowState_Completed EscrowState = 1
	// ESCROW_STATE_REFUNDED defines a refunded state.
	EscrowState_Refunded EscrowState = 2
	// ESCROW_STATE_REFUNDED defines an expired state.
	EscrowState_Expired EscrowState = 3
)

var EscrowState_name = map[int32]string{
	0: "ESCROW_STATE_OPEN_UNSPECIFIED",
	1: "ESCROW_STATE_COMPLETED",
	2: "ESCROW_STATE_REFUNDED",
	3: "ESCROW_STATE_EXPIRED",
}

var EscrowState_value = map[string]int32{
	"ESCROW_STATE_OPEN_UNSPECIFIED": 0,
	"ESCROW_STATE_COMPLETED":        1,
	"ESCROW_STATE_REFUNDED":         2,
	"ESCROW_STATE_EXPIRED":          3,
}

func (x EscrowState) String() string {
	return proto.EnumName(EscrowState_name, int32(x))
}

func (EscrowState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_06970306f8aa7966, []int{0}
}

// Escrow defines the struct of an escrow
type Escrow struct {
	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Seller string     `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	Object *types.Any `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	// TODO: refactor this to use sdk.Coin instead of sdk.Coins
	// Although the price contains multiple coins, for now we enforce a specific
	// denomination, so there will be only one coin type in a valid escrow
	Price            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	State            EscrowState                              `protobuf:"varint,5,opt,name=state,proto3,enum=ummachain.ummacore.escrow.v1beta1.EscrowState" json:"state,omitempty"`
	Deadline         uint64                                   `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	BrokerAddress    string                                   `protobuf:"bytes,7,opt,name=broker_address,json=brokerAddress,proto3" json:"broker_address,omitempty"`
	BrokerCommission github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,8,opt,name=broker_commission,json=brokerCommission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"broker_commission"`
}

func (m *Escrow) Reset()         { *m = Escrow{} }
func (m *Escrow) String() string { return proto.CompactTextString(m) }
func (*Escrow) ProtoMessage()    {}
func (*Escrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_06970306f8aa7966, []int{0}
}
func (m *Escrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Escrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Escrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Escrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Escrow.Merge(m, src)
}
func (m *Escrow) XXX_Size() int {
	return m.Size()
}
func (m *Escrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Escrow.DiscardUnknown(m)
}

var xxx_messageInfo_Escrow proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ummachain.ummacore.escrow.v1beta1.EscrowState", EscrowState_name, EscrowState_value)
	proto.RegisterType((*Escrow)(nil), "ummachain.ummacore.escrow.v1beta1.Escrow")
}

func init() { proto.RegisterFile("iov/escrow/v1beta1/types.proto", fileDescriptor_06970306f8aa7966) }

var fileDescriptor_06970306f8aa7966 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0x24, 0x4d, 0xd3, 0x0b, 0xad, 0xc2, 0xa9, 0x54, 0xae, 0x25, 0x5c, 0x0b, 0xa9,
	0x34, 0x50, 0xf5, 0x4c, 0xcb, 0xcc, 0xd0, 0xc4, 0x57, 0x29, 0x12, 0x34, 0x91, 0x93, 0x0a, 0x04,
	0x43, 0x64, 0xfb, 0xde, 0x86, 0xa3, 0x89, 0x2f, 0xf2, 0xb9, 0xa5, 0xfd, 0x06, 0x28, 0x13, 0x5f,
	0x20, 0x13, 0x1b, 0x33, 0x03, 0x12, 0x5f, 0xa0, 0x62, 0xea, 0x88, 0x18, 0x0a, 0xb4, 0x5f, 0x04,
	0xd9, 0xe7, 0x44, 0xed, 0x80, 0xc4, 0xe4, 0x7b, 0xdf, 0xfb, 0x3d, 0xef, 0x9f, 0xc7, 0x36, 0x32,
	0xb9, 0x38, 0xb1, 0x41, 0x06, 0x91, 0x78, 0x6f, 0x9f, 0x6c, 0xfb, 0x10, 0x7b, 0xdb, 0x76, 0x7c,
	0x36, 0x02, 0x49, 0x46, 0x91, 0x88, 0x05, 0x36, 0x64, 0xec, 0x45, 0xa1, 0x37, 0x04, 0x46, 0x4e,
	0x89, 0xe2, 0x48, 0xc6, 0x19, 0x66, 0x20, 0xe4, 0x50, 0x48, 0xdb, 0xf7, 0x24, 0xcc, 0xc4, 0x81,
	0xe0, 0xa1, 0xd2, 0x1a, 0xcb, 0x7d, 0xd1, 0x17, 0xe9, 0xd1, 0x4e, 0x4e, 0x59, 0x76, 0xb5, 0x2f,
	0x44, 0x7f, 0x00, 0x76, 0x1a, 0xf9, 0xc7, 0x87, 0xb6, 0x17, 0x9e, 0x4d, 0xaf, 0x54, 0xc1, 0x9e,
	0xd2, 0xa8, 0x40, 0x5d, 0x3d, 0xf8, 0x5a, 0x40, 0x25, 0x9a, 0xb6, 0xc7, 0x4b, 0x28, 0xcf, 0x99,
	0xae, 0x59, 0x5a, 0x6d, 0xc1, 0xcd, 0x73, 0x86, 0x57, 0x50, 0x49, 0xc2, 0x60, 0x00, 0x91, 0x9e,
	0x4f, 0x73, 0x59, 0x84, 0x1d, 0x54, 0x12, 0xfe, 0x3b, 0x08, 0x62, 0xbd, 0x60, 0x69, 0xb5, 0xca,
	0xce, 0x32, 0x51, 0x9d, 0xc9, 0xb4, 0x33, 0xd9, 0x0d, 0xcf, 0xea, 0x2b, 0xdf, 0xbf, 0x6c, 0xe1,
	0x6e, 0xe4, 0x85, 0xf2, 0x10, 0x22, 0xcf, 0x1f, 0x40, 0x2b, 0xd5, 0xb8, 0x99, 0x16, 0x7b, 0x68,
	0x6e, 0x14, 0xf1, 0x00, 0xf4, 0xa2, 0x55, 0xa8, 0x55, 0x76, 0x56, 0x49, 0x36, 0x56, 0xb2, 0xf4,
	0xd4, 0x09, 0xd2, 0x10, 0x3c, 0xac, 0x3f, 0x39, 0xbf, 0x5c, 0xcb, 0x7d, 0xfe, 0xb5, 0x56, 0xeb,
	0xf3, 0xf8, 0xed, 0xb1, 0x4f, 0x02, 0x31, 0xcc, 0x76, 0xc8, 0x1e, 0x5b, 0x92, 0x1d, 0x65, 0xe6,
	0x26, 0x02, 0xe9, 0xaa, 0xca, 0xf8, 0x19, 0x9a, 0x93, 0xb1, 0x17, 0x83, 0x3e, 0x67, 0x69, 0xb5,
	0xa5, 0x9d, 0x0d, 0xf2, 0x6f, 0xcf, 0x89, 0xf2, 0xa0, 0x93, 0xe0, 0xae, 0x52, 0x61, 0x03, 0x95,
	0x19, 0x78, 0x6c, 0xc0, 0x43, 0xd0, 0x4b, 0x96, 0x56, 0x2b, 0xba, 0xb3, 0x18, 0xaf, 0xa3, 0x25,
	0x3f, 0x12, 0x47, 0x10, 0xf5, 0x3c, 0xc6, 0x22, 0x90, 0x52, 0x9f, 0x4f, 0x3d, 0x5a, 0x54, 0xd9,
	0x5d, 0x95, 0xc4, 0x6f, 0xd0, 0xdd, 0x0c, 0x0b, 0xc4, 0x70, 0xc8, 0xa5, 0xe4, 0x22, 0xd4, 0xcb,
	0x09, 0x59, 0x27, 0xc9, 0x56, 0x3f, 0x2f, 0xd7, 0x1e, 0xfe, 0xc7, 0x56, 0x0e, 0x04, 0x6e, 0x55,
	0x15, 0x6a, 0xcc, 0xea, 0x3c, 0xfe, 0xa6, 0xa1, 0xca, 0x8d, 0xb1, 0xf1, 0x26, 0xba, 0x4f, 0x3b,
	0x0d, 0xb7, 0xf5, 0xb2, 0xd7, 0xe9, 0xee, 0x76, 0x69, 0xaf, 0xd5, 0xa6, 0xfb, 0xbd, 0x83, 0xfd,
	0x4e, 0x9b, 0x36, 0x9a, 0x7b, 0x4d, 0xea, 0x54, 0x73, 0x46, 0x79, 0x3c, 0xb1, 0x8a, 0xad, 0x11,
	0x84, 0xf8, 0x11, 0x5a, 0xb9, 0x05, 0x37, 0x5a, 0x2f, 0xda, 0xcf, 0x69, 0x97, 0x3a, 0x55, 0xcd,
	0x58, 0x1c, 0x4f, 0xac, 0x85, 0x86, 0x18, 0x8e, 0x06, 0x10, 0x03, 0xc3, 0x1b, 0xe8, 0xde, 0x2d,
	0xd4, 0xa5, 0x7b, 0x07, 0xfb, 0x0e, 0x75, 0xaa, 0x79, 0xe3, 0xce, 0x78, 0x62, 0x95, 0x5d, 0x38,
	0x3c, 0x0e, 0x19, 0x30, 0xbc, 0x8e, 0x96, 0x6f, 0x81, 0xf4, 0x55, 0xbb, 0xe9, 0x52, 0xa7, 0x5a,
	0x30, 0x2a, 0xe3, 0x89, 0x35, 0x4f, 0x4f, 0x47, 0x3c, 0x02, 0x66, 0x14, 0x3f, 0x7c, 0x32, 0xb5,
	0x7a, 0xf3, 0xfc, 0x8f, 0x99, 0x3b, 0xbf, 0x32, 0xb5, 0x8b, 0x2b, 0x53, 0xfb, 0x7d, 0x65, 0x6a,
	0x1f, 0xaf, 0xcd, 0xdc, 0xc5, 0xb5, 0x99, 0xfb, 0x71, 0x6d, 0xe6, 0x5e, 0x6f, 0xde, 0x70, 0x85,
	0x8b, 0x93, 0x2d, 0x11, 0x82, 0x3d, 0x7b, 0x7b, 0xf6, 0xe9, 0xf4, 0xcf, 0x4a, 0xed, 0xf1, 0x4b,
	0xe9, 0x87, 0xf7, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa2, 0x24, 0xc6, 0x74, 0x03,
	0x00, 0x00,
}

func (m *Escrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Escrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Escrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BrokerCommission.Size()
		i -= size
		if _, err := m.BrokerCommission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.BrokerAddress) > 0 {
		i -= len(m.BrokerAddress)
		copy(dAtA[i:], m.BrokerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BrokerAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Deadline != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x30
	}
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Object != nil {
		{
			size, err := m.Object.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Escrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Object != nil {
		l = m.Object.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	if m.Deadline != 0 {
		n += 1 + sovTypes(uint64(m.Deadline))
	}
	l = len(m.BrokerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.BrokerCommission.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Escrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Escrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Escrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Object == nil {
				m.Object = &types.Any{}
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types1.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= EscrowState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BrokerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BrokerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BrokerCommission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BrokerCommission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)