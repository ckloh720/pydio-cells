// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "github.com/pydio/cells/v4/common/proto/auth"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SettingsMenuRequest struct {
}

func (m *SettingsMenuRequest) Reset()                    { *m = SettingsMenuRequest{} }
func (m *SettingsMenuRequest) String() string            { return proto.CompactTextString(m) }
func (*SettingsMenuRequest) ProtoMessage()               {}
func (*SettingsMenuRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type SettingsEntryMeta struct {
	IconClass string   `protobuf:"bytes,1,opt,name=IconClass,json=icon_class" json:"IconClass,omitempty"`
	Component string   `protobuf:"bytes,2,opt,name=Component,json=component" json:"Component,omitempty"`
	Props     string   `protobuf:"bytes,3,opt,name=Props,json=props" json:"Props,omitempty"`
	Advanced  bool     `protobuf:"varint,4,opt,name=Advanced,json=advanced" json:"Advanced,omitempty"`
	Indexed   []string `protobuf:"bytes,5,rep,name=Indexed,json=indexed" json:"Indexed,omitempty"`
}

func (m *SettingsEntryMeta) Reset()                    { *m = SettingsEntryMeta{} }
func (m *SettingsEntryMeta) String() string            { return proto.CompactTextString(m) }
func (*SettingsEntryMeta) ProtoMessage()               {}
func (*SettingsEntryMeta) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *SettingsEntryMeta) GetIconClass() string {
	if m != nil {
		return m.IconClass
	}
	return ""
}

func (m *SettingsEntryMeta) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SettingsEntryMeta) GetProps() string {
	if m != nil {
		return m.Props
	}
	return ""
}

func (m *SettingsEntryMeta) GetAdvanced() bool {
	if m != nil {
		return m.Advanced
	}
	return false
}

func (m *SettingsEntryMeta) GetIndexed() []string {
	if m != nil {
		return m.Indexed
	}
	return nil
}

type SettingsEntry struct {
	Key         string                     `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Label       string                     `protobuf:"bytes,2,opt,name=Label,json=LABEL" json:"Label,omitempty"`
	Description string                     `protobuf:"bytes,3,opt,name=Description,json=DESCRIPTION" json:"Description,omitempty"`
	Manager     string                     `protobuf:"bytes,4,opt,name=Manager,json=MANAGER" json:"Manager,omitempty"`
	Alias       string                     `protobuf:"bytes,5,opt,name=Alias,json=ALIAS" json:"Alias,omitempty"`
	Metadata    *SettingsEntryMeta         `protobuf:"bytes,6,opt,name=Metadata,json=METADATA" json:"Metadata,omitempty"`
	Accesses    map[string]*SettingsAccess `protobuf:"bytes,7,rep,name=Accesses" json:"Accesses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feature     string                     `protobuf:"bytes,8,opt,name=Feature" json:"Feature,omitempty"`
}

func (m *SettingsEntry) Reset()                    { *m = SettingsEntry{} }
func (m *SettingsEntry) String() string            { return proto.CompactTextString(m) }
func (*SettingsEntry) ProtoMessage()               {}
func (*SettingsEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *SettingsEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SettingsEntry) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SettingsEntry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SettingsEntry) GetManager() string {
	if m != nil {
		return m.Manager
	}
	return ""
}

func (m *SettingsEntry) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SettingsEntry) GetMetadata() *SettingsEntryMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SettingsEntry) GetAccesses() map[string]*SettingsAccess {
	if m != nil {
		return m.Accesses
	}
	return nil
}

func (m *SettingsEntry) GetFeature() string {
	if m != nil {
		return m.Feature
	}
	return ""
}

type SettingsAccess struct {
	Label       string                      `protobuf:"bytes,1,opt,name=Label" json:"Label,omitempty"`
	Description string                      `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	Policies    []*SettingsAccessRestPolicy `protobuf:"bytes,2,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *SettingsAccess) Reset()                    { *m = SettingsAccess{} }
func (m *SettingsAccess) String() string            { return proto.CompactTextString(m) }
func (*SettingsAccess) ProtoMessage()               {}
func (*SettingsAccess) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *SettingsAccess) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SettingsAccess) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SettingsAccess) GetPolicies() []*SettingsAccessRestPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type SettingsAccessRestPolicy struct {
	Action   string `protobuf:"bytes,2,opt,name=Action" json:"Action,omitempty"`
	Resource string `protobuf:"bytes,3,opt,name=Resource" json:"Resource,omitempty"`
}

