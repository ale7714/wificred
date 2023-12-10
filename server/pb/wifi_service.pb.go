// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: wifi_service.proto

package pb

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

// The request message containing the WiFi credentials.
type WifiCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssid     string `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *WifiCredentials) Reset() {
	*x = WifiCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wifi_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiCredentials) ProtoMessage() {}

func (x *WifiCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_wifi_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiCredentials.ProtoReflect.Descriptor instead.
func (*WifiCredentials) Descriptor() ([]byte, []int) {
	return file_wifi_service_proto_rawDescGZIP(), []int{0}
}

func (x *WifiCredentials) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *WifiCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// The response message containing the confirmation.
type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wifi_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_wifi_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_wifi_service_proto_rawDescGZIP(), []int{1}
}

func (x *Confirmation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wifi_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_wifi_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_wifi_service_proto_rawDescGZIP(), []int{2}
}

var File_wifi_service_proto protoreflect.FileDescriptor

var file_wifi_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x77, 0x69, 0x66, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0f, 0x57, 0x69, 0x66, 0x69, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x73, 0x0a, 0x0b,
	0x57, 0x69, 0x66, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x10,
	0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x57, 0x69, 0x66, 0x69, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22,
	0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wifi_service_proto_rawDescOnce sync.Once
	file_wifi_service_proto_rawDescData = file_wifi_service_proto_rawDesc
)

func file_wifi_service_proto_rawDescGZIP() []byte {
	file_wifi_service_proto_rawDescOnce.Do(func() {
		file_wifi_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wifi_service_proto_rawDescData)
	})
	return file_wifi_service_proto_rawDescData
}

var file_wifi_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wifi_service_proto_goTypes = []interface{}{
	(*WifiCredentials)(nil), // 0: WifiCredentials
	(*Confirmation)(nil),    // 1: Confirmation
	(*Request)(nil),         // 2: Request
}
var file_wifi_service_proto_depIdxs = []int32{
	0, // 0: WifiService.SendCredentials:input_type -> WifiCredentials
	2, // 1: WifiService.GetCredentials:input_type -> Request
	1, // 2: WifiService.SendCredentials:output_type -> Confirmation
	0, // 3: WifiService.GetCredentials:output_type -> WifiCredentials
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wifi_service_proto_init() }
func file_wifi_service_proto_init() {
	if File_wifi_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wifi_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiCredentials); i {
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
		file_wifi_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_wifi_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_wifi_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wifi_service_proto_goTypes,
		DependencyIndexes: file_wifi_service_proto_depIdxs,
		MessageInfos:      file_wifi_service_proto_msgTypes,
	}.Build()
	File_wifi_service_proto = out.File
	file_wifi_service_proto_rawDesc = nil
	file_wifi_service_proto_goTypes = nil
	file_wifi_service_proto_depIdxs = nil
}
