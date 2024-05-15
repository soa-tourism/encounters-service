// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0--rc2
// source: encounter.proto

package encounter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EncounterId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EncounterId) Reset() {
	*x = EncounterId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterId) ProtoMessage() {}

func (x *EncounterId) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterId.ProtoReflect.Descriptor instead.
func (*EncounterId) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{0}
}

func (x *EncounterId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encounter *EncounterDto `protobuf:"bytes,1,opt,name=encounter,proto3" json:"encounter,omitempty"`
	ImageF    [][]byte      `protobuf:"bytes,2,rep,name=image_f,json=imageF,proto3" json:"image_f,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetEncounter() *EncounterDto {
	if x != nil {
		return x.Encounter
	}
	return nil
}

func (x *UpdateRequest) GetImageF() [][]byte {
	if x != nil {
		return x.ImageF
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encounter            *EncounterDto `protobuf:"bytes,1,opt,name=encounter,proto3" json:"encounter,omitempty"`
	CheckpointId         int64         `protobuf:"varint,2,opt,name=checkpoint_id,json=checkpointId,proto3" json:"checkpoint_id,omitempty"`
	IsSecretPrerequisite bool          `protobuf:"varint,3,opt,name=is_secret_prerequisite,json=isSecretPrerequisite,proto3" json:"is_secret_prerequisite,omitempty"`
	ImageF               [][]byte      `protobuf:"bytes,4,rep,name=image_f,json=imageF,proto3" json:"image_f,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetEncounter() *EncounterDto {
	if x != nil {
		return x.Encounter
	}
	return nil
}

func (x *CreateRequest) GetCheckpointId() int64 {
	if x != nil {
		return x.CheckpointId
	}
	return 0
}

func (x *CreateRequest) GetIsSecretPrerequisite() bool {
	if x != nil {
		return x.IsSecretPrerequisite
	}
	return false
}

func (x *CreateRequest) GetImageF() [][]byte {
	if x != nil {
		return x.ImageF
	}
	return nil
}

type EncounterDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId          int64   `protobuf:"varint,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Id                int32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name              string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description       string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XP                int32   `protobuf:"varint,5,opt,name=xP,proto3" json:"xP,omitempty"`
	Status            string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type              string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Longitude         float64 `protobuf:"fixed64,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude          float64 `protobuf:"fixed64,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	LocationLongitude float64 `protobuf:"fixed64,10,opt,name=locationLongitude,proto3" json:"locationLongitude,omitempty"`
	LocationLatitude  float64 `protobuf:"fixed64,11,opt,name=locationLatitude,proto3" json:"locationLatitude,omitempty"`
	Image             string  `protobuf:"bytes,12,opt,name=image,proto3" json:"image,omitempty"`
	Range             float64 `protobuf:"fixed64,13,opt,name=range,proto3" json:"range,omitempty"`
	RequiredPeople    int32   `protobuf:"varint,14,opt,name=requiredPeople,proto3" json:"requiredPeople,omitempty"`
	ActiveTouristsIds []int32 `protobuf:"varint,15,rep,packed,name=activeTouristsIds,proto3" json:"activeTouristsIds,omitempty"`
}

func (x *EncounterDto) Reset() {
	*x = EncounterDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterDto) ProtoMessage() {}

func (x *EncounterDto) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterDto.ProtoReflect.Descriptor instead.
func (*EncounterDto) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{3}
}

func (x *EncounterDto) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *EncounterDto) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EncounterDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EncounterDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EncounterDto) GetXP() int32 {
	if x != nil {
		return x.XP
	}
	return 0
}

func (x *EncounterDto) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EncounterDto) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EncounterDto) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *EncounterDto) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EncounterDto) GetLocationLongitude() float64 {
	if x != nil {
		return x.LocationLongitude
	}
	return 0
}

func (x *EncounterDto) GetLocationLatitude() float64 {
	if x != nil {
		return x.LocationLatitude
	}
	return 0
}

func (x *EncounterDto) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *EncounterDto) GetRange() float64 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *EncounterDto) GetRequiredPeople() int32 {
	if x != nil {
		return x.RequiredPeople
	}
	return 0
}

func (x *EncounterDto) GetActiveTouristsIds() []int32 {
	if x != nil {
		return x.ActiveTouristsIds
	}
	return nil
}

var File_encounter_proto protoreflect.FileDescriptor

var file_encounter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f,
	0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x46, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x22, 0xc2, 0x03, 0x0a, 0x0c, 0x45, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x50,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x73,
	0x49, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x73, 0x49, 0x64, 0x73, 0x32, 0xbd, 0x01, 0x0a,
	0x09, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x45, 0x6e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e,
	0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encounter_proto_rawDescOnce sync.Once
	file_encounter_proto_rawDescData = file_encounter_proto_rawDesc
)

func file_encounter_proto_rawDescGZIP() []byte {
	file_encounter_proto_rawDescOnce.Do(func() {
		file_encounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_encounter_proto_rawDescData)
	})
	return file_encounter_proto_rawDescData
}

var file_encounter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_encounter_proto_goTypes = []interface{}{
	(*EncounterId)(nil),   // 0: EncounterId
	(*UpdateRequest)(nil), // 1: UpdateRequest
	(*CreateRequest)(nil), // 2: CreateRequest
	(*EncounterDto)(nil),  // 3: EncounterDto
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_encounter_proto_depIdxs = []int32{
	3, // 0: UpdateRequest.encounter:type_name -> EncounterDto
	3, // 1: CreateRequest.encounter:type_name -> EncounterDto
	2, // 2: Encounter.Create:input_type -> CreateRequest
	1, // 3: Encounter.Update:input_type -> UpdateRequest
	0, // 4: Encounter.Delete:input_type -> EncounterId
	0, // 5: Encounter.GetById:input_type -> EncounterId
	3, // 6: Encounter.Create:output_type -> EncounterDto
	3, // 7: Encounter.Update:output_type -> EncounterDto
	4, // 8: Encounter.Delete:output_type -> google.protobuf.Empty
	3, // 9: Encounter.GetById:output_type -> EncounterDto
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_encounter_proto_init() }
func file_encounter_proto_init() {
	if File_encounter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encounter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterId); i {
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
		file_encounter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_encounter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_encounter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterDto); i {
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
			RawDescriptor: file_encounter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encounter_proto_goTypes,
		DependencyIndexes: file_encounter_proto_depIdxs,
		MessageInfos:      file_encounter_proto_msgTypes,
	}.Build()
	File_encounter_proto = out.File
	file_encounter_proto_rawDesc = nil
	file_encounter_proto_goTypes = nil
	file_encounter_proto_depIdxs = nil
}
