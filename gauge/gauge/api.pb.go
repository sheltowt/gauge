// Code generated by protoc-gen-go.
// source: gauge/api.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	gauge/api.proto

It has these top-level messages:
	GetProjectRootRequest
	GetProjectRootResponse
	GetAllStepsRequest
	GetAllStepsResponse
	GetAllSpecsRequest
	GetAllSpecsResponse
	ProtoSpec
	ProtoItem
	ProtoHeading
	ProtoScenario
	ProtoStep
	ProtoStepExecutionStatus
	ProtoConcept
	Fragment
	Parameter
	ProtoComment
	ProtoTableParam
	ProtoTableRow
	APIMessage
	ProtoTags
*/
package main

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import main1 "gauge/messages.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtoItem_ItemType int32

const (
	ProtoItem_Heading  ProtoItem_ItemType = 1
	ProtoItem_Step     ProtoItem_ItemType = 2
	ProtoItem_Concept  ProtoItem_ItemType = 4
	ProtoItem_Scenario ProtoItem_ItemType = 5
	ProtoItem_Comment  ProtoItem_ItemType = 6
	ProtoItem_Table    ProtoItem_ItemType = 7
	ProtoItem_Tags     ProtoItem_ItemType = 8
)

var ProtoItem_ItemType_name = map[int32]string{
	1: "Heading",
	2: "Step",
	4: "Concept",
	5: "Scenario",
	6: "Comment",
	7: "Table",
	8: "Tags",
}
var ProtoItem_ItemType_value = map[string]int32{
	"Heading":  1,
	"Step":     2,
	"Concept":  4,
	"Scenario": 5,
	"Comment":  6,
	"Table":    7,
	"Tags":     8,
}

func (x ProtoItem_ItemType) Enum() *ProtoItem_ItemType {
	p := new(ProtoItem_ItemType)
	*p = x
	return p
}
func (x ProtoItem_ItemType) String() string {
	return proto.EnumName(ProtoItem_ItemType_name, int32(x))
}
func (x *ProtoItem_ItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtoItem_ItemType_value, data, "ProtoItem_ItemType")
	if err != nil {
		return err
	}
	*x = ProtoItem_ItemType(value)
	return nil
}

type ProtoHeading_HeadingType int32

const (
	ProtoHeading_Spec     ProtoHeading_HeadingType = 1
	ProtoHeading_Scenario ProtoHeading_HeadingType = 2
)

var ProtoHeading_HeadingType_name = map[int32]string{
	1: "Spec",
	2: "Scenario",
}
var ProtoHeading_HeadingType_value = map[string]int32{
	"Spec":     1,
	"Scenario": 2,
}

func (x ProtoHeading_HeadingType) Enum() *ProtoHeading_HeadingType {
	p := new(ProtoHeading_HeadingType)
	*p = x
	return p
}
func (x ProtoHeading_HeadingType) String() string {
	return proto.EnumName(ProtoHeading_HeadingType_name, int32(x))
}
func (x *ProtoHeading_HeadingType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtoHeading_HeadingType_value, data, "ProtoHeading_HeadingType")
	if err != nil {
		return err
	}
	*x = ProtoHeading_HeadingType(value)
	return nil
}

type Fragment_FragmentType int32

const (
	Fragment_Text      Fragment_FragmentType = 1
	Fragment_Parameter Fragment_FragmentType = 2
)

var Fragment_FragmentType_name = map[int32]string{
	1: "Text",
	2: "Parameter",
}
var Fragment_FragmentType_value = map[string]int32{
	"Text":      1,
	"Parameter": 2,
}

func (x Fragment_FragmentType) Enum() *Fragment_FragmentType {
	p := new(Fragment_FragmentType)
	*p = x
	return p
}
func (x Fragment_FragmentType) String() string {
	return proto.EnumName(Fragment_FragmentType_name, int32(x))
}
func (x *Fragment_FragmentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Fragment_FragmentType_value, data, "Fragment_FragmentType")
	if err != nil {
		return err
	}
	*x = Fragment_FragmentType(value)
	return nil
}

type Parameter_ParameterType int32

