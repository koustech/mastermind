// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: proto/mastermind/v1/mastermind.proto

package mastermindv1

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

// A MissionState is a state the mission can be in at any point in time
type MissionState int32

const (
	MissionState_MISSION_STATE_UNSPECIFIED MissionState = 0
	MissionState_MISSION_STATE_APPROACH    MissionState = 1
	MissionState_MISSION_STATE_FOLLOWING   MissionState = 2
	MissionState_MISSION_STATE_KAMIKAZE    MissionState = 3
)

// Enum value maps for MissionState.
var (
	MissionState_name = map[int32]string{
		0: "MISSION_STATE_UNSPECIFIED",
		1: "MISSION_STATE_APPROACH",
		2: "MISSION_STATE_FOLLOWING",
		3: "MISSION_STATE_KAMIKAZE",
	}
	MissionState_value = map[string]int32{
		"MISSION_STATE_UNSPECIFIED": 0,
		"MISSION_STATE_APPROACH":    1,
		"MISSION_STATE_FOLLOWING":   2,
		"MISSION_STATE_KAMIKAZE":    3,
	}
)

func (x MissionState) Enum() *MissionState {
	p := new(MissionState)
	*p = x
	return p
}

func (x MissionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MissionState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_mastermind_v1_mastermind_proto_enumTypes[0].Descriptor()
}

func (MissionState) Type() protoreflect.EnumType {
	return &file_proto_mastermind_v1_mastermind_proto_enumTypes[0]
}

func (x MissionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MissionState.Descriptor instead.
func (MissionState) EnumDescriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{0}
}

// A StateTransition is an action that causes the current MissonState to change
//
// e.g. STATE_TRANSITION_TARGET_APPROACHED changes mission state
// from MISSION_STATE_APPROACH to MISSION_STATE_FOLLOWING
type StateTransition int32

const (
	StateTransition_STATE_TRANSITION_UNSPECIFIED            StateTransition = 0
	StateTransition_STATE_TRANSITION_TARGET_APPROACHED      StateTransition = 1
	StateTransition_STATE_TRANSITION_LOCK_FAILED            StateTransition = 2
	StateTransition_STATE_TRANSITION_LOCK_SUCCESS           StateTransition = 3
	StateTransition_STATE_TRANSITION_QR_SUCCESS             StateTransition = 4
	StateTransition_STATE_TRANSITION_QR_FAILED              StateTransition = 5
	StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED StateTransition = 6
	StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED StateTransition = 7
)

// Enum value maps for StateTransition.
var (
	StateTransition_name = map[int32]string{
		0: "STATE_TRANSITION_UNSPECIFIED",
		1: "STATE_TRANSITION_TARGET_APPROACHED",
		2: "STATE_TRANSITION_LOCK_FAILED",
		3: "STATE_TRANSITION_LOCK_SUCCESS",
		4: "STATE_TRANSITION_QR_SUCCESS",
		5: "STATE_TRANSITION_QR_FAILED",
		6: "STATE_TRANSITION_MODE_KAMIKAZE_SELECTED",
		7: "STATE_TRANSITION_MODE_APPROACH_SELECTED",
	}
	StateTransition_value = map[string]int32{
		"STATE_TRANSITION_UNSPECIFIED":            0,
		"STATE_TRANSITION_TARGET_APPROACHED":      1,
		"STATE_TRANSITION_LOCK_FAILED":            2,
		"STATE_TRANSITION_LOCK_SUCCESS":           3,
		"STATE_TRANSITION_QR_SUCCESS":             4,
		"STATE_TRANSITION_QR_FAILED":              5,
		"STATE_TRANSITION_MODE_KAMIKAZE_SELECTED": 6,
		"STATE_TRANSITION_MODE_APPROACH_SELECTED": 7,
	}
)

func (x StateTransition) Enum() *StateTransition {
	p := new(StateTransition)
	*p = x
	return p
}

func (x StateTransition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateTransition) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_mastermind_v1_mastermind_proto_enumTypes[1].Descriptor()
}

