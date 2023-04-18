// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cells-mailer.proto

package mailer

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Language string `protobuf:"bytes,4,opt,name=Language,proto3" json:"Language,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User object used to compute the From header
	From *User `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	// List of target users to send the mail to
	To []*User `protobuf:"bytes,3,rep,name=To,proto3" json:"To,omitempty"`
	// List of target users to put in CC
	Cc []*User `protobuf:"bytes,4,rep,name=Cc,proto3" json:"Cc,omitempty"`
	// Date of sending
	DateSent int64 `protobuf:"varint,5,opt,name=DateSent,proto3" json:"DateSent,omitempty"`
	// String used as subject for the email
	Subject string `protobuf:"bytes,6,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// Plain-text content used for the body, if not set will be generated from the ContentHtml
	ContentPlain string `protobuf:"bytes,7,opt,name=ContentPlain,proto3" json:"ContentPlain,omitempty"`
	// HTML content used for the body
	ContentHtml string `protobuf:"bytes,8,opt,name=ContentHtml,proto3" json:"ContentHtml,omitempty"`
	// Markdown content used for the body
	ContentMarkdown string `protobuf:"bytes,9,opt,name=ContentMarkdown,proto3" json:"ContentMarkdown,omitempty"`
	// List of attachments
	Attachments []string `protobuf:"bytes,10,rep,name=Attachments,proto3" json:"Attachments,omitempty"`
	// Not used, could be used to create conversations
	ThreadUuid string `protobuf:"bytes,11,opt,name=ThreadUuid,proto3" json:"ThreadUuid,omitempty"`
	// Not used, could be used to create conversations
	ThreadIndex string `protobuf:"bytes,12,opt,name=ThreadIndex,proto3" json:"ThreadIndex,omitempty"`
	// Mail Template Id refers to predefined templates
	TemplateId string `protobuf:"bytes,13,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
	// Key/values to pass to the template
	TemplateData map[string]string `protobuf:"bytes,14,rep,name=TemplateData,proto3" json:"TemplateData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of retries after failed attempts (used internally)
	Retries int32 `protobuf:"varint,15,opt,name=Retries,proto3" json:"Retries,omitempty"`
	// Errors stacked on failed attempts
	SendErrors []string `protobuf:"bytes,16,rep,name=sendErrors,proto3" json:"sendErrors,omitempty"`
	// User object used to compute the Sender header
	Sender *User `protobuf:"bytes,17,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *Mail) GetFrom() *User {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Mail) GetTo() []*User {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Mail) GetCc() []*User {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *Mail) GetDateSent() int64 {
	if x != nil {
		return x.DateSent
	}
	return 0
}

func (x *Mail) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Mail) GetContentPlain() string {
	if x != nil {
		return x.ContentPlain
	}
	return ""
}

func (x *Mail) GetContentHtml() string {
	if x != nil {
		return x.ContentHtml
	}
	return ""
}

func (x *Mail) GetContentMarkdown() string {
	if x != nil {
		return x.ContentMarkdown
	}
	return ""
}

func (x *Mail) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Mail) GetThreadUuid() string {
	if x != nil {
		return x.ThreadUuid
	}
	return ""
}

func (x *Mail) GetThreadIndex() string {
	if x != nil {
		return x.ThreadIndex
	}
	return ""
}

func (x *Mail) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *Mail) GetTemplateData() map[string]string {
	if x != nil {
		return x.TemplateData
	}
	return nil
}

func (x *Mail) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *Mail) GetSendErrors() []string {
	if x != nil {
		return x.SendErrors
	}
	return nil
}

