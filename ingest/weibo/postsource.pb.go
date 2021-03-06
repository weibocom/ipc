// Code generated by protoc-gen-gogo.
// source: postsource.proto
// DO NOT EDIT!

package weibo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PostSource struct {
	// Types that are valid to be assigned to IdPresent:
	//	*PostSource_Id
	IdPresent isPostSource_IdPresent `protobuf_oneof:"id_present"`
	// Types that are valid to be assigned to NamePresent:
	//	*PostSource_Name
	NamePresent isPostSource_NamePresent `protobuf_oneof:"name_present"`
	// Types that are valid to be assigned to UrlPresent:
	//	*PostSource_Url
	UrlPresent isPostSource_UrlPresent `protobuf_oneof:"url_present"`
	// Types that are valid to be assigned to AppkeyPresent:
	//	*PostSource_Appkey
	AppkeyPresent isPostSource_AppkeyPresent `protobuf_oneof:"appkey_present"`
}

func (m *PostSource) Reset()                    { *m = PostSource{} }
func (m *PostSource) String() string            { return proto.CompactTextString(m) }
func (*PostSource) ProtoMessage()               {}
func (*PostSource) Descriptor() ([]byte, []int) { return fileDescriptorPostsource, []int{0} }

type isPostSource_IdPresent interface {
	isPostSource_IdPresent()
	MarshalTo([]byte) (int, error)
	Size() int
}
type isPostSource_NamePresent interface {
	isPostSource_NamePresent()
	MarshalTo([]byte) (int, error)
	Size() int
}
type isPostSource_UrlPresent interface {
	isPostSource_UrlPresent()
	MarshalTo([]byte) (int, error)
	Size() int
}
type isPostSource_AppkeyPresent interface {
	isPostSource_AppkeyPresent()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PostSource_Id struct {
	Id uint32 `protobuf:"varint,501,opt,name=id,proto3,oneof"`
}
type PostSource_Name struct {
	Name string `protobuf:"bytes,502,opt,name=name,proto3,oneof"`
}
type PostSource_Url struct {
	Url string `protobuf:"bytes,503,opt,name=url,proto3,oneof"`
}
type PostSource_Appkey struct {
	Appkey string `protobuf:"bytes,504,opt,name=appkey,proto3,oneof"`
}

func (*PostSource_Id) isPostSource_IdPresent()         {}
func (*PostSource_Name) isPostSource_NamePresent()     {}
func (*PostSource_Url) isPostSource_UrlPresent()       {}
func (*PostSource_Appkey) isPostSource_AppkeyPresent() {}

func (m *PostSource) GetIdPresent() isPostSource_IdPresent {
	if m != nil {
		return m.IdPresent
	}
	return nil
}
func (m *PostSource) GetNamePresent() isPostSource_NamePresent {
	if m != nil {
		return m.NamePresent
	}
	return nil
}
func (m *PostSource) GetUrlPresent() isPostSource_UrlPresent {
	if m != nil {
		return m.UrlPresent
	}
	return nil
}
func (m *PostSource) GetAppkeyPresent() isPostSource_AppkeyPresent {
	if m != nil {
		return m.AppkeyPresent
	}
	return nil
}

func (m *PostSource) GetId() uint32 {
	if x, ok := m.GetIdPresent().(*PostSource_Id); ok {
		return x.Id
	}
	return 0
}

func (m *PostSource) GetName() string {
	if x, ok := m.GetNamePresent().(*PostSource_Name); ok {
		return x.Name
	}
	return ""
}

func (m *PostSource) GetUrl() string {
	if x, ok := m.GetUrlPresent().(*PostSource_Url); ok {
		return x.Url
	}
	return ""
}

func (m *PostSource) GetAppkey() string {
	if x, ok := m.GetAppkeyPresent().(*PostSource_Appkey); ok {
		return x.Appkey
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PostSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PostSource_OneofMarshaler, _PostSource_OneofUnmarshaler, _PostSource_OneofSizer, []interface{}{
		(*PostSource_Id)(nil),
		(*PostSource_Name)(nil),
		(*PostSource_Url)(nil),
		(*PostSource_Appkey)(nil),
	}
}

func _PostSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PostSource)
	// id_present
	switch x := m.IdPresent.(type) {
	case *PostSource_Id:
		_ = b.EncodeVarint(501<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Id))
	case nil:
	default:
		return fmt.Errorf("PostSource.IdPresent has unexpected type %T", x)
	}
	// name_present
	switch x := m.NamePresent.(type) {
	case *PostSource_Name:
		_ = b.EncodeVarint(502<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("PostSource.NamePresent has unexpected type %T", x)
	}
	// url_present
	switch x := m.UrlPresent.(type) {
	case *PostSource_Url:
		_ = b.EncodeVarint(503<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Url)
	case nil:
	default:
		return fmt.Errorf("PostSource.UrlPresent has unexpected type %T", x)
	}
	// appkey_present
	switch x := m.AppkeyPresent.(type) {
	case *PostSource_Appkey:
		_ = b.EncodeVarint(504<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Appkey)
	case nil:
	default:
		return fmt.Errorf("PostSource.AppkeyPresent has unexpected type %T", x)
	}
	return nil
}

func _PostSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PostSource)
	switch tag {
	case 501: // id_present.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.IdPresent = &PostSource_Id{uint32(x)}
		return true, err
	case 502: // name_present.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.NamePresent = &PostSource_Name{x}
		return true, err
	case 503: // url_present.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.UrlPresent = &PostSource_Url{x}
		return true, err
	case 504: // appkey_present.appkey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AppkeyPresent = &PostSource_Appkey{x}
		return true, err
	default:
		return false, nil
	}
}

