// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: .proto

package grpc

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file___proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file___proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Buyer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName  string `protobuf:"bytes,3,opt,name=middleName,proto3" json:"middleName,omitempty"`
	Age         int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Registation string `protobuf:"bytes,5,opt,name=registation,proto3" json:"registation,omitempty"`
}

func (x *Buyer) Reset() {
	*x = Buyer{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buyer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buyer) ProtoMessage() {}

func (x *Buyer) ProtoReflect() protoreflect.Message {
	mi := &file___proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buyer.ProtoReflect.Descriptor instead.
func (*Buyer) Descriptor() ([]byte, []int) {
	return file___proto_rawDescGZIP(), []int{1}
}

func (x *Buyer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Buyer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Buyer) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Buyer) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Buyer) GetRegistation() string {
	if x != nil {
		return x.Registation
	}
	return ""
}

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Working bool   `protobuf:"varint,3,opt,name=working,proto3" json:"working,omitempty"`
	Owner   string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file___proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file___proto_rawDescGZIP(), []int{2}
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Shop) GetWorking() bool {
	if x != nil {
		return x.Working
	}
	return false
}

func (x *Shop) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

var File___proto protoreflect.FileDescriptor

var file___proto_rawDesc = []byte{
	0x0a, 0x06, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x79,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x64, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x32, 0xcd, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x65, 0x72,
	0x73, 0x41, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x42, 0x75, 0x79, 0x65, 0x72,
	0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x42, 0x75, 0x79, 0x65, 0x72, 0x1a,
	0x06, 0x2e, 0x42, 0x75, 0x79, 0x65, 0x72, 0x22, 0x00, 0x12, 0x19, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x1a, 0x05, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x42, 0x75, 0x79, 0x65, 0x72, 0x1a, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x1a, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file___proto_rawDescOnce sync.Once
	file___proto_rawDescData = file___proto_rawDesc
)

func file___proto_rawDescGZIP() []byte {
	file___proto_rawDescOnce.Do(func() {
		file___proto_rawDescData = protoimpl.X.CompressGZIP(file___proto_rawDescData)
	})
	return file___proto_rawDescData
}

var file___proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file___proto_goTypes = []interface{}{
	(*Status)(nil), // 0: Status
	(*Buyer)(nil),  // 1: Buyer
	(*Shop)(nil),   // 2: Shop
}
var file___proto_depIdxs = []int32{
	1, // 0: BuyersAndShops.CreateBuyer:input_type -> Buyer
	2, // 1: BuyersAndShops.CreateShop:input_type -> Shop
	1, // 2: BuyersAndShops.GetBuyer:input_type -> Buyer
	2, // 3: BuyersAndShops.GetShop:input_type -> Shop
	1, // 4: BuyersAndShops.DeleteBuyer:input_type -> Buyer
	2, // 5: BuyersAndShops.DeleteShop:input_type -> Shop
	0, // 6: BuyersAndShops.CreateBuyer:output_type -> Status
	0, // 7: BuyersAndShops.CreateShop:output_type -> Status
	1, // 8: BuyersAndShops.GetBuyer:output_type -> Buyer
	2, // 9: BuyersAndShops.GetShop:output_type -> Shop
	0, // 10: BuyersAndShops.DeleteBuyer:output_type -> Status
	0, // 11: BuyersAndShops.DeleteShop:output_type -> Status
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file___proto_init() }
func file___proto_init() {
	if File___proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file___proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Buyer); i {
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
		file___proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
			RawDescriptor: file___proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___proto_goTypes,
		DependencyIndexes: file___proto_depIdxs,
		MessageInfos:      file___proto_msgTypes,
	}.Build()
	File___proto = out.File
	file___proto_rawDesc = nil
	file___proto_goTypes = nil
	file___proto_depIdxs = nil
}
