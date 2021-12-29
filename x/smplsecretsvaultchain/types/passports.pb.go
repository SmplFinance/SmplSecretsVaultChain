// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smplsecretsvaultchain/passports.proto

package types

import (
	fmt "fmt"
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

type Passports struct {
	PassportStorageMap  map[string]*PassportStorage `protobuf:"bytes,1,rep,name=passportStorageMap,proto3" json:"passportStorageMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TestPhrase          string                      `protobuf:"bytes,2,opt,name=testPhrase,proto3" json:"testPhrase,omitempty"`
	EncryptedTestPhrase string                      `protobuf:"bytes,3,opt,name=encryptedTestPhrase,proto3" json:"encryptedTestPhrase,omitempty"`
}

func (m *Passports) Reset()         { *m = Passports{} }
func (m *Passports) String() string { return proto.CompactTextString(m) }
func (*Passports) ProtoMessage()    {}
func (*Passports) Descriptor() ([]byte, []int) {
	return fileDescriptor_30093bff18060f68, []int{0}
}
func (m *Passports) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Passports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Passports.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Passports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Passports.Merge(m, src)
}
func (m *Passports) XXX_Size() int {
	return m.Size()
}
func (m *Passports) XXX_DiscardUnknown() {
	xxx_messageInfo_Passports.DiscardUnknown(m)
}

var xxx_messageInfo_Passports proto.InternalMessageInfo

func (m *Passports) GetPassportStorageMap() map[string]*PassportStorage {
	if m != nil {
		return m.PassportStorageMap
	}
	return nil
}

func (m *Passports) GetTestPhrase() string {
	if m != nil {
		return m.TestPhrase
	}
	return ""
}

func (m *Passports) GetEncryptedTestPhrase() string {
	if m != nil {
		return m.EncryptedTestPhrase
	}
	return ""
}

func init() {
	proto.RegisterType((*Passports)(nil), "SmplFinance.smplsecretsvaultchain.smplsecretsvaultchain.Passports")
	proto.RegisterMapType((map[string]*PassportStorage)(nil), "SmplFinance.smplsecretsvaultchain.smplsecretsvaultchain.Passports.PassportStorageMapEntry")
}

func init() {
	proto.RegisterFile("smplsecretsvaultchain/passports.proto", fileDescriptor_30093bff18060f68)
}

var fileDescriptor_30093bff18060f68 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xce, 0x2d, 0xc8,
	0x29, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x29, 0x2e, 0x4b, 0x2c, 0xcd, 0x29, 0x49, 0xce, 0x48, 0xcc,
	0xcc, 0xd3, 0x2f, 0x48, 0x2c, 0x2e, 0x2e, 0xc8, 0x2f, 0x2a, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x32, 0x0f, 0xce, 0x2d, 0xc8, 0x71, 0xcb, 0xcc, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0xc3,
	0xaa, 0x05, 0xbb, 0xa8, 0x94, 0x0e, 0x7e, 0xf3, 0xe3, 0x8b, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53,
	0x21, 0xd6, 0x28, 0xfd, 0x61, 0xe2, 0xe2, 0x0c, 0x80, 0x59, 0x2d, 0xd4, 0xc5, 0xc8, 0x25, 0x04,
	0x53, 0x18, 0x0c, 0x51, 0xe7, 0x9b, 0x58, 0x20, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x14, 0xa5,
	0x47, 0xa6, 0x93, 0xf4, 0xe0, 0x16, 0xc0, 0x59, 0x08, 0xc3, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83,
	0xb0, 0xd8, 0x2a, 0x24, 0xc7, 0xc5, 0x55, 0x92, 0x5a, 0x5c, 0x12, 0x90, 0x51, 0x94, 0x58, 0x9c,
	0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x64, 0xc0, 0x25, 0x9c, 0x9a, 0x97,
	0x5c, 0x54, 0x59, 0x50, 0x92, 0x9a, 0x12, 0x82, 0x50, 0xc8, 0x0c, 0x56, 0x88, 0x4d, 0x4a, 0xaa,
	0x9f, 0x91, 0x4b, 0x1c, 0x87, 0x0b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0xc1,
	0xba, 0x41, 0x4c, 0xa1, 0x38, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0x52, 0x88, 0xd5, 0xdc, 0x46, 0x1e,
	0x14, 0x7b, 0x1f, 0x6a, 0x65, 0x10, 0xc4, 0x58, 0x2b, 0x26, 0x0b, 0x46, 0xa7, 0x94, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xb2, 0x18, 0xcc, 0x0e, 0x86, 0x58, 0x11, 0x06, 0xb2, 0xc2, 0x19,
	0x1c, 0xbb, 0x15, 0xfa, 0xd8, 0x63, 0xbd, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0xd7,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xbd, 0xe8, 0x4e, 0x7b, 0x02, 0x00, 0x00,
}

func (m *Passports) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Passports) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Passports) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedTestPhrase) > 0 {
		i -= len(m.EncryptedTestPhrase)
		copy(dAtA[i:], m.EncryptedTestPhrase)
		i = encodeVarintPassports(dAtA, i, uint64(len(m.EncryptedTestPhrase)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TestPhrase) > 0 {
		i -= len(m.TestPhrase)
		copy(dAtA[i:], m.TestPhrase)
		i = encodeVarintPassports(dAtA, i, uint64(len(m.TestPhrase)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PassportStorageMap) > 0 {
		for k := range m.PassportStorageMap {
			v := m.PassportStorageMap[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintPassports(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintPassports(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintPassports(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPassports(dAtA []byte, offset int, v uint64) int {
	offset -= sovPassports(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Passports) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PassportStorageMap) > 0 {
		for k, v := range m.PassportStorageMap {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovPassports(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovPassports(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovPassports(uint64(mapEntrySize))
		}
	}
	l = len(m.TestPhrase)
	if l > 0 {
		n += 1 + l + sovPassports(uint64(l))
	}
	l = len(m.EncryptedTestPhrase)
	if l > 0 {
		n += 1 + l + sovPassports(uint64(l))
	}
	return n
}

func sovPassports(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPassports(x uint64) (n int) {
	return sovPassports(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Passports) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassports
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
			return fmt.Errorf("proto: Passports: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Passports: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PassportStorageMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassports
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
				return ErrInvalidLengthPassports
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPassports
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PassportStorageMap == nil {
				m.PassportStorageMap = make(map[string]*PassportStorage)
			}
			var mapkey string
			var mapvalue *PassportStorage
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPassports
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPassports
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPassports
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthPassports
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPassports
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthPassports
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthPassports
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &PassportStorage{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPassports(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthPassports
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PassportStorageMap[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestPhrase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassports
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
				return ErrInvalidLengthPassports
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPassports
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestPhrase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedTestPhrase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassports
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
				return ErrInvalidLengthPassports
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPassports
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedTestPhrase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPassports(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPassports
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
func skipPassports(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPassports
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
					return 0, ErrIntOverflowPassports
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
					return 0, ErrIntOverflowPassports
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
				return 0, ErrInvalidLengthPassports
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPassports
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPassports
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPassports        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPassports          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPassports = fmt.Errorf("proto: unexpected end of group")
)