const (
	Parameter_Static  Parameter_ParameterType = 1
	Parameter_Dynamic Parameter_ParameterType = 2
	Parameter_Special Parameter_ParameterType = 3
	Parameter_Table   Parameter_ParameterType = 4
)

var Parameter_ParameterType_name = map[int32]string{
	1: "Static",
	2: "Dynamic",
	3: "Special",
	4: "Table",
}
var Parameter_ParameterType_value = map[string]int32{
	"Static":  1,
	"Dynamic": 2,
	"Special": 3,
	"Table":   4,
}

func (x Parameter_ParameterType) Enum() *Parameter_ParameterType {
	p := new(Parameter_ParameterType)
	*p = x
	return p
}
func (x Parameter_ParameterType) String() string {
	return proto.EnumName(Parameter_ParameterType_name, int32(x))
}
func (x *Parameter_ParameterType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Parameter_ParameterType_value, data, "Parameter_ParameterType")
	if err != nil {
		return err
	}
	*x = Parameter_ParameterType(value)
	return nil
}

type APIMessage_APIMessageType int32

const (
	APIMessage_GetProjectRootRequest  APIMessage_APIMessageType = 1
	APIMessage_GetProjectRootResponse APIMessage_APIMessageType = 2
	APIMessage_GetAllStepsRequest     APIMessage_APIMessageType = 3
	APIMessage_GetAllStepResponse     APIMessage_APIMessageType = 4
	APIMessage_GetAllSpecsRequest     APIMessage_APIMessageType = 5
	APIMessage_GetAllSpecsResponse    APIMessage_APIMessageType = 6
)

var APIMessage_APIMessageType_name = map[int32]string{
	1: "GetProjectRootRequest",
	2: "GetProjectRootResponse",
	3: "GetAllStepsRequest",
	4: "GetAllStepResponse",
	5: "GetAllSpecsRequest",
	6: "GetAllSpecsResponse",
}
var APIMessage_APIMessageType_value = map[string]int32{
	"GetProjectRootRequest":  1,
	"GetProjectRootResponse": 2,
	"GetAllStepsRequest":     3,
	"GetAllStepResponse":     4,
	"GetAllSpecsRequest":     5,
	"GetAllSpecsResponse":    6,
}

func (x APIMessage_APIMessageType) Enum() *APIMessage_APIMessageType {
	p := new(APIMessage_APIMessageType)
	*p = x
	return p
}
func (x APIMessage_APIMessageType) String() string {
	return proto.EnumName(APIMessage_APIMessageType_name, int32(x))
}
func (x *APIMessage_APIMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APIMessage_APIMessageType_value, data, "APIMessage_APIMessageType")
	if err != nil {
		return err
	}
	*x = APIMessage_APIMessageType(value)
	return nil
}

type GetProjectRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetProjectRootRequest) Reset()         { *m = GetProjectRootRequest{} }
func (m *GetProjectRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootRequest) ProtoMessage()    {}

type GetProjectRootResponse struct {
	ProjectRoot      *string `protobuf:"bytes,1,req,name=projectRoot" json:"projectRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetProjectRootResponse) Reset()         { *m = GetProjectRootResponse{} }
func (m *GetProjectRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootResponse) ProtoMessage()    {}

func (m *GetProjectRootResponse) GetProjectRoot() string {
	if m != nil && m.ProjectRoot != nil {
		return *m.ProjectRoot
	}
	return ""
}

type GetAllStepsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllStepsRequest) Reset()         { *m = GetAllStepsRequest{} }
func (m *GetAllStepsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsRequest) ProtoMessage()    {}

type GetAllStepsResponse struct {
	Steps            []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetAllStepsResponse) Reset()         { *m = GetAllStepsResponse{} }
func (m *GetAllStepsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsResponse) ProtoMessage()    {}

func (m *GetAllStepsResponse) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

type GetAllSpecsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllSpecsRequest) Reset()         { *m = GetAllSpecsRequest{} }
func (m *GetAllSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsRequest) ProtoMessage()    {}

type GetAllSpecsResponse struct {
	Specs            []*ProtoSpec `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetAllSpecsResponse) Reset()         { *m = GetAllSpecsResponse{} }
func (m *GetAllSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsResponse) ProtoMessage()    {}

