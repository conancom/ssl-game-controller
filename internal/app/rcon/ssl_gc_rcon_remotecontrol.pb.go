// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: ssl_gc_rcon_remotecontrol.proto

package rcon

import (
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
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

type RemoteControlToController_Request int32

const (
	RemoteControlToController_UNKNOWN RemoteControlToController_Request = 0
	// Ping the GC to test the connection. The GC will respond with OK.
	RemoteControlToController_PING RemoteControlToController_Request = 1
	// Request the current team state
	RemoteControlToController_GET_STATE RemoteControlToController_Request = 2
	// Raise a challenge flag (this is not revocable)
	RemoteControlToController_CHALLENGE_FLAG RemoteControlToController_Request = 3
)

// Enum value maps for RemoteControlToController_Request.
var (
	RemoteControlToController_Request_name = map[int32]string{
		0: "UNKNOWN",
		1: "PING",
		2: "GET_STATE",
		3: "CHALLENGE_FLAG",
	}
	RemoteControlToController_Request_value = map[string]int32{
		"UNKNOWN":        0,
		"PING":           1,
		"GET_STATE":      2,
		"CHALLENGE_FLAG": 3,
	}
)

func (x RemoteControlToController_Request) Enum() *RemoteControlToController_Request {
	p := new(RemoteControlToController_Request)
	*p = x
	return p
}

func (x RemoteControlToController_Request) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteControlToController_Request) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_rcon_remotecontrol_proto_enumTypes[0].Descriptor()
}

func (RemoteControlToController_Request) Type() protoreflect.EnumType {
	return &file_ssl_gc_rcon_remotecontrol_proto_enumTypes[0]
}

func (x RemoteControlToController_Request) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RemoteControlToController_Request) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RemoteControlToController_Request(num)
	return nil
}

// Deprecated: Use RemoteControlToController_Request.Descriptor instead.
func (RemoteControlToController_Request) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{1, 0}
}

