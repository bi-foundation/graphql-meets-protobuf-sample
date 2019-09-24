// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models.proto

package models

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/bi-foundation/protobuf-graphql-extension/graphqlproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_graphql_go_graphql "github.com/graphql-go/graphql"
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

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}

type Person struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string       `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *PhoneNumber `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Person.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return m.Size()
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhone() *PhoneNumber {
	if m != nil {
		return m.Phone
	}
	return nil
}

type PhoneNumber struct {
	Number               string    `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=models.PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}
func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(m, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return m.Size()
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *PhoneNumber) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_MOBILE
}

func init() {
	proto.RegisterEnum("models.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Person)(nil), "models.Person")
	proto.RegisterType((*PhoneNumber)(nil), "models.PhoneNumber")
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_0b5431a010549573) }

var fileDescriptor_0b5431a010549573 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x7d, 0x31, 0x0d, 0xf6, 0x55, 0x4a, 0x1d, 0x45, 0x82, 0xc8, 0x58, 0x0a, 0x42, 0x55,
	0xda, 0x42, 0xbd, 0x41, 0xa1, 0xa0, 0xd8, 0x36, 0x25, 0x08, 0xae, 0x13, 0x33, 0x6d, 0x06, 0x32,
	0x33, 0x31, 0x7f, 0xc0, 0x5e, 0xc7, 0x95, 0x47, 0x70, 0xe9, 0xd2, 0xa5, 0x47, 0xd0, 0x9c, 0xc2,
	0xa5, 0x64, 0x12, 0x45, 0x77, 0xdf, 0xf7, 0x7b, 0x3f, 0xde, 0x1b, 0x06, 0x77, 0x85, 0x0a, 0x58,
	0x94, 0x0e, 0xe3, 0x44, 0x65, 0x8a, 0x58, 0x55, 0x3b, 0x5a, 0xac, 0x79, 0x16, 0xe6, 0xfe, 0xf0,
	0x5e, 0x89, 0x91, 0xcf, 0x07, 0x2b, 0x95, 0xcb, 0xc0, 0xcb, 0xb8, 0x92, 0x23, 0xed, 0xf9, 0xf9,
	0x6a, 0xb0, 0x4e, 0xbc, 0x38, 0x7c, 0x88, 0x06, 0xec, 0x31, 0x63, 0x32, 0x2d, 0x47, 0x35, 0xd1,
	0xc6, 0x4f, 0xa9, 0xf6, 0xf6, 0x04, 0x5a, 0x4b, 0x96, 0xa4, 0x4a, 0x12, 0x82, 0xa6, 0xf4, 0x04,
	0xb3, 0xa1, 0x0b, 0xfd, 0xa6, 0xab, 0x33, 0x69, 0xa3, 0xc1, 0x03, 0xdb, 0xe8, 0x42, 0xbf, 0xe1,
	0x1a, 0x3c, 0x20, 0x07, 0xd8, 0x60, 0xc2, 0xe3, 0x91, 0xbd, 0xad, 0xa5, 0xaa, 0x90, 0x33, 0x6c,
	0xc4, 0xa1, 0x92, 0xcc, 0x36, 0xbb, 0xd0, 0x6f, 0x8d, 0xf7, 0x87, 0xf5, 0xcb, 0x97, 0x25, 0x5c,
	0xe4, 0xc2, 0x67, 0x89, 0x5b, 0x19, 0xbd, 0x19, 0xb6, 0xfe, 0x50, 0x72, 0x88, 0x96, 0xd4, 0xa9,
	0xbe, 0x5a, 0x37, 0x72, 0x8a, 0x66, 0xb6, 0x89, 0x99, 0xbe, 0xdc, 0x1e, 0xef, 0xfd, 0x5b, 0x78,
	0xbb, 0x89, 0x99, 0xab, 0xc7, 0xe7, 0x17, 0xd8, 0xfc, 0x45, 0x04, 0xd1, 0x9a, 0x3b, 0x93, 0xeb,
	0xd9, 0xb4, 0xb3, 0x45, 0x76, 0xd0, 0xbc, 0x72, 0xe6, 0xd3, 0x0e, 0x94, 0xe9, 0xce, 0x71, 0x6f,
	0x3a, 0xc6, 0xe4, 0xf8, 0xeb, 0x93, 0xc2, 0x73, 0x41, 0xe1, 0xa5, 0xa0, 0xf0, 0x56, 0x50, 0x78,
	0x2f, 0x28, 0x7c, 0x14, 0x14, 0x5e, 0x9f, 0x4e, 0xc0, 0xb7, 0xf4, 0x77, 0x5c, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0x47, 0x4b, 0x73, 0x76, 0x01, 0x00, 0x00,
}

func (this *Person) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Person)
	if !ok {
		that2, ok := that.(Person)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	if !this.Phone.Equal(that1.Phone) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PhoneNumber) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PhoneNumber)
	if !ok {
		that2, ok := that.(PhoneNumber)
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
	if this.Number != that1.Number {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

var GraphQLPhoneTypeEnum *github_com_graphql_go_graphql.Enum

type PersonGetter interface {
	GetPerson() *Person
}

var GraphQLPersonType *github_com_graphql_go_graphql.Object

type PhoneNumberGetter interface {
	GetPhoneNumber() *PhoneNumber
}

var GraphQLPhoneNumberType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLPhoneTypeEnum = github_com_graphql_go_graphql.NewEnum(github_com_graphql_go_graphql.EnumConfig{
		Name: "PhoneType",
		Values: github_com_graphql_go_graphql.EnumValueConfigMap{
			"MOBILE": &github_com_graphql_go_graphql.EnumValueConfig{
				Value: 0,
			},
			"HOME": &github_com_graphql_go_graphql.EnumValueConfig{
				Value: 1,
			},
			"WORK": &github_com_graphql_go_graphql.EnumValueConfig{
				Value: 2,
			},
		},
	})
	GraphQLPersonType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "Person",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"name": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Person)
						if ok {
							return obj.Name, nil
						}
						inter, ok := p.Source.(PersonGetter)
						if ok {
							face := inter.GetPerson()
							if face == nil {
								return nil, nil
							}
							return face.Name, nil
						}
						return nil, fmt.Errorf("field name not resolved")
					},
				},
				"id": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Int,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Person)
						if ok {
							return obj.Id, nil
						}
						inter, ok := p.Source.(PersonGetter)
						if ok {
							face := inter.GetPerson()
							if face == nil {
								return nil, nil
							}
							return face.Id, nil
						}
						return nil, fmt.Errorf("field id not resolved")
					},
				},
				"email": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Person)
						if ok {
							return obj.Email, nil
						}
						inter, ok := p.Source.(PersonGetter)
						if ok {
							face := inter.GetPerson()
							if face == nil {
								return nil, nil
							}
							return face.Email, nil
						}
						return nil, fmt.Errorf("field email not resolved")
					},
				},
				"phone": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLPhoneNumberType,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*Person)
						if ok {
							if obj.Phone == nil {
								return nil, nil
							}
							return obj.GetPhone(), nil
						}
						inter, ok := p.Source.(PersonGetter)
						if ok {
							face := inter.GetPerson()
							if face == nil {
								return nil, nil
							}
							if face.Phone == nil {
								return nil, nil
							}
							return face.GetPhone(), nil
						}
						return nil, fmt.Errorf("field phone not resolved")
					},
				},
			}
		}),
	})
	GraphQLPhoneNumberType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "PhoneNumber",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"number": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.String,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*PhoneNumber)
						if ok {
							return obj.Number, nil
						}
						inter, ok := p.Source.(PhoneNumberGetter)
						if ok {
							face := inter.GetPhoneNumber()
							if face == nil {
								return nil, nil
							}
							return face.Number, nil
						}
						return nil, fmt.Errorf("field number not resolved")
					},
				},
				"type": &github_com_graphql_go_graphql.Field{
					Type:        GraphQLPhoneTypeEnum,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*PhoneNumber)
						if ok {
							return int(PhoneType_value[obj.Type.String()]), nil
						}
						inter, ok := p.Source.(PhoneNumberGetter)
						if ok {
							face := inter.GetPhoneNumber()
							if face == nil {
								return nil, nil
							}
							return int(PhoneType_value[face.Type.String()]), nil
						}
						return nil, fmt.Errorf("field type not resolved")
					},
				},
			}
		}),
	})
}
func (m *Person) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Person) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Person) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Phone != nil {
		{
			size, err := m.Phone.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModels(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PhoneNumber) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PhoneNumber) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PhoneNumber) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Number) > 0 {
		i -= len(m.Number)
		copy(dAtA[i:], m.Number)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Number)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedPerson(r randyModels, easy bool) *Person {
	this := &Person{}
	this.Name = string(randStringModels(r))
	this.Id = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Id *= -1
	}
	this.Email = string(randStringModels(r))
	if r.Intn(5) != 0 {
		this.Phone = NewPopulatedPhoneNumber(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedModels(r, 5)
	}
	return this
}

func NewPopulatedPhoneNumber(r randyModels, easy bool) *PhoneNumber {
	this := &PhoneNumber{}
	this.Number = string(randStringModels(r))
	this.Type = PhoneType([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedModels(r, 3)
	}
	return this
}

type randyModels interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneModels(r randyModels) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringModels(r randyModels) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneModels(r)
	}
	return string(tmps)
}
func randUnrecognizedModels(r randyModels, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldModels(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldModels(dAtA []byte, r randyModels, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateModels(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateModels(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateModels(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateModels(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateModels(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateModels(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateModels(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Person) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovModels(uint64(m.Id))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.Phone != nil {
		l = m.Phone.Size()
		n += 1 + l + sovModels(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PhoneNumber) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Number)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovModels(uint64(m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Person) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Person: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Person: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Phone == nil {
				m.Phone = &PhoneNumber{}
			}
			if err := m.Phone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *PhoneNumber) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: PhoneNumber: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PhoneNumber: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Number = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PhoneType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthModels
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModels
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
				next, err := skipModels(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthModels
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
	ErrInvalidLengthModels = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels   = fmt.Errorf("proto: integer overflow")
)