func (m *GetAllSpecsResponse) GetSpecs() []*ProtoSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

type ProtoSpec struct {
	Items            []*ProtoItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ProtoSpec) Reset()         { *m = ProtoSpec{} }
func (m *ProtoSpec) String() string { return proto.CompactTextString(m) }
func (*ProtoSpec) ProtoMessage()    {}

func (m *ProtoSpec) GetItems() []*ProtoItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ProtoItem struct {
	ItemType         *ProtoItem_ItemType `protobuf:"varint,1,req,name=itemType,enum=main.ProtoItem_ItemType" json:"itemType,omitempty"`
	Heading          *ProtoHeading       `protobuf:"bytes,2,opt,name=heading" json:"heading,omitempty"`
	Step             *ProtoStep          `protobuf:"bytes,3,opt,name=step" json:"step,omitempty"`
	Concept          *ProtoConcept       `protobuf:"bytes,4,opt,name=concept" json:"concept,omitempty"`
	Scenario         *ProtoScenario      `protobuf:"bytes,5,opt,name=scenario" json:"scenario,omitempty"`
	Comment          *ProtoComment       `protobuf:"bytes,6,opt,name=comment" json:"comment,omitempty"`
	Table            *ProtoTableParam    `protobuf:"bytes,7,opt,name=table" json:"table,omitempty"`
	Tags             *ProtoTags          `protobuf:"bytes,8,opt,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ProtoItem) Reset()         { *m = ProtoItem{} }
func (m *ProtoItem) String() string { return proto.CompactTextString(m) }
func (*ProtoItem) ProtoMessage()    {}

func (m *ProtoItem) GetItemType() ProtoItem_ItemType {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return ProtoItem_Heading
}

func (m *ProtoItem) GetHeading() *ProtoHeading {
	if m != nil {
		return m.Heading
	}
	return nil
}

func (m *ProtoItem) GetStep() *ProtoStep {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *ProtoItem) GetConcept() *ProtoConcept {
	if m != nil {
		return m.Concept
	}
	return nil
}

func (m *ProtoItem) GetScenario() *ProtoScenario {
	if m != nil {
		return m.Scenario
	}
	return nil
}

func (m *ProtoItem) GetComment() *ProtoComment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ProtoItem) GetTable() *ProtoTableParam {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *ProtoItem) GetTags() *ProtoTags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ProtoHeading struct {
	HeadingType      *ProtoHeading_HeadingType `protobuf:"varint,1,req,name=headingType,enum=main.ProtoHeading_HeadingType" json:"headingType,omitempty"`
	Text             *string                   `protobuf:"bytes,2,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ProtoHeading) Reset()         { *m = ProtoHeading{} }
func (m *ProtoHeading) String() string { return proto.CompactTextString(m) }
func (*ProtoHeading) ProtoMessage()    {}

func (m *ProtoHeading) GetHeadingType() ProtoHeading_HeadingType {
	if m != nil && m.HeadingType != nil {
		return *m.HeadingType
	}
	return ProtoHeading_Spec
}

func (m *ProtoHeading) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type ProtoScenario struct {
	ScenarioItems    []*ProtoItem `protobuf:"bytes,1,rep,name=scenarioItems" json:"scenarioItems,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ProtoScenario) Reset()         { *m = ProtoScenario{} }
func (m *ProtoScenario) String() string { return proto.CompactTextString(m) }
func (*ProtoScenario) ProtoMessage()    {}

func (m *ProtoScenario) GetScenarioItems() []*ProtoItem {
	if m != nil {
		return m.ScenarioItems
	}
	return nil
}

type ProtoStep struct {
	Text             *string                     `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	Parameters       []*Parameter                `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
	Fragments        []*Fragment                 `protobuf:"bytes,3,rep,name=fragments" json:"fragments,omitempty"`
	ExecStatus       []*ProtoStepExecutionStatus `protobuf:"bytes,4,rep,name=execStatus" json:"execStatus,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ProtoStep) Reset()         { *m = ProtoStep{} }
func (m *ProtoStep) String() string { return proto.CompactTextString(m) }
func (*ProtoStep) ProtoMessage()    {}

func (m *ProtoStep) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *ProtoStep) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ProtoStep) GetFragments() []*Fragment {
	if m != nil {
		return m.Fragments
	}
	return nil
}

