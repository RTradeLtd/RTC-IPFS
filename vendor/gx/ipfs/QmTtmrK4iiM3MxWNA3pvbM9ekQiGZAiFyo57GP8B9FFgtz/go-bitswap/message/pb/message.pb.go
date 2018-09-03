// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package bitswap_message_pb

import proto "gx/ipfs/QmdxUuburamoF6zF9qjeQC4WYcWGbWuRmdLacMEsW8ioD8/gogo-protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Wantlist             *Message_Wantlist `protobuf:"bytes,1,opt,name=wantlist" json:"wantlist,omitempty"`
	Blocks               [][]byte          `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
	Payload              []*Message_Block  `protobuf:"bytes,3,rep,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_1e228ff77b8fb7b4, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetWantlist() *Message_Wantlist {
	if m != nil {
		return m.Wantlist
	}
	return nil
}

func (m *Message) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Message) GetPayload() []*Message_Block {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Message_Wantlist struct {
	Entries              []*Message_Wantlist_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	Full                 bool                      `protobuf:"varint,2,opt,name=full,proto3" json:"full,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Message_Wantlist) Reset()         { *m = Message_Wantlist{} }
func (m *Message_Wantlist) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist) ProtoMessage()    {}
func (*Message_Wantlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_1e228ff77b8fb7b4, []int{0, 0}
}
func (m *Message_Wantlist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Wantlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Wantlist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message_Wantlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist.Merge(dst, src)
}
func (m *Message_Wantlist) XXX_Size() int {
	return m.Size()
}
func (m *Message_Wantlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist proto.InternalMessageInfo

func (m *Message_Wantlist) GetEntries() []*Message_Wantlist_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Message_Wantlist) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type Message_Wantlist_Entry struct {
	Block                []byte   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Cancel               bool     `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Wantlist_Entry) Reset()         { *m = Message_Wantlist_Entry{} }
func (m *Message_Wantlist_Entry) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist_Entry) ProtoMessage()    {}
func (*Message_Wantlist_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_1e228ff77b8fb7b4, []int{0, 0, 0}
}
func (m *Message_Wantlist_Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Wantlist_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Wantlist_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message_Wantlist_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist_Entry.Merge(dst, src)
}
func (m *Message_Wantlist_Entry) XXX_Size() int {
	return m.Size()
}
func (m *Message_Wantlist_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist_Entry proto.InternalMessageInfo

func (m *Message_Wantlist_Entry) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Message_Wantlist_Entry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Message_Wantlist_Entry) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

type Message_Block struct {
	Prefix               []byte   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Block) Reset()         { *m = Message_Block{} }