func (StateTransition) Type() protoreflect.EnumType {
	return &file_proto_mastermind_v1_mastermind_proto_enumTypes[1]
}

func (x StateTransition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateTransition.Descriptor instead.
func (StateTransition) EnumDescriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{1}
}

type AgilityLevel int32

const (
	AgilityLevel_AGILITY_LEVEL_UNSPECIFIED AgilityLevel = 0
	AgilityLevel_AGILITY_LEVEL_HIGH        AgilityLevel = 1
	AgilityLevel_AGILITY_LEVEL_MEDIUM      AgilityLevel = 2
	AgilityLevel_AGILITY_LEVEL_LOW         AgilityLevel = 3
)

// Enum value maps for AgilityLevel.
var (
	AgilityLevel_name = map[int32]string{
		0: "AGILITY_LEVEL_UNSPECIFIED",
		1: "AGILITY_LEVEL_HIGH",
		2: "AGILITY_LEVEL_MEDIUM",
		3: "AGILITY_LEVEL_LOW",
	}
	AgilityLevel_value = map[string]int32{
		"AGILITY_LEVEL_UNSPECIFIED": 0,
		"AGILITY_LEVEL_HIGH":        1,
		"AGILITY_LEVEL_MEDIUM":      2,
		"AGILITY_LEVEL_LOW":         3,
	}
)

func (x AgilityLevel) Enum() *AgilityLevel {
	p := new(AgilityLevel)
	*p = x
	return p
}

func (x AgilityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgilityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_mastermind_v1_mastermind_proto_enumTypes[2].Descriptor()
}

func (AgilityLevel) Type() protoreflect.EnumType {
	return &file_proto_mastermind_v1_mastermind_proto_enumTypes[2]
}

func (x AgilityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgilityLevel.Descriptor instead.
func (AgilityLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{2}
}

type UpdateStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTransition StateTransition `protobuf:"varint,1,opt,name=state_transition,json=stateTransition,proto3,enum=mastermind.v1.StateTransition" json:"state_transition,omitempty"`
}

func (x *UpdateStateRequest) Reset() {
	*x = UpdateStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateRequest) ProtoMessage() {}

func (x *UpdateStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateStateRequest) GetStateTransition() StateTransition {
	if x != nil {
		return x.StateTransition
	}
	return StateTransition_STATE_TRANSITION_UNSPECIFIED
}

type UpdateStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldState        MissionState    `protobuf:"varint,1,opt,name=old_state,json=oldState,proto3,enum=mastermind.v1.MissionState" json:"old_state,omitempty"`
	StateTransition StateTransition `protobuf:"varint,2,opt,name=state_transition,json=stateTransition,proto3,enum=mastermind.v1.StateTransition" json:"state_transition,omitempty"`
	CurrentState    MissionState    `protobuf:"varint,3,opt,name=current_state,json=currentState,proto3,enum=mastermind.v1.MissionState" json:"current_state,omitempty"`
}

func (x *UpdateStateResponse) Reset() {
	*x = UpdateStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateResponse) ProtoMessage() {}

func (x *UpdateStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStateResponse) GetOldState() MissionState {
	if x != nil {
		return x.OldState
	}
	return MissionState_MISSION_STATE_UNSPECIFIED
}

func (x *UpdateStateResponse) GetStateTransition() StateTransition {
	if x != nil {
		return x.StateTransition
	}
	return StateTransition_STATE_TRANSITION_UNSPECIFIED
}

func (x *UpdateStateResponse) GetCurrentState() MissionState {
	if x != nil {
		return x.CurrentState
	}
	return MissionState_MISSION_STATE_UNSPECIFIED
}

