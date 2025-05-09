// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: halo/evmstaking/keeper/evmstaking.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	types "github.com/omni-network/omni/octane/evmengine/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EVMEvent is an unparsed EVM event.
type EVMEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Event         *types.EVMEvent        `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EVMEvent) Reset() {
	*x = EVMEvent{}
	mi := &file_halo_evmstaking_keeper_evmstaking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EVMEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVMEvent) ProtoMessage() {}

func (x *EVMEvent) ProtoReflect() protoreflect.Message {
	mi := &file_halo_evmstaking_keeper_evmstaking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVMEvent.ProtoReflect.Descriptor instead.
func (*EVMEvent) Descriptor() ([]byte, []int) {
	return file_halo_evmstaking_keeper_evmstaking_proto_rawDescGZIP(), []int{0}
}

func (x *EVMEvent) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EVMEvent) GetEvent() *types.EVMEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_halo_evmstaking_keeper_evmstaking_proto protoreflect.FileDescriptor

const file_halo_evmstaking_keeper_evmstaking_proto_rawDesc = "" +
	"\n" +
	"'halo/evmstaking/keeper/evmstaking.proto\x12\x16halo.evmstaking.keeper\x1a\x17cosmos/orm/v1/orm.proto\x1a\x1foctane/evmengine/types/tx.proto\"d\n" +
	"\bEVMEvent\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x126\n" +
	"\x05event\x18\x02 \x01(\v2 .octane.evmengine.types.EVMEventR\x05event:\x10\xf2\x9eӎ\x03\n" +
	"\n" +
	"\x06\n" +
	"\x02id\x10\x01\x18\x01B\xdc\x01\n" +
	"\x1acom.halo.evmstaking.keeperB\x0fEvmstakingProtoP\x01Z3github.com/omni-network/omni/halo/evmstaking/keeper\xa2\x02\x03HEK\xaa\x02\x16Halo.Evmstaking.Keeper\xca\x02\x16Halo\\Evmstaking\\Keeper\xe2\x02\"Halo\\Evmstaking\\Keeper\\GPBMetadata\xea\x02\x18Halo::Evmstaking::Keeperb\x06proto3"

var (
	file_halo_evmstaking_keeper_evmstaking_proto_rawDescOnce sync.Once
	file_halo_evmstaking_keeper_evmstaking_proto_rawDescData []byte
)

func file_halo_evmstaking_keeper_evmstaking_proto_rawDescGZIP() []byte {
	file_halo_evmstaking_keeper_evmstaking_proto_rawDescOnce.Do(func() {
		file_halo_evmstaking_keeper_evmstaking_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_halo_evmstaking_keeper_evmstaking_proto_rawDesc), len(file_halo_evmstaking_keeper_evmstaking_proto_rawDesc)))
	})
	return file_halo_evmstaking_keeper_evmstaking_proto_rawDescData
}

var file_halo_evmstaking_keeper_evmstaking_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_halo_evmstaking_keeper_evmstaking_proto_goTypes = []any{
	(*EVMEvent)(nil),       // 0: halo.evmstaking.keeper.EVMEvent
	(*types.EVMEvent)(nil), // 1: octane.evmengine.types.EVMEvent
}
var file_halo_evmstaking_keeper_evmstaking_proto_depIdxs = []int32{
	1, // 0: halo.evmstaking.keeper.EVMEvent.event:type_name -> octane.evmengine.types.EVMEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_halo_evmstaking_keeper_evmstaking_proto_init() }
func file_halo_evmstaking_keeper_evmstaking_proto_init() {
	if File_halo_evmstaking_keeper_evmstaking_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_halo_evmstaking_keeper_evmstaking_proto_rawDesc), len(file_halo_evmstaking_keeper_evmstaking_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_evmstaking_keeper_evmstaking_proto_goTypes,
		DependencyIndexes: file_halo_evmstaking_keeper_evmstaking_proto_depIdxs,
		MessageInfos:      file_halo_evmstaking_keeper_evmstaking_proto_msgTypes,
	}.Build()
	File_halo_evmstaking_keeper_evmstaking_proto = out.File
	file_halo_evmstaking_keeper_evmstaking_proto_goTypes = nil
	file_halo_evmstaking_keeper_evmstaking_proto_depIdxs = nil
}