func (m *SettingsAccessRestPolicy) Reset()                    { *m = SettingsAccessRestPolicy{} }
func (m *SettingsAccessRestPolicy) String() string            { return proto.CompactTextString(m) }
func (*SettingsAccessRestPolicy) ProtoMessage()               {}
func (*SettingsAccessRestPolicy) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SettingsAccessRestPolicy) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SettingsAccessRestPolicy) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type SettingsSection struct {
	Key         string           `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Label       string           `protobuf:"bytes,2,opt,name=Label,json=LABEL" json:"Label,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=Description,json=DESCRIPTION" json:"Description,omitempty"`
	Children    []*SettingsEntry `protobuf:"bytes,4,rep,name=Children,json=CHILDREN" json:"Children,omitempty"`
}

func (m *SettingsSection) Reset()                    { *m = SettingsSection{} }
func (m *SettingsSection) String() string            { return proto.CompactTextString(m) }
func (*SettingsSection) ProtoMessage()               {}
func (*SettingsSection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *SettingsSection) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SettingsSection) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SettingsSection) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SettingsSection) GetChildren() []*SettingsEntry {
	if m != nil {
		return m.Children
	}
	return nil
}

type SettingsMenuResponse struct {
	RootMetadata *SettingsEntryMeta `protobuf:"bytes,1,opt,name=RootMetadata,json=__metadata__" json:"RootMetadata,omitempty"`
	Sections     []*SettingsSection `protobuf:"bytes,2,rep,name=Sections" json:"Sections,omitempty"`
}

func (m *SettingsMenuResponse) Reset()                    { *m = SettingsMenuResponse{} }
func (m *SettingsMenuResponse) String() string            { return proto.CompactTextString(m) }
func (*SettingsMenuResponse) ProtoMessage()               {}
func (*SettingsMenuResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *SettingsMenuResponse) GetRootMetadata() *SettingsEntryMeta {
	if m != nil {
		return m.RootMetadata
	}
	return nil
}

func (m *SettingsMenuResponse) GetSections() []*SettingsSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

type FrontStateRequest struct {
}

func (m *FrontStateRequest) Reset()                    { *m = FrontStateRequest{} }
func (m *FrontStateRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontStateRequest) ProtoMessage()               {}
func (*FrontStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type FrontStateResponse struct {
}

func (m *FrontStateResponse) Reset()                    { *m = FrontStateResponse{} }
func (m *FrontStateResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontStateResponse) ProtoMessage()               {}
func (*FrontStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

type FrontPluginsRequest struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
}

func (m *FrontPluginsRequest) Reset()                    { *m = FrontPluginsRequest{} }
func (m *FrontPluginsRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontPluginsRequest) ProtoMessage()               {}
func (*FrontPluginsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *FrontPluginsRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type FrontPluginsResponse struct {
}

func (m *FrontPluginsResponse) Reset()                    { *m = FrontPluginsResponse{} }
func (m *FrontPluginsResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontPluginsResponse) ProtoMessage()               {}
func (*FrontPluginsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

type FrontMessagesRequest struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
}

func (m *FrontMessagesRequest) Reset()                    { *m = FrontMessagesRequest{} }
func (m *FrontMessagesRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontMessagesRequest) ProtoMessage()               {}
func (*FrontMessagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *FrontMessagesRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type FrontMessagesResponse struct {
	Messages map[string]string `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontMessagesResponse) Reset()                    { *m = FrontMessagesResponse{} }
func (m *FrontMessagesResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontMessagesResponse) ProtoMessage()               {}
func (*FrontMessagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *FrontMessagesResponse) GetMessages() map[string]string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type FrontSessionGetRequest struct {
}

