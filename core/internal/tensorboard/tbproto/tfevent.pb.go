// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.4
// source: core/internal/tensorboard/tbproto/tfevent.proto

package tbproto

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

// An Event proto written by TensorBoard into a tfevents file.
//
// https://github.com/tensorflow/tensorboard/blob/master/tensorboard/compat/proto/event.proto
//
// We only include fields that are relevant to us.
type TFEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the event.
	WallTime float64 `protobuf:"fixed64,1,opt,name=wall_time,json=wallTime,proto3" json:"wall_time,omitempty"`
	// An event-specific "step" number, often used as the X axis in charts.
	//
	// The linked source documents this as "Global step of the event."
	//
	// The meaning of this seems to depend on the event itself. For example,
	// a summary event with a single value tagged "epoch_learning_rate"
	// may use the step as an epoch number, but a summary event for
	// "evaluation_loss_vs_iterations" may use as an iteration number.
	// Because of this, consecutive events in a tfevents file may have
	// unrelated step numbers.
	Step int64 `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	// Types that are assignable to What:
	//
	//	*TFEvent_Summary
	What isTFEvent_What `protobuf_oneof:"what"`
}

func (x *TFEvent) Reset() {
	*x = TFEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFEvent) ProtoMessage() {}

func (x *TFEvent) ProtoReflect() protoreflect.Message {
	mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFEvent.ProtoReflect.Descriptor instead.
func (*TFEvent) Descriptor() ([]byte, []int) {
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP(), []int{0}
}

func (x *TFEvent) GetWallTime() float64 {
	if x != nil {
		return x.WallTime
	}
	return 0
}

func (x *TFEvent) GetStep() int64 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (m *TFEvent) GetWhat() isTFEvent_What {
	if m != nil {
		return m.What
	}
	return nil
}

func (x *TFEvent) GetSummary() *Summary {
	if x, ok := x.GetWhat().(*TFEvent_Summary); ok {
		return x.Summary
	}
	return nil
}

type isTFEvent_What interface {
	isTFEvent_What()
}

type TFEvent_Summary struct {
	// A summary was generated.
	Summary *Summary `protobuf:"bytes,5,opt,name=summary,proto3,oneof"`
}

func (*TFEvent_Summary) isTFEvent_What() {}

// A TensorBoard "summary" event.
//
// https://github.com/tensorflow/tensorboard/blob/master/tensorboard/compat/proto/summary.proto
//
// We only include fields that are relevant to us.
type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of values for the summary.
	Value []*Summary_Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP(), []int{1}
}

func (x *Summary) GetValue() []*Summary_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

// Information about which plugins are able to make use of a certain
// summary value.
type SummaryMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data that associates a summary with a certain plugin.
	PluginData *SummaryMetadata_PluginData `protobuf:"bytes,1,opt,name=plugin_data,json=pluginData,proto3" json:"plugin_data,omitempty"`
}

func (x *SummaryMetadata) Reset() {
	*x = SummaryMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryMetadata) ProtoMessage() {}

func (x *SummaryMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryMetadata.ProtoReflect.Descriptor instead.
func (*SummaryMetadata) Descriptor() ([]byte, []int) {
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP(), []int{2}
}

func (x *SummaryMetadata) GetPluginData() *SummaryMetadata_PluginData {
	if x != nil {
		return x.PluginData
	}
	return nil
}

type Summary_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag name for the data.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Which plugins may use this data.
	//
	// See the note in the original source file: the metadata is only present
	// on the first summary value of each tag.
	Metadata *SummaryMetadata `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Value associated with the tag.
	//
	// Types that are assignable to Value:
	//
	//	*Summary_Value_SimpleValue
	//	*Summary_Value_Tensor
	Value isSummary_Value_Value `protobuf_oneof:"value"`
}

func (x *Summary_Value) Reset() {
	*x = Summary_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary_Value) ProtoMessage() {}

func (x *Summary_Value) ProtoReflect() protoreflect.Message {
	mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary_Value.ProtoReflect.Descriptor instead.
func (*Summary_Value) Descriptor() ([]byte, []int) {
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Summary_Value) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Summary_Value) GetMetadata() *SummaryMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (m *Summary_Value) GetValue() isSummary_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Summary_Value) GetSimpleValue() float32 {
	if x, ok := x.GetValue().(*Summary_Value_SimpleValue); ok {
		return x.SimpleValue
	}
	return 0
}

