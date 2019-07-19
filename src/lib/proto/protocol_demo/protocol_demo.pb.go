// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol_demo.proto

package protocol_demo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloWorldReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldReq) Reset()         { *m = HelloWorldReq{} }
func (m *HelloWorldReq) String() string { return proto.CompactTextString(m) }
func (*HelloWorldReq) ProtoMessage()    {}
func (*HelloWorldReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95c63798f1a0ff5, []int{0}
}
func (m *HelloWorldReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloWorldReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloWorldReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloWorldReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldReq.Merge(m, src)
}
func (m *HelloWorldReq) XXX_Size() int {
	return m.Size()
}
func (m *HelloWorldReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldReq proto.InternalMessageInfo

func (m *HelloWorldReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type HelloWorldRsp struct {
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRsp) Reset()         { *m = HelloWorldRsp{} }
func (m *HelloWorldRsp) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRsp) ProtoMessage()    {}
func (*HelloWorldRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95c63798f1a0ff5, []int{1}
}
func (m *HelloWorldRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloWorldRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloWorldRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloWorldRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRsp.Merge(m, src)
}
func (m *HelloWorldRsp) XXX_Size() int {
	return m.Size()
}
func (m *HelloWorldRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRsp proto.InternalMessageInfo

func (m *HelloWorldRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloWorldReq)(nil), "protocol_demo.HelloWorldReq")
	proto.RegisterType((*HelloWorldRsp)(nil), "protocol_demo.HelloWorldRsp")
}

func init() { proto.RegisterFile("protocol_demo.proto", fileDescriptor_d95c63798f1a0ff5) }

var fileDescriptor_d95c63798f1a0ff5 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0x89, 0x4f, 0x49, 0xcd, 0xcd, 0xd7, 0x03, 0xf3, 0x84, 0x78, 0x51, 0x04,
	0x95, 0xe4, 0xb9, 0x78, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0xc3, 0xf3, 0x8b, 0x72, 0x52, 0x82, 0x52,
	0x0b, 0x85, 0xf8, 0xb8, 0x98, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x98, 0x3c,
	0x53, 0x94, 0x34, 0x51, 0x14, 0x14, 0x17, 0x08, 0x49, 0x70, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xd1, 0x5c, 0x5c, 0x08,
	0xa5, 0x42, 0xbe, 0x5c, 0xbc, 0xee, 0xa9, 0x25, 0x48, 0x02, 0x32, 0x7a, 0xa8, 0xee, 0x41, 0xb1,
	0x57, 0x0a, 0x8f, 0x6c, 0x71, 0x81, 0x12, 0x83, 0x93, 0xe6, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x78, 0x4e, 0x66, 0x92,
	0x3e, 0x58, 0x93, 0x3e, 0x8a, 0xd6, 0x24, 0x36, 0x30, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x95, 0xa6, 0x6c, 0xa5, 0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldClient interface {
	GetHelloWorld(ctx context.Context, in *HelloWorldReq, opts ...grpc.CallOption) (*HelloWorldRsp, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) GetHelloWorld(ctx context.Context, in *HelloWorldReq, opts ...grpc.CallOption) (*HelloWorldRsp, error) {
	out := new(HelloWorldRsp)
	err := c.cc.Invoke(ctx, "/protocol_demo.HelloWorld/GetHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
type HelloWorldServer interface {
	GetHelloWorld(context.Context, *HelloWorldReq) (*HelloWorldRsp, error)
}

// UnimplementedHelloWorldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (*UnimplementedHelloWorldServer) GetHelloWorld(ctx context.Context, req *HelloWorldReq) (*HelloWorldRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelloWorld not implemented")
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_GetHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GetHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol_demo.HelloWorld/GetHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GetHelloWorld(ctx, req.(*HelloWorldReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol_demo.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHelloWorld",
			Handler:    _HelloWorld_GetHelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol_demo.proto",
}

func (m *HelloWorldReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloWorldReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocolDemo(dAtA, i, uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HelloWorldRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloWorldRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtocolDemo(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintProtocolDemo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HelloWorldReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProtocolDemo(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HelloWorldRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovProtocolDemo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocolDemo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocolDemo(x uint64) (n int) {
	return sovProtocolDemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HelloWorldReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolDemo
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
			return fmt.Errorf("proto: HelloWorldReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloWorldReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocolDemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocolDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloWorldRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocolDemo
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
			return fmt.Errorf("proto: HelloWorldRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloWorldRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocolDemo
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
				return ErrInvalidLengthProtocolDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocolDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocolDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocolDemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocolDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocolDemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocolDemo
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
					return 0, ErrIntOverflowProtocolDemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocolDemo
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
				return 0, ErrInvalidLengthProtocolDemo
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProtocolDemo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocolDemo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtocolDemo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProtocolDemo
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtocolDemo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocolDemo   = fmt.Errorf("proto: integer overflow")
)
