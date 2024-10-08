// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: org/org.proto

package org

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrgEnum_Status int32

const (
	OrgEnum_Active   OrgEnum_Status = 0
	OrgEnum_Inactive OrgEnum_Status = 1
	OrgEnum_Deleted  OrgEnum_Status = 2
)

// Enum value maps for OrgEnum_Status.
var (
	OrgEnum_Status_name = map[int32]string{
		0: "Active",
		1: "Inactive",
		2: "Deleted",
	}
	OrgEnum_Status_value = map[string]int32{
		"Active":   0,
		"Inactive": 1,
		"Deleted":  2,
	}
)

func (x OrgEnum_Status) Enum() *OrgEnum_Status {
	p := new(OrgEnum_Status)
	*p = x
	return p
}

func (x OrgEnum_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrgEnum_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_org_org_proto_enumTypes[0].Descriptor()
}

func (OrgEnum_Status) Type() protoreflect.EnumType {
	return &file_org_org_proto_enumTypes[0]
}

func (x OrgEnum_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *OrgEnum_Status) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = OrgEnum_Status(num)
	return nil
}

// Deprecated: Use OrgEnum_Status.Descriptor instead.
func (OrgEnum_Status) EnumDescriptor() ([]byte, []int) {
	return file_org_org_proto_rawDescGZIP(), []int{0, 0}
}

type OrgEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrgEnum) Reset() {
	*x = OrgEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_org_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgEnum) ProtoMessage() {}

func (x *OrgEnum) ProtoReflect() protoreflect.Message {
	mi := &file_org_org_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgEnum.ProtoReflect.Descriptor instead.
func (*OrgEnum) Descriptor() ([]byte, []int) {
	return file_org_org_proto_rawDescGZIP(), []int{0}
}

type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string                `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name        *string                `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Description *string                `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	OrgId       *string                `protobuf:"bytes,4,req,name=orgId" json:"orgId,omitempty"`
	Status      *OrgEnum_Status        `protobuf:"varint,5,req,name=status,enum=OrgEnum_Status" json:"status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updatedAt" json:"updatedAt,omitempty"`
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deletedAt" json:"deletedAt,omitempty"`
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_org_org_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_org_org_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_org_org_proto_rawDescGZIP(), []int{1}
}

func (x *Org) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Org) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Org) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Org) GetOrgId() string {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return ""
}

func (x *Org) GetStatus() OrgEnum_Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return OrgEnum_Active
}

func (x *Org) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Org) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Org) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_org_org_proto protoreflect.FileDescriptor

var file_org_org_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3a, 0x0a, 0x07, 0x4f, 0x72, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x2f, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x02, 0x22, 0xb8, 0x02, 0x0a,
	0x03, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x4f, 0x72, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x65, 0x74, 0x6e, 0x61, 0x73,
	0x79, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6f, 0x72, 0x67,
}

var (
	file_org_org_proto_rawDescOnce sync.Once
	file_org_org_proto_rawDescData = file_org_org_proto_rawDesc
)

func file_org_org_proto_rawDescGZIP() []byte {
	file_org_org_proto_rawDescOnce.Do(func() {
		file_org_org_proto_rawDescData = protoimpl.X.CompressGZIP(file_org_org_proto_rawDescData)
	})
	return file_org_org_proto_rawDescData
}

var file_org_org_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_org_org_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_org_org_proto_goTypes = []any{
	(OrgEnum_Status)(0),           // 0: OrgEnum.Status
	(*OrgEnum)(nil),               // 1: OrgEnum
	(*Org)(nil),                   // 2: Org
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_org_org_proto_depIdxs = []int32{
	0, // 0: Org.status:type_name -> OrgEnum.Status
	3, // 1: Org.createdAt:type_name -> google.protobuf.Timestamp
	3, // 2: Org.updatedAt:type_name -> google.protobuf.Timestamp
	3, // 3: Org.deletedAt:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_org_org_proto_init() }
func file_org_org_proto_init() {
	if File_org_org_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_org_org_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OrgEnum); i {
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
		file_org_org_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Org); i {
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
			RawDescriptor: file_org_org_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_org_org_proto_goTypes,
		DependencyIndexes: file_org_org_proto_depIdxs,
		EnumInfos:         file_org_org_proto_enumTypes,
		MessageInfos:      file_org_org_proto_msgTypes,
	}.Build()
	File_org_org_proto = out.File
	file_org_org_proto_rawDesc = nil
	file_org_org_proto_goTypes = nil
	file_org_org_proto_depIdxs = nil
}
