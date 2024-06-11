// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: reservation_orders.proto

package reservation

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

type ReservationOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReservationId string `protobuf:"bytes,2,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	MenuItemId    string `protobuf:"bytes,3,opt,name=menu_item_id,json=menuItemId,proto3" json:"menu_item_id,omitempty"`
	Quantity      string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ReservationOrderReq) Reset() {
	*x = ReservationOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationOrderReq) ProtoMessage() {}

func (x *ReservationOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationOrderReq.ProtoReflect.Descriptor instead.
func (*ReservationOrderReq) Descriptor() ([]byte, []int) {
	return file_reservation_orders_proto_rawDescGZIP(), []int{0}
}

func (x *ReservationOrderReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationOrderReq) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *ReservationOrderReq) GetMenuItemId() string {
	if x != nil {
		return x.MenuItemId
	}
	return ""
}

func (x *ReservationOrderReq) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type ReservationOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reservation *ReservationRes `protobuf:"bytes,2,opt,name=reservation,proto3" json:"reservation,omitempty"`
	MenuItem    *MenuRes        `protobuf:"bytes,3,opt,name=menu_item,json=menuItem,proto3" json:"menu_item,omitempty"`
	Quantity    string          `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ReservationOrderRes) Reset() {
	*x = ReservationOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationOrderRes) ProtoMessage() {}

func (x *ReservationOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationOrderRes.ProtoReflect.Descriptor instead.
func (*ReservationOrderRes) Descriptor() ([]byte, []int) {
	return file_reservation_orders_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationOrderRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationOrderRes) GetReservation() *ReservationRes {
	if x != nil {
		return x.Reservation
	}
	return nil
}

func (x *ReservationOrderRes) GetMenuItem() *MenuRes {
	if x != nil {
		return x.MenuItem
	}
	return nil
}

func (x *ReservationOrderRes) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type GetAllReservationOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price  float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	Filter *Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllReservationOrderReq) Reset() {
	*x = GetAllReservationOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReservationOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReservationOrderReq) ProtoMessage() {}