// Plane is a UAV participating in the round. Can be the sender's own UAV
type Plane struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaneId int32 `protobuf:"varint,1,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"` // Plane's ID in competition server. Exception is sender's
	// UAV, which is changed to a negative int
	Lat          float64      `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`                                                                      // Plane's latitude in degrees
	Lon          float64      `protobuf:"fixed64,3,opt,name=lon,proto3" json:"lon,omitempty"`                                                                      // Plane's longitude in degrees
	Alt          float64      `protobuf:"fixed64,4,opt,name=alt,proto3" json:"alt,omitempty"`                                                                      // Plane's altitude in metres
	Heading      float64      `protobuf:"fixed64,5,opt,name=heading,proto3" json:"heading,omitempty"`                                                              // Plane's heading in degrees from North
	Speed        float64      `protobuf:"fixed64,6,opt,name=speed,proto3" json:"speed,omitempty"`                                                                  // Plane's groundspeend in m/s
	AgilityLevel AgilityLevel `protobuf:"varint,7,opt,name=agility_level,json=agilityLevel,proto3,enum=mastermind.v1.AgilityLevel" json:"agility_level,omitempty"` // Plane's agility level. High or 1 is most
	// agile, low or 3 is least agile
	Score float64 `protobuf:"fixed64,8,opt,name=score,proto3" json:"score,omitempty"` // Probability of locking on plane, between 0 and 100. 100
}

func (x *Plane) Reset() {
	*x = Plane{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plane) ProtoMessage() {}

func (x *Plane) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plane.ProtoReflect.Descriptor instead.
func (*Plane) Descriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{2}
}

func (x *Plane) GetPlaneId() int32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *Plane) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Plane) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Plane) GetAlt() float64 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *Plane) GetHeading() float64 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *Plane) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Plane) GetAgilityLevel() AgilityLevel {
	if x != nil {
		return x.AgilityLevel
	}
	return AgilityLevel_AGILITY_LEVEL_UNSPECIFIED
}

func (x *Plane) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

// GetPlanesRequest is the request that initializes a stream of
// GetPlanesResponses
type GetPlanesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlanesRequest) Reset() {
	*x = GetPlanesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanesRequest) ProtoMessage() {}

func (x *GetPlanesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanesRequest.ProtoReflect.Descriptor instead.
func (*GetPlanesRequest) Descriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{3}
}

// GetPlanesResponse has a list of planes and the selected plane id for any
// point in time
type GetPlanesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelectedPlaneId int32    `protobuf:"varint,1,opt,name=selected_plane_id,json=selectedPlaneId,proto3" json:"selected_plane_id,omitempty"`
	Planes          []*Plane `protobuf:"bytes,2,rep,name=planes,proto3" json:"planes,omitempty"`
}

func (x *GetPlanesResponse) Reset() {
	*x = GetPlanesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanesResponse) ProtoMessage() {}

func (x *GetPlanesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mastermind_v1_mastermind_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanesResponse.ProtoReflect.Descriptor instead.
func (*GetPlanesResponse) Descriptor() ([]byte, []int) {
	return file_proto_mastermind_v1_mastermind_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlanesResponse) GetSelectedPlaneId() int32 {
	if x != nil {
		return x.SelectedPlaneId
	}
	return 0
}

func (x *GetPlanesResponse) GetPlanes() []*Plane {
	if x != nil {
		return x.Planes
	}
	return nil
}

var File_proto_mastermind_v1_mastermind_proto protoreflect.FileDescriptor

var file_proto_mastermind_v1_mastermind_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x5f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x10, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x40, 0x0a, 0x0d, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2a, 0x82, 0x01, 0x0a, 0x0c,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x41, 0x43, 0x48, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4b, 0x41, 0x4d, 0x49, 0x4b, 0x41, 0x5a, 0x45, 0x10, 0x03,
	0x2a, 0xbb, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x52, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4b, 0x41,
	0x4d, 0x49, 0x4b, 0x41, 0x5a, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f,
	0x41, 0x43, 0x48, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x07, 0x2a, 0x76,
	0x0a, 0x0c, 0x41, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d,
	0x0a, 0x19, 0x41, 0x47, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x47, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x47, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x41, 0x47, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x03, 0x32, 0x6d, 0x0a, 0x11, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x32, 0x63, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x61, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xc7, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6f, 0x75, 0x73, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mastermind_v1_mastermind_proto_rawDescOnce sync.Once
	file_proto_mastermind_v1_mastermind_proto_rawDescData = file_proto_mastermind_v1_mastermind_proto_rawDesc
)

