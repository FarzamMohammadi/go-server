// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/apps/script/type/addon_widget_set.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The Widget type. DEFAULT is the basic widget set.
type AddOnWidgetSet_WidgetType int32

const (
	// The default widget set.
	AddOnWidgetSet_WIDGET_TYPE_UNSPECIFIED AddOnWidgetSet_WidgetType = 0
	// The date picker.
	AddOnWidgetSet_DATE_PICKER AddOnWidgetSet_WidgetType = 1
	// Styled buttons include filled buttons and disabled buttons.
	AddOnWidgetSet_STYLED_BUTTONS AddOnWidgetSet_WidgetType = 2
	// Persistent forms allow persisting form values during actions.
	AddOnWidgetSet_PERSISTENT_FORMS AddOnWidgetSet_WidgetType = 3
	// Fixed footer in card.
	AddOnWidgetSet_FIXED_FOOTER AddOnWidgetSet_WidgetType = 4
	// Update the subject and recipients of a draft.
	AddOnWidgetSet_UPDATE_SUBJECT_AND_RECIPIENTS AddOnWidgetSet_WidgetType = 5
	// The grid widget.
	AddOnWidgetSet_GRID_WIDGET AddOnWidgetSet_WidgetType = 6
	// A Gmail add-on action that applies to the addon compose UI.
	AddOnWidgetSet_ADDON_COMPOSE_UI_ACTION AddOnWidgetSet_WidgetType = 7
)

// Enum value maps for AddOnWidgetSet_WidgetType.
var (
	AddOnWidgetSet_WidgetType_name = map[int32]string{
		0: "WIDGET_TYPE_UNSPECIFIED",
		1: "DATE_PICKER",
		2: "STYLED_BUTTONS",
		3: "PERSISTENT_FORMS",
		4: "FIXED_FOOTER",
		5: "UPDATE_SUBJECT_AND_RECIPIENTS",
		6: "GRID_WIDGET",
		7: "ADDON_COMPOSE_UI_ACTION",
	}
	AddOnWidgetSet_WidgetType_value = map[string]int32{
		"WIDGET_TYPE_UNSPECIFIED":       0,
		"DATE_PICKER":                   1,
		"STYLED_BUTTONS":                2,
		"PERSISTENT_FORMS":              3,
		"FIXED_FOOTER":                  4,
		"UPDATE_SUBJECT_AND_RECIPIENTS": 5,
		"GRID_WIDGET":                   6,
		"ADDON_COMPOSE_UI_ACTION":       7,
	}
)

func (x AddOnWidgetSet_WidgetType) Enum() *AddOnWidgetSet_WidgetType {
	p := new(AddOnWidgetSet_WidgetType)
	*p = x
	return p
}

func (x AddOnWidgetSet_WidgetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddOnWidgetSet_WidgetType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_script_type_addon_widget_set_proto_enumTypes[0].Descriptor()
}

func (AddOnWidgetSet_WidgetType) Type() protoreflect.EnumType {
	return &file_google_apps_script_type_addon_widget_set_proto_enumTypes[0]
}

func (x AddOnWidgetSet_WidgetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddOnWidgetSet_WidgetType.Descriptor instead.
func (AddOnWidgetSet_WidgetType) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_script_type_addon_widget_set_proto_rawDescGZIP(), []int{0, 0}
}

// The widget subset used by an add-on.
type AddOnWidgetSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of widgets used in an add-on.
	UsedWidgets []AddOnWidgetSet_WidgetType `protobuf:"varint,1,rep,packed,name=used_widgets,json=usedWidgets,proto3,enum=google.apps.script.type.AddOnWidgetSet_WidgetType" json:"used_widgets,omitempty"`
}

func (x *AddOnWidgetSet) Reset() {
	*x = AddOnWidgetSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_addon_widget_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOnWidgetSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOnWidgetSet) ProtoMessage() {}

func (x *AddOnWidgetSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_addon_widget_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOnWidgetSet.ProtoReflect.Descriptor instead.
func (*AddOnWidgetSet) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_addon_widget_set_proto_rawDescGZIP(), []int{0}
}

func (x *AddOnWidgetSet) GetUsedWidgets() []AddOnWidgetSet_WidgetType {
	if x != nil {
		return x.UsedWidgets
	}
	return nil
}

var File_google_apps_script_type_addon_widget_set_proto protoreflect.FileDescriptor

var file_google_apps_script_type_addon_widget_set_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb1, 0x02, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x4f, 0x6e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x55, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x4f, 0x6e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x2e, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x0a, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x57, 0x49, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x54, 0x54, 0x4f,
	0x4e, 0x53, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x45, 0x52, 0x53, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x53, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49,
	0x58, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x4f, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x49, 0x50, 0x49, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x05, 0x12,
	0x0f, 0x0a, 0x0b, 0x47, 0x52, 0x49, 0x44, 0x5f, 0x57, 0x49, 0x44, 0x47, 0x45, 0x54, 0x10, 0x06,
	0x12, 0x1b, 0x0a, 0x17, 0x41, 0x44, 0x44, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53,
	0x45, 0x5f, 0x55, 0x49, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x42, 0xa8, 0x01,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x50, 0x01, 0x5a,
	0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0xea, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_script_type_addon_widget_set_proto_rawDescOnce sync.Once
	file_google_apps_script_type_addon_widget_set_proto_rawDescData = file_google_apps_script_type_addon_widget_set_proto_rawDesc
)

func file_google_apps_script_type_addon_widget_set_proto_rawDescGZIP() []byte {
	file_google_apps_script_type_addon_widget_set_proto_rawDescOnce.Do(func() {
		file_google_apps_script_type_addon_widget_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_script_type_addon_widget_set_proto_rawDescData)
	})
	return file_google_apps_script_type_addon_widget_set_proto_rawDescData
}

var file_google_apps_script_type_addon_widget_set_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_script_type_addon_widget_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_apps_script_type_addon_widget_set_proto_goTypes = []interface{}{
	(AddOnWidgetSet_WidgetType)(0), // 0: google.apps.script.type.AddOnWidgetSet.WidgetType
	(*AddOnWidgetSet)(nil),         // 1: google.apps.script.type.AddOnWidgetSet
}
var file_google_apps_script_type_addon_widget_set_proto_depIdxs = []int32{
	0, // 0: google.apps.script.type.AddOnWidgetSet.used_widgets:type_name -> google.apps.script.type.AddOnWidgetSet.WidgetType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_apps_script_type_addon_widget_set_proto_init() }
func file_google_apps_script_type_addon_widget_set_proto_init() {
	if File_google_apps_script_type_addon_widget_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_script_type_addon_widget_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOnWidgetSet); i {
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
			RawDescriptor: file_google_apps_script_type_addon_widget_set_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_script_type_addon_widget_set_proto_goTypes,
		DependencyIndexes: file_google_apps_script_type_addon_widget_set_proto_depIdxs,
		EnumInfos:         file_google_apps_script_type_addon_widget_set_proto_enumTypes,
		MessageInfos:      file_google_apps_script_type_addon_widget_set_proto_msgTypes,
	}.Build()
	File_google_apps_script_type_addon_widget_set_proto = out.File
	file_google_apps_script_type_addon_widget_set_proto_rawDesc = nil
	file_google_apps_script_type_addon_widget_set_proto_goTypes = nil
	file_google_apps_script_type_addon_widget_set_proto_depIdxs = nil
}
