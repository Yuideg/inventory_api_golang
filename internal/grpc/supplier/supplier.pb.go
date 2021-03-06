// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: supplier/supplier.proto

package supplier

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

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SupplierName string  `protobuf:"bytes,2,opt,name=supplier_name,json=supplierName,proto3" json:"supplier_name,omitempty"`
	Country      string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	City         string  `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State        string  `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Zipcode      int32   `protobuf:"varint,6,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	Street       string  `protobuf:"bytes,7,opt,name=street,proto3" json:"street,omitempty"`
	Latitude     float32 `protobuf:"fixed32,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude    float32 `protobuf:"fixed32,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Fax          string  `protobuf:"bytes,10,opt,name=fax,proto3" json:"fax,omitempty"`
	Pobox        string  `protobuf:"bytes,11,opt,name=pobox,proto3" json:"pobox,omitempty"`
	Email        string  `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Website      string  `protobuf:"bytes,13,opt,name=website,proto3" json:"website,omitempty"`
	Phone        string  `protobuf:"bytes,14,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedOn    string  `protobuf:"bytes,15,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn    string  `protobuf:"bytes,16,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *Supplier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supplier) GetSupplierName() string {
	if x != nil {
		return x.SupplierName
	}
	return ""
}

func (x *Supplier) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Supplier) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Supplier) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Supplier) GetZipcode() int32 {
	if x != nil {
		return x.Zipcode
	}
	return 0
}

func (x *Supplier) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Supplier) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Supplier) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Supplier) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *Supplier) GetPobox() string {
	if x != nil {
		return x.Pobox
	}
	return ""
}

func (x *Supplier) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Supplier) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Supplier) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Supplier) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Supplier) GetUpdatedOn() string {
	if x != nil {
		return x.UpdatedOn
	}
	return ""
}

type CreateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *CreateSupplierRequest) Reset() {
	*x = CreateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierRequest) ProtoMessage() {}

func (x *CreateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierRequest.ProtoReflect.Descriptor instead.
func (*CreateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSupplierRequest) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type CreateSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSupplierResponse) Reset() {
	*x = CreateSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierResponse) ProtoMessage() {}

func (x *CreateSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierResponse.ProtoReflect.Descriptor instead.
func (*CreateSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSupplierResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupplierByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SupplierByIDRequest) Reset() {
	*x = SupplierByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierByIDRequest) ProtoMessage() {}

func (x *SupplierByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierByIDRequest.ProtoReflect.Descriptor instead.
func (*SupplierByIDRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *SupplierByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupplierByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *SupplierByIDResponse) Reset() {
	*x = SupplierByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierByIDResponse) ProtoMessage() {}

func (x *SupplierByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierByIDResponse.ProtoReflect.Descriptor instead.
func (*SupplierByIDResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *SupplierByIDResponse) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type SupplierBySupplierIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SupplierBySupplierIDRequest) Reset() {
	*x = SupplierBySupplierIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierBySupplierIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierBySupplierIDRequest) ProtoMessage() {}

func (x *SupplierBySupplierIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierBySupplierIDRequest.ProtoReflect.Descriptor instead.
func (*SupplierBySupplierIDRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{5}
}

func (x *SupplierBySupplierIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupplierBySupplierIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *SupplierBySupplierIDResponse) Reset() {
	*x = SupplierBySupplierIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierBySupplierIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierBySupplierIDResponse) ProtoMessage() {}

func (x *SupplierBySupplierIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierBySupplierIDResponse.ProtoReflect.Descriptor instead.
func (*SupplierBySupplierIDResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{6}
}

func (x *SupplierBySupplierIDResponse) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type SupplierListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SupplierListRequest) Reset() {
	*x = SupplierListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierListRequest) ProtoMessage() {}

func (x *SupplierListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierListRequest.ProtoReflect.Descriptor instead.
func (*SupplierListRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{7}
}

type SupplierListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier []*Supplier `protobuf:"bytes,1,rep,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *SupplierListResponse) Reset() {
	*x = SupplierListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierListResponse) ProtoMessage() {}

func (x *SupplierListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierListResponse.ProtoReflect.Descriptor instead.
func (*SupplierListResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{8}
}

func (x *SupplierListResponse) GetSupplier() []*Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type UpdateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *UpdateSupplierRequest) Reset() {
	*x = UpdateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierRequest) ProtoMessage() {}

func (x *UpdateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateSupplierRequest) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type UpdateSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSupplierResponse) Reset() {
	*x = UpdateSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierResponse) ProtoMessage() {}

func (x *UpdateSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierResponse.ProtoReflect.Descriptor instead.
func (*UpdateSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{10}
}

type DeleteSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSupplierRequest) Reset() {
	*x = DeleteSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierRequest) ProtoMessage() {}

