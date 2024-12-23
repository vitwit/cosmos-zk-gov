// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/zkgov/v1beta1/zkgov.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type VoteOption int32

const (
	VoteOption_VOTE_OPTION_NO  VoteOption = 0
	VoteOption_VOTE_OPTION_YES VoteOption = 1
)

var VoteOption_name = map[int32]string{
	0: "VOTE_OPTION_NO",
	1: "VOTE_OPTION_YES",
}

var VoteOption_value = map[string]int32{
	"VOTE_OPTION_NO":  0,
	"VOTE_OPTION_YES": 1,
}

func (x VoteOption) String() string {
	return proto.EnumName(VoteOption_name, int32(x))
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78f302300f92129, []int{0}
}

// commitment
type Commitment struct {
	Commitment   string `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	CommitmentId uint64 `protobuf:"varint,2,opt,name=commitment_id,json=commitmentId,proto3" json:"commitment_id,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78f302300f92129, []int{0}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Commitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return m.Size()
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

func (m *Commitment) GetCommitment() string {
	if m != nil {
		return m.Commitment
	}
	return ""
}

func (m *Commitment) GetCommitmentId() uint64 {
	if m != nil {
		return m.CommitmentId
	}
	return 0
}

// user
type User struct {
	Userid       uint64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Commitment   []byte `protobuf:"bytes,2,opt,name=commitment,proto3" json:"commitment,omitempty"`
	RandomNumber uint64 `protobuf:"varint,3,opt,name=random_number,json=randomNumber,proto3" json:"random_number,omitempty"`
	Nullifier    []byte `protobuf:"bytes,4,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78f302300f92129, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *User) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *User) GetRandomNumber() uint64 {
	if m != nil {
		return m.RandomNumber
	}
	return 0
}

func (m *User) GetNullifier() []byte {
	if m != nil {
		return m.Nullifier
	}
	return nil
}

func init() {
	proto.RegisterEnum("sdk.zkgov.v1beta1.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterType((*Commitment)(nil), "sdk.zkgov.v1beta1.Commitment")
	proto.RegisterType((*User)(nil), "sdk.zkgov.v1beta1.User")
}

func init() { proto.RegisterFile("sdk/zkgov/v1beta1/zkgov.proto", fileDescriptor_e78f302300f92129) }

var fileDescriptor_e78f302300f92129 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0x02, 0x41,
	0x18, 0xc7, 0x77, 0x6a, 0x11, 0xfc, 0xb0, 0xb2, 0x09, 0x62, 0x0f, 0x35, 0x88, 0x5d, 0x24, 0xd0,
	0x45, 0xa2, 0x17, 0x28, 0x3c, 0x78, 0x71, 0xcb, 0x4c, 0xa8, 0xcb, 0xb2, 0xeb, 0x4c, 0x3a, 0xac,
	0x33, 0x23, 0x33, 0xb3, 0x52, 0xde, 0x7a, 0x83, 0x1e, 0xab, 0xa3, 0xc7, 0x8e, 0xa1, 0x2f, 0x12,
	0xee, 0x4a, 0x2b, 0x1e, 0x7f, 0xbf, 0x0f, 0xbe, 0x3f, 0xfc, 0xe0, 0xd2, 0xd0, 0xc4, 0x5f, 0x24,
	0x63, 0x35, 0xf7, 0xe7, 0xed, 0x98, 0xd9, 0xa8, 0x9d, 0x53, 0x6b, 0xa6, 0x95, 0x55, 0xf8, 0xd4,
	0xd0, 0xa4, 0x95, 0x8b, 0xed, 0xb9, 0xfe, 0x08, 0x70, 0xaf, 0x84, 0xe0, 0x56, 0x30, 0x69, 0x31,
	0x01, 0x18, 0xfd, 0x93, 0x87, 0x6a, 0xa8, 0x51, 0xee, 0xef, 0x18, 0x7c, 0x05, 0x47, 0x05, 0x85,
	0x9c, 0x7a, 0x07, 0x35, 0xd4, 0x70, 0xfb, 0x95, 0x42, 0x76, 0x69, 0xfd, 0x13, 0x81, 0xfb, 0x6c,
	0x98, 0xc6, 0xe7, 0x50, 0x4a, 0x0d, 0xd3, 0x9c, 0x66, 0x9f, 0xdc, 0xfe, 0x96, 0xf6, 0x56, 0x36,
	0x2f, 0x2a, 0xfb, 0x2b, 0x3a, 0x92, 0x54, 0x89, 0x50, 0xa6, 0x22, 0x66, 0xda, 0x3b, 0xcc, 0x57,
	0x72, 0xd9, 0xcb, 0x1c, 0xbe, 0x80, 0xb2, 0x4c, 0xa7, 0x53, 0xfe, 0xc6, 0x99, 0xf6, 0xdc, 0xec,
	0x47, 0x21, 0xae, 0x6f, 0x01, 0x86, 0xca, 0xb2, 0x60, 0x66, 0xb9, 0x92, 0x18, 0xc3, 0xf1, 0x30,
	0x18, 0x74, 0xc2, 0xe0, 0x61, 0xd0, 0x0d, 0x7a, 0x61, 0x2f, 0xa8, 0x3a, 0xf8, 0x0c, 0x4e, 0x76,
	0xdd, 0x4b, 0xe7, 0xa9, 0x8a, 0xee, 0xba, 0xdf, 0x2b, 0x82, 0x96, 0x2b, 0x82, 0x7e, 0x57, 0x04,
	0x7d, 0xad, 0x89, 0xb3, 0x5c, 0x13, 0xe7, 0x67, 0x4d, 0x9c, 0x57, 0x7f, 0xcc, 0xed, 0x24, 0x8d,
	0x5b, 0x23, 0x25, 0xfc, 0x39, 0x37, 0x93, 0x68, 0xda, 0x4c, 0x22, 0x29, 0x23, 0x7f, 0xb1, 0x09,
	0xde, 0xdc, 0x14, 0x7f, 0xdf, 0x96, 0xb7, 0x1f, 0x33, 0x66, 0xe2, 0x52, 0x96, 0xfc, 0xe6, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0xcb, 0x76, 0x5a, 0x93, 0x01, 0x00, 0x00,
}

func (m *Commitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommitmentId != 0 {
		i = encodeVarintZkgov(dAtA, i, uint64(m.CommitmentId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Commitment) > 0 {
		i -= len(m.Commitment)
		copy(dAtA[i:], m.Commitment)
		i = encodeVarintZkgov(dAtA, i, uint64(len(m.Commitment)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nullifier) > 0 {
		i -= len(m.Nullifier)
		copy(dAtA[i:], m.Nullifier)
		i = encodeVarintZkgov(dAtA, i, uint64(len(m.Nullifier)))
		i--
		dAtA[i] = 0x22
	}
	if m.RandomNumber != 0 {
		i = encodeVarintZkgov(dAtA, i, uint64(m.RandomNumber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Commitment) > 0 {
		i -= len(m.Commitment)
		copy(dAtA[i:], m.Commitment)
		i = encodeVarintZkgov(dAtA, i, uint64(len(m.Commitment)))
		i--
		dAtA[i] = 0x12
	}
	if m.Userid != 0 {
		i = encodeVarintZkgov(dAtA, i, uint64(m.Userid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintZkgov(dAtA []byte, offset int, v uint64) int {
	offset -= sovZkgov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Commitment)
	if l > 0 {
		n += 1 + l + sovZkgov(uint64(l))
	}
	if m.CommitmentId != 0 {
		n += 1 + sovZkgov(uint64(m.CommitmentId))
	}
	return n
}

func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Userid != 0 {
		n += 1 + sovZkgov(uint64(m.Userid))
	}
	l = len(m.Commitment)
	if l > 0 {
		n += 1 + l + sovZkgov(uint64(l))
	}
	if m.RandomNumber != 0 {
		n += 1 + sovZkgov(uint64(m.RandomNumber))
	}
	l = len(m.Nullifier)
	if l > 0 {
		n += 1 + l + sovZkgov(uint64(l))
	}
	return n
}

func sovZkgov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZkgov(x uint64) (n int) {
	return sovZkgov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZkgov
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
			return fmt.Errorf("proto: Commitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
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
				return ErrInvalidLengthZkgov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZkgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitmentId", wireType)
			}
			m.CommitmentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitmentId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipZkgov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZkgov
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZkgov
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			m.Userid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Userid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthZkgov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZkgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitment = append(m.Commitment[:0], dAtA[iNdEx:postIndex]...)
			if m.Commitment == nil {
				m.Commitment = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomNumber", wireType)
			}
			m.RandomNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nullifier", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkgov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthZkgov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZkgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nullifier = append(m.Nullifier[:0], dAtA[iNdEx:postIndex]...)
			if m.Nullifier == nil {
				m.Nullifier = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZkgov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZkgov
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
func skipZkgov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZkgov
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
					return 0, ErrIntOverflowZkgov
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
					return 0, ErrIntOverflowZkgov
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
				return 0, ErrInvalidLengthZkgov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupZkgov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthZkgov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthZkgov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZkgov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupZkgov = fmt.Errorf("proto: unexpected end of group")
)