func (m *ProtoStep) GetExecStatus() []*ProtoStepExecutionStatus {
	if m != nil {
		return m.ExecStatus
	}
	return nil
}

type ProtoStepExecutionStatus struct {
	IsPassed         *bool                     `protobuf:"varint,1,req,name=isPassed" json:"isPassed,omitempty"`
	StackTrace       *string                   `protobuf:"bytes,2,opt,name=stackTrace" json:"stackTrace,omitempty"`
	StepExecRequest  *main1.ExecuteStepRequest `protobuf:"bytes,3,req,name=stepExecRequest" json:"stepExecRequest,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ProtoStepExecutionStatus) Reset()         { *m = ProtoStepExecutionStatus{} }
func (m *ProtoStepExecutionStatus) String() string { return proto.CompactTextString(m) }
func (*ProtoStepExecutionStatus) ProtoMessage()    {}

func (m *ProtoStepExecutionStatus) GetIsPassed() bool {
	if m != nil && m.IsPassed != nil {
		return *m.IsPassed
	}
	return false
}

func (m *ProtoStepExecutionStatus) GetStackTrace() string {
	if m != nil && m.StackTrace != nil {
		return *m.StackTrace
	}
	return ""
}

func (m *ProtoStepExecutionStatus) GetStepExecRequest() *main1.ExecuteStepRequest {
	if m != nil {
		return m.StepExecRequest
	}
	return nil
}

type ProtoConcept struct {
	ConceptStep      *ProtoStep   `protobuf:"bytes,1,req,name=conceptStep" json:"conceptStep,omitempty"`
	Steps            []*ProtoStep `protobuf:"bytes,2,rep,name=steps" json:"steps,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ProtoConcept) Reset()         { *m = ProtoConcept{} }
func (m *ProtoConcept) String() string { return proto.CompactTextString(m) }
func (*ProtoConcept) ProtoMessage()    {}

func (m *ProtoConcept) GetConceptStep() *ProtoStep {
	if m != nil {
		return m.ConceptStep
	}
	return nil
}

func (m *ProtoConcept) GetSteps() []*ProtoStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

type Fragment struct {
	FragmentType     *Fragment_FragmentType `protobuf:"varint,1,req,name=fragmentType,enum=main.Fragment_FragmentType" json:"fragmentType,omitempty"`
	Text             *string                `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Parameter        *Parameter             `protobuf:"bytes,3,opt,name=parameter" json:"parameter,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Fragment) Reset()         { *m = Fragment{} }
func (m *Fragment) String() string { return proto.CompactTextString(m) }
func (*Fragment) ProtoMessage()    {}

func (m *Fragment) GetFragmentType() Fragment_FragmentType {
	if m != nil && m.FragmentType != nil {
		return *m.FragmentType
	}
	return Fragment_Text
}

func (m *Fragment) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Fragment) GetParameter() *Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

type Parameter struct {
	ParameterType    *Parameter_ParameterType `protobuf:"varint,1,req,name=parameterType,enum=main.Parameter_ParameterType" json:"parameterType,omitempty"`
	Value            *string                  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Table            *ProtoTableParam         `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}

func (m *Parameter) GetParameterType() Parameter_ParameterType {
	if m != nil && m.ParameterType != nil {
		return *m.ParameterType
	}
	return Parameter_Static
}

func (m *Parameter) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Parameter) GetTable() *ProtoTableParam {
	if m != nil {
		return m.Table
	}
	return nil
}

type ProtoComment struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoComment) Reset()         { *m = ProtoComment{} }
func (m *ProtoComment) String() string { return proto.CompactTextString(m) }
func (*ProtoComment) ProtoMessage()    {}

func (m *ProtoComment) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type ProtoTableParam struct {
	Headers          *ProtoTableRow   `protobuf:"bytes,1,req,name=headers" json:"headers,omitempty"`
	Rows             []*ProtoTableRow `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ProtoTableParam) Reset()         { *m = ProtoTableParam{} }
func (m *ProtoTableParam) String() string { return proto.CompactTextString(m) }
func (*ProtoTableParam) ProtoMessage()    {}

func (m *ProtoTableParam) GetHeaders() *ProtoTableRow {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *ProtoTableParam) GetRows() []*ProtoTableRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ProtoTableRow struct {
	Cells            []string `protobuf:"bytes,1,rep,name=cells" json:"cells,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoTableRow) Reset()         { *m = ProtoTableRow{} }
func (m *ProtoTableRow) String() string { return proto.CompactTextString(m) }
func (*ProtoTableRow) ProtoMessage()    {}

func (m *ProtoTableRow) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

type APIMessage struct {
	MessageType         *APIMessage_APIMessageType `protobuf:"varint,1,req,name=messageType,enum=main.APIMessage_APIMessageType" json:"messageType,omitempty"`
	MessageId           *int64                     `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	ProjectRootRequest  *GetProjectRootRequest     `protobuf:"bytes,3,opt,name=projectRootRequest" json:"projectRootRequest,omitempty"`
	ProjectRootResponse *GetProjectRootResponse    `protobuf:"bytes,4,opt,name=projectRootResponse" json:"projectRootResponse,omitempty"`
	AllStepsRequest     *GetAllStepsRequest        `protobuf:"bytes,5,opt,name=allStepsRequest" json:"allStepsRequest,omitempty"`
	AllStepsResponse    *GetAllStepsResponse       `protobuf:"bytes,6,opt,name=allStepsResponse" json:"allStepsResponse,omitempty"`
	AllSpecsRequest     *GetAllSpecsRequest        `protobuf:"bytes,7,opt,name=allSpecsRequest" json:"allSpecsRequest,omitempty"`
	AllSpecsResponse    *GetAllSpecsResponse       `protobuf:"bytes,8,opt,name=allSpecsResponse" json:"allSpecsResponse,omitempty"`
	XXX_unrecognized    []byte                     `json:"-"`
}

func (m *APIMessage) Reset()         { *m = APIMessage{} }
func (m *APIMessage) String() string { return proto.CompactTextString(m) }
func (*APIMessage) ProtoMessage()    {}

func (m *APIMessage) GetMessageType() APIMessage_APIMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return APIMessage_GetProjectRootRequest
}

