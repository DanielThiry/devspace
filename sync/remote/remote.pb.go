// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package remote

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChangeType int32

const (
	ChangeType_CHANGE ChangeType = 0
	ChangeType_DELETE ChangeType = 1
)

var ChangeType_name = map[int32]string{
	0: "CHANGE",
	1: "DELETE",
}

var ChangeType_value = map[string]int32{
	"CHANGE": 0,
	"DELETE": 1,
}

func (x ChangeType) String() string {
	return proto.EnumName(ChangeType_name, int32(x))
}

func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{0}
}

type Watch struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Exclude              []string `protobuf:"bytes,2,rep,name=Exclude,proto3" json:"Exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Watch) Reset()         { *m = Watch{} }
func (m *Watch) String() string { return proto.CompactTextString(m) }
func (*Watch) ProtoMessage()    {}
func (*Watch) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{0}
}

func (m *Watch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Watch.Unmarshal(m, b)
}
func (m *Watch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Watch.Marshal(b, m, deterministic)
}
func (m *Watch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Watch.Merge(m, src)
}
func (m *Watch) XXX_Size() int {
	return xxx_messageInfo_Watch.Size(m)
}
func (m *Watch) XXX_DiscardUnknown() {
	xxx_messageInfo_Watch.DiscardUnknown(m)
}

var xxx_messageInfo_Watch proto.InternalMessageInfo

func (m *Watch) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Watch) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type ChangeAmount struct {
	Amount               int64    `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAmount) Reset()         { *m = ChangeAmount{} }
func (m *ChangeAmount) String() string { return proto.CompactTextString(m) }
func (*ChangeAmount) ProtoMessage()    {}
func (*ChangeAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{1}
}

func (m *ChangeAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAmount.Unmarshal(m, b)
}
func (m *ChangeAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAmount.Marshal(b, m, deterministic)
}
func (m *ChangeAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAmount.Merge(m, src)
}
func (m *ChangeAmount) XXX_Size() int {
	return xxx_messageInfo_ChangeAmount.Size(m)
}
func (m *ChangeAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAmount.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAmount proto.InternalMessageInfo

func (m *ChangeAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ChangeChunk struct {
	Changes              []*Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ChangeChunk) Reset()         { *m = ChangeChunk{} }
func (m *ChangeChunk) String() string { return proto.CompactTextString(m) }
func (*ChangeChunk) ProtoMessage()    {}
func (*ChangeChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{2}
}

func (m *ChangeChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeChunk.Unmarshal(m, b)
}
func (m *ChangeChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeChunk.Marshal(b, m, deterministic)
}
func (m *ChangeChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeChunk.Merge(m, src)
}
func (m *ChangeChunk) XXX_Size() int {
	return xxx_messageInfo_ChangeChunk.Size(m)
}
func (m *ChangeChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeChunk.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeChunk proto.InternalMessageInfo

func (m *ChangeChunk) GetChanges() []*Change {
	if m != nil {
		return m.Changes
	}
	return nil
}

type Change struct {
	ChangeType           ChangeType `protobuf:"varint,1,opt,name=ChangeType,proto3,enum=remote.ChangeType" json:"ChangeType,omitempty"`
	Path                 string     `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	MtimeUnix            int64      `protobuf:"varint,3,opt,name=MtimeUnix,proto3" json:"MtimeUnix,omitempty"`
	MtimeUnixNano        int64      `protobuf:"varint,4,opt,name=MtimeUnixNano,proto3" json:"MtimeUnixNano,omitempty"`
	Size                 int64      `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	IsDir                bool       `protobuf:"varint,6,opt,name=IsDir,proto3" json:"IsDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{3}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetChangeType() ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return ChangeType_CHANGE
}

func (m *Change) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Change) GetMtimeUnix() int64 {
	if m != nil {
		return m.MtimeUnix
	}
	return 0
}

func (m *Change) GetMtimeUnixNano() int64 {
	if m != nil {
		return m.MtimeUnixNano
	}
	return 0
}

func (m *Change) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Change) GetIsDir() bool {
	if m != nil {
		return m.IsDir
	}
	return false
}

type Paths struct {
	Paths                []string `protobuf:"bytes,1,rep,name=Paths,proto3" json:"Paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paths) Reset()         { *m = Paths{} }
func (m *Paths) String() string { return proto.CompactTextString(m) }
func (*Paths) ProtoMessage()    {}
func (*Paths) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{4}
}

func (m *Paths) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paths.Unmarshal(m, b)
}
func (m *Paths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paths.Marshal(b, m, deterministic)
}
func (m *Paths) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paths.Merge(m, src)
}
func (m *Paths) XXX_Size() int {
	return xxx_messageInfo_Paths.Size(m)
}
func (m *Paths) XXX_DiscardUnknown() {
	xxx_messageInfo_Paths.DiscardUnknown(m)
}

var xxx_messageInfo_Paths proto.InternalMessageInfo

func (m *Paths) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{5}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("remote.ChangeType", ChangeType_name, ChangeType_value)
	proto.RegisterType((*Watch)(nil), "remote.Watch")
	proto.RegisterType((*ChangeAmount)(nil), "remote.ChangeAmount")
	proto.RegisterType((*ChangeChunk)(nil), "remote.ChangeChunk")
	proto.RegisterType((*Change)(nil), "remote.Change")
	proto.RegisterType((*Paths)(nil), "remote.Paths")
	proto.RegisterType((*Chunk)(nil), "remote.Chunk")
	proto.RegisterType((*Empty)(nil), "remote.Empty")
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor_eefc82927d57d89b) }