func file_proto_mastermind_v1_mastermind_proto_rawDescGZIP() []byte {
	file_proto_mastermind_v1_mastermind_proto_rawDescOnce.Do(func() {
		file_proto_mastermind_v1_mastermind_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mastermind_v1_mastermind_proto_rawDescData)
	})
	return file_proto_mastermind_v1_mastermind_proto_rawDescData
}

var file_proto_mastermind_v1_mastermind_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_mastermind_v1_mastermind_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_mastermind_v1_mastermind_proto_goTypes = []interface{}{
	(MissionState)(0),           // 0: mastermind.v1.MissionState
	(StateTransition)(0),        // 1: mastermind.v1.StateTransition
	(AgilityLevel)(0),           // 2: mastermind.v1.AgilityLevel
	(*UpdateStateRequest)(nil),  // 3: mastermind.v1.UpdateStateRequest
	(*UpdateStateResponse)(nil), // 4: mastermind.v1.UpdateStateResponse
	(*Plane)(nil),               // 5: mastermind.v1.Plane
	(*GetPlanesRequest)(nil),    // 6: mastermind.v1.GetPlanesRequest
	(*GetPlanesResponse)(nil),   // 7: mastermind.v1.GetPlanesResponse
}
var file_proto_mastermind_v1_mastermind_proto_depIdxs = []int32{
	1, // 0: mastermind.v1.UpdateStateRequest.state_transition:type_name -> mastermind.v1.StateTransition
	0, // 1: mastermind.v1.UpdateStateResponse.old_state:type_name -> mastermind.v1.MissionState
	1, // 2: mastermind.v1.UpdateStateResponse.state_transition:type_name -> mastermind.v1.StateTransition
	0, // 3: mastermind.v1.UpdateStateResponse.current_state:type_name -> mastermind.v1.MissionState
	2, // 4: mastermind.v1.Plane.agility_level:type_name -> mastermind.v1.AgilityLevel
	5, // 5: mastermind.v1.GetPlanesResponse.planes:type_name -> mastermind.v1.Plane
	3, // 6: mastermind.v1.MastermindService.UpdateState:input_type -> mastermind.v1.UpdateStateRequest
	6, // 7: mastermind.v1.ApproachService.GetPlanes:input_type -> mastermind.v1.GetPlanesRequest
	4, // 8: mastermind.v1.MastermindService.UpdateState:output_type -> mastermind.v1.UpdateStateResponse
	7, // 9: mastermind.v1.ApproachService.GetPlanes:output_type -> mastermind.v1.GetPlanesResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_mastermind_v1_mastermind_proto_init() }
func file_proto_mastermind_v1_mastermind_proto_init() {
	if File_proto_mastermind_v1_mastermind_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mastermind_v1_mastermind_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStateRequest); i {
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
		file_proto_mastermind_v1_mastermind_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStateResponse); i {
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
		file_proto_mastermind_v1_mastermind_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plane); i {
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
		file_proto_mastermind_v1_mastermind_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanesRequest); i {
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
		file_proto_mastermind_v1_mastermind_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanesResponse); i {
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
			RawDescriptor: file_proto_mastermind_v1_mastermind_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_mastermind_v1_mastermind_proto_goTypes,
		DependencyIndexes: file_proto_mastermind_v1_mastermind_proto_depIdxs,
		EnumInfos:         file_proto_mastermind_v1_mastermind_proto_enumTypes,
		MessageInfos:      file_proto_mastermind_v1_mastermind_proto_msgTypes,
	}.Build()
	File_proto_mastermind_v1_mastermind_proto = out.File
	file_proto_mastermind_v1_mastermind_proto_rawDesc = nil
	file_proto_mastermind_v1_mastermind_proto_goTypes = nil
	file_proto_mastermind_v1_mastermind_proto_depIdxs = nil
}
