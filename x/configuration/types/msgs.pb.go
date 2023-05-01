// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/configuration/v1beta1/msgs.proto

package types

import (
	fmt "fmt"
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

// MsgUpdateConfig is used to update starname configuration
type MsgUpdateConfig struct {
	// Signer is the address of the entity who is doing the transaction
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
	// NewConfiguration contains the new configuration data
	NewConfiguration *Config `protobuf:"bytes,2,opt,name=new_configuration,json=newConfiguration,proto3" json:"new_configuration,omitempty" yaml:"new_configuration"`
}

func (m *MsgUpdateConfig) Reset()         { *m = MsgUpdateConfig{} }
func (m *MsgUpdateConfig) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateConfig) ProtoMessage()    {}
func (*MsgUpdateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae93b9835c563ab6, []int{0}
}
func (m *MsgUpdateConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateConfig.Merge(m, src)
}
func (m *MsgUpdateConfig) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateConfig proto.InternalMessageInfo

func (m *MsgUpdateConfig) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgUpdateConfig) GetNewConfiguration() *Config {
	if m != nil {
		return m.NewConfiguration
	}
	return nil
}

// MsgUpdateFees is used to update the starname product fees in the starname
// module.
type MsgUpdateFees struct {
	Fees       *Fees  `protobuf:"bytes,1,opt,name=fees,proto3" json:"fees,omitempty"`
	Configurer string `protobuf:"bytes,2,opt,name=configurer,proto3" json:"configurer,omitempty" yaml:"configurer"`
}

func (m *MsgUpdateFees) Reset()         { *m = MsgUpdateFees{} }
func (m *MsgUpdateFees) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateFees) ProtoMessage()    {}
func (*MsgUpdateFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae93b9835c563ab6, []int{1}
}
func (m *MsgUpdateFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateFees.Merge(m, src)
}
func (m *MsgUpdateFees) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateFees) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateFees.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateFees proto.InternalMessageInfo

func (m *MsgUpdateFees) GetFees() *Fees {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *MsgUpdateFees) GetConfigurer() string {
	if m != nil {
		return m.Configurer
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgUpdateConfig)(nil), "ummachain.ummacore.configuration.v1beta1.MsgUpdateConfig")
	proto.RegisterType((*MsgUpdateFees)(nil), "ummachain.ummacore.configuration.v1beta1.MsgUpdateFees")
}

func init() {
	proto.RegisterFile("iov/configuration/v1beta1/msgs.proto", fileDescriptor_ae93b9835c563ab6)
}

var fileDescriptor_ae93b9835c563ab6 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0xcc, 0x2f, 0xd3,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x2d, 0x4e, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x2c, 0x2e, 0x49, 0x2c, 0xca, 0x4b, 0xcc, 0x4d, 0x4d, 0xd1, 0xab, 0xd0, 0x43, 0x51,
	0xad, 0x07, 0x55, 0x2d, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xad, 0x0f, 0x62, 0x41, 0x34,
	0x4a, 0xa9, 0xe2, 0x36, 0xbe, 0xa4, 0xb2, 0x20, 0x15, 0x6a, 0xbe, 0xd2, 0x36, 0x46, 0x2e, 0x7e,
	0xdf, 0xe2, 0xf4, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0x67, 0xb0, 0x72, 0x21, 0x4d, 0x2e, 0xb6,
	0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xc1, 0x4f, 0xf7,
	0xe4, 0x79, 0x2b, 0x13, 0x73, 0x73, 0xac, 0x94, 0x20, 0xe2, 0x4a, 0x41, 0x50, 0x05, 0x42, 0x15,
	0x5c, 0x82, 0x79, 0xa9, 0xe5, 0xf1, 0x28, 0xf6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x69,
	0xea, 0x11, 0x74, 0xba, 0x1e, 0xc4, 0x42, 0x27, 0x85, 0x13, 0xf7, 0xe4, 0x19, 0x3f, 0xdd, 0x93,
	0x97, 0x80, 0x58, 0x82, 0x61, 0xa2, 0x52, 0x90, 0x40, 0x5e, 0x6a, 0xb9, 0x33, 0x8a, 0x50, 0x33,
	0x23, 0x17, 0x2f, 0xdc, 0xe1, 0x6e, 0xa9, 0xa9, 0xc5, 0x42, 0xd6, 0x5c, 0x2c, 0x69, 0xa9, 0xa9,
	0xc5, 0x60, 0x47, 0x73, 0x1b, 0xa9, 0x13, 0x61, 0x3d, 0x48, 0x5b, 0x10, 0x58, 0x93, 0x90, 0x29,
	0x17, 0x17, 0x4c, 0x4d, 0x6a, 0x11, 0xd8, 0x07, 0x9c, 0x4e, 0xa2, 0x9f, 0xee, 0xc9, 0x0b, 0x42,
	0x9c, 0x84, 0x90, 0x53, 0x0a, 0x42, 0x52, 0xe8, 0xe4, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x46, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x99, 0xf9, 0x65, 0xba, 0xf9, 0x79, 0xa9, 0xfa, 0x70, 0x17, 0xe9, 0x57, 0xa0, 0x45, 0x0d, 0x38,
	0x4a, 0x92, 0xd8, 0xc0, 0x71, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xf0, 0x5e, 0x48,
	0x1b, 0x02, 0x00, 0x00,
}

func (m *MsgUpdateConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewConfiguration != nil {
		{
			size, err := m.NewConfiguration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Configurer) > 0 {
		i -= len(m.Configurer)
		copy(dAtA[i:], m.Configurer)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Configurer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Fees != nil {
		{
			size, err := m.Fees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	if m.NewConfiguration != nil {
		l = m.NewConfiguration.Size()
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func (m *MsgUpdateFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.Configurer)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgUpdateConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewConfiguration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NewConfiguration == nil {
				m.NewConfiguration = &Config{}
			}
			if err := m.NewConfiguration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgUpdateFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgUpdateFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &Fees{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configurer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configurer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