func (m *APIMessage) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *APIMessage) GetProjectRootRequest() *GetProjectRootRequest {
	if m != nil {
		return m.ProjectRootRequest
	}
	return nil
}

func (m *APIMessage) GetProjectRootResponse() *GetProjectRootResponse {
	if m != nil {
		return m.ProjectRootResponse
	}
	return nil
}

func (m *APIMessage) GetAllStepsRequest() *GetAllStepsRequest {
	if m != nil {
		return m.AllStepsRequest
	}
	return nil
}

func (m *APIMessage) GetAllStepsResponse() *GetAllStepsResponse {
	if m != nil {
		return m.AllStepsResponse
	}
	return nil
}

func (m *APIMessage) GetAllSpecsRequest() *GetAllSpecsRequest {
	if m != nil {
		return m.AllSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetAllSpecsResponse() *GetAllSpecsResponse {
	if m != nil {
		return m.AllSpecsResponse
	}
	return nil
}

type ProtoTags struct {
	Tags             []string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoTags) Reset()         { *m = ProtoTags{} }
func (m *ProtoTags) String() string { return proto.CompactTextString(m) }
func (*ProtoTags) ProtoMessage()    {}

func (m *ProtoTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.ProtoItem_ItemType", ProtoItem_ItemType_name, ProtoItem_ItemType_value)
	proto.RegisterEnum("main.ProtoHeading_HeadingType", ProtoHeading_HeadingType_name, ProtoHeading_HeadingType_value)
	proto.RegisterEnum("main.Fragment_FragmentType", Fragment_FragmentType_name, Fragment_FragmentType_value)
	proto.RegisterEnum("main.Parameter_ParameterType", Parameter_ParameterType_name, Parameter_ParameterType_value)
	proto.RegisterEnum("main.APIMessage_APIMessageType", APIMessage_APIMessageType_name, APIMessage_APIMessageType_value)
}