func (x *DeleteSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierRequest.ProtoReflect.Descriptor instead.
func (*DeleteSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteSupplierRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSupplierResponse) Reset() {
	*x = DeleteSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_supplier_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierResponse) ProtoMessage() {}

func (x *DeleteSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_supplier_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierResponse.ProtoReflect.Descriptor instead.
func (*DeleteSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_supplier_proto_rawDescGZIP(), []int{12}
}

var File_supplier_supplier_proto protoreflect.FileDescriptor

var file_supplier_supplier_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x22, 0x9b, 0x03, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x66, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x62, 0x6f, 0x78,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x62, 0x6f, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f,
	0x6e, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x1b, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42,
	0x79, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x4e, 0x0a, 0x1c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x04, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x1d, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x68, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supplier_supplier_proto_rawDescOnce sync.Once
	file_supplier_supplier_proto_rawDescData = file_supplier_supplier_proto_rawDesc
)

func file_supplier_supplier_proto_rawDescGZIP() []byte {
	file_supplier_supplier_proto_rawDescOnce.Do(func() {
		file_supplier_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_supplier_supplier_proto_rawDescData)
	})
	return file_supplier_supplier_proto_rawDescData
}

var file_supplier_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_supplier_supplier_proto_goTypes = []interface{}{
	(*Supplier)(nil),                     // 0: supplier.Supplier
	(*CreateSupplierRequest)(nil),        // 1: supplier.CreateSupplierRequest
	(*CreateSupplierResponse)(nil),       // 2: supplier.CreateSupplierResponse
	(*SupplierByIDRequest)(nil),          // 3: supplier.SupplierByIDRequest
	(*SupplierByIDResponse)(nil),         // 4: supplier.SupplierByIDResponse
	(*SupplierBySupplierIDRequest)(nil),  // 5: supplier.SupplierBySupplierIDRequest
	(*SupplierBySupplierIDResponse)(nil), // 6: supplier.SupplierBySupplierIDResponse
	(*SupplierListRequest)(nil),          // 7: supplier.SupplierListRequest
	(*SupplierListResponse)(nil),         // 8: supplier.SupplierListResponse
	(*UpdateSupplierRequest)(nil),        // 9: supplier.UpdateSupplierRequest
	(*UpdateSupplierResponse)(nil),       // 10: supplier.UpdateSupplierResponse
	(*DeleteSupplierRequest)(nil),        // 11: supplier.DeleteSupplierRequest
	(*DeleteSupplierResponse)(nil),       // 12: supplier.DeleteSupplierResponse
}
var file_supplier_supplier_proto_depIdxs = []int32{
	0,  // 0: supplier.CreateSupplierRequest.supplier:type_name -> supplier.Supplier
	0,  // 1: supplier.SupplierByIDResponse.supplier:type_name -> supplier.Supplier
	0,  // 2: supplier.SupplierBySupplierIDResponse.supplier:type_name -> supplier.Supplier
	0,  // 3: supplier.SupplierListResponse.supplier:type_name -> supplier.Supplier
	0,  // 4: supplier.UpdateSupplierRequest.supplier:type_name -> supplier.Supplier
	1,  // 5: supplier.SupplierService.CreateSupplier:input_type -> supplier.CreateSupplierRequest
	3,  // 6: supplier.SupplierService.SupplierByID:input_type -> supplier.SupplierByIDRequest
	5,  // 7: supplier.SupplierService.GetSupplierBySupplierID:input_type -> supplier.SupplierBySupplierIDRequest
	7,  // 8: supplier.SupplierService.SupplierList:input_type -> supplier.SupplierListRequest
	9,  // 9: supplier.SupplierService.UpdateSupplier:input_type -> supplier.UpdateSupplierRequest
	11, // 10: supplier.SupplierService.DeleteSupplier:input_type -> supplier.DeleteSupplierRequest
	2,  // 11: supplier.SupplierService.CreateSupplier:output_type -> supplier.CreateSupplierResponse
	4,  // 12: supplier.SupplierService.SupplierByID:output_type -> supplier.SupplierByIDResponse
	6,  // 13: supplier.SupplierService.GetSupplierBySupplierID:output_type -> supplier.SupplierBySupplierIDResponse
	8,  // 14: supplier.SupplierService.SupplierList:output_type -> supplier.SupplierListResponse
	10, // 15: supplier.SupplierService.UpdateSupplier:output_type -> supplier.UpdateSupplierResponse
	12, // 16: supplier.SupplierService.DeleteSupplier:output_type -> supplier.DeleteSupplierResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_supplier_supplier_proto_init() }
func file_supplier_supplier_proto_init() {
	if File_supplier_supplier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supplier_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
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
		file_supplier_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplierRequest); i {
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
		file_supplier_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplierResponse); i {
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
		file_supplier_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierByIDRequest); i {
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
		file_supplier_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierByIDResponse); i {
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
		file_supplier_supplier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierBySupplierIDRequest); i {
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
		file_supplier_supplier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierBySupplierIDResponse); i {
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
		file_supplier_supplier_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierListRequest); i {
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
		file_supplier_supplier_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierListResponse); i {
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
		file_supplier_supplier_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplierRequest); i {
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
		file_supplier_supplier_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplierResponse); i {
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
		file_supplier_supplier_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierRequest); i {
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
		file_supplier_supplier_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierResponse); i {
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
			RawDescriptor: file_supplier_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supplier_supplier_proto_goTypes,
		DependencyIndexes: file_supplier_supplier_proto_depIdxs,
		MessageInfos:      file_supplier_supplier_proto_msgTypes,
	}.Build()
	File_supplier_supplier_proto = out.File
	file_supplier_supplier_proto_rawDesc = nil
	file_supplier_supplier_proto_goTypes = nil
	file_supplier_supplier_proto_depIdxs = nil
}
