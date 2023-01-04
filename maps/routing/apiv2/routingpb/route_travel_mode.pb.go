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
// source: google/maps/routing/v2/route_travel_mode.proto

package routingpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A set of values used to specify the mode of travel.
// NOTE: WALK, BICYCLE, and TWO_WHEELER routes are in beta and might sometimes
// be missing clear sidewalks, pedestrian paths, or bicycling paths.
// You must display this warning to the user for all walking, bicycling, and
// two-wheel routes that you display in your app.
type RouteTravelMode int32

const (
	// No travel mode specified. Defaults to `DRIVE`.
	RouteTravelMode_TRAVEL_MODE_UNSPECIFIED RouteTravelMode = 0
	// Travel by passenger car.
	RouteTravelMode_DRIVE RouteTravelMode = 1
	// Travel by bicycle.
	RouteTravelMode_BICYCLE RouteTravelMode = 2
	// Travel by walking.
	RouteTravelMode_WALK RouteTravelMode = 3
	// Two-wheeled, motorized vehicle. For example, motorcycle. Note that this
	// differs from the `BICYCLE` travel mode which covers human-powered mode.
	RouteTravelMode_TWO_WHEELER RouteTravelMode = 4
)

// Enum value maps for RouteTravelMode.
var (
	RouteTravelMode_name = map[int32]string{
		0: "TRAVEL_MODE_UNSPECIFIED",
		1: "DRIVE",
		2: "BICYCLE",
		3: "WALK",
		4: "TWO_WHEELER",
	}
	RouteTravelMode_value = map[string]int32{
		"TRAVEL_MODE_UNSPECIFIED": 0,
		"DRIVE":                   1,
		"BICYCLE":                 2,
		"WALK":                    3,
		"TWO_WHEELER":             4,
	}
)

func (x RouteTravelMode) Enum() *RouteTravelMode {
	p := new(RouteTravelMode)
	*p = x
	return p
}

func (x RouteTravelMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteTravelMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_route_travel_mode_proto_enumTypes[0].Descriptor()
}

func (RouteTravelMode) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_route_travel_mode_proto_enumTypes[0]
}

func (x RouteTravelMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteTravelMode.Descriptor instead.
func (RouteTravelMode) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_route_travel_mode_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routing_v2_route_travel_mode_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_route_travel_mode_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2a, 0x61, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54,
	0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x57,
	0x4f, 0x5f, 0x57, 0x48, 0x45, 0x45, 0x4c, 0x45, 0x52, 0x10, 0x04, 0x42, 0xcc, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x14, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d,
	0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02,
	0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_route_travel_mode_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_route_travel_mode_proto_rawDescData = file_google_maps_routing_v2_route_travel_mode_proto_rawDesc
)

func file_google_maps_routing_v2_route_travel_mode_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_route_travel_mode_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_route_travel_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_route_travel_mode_proto_rawDescData)
	})
	return file_google_maps_routing_v2_route_travel_mode_proto_rawDescData
}

var file_google_maps_routing_v2_route_travel_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_route_travel_mode_proto_goTypes = []interface{}{
	(RouteTravelMode)(0), // 0: google.maps.routing.v2.RouteTravelMode
}
var file_google_maps_routing_v2_route_travel_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_route_travel_mode_proto_init() }
func file_google_maps_routing_v2_route_travel_mode_proto_init() {
	if File_google_maps_routing_v2_route_travel_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_route_travel_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_route_travel_mode_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_route_travel_mode_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_route_travel_mode_proto_enumTypes,
	}.Build()
	File_google_maps_routing_v2_route_travel_mode_proto = out.File
	file_google_maps_routing_v2_route_travel_mode_proto_rawDesc = nil
	file_google_maps_routing_v2_route_travel_mode_proto_goTypes = nil
	file_google_maps_routing_v2_route_travel_mode_proto_depIdxs = nil
}
