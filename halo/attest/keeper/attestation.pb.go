// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: halo/attest/keeper/attestation.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

type Status int32

const (
	Status_Unknown  Status = 0
	Status_Pending  Status = 1
	Status_Approved Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Pending",
		2: "Approved",
	}
	Status_value = map[string]int32{
		"Unknown":  0,
		"Pending":  1,
		"Approved": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_halo_attest_keeper_attestation_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_halo_attest_keeper_attestation_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{0}
}

type Attestation struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                  // Auto-incremented ID
	ChainId         uint64                 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                         // Chain ID as per https://chainlist.org
	ConfLevel       uint32                 `protobuf:"varint,3,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`                   // Confirmation level of the cross-chain block
	AttestOffset    uint64                 `protobuf:"varint,4,opt,name=attest_offset,json=attestOffset,proto3" json:"attest_offset,omitempty"`          // Offset of the cross-chain block
	BlockHeight     uint64                 `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`             // Height of the source-chain block
	BlockHash       []byte                 `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                    // Hash of the source-chain block
	MsgRoot         []byte                 `protobuf:"bytes,7,opt,name=msg_root,json=msgRoot,proto3" json:"msg_root,omitempty"`                          // Merkle root of all the messages in the cross-chain Block
	AttestationRoot []byte                 `protobuf:"bytes,8,opt,name=attestation_root,json=attestationRoot,proto3" json:"attestation_root,omitempty"`  // Attestation merkle root of the cross-chain Block
	Status          uint32                 `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`                                          // Status of the block; pending, approved.
	ValidatorSetId  uint64                 `protobuf:"varint,10,opt,name=validator_set_id,json=validatorSetId,proto3" json:"validator_set_id,omitempty"` // Validator set that approved this attestation.
	CreatedHeight   uint64                 `protobuf:"varint,11,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"`      // Consensus height at which this attestation was created.
	FinalizedAttId  uint64                 `protobuf:"varint,12,opt,name=finalized_att_id,json=finalizedAttId,proto3" json:"finalized_att_id,omitempty"` // Approved finalized attestation for same chain_id and offset.
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attestation) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Attestation) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Attestation) GetAttestOffset() uint64 {
	if x != nil {
		return x.AttestOffset
	}
	return 0
}

func (x *Attestation) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Attestation) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Attestation) GetMsgRoot() []byte {
	if x != nil {
		return x.MsgRoot
	}
	return nil
}

func (x *Attestation) GetAttestationRoot() []byte {
	if x != nil {
		return x.AttestationRoot
	}
	return nil
}

func (x *Attestation) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Attestation) GetValidatorSetId() uint64 {
	if x != nil {
		return x.ValidatorSetId
	}
	return 0
}

func (x *Attestation) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

func (x *Attestation) GetFinalizedAttId() uint64 {
	if x != nil {
		return x.FinalizedAttId
	}
	return 0
}

// Signature is the attestation signature of the validator over the block root.
type Signature struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                    // Auto-incremented ID
	Signature        []byte                 `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`                                       // Validator signature over XBlockRoot; Ethereum 65 bytes [R || S || V] format.
	ValidatorAddress []byte                 `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"` // Validator ethereum address; 20 bytes.
	AttId            uint64                 `protobuf:"varint,4,opt,name=att_id,json=attId,proto3" json:"att_id,omitempty"`                                 // Attestation ID to which this signature belongs.
	ChainId          uint64                 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                           // Chain ID as per https://chainlist.org
	ConfLevel        uint32                 `protobuf:"varint,6,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`                     // Confirmation level of the cross-chain block
	AttestOffset     uint64                 `protobuf:"varint,7,opt,name=attest_offset,json=attestOffset,proto3" json:"attest_offset,omitempty"`            // Offset of the cross-chain block
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Signature) Reset() {
	*x = Signature{}
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_halo_attest_keeper_attestation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_halo_attest_keeper_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Signature) GetValidatorAddress() []byte {
	if x != nil {
		return x.ValidatorAddress
	}
	return nil
}

func (x *Signature) GetAttId() uint64 {
	if x != nil {
		return x.AttId
	}
	return 0
}

func (x *Signature) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Signature) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Signature) GetAttestOffset() uint64 {
	if x != nil {
		return x.AttestOffset
	}
	return 0
}

var File_halo_attest_keeper_attestation_proto protoreflect.FileDescriptor

