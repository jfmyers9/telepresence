// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: rpc/connector/connector.proto

package connector

import (
	manager "github.com/datawire/telepresence2/pkg/rpc/manager"
	version "github.com/datawire/telepresence2/pkg/rpc/version"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// InterceptError is a common error type used by the intercept call family (add,
// remove, list, available).
type InterceptError int32

const (
	InterceptError_UNSPECIFIED                InterceptError = 0
	InterceptError_NO_PREVIEW_HOST            InterceptError = 1
	InterceptError_NO_CONNECTION              InterceptError = 2
	InterceptError_NO_TRAFFIC_MANAGER         InterceptError = 3
	InterceptError_TRAFFIC_MANAGER_CONNECTING InterceptError = 4
	InterceptError_TRAFFIC_MANAGER_ERROR      InterceptError = 5
	InterceptError_ALREADY_EXISTS             InterceptError = 6
	InterceptError_NO_ACCEPTABLE_DEPLOYMENT   InterceptError = 7
	InterceptError_AMBIGUOUS_MATCH            InterceptError = 8
	InterceptError_FAILED_TO_ESTABLISH        InterceptError = 9
	InterceptError_FAILED_TO_REMOVE           InterceptError = 10
	InterceptError_NOT_FOUND                  InterceptError = 11
)

// Enum value maps for InterceptError.
var (
	InterceptError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "NO_PREVIEW_HOST",
		2:  "NO_CONNECTION",
		3:  "NO_TRAFFIC_MANAGER",
		4:  "TRAFFIC_MANAGER_CONNECTING",
		5:  "TRAFFIC_MANAGER_ERROR",
		6:  "ALREADY_EXISTS",
		7:  "NO_ACCEPTABLE_DEPLOYMENT",
		8:  "AMBIGUOUS_MATCH",
		9:  "FAILED_TO_ESTABLISH",
		10: "FAILED_TO_REMOVE",
		11: "NOT_FOUND",
	}
	InterceptError_value = map[string]int32{
		"UNSPECIFIED":                0,
		"NO_PREVIEW_HOST":            1,
		"NO_CONNECTION":              2,
		"NO_TRAFFIC_MANAGER":         3,
		"TRAFFIC_MANAGER_CONNECTING": 4,
		"TRAFFIC_MANAGER_ERROR":      5,
		"ALREADY_EXISTS":             6,
		"NO_ACCEPTABLE_DEPLOYMENT":   7,
		"AMBIGUOUS_MATCH":            8,
		"FAILED_TO_ESTABLISH":        9,
		"FAILED_TO_REMOVE":           10,
		"NOT_FOUND":                  11,
	}
)

func (x InterceptError) Enum() *InterceptError {
	p := new(InterceptError)
	*p = x
	return p
}

func (x InterceptError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterceptError) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_connector_connector_proto_enumTypes[0].Descriptor()
}

func (InterceptError) Type() protoreflect.EnumType {
	return &file_rpc_connector_connector_proto_enumTypes[0]
}

func (x InterceptError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InterceptError.Descriptor instead.
func (InterceptError) EnumDescriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{0}
}

type ConnectInfo_ErrType int32

const (
	ConnectInfo_UNSPECIFIED            ConnectInfo_ErrType = 0
	ConnectInfo_ALREADY_CONNECTED      ConnectInfo_ErrType = 1
	ConnectInfo_DISCONNECTING          ConnectInfo_ErrType = 2
	ConnectInfo_CLUSTER_FAILED         ConnectInfo_ErrType = 3
	ConnectInfo_BRIDGE_FAILED          ConnectInfo_ErrType = 4
	ConnectInfo_TRAFFIC_MANAGER_FAILED ConnectInfo_ErrType = 5
)

// Enum value maps for ConnectInfo_ErrType.
var (
	ConnectInfo_ErrType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ALREADY_CONNECTED",
		2: "DISCONNECTING",
		3: "CLUSTER_FAILED",
		4: "BRIDGE_FAILED",
		5: "TRAFFIC_MANAGER_FAILED",
	}
	ConnectInfo_ErrType_value = map[string]int32{
		"UNSPECIFIED":            0,
		"ALREADY_CONNECTED":      1,
		"DISCONNECTING":          2,
		"CLUSTER_FAILED":         3,
		"BRIDGE_FAILED":          4,
		"TRAFFIC_MANAGER_FAILED": 5,
	}
)

func (x ConnectInfo_ErrType) Enum() *ConnectInfo_ErrType {
	p := new(ConnectInfo_ErrType)
	*p = x
	return p
}

func (x ConnectInfo_ErrType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectInfo_ErrType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_connector_connector_proto_enumTypes[1].Descriptor()
}

