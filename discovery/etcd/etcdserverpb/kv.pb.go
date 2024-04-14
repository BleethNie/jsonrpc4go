// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: discovery/etcd/etcdserverpb/kv.proto

package etcdserverpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Lease int64  `protobuf:"varint,3,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {
	// Do nothing.
}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP(), []int{0}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PutRequest) GetLease() int64 {
	if x != nil {
		return x.Lease
	}
	return 0
}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {

}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP(), []int{1}
}

type RangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRequest) ProtoMessage() {}

func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRequest.ProtoReflect.Descriptor instead.
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP(), []int{2}
}

func (x *RangeRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kvs []*KeyValue `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (x *RangeResponse) Reset() {
	*x = RangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeResponse) ProtoMessage() {}

func (x *RangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeResponse.ProtoReflect.Descriptor instead.
func (*RangeResponse) Descriptor() ([]byte, []int) {
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP(), []int{3}
}

func (x *RangeResponse) GetKvs() []*KeyValue {
	if x != nil {
		return x.Kvs
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Lease int64  `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[4]
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
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP(), []int{4}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValue) GetLease() int64 {
	if x != nil {
		return x.Lease
	}
	return 0
}

var File_discovery_etcd_etcdserverpb_kv_proto protoreflect.FileDescriptor

var file_discovery_etcd_etcdserverpb_kv_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x74, 0x63, 0x64,
	0x2f, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6b, 0x76,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x70, 0x62, 0x22, 0x4a, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x39, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x22, 0x48, 0x0a, 0x08,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x32, 0x82, 0x01, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x3a, 0x0a,
	0x03, 0x50, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x65, 0x74, 0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x65, 0x74,
	0x63, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_discovery_etcd_etcdserverpb_kv_proto_rawDescOnce sync.Once
	file_discovery_etcd_etcdserverpb_kv_proto_rawDescData = file_discovery_etcd_etcdserverpb_kv_proto_rawDesc
)

func file_discovery_etcd_etcdserverpb_kv_proto_rawDescGZIP() []byte {
	file_discovery_etcd_etcdserverpb_kv_proto_rawDescOnce.Do(func() {
		file_discovery_etcd_etcdserverpb_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_etcd_etcdserverpb_kv_proto_rawDescData)
	})
	return file_discovery_etcd_etcdserverpb_kv_proto_rawDescData
}

var file_discovery_etcd_etcdserverpb_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_discovery_etcd_etcdserverpb_kv_proto_goTypes = []interface{}{
	(*PutRequest)(nil),    // 0: etcdserverpb.PutRequest
	(*PutResponse)(nil),   // 1: etcdserverpb.PutResponse
	(*RangeRequest)(nil),  // 2: etcdserverpb.RangeRequest
	(*RangeResponse)(nil), // 3: etcdserverpb.RangeResponse
	(*KeyValue)(nil),      // 4: etcdserverpb.KeyValue
}
var file_discovery_etcd_etcdserverpb_kv_proto_depIdxs = []int32{
	4, // 0: etcdserverpb.RangeResponse.kvs:type_name -> etcdserverpb.KeyValue
	0, // 1: etcdserverpb.KV.Put:input_type -> etcdserverpb.PutRequest
	2, // 2: etcdserverpb.KV.Range:input_type -> etcdserverpb.RangeRequest
	1, // 3: etcdserverpb.KV.Put:output_type -> etcdserverpb.PutResponse
	3, // 4: etcdserverpb.KV.Range:output_type -> etcdserverpb.RangeResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_discovery_etcd_etcdserverpb_kv_proto_init() }
func file_discovery_etcd_etcdserverpb_kv_proto_init() {
	if File_discovery_etcd_etcdserverpb_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
		file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeRequest); i {
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
		file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeResponse); i {
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
		file_discovery_etcd_etcdserverpb_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discovery_etcd_etcdserverpb_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discovery_etcd_etcdserverpb_kv_proto_goTypes,
		DependencyIndexes: file_discovery_etcd_etcdserverpb_kv_proto_depIdxs,
		MessageInfos:      file_discovery_etcd_etcdserverpb_kv_proto_msgTypes,
	}.Build()
	File_discovery_etcd_etcdserverpb_kv_proto = out.File
	file_discovery_etcd_etcdserverpb_kv_proto_rawDesc = nil
	file_discovery_etcd_etcdserverpb_kv_proto_goTypes = nil
	file_discovery_etcd_etcdserverpb_kv_proto_depIdxs = nil
}
