// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: common.proto

package helpy

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

type Operation int32

const (
	// Default
	Operation__ Operation = 0
	// =
	Operation_eq Operation = 1
	// !=
	Operation_neq Operation = 2
	// >
	Operation_gt Operation = 3
	// >=
	Operation_gte Operation = 4
	// <
	Operation_lt Operation = 5
	// <=
	Operation_lte Operation = 6
	// ILIKE '...%' (strings only)
	Operation_begins Operation = 7
	// ILIKE '%...%' (strings only)
	Operation_contains Operation = 8
	// ILIKE '...' (strings only)
	Operation_equal Operation = 9
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "_",
		1: "eq",
		2: "neq",
		3: "gt",
		4: "gte",
		5: "lt",
		6: "lte",
		7: "begins",
		8: "contains",
		9: "equal",
	}
	Operation_value = map[string]int32{
		"_":        0,
		"eq":       1,
		"neq":      2,
		"gt":       3,
		"gte":      4,
		"lt":       5,
		"lte":      6,
		"begins":   7,
		"contains": 8,
		"equal":    9,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type OperationMulti int32

const (
	// IN
	OperationMulti_in OperationMulti = 0
	// NOT IN
	OperationMulti_nin OperationMulti = 1
)

// Enum value maps for OperationMulti.
var (
	OperationMulti_name = map[int32]string{
		0: "in",
		1: "nin",
	}
	OperationMulti_value = map[string]int32{
		"in":  0,
		"nin": 1,
	}
)

func (x OperationMulti) Enum() *OperationMulti {
	p := new(OperationMulti)
	*p = x
	return p
}

func (x OperationMulti) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationMulti) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (OperationMulti) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x OperationMulti) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationMulti.Descriptor instead.
func (OperationMulti) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type SortDirection int32

const (
	SortDirection_asc  SortDirection = 0
	SortDirection_desc SortDirection = 1
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "asc",
		1: "desc",
	}
	SortDirection_value = map[string]int32{
		"asc":  0,
		"desc": 1,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

// Filter by single value
type FilterItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string    `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Op    Operation `protobuf:"varint,2,opt,name=op,proto3,enum=helpy.Operation" json:"op,omitempty"`
}

func (x *FilterItem) Reset() {
	*x = FilterItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterItem) ProtoMessage() {}

func (x *FilterItem) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterItem.ProtoReflect.Descriptor instead.
func (*FilterItem) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *FilterItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *FilterItem) GetOp() Operation {
	if x != nil {
		return x.Op
	}
	return Operation__
}

// Filter by multiple values
type FilterItemMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string       `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Op     OperationMulti `protobuf:"varint,2,opt,name=op,proto3,enum=helpy.OperationMulti" json:"op,omitempty"`
}

func (x *FilterItemMulti) Reset() {
	*x = FilterItemMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterItemMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterItemMulti) ProtoMessage() {}

func (x *FilterItemMulti) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterItemMulti.ProtoReflect.Descriptor instead.
func (*FilterItemMulti) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *FilterItemMulti) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *FilterItemMulti) GetOp() OperationMulti {
	if x != nil {
		return x.Op
	}
	return OperationMulti_in
}

// Filter by range values
type FilterItemRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *FilterItemRange) Reset() {
	*x = FilterItemRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterItemRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterItemRange) ProtoMessage() {}

func (x *FilterItemRange) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterItemRange.ProtoReflect.Descriptor instead.
func (*FilterItemRange) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *FilterItemRange) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *FilterItemRange) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type FilterSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string        `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Dir   SortDirection `protobuf:"varint,2,opt,name=dir,proto3,enum=helpy.SortDirection" json:"dir,omitempty"`
}

func (x *FilterSort) Reset() {
	*x = FilterSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSort) ProtoMessage() {}

func (x *FilterSort) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterSort.ProtoReflect.Descriptor instead.
func (*FilterSort) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *FilterSort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FilterSort) GetDir() SortDirection {
	if x != nil {
		return x.Dir
	}
	return SortDirection_asc
}

type FilterPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    *int32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	PerPage *int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3,oneof" json:"per_page,omitempty"`
}

func (x *FilterPage) Reset() {
	*x = FilterPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterPage) ProtoMessage() {}

func (x *FilterPage) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterPage.ProtoReflect.Descriptor instead.
func (*FilterPage) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *FilterPage) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *FilterPage) GetPerPage() int32 {
	if x != nil && x.PerPage != nil {
		return *x.PerPage
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x68, 0x65, 0x6c, 0x70, 0x79, 0x22, 0x44, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x79, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x50, 0x0a, 0x0f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x79, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x35, 0x0a,
	0x0f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x79, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x64, 0x69, 0x72,
	0x22, 0x5b, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2a, 0x6a, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x65, 0x71, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x65, 0x71,
	0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x67, 0x74, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x67, 0x74,
	0x65, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x6c, 0x74, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x6c,
	0x74, 0x65, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x73, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x10, 0x08, 0x12, 0x09,
	0x0a, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x10, 0x09, 0x2a, 0x21, 0x0a, 0x0e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x06, 0x0a, 0x02, 0x69,
	0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x69, 0x6e, 0x10, 0x01, 0x2a, 0x22, 0x0a, 0x0d,
	0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a,
	0x03, 0x61, 0x73, 0x63, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x10, 0x01,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x2e, 0x69, 0x64, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x72, 0x75, 0x2f, 0x67, 0x6f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x6c,
	0x70, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_proto_goTypes = []interface{}{
	(Operation)(0),          // 0: helpy.Operation
	(OperationMulti)(0),     // 1: helpy.OperationMulti
	(SortDirection)(0),      // 2: helpy.SortDirection
	(*FilterItem)(nil),      // 3: helpy.FilterItem
	(*FilterItemMulti)(nil), // 4: helpy.FilterItemMulti
	(*FilterItemRange)(nil), // 5: helpy.FilterItemRange
	(*FilterSort)(nil),      // 6: helpy.FilterSort
	(*FilterPage)(nil),      // 7: helpy.FilterPage
}
var file_common_proto_depIdxs = []int32{
	0, // 0: helpy.FilterItem.op:type_name -> helpy.Operation
	1, // 1: helpy.FilterItemMulti.op:type_name -> helpy.OperationMulti
	2, // 2: helpy.FilterSort.dir:type_name -> helpy.SortDirection
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterItem); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterItemMulti); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterItemRange); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterSort); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterPage); i {
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
	file_common_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