func (ConnectInfo_ErrType) Type() protoreflect.EnumType {
	return &file_rpc_connector_connector_proto_enumTypes[1]
}

func (x ConnectInfo_ErrType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectInfo_ErrType.Descriptor instead.
func (ConnectInfo_ErrType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{1, 0}
}

type ConnectorStatus_ErrType int32

const (
	ConnectorStatus_UNSPECIFIED  ConnectorStatus_ErrType = 0
	ConnectorStatus_NOT_STARTED  ConnectorStatus_ErrType = 1
	ConnectorStatus_DISCONNECTED ConnectorStatus_ErrType = 2
)

// Enum value maps for ConnectorStatus_ErrType.
var (
	ConnectorStatus_ErrType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NOT_STARTED",
		2: "DISCONNECTED",
	}
	ConnectorStatus_ErrType_value = map[string]int32{
		"UNSPECIFIED":  0,
		"NOT_STARTED":  1,
		"DISCONNECTED": 2,
	}
)

func (x ConnectorStatus_ErrType) Enum() *ConnectorStatus_ErrType {
	p := new(ConnectorStatus_ErrType)
	*p = x
	return p
}

func (x ConnectorStatus_ErrType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectorStatus_ErrType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_connector_connector_proto_enumTypes[2].Descriptor()
}

func (ConnectorStatus_ErrType) Type() protoreflect.EnumType {
	return &file_rpc_connector_connector_proto_enumTypes[2]
}

func (x ConnectorStatus_ErrType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectorStatus_ErrType.Descriptor instead.
func (ConnectorStatus_ErrType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{2, 0}
}

// ConnectRequest contains the information needed to connect ot a cluster.
type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context   string   `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	InstallId string   `protobuf:"bytes,3,opt,name=install_id,json=installId,proto3" json:"install_id,omitempty"`
	Args      []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	// true if this is part of a CI run
	IsCi bool `protobuf:"varint,5,opt,name=is_ci,json=isCi,proto3" json:"is_ci,omitempty"`
	// true if the call must wait for a running traffic manager before returning
	InterceptEnabled bool `protobuf:"varint,6,opt,name=intercept_enabled,json=interceptEnabled,proto3" json:"intercept_enabled,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_connector_connector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_connector_connector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *ConnectRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ConnectRequest) GetInstallId() string {
	if x != nil {
		return x.InstallId
	}
	return ""
}

func (x *ConnectRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ConnectRequest) GetIsCi() bool {
	if x != nil {
		return x.IsCi
	}
	return false
}

func (x *ConnectRequest) GetInterceptEnabled() bool {
	if x != nil {
		return x.InterceptEnabled
	}
	return false
}

type ConnectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error          ConnectInfo_ErrType `protobuf:"varint,1,opt,name=error,proto3,enum=telepresence.ConnectInfo_ErrType" json:"error,omitempty"`
	ErrorText      string              `protobuf:"bytes,2,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
	ClusterContext string              `protobuf:"bytes,3,opt,name=cluster_context,json=clusterContext,proto3" json:"cluster_context,omitempty"`
	ClusterServer  string              `protobuf:"bytes,4,opt,name=cluster_server,json=clusterServer,proto3" json:"cluster_server,omitempty"`
}

func (x *ConnectInfo) Reset() {
	*x = ConnectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_connector_connector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectInfo) ProtoMessage() {}

func (x *ConnectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_connector_connector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectInfo.ProtoReflect.Descriptor instead.
func (*ConnectInfo) Descriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectInfo) GetError() ConnectInfo_ErrType {
	if x != nil {
		return x.Error
	}
	return ConnectInfo_UNSPECIFIED
}

func (x *ConnectInfo) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

func (x *ConnectInfo) GetClusterContext() string {
	if x != nil {
		return x.ClusterContext
	}
	return ""
}

func (x *ConnectInfo) GetClusterServer() string {
	if x != nil {
		return x.ClusterServer
	}
	return ""
}

type ConnectorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      ConnectorStatus_ErrType        `protobuf:"varint,1,opt,name=error,proto3,enum=telepresence.ConnectorStatus_ErrType" json:"error,omitempty"`
	ErrorText  string                         `protobuf:"bytes,2,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
	Bridge     bool                           `protobuf:"varint,3,opt,name=bridge,proto3" json:"bridge,omitempty"`
	Cluster    *ConnectorStatus_ClusterInfo   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Agents     *manager.AgentInfoSnapshot     `protobuf:"bytes,5,opt,name=agents,proto3" json:"agents,omitempty"`
	Intercepts *manager.InterceptInfoSnapshot `protobuf:"bytes,6,opt,name=intercepts,proto3" json:"intercepts,omitempty"`
}

