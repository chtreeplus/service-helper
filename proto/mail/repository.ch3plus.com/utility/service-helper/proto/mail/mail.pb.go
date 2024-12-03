package mail

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	From     string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Receiver string   `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	To       []string `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	Cc       []string `protobuf:"bytes,5,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc      []string `protobuf:"bytes,6,rep,name=bcc,proto3" json:"bcc,omitempty"`
	Subject  string   `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	Link     string   `protobuf:"bytes,8,opt,name=link,proto3" json:"link,omitempty"`
	Body     string   `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	Attach   []string `protobuf:"bytes,10,rep,name=attach,proto3" json:"attach,omitempty"`
}

func (x *MailMessage) Reset() {
	*x = MailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailMessage) ProtoMessage() {}

func (x *MailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailMessage.ProtoReflect.Descriptor instead.
func (*MailMessage) Descriptor() ([]byte, []int) {
	return file_proto_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MailMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MailMessage) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *MailMessage) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MailMessage) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *MailMessage) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *MailMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MailMessage) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *MailMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MailMessage) GetAttach() []string {
	if x != nil {
		return x.Attach
	}
	return nil
}

var File_proto_mail_mail_proto protoreflect.FileDescriptor

var file_proto_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x64, 0x65, 0x63, 0x65, 0x6d, 0x74, 0x65,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x63, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x63, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x63, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x63, 0x65, 0x6d, 0x2d, 0x54, 0x65,
	0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61,
	0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mail_mail_proto_rawDescOnce sync.Once
	file_proto_mail_mail_proto_rawDescData = file_proto_mail_mail_proto_rawDesc
)

func file_proto_mail_mail_proto_rawDescGZIP() []byte {
	file_proto_mail_mail_proto_rawDescOnce.Do(func() {
		file_proto_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mail_mail_proto_rawDescData)
	})
	return file_proto_mail_mail_proto_rawDescData
}

var file_proto_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_mail_mail_proto_goTypes = []interface{}{
	(*MailMessage)(nil), // 0: decemtech.proto.helpers.mail.MailMessage
}
var file_proto_mail_mail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mail_mail_proto_init() }
func file_proto_mail_mail_proto_init() {
	if File_proto_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mail_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailMessage); i {
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
			RawDescriptor: file_proto_mail_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_mail_mail_proto_goTypes,
		DependencyIndexes: file_proto_mail_mail_proto_depIdxs,
		MessageInfos:      file_proto_mail_mail_proto_msgTypes,
	}.Build()
	File_proto_mail_mail_proto = out.File
	file_proto_mail_mail_proto_rawDesc = nil
	file_proto_mail_mail_proto_goTypes = nil
	file_proto_mail_mail_proto_depIdxs = nil
}
