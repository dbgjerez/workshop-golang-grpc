// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: comment/grpc/comment.proto

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdComment  int64  `protobuf:"varint,1,opt,name=idComment,proto3" json:"idComment,omitempty"`
	IdObject   int32  `protobuf:"varint,2,opt,name=idObject,proto3" json:"idObject,omitempty"`
	TypeObject string `protobuf:"bytes,3,opt,name=typeObject,proto3" json:"typeObject,omitempty"`
	IdUser     int32  `protobuf:"varint,4,opt,name=idUser,proto3" json:"idUser,omitempty"`
	Comment    string `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_grpc_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_grpc_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_grpc_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetIdComment() int64 {
	if x != nil {
		return x.IdComment
	}
	return 0
}

func (x *Comment) GetIdObject() int32 {
	if x != nil {
		return x.IdObject
	}
	return 0
}

func (x *Comment) GetTypeObject() string {
	if x != nil {
		return x.TypeObject
	}
	return ""
}

func (x *Comment) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *Comment) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdObject   int32  `protobuf:"varint,1,opt,name=idObject,proto3" json:"idObject,omitempty"`
	TypeObject string `protobuf:"bytes,2,opt,name=typeObject,proto3" json:"typeObject,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_grpc_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_grpc_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_comment_grpc_comment_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveRequest) GetIdObject() int32 {
	if x != nil {
		return x.IdObject
	}
	return 0
}

func (x *RetrieveRequest) GetTypeObject() string {
	if x != nil {
		return x.TypeObject
	}
	return ""
}

type Comments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Comments) Reset() {
	*x = Comments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_grpc_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comments) ProtoMessage() {}

func (x *Comments) ProtoReflect() protoreflect.Message {
	mi := &file_comment_grpc_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comments.ProtoReflect.Descriptor instead.
func (*Comments) Descriptor() ([]byte, []int) {
	return file_comment_grpc_comment_proto_rawDescGZIP(), []int{2}
}

func (x *Comments) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

var File_comment_grpc_comment_proto protoreflect.FileDescriptor

var file_comment_grpc_comment_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x3d, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x62, 0x67, 0x6a, 0x65, 0x72, 0x65, 0x7a, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x68, 0x6f, 0x70, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_grpc_comment_proto_rawDescOnce sync.Once
	file_comment_grpc_comment_proto_rawDescData = file_comment_grpc_comment_proto_rawDesc
)

func file_comment_grpc_comment_proto_rawDescGZIP() []byte {
	file_comment_grpc_comment_proto_rawDescOnce.Do(func() {
		file_comment_grpc_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_grpc_comment_proto_rawDescData)
	})
	return file_comment_grpc_comment_proto_rawDescData
}

var file_comment_grpc_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_comment_grpc_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),         // 0: Comment
	(*RetrieveRequest)(nil), // 1: RetrieveRequest
	(*Comments)(nil),        // 2: Comments
}
var file_comment_grpc_comment_proto_depIdxs = []int32{
	0, // 0: Comments.comments:type_name -> Comment
	1, // 1: CommentService.Retrieve:input_type -> RetrieveRequest
	2, // 2: CommentService.Retrieve:output_type -> Comments
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comment_grpc_comment_proto_init() }
func file_comment_grpc_comment_proto_init() {
	if File_comment_grpc_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_grpc_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comment_grpc_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_comment_grpc_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comments); i {
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
			RawDescriptor: file_comment_grpc_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_grpc_comment_proto_goTypes,
		DependencyIndexes: file_comment_grpc_comment_proto_depIdxs,
		MessageInfos:      file_comment_grpc_comment_proto_msgTypes,
	}.Build()
	File_comment_grpc_comment_proto = out.File
	file_comment_grpc_comment_proto_rawDesc = nil
	file_comment_grpc_comment_proto_goTypes = nil
	file_comment_grpc_comment_proto_depIdxs = nil
}
