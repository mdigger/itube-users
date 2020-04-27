// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokens.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// поддерживаемые типы токенов
type TokenType int32

const (
	EMAIL    TokenType = 0
	PASSWORD TokenType = 1
)

var TokenType_name = map[int32]string{
	0: "EMAIL",
	1: "PASSWORD",
}

var TokenType_value = map[string]int32{
	"EMAIL":    0,
	"PASSWORD": 1,
}

func (x TokenType) String() string {
	return proto.EnumName(TokenType_name, int32(x))
}

func (TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7213d78cc820f18a, []int{0}
}

// VerifyRequest используется для изменения запроса на проверку почтового адреса
// пользователя или для замены пароля. В данном случае domain влияет на
// формируемую ссылку для проверки токена и на быбор шаблона письма для
// отправки.
type VerifyRequest struct {
	// домен
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// логин пользователя
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// тип проверки
	Type TokenType `protobuf:"varint,3,opt,name=type,proto3,enum=itube.users.TokenType" json:"type,omitempty"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7213d78cc820f18a, []int{0}
}
func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return m.Size()
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

// TokenInfo описывает данные для проверки почтового адреса или сброса пароля
// по токену. Физически токен не привязан к домену и, чисто теоретически,
// может быть подтвержден на любом сайте. Тип проверки является чисто
// информационным и не влияет на уникальность токена.
type TokenInfo struct {
	// домен
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// полученный токен
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// тип проверки
	Type TokenType `protobuf:"varint,3,opt,name=type,proto3,enum=itube.users.TokenType" json:"type,omitempty"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7213d78cc820f18a, []int{1}
}
func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(m, src)
}
func (m *TokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("itube.users.TokenType", TokenType_name, TokenType_value)
	golang_proto.RegisterEnum("itube.users.TokenType", TokenType_name, TokenType_value)
	proto.RegisterType((*VerifyRequest)(nil), "itube.users.VerifyRequest")
	golang_proto.RegisterType((*VerifyRequest)(nil), "itube.users.VerifyRequest")
	proto.RegisterType((*TokenInfo)(nil), "itube.users.TokenInfo")
	golang_proto.RegisterType((*TokenInfo)(nil), "itube.users.TokenInfo")
}

func init() { proto.RegisterFile("tokens.proto", fileDescriptor_7213d78cc820f18a) }
func init() { golang_proto.RegisterFile("tokens.proto", fileDescriptor_7213d78cc820f18a) }

var fileDescriptor_7213d78cc820f18a = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x9b, 0xe9, 0xea, 0x16, 0xa7, 0xcc, 0x1c, 0x64, 0x14, 0x89, 0x63, 0x78, 0x18, 0x83,
	0xb5, 0xb8, 0x81, 0x47, 0x61, 0x43, 0x91, 0x81, 0xa2, 0x74, 0xf3, 0x0f, 0xde, 0x5a, 0x97, 0xd5,
	0xb0, 0xb5, 0xa9, 0x6d, 0xea, 0xe8, 0xc1, 0xbb, 0x47, 0x3f, 0x92, 0xc7, 0x1d, 0x77, 0xf4, 0xa6,
	0x6b, 0xbf, 0x88, 0x2c, 0x1d, 0xb2, 0x8a, 0x82, 0x9e, 0xf2, 0xbe, 0xbf, 0xf7, 0x49, 0x1e, 0x9e,
	0x37, 0xb0, 0xc0, 0xd9, 0x90, 0x38, 0xbe, 0xea, 0x7a, 0x8c, 0x33, 0xb4, 0x4e, 0x79, 0x60, 0x12,
	0x35, 0xf0, 0x89, 0xe7, 0x2b, 0x70, 0x7e, 0x24, 0x03, 0xa5, 0x6e, 0x51, 0x7e, 0x1f, 0x98, 0xea,
	0x1d, 0xb3, 0x35, 0x8b, 0x59, 0x4c, 0x13, 0xd8, 0x0c, 0x06, 0xa2, 0x13, 0x8d, 0xa8, 0x16, 0xf2,
	0x83, 0x25, 0xb9, 0x3d, 0xa6, 0x7c, 0xc8, 0xc6, 0x9a, 0xc5, 0xea, 0x62, 0x58, 0x7f, 0x34, 0x46,
	0xb4, 0x6f, 0x70, 0xe6, 0xf9, 0xda, 0x57, 0x99, 0xdc, 0xab, 0x84, 0x70, 0xe3, 0x8a, 0x78, 0x74,
	0x10, 0xea, 0xe4, 0x21, 0x20, 0x3e, 0x47, 0x18, 0xca, 0x7d, 0x66, 0x1b, 0xd4, 0x29, 0x81, 0x32,
	0xa8, 0xe6, 0xdb, 0x72, 0xf4, 0xbe, 0x9b, 0xb9, 0x01, 0xfa, 0x82, 0xa2, 0x1d, 0x98, 0x25, 0xb6,
	0x41, 0x47, 0xa5, 0x4c, 0x6a, 0x9c, 0x40, 0x54, 0x83, 0xab, 0x3c, 0x74, 0x49, 0x69, 0xa5, 0x0c,
	0xaa, 0x9b, 0x8d, 0x6d, 0x75, 0x29, 0x9d, 0xda, 0x9b, 0xe7, 0xee, 0x85, 0x2e, 0xd1, 0x85, 0xa6,
	0x12, 0xc0, 0xbc, 0x40, 0x1d, 0x67, 0xc0, 0xfe, 0x62, 0x2b, 0xf6, 0xf6, 0xdd, 0x56, 0xc0, 0xff,
	0xd8, 0xd6, 0xf6, 0x16, 0xb6, 0x73, 0x84, 0xf2, 0x30, 0x7b, 0x7c, 0xd6, 0xea, 0x9c, 0x16, 0x25,
	0x54, 0x80, 0xb9, 0x8b, 0x56, 0xb7, 0x7b, 0x7d, 0xae, 0x1f, 0x15, 0x41, 0xe3, 0x09, 0xca, 0x42,
	0xe5, 0xa3, 0x43, 0x98, 0x3b, 0x21, 0x0e, 0xf1, 0x0c, 0x4e, 0x90, 0x92, 0x7a, 0x39, 0xb5, 0x38,
	0xe5, 0x07, 0x57, 0x91, 0xac, 0x09, 0xe5, 0x44, 0x88, 0x7e, 0x51, 0x28, 0x5b, 0x29, 0x7e, 0xe9,
	0x13, 0xaf, 0xbd, 0x3f, 0x99, 0x61, 0x69, 0x3a, 0xc3, 0xd2, 0x24, 0xc2, 0x60, 0x1a, 0x61, 0xf0,
	0x11, 0x61, 0xf0, 0x1c, 0x63, 0xe9, 0x25, 0xc6, 0xd2, 0x6b, 0x8c, 0xc1, 0x34, 0xc6, 0xd2, 0x5b,
	0x8c, 0xa5, 0xdb, 0x35, 0x77, 0x68, 0x69, 0x86, 0x4b, 0x4d, 0x59, 0x7c, 0x68, 0xf3, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x3c, 0x8b, 0xe4, 0x60, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokensClient is the client API for Tokens service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensClient interface {
	// Generate создает запрос для проверки адреса email пользователя или
	// сброса пароля. При вызове сервер отправляет соответствующее письмо
	// на email адрес пользователя с токеном для верификации.
	// Повторный вызов с теми же значениями параметров заменяет токен на новый,
	// а действие старого отменяет.
	//
	// Возвращает ошибки:
	//  - InvalidArgument - неверный формат данных входящего запроса
	//  - Internal - внутренние ошибки
	Generate(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*TokenInfo, error)
	// Verify проверяет токен и возвращает зарегистрированного пользователя.
	// Если токен неверен, то возвращается ошибка NotFound. После проверки
	// токен автоматически удаляется и повторное его использование невозможно.
	//
	// Так же автоматически подтверждает почтовый адрес, через который был
	// отправлен данный токен.
	//
	// Возвращает ошибки:
	//  - NotFound - пользователь не зарегистрирован
	//  - InvalidArgument - неверный формат данных входящего запроса
	//  - Internal - внутренние ошибки
	Verify(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*User, error)
}