func (x *Mail) GetSender() *User {
	if x != nil {
		return x.Sender
	}
	return nil
}

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Complete mail object to send
	Mail *Mail `protobuf:"bytes,1,opt,name=Mail,proto3" json:"Mail,omitempty"`
	// Whether to place in mails queue or to send right away
	InQueue bool `protobuf:"varint,2,opt,name=InQueue,proto3" json:"InQueue,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{2}
}

func (x *SendMailRequest) GetMail() *Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

func (x *SendMailRequest) GetInQueue() bool {
	if x != nil {
		return x.InQueue
	}
	return false
}

type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{3}
}

func (x *SendMailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ConsumeQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxEmails int64 `protobuf:"varint,1,opt,name=MaxEmails,proto3" json:"MaxEmails,omitempty"`
}

func (x *ConsumeQueueRequest) Reset() {
	*x = ConsumeQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeQueueRequest) ProtoMessage() {}

func (x *ConsumeQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeQueueRequest.ProtoReflect.Descriptor instead.
func (*ConsumeQueueRequest) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumeQueueRequest) GetMaxEmails() int64 {
	if x != nil {
		return x.MaxEmails
	}
	return 0
}

type ConsumeQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	EmailsSent int64  `protobuf:"varint,2,opt,name=EmailsSent,proto3" json:"EmailsSent,omitempty"`
}

func (x *ConsumeQueueResponse) Reset() {
	*x = ConsumeQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cells_mailer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeQueueResponse) ProtoMessage() {}

func (x *ConsumeQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cells_mailer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeQueueResponse.ProtoReflect.Descriptor instead.
func (*ConsumeQueueResponse) Descriptor() ([]byte, []int) {
	return file_cells_mailer_proto_rawDescGZIP(), []int{5}
}

func (x *ConsumeQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConsumeQueueResponse) GetEmailsSent() int64 {
	if x != nil {
		return x.EmailsSent
	}
	return 0
}

var File_cells_mailer_proto protoreflect.FileDescriptor

var file_cells_mailer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0xf3, 0x04, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a,
	0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x1c, 0x0a, 0x02, 0x43,
	0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x02, 0x43, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x6c,
	0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x74,
	0x6d, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x55, 0x75, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x49, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x4d, 0x61, 0x78, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x4d, 0x61, 0x78, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x32, 0x9d, 0x01, 0x0a,
	0x0d, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x1b, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79, 0x64, 0x69, 0x6f,
	0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cells_mailer_proto_rawDescOnce sync.Once
	file_cells_mailer_proto_rawDescData = file_cells_mailer_proto_rawDesc
)

func file_cells_mailer_proto_rawDescGZIP() []byte {
	file_cells_mailer_proto_rawDescOnce.Do(func() {
		file_cells_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_cells_mailer_proto_rawDescData)
	})
	return file_cells_mailer_proto_rawDescData
}

var file_cells_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cells_mailer_proto_goTypes = []interface{}{
	(*User)(nil),                 // 0: mailer.User
	(*Mail)(nil),                 // 1: mailer.Mail
	(*SendMailRequest)(nil),      // 2: mailer.SendMailRequest
	(*SendMailResponse)(nil),     // 3: mailer.SendMailResponse
	(*ConsumeQueueRequest)(nil),  // 4: mailer.ConsumeQueueRequest
	(*ConsumeQueueResponse)(nil), // 5: mailer.ConsumeQueueResponse
	nil,                          // 6: mailer.Mail.TemplateDataEntry
}
var file_cells_mailer_proto_depIdxs = []int32{
	0, // 0: mailer.Mail.From:type_name -> mailer.User
	0, // 1: mailer.Mail.To:type_name -> mailer.User
	0, // 2: mailer.Mail.Cc:type_name -> mailer.User
	6, // 3: mailer.Mail.TemplateData:type_name -> mailer.Mail.TemplateDataEntry
	0, // 4: mailer.Mail.Sender:type_name -> mailer.User
	1, // 5: mailer.SendMailRequest.Mail:type_name -> mailer.Mail
	2, // 6: mailer.MailerService.SendMail:input_type -> mailer.SendMailRequest
	4, // 7: mailer.MailerService.ConsumeQueue:input_type -> mailer.ConsumeQueueRequest
	3, // 8: mailer.MailerService.SendMail:output_type -> mailer.SendMailResponse
	5, // 9: mailer.MailerService.ConsumeQueue:output_type -> mailer.ConsumeQueueResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cells_mailer_proto_init() }
func file_cells_mailer_proto_init() {
	if File_cells_mailer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cells_mailer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cells_mailer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
		file_cells_mailer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_cells_mailer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResponse); i {
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
		file_cells_mailer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeQueueRequest); i {
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
		file_cells_mailer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeQueueResponse); i {
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
			RawDescriptor: file_cells_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cells_mailer_proto_goTypes,
		DependencyIndexes: file_cells_mailer_proto_depIdxs,
		MessageInfos:      file_cells_mailer_proto_msgTypes,
	}.Build()
	File_cells_mailer_proto = out.File
	file_cells_mailer_proto_rawDesc = nil
	file_cells_mailer_proto_goTypes = nil
	file_cells_mailer_proto_depIdxs = nil
}
