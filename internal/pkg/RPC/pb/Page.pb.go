// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: Page.proto

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

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string       `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Route         string       `protobuf:"bytes,2,opt,name=Route,proto3" json:"Route,omitempty"`
	Template      *Template    `protobuf:"bytes,3,opt,name=Template,proto3" json:"Template,omitempty"`
	UserAccountId string       `protobuf:"bytes,4,opt,name=UserAccountId,proto3" json:"UserAccountId,omitempty"`
	Links         []*PageLinks `protobuf:"bytes,5,rep,name=Links,proto3" json:"Links,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Page) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *Page) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *Page) GetUserAccountId() string {
	if x != nil {
		return x.UserAccountId
	}
	return ""
}

func (x *Page) GetLinks() []*PageLinks {
	if x != nil {
		return x.Links
	}
	return nil
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name           string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc           string  `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Image          string  `protobuf:"bytes,4,opt,name=Image,proto3" json:"Image,omitempty"`
	Button         string  `protobuf:"bytes,5,opt,name=Button,proto3" json:"Button,omitempty"`
	Background     string  `protobuf:"bytes,6,opt,name=Background,proto3" json:"Background,omitempty"`
	Font           string  `protobuf:"bytes,7,opt,name=Font,proto3" json:"Font,omitempty"`
	FontColor      string  `protobuf:"bytes,8,opt,name=FontColor,proto3" json:"FontColor,omitempty"`
	MetaTags       []*Meta `protobuf:"bytes,9,rep,name=MetaTags,proto3" json:"MetaTags,omitempty"`
	PageId         string  `protobuf:"bytes,10,opt,name=PageId,proto3" json:"PageId,omitempty"`
	Social         bool    `protobuf:"varint,11,opt,name=Social,proto3" json:"Social,omitempty"`
	SocialPosition string  `protobuf:"bytes,12,opt,name=SocialPosition,proto3" json:"SocialPosition,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{1}
}

func (x *Template) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Template) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Template) GetButton() string {
	if x != nil {
		return x.Button
	}
	return ""
}

func (x *Template) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *Template) GetFont() string {
	if x != nil {
		return x.Font
	}
	return ""
}

func (x *Template) GetFontColor() string {
	if x != nil {
		return x.FontColor
	}
	return ""
}

func (x *Template) GetMetaTags() []*Meta {
	if x != nil {
		return x.MetaTags
	}
	return nil
}

func (x *Template) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *Template) GetSocial() bool {
	if x != nil {
		return x.Social
	}
	return false
}

func (x *Template) GetSocialPosition() string {
	if x != nil {
		return x.SocialPosition
	}
	return ""
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Value      string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	TemplateId string `protobuf:"bytes,4,opt,name=TemplateId,proto3" json:"TemplateId,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meta) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Meta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Meta) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

type PageLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Link   string `protobuf:"bytes,3,opt,name=Link,proto3" json:"Link,omitempty"`
	Icon   string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Social bool   `protobuf:"varint,5,opt,name=Social,proto3" json:"Social,omitempty"`
	PageId string `protobuf:"bytes,6,opt,name=PageId,proto3" json:"PageId,omitempty"`
}

func (x *PageLinks) Reset() {
	*x = PageLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageLinks) ProtoMessage() {}

func (x *PageLinks) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageLinks.ProtoReflect.Descriptor instead.
func (*PageLinks) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{3}
}

func (x *PageLinks) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PageLinks) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PageLinks) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *PageLinks) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *PageLinks) GetSocial() bool {
	if x != nil {
		return x.Social
	}
	return false
}

func (x *PageLinks) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{4}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route         string `protobuf:"bytes,1,opt,name=Route,proto3" json:"Route,omitempty"`
	UserAccountId string `protobuf:"bytes,2,opt,name=UserAccountId,proto3" json:"UserAccountId,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{5}
}

func (x *PageRequest) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *PageRequest) GetUserAccountId() string {
	if x != nil {
		return x.UserAccountId
	}
	return ""
}

type TemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc           string `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Image          string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	Button         string `protobuf:"bytes,4,opt,name=Button,proto3" json:"Button,omitempty"`
	Background     string `protobuf:"bytes,5,opt,name=Background,proto3" json:"Background,omitempty"`
	Font           string `protobuf:"bytes,6,opt,name=Font,proto3" json:"Font,omitempty"`
	FontColor      string `protobuf:"bytes,7,opt,name=FontColor,proto3" json:"FontColor,omitempty"`
	PageId         string `protobuf:"bytes,8,opt,name=PageId,proto3" json:"PageId,omitempty"`
	Social         bool   `protobuf:"varint,9,opt,name=Social,proto3" json:"Social,omitempty"`
	SocialPosition string `protobuf:"bytes,10,opt,name=SocialPosition,proto3" json:"SocialPosition,omitempty"`
}

func (x *TemplateRequest) Reset() {
	*x = TemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRequest) ProtoMessage() {}

func (x *TemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRequest.ProtoReflect.Descriptor instead.
func (*TemplateRequest) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{6}
}

func (x *TemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *TemplateRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *TemplateRequest) GetButton() string {
	if x != nil {
		return x.Button
	}
	return ""
}

func (x *TemplateRequest) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *TemplateRequest) GetFont() string {
	if x != nil {
		return x.Font
	}
	return ""
}

func (x *TemplateRequest) GetFontColor() string {
	if x != nil {
		return x.FontColor
	}
	return ""
}

func (x *TemplateRequest) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

func (x *TemplateRequest) GetSocial() bool {
	if x != nil {
		return x.Social
	}
	return false
}

func (x *TemplateRequest) GetSocialPosition() string {
	if x != nil {
		return x.SocialPosition
	}
	return ""
}

type CreateLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Link         string `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"`
	Icon         string `protobuf:"bytes,3,opt,name=Icon,proto3" json:"Icon,omitempty"`
	IsSocialIcon bool   `protobuf:"varint,4,opt,name=IsSocialIcon,proto3" json:"IsSocialIcon,omitempty"`
	PageId       string `protobuf:"bytes,5,opt,name=pageId,proto3" json:"pageId,omitempty"`
}