func (x *Summary_Value) GetTensor() *TensorProto {
	if x, ok := x.GetValue().(*Summary_Value_Tensor); ok {
		return x.Tensor
	}
	return nil
}

type isSummary_Value_Value interface {
	isSummary_Value_Value()
}

type Summary_Value_SimpleValue struct {
	SimpleValue float32 `protobuf:"fixed32,2,opt,name=simple_value,json=simpleValue,proto3,oneof"`
}

type Summary_Value_Tensor struct {
	Tensor *TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

func (*Summary_Value_SimpleValue) isSummary_Value_Value() {}

func (*Summary_Value_Tensor) isSummary_Value_Value() {}

type SummaryMetadata_PluginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the plugin this data pertains to.
	PluginName string `protobuf:"bytes,1,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
}

func (x *SummaryMetadata_PluginData) Reset() {
	*x = SummaryMetadata_PluginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryMetadata_PluginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryMetadata_PluginData) ProtoMessage() {}

func (x *SummaryMetadata_PluginData) ProtoReflect() protoreflect.Message {
	mi := &file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryMetadata_PluginData.ProtoReflect.Descriptor instead.
func (*SummaryMetadata_PluginData) Descriptor() ([]byte, []int) {
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SummaryMetadata_PluginData) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

var File_core_internal_tensorboard_tbproto_tfevent_proto protoreflect.FileDescriptor

var file_core_internal_tensorboard_tbproto_tfevent_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x74, 0x62, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x66, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x74, 0x62, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x68, 0x0a, 0x07, 0x54, 0x46, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x24, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x77, 0x68, 0x61, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x07,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x9d, 0x01,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7e, 0x0a,
	0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x2d,
	0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x64,
	0x62, 0x2f, 0x77, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x74, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescOnce sync.Once
	file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescData = file_core_internal_tensorboard_tbproto_tfevent_proto_rawDesc
)

func file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescGZIP() []byte {
	file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescOnce.Do(func() {
		file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescData)
	})
	return file_core_internal_tensorboard_tbproto_tfevent_proto_rawDescData
}

var file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_internal_tensorboard_tbproto_tfevent_proto_goTypes = []interface{}{
	(*TFEvent)(nil),                    // 0: TFEvent
	(*Summary)(nil),                    // 1: Summary
	(*SummaryMetadata)(nil),            // 2: SummaryMetadata
	(*Summary_Value)(nil),              // 3: Summary.Value
	(*SummaryMetadata_PluginData)(nil), // 4: SummaryMetadata.PluginData
	(*TensorProto)(nil),                // 5: TensorProto
}
var file_core_internal_tensorboard_tbproto_tfevent_proto_depIdxs = []int32{
	1, // 0: TFEvent.summary:type_name -> Summary
	3, // 1: Summary.value:type_name -> Summary.Value
	4, // 2: SummaryMetadata.plugin_data:type_name -> SummaryMetadata.PluginData
	2, // 3: Summary.Value.metadata:type_name -> SummaryMetadata
	5, // 4: Summary.Value.tensor:type_name -> TensorProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_core_internal_tensorboard_tbproto_tfevent_proto_init() }
func file_core_internal_tensorboard_tbproto_tfevent_proto_init() {
	if File_core_internal_tensorboard_tbproto_tfevent_proto != nil {
		return
	}
	file_core_internal_tensorboard_tbproto_tensor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFEvent); i {
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
		file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryMetadata); i {
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
		file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary_Value); i {
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
		file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryMetadata_PluginData); i {
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
	file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TFEvent_Summary)(nil),
	}
	file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Summary_Value_SimpleValue)(nil),
		(*Summary_Value_Tensor)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_internal_tensorboard_tbproto_tfevent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_internal_tensorboard_tbproto_tfevent_proto_goTypes,
		DependencyIndexes: file_core_internal_tensorboard_tbproto_tfevent_proto_depIdxs,
		MessageInfos:      file_core_internal_tensorboard_tbproto_tfevent_proto_msgTypes,
	}.Build()
	File_core_internal_tensorboard_tbproto_tfevent_proto = out.File
	file_core_internal_tensorboard_tbproto_tfevent_proto_rawDesc = nil
	file_core_internal_tensorboard_tbproto_tfevent_proto_goTypes = nil
	file_core_internal_tensorboard_tbproto_tfevent_proto_depIdxs = nil
}