type tokensClient struct {
	cc *grpc.ClientConn
}

func NewTokensClient(cc *grpc.ClientConn) TokensClient {
	return &tokensClient{cc}
}

func (c *tokensClient) Generate(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*TokenInfo, error) {
	out := new(TokenInfo)
	err := c.cc.Invoke(ctx, "/itube.users.Tokens/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) Verify(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/itube.users.Tokens/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensServer is the server API for Tokens service.
type TokensServer interface {
	// Generate создает запрос для проверки адреса email пользователя или
	// сброса пароля. При вызове сервер отправляет соответствующее письмо
	// на email адрес пользователя с токеном для верификации.
	// Повторный вызов с теми же значениями параметров заменяет токен на новый,
	// а действие старого отменяет.
	//
	// Возвращает ошибки:
	//  - InvalidArgument - неверный формат данных входящего запроса
	//  - Internal - внутренние ошибки
	Generate(context.Context, *VerifyRequest) (*TokenInfo, error)
	// Verify проверяет токен и возвращает зарегистрированного пользователя.
	// Если токен неверен, то возвращается ошибка NotFound. После проверки
	// токен автоматически удаляется и повторное его использование невозможно.
	//
	// Так же автоматически подтверждает почтовый адрес, через который был
	// отправлен данный токен.
	//
	// Возвращает ошибки:
	//  - NotFound - пользователь не зарегистрирован
	//  - InvalidArgument - неверный формат данных входящего запроса
	//  - Internal - внутренние ошибки
	Verify(context.Context, *TokenInfo) (*User, error)
}

// UnimplementedTokensServer can be embedded to have forward compatible implementations.
type UnimplementedTokensServer struct {
}

func (*UnimplementedTokensServer) Generate(ctx context.Context, req *VerifyRequest) (*TokenInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedTokensServer) Verify(ctx context.Context, req *TokenInfo) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterTokensServer(s *grpc.Server, srv TokensServer) {
	s.RegisterService(&_Tokens_serviceDesc, srv)
}

func _Tokens_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itube.users.Tokens/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).Generate(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itube.users.Tokens/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).Verify(ctx, req.(*TokenInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokens_serviceDesc = grpc.ServiceDesc{
	ServiceName: "itube.users.Tokens",
	HandlerType: (*TokensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Tokens_Generate_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Tokens_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokens.proto",
}

func (m *VerifyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTokens(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTokens(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokens(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokens(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VerifyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTokens(uint64(m.Type))
	}
	return n
}

func (m *TokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTokens(uint64(m.Type))
	}
	return n
}

func sovTokens(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokens(x uint64) (n int) {
	return sovTokens(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokens
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
			return fmt.Errorf("proto: VerifyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= TokenType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokens(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokens
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTokens
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
func (m *TokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokens
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
			return fmt.Errorf("proto: TokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= TokenType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokens(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokens
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTokens
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
func skipTokens(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokens
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
					return 0, ErrIntOverflowTokens
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
					return 0, ErrIntOverflowTokens
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
				return 0, ErrInvalidLengthTokens
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokens
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokens
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokens        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokens          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokens = fmt.Errorf("proto: unexpected end of group")
)