func (x *GetAllReservationOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReservationOrderReq.ProtoReflect.Descriptor instead.
func (*GetAllReservationOrderReq) Descriptor() ([]byte, []int) {
	return file_reservation_orders_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllReservationOrderReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetAllReservationOrderReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAllReservationOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationOrder []*ReservationOrderRes `protobuf:"bytes,1,rep,name=reservation_order,json=reservationOrder,proto3" json:"reservation_order,omitempty"`
}

func (x *GetAllReservationOrderRes) Reset() {
	*x = GetAllReservationOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReservationOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReservationOrderRes) ProtoMessage() {}

func (x *GetAllReservationOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReservationOrderRes.ProtoReflect.Descriptor instead.
func (*GetAllReservationOrderRes) Descriptor() ([]byte, []int) {
	return file_reservation_orders_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllReservationOrderRes) GetReservationOrder() []*ReservationOrderRes {
	if x != nil {
		return x.ReservationOrder
	}
	return nil
}

type ReservationOrderUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReservationId string `protobuf:"bytes,2,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	MenuItemId    string `protobuf:"bytes,3,opt,name=menu_item_id,json=menuItemId,proto3" json:"menu_item_id,omitempty"`
	Quantity      string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ReservationOrderUpdate) Reset() {
	*x = ReservationOrderUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationOrderUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationOrderUpdate) ProtoMessage() {}

func (x *ReservationOrderUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationOrderUpdate.ProtoReflect.Descriptor instead.
func (*ReservationOrderUpdate) Descriptor() ([]byte, []int) {
	return file_reservation_orders_proto_rawDescGZIP(), []int{4}
}

func (x *ReservationOrderUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationOrderUpdate) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *ReservationOrderUpdate) GetMenuItemId() string {
	if x != nil {
		return x.MenuItemId
	}
	return ""
}

func (x *ReservationOrderUpdate) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

var File_reservation_orders_proto protoreflect.FileDescriptor

var file_reservation_orders_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xb3, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x5e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x52, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6e,
	0x75, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x32, 0x94, 0x03, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x20, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x26, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x20,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_orders_proto_rawDescOnce sync.Once
	file_reservation_orders_proto_rawDescData = file_reservation_orders_proto_rawDesc
)

func file_reservation_orders_proto_rawDescGZIP() []byte {
	file_reservation_orders_proto_rawDescOnce.Do(func() {
		file_reservation_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_orders_proto_rawDescData)
	})
	return file_reservation_orders_proto_rawDescData
}

var file_reservation_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reservation_orders_proto_goTypes = []interface{}{
	(*ReservationOrderReq)(nil),       // 0: reservation.ReservationOrderReq
	(*ReservationOrderRes)(nil),       // 1: reservation.ReservationOrderRes
	(*GetAllReservationOrderReq)(nil), // 2: reservation.GetAllReservationOrderReq
	(*GetAllReservationOrderRes)(nil), // 3: reservation.GetAllReservationOrderRes
	(*ReservationOrderUpdate)(nil),    // 4: reservation.ReservationOrderUpdate
	(*ReservationRes)(nil),            // 5: reservation.ReservationRes
	(*MenuRes)(nil),                   // 6: reservation.MenuRes
	(*Filter)(nil),                    // 7: reservation.Filter
	(*GetByIdReq)(nil),                // 8: reservation.GetByIdReq
	(*Void)(nil),                      // 9: reservation.Void
}
var file_reservation_orders_proto_depIdxs = []int32{
	5, // 0: reservation.ReservationOrderRes.reservation:type_name -> reservation.ReservationRes
	6, // 1: reservation.ReservationOrderRes.menu_item:type_name -> reservation.MenuRes
	7, // 2: reservation.GetAllReservationOrderReq.filter:type_name -> reservation.Filter
	1, // 3: reservation.GetAllReservationOrderRes.reservation_order:type_name -> reservation.ReservationOrderRes
	0, // 4: reservation.ReservationOrderService.Create:input_type -> reservation.ReservationOrderReq
	8, // 5: reservation.ReservationOrderService.Get:input_type -> reservation.GetByIdReq
	2, // 6: reservation.ReservationOrderService.GetAll:input_type -> reservation.GetAllReservationOrderReq
	4, // 7: reservation.ReservationOrderService.Update:input_type -> reservation.ReservationOrderUpdate
	8, // 8: reservation.ReservationOrderService.Delete:input_type -> reservation.GetByIdReq
	1, // 9: reservation.ReservationOrderService.Create:output_type -> reservation.ReservationOrderRes
	1, // 10: reservation.ReservationOrderService.Get:output_type -> reservation.ReservationOrderRes
	3, // 11: reservation.ReservationOrderService.GetAll:output_type -> reservation.GetAllReservationOrderRes
	1, // 12: reservation.ReservationOrderService.Update:output_type -> reservation.ReservationOrderRes
	9, // 13: reservation.ReservationOrderService.Delete:output_type -> reservation.Void
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_reservation_orders_proto_init() }
func file_reservation_orders_proto_init() {
	if File_reservation_orders_proto != nil {
		return
	}
	file_common_proto_init()
	file_menu_proto_init()
	file_reservation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reservation_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationOrderReq); i {
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
		file_reservation_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationOrderRes); i {
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
		file_reservation_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReservationOrderReq); i {
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
		file_reservation_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReservationOrderRes); i {
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
		file_reservation_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationOrderUpdate); i {
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
			RawDescriptor: file_reservation_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_orders_proto_goTypes,
		DependencyIndexes: file_reservation_orders_proto_depIdxs,
		MessageInfos:      file_reservation_orders_proto_msgTypes,
	}.Build()
	File_reservation_orders_proto = out.File
	file_reservation_orders_proto_rawDesc = nil
	file_reservation_orders_proto_goTypes = nil
	file_reservation_orders_proto_depIdxs = nil
}