func (x *CreateLinkRequest) Reset() {
	*x = CreateLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkRequest) ProtoMessage() {}

func (x *CreateLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{7}
}

func (x *CreateLinkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateLinkRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateLinkRequest) GetIsSocialIcon() bool {
	if x != nil {
		return x.IsSocialIcon
	}
	return false
}

func (x *CreateLinkRequest) GetPageId() string {
	if x != nil {
		return x.PageId
	}
	return ""
}

type IdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResponse) Reset() {
	*x = IdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponse) ProtoMessage() {}

func (x *IdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponse.ProtoReflect.Descriptor instead.
func (*IdResponse) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{8}
}

func (x *IdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VoidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidResponse) Reset() {
	*x = VoidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Page_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidResponse) ProtoMessage() {}

func (x *VoidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Page_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidResponse.ProtoReflect.Descriptor instead.
func (*VoidResponse) Descriptor() ([]byte, []int) {
	return file_Page_proto_rawDescGZIP(), []int{9}
}

var File_Page_proto protoreflect.FileDescriptor

var file_Page_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0xca, 0x02, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6f, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x6f, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x6f, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x6f,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x54,
	0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x4d,
	0x65, 0x74, 0x61, 0x54, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x60, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x46, 0x6f, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x6f, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x6f, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x6f, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12,
	0x26, 0x0a, 0x0e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x49, 0x73, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x9a, 0x04, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1a, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x17, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x40,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x12, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Page_proto_rawDescOnce sync.Once
	file_Page_proto_rawDescData = file_Page_proto_rawDesc
)

func file_Page_proto_rawDescGZIP() []byte {
	file_Page_proto_rawDescOnce.Do(func() {
		file_Page_proto_rawDescData = protoimpl.X.CompressGZIP(file_Page_proto_rawDescData)
	})
	return file_Page_proto_rawDescData
}

var file_Page_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Page_proto_goTypes = []interface{}{
	(*Page)(nil),              // 0: page_package.Page
	(*Template)(nil),          // 1: page_package.Template
	(*Meta)(nil),              // 2: page_package.Meta
	(*PageLinks)(nil),         // 3: page_package.PageLinks
	(*IdRequest)(nil),         // 4: page_package.IdRequest
	(*PageRequest)(nil),       // 5: page_package.PageRequest
	(*TemplateRequest)(nil),   // 6: page_package.TemplateRequest
	(*CreateLinkRequest)(nil), // 7: page_package.CreateLinkRequest
	(*IdResponse)(nil),        // 8: page_package.IdResponse
	(*VoidResponse)(nil),      // 9: page_package.VoidResponse
}
var file_Page_proto_depIdxs = []int32{
	1,  // 0: page_package.Page.Template:type_name -> page_package.Template
	3,  // 1: page_package.Page.Links:type_name -> page_package.PageLinks
	2,  // 2: page_package.Template.MetaTags:type_name -> page_package.Meta
	4,  // 3: page_package.PagePackage.GetPage:input_type -> page_package.IdRequest
	5,  // 4: page_package.PagePackage.CreatePage:input_type -> page_package.PageRequest
	6,  // 5: page_package.PagePackage.CreateTemplate:input_type -> page_package.TemplateRequest
	8,  // 6: page_package.PagePackage.UpdateTemplate:input_type -> page_package.IdResponse
	7,  // 7: page_package.PagePackage.CreateLink:input_type -> page_package.CreateLinkRequest
	3,  // 8: page_package.PagePackage.UpdateLink:input_type -> page_package.PageLinks
	2,  // 9: page_package.PagePackage.CreateMetaLink:input_type -> page_package.Meta
	2,  // 10: page_package.PagePackage.UpdateMetaLink:input_type -> page_package.Meta
	0,  // 11: page_package.PagePackage.GetPage:output_type -> page_package.Page
	0,  // 12: page_package.PagePackage.CreatePage:output_type -> page_package.Page
	1,  // 13: page_package.PagePackage.CreateTemplate:output_type -> page_package.Template
	9,  // 14: page_package.PagePackage.UpdateTemplate:output_type -> page_package.VoidResponse
	3,  // 15: page_package.PagePackage.CreateLink:output_type -> page_package.PageLinks
	9,  // 16: page_package.PagePackage.UpdateLink:output_type -> page_package.VoidResponse
	2,  // 17: page_package.PagePackage.CreateMetaLink:output_type -> page_package.Meta
	9,  // 18: page_package.PagePackage.UpdateMetaLink:output_type -> page_package.VoidResponse
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_Page_proto_init() }
func file_Page_proto_init() {
	if File_Page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_Page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_Page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_Page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageLinks); i {
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
		file_Page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_Page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_Page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRequest); i {
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
		file_Page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkRequest); i {
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
		file_Page_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdResponse); i {
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
		file_Page_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidResponse); i {
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
			RawDescriptor: file_Page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Page_proto_goTypes,
		DependencyIndexes: file_Page_proto_depIdxs,
		MessageInfos:      file_Page_proto_msgTypes,
	}.Build()
	File_Page_proto = out.File
	file_Page_proto_rawDesc = nil
	file_Page_proto_goTypes = nil
	file_Page_proto_depIdxs = nil
}