func (m *Message_Block) String() string { return proto.CompactTextString(m) }
func (*Message_Block) ProtoMessage()    {}
func (*Message_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_1e228ff77b8fb7b4, []int{0, 1}
}
func (m *Message_Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Block.Merge(dst, src)
}
func (m *Message_Block) XXX_Size() int {
	return m.Size()
}
func (m *Message_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Block proto.InternalMessageInfo

func (m *Message_Block) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Message_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "bitswap.message.pb.Message")
	proto.RegisterType((*Message_Wantlist)(nil), "bitswap.message.pb.Message.Wantlist")
	proto.RegisterType((*Message_Wantlist_Entry)(nil), "bitswap.message.pb.Message.Wantlist.Entry")
	proto.RegisterType((*Message_Block)(nil), "bitswap.message.pb.Message.Block")
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Wantlist != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Wantlist.Size()))
		n1, err := m.Wantlist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Blocks) > 0 {
		for _, b := range m.Blocks {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMessage(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Payload) > 0 {
		for _, msg := range m.Payload {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintMessage(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Message_Wantlist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Wantlist) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMessage(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Full {
		dAtA[i] = 0x10
		i++
		if m.Full {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Message_Wantlist_Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Wantlist_Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Block) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Block)))
		i += copy(dAtA[i:], m.Block)
	}
	if m.Priority != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Priority))
	}
	if m.Cancel {
		dAtA[i] = 0x18
		i++
		if m.Cancel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Message_Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Block) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.Wantlist != nil {
		l = m.Wantlist.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Blocks) > 0 {
		for _, b := range m.Blocks {
			l = len(b)
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if len(m.Payload) > 0 {
		for _, e := range m.Payload {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Message_Wantlist) Size() (n int) {
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.Full {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Message_Wantlist_Entry) Size() (n int) {
	var l int
	_ = l
	l = len(m.Block)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovMessage(uint64(m.Priority))
	}
	if m.Cancel {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Message_Block) Size() (n int) {
	var l int
	_ = l
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wantlist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Wantlist == nil {
				m.Wantlist = &Message_Wantlist{}
			}
			if err := m.Wantlist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, make([]byte, postIndex-iNdEx))
			copy(m.Blocks[len(m.Blocks)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload, &Message_Block{})
			if err := m.Payload[len(m.Payload)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Message_Wantlist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Wantlist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wantlist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &Message_Wantlist_Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Full", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Full = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Message_Wantlist_Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Block = append(m.Block[:0], dAtA[iNdEx:postIndex]...)
			if m.Block == nil {
				m.Block = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Cancel = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Message_Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_1e228ff77b8fb7b4) }

var fileDescriptor_message_1e228ff77b8fb7b4 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xe6, 0x4f, 0x1b, 0xdd, 0xe6, 0x5f, 0x2c, 0x84, 0xac, 0x0c, 0x55, 0x40, 0x0c,
	0x11, 0x83, 0x87, 0x76, 0x64, 0x41, 0x15, 0x8c, 0x0c, 0x78, 0x61, 0x76, 0x52, 0x17, 0x59, 0x98,
	0x24, 0xb2, 0x8d, 0x4a, 0x9e, 0x82, 0xc7, 0xe1, 0x15, 0x18, 0x79, 0x04, 0x94, 0x27, 0x41, 0xb9,
	0x75, 0xb2, 0x20, 0x21, 0xb6, 0x7b, 0xac, 0xf3, 0x1d, 0x9f, 0x6b, 0xc3, 0xff, 0x67, 0xe5, 0x9c,
	0x7c, 0x54, 0xbc, 0xb5, 0x8d, 0x6f, 0x28, 0x2d, 0xb5, 0x77, 0x07, 0xd9, 0xf2, 0xe9, 0xb8, 0x3c,
	0x7f, 0x8b, 0x60, 0x71, 0x77, 0x94, 0xf4, 0x1a, 0x92, 0x83, 0xac, 0xbd, 0xd1, 0xce, 0x33, 0x92,
	0x93, 0x62, 0xb9, 0xbe, 0xe0, 0x3f, 0x11, 0x1e, 0xec, 0xfc, 0x21, 0x78, 0xc5, 0x44, 0xd1, 0x53,
	0x98, 0x97, 0xa6, 0xa9, 0x9e, 0x1c, 0x9b, 0xe5, 0x51, 0x91, 0x8a, 0xa0, 0xe8, 0x15, 0x2c, 0x5a,
	0xd9, 0x99, 0x46, 0xee, 0x58, 0x94, 0x47, 0xc5, 0x72, 0x7d, 0xf6, 0x5b, 0xf0, 0x76, 0x80, 0xc4,
	0x48, 0x64, 0xef, 0x04, 0x92, 0xf1, 0x2e, 0x7a, 0x03, 0x0b, 0x55, 0x7b, 0xab, 0x95, 0x63, 0x04,
	0x93, 0x2e, 0xff, 0x52, 0x91, 0xdf, 0xd6, 0xde, 0x76, 0x62, 0x44, 0x29, 0x85, 0x7f, 0xfb, 0x17,
	0x63, 0xd8, 0x2c, 0x27, 0x45, 0x22, 0x70, 0xce, 0xee, 0x21, 0x46, 0x17, 0x3d, 0x81, 0x18, 0x6b,
	0xe3, 0x1b, 0xa4, 0xe2, 0x28, 0x68, 0x06, 0x49, 0x6b, 0x75, 0x63, 0xb5, 0xef, 0x10, 0x8b, 0xc5,
	0xa4, 0x87, 0xb5, 0x2b, 0x59, 0x57, 0xca, 0xb0, 0x08, 0x03, 0x83, 0xca, 0x36, 0x10, 0xe3, 0x2e,
	0x83, 0xa1, 0xb5, 0x6a, 0xaf, 0x5f, 0x43, 0x66, 0x50, 0x43, 0x8f, 0x9d, 0xf4, 0x12, 0x03, 0x53,
	0x81, 0xf3, 0x36, 0xfd, 0xe8, 0x57, 0xe4, 0xb3, 0x5f, 0x91, 0xaf, 0x7e, 0x45, 0xca, 0x39, 0x7e,
	0xdd, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x95, 0x9b, 0xc1, 0xcb, 0x01, 0x00, 0x00,
}