// a registration that must be send by the remote control to the controller as the very first message
type RemoteControlRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the team to be controlled
	Team *state.Team `protobuf:"varint,1,req,name=team,enum=Team" json:"team,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (x *RemoteControlRegistration) Reset() {
	*x = RemoteControlRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteControlRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteControlRegistration) ProtoMessage() {}

func (x *RemoteControlRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteControlRegistration.ProtoReflect.Descriptor instead.
func (*RemoteControlRegistration) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteControlRegistration) GetTeam() state.Team {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return state.Team_UNKNOWN
}

func (x *RemoteControlRegistration) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// wrapper for all messages from the remote control to the controller
type RemoteControlToController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Types that are assignable to Msg:
	//	*RemoteControlToController_Request_
	//	*RemoteControlToController_DesiredKeeper
	//	*RemoteControlToController_SubstituteBot
	//	*RemoteControlToController_Timeout
	//	*RemoteControlToController_EmergencyStop
	Msg isRemoteControlToController_Msg `protobuf_oneof:"msg"`
}

func (x *RemoteControlToController) Reset() {
	*x = RemoteControlToController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteControlToController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteControlToController) ProtoMessage() {}

func (x *RemoteControlToController) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteControlToController.ProtoReflect.Descriptor instead.
func (*RemoteControlToController) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteControlToController) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *RemoteControlToController) GetMsg() isRemoteControlToController_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *RemoteControlToController) GetRequest() RemoteControlToController_Request {
	if x, ok := x.GetMsg().(*RemoteControlToController_Request_); ok {
		return x.Request
	}
	return RemoteControlToController_UNKNOWN
}

func (x *RemoteControlToController) GetDesiredKeeper() int32 {
	if x, ok := x.GetMsg().(*RemoteControlToController_DesiredKeeper); ok {
		return x.DesiredKeeper
	}
	return 0
}

func (x *RemoteControlToController) GetSubstituteBot() bool {
	if x, ok := x.GetMsg().(*RemoteControlToController_SubstituteBot); ok {
		return x.SubstituteBot
	}
	return false
}

func (x *RemoteControlToController) GetTimeout() bool {
	if x, ok := x.GetMsg().(*RemoteControlToController_Timeout); ok {
		return x.Timeout
	}
	return false
}

func (x *RemoteControlToController) GetEmergencyStop() bool {
	if x, ok := x.GetMsg().(*RemoteControlToController_EmergencyStop); ok {
		return x.EmergencyStop
	}
	return false
}

type isRemoteControlToController_Msg interface {
	isRemoteControlToController_Msg()
}

type RemoteControlToController_Request_ struct {
	// send a ping to the GC to test if the connection is still open.
	// the value is ignored and a reply is sent back
	Request RemoteControlToController_Request `protobuf:"varint,2,opt,name=request,enum=RemoteControlToController_Request,oneof"`
}

type RemoteControlToController_DesiredKeeper struct {
	// request a new desired keeper id
	DesiredKeeper int32 `protobuf:"varint,3,opt,name=desired_keeper,json=desiredKeeper,oneof"`
}

type RemoteControlToController_SubstituteBot struct {
	// true: request to substitute a robot at the next possibility
	// false: cancel request
	SubstituteBot bool `protobuf:"varint,4,opt,name=substitute_bot,json=substituteBot,oneof"`
}

type RemoteControlToController_Timeout struct {
	// true: request a timeout with the next stoppage
	// false: cancel the reuqest
	Timeout bool `protobuf:"varint,5,opt,name=timeout,oneof"`
}

type RemoteControlToController_EmergencyStop struct {
	// true: initiate an emergency stop
	// false: ignored (emergency_stop can not be canceled)
	EmergencyStop bool `protobuf:"varint,6,opt,name=emergency_stop,json=emergencyStop,oneof"`
}

func (*RemoteControlToController_Request_) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_DesiredKeeper) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_SubstituteBot) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_Timeout) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_EmergencyStop) isRemoteControlToController_Msg() {}

// wrapper for all messages from controller to a team's computer
type ControllerToRemoteControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply" json:"controller_reply,omitempty"`
	// currently set keeper id
	Keeper *int32 `protobuf:"varint,2,opt,name=keeper" json:"keeper,omitempty"`
	// true, if substitution request pending
	SubstituteBot *bool `protobuf:"varint,3,opt,name=substitute_bot,json=substituteBot" json:"substitute_bot,omitempty"`
	// true, if emergency stop pending
	EmergencyStop *bool `protobuf:"varint,4,opt,name=emergency_stop,json=emergencyStop" json:"emergency_stop,omitempty"`
	// number of seconds till emergency stop is executed
	EmergencyStopIn *float32 `protobuf:"fixed32,5,opt,name=emergency_stop_in,json=emergencyStopIn" json:"emergency_stop_in,omitempty"`
	// true, if timeout request pending
	Timeout *bool `protobuf:"varint,6,opt,name=timeout" json:"timeout,omitempty"`
	// true, if challenge flag pending
	ChallengeFlag *bool `protobuf:"varint,7,opt,name=challenge_flag,json=challengeFlag" json:"challenge_flag,omitempty"`
}

func (x *ControllerToRemoteControl) Reset() {
	*x = ControllerToRemoteControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerToRemoteControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerToRemoteControl) ProtoMessage() {}

func (x *ControllerToRemoteControl) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_remotecontrol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerToRemoteControl.ProtoReflect.Descriptor instead.
func (*ControllerToRemoteControl) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP(), []int{2}
}

func (x *ControllerToRemoteControl) GetControllerReply() *ControllerReply {
	if x != nil {
		return x.ControllerReply
	}
	return nil
}

func (x *ControllerToRemoteControl) GetKeeper() int32 {
	if x != nil && x.Keeper != nil {
		return *x.Keeper
	}
	return 0
}

func (x *ControllerToRemoteControl) GetSubstituteBot() bool {
	if x != nil && x.SubstituteBot != nil {
		return *x.SubstituteBot
	}
	return false
}

func (x *ControllerToRemoteControl) GetEmergencyStop() bool {
	if x != nil && x.EmergencyStop != nil {
		return *x.EmergencyStop
	}
	return false
}

func (x *ControllerToRemoteControl) GetEmergencyStopIn() float32 {
	if x != nil && x.EmergencyStopIn != nil {
		return *x.EmergencyStopIn
	}
	return 0
}