func (x *ConnectorStatus) Reset() {
	*x = ConnectorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_connector_connector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorStatus) ProtoMessage() {}

func (x *ConnectorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_connector_connector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorStatus.ProtoReflect.Descriptor instead.
func (*ConnectorStatus) Descriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectorStatus) GetError() ConnectorStatus_ErrType {
	if x != nil {
		return x.Error
	}
	return ConnectorStatus_UNSPECIFIED
}

func (x *ConnectorStatus) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

func (x *ConnectorStatus) GetBridge() bool {
	if x != nil {
		return x.Bridge
	}
	return false
}

func (x *ConnectorStatus) GetCluster() *ConnectorStatus_ClusterInfo {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ConnectorStatus) GetAgents() *manager.AgentInfoSnapshot {
	if x != nil {
		return x.Agents
	}
	return nil
}

func (x *ConnectorStatus) GetIntercepts() *manager.InterceptInfoSnapshot {
	if x != nil {
		return x.Intercepts
	}
	return nil
}

type InterceptResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterceptInfo *manager.InterceptInfo `protobuf:"bytes,1,opt,name=intercept_info,json=interceptInfo,proto3" json:"intercept_info,omitempty"`
	Error         InterceptError         `protobuf:"varint,2,opt,name=error,proto3,enum=telepresence.InterceptError" json:"error,omitempty"`
	ErrorText     string                 `protobuf:"bytes,3,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
}

func (x *InterceptResult) Reset() {
	*x = InterceptResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_connector_connector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterceptResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterceptResult) ProtoMessage() {}

func (x *InterceptResult) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_connector_connector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterceptResult.ProtoReflect.Descriptor instead.
func (*InterceptResult) Descriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{3}
}

func (x *InterceptResult) GetInterceptInfo() *manager.InterceptInfo {
	if x != nil {
		return x.InterceptInfo
	}
	return nil
}

func (x *InterceptResult) GetError() InterceptError {
	if x != nil {
		return x.Error
	}
	return InterceptError_UNSPECIFIED
}

func (x *InterceptResult) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

type ConnectorStatus_ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connected bool   `protobuf:"varint,1,opt,name=connected,proto3" json:"connected,omitempty"`
	Server    string `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Context   string `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ConnectorStatus_ClusterInfo) Reset() {
	*x = ConnectorStatus_ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_connector_connector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectorStatus_ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectorStatus_ClusterInfo) ProtoMessage() {}

func (x *ConnectorStatus_ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_connector_connector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectorStatus_ClusterInfo.ProtoReflect.Descriptor instead.
func (*ConnectorStatus_ClusterInfo) Descriptor() ([]byte, []int) {
	return file_rpc_connector_connector_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ConnectorStatus_ClusterInfo) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *ConnectorStatus_ClusterInfo) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ConnectorStatus_ClusterInfo) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

var File_rpc_connector_connector_proto protoreflect.FileDescriptor

var file_rpc_connector_connector_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbd, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x13,
	0x0a, 0x05, 0x69, 0x73, 0x5f, 0x63, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x43, 0x69, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0xbf, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x37, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x45, 0x72, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49,
	0x43, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x05, 0x22, 0xdc, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x06, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65,
	0x70, 0x74, 0x73, 0x1a, 0x5d, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x3d, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65,
	0x70, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x2a, 0xa1, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4e,
	0x4f, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x4f, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49,
	0x43, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x54,
	0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x4f,
	0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x4d, 0x42, 0x49,
	0x47, 0x55, 0x4f, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x08, 0x12, 0x17, 0x0a,
	0x13, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x53, 0x54, 0x41, 0x42,
	0x4c, 0x49, 0x53, 0x48, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x32, 0xbd, 0x04, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1c,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x52, 0x0a, 0x0f, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x12, 0x20, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x1a,
	0x1d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x49,
	0x0a, 0x13, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x63, 0x65, 0x70, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x48, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69,
	0x72, 0x65, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x32,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_connector_connector_proto_rawDescOnce sync.Once
	file_rpc_connector_connector_proto_rawDescData = file_rpc_connector_connector_proto_rawDesc
)

func file_rpc_connector_connector_proto_rawDescGZIP() []byte {
	file_rpc_connector_connector_proto_rawDescOnce.Do(func() {
		file_rpc_connector_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_connector_connector_proto_rawDescData)
	})
	return file_rpc_connector_connector_proto_rawDescData
}

var file_rpc_connector_connector_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_rpc_connector_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_connector_connector_proto_goTypes = []interface{}{
	(InterceptError)(0),                     // 0: telepresence.InterceptError
	(ConnectInfo_ErrType)(0),                // 1: telepresence.ConnectInfo.ErrType
	(ConnectorStatus_ErrType)(0),            // 2: telepresence.ConnectorStatus.ErrType
	(*ConnectRequest)(nil),                  // 3: telepresence.ConnectRequest
	(*ConnectInfo)(nil),                     // 4: telepresence.ConnectInfo
	(*ConnectorStatus)(nil),                 // 5: telepresence.ConnectorStatus
	(*InterceptResult)(nil),                 // 6: telepresence.InterceptResult
	(*ConnectorStatus_ClusterInfo)(nil),     // 7: telepresence.ConnectorStatus.ClusterInfo
	(*manager.AgentInfoSnapshot)(nil),       // 8: manager.AgentInfoSnapshot
	(*manager.InterceptInfoSnapshot)(nil),   // 9: manager.InterceptInfoSnapshot
	(*manager.InterceptInfo)(nil),           // 10: manager.InterceptInfo
	(*empty.Empty)(nil),                     // 11: google.protobuf.Empty
	(*manager.CreateInterceptRequest)(nil),  // 12: manager.CreateInterceptRequest
	(*manager.RemoveInterceptRequest2)(nil), // 13: manager.RemoveInterceptRequest2
	(*version.VersionInfo)(nil),             // 14: version.VersionInfo
}
var file_rpc_connector_connector_proto_depIdxs = []int32{
	1,  // 0: telepresence.ConnectInfo.error:type_name -> telepresence.ConnectInfo.ErrType
	2,  // 1: telepresence.ConnectorStatus.error:type_name -> telepresence.ConnectorStatus.ErrType
	7,  // 2: telepresence.ConnectorStatus.cluster:type_name -> telepresence.ConnectorStatus.ClusterInfo
	8,  // 3: telepresence.ConnectorStatus.agents:type_name -> manager.AgentInfoSnapshot
	9,  // 4: telepresence.ConnectorStatus.intercepts:type_name -> manager.InterceptInfoSnapshot
	10, // 5: telepresence.InterceptResult.intercept_info:type_name -> manager.InterceptInfo
	0,  // 6: telepresence.InterceptResult.error:type_name -> telepresence.InterceptError
	11, // 7: telepresence.Connector.Version:input_type -> google.protobuf.Empty
	11, // 8: telepresence.Connector.Status:input_type -> google.protobuf.Empty
	3,  // 9: telepresence.Connector.Connect:input_type -> telepresence.ConnectRequest
	12, // 10: telepresence.Connector.CreateIntercept:input_type -> manager.CreateInterceptRequest
	13, // 11: telepresence.Connector.RemoveIntercept:input_type -> manager.RemoveInterceptRequest2
	11, // 12: telepresence.Connector.AvailableIntercepts:input_type -> google.protobuf.Empty
	11, // 13: telepresence.Connector.ListIntercepts:input_type -> google.protobuf.Empty
	11, // 14: telepresence.Connector.Quit:input_type -> google.protobuf.Empty
	14, // 15: telepresence.Connector.Version:output_type -> version.VersionInfo
	5,  // 16: telepresence.Connector.Status:output_type -> telepresence.ConnectorStatus
	4,  // 17: telepresence.Connector.Connect:output_type -> telepresence.ConnectInfo
	6,  // 18: telepresence.Connector.CreateIntercept:output_type -> telepresence.InterceptResult
	6,  // 19: telepresence.Connector.RemoveIntercept:output_type -> telepresence.InterceptResult
	8,  // 20: telepresence.Connector.AvailableIntercepts:output_type -> manager.AgentInfoSnapshot
	9,  // 21: telepresence.Connector.ListIntercepts:output_type -> manager.InterceptInfoSnapshot
	11, // 22: telepresence.Connector.Quit:output_type -> google.protobuf.Empty
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_rpc_connector_connector_proto_init() }
func file_rpc_connector_connector_proto_init() {
	if File_rpc_connector_connector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_connector_connector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_rpc_connector_connector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectInfo); i {
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
		file_rpc_connector_connector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorStatus); i {
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
		file_rpc_connector_connector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterceptResult); i {
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
		file_rpc_connector_connector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectorStatus_ClusterInfo); i {
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
			RawDescriptor: file_rpc_connector_connector_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_connector_connector_proto_goTypes,
		DependencyIndexes: file_rpc_connector_connector_proto_depIdxs,
		EnumInfos:         file_rpc_connector_connector_proto_enumTypes,
		MessageInfos:      file_rpc_connector_connector_proto_msgTypes,
	}.Build()
	File_rpc_connector_connector_proto = out.File
	file_rpc_connector_connector_proto_rawDesc = nil
	file_rpc_connector_connector_proto_goTypes = nil
	file_rpc_connector_connector_proto_depIdxs = nil
}