func (m *FrontSessionGetRequest) Reset()                    { *m = FrontSessionGetRequest{} }
func (m *FrontSessionGetRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionGetRequest) ProtoMessage()               {}
func (*FrontSessionGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

type FrontSessionGetResponse struct {
	Token *auth.Token `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
}

func (m *FrontSessionGetResponse) Reset()                    { *m = FrontSessionGetResponse{} }
func (m *FrontSessionGetResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionGetResponse) ProtoMessage()               {}
func (*FrontSessionGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *FrontSessionGetResponse) GetToken() *auth.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type FrontSessionRequest struct {
	// Time reference for computing jwt expiry
	ClientTime int32 `protobuf:"varint,1,opt,name=ClientTime" json:"ClientTime,omitempty"`
	// Data sent back by specific auth steps
	AuthInfo map[string]string `protobuf:"bytes,2,rep,name=AuthInfo" json:"AuthInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Kill session now
	Logout bool `protobuf:"varint,3,opt,name=Logout" json:"Logout,omitempty"`
}

func (m *FrontSessionRequest) Reset()                    { *m = FrontSessionRequest{} }
func (m *FrontSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionRequest) ProtoMessage()               {}
func (*FrontSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *FrontSessionRequest) GetClientTime() int32 {
	if m != nil {
		return m.ClientTime
	}
	return 0
}

func (m *FrontSessionRequest) GetAuthInfo() map[string]string {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *FrontSessionRequest) GetLogout() bool {
	if m != nil {
		return m.Logout
	}
	return false
}

type FrontSessionResponse struct {
	// Legacy information (now in token)
	JWT        string `protobuf:"bytes,1,opt,name=JWT" json:"JWT,omitempty"`
	ExpireTime int32  `protobuf:"varint,2,opt,name=ExpireTime" json:"ExpireTime,omitempty"`
	// Trigger a specific Auth step
	Trigger string `protobuf:"bytes,3,opt,name=Trigger" json:"Trigger,omitempty"`
	// Additional data for the trigger
	TriggerInfo map[string]string `protobuf:"bytes,4,rep,name=TriggerInfo" json:"TriggerInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Token       *auth.Token       `protobuf:"bytes,5,opt,name=Token" json:"Token,omitempty"`
	RedirectTo  string            `protobuf:"bytes,6,opt,name=RedirectTo" json:"RedirectTo,omitempty"`
	Error       string            `protobuf:"bytes,7,opt,name=Error" json:"Error,omitempty"`
}

func (m *FrontSessionResponse) Reset()                    { *m = FrontSessionResponse{} }
func (m *FrontSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionResponse) ProtoMessage()               {}
func (*FrontSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *FrontSessionResponse) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

func (m *FrontSessionResponse) GetExpireTime() int32 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *FrontSessionResponse) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *FrontSessionResponse) GetTriggerInfo() map[string]string {
	if m != nil {
		return m.TriggerInfo
	}
	return nil
}

func (m *FrontSessionResponse) GetToken() *auth.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *FrontSessionResponse) GetRedirectTo() string {
	if m != nil {
		return m.RedirectTo
	}
	return ""
}

func (m *FrontSessionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FrontSessionDelRequest struct {
}

func (m *FrontSessionDelRequest) Reset()                    { *m = FrontSessionDelRequest{} }
func (m *FrontSessionDelRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionDelRequest) ProtoMessage()               {}
func (*FrontSessionDelRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

type FrontSessionDelResponse struct {
}

func (m *FrontSessionDelResponse) Reset()                    { *m = FrontSessionDelResponse{} }
func (m *FrontSessionDelResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionDelResponse) ProtoMessage()               {}
func (*FrontSessionDelResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

type FrontAuthRequest struct {
	RequestID string `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
}

func (m *FrontAuthRequest) Reset()                    { *m = FrontAuthRequest{} }
func (m *FrontAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontAuthRequest) ProtoMessage()               {}
func (*FrontAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func (m *FrontAuthRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

type FrontAuthResponse struct {
}

func (m *FrontAuthResponse) Reset()                    { *m = FrontAuthResponse{} }
func (m *FrontAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontAuthResponse) ProtoMessage()               {}
func (*FrontAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{20} }

type FrontEnrollAuthRequest struct {
	EnrollType string            `protobuf:"bytes,1,opt,name=EnrollType" json:"EnrollType,omitempty"`
	EnrollInfo map[string]string `protobuf:"bytes,2,rep,name=EnrollInfo" json:"EnrollInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontEnrollAuthRequest) Reset()                    { *m = FrontEnrollAuthRequest{} }
func (m *FrontEnrollAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontEnrollAuthRequest) ProtoMessage()               {}
func (*FrontEnrollAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{21} }

func (m *FrontEnrollAuthRequest) GetEnrollType() string {
	if m != nil {
		return m.EnrollType
	}
	return ""
}

func (m *FrontEnrollAuthRequest) GetEnrollInfo() map[string]string {
	if m != nil {
		return m.EnrollInfo
	}
	return nil
}

type FrontEnrollAuthResponse struct {
	// Any parameters can be returned
	Info map[string]string `protobuf:"bytes,1,rep,name=Info" json:"Info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontEnrollAuthResponse) Reset()                    { *m = FrontEnrollAuthResponse{} }
func (m *FrontEnrollAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontEnrollAuthResponse) ProtoMessage()               {}
func (*FrontEnrollAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{22} }

func (m *FrontEnrollAuthResponse) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

// Download binary
type FrontBinaryRequest struct {
	// Currently supported values are USER and GLOBAL
	BinaryType string `protobuf:"bytes,1,opt,name=BinaryType" json:"BinaryType,omitempty"`
	// Id of the binary
	Uuid string `protobuf:"bytes,2,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *FrontBinaryRequest) Reset()                    { *m = FrontBinaryRequest{} }
func (m *FrontBinaryRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontBinaryRequest) ProtoMessage()               {}
func (*FrontBinaryRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{23} }

func (m *FrontBinaryRequest) GetBinaryType() string {
	if m != nil {
		return m.BinaryType
	}
	return ""
}

func (m *FrontBinaryRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Not used, endpoint returns octet-stream
type FrontBinaryResponse struct {
}

func (m *FrontBinaryResponse) Reset()                    { *m = FrontBinaryResponse{} }
func (m *FrontBinaryResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontBinaryResponse) ProtoMessage()               {}
func (*FrontBinaryResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{24} }

type FrontBootConfRequest struct {
}

func (m *FrontBootConfRequest) Reset()                    { *m = FrontBootConfRequest{} }
func (m *FrontBootConfRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontBootConfRequest) ProtoMessage()               {}
func (*FrontBootConfRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{25} }

type FrontBootConfResponse struct {
	JsonData map[string]string `protobuf:"bytes,1,rep,name=JsonData" json:"JsonData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontBootConfResponse) Reset()                    { *m = FrontBootConfResponse{} }
func (m *FrontBootConfResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontBootConfResponse) ProtoMessage()               {}
func (*FrontBootConfResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{26} }

func (m *FrontBootConfResponse) GetJsonData() map[string]string {
	if m != nil {
		return m.JsonData
	}
	return nil
}

func init() {
	proto.RegisterType((*SettingsMenuRequest)(nil), "rest.SettingsMenuRequest")
	proto.RegisterType((*SettingsEntryMeta)(nil), "rest.SettingsEntryMeta")
	proto.RegisterType((*SettingsEntry)(nil), "rest.SettingsEntry")
	proto.RegisterType((*SettingsAccess)(nil), "rest.SettingsAccess")
	proto.RegisterType((*SettingsAccessRestPolicy)(nil), "rest.SettingsAccessRestPolicy")
	proto.RegisterType((*SettingsSection)(nil), "rest.SettingsSection")
	proto.RegisterType((*SettingsMenuResponse)(nil), "rest.SettingsMenuResponse")
	proto.RegisterType((*FrontStateRequest)(nil), "rest.FrontStateRequest")
	proto.RegisterType((*FrontStateResponse)(nil), "rest.FrontStateResponse")
	proto.RegisterType((*FrontPluginsRequest)(nil), "rest.FrontPluginsRequest")
	proto.RegisterType((*FrontPluginsResponse)(nil), "rest.FrontPluginsResponse")
	proto.RegisterType((*FrontMessagesRequest)(nil), "rest.FrontMessagesRequest")
	proto.RegisterType((*FrontMessagesResponse)(nil), "rest.FrontMessagesResponse")
	proto.RegisterType((*FrontSessionGetRequest)(nil), "rest.FrontSessionGetRequest")
	proto.RegisterType((*FrontSessionGetResponse)(nil), "rest.FrontSessionGetResponse")
	proto.RegisterType((*FrontSessionRequest)(nil), "rest.FrontSessionRequest")
	proto.RegisterType((*FrontSessionResponse)(nil), "rest.FrontSessionResponse")
	proto.RegisterType((*FrontSessionDelRequest)(nil), "rest.FrontSessionDelRequest")
	proto.RegisterType((*FrontSessionDelResponse)(nil), "rest.FrontSessionDelResponse")
	proto.RegisterType((*FrontAuthRequest)(nil), "rest.FrontAuthRequest")
	proto.RegisterType((*FrontAuthResponse)(nil), "rest.FrontAuthResponse")
	proto.RegisterType((*FrontEnrollAuthRequest)(nil), "rest.FrontEnrollAuthRequest")
	proto.RegisterType((*FrontEnrollAuthResponse)(nil), "rest.FrontEnrollAuthResponse")
	proto.RegisterType((*FrontBinaryRequest)(nil), "rest.FrontBinaryRequest")
	proto.RegisterType((*FrontBinaryResponse)(nil), "rest.FrontBinaryResponse")
	proto.RegisterType((*FrontBootConfRequest)(nil), "rest.FrontBootConfRequest")
	proto.RegisterType((*FrontBootConfResponse)(nil), "rest.FrontBootConfResponse")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1087 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x65, 0x2b, 0xa6, 0x46, 0x71, 0xe2, 0xd0, 0xb2, 0xc3, 0x0a, 0x69, 0xa0, 0xf0, 0x52,
	0x27, 0x2d, 0xa8, 0xd6, 0x3e, 0xb4, 0x48, 0x9a, 0x02, 0xb4, 0xa4, 0x24, 0x72, 0x25, 0xd7, 0x5d,
	0xab, 0xc8, 0xd1, 0xa0, 0xa9, 0xb5, 0x4c, 0x98, 0xda, 0x55, 0xb9, 0xcb, 0x20, 0xba, 0x17, 0x3d,
	0x15, 0xe8, 0x03, 0xb4, 0xaf, 0xd3, 0x4b, 0x2f, 0x7d, 0x8f, 0x3e, 0x45, 0xb1, 0x7f, 0x22, 0x69,
	0xc9, 0x01, 0xdc, 0xf6, 0x22, 0xec, 0xcc, 0x7c, 0x33, 0x3b, 0xfb, 0xcd, 0x8f, 0x08, 0xf7, 0x2e,
	0x52, 0x4a, 0x38, 0x26, 0x63, 0x7f, 0x96, 0x52, 0x4e, 0x9d, 0xf5, 0x14, 0x33, 0xde, 0x3c, 0x98,
	0xc4, 0xfc, 0x32, 0x3b, 0xf7, 0x23, 0x3a, 0x6d, 0xcf, 0xe6, 0xe3, 0x98, 0xb6, 0x23, 0x9c, 0x24,
	0xac, 0x1d, 0xd1, 0xe9, 0x94, 0x92, 0xb6, 0x84, 0xb6, 0xc3, 0x8c, 0x5f, 0xca, 0x1f, 0xe5, 0xea,
	0xed, 0xc0, 0xf6, 0x29, 0xe6, 0x3c, 0x26, 0x13, 0x36, 0xc4, 0x24, 0x43, 0xf8, 0xc7, 0x0c, 0x33,
	0xee, 0xfd, 0x6e, 0xc1, 0x03, 0xa3, 0xef, 0x11, 0x9e, 0xce, 0x87, 0x98, 0x87, 0xce, 0xc7, 0x50,
	0xeb, 0x47, 0x94, 0x74, 0x92, 0x90, 0x31, 0xd7, 0x6a, 0x59, 0x7b, 0x35, 0x04, 0x71, 0x44, 0xc9,
	0x59, 0x24, 0x34, 0xce, 0x23, 0xa8, 0x75, 0xe8, 0x74, 0x46, 0x09, 0x26, 0xdc, 0xad, 0x48, 0x73,
	0x2d, 0x32, 0x0a, 0xa7, 0x01, 0xd5, 0x93, 0x94, 0xce, 0x98, 0xbb, 0x26, 0x2d, 0xd5, 0x99, 0x10,
	0x9c, 0x26, 0xd8, 0xc1, 0xf8, 0x5d, 0x48, 0x22, 0x3c, 0x76, 0xd7, 0x5b, 0xd6, 0x9e, 0x8d, 0xec,
	0x50, 0xcb, 0x8e, 0x0b, 0x1b, 0x7d, 0x32, 0xc6, 0xef, 0xf1, 0xd8, 0xad, 0xb6, 0xd6, 0xf6, 0x6a,
	0x68, 0x23, 0x56, 0xa2, 0xf7, 0x77, 0x05, 0x36, 0x4b, 0xe9, 0x39, 0x5b, 0xb0, 0xf6, 0x2d, 0x9e,
	0xeb, 0xa4, 0xc4, 0x51, 0xdc, 0x37, 0x08, 0xcf, 0x71, 0xa2, 0x33, 0xa9, 0x0e, 0x82, 0xc3, 0xde,
	0xc0, 0x69, 0x41, 0xbd, 0x8b, 0x59, 0x94, 0xc6, 0x33, 0x1e, 0x53, 0xa2, 0x73, 0xa9, 0x77, 0x7b,
	0xa7, 0x1d, 0xd4, 0x3f, 0x19, 0xf5, 0xbf, 0x3b, 0x16, 0xb7, 0x0e, 0x43, 0x12, 0x4e, 0x70, 0x2a,
	0x13, 0xaa, 0xa1, 0x8d, 0x61, 0x70, 0x1c, 0xbc, 0xee, 0x21, 0x11, 0x31, 0x48, 0xe2, 0x90, 0xb9,
	0x55, 0x15, 0x31, 0x18, 0xf4, 0x83, 0x53, 0xe7, 0x00, 0x6c, 0x41, 0xce, 0x38, 0xe4, 0xa1, 0x7b,
	0xa7, 0x65, 0xed, 0xd5, 0xf7, 0x1f, 0xfa, 0xa2, 0x1e, 0xfe, 0x12, 0x7f, 0xc8, 0x1e, 0xf6, 0x46,
	0x41, 0x37, 0x18, 0x05, 0xce, 0x4b, 0xb0, 0x83, 0x28, 0xc2, 0x8c, 0x61, 0xe6, 0x6e, 0xb4, 0xd6,
	0xf6, 0xea, 0xfb, 0x4f, 0x56, 0x38, 0xf9, 0x06, 0x23, 0x25, 0xb4, 0x70, 0x11, 0x39, 0xbe, 0xc2,
	0x21, 0xcf, 0x52, 0xec, 0xda, 0x2a, 0x47, 0x2d, 0x36, 0xbf, 0x87, 0xcd, 0x92, 0x93, 0x20, 0xe6,
	0x2a, 0x27, 0xe6, 0x0a, 0xcf, 0x9d, 0x67, 0x50, 0x7d, 0x17, 0x26, 0x19, 0x96, 0xc4, 0xd4, 0xf7,
	0x1b, 0xe5, 0x8b, 0x95, 0x37, 0x52, 0x90, 0xe7, 0x95, 0xaf, 0x2c, 0xef, 0x27, 0x0b, 0xee, 0x95,
	0xad, 0x39, 0xb7, 0x96, 0xe6, 0x56, 0x08, 0x37, 0x71, 0x9b, 0xab, 0x9c, 0xe7, 0x60, 0x9f, 0xd0,
	0x24, 0x8e, 0x62, 0xcc, 0xdc, 0x8a, 0x7c, 0xf6, 0xe3, 0x95, 0xb7, 0x63, 0xc6, 0x25, 0x6e, 0x8e,
	0x16, 0x78, 0xef, 0x18, 0xdc, 0x9b, 0x50, 0xce, 0x2e, 0xdc, 0x09, 0x22, 0x79, 0xa9, 0x2a, 0xb6,
	0x96, 0x44, 0x77, 0x21, 0xcc, 0x68, 0x96, 0x46, 0x58, 0xa7, 0xb3, 0x90, 0xbd, 0x5f, 0x2c, 0xb8,
	0x6f, 0x02, 0x9e, 0x62, 0x85, 0xff, 0xff, 0xba, 0xa8, 0x0d, 0x76, 0xe7, 0x32, 0x4e, 0xc6, 0x29,
	0x26, 0xee, 0xba, 0x7c, 0xe9, 0xf6, 0x8a, 0x02, 0x23, 0xbb, 0xf3, 0xa6, 0x3f, 0xe8, 0xa2, 0xde,
	0xb1, 0xf7, 0xb3, 0x05, 0x8d, 0xf2, 0x24, 0xb2, 0x19, 0x25, 0x0c, 0x3b, 0x2f, 0xe0, 0x2e, 0xa2,
	0x94, 0x2f, 0x7a, 0xcc, 0xfa, 0x70, 0x8f, 0xdd, 0x3d, 0x3b, 0x9b, 0x6a, 0xe8, 0xd9, 0x99, 0xf3,
	0x05, 0xd8, 0xfa, 0x6d, 0x86, 0xf0, 0x9d, 0xb2, 0xa3, 0xb6, 0xa2, 0x05, 0xcc, 0xdb, 0x86, 0x07,
	0xaf, 0xc4, 0x7a, 0x39, 0xe5, 0x21, 0xc7, 0x66, 0x1f, 0x34, 0xc0, 0x29, 0x2a, 0x55, 0x6a, 0xde,
	0x53, 0xd8, 0x96, 0xda, 0x93, 0x24, 0x9b, 0xc4, 0x84, 0x69, 0xb0, 0xe3, 0xc0, 0xfa, 0x20, 0x24,
	0x13, 0x4d, 0xa3, 0x3c, 0x7b, 0xbb, 0xd0, 0x28, 0x43, 0x75, 0x88, 0x67, 0x5a, 0x3f, 0xc4, 0x8c,
	0x85, 0x13, 0xfc, 0xc1, 0x18, 0xbf, 0x59, 0xb0, 0x73, 0x0d, 0xac, 0x39, 0xea, 0x89, 0x19, 0x54,
	0x3a, 0xd7, 0x92, 0xcf, 0x7c, 0xaa, 0x9e, 0xb9, 0x12, 0xee, 0x1b, 0x85, 0xae, 0x81, 0x11, 0x9b,
	0x2f, 0x60, 0xb3, 0x64, 0x5a, 0x31, 0x3c, 0x8d, 0xe2, 0xf0, 0xd4, 0x8a, 0x63, 0xe2, 0xc2, 0xae,
	0xa2, 0x08, 0x33, 0x16, 0x53, 0xf2, 0x1a, 0x73, 0x43, 0xde, 0xd7, 0xf0, 0x70, 0xc9, 0xa2, 0x13,
	0x7f, 0x02, 0xd5, 0x11, 0xbd, 0xc2, 0x44, 0x57, 0xb5, 0xee, 0xcb, 0xd5, 0x2c, 0x55, 0x48, 0x59,
	0xbc, 0xbf, 0x2c, 0xcd, 0xb2, 0x76, 0x37, 0x0c, 0x3d, 0x06, 0xe8, 0x24, 0x31, 0x26, 0x7c, 0x14,
	0x4f, 0xb1, 0xf4, 0xaf, 0xa2, 0x82, 0xc6, 0xe9, 0x80, 0x1d, 0x64, 0xfc, 0xb2, 0x4f, 0x2e, 0xa8,
	0x2e, 0xfd, 0x27, 0x05, 0x4e, 0xca, 0xc1, 0x7c, 0x83, 0x34, 0x8b, 0x46, 0x8b, 0x62, 0xb0, 0x06,
	0x74, 0x42, 0x33, 0x2e, 0x7b, 0xdc, 0x46, 0x5a, 0x12, 0x4c, 0x95, 0x5c, 0x6e, 0xc5, 0xd4, 0x9f,
	0x15, 0x5d, 0xf4, 0x45, 0x12, 0x9a, 0x8d, 0x2d, 0x58, 0x3b, 0x7a, 0x3b, 0x32, 0x41, 0x8e, 0xde,
	0x8e, 0xc4, 0x23, 0x7b, 0xef, 0x67, 0x71, 0x8a, 0xe5, 0x23, 0x2b, 0xea, 0x91, 0xb9, 0x46, 0x2c,
	0xc2, 0x51, 0x1a, 0x4f, 0xc4, 0xb2, 0x56, 0x43, 0x68, 0x44, 0x67, 0x08, 0x75, 0x7d, 0x94, 0x0c,
	0xa8, 0x19, 0xfc, 0x74, 0x15, 0x03, 0xba, 0x29, 0x0a, 0x68, 0xc5, 0x42, 0xd1, 0x3f, 0x2f, 0x54,
	0xf5, 0xa6, 0x42, 0x89, 0x5c, 0x11, 0x1e, 0xc7, 0x29, 0x8e, 0xf8, 0x88, 0xca, 0xbf, 0x82, 0x1a,
	0x2a, 0x68, 0x04, 0x21, 0xbd, 0x34, 0xa5, 0xa9, 0xbb, 0xa1, 0x08, 0x91, 0x42, 0xf3, 0x1b, 0xd8,
	0xba, 0x7e, 0xf3, 0x7f, 0x69, 0xbb, 0x2e, 0x4e, 0x4c, 0xdb, 0x7d, 0x54, 0x6e, 0x3b, 0x69, 0xd1,
	0x53, 0xf7, 0x39, 0x6c, 0x49, 0x93, 0xa8, 0xa1, 0xe9, 0xa7, 0x47, 0x50, 0xd3, 0xc7, 0x7e, 0x57,
	0x5f, 0x9d, 0x2b, 0x16, 0x5b, 0x41, 0x79, 0xe8, 0x30, 0x7f, 0x58, 0xfa, 0xf2, 0x1e, 0x49, 0x69,
	0x92, 0x14, 0xa3, 0x89, 0xc2, 0x49, 0xe5, 0x68, 0x3e, 0xc3, 0xe6, 0x5b, 0x21, 0xd7, 0x38, 0x03,
	0x63, 0x2f, 0xf4, 0xe7, 0x67, 0x85, 0xea, 0x2c, 0x45, 0xf4, 0x73, 0xb8, 0x2a, 0x4f, 0xc1, 0xbf,
	0xf9, 0x12, 0xee, 0x5f, 0x33, 0xdf, 0x8a, 0xc3, 0x5f, 0x2d, 0x4d, 0x55, 0xf1, 0xd6, 0xc5, 0xfa,
	0x5d, 0x97, 0x29, 0x5a, 0x4b, 0x23, 0xb4, 0x0c, 0xf6, 0xf3, 0xec, 0xa4, 0x53, 0xf3, 0x4b, 0xa8,
	0xfd, 0xbb, 0x8c, 0xde, 0xe8, 0x7d, 0x7b, 0x18, 0x93, 0x30, 0x9d, 0x17, 0x48, 0x55, 0x8a, 0x22,
	0xa9, 0xb9, 0x46, 0x2c, 0xcd, 0x1f, 0xb2, 0x78, 0xac, 0xc3, 0xc9, 0xb3, 0xf8, 0xc0, 0x2b, 0x45,
	0xd2, 0xa5, 0x33, 0xfb, 0xf8, 0x90, 0x52, 0xde, 0xa1, 0xe4, 0xc2, 0x34, 0xcd, 0x62, 0xc7, 0xe6,
	0x86, 0x7c, 0xc7, 0x1e, 0x31, 0x4a, 0xba, 0xea, 0x3f, 0xe8, 0xfa, 0x8e, 0xbd, 0x0e, 0xf7, 0x0d,
	0x56, 0x6f, 0x14, 0x23, 0x8a, 0xcd, 0x51, 0x32, 0xdd, 0x86, 0x96, 0xf3, 0x3b, 0xf2, 0xa3, 0xf5,
	0xe0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xd9, 0x35, 0x6f, 0x01, 0x0b, 0x00, 0x00,
}
