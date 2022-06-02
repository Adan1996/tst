// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: wallet/wallet.proto

package wallet

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

type WalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId int64   `protobuf:"varint,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Amount   float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WalletRequest) Reset() {
	*x = WalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletRequest) ProtoMessage() {}

func (x *WalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletRequest.ProtoReflect.Descriptor instead.
func (*WalletRequest) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *WalletRequest) GetWalletId() int64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *WalletRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type WalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId int64   `protobuf:"varint,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Amount   float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WalletResponse) Reset() {
	*x = WalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletResponse) ProtoMessage() {}

func (x *WalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletResponse.ProtoReflect.Descriptor instead.
func (*WalletResponse) Descriptor() ([]byte, []int) {
	return file_wallet_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *WalletResponse) GetWalletId() int64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *WalletResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_wallet_wallet_proto protoreflect.FileDescriptor

var file_wallet_wallet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0e, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x3b, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x0e,
	0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1b, 0x5a, 0x19, 0x73, 0x79, 0x61, 0x68, 0x64, 0x61, 0x6e, 0x2e, 0x69, 0x64, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x69, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_wallet_proto_rawDescOnce sync.Once
	file_wallet_wallet_proto_rawDescData = file_wallet_wallet_proto_rawDesc
)

func file_wallet_wallet_proto_rawDescGZIP() []byte {
	file_wallet_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_wallet_proto_rawDescData)
	})
	return file_wallet_wallet_proto_rawDescData
}

var file_wallet_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wallet_wallet_proto_goTypes = []interface{}{
	(*WalletRequest)(nil),  // 0: WalletRequest
	(*WalletResponse)(nil), // 1: WalletResponse
}
var file_wallet_wallet_proto_depIdxs = []int32{
	0, // 0: Wallet.FindWalletById:input_type -> WalletRequest
	1, // 1: Wallet.FindWalletById:output_type -> WalletResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_wallet_proto_init() }
func file_wallet_wallet_proto_init() {
	if File_wallet_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletRequest); i {
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
		file_wallet_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletResponse); i {
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
			RawDescriptor: file_wallet_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_wallet_proto_msgTypes,
	}.Build()
	File_wallet_wallet_proto = out.File
	file_wallet_wallet_proto_rawDesc = nil
	file_wallet_wallet_proto_goTypes = nil
	file_wallet_wallet_proto_depIdxs = nil
}