func (x *ControllerToRemoteControl) GetTimeout() bool {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return false
}

func (x *ControllerToRemoteControl) GetChallengeFlag() bool {
	if x != nil && x.ChallengeFlag != nil {
		return *x.ChallengeFlag
	}
	return false
}

var File_ssl_gc_rcon_remotecontrol_proto protoreflect.FileDescriptor

var file_ssl_gc_rcon_remotecontrol_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72,
	0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x19, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x19,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x6f, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x27, 0x0a, 0x0e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73,
	0x74, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x22, 0x43, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x10, 0x03, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xab, 0x02, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x3b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x74,
	0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0f, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x6f,
	0x70, 0x49, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73,
	0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_ssl_gc_rcon_remotecontrol_proto_rawDescOnce sync.Once
	file_ssl_gc_rcon_remotecontrol_proto_rawDescData = file_ssl_gc_rcon_remotecontrol_proto_rawDesc
)

func file_ssl_gc_rcon_remotecontrol_proto_rawDescGZIP() []byte {
	file_ssl_gc_rcon_remotecontrol_proto_rawDescOnce.Do(func() {
		file_ssl_gc_rcon_remotecontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_rcon_remotecontrol_proto_rawDescData)
	})
	return file_ssl_gc_rcon_remotecontrol_proto_rawDescData
}

var file_ssl_gc_rcon_remotecontrol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ssl_gc_rcon_remotecontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ssl_gc_rcon_remotecontrol_proto_goTypes = []interface{}{
	(RemoteControlToController_Request)(0), // 0: RemoteControlToController.Request
	(*RemoteControlRegistration)(nil),      // 1: RemoteControlRegistration
	(*RemoteControlToController)(nil),      // 2: RemoteControlToController
	(*ControllerToRemoteControl)(nil),      // 3: ControllerToRemoteControl
	(state.Team)(0),                        // 4: Team
	(*Signature)(nil),                      // 5: Signature
	(*ControllerReply)(nil),                // 6: ControllerReply
}
var file_ssl_gc_rcon_remotecontrol_proto_depIdxs = []int32{
	4, // 0: RemoteControlRegistration.team:type_name -> Team
	5, // 1: RemoteControlRegistration.signature:type_name -> Signature
	5, // 2: RemoteControlToController.signature:type_name -> Signature
	0, // 3: RemoteControlToController.request:type_name -> RemoteControlToController.Request
	6, // 4: ControllerToRemoteControl.controller_reply:type_name -> ControllerReply
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ssl_gc_rcon_remotecontrol_proto_init() }
func file_ssl_gc_rcon_remotecontrol_proto_init() {
	if File_ssl_gc_rcon_remotecontrol_proto != nil {
		return
	}
	file_ssl_gc_rcon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_rcon_remotecontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteControlRegistration); i {
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
		file_ssl_gc_rcon_remotecontrol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteControlToController); i {
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
		file_ssl_gc_rcon_remotecontrol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerToRemoteControl); i {
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
	file_ssl_gc_rcon_remotecontrol_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RemoteControlToController_Request_)(nil),
		(*RemoteControlToController_DesiredKeeper)(nil),
		(*RemoteControlToController_SubstituteBot)(nil),
		(*RemoteControlToController_Timeout)(nil),
		(*RemoteControlToController_EmergencyStop)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_gc_rcon_remotecontrol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_rcon_remotecontrol_proto_goTypes,
		DependencyIndexes: file_ssl_gc_rcon_remotecontrol_proto_depIdxs,
		EnumInfos:         file_ssl_gc_rcon_remotecontrol_proto_enumTypes,
		MessageInfos:      file_ssl_gc_rcon_remotecontrol_proto_msgTypes,
	}.Build()
	File_ssl_gc_rcon_remotecontrol_proto = out.File
	file_ssl_gc_rcon_remotecontrol_proto_rawDesc = nil
	file_ssl_gc_rcon_remotecontrol_proto_goTypes = nil
	file_ssl_gc_rcon_remotecontrol_proto_depIdxs = nil
}