func _PostSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PostSource)
	// id_present
	switch x := m.IdPresent.(type) {
	case *PostSource_Id:
		n += proto.SizeVarint(501<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Id))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// name_present
	switch x := m.NamePresent.(type) {
	case *PostSource_Name:
		n += proto.SizeVarint(502<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// url_present
	switch x := m.UrlPresent.(type) {
	case *PostSource_Url:
		n += proto.SizeVarint(503<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// appkey_present
	switch x := m.AppkeyPresent.(type) {
	case *PostSource_Appkey:
		n += proto.SizeVarint(504<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Appkey)))
		n += len(x.Appkey)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*PostSource)(nil), "weibo.PostSource")
}
func (m *PostSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IdPresent != nil {
		nn1, err := m.IdPresent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.NamePresent != nil {
		nn2, err := m.NamePresent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.UrlPresent != nil {
		nn3, err := m.UrlPresent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	if m.AppkeyPresent != nil {
		nn4, err := m.AppkeyPresent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *PostSource_Id) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa8
	i++
	dAtA[i] = 0x1f
	i++
	i = encodeVarintPostsource(dAtA, i, uint64(m.Id))
	return i, nil
}
func (m *PostSource_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xb2
	i++
	dAtA[i] = 0x1f
	i++
	i = encodeVarintPostsource(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func (m *PostSource_Url) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xba
	i++
	dAtA[i] = 0x1f
	i++
	i = encodeVarintPostsource(dAtA, i, uint64(len(m.Url)))
	i += copy(dAtA[i:], m.Url)
	return i, nil
}
func (m *PostSource_Appkey) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xc2
	i++
	dAtA[i] = 0x1f
	i++
	i = encodeVarintPostsource(dAtA, i, uint64(len(m.Appkey)))
	i += copy(dAtA[i:], m.Appkey)
	return i, nil
}
func encodeFixed64Postsource(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Postsource(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPostsource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PostSource) Size() (n int) {
	var l int
	_ = l
	if m.IdPresent != nil {
		n += m.IdPresent.Size()
	}
	if m.NamePresent != nil {
		n += m.NamePresent.Size()
	}
	if m.UrlPresent != nil {
		n += m.UrlPresent.Size()
	}
	if m.AppkeyPresent != nil {
		n += m.AppkeyPresent.Size()
	}
	return n
}

func (m *PostSource_Id) Size() (n int) {
	var l int
	_ = l
	n += 2 + sovPostsource(uint64(m.Id))
	return n
}
func (m *PostSource_Name) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 2 + l + sovPostsource(uint64(l))
	return n
}
func (m *PostSource_Url) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	n += 2 + l + sovPostsource(uint64(l))
	return n
}
func (m *PostSource_Appkey) Size() (n int) {
	var l int
	_ = l
	l = len(m.Appkey)
	n += 2 + l + sovPostsource(uint64(l))
	return n
}

func sovPostsource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPostsource(x uint64) (n int) {
	return sovPostsource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PostSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostsource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PostSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 501:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IdPresent = &PostSource_Id{v}
		case 502:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamePresent = &PostSource_Name{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 503:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UrlPresent = &PostSource_Url{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 504:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Appkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostsource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostsource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppkeyPresent = &PostSource_Appkey{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostsource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPostsource
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
func skipPostsource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPostsource
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
					return 0, ErrIntOverflowPostsource
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
					return 0, ErrIntOverflowPostsource
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPostsource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPostsource
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
				next, err := skipPostsource(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthPostsource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPostsource   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("postsource.proto", fileDescriptorPostsource) }

var fileDescriptorPostsource = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc8, 0x2f, 0x2e,
	0x29, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x4f,
	0xcd, 0x4c, 0xca, 0x57, 0x5a, 0xca, 0xc8, 0xc5, 0x15, 0x90, 0x5f, 0x5c, 0x12, 0x0c, 0x96, 0x13,
	0x12, 0xe4, 0x62, 0xca, 0x4c, 0x91, 0xf8, 0xca, 0xac, 0xc0, 0xa8, 0xc1, 0xeb, 0xc1, 0x10, 0xc4,
	0x94, 0x99, 0x22, 0x24, 0xca, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xf1, 0x0d, 0x24, 0xc8, 0xe9,
	0xc1, 0x18, 0x04, 0xe6, 0x0a, 0x09, 0x73, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x7c, 0x87, 0x88, 0x32,
	0x05, 0x81, 0x78, 0x42, 0x92, 0x5c, 0x6c, 0x89, 0x05, 0x05, 0xd9, 0xa9, 0x95, 0x12, 0x3f, 0x20,
	0xe2, 0xcc, 0x41, 0x50, 0x01, 0x27, 0x1e, 0x2e, 0xae, 0xcc, 0x94, 0xf8, 0x82, 0xa2, 0xd4, 0xe2,
	0xd4, 0xbc, 0x12, 0x27, 0x3e, 0x2e, 0x1e, 0x90, 0x29, 0x70, 0x3e, 0x2f, 0x17, 0x77, 0x69, 0x51,
	0x0e, 0x9c, 0x2b, 0xc0, 0xc5, 0x07, 0xd1, 0x06, 0x17, 0xb1, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0xe0, 0x92, 0x4c, 0xce, 0xd3,
	0x2b, 0xce, 0xcc, 0x4b, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x4b, 0x49, 0x2c, 0x49, 0x84, 0x78, 0x2c,
	0xa9, 0x34, 0xcd, 0x89, 0x0f, 0xe1, 0xa3, 0xf0, 0xa2, 0xc4, 0x82, 0x24, 0x36, 0xb0, 0x8c, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0xdb, 0xe8, 0x84, 0x06, 0x01, 0x00, 0x00,
}