var fileDescriptor_eefc82927d57d89b = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x8d, 0xea, 0xda, 0x8e, 0x6f, 0x93, 0x52, 0xb4, 0x32, 0x44, 0xd8, 0xc0, 0x13, 0x65, 0x98,
	0x3e, 0x84, 0xce, 0xa3, 0xec, 0xb9, 0x38, 0x66, 0x1b, 0x6c, 0x65, 0x68, 0x0d, 0x7b, 0x76, 0x53,
	0x31, 0x9b, 0xc4, 0x92, 0xb1, 0xe5, 0x2d, 0xd9, 0x0f, 0xda, 0xcf, 0xd8, 0x6f, 0x1b, 0x92, 0xec,
	0x24, 0x1e, 0xec, 0xed, 0x9c, 0x7b, 0xcf, 0xfd, 0x3a, 0x5c, 0x98, 0xd4, 0xbc, 0x94, 0x8a, 0xcf,
	0xab, 0x5a, 0x2a, 0x89, 0x3d, 0xcb, 0xe8, 0x2d, 0xb8, 0xdf, 0x32, 0xb5, 0xca, 0x31, 0x86, 0xd3,
	0x2f, 0x99, 0xca, 0x09, 0x0a, 0x51, 0x14, 0x30, 0x83, 0x31, 0x01, 0x3f, 0xdd, 0xae, 0x36, 0xed,
	0x13, 0x27, 0x27, 0xa1, 0x13, 0x05, 0xac, 0xa7, 0xf4, 0x35, 0x4c, 0x92, 0x3c, 0x13, 0xdf, 0xf9,
	0x5d, 0x29, 0x5b, 0xa1, 0xf0, 0x73, 0xf0, 0x2c, 0x32, 0xf5, 0x0e, 0xeb, 0x18, 0x7d, 0x07, 0x67,
	0x56, 0x97, 0xe4, 0xad, 0x58, 0xe3, 0x08, 0xfc, 0x95, 0xa1, 0x0d, 0x41, 0xa1, 0x13, 0x9d, 0xc5,
	0xe7, 0xf3, 0x6e, 0x2b, 0xab, 0x62, 0x7d, 0x9a, 0xfe, 0x41, 0xe0, 0xd9, 0x18, 0x8e, 0x01, 0x2c,
	0x7a, 0xd8, 0x55, 0xdc, 0xf4, 0x3f, 0x8f, 0xf1, 0xb0, 0x4e, 0x67, 0xd8, 0x91, 0x6a, 0x7f, 0xcd,
	0xc9, 0xd1, 0x35, 0x2f, 0x20, 0xf8, 0xac, 0x8a, 0x92, 0x2f, 0x45, 0xb1, 0x25, 0x8e, 0x59, 0xf3,
	0x10, 0xc0, 0x57, 0x30, 0xdd, 0x93, 0xfb, 0x4c, 0x48, 0x72, 0x6a, 0x14, 0xc3, 0xa0, 0xee, 0xfb,
	0xb5, 0xf8, 0xc5, 0x89, 0x6b, 0x92, 0x06, 0xe3, 0x4b, 0x70, 0x3f, 0x36, 0x8b, 0xa2, 0x26, 0x5e,
	0x88, 0xa2, 0x31, 0xb3, 0x84, 0xbe, 0x04, 0x57, 0x4f, 0x6d, 0x74, 0xda, 0x00, 0x73, 0x71, 0xc0,
	0x2c, 0xa1, 0xaf, 0xc0, 0xb5, 0x96, 0x10, 0xf0, 0x13, 0x29, 0x14, 0xef, 0xac, 0x9b, 0xb0, 0x9e,
	0x52, 0x1f, 0xdc, 0xb4, 0xac, 0xd4, 0xee, 0xfa, 0xea, 0xd8, 0x00, 0x0c, 0xe0, 0x25, 0x1f, 0xee,
	0xee, 0xdf, 0xa7, 0x17, 0x23, 0x8d, 0x17, 0xe9, 0xa7, 0xf4, 0x21, 0xbd, 0x40, 0xf1, 0x6f, 0x04,
	0xb0, 0x90, 0x3f, 0x45, 0xa3, 0x6a, 0x9e, 0x95, 0x78, 0x0e, 0x63, 0xcd, 0x36, 0x32, 0x7b, 0xc2,
	0xd3, 0xde, 0x2d, 0x33, 0x7b, 0x36, 0x3d, 0x98, 0xd7, 0x8a, 0x35, 0x1d, 0x45, 0xe8, 0x06, 0xe1,
	0x37, 0xe0, 0xdb, 0x21, 0xcd, 0x41, 0x6e, 0xc6, 0xcf, 0x9e, 0x0d, 0xbd, 0xee, 0x8a, 0x6e, 0x10,
	0xbe, 0xed, 0x9f, 0xa0, 0x49, 0xcc, 0x13, 0xfc, 0x53, 0x77, 0x39, 0xac, 0xeb, 0x3e, 0x62, 0x14,
	0x3f, 0xc2, 0x78, 0x59, 0x75, 0x5b, 0x5e, 0x83, 0xb7, 0xac, 0x86, 0x3b, 0x9a, 0xfe, 0xb3, 0x61,
	0x2f, 0xbd, 0xa3, 0xd6, 0x32, 0x5e, 0xca, 0x1f, 0xfc, 0xbf, 0xf7, 0xec, 0xb5, 0x8f, 0x9e, 0xf9,
	0xf2, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0xa8, 0x8d, 0xf9, 0xf5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownstreamClient is the client API for Downstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownstreamClient interface {
	Download(ctx context.Context, opts ...grpc.CallOption) (Downstream_DownloadClient, error)
	Changes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Downstream_ChangesClient, error)
	ChangesCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChangeAmount, error)
}

