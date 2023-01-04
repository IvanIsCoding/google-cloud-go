// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/cloud/documentai/v1/processor_type.proto

package documentaipb

import (
	reflect "reflect"
	sync "sync"

	api "google.golang.org/genproto/googleapis/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A processor type is responsible for performing a certain document
// understanding task on a certain type of document.
type ProcessorType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the processor type.
	// Format: `projects/{project}/processorTypes/{processor_type}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The processor type, e.g., `OCR_PROCESSOR`, `INVOICE_PROCESSOR`, etc.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The processor category, used by UI to group processor types.
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// The locations in which this processor is available.
	AvailableLocations []*ProcessorType_LocationInfo `protobuf:"bytes,4,rep,name=available_locations,json=availableLocations,proto3" json:"available_locations,omitempty"`
	// Whether the processor type allows creation. If true, users can create a
	// processor of this processor type. Otherwise, users need to request access.
	AllowCreation bool `protobuf:"varint,6,opt,name=allow_creation,json=allowCreation,proto3" json:"allow_creation,omitempty"`
	// Launch stage of the processor type
	LaunchStage api.LaunchStage `protobuf:"varint,8,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
	// A set of Cloud Storage URIs of sample documents for this processor.
	SampleDocumentUris []string `protobuf:"bytes,9,rep,name=sample_document_uris,json=sampleDocumentUris,proto3" json:"sample_document_uris,omitempty"`
}

func (x *ProcessorType) Reset() {
	*x = ProcessorType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_processor_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessorType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorType) ProtoMessage() {}

func (x *ProcessorType) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_processor_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorType.ProtoReflect.Descriptor instead.
func (*ProcessorType) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_type_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessorType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessorType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProcessorType) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProcessorType) GetAvailableLocations() []*ProcessorType_LocationInfo {
	if x != nil {
		return x.AvailableLocations
	}
	return nil
}

func (x *ProcessorType) GetAllowCreation() bool {
	if x != nil {
		return x.AllowCreation
	}
	return false
}

func (x *ProcessorType) GetLaunchStage() api.LaunchStage {
	if x != nil {
		return x.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

func (x *ProcessorType) GetSampleDocumentUris() []string {
	if x != nil {
		return x.SampleDocumentUris
	}
	return nil
}

// The location information about where the processor is available.
type ProcessorType_LocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location id, currently must be one of [us, eu].
	LocationId string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *ProcessorType_LocationInfo) Reset() {
	*x = ProcessorType_LocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_processor_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessorType_LocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessorType_LocationInfo) ProtoMessage() {}

func (x *ProcessorType_LocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_processor_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessorType_LocationInfo.ProtoReflect.Descriptor instead.
func (*ProcessorType_LocationInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_processor_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProcessorType_LocationInfo) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

var File_google_cloud_documentai_v1_processor_type_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1_processor_type_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x67, 0x0a,
	0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a,
	0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x73, 0x1a, 0x2f, 0x0a, 0x0c, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3a, 0x75, 0xea, 0x41,
	0x72, 0x0a, 0x27, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x47, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x42, 0xdb, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x69, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x49, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1_processor_type_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1_processor_type_proto_rawDescData = file_google_cloud_documentai_v1_processor_type_proto_rawDesc
)

func file_google_cloud_documentai_v1_processor_type_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1_processor_type_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1_processor_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1_processor_type_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1_processor_type_proto_rawDescData
}

var file_google_cloud_documentai_v1_processor_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_documentai_v1_processor_type_proto_goTypes = []interface{}{
	(*ProcessorType)(nil),              // 0: google.cloud.documentai.v1.ProcessorType
	(*ProcessorType_LocationInfo)(nil), // 1: google.cloud.documentai.v1.ProcessorType.LocationInfo
	(api.LaunchStage)(0),               // 2: google.api.LaunchStage
}
var file_google_cloud_documentai_v1_processor_type_proto_depIdxs = []int32{
	1, // 0: google.cloud.documentai.v1.ProcessorType.available_locations:type_name -> google.cloud.documentai.v1.ProcessorType.LocationInfo
	2, // 1: google.cloud.documentai.v1.ProcessorType.launch_stage:type_name -> google.api.LaunchStage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1_processor_type_proto_init() }
func file_google_cloud_documentai_v1_processor_type_proto_init() {
	if File_google_cloud_documentai_v1_processor_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1_processor_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessorType); i {
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
		file_google_cloud_documentai_v1_processor_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessorType_LocationInfo); i {
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
			RawDescriptor: file_google_cloud_documentai_v1_processor_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1_processor_type_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1_processor_type_proto_depIdxs,
		MessageInfos:      file_google_cloud_documentai_v1_processor_type_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1_processor_type_proto = out.File
	file_google_cloud_documentai_v1_processor_type_proto_rawDesc = nil
	file_google_cloud_documentai_v1_processor_type_proto_goTypes = nil
	file_google_cloud_documentai_v1_processor_type_proto_depIdxs = nil
}
