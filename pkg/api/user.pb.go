// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// User описывает информацию о пользователе.
//
// При обновлении domain, verified и updated игнорируются. При авторизации
// или запросе информации о пользователе в поле domain возвращается тоже
// значение, что и было в запросе.
type User struct {
	// домен (возвращает тот, который был указан в запросе)
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// уникальный идентификатор пользователя
	UID string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// email-адрес пользователя
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// флаг, что email-адрес подтвержден
	Verified bool `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	// дата и время последнего обновления
	Updated *time.Time `protobuf:"bytes,5,opt,name=updated,proto3,stdtime" json:"updated,omitempty"`
	// расширенные свойства
	Properties *types.Struct `protobuf:"bytes,10,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

// RegInfo описывает дополнительную информацию, используемую при регистрации.
type RegInfo struct {
	// точка перехода
	Referer string `protobuf:"bytes,1,opt,name=referer,proto3" json:"referer,omitempty"`
	// маркетинговая информация (https://ru.wikipedia.org/wiki/UTM-метки)
	// желательно имена меток давать без префикса "utm_"
	UTM map[string]string `protobuf:"bytes,2,rep,name=utm,proto3" json:"utm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *RegInfo) Reset()         { *m = RegInfo{} }
func (m *RegInfo) String() string { return proto.CompactTextString(m) }
func (*RegInfo) ProtoMessage()    {}
func (*RegInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}
func (m *RegInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegInfo.Merge(m, src)
}
func (m *RegInfo) XXX_Size() int {
	return m.Size()
}
func (m *RegInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RegInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "itube.users.User")
	golang_proto.RegisterType((*User)(nil), "itube.users.User")
	proto.RegisterType((*RegInfo)(nil), "itube.users.RegInfo")
	golang_proto.RegisterType((*RegInfo)(nil), "itube.users.RegInfo")
	proto.RegisterMapType((map[string]string)(nil), "itube.users.RegInfo.UtmEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "itube.users.RegInfo.UtmEntry")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }
func init() { golang_proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x7d, 0x71, 0x9a, 0x84, 0xcb, 0x82, 0x4e, 0x48, 0x58, 0x56, 0x39, 0x5b, 0x5d, 0x88,
	0x90, 0x62, 0x8b, 0x20, 0x95, 0xaa, 0x63, 0x54, 0x86, 0x0e, 0x2c, 0xa6, 0x91, 0x2a, 0x36, 0xa7,
	0x7e, 0x36, 0xa7, 0xc4, 0x3e, 0x73, 0xbe, 0x4b, 0xd5, 0xff, 0x80, 0x31, 0x0b, 0x13, 0xff, 0x0c,
	0x63, 0xc7, 0x8c, 0x4c, 0x84, 0xda, 0xff, 0x08, 0xf2, 0xd9, 0x2e, 0x01, 0xb6, 0xf7, 0xbd, 0xef,
	0xfb, 0x3d, 0xcb, 0x9f, 0x8d, 0xb1, 0x2a, 0x40, 0x78, 0xb9, 0xe0, 0x92, 0x93, 0x31, 0x93, 0x6a,
	0x09, 0x5e, 0xbd, 0x29, 0xec, 0xe3, 0x84, 0xf3, 0x64, 0x0d, 0xbe, 0xb6, 0x96, 0x2a, 0xf6, 0x0b,
	0x29, 0xd4, 0x8d, 0x6c, 0xa2, 0xb6, 0xf3, 0xaf, 0x2b, 0x59, 0x0a, 0x85, 0x0c, 0xd3, 0xbc, 0x0d,
	0x4c, 0x13, 0x26, 0x3f, 0xa9, 0xa5, 0x77, 0xc3, 0x53, 0x3f, 0xe1, 0x09, 0xff, 0x93, 0xac, 0x95,
	0x16, 0x7a, 0x6a, 0xe3, 0xa7, 0x07, 0xf1, 0xf4, 0x96, 0xc9, 0x15, 0xbf, 0xf5, 0x13, 0x3e, 0xd5,
	0xe6, 0x74, 0x13, 0xae, 0x59, 0x14, 0x4a, 0x2e, 0x0a, 0xff, 0x71, 0x6c, 0xb8, 0x93, 0x6f, 0x3d,
	0xdc, 0x5f, 0x14, 0x20, 0x08, 0xc5, 0x83, 0x88, 0xa7, 0x21, 0xcb, 0x2c, 0xe4, 0xa2, 0xc9, 0x93,
	0xf9, 0xa0, 0xdc, 0x3b, 0xbd, 0x6b, 0x14, 0xb4, 0x5b, 0x72, 0x81, 0x4d, 0xc5, 0x22, 0xab, 0xa7,
	0xcd, 0x59, 0xf9, 0xd3, 0x31, 0x17, 0x97, 0x17, 0xe5, 0xde, 0x79, 0xf9, 0xca, 0x65, 0x99, 0xbe,
	0xea, 0xaa, 0x8c, 0x7d, 0x56, 0xe0, 0xb2, 0x08, 0x32, 0xc9, 0x62, 0x06, 0xc2, 0x8d, 0xb9, 0x48,
	0x43, 0x79, 0x8d, 0xb6, 0xa8, 0x1f, 0xd4, 0x38, 0x39, 0xc6, 0x47, 0x90, 0x86, 0x6c, 0x6d, 0x99,
	0x7f, 0x3d, 0xa4, 0x59, 0x12, 0x1b, 0x8f, 0x36, 0x20, 0x6a, 0x34, 0xb2, 0xfa, 0x2e, 0x9a, 0x8c,
	0x82, 0x47, 0x4d, 0xce, 0xf1, 0x50, 0xe5, 0x51, 0x28, 0x21, 0xb2, 0x8e, 0x5c, 0x34, 0x19, 0xcf,
	0x6c, 0xaf, 0xa9, 0xd0, 0xeb, 0x8a, 0xf1, 0xae, 0xba, 0x0a, 0xe7, 0xfd, 0xed, 0xde, 0x41, 0x41,
	0x07, 0x90, 0xb7, 0x18, 0xe7, 0x82, 0xe7, 0x20, 0x24, 0x83, 0xc2, 0xc2, 0x1a, 0x7f, 0xfe, 0x1f,
	0xfe, 0x41, 0x7f, 0x9f, 0xe0, 0x20, 0x7a, 0xf2, 0x15, 0xe1, 0x61, 0x00, 0xc9, 0x65, 0x16, 0x73,
	0x62, 0xe1, 0xa1, 0x80, 0x18, 0x04, 0x88, 0xa6, 0xa1, 0xa0, 0x93, 0xe4, 0x0c, 0x9b, 0x4a, 0xa6,
	0x56, 0xcf, 0x35, 0x27, 0xe3, 0xd9, 0x0b, 0xef, 0xe0, 0x27, 0xf0, 0x5a, 0xd8, 0x5b, 0xc8, 0xf4,
	0x5d, 0x26, 0xc5, 0xdd, 0x7c, 0xa8, 0x9b, 0xbb, 0x7a, 0x1f, 0xd4, 0x88, 0x7d, 0x8a, 0x47, 0x9d,
	0x43, 0x9e, 0x62, 0x73, 0x05, 0x77, 0xed, 0xed, 0x7a, 0x24, 0xcf, 0xf0, 0xd1, 0x26, 0x5c, 0x2b,
	0x68, 0x4a, 0x0f, 0x1a, 0x71, 0xde, 0x3b, 0x43, 0xf3, 0xd7, 0xf7, 0x0f, 0xd4, 0xd8, 0x3d, 0x50,
	0xe3, 0xbe, 0xa4, 0x68, 0x57, 0x52, 0xf4, 0xab, 0xa4, 0xe8, 0x4b, 0x45, 0x8d, 0x6d, 0x45, 0x8d,
	0xef, 0x15, 0x45, 0xbb, 0x8a, 0x1a, 0x3f, 0x2a, 0x6a, 0x7c, 0x1c, 0xe6, 0xab, 0xc4, 0x0f, 0x73,
	0xb6, 0x1c, 0xe8, 0xf7, 0x7c, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x61, 0x1e, 0x74, 0xb0,
	0x02, 0x00, 0x00,
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
	if m.Properties != nil {
		{
			size, err := m.Properties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUser(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Updated != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Updated):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintUser(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x2a
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UTM) > 0 {
		for k := range m.UTM {
			v := m.UTM[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintUser(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintUser(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintUser(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Referer) > 0 {
		i -= len(m.Referer)
		copy(dAtA[i:], m.Referer)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Referer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Verified {
		n += 2
	}
	if m.Updated != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Updated)
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Properties != nil {
		l = m.Properties.Size()
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *RegInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Referer)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if len(m.UTM) > 0 {
		for k, v := range m.UTM {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovUser(uint64(len(k))) + 1 + len(v) + sovUser(uint64(len(v)))
			n += mapEntrySize + 1 + sovUser(uint64(mapEntrySize))
		}
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Updated == nil {
				m.Updated = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Properties == nil {
				m.Properties = &types.Struct{}
			}
			if err := m.Properties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *RegInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: RegInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Referer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Referer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UTM", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UTM == nil {
				m.UTM = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
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
							return ErrIntOverflowUser
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
						return ErrInvalidLengthUser
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthUser
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUser
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthUser
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthUser
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipUser(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthUser
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.UTM[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
