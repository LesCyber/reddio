// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v4.25.2
// source: statedb.proto

package goproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripodName string `protobuf:"bytes,1,opt,name=tripod_name,json=tripodName,proto3" json:"tripod_name,omitempty"`
	Key        []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statedb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_statedb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_statedb_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *Key) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripodName string `protobuf:"bytes,1,opt,name=tripod_name,json=tripodName,proto3" json:"tripod_name,omitempty"`
	Key        []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statedb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_statedb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_statedb_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *KeyValue) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *ValueResponse) Reset() {
	*x = ValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statedb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueResponse) ProtoMessage() {}

func (x *ValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statedb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueResponse.ProtoReflect.Descriptor instead.
func (*ValueResponse) Descriptor() ([]byte, []int) {
	return file_statedb_proto_rawDescGZIP(), []int{2}
}

func (x *ValueResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ValueResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type KeyByHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripodName string `protobuf:"bytes,1,opt,name=tripod_name,json=tripodName,proto3" json:"tripod_name,omitempty"`
	Key        []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	BlockHash  []byte `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *KeyByHash) Reset() {
	*x = KeyByHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statedb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyByHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyByHash) ProtoMessage() {}

func (x *KeyByHash) ProtoReflect() protoreflect.Message {
	mi := &file_statedb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyByHash.ProtoReflect.Descriptor instead.
func (*KeyByHash) Descriptor() ([]byte, []int) {
	return file_statedb_proto_rawDescGZIP(), []int{3}
}

func (x *KeyByHash) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *KeyByHash) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyByHash) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

var File_statedb_proto protoreflect.FileDescriptor

var file_statedb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x53, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a,
	0x0d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x5d, 0x0a,
	0x09, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72,
	0x69, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x32, 0xfa, 0x03, 0x0a,
	0x07, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x42, 0x12, 0x1b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x04, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x09, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x04, 0x2e, 0x4b, 0x65, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x12, 0x04, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x2c, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x0a, 0x2e, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x0e, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x04, 0x2e, 0x4b, 0x65,
	0x79, 0x1a, 0x0e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x08, 0x2e, 0x54, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x54, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3c, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x54, 0x78, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statedb_proto_rawDescOnce sync.Once
	file_statedb_proto_rawDescData = file_statedb_proto_rawDesc
)

func file_statedb_proto_rawDescGZIP() []byte {
	file_statedb_proto_rawDescOnce.Do(func() {
		file_statedb_proto_rawDescData = protoimpl.X.CompressGZIP(file_statedb_proto_rawDescData)
	})
	return file_statedb_proto_rawDescData
}

var file_statedb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_statedb_proto_goTypes = []interface{}{
	(*Key)(nil),             // 0: Key
	(*KeyValue)(nil),        // 1: KeyValue
	(*ValueResponse)(nil),   // 2: ValueResponse
	(*KeyByHash)(nil),       // 3: KeyByHash
	(*TxnHash)(nil),         // 4: TxnHash
	(*emptypb.Empty)(nil),   // 5: google.protobuf.Empty
	(*Bool)(nil),            // 6: Bool
	(*TxnHashResponse)(nil), // 7: TxnHashResponse
}
var file_statedb_proto_depIdxs = []int32{
	0,  // 0: StateDB.Get:input_type -> Key
	1,  // 1: StateDB.Set:input_type -> KeyValue
	0,  // 2: StateDB.Delete:input_type -> Key
	0,  // 3: StateDB.Exist:input_type -> Key
	3,  // 4: StateDB.GetByBlockHash:input_type -> KeyByHash
	0,  // 5: StateDB.GetFinalized:input_type -> Key
	4,  // 6: StateDB.StartBlock:input_type -> TxnHash
	5,  // 7: StateDB.Commit:input_type -> google.protobuf.Empty
	5,  // 8: StateDB.Discard:input_type -> google.protobuf.Empty
	5,  // 9: StateDB.DiscardAll:input_type -> google.protobuf.Empty
	5,  // 10: StateDB.NextTxn:input_type -> google.protobuf.Empty
	2,  // 11: StateDB.Get:output_type -> ValueResponse
	5,  // 12: StateDB.Set:output_type -> google.protobuf.Empty
	5,  // 13: StateDB.Delete:output_type -> google.protobuf.Empty
	6,  // 14: StateDB.Exist:output_type -> Bool
	2,  // 15: StateDB.GetByBlockHash:output_type -> ValueResponse
	2,  // 16: StateDB.GetFinalized:output_type -> ValueResponse
	5,  // 17: StateDB.StartBlock:output_type -> google.protobuf.Empty
	7,  // 18: StateDB.Commit:output_type -> TxnHashResponse
	5,  // 19: StateDB.Discard:output_type -> google.protobuf.Empty
	5,  // 20: StateDB.DiscardAll:output_type -> google.protobuf.Empty
	5,  // 21: StateDB.NextTxn:output_type -> google.protobuf.Empty
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_statedb_proto_init() }
func file_statedb_proto_init() {
	if File_statedb_proto != nil {
		return
	}
	file_base_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_statedb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statedb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statedb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statedb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyByHash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_statedb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statedb_proto_goTypes,
		DependencyIndexes: file_statedb_proto_depIdxs,
		MessageInfos:      file_statedb_proto_msgTypes,
	}.Build()
	File_statedb_proto = out.File
	file_statedb_proto_rawDesc = nil
	file_statedb_proto_goTypes = nil
	file_statedb_proto_depIdxs = nil
}