type downstreamClient struct {
	cc *grpc.ClientConn
}

func NewDownstreamClient(cc *grpc.ClientConn) DownstreamClient {
	return &downstreamClient{cc}
}

func (c *downstreamClient) Download(ctx context.Context, opts ...grpc.CallOption) (Downstream_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Downstream_serviceDesc.Streams[0], "/remote.Downstream/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &downstreamDownloadClient{stream}
	return x, nil
}

type Downstream_DownloadClient interface {
	Send(*Paths) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type downstreamDownloadClient struct {
	grpc.ClientStream
}

func (x *downstreamDownloadClient) Send(m *Paths) error {
	return x.ClientStream.SendMsg(m)
}

func (x *downstreamDownloadClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *downstreamClient) Changes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Downstream_ChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Downstream_serviceDesc.Streams[1], "/remote.Downstream/Changes", opts...)
	if err != nil {
		return nil, err
	}
	x := &downstreamChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Downstream_ChangesClient interface {
	Recv() (*ChangeChunk, error)
	grpc.ClientStream
}

type downstreamChangesClient struct {
	grpc.ClientStream
}

func (x *downstreamChangesClient) Recv() (*ChangeChunk, error) {
	m := new(ChangeChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *downstreamClient) ChangesCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChangeAmount, error) {
	out := new(ChangeAmount)
	err := c.cc.Invoke(ctx, "/remote.Downstream/ChangesCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownstreamServer is the server API for Downstream service.
type DownstreamServer interface {
	Download(Downstream_DownloadServer) error
	Changes(*Empty, Downstream_ChangesServer) error
	ChangesCount(context.Context, *Empty) (*ChangeAmount, error)
}

func RegisterDownstreamServer(s *grpc.Server, srv DownstreamServer) {
	s.RegisterService(&_Downstream_serviceDesc, srv)
}

func _Downstream_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DownstreamServer).Download(&downstreamDownloadServer{stream})
}

type Downstream_DownloadServer interface {
	Send(*Chunk) error
	Recv() (*Paths, error)
	grpc.ServerStream
}

type downstreamDownloadServer struct {
	grpc.ServerStream
}

func (x *downstreamDownloadServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *downstreamDownloadServer) Recv() (*Paths, error) {
	m := new(Paths)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Downstream_Changes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownstreamServer).Changes(m, &downstreamChangesServer{stream})
}

type Downstream_ChangesServer interface {
	Send(*ChangeChunk) error
	grpc.ServerStream
}

type downstreamChangesServer struct {
	grpc.ServerStream
}

func (x *downstreamChangesServer) Send(m *ChangeChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Downstream_ChangesCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownstreamServer).ChangesCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Downstream/ChangesCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownstreamServer).ChangesCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Downstream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.Downstream",
	HandlerType: (*DownstreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangesCount",
			Handler:    _Downstream_ChangesCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Downstream_Download_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Changes",
			Handler:       _Downstream_Changes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote.proto",
}

// UpstreamClient is the client API for Upstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpstreamClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Upstream_UploadClient, error)
	Remove(ctx context.Context, opts ...grpc.CallOption) (Upstream_RemoveClient, error)
}

type upstreamClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamClient(cc *grpc.ClientConn) UpstreamClient {
	return &upstreamClient{cc}
}

func (c *upstreamClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Upstream_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upstream_serviceDesc.Streams[0], "/remote.Upstream/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamUploadClient{stream}
	return x, nil
}

type Upstream_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type upstreamUploadClient struct {
	grpc.ClientStream
}

func (x *upstreamUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *upstreamUploadClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *upstreamClient) Remove(ctx context.Context, opts ...grpc.CallOption) (Upstream_RemoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upstream_serviceDesc.Streams[1], "/remote.Upstream/Remove", opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamRemoveClient{stream}
	return x, nil
}

type Upstream_RemoveClient interface {
	Send(*Paths) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type upstreamRemoveClient struct {
	grpc.ClientStream
}

func (x *upstreamRemoveClient) Send(m *Paths) error {
	return x.ClientStream.SendMsg(m)
}

func (x *upstreamRemoveClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UpstreamServer is the server API for Upstream service.
type UpstreamServer interface {
	Upload(Upstream_UploadServer) error
	Remove(Upstream_RemoveServer) error
}

func RegisterUpstreamServer(s *grpc.Server, srv UpstreamServer) {
	s.RegisterService(&_Upstream_serviceDesc, srv)
}

func _Upstream_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpstreamServer).Upload(&upstreamUploadServer{stream})
}

type Upstream_UploadServer interface {
	SendAndClose(*Empty) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type upstreamUploadServer struct {
	grpc.ServerStream
}

func (x *upstreamUploadServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *upstreamUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Upstream_Remove_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpstreamServer).Remove(&upstreamRemoveServer{stream})
}

type Upstream_RemoveServer interface {
	SendAndClose(*Empty) error
	Recv() (*Paths, error)
	grpc.ServerStream
}

type upstreamRemoveServer struct {
	grpc.ServerStream
}

func (x *upstreamRemoveServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *upstreamRemoveServer) Recv() (*Paths, error) {
	m := new(Paths)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Upstream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.Upstream",
	HandlerType: (*UpstreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Upstream_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Remove",
			Handler:       _Upstream_Remove_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "remote.proto",
}