const file_halo_attest_keeper_attestation_proto_rawDesc = "" +
	"\n" +
	"$halo/attest/keeper/attestation.proto\x12\x12halo.attest.keeper\x1a\x17cosmos/orm/v1/orm.proto\"\x83\x04\n" +
	"\vAttestation\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x19\n" +
	"\bchain_id\x18\x02 \x01(\x04R\achainId\x12\x1d\n" +
	"\n" +
	"conf_level\x18\x03 \x01(\rR\tconfLevel\x12#\n" +
	"\rattest_offset\x18\x04 \x01(\x04R\fattestOffset\x12!\n" +
	"\fblock_height\x18\x05 \x01(\x04R\vblockHeight\x12\x1d\n" +
	"\n" +
	"block_hash\x18\x06 \x01(\fR\tblockHash\x12\x19\n" +
	"\bmsg_root\x18\a \x01(\fR\amsgRoot\x12)\n" +
	"\x10attestation_root\x18\b \x01(\fR\x0fattestationRoot\x12\x16\n" +
	"\x06status\x18\t \x01(\rR\x06status\x12(\n" +
	"\x10validator_set_id\x18\n" +
	" \x01(\x04R\x0evalidatorSetId\x12%\n" +
	"\x0ecreated_height\x18\v \x01(\x04R\rcreatedHeight\x12(\n" +
	"\x10finalized_att_id\x18\f \x01(\x04R\x0efinalizedAttId:j\xf2\x9eӎ\x03d\n" +
	"\x06\n" +
	"\x02id\x10\x01\x12\x16\n" +
	"\x10attestation_root\x10\x01\x18\x01\x12,\n" +
	"(status,chain_id,conf_level,attest_offset\x10\x02\x12\x12\n" +
	"\x0ecreated_height\x10\x03\x18\x01\"\xc9\x02\n" +
	"\tSignature\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x1c\n" +
	"\tsignature\x18\x03 \x01(\fR\tsignature\x12+\n" +
	"\x11validator_address\x18\x02 \x01(\fR\x10validatorAddress\x12\x15\n" +
	"\x06att_id\x18\x04 \x01(\x04R\x05attId\x12\x19\n" +
	"\bchain_id\x18\x05 \x01(\x04R\achainId\x12\x1d\n" +
	"\n" +
	"conf_level\x18\x06 \x01(\rR\tconfLevel\x12#\n" +
	"\rattest_offset\x18\a \x01(\x04R\fattestOffset:k\xf2\x9eӎ\x03e\n" +
	"\x06\n" +
	"\x02id\x10\x01\x12\x1e\n" +
	"\x18att_id,validator_address\x10\x01\x18\x01\x129\n" +
	"3chain_id,conf_level,attest_offset,validator_address\x10\x02\x18\x01\x18\x02*0\n" +
	"\x06Status\x12\v\n" +
	"\aUnknown\x10\x00\x12\v\n" +
	"\aPending\x10\x01\x12\f\n" +
	"\bApproved\x10\x02B\xc5\x01\n" +
	"\x16com.halo.attest.keeperB\x10AttestationProtoP\x01Z/github.com/omni-network/omni/halo/attest/keeper\xa2\x02\x03HAK\xaa\x02\x12Halo.Attest.Keeper\xca\x02\x12Halo\\Attest\\Keeper\xe2\x02\x1eHalo\\Attest\\Keeper\\GPBMetadata\xea\x02\x14Halo::Attest::Keeperb\x06proto3"

var (
	file_halo_attest_keeper_attestation_proto_rawDescOnce sync.Once
	file_halo_attest_keeper_attestation_proto_rawDescData []byte
)

func file_halo_attest_keeper_attestation_proto_rawDescGZIP() []byte {
	file_halo_attest_keeper_attestation_proto_rawDescOnce.Do(func() {
		file_halo_attest_keeper_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_halo_attest_keeper_attestation_proto_rawDesc), len(file_halo_attest_keeper_attestation_proto_rawDesc)))
	})
	return file_halo_attest_keeper_attestation_proto_rawDescData
}

var file_halo_attest_keeper_attestation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_halo_attest_keeper_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_halo_attest_keeper_attestation_proto_goTypes = []any{
	(Status)(0),         // 0: halo.attest.keeper.Status
	(*Attestation)(nil), // 1: halo.attest.keeper.Attestation
	(*Signature)(nil),   // 2: halo.attest.keeper.Signature
}
var file_halo_attest_keeper_attestation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_halo_attest_keeper_attestation_proto_init() }
func file_halo_attest_keeper_attestation_proto_init() {
	if File_halo_attest_keeper_attestation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_halo_attest_keeper_attestation_proto_rawDesc), len(file_halo_attest_keeper_attestation_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_attest_keeper_attestation_proto_goTypes,
		DependencyIndexes: file_halo_attest_keeper_attestation_proto_depIdxs,
		EnumInfos:         file_halo_attest_keeper_attestation_proto_enumTypes,
		MessageInfos:      file_halo_attest_keeper_attestation_proto_msgTypes,
	}.Build()
	File_halo_attest_keeper_attestation_proto = out.File
	file_halo_attest_keeper_attestation_proto_goTypes = nil
	file_halo_attest_keeper_attestation_proto_depIdxs = nil
}
