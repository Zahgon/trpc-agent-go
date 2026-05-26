//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package telemetry provides telemetry and observability functionality for the trpc-agent-go framework.
// It includes tracing, metrics, and monitoring capabilities for agent operations.
package telemetry

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// grpcDial is a package-level variable to allow test injection of a custom dialer.
// In production, this points to grpc.Dial.
var grpcDial = grpc.Dial

// telemetry service constants.
const (
	ServiceName      = "telemetry"
	ServiceVersion   = "v0.1.0"
	ServiceNamespace = "trpc-go-agent"
	InstrumentName   = "trpc.agent.go"

	SpanNamePrefixExecuteTool = "execute_tool"

	OperationExecuteTool     = "execute_tool"
	OperationChat            = "chat"
	OperationGenerateContent = "generate_content"
	OperationInvokeAgent     = "invoke_agent"
	OperationCreateAgent     = "create_agent"
	OperationEmbeddings      = "embeddings"
	OperationWorkflow        = "workflow"
)

// NewChatSpanName creates a new chat span name.
func NewChatSpanName(requestModel string) string { _ = "STUB: not implemented"; return "" }

// NewExecuteToolSpanName creates a new execute tool span name.
func NewExecuteToolSpanName(toolName string) string { _ = "STUB: not implemented"; return "" }

// WorkflowType is the normalized type vocabulary used by workflow spans.
type WorkflowType string

// Standard workflow type values.
const (
	WorkflowTypeGraph    WorkflowType = "graph"
	WorkflowTypeFunction WorkflowType = "function"
	WorkflowTypeLLM      WorkflowType = "llm"
	WorkflowTypeTool     WorkflowType = "tool"
	WorkflowTypeAgent    WorkflowType = "agent"
	WorkflowTypeJoin     WorkflowType = "join"
	WorkflowTypeRouter   WorkflowType = "router"
)

// String returns the string representation of the workflow type.
func (wt WorkflowType) String() string {
	_ = "STUB: not implemented"

	// Workflow is the workflow information.
	return ""
}

type Workflow struct {
	Name     string
	ID       string
	Type     WorkflowType
	Request  any
	Response any
	Error    error
}

// NewWorkflowSpanName creates a new workflow span name.
func NewWorkflowSpanName(workflowName string) string { _ = "STUB: not implemented"; return "" }

type telemetryMessage struct {
	Role             model.Role          `json:"role"`
	Content          string              `json:"content,omitempty"`
	ContentParts     []model.ContentPart `json:"content_parts,omitempty"`
	ToolCallID       string              `json:"tool_call_id,omitempty"`
	Name             string              `json:"name,omitempty"`
	ToolCalls        []model.ToolCall    `json:"tool_calls,omitempty"`
	ReasoningContent string              `json:"reasoning_content,omitempty"`
}

type telemetryChoice struct {
	Index        int              `json:"index"`
	Message      telemetryMessage `json:"message,omitempty"`
	Delta        telemetryMessage `json:"delta,omitempty"`
	FinishReason *string          `json:"finish_reason,omitempty"`
}

func telemetryMessageFromModel(msg model.Message) telemetryMessage {
	_ = "STUB: not implemented"
	return *new(telemetryMessage)
}

func marshalTelemetryMessages(messages []model.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalTelemetryChoices(choices []model.Choice) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TraceWorkflow traces the workflow.
func TraceWorkflow(span trace.Span, workflow *Workflow) { _ = "STUB: not implemented"; return }

// newInferenceSpanName creates a new inference span name.
// inference operation name: "chat" for openai, "generate_content" for gemini.
// For example, "chat gpt-4.0".
func newInferenceSpanName(operationNames, requestModel string) string {
	_ = "STUB: not implemented"
	return ""
}

const (
	// ProtocolGRPC uses gRPC protocol for OTLP exporter.
	ProtocolGRPC string = "grpc"
	// ProtocolHTTP uses HTTP protocol for OTLP exporter.
	ProtocolHTTP string = "http"
)

// TraceToolCall traces the invocation of a tool call.
func TraceToolCall(span trace.Span, sess *session.Session, declaration *tool.Declaration, args []byte, rspEvent *event.Event, err error) {
	_ = "STUB: not implemented"
	return
}

// args is json-encoded.

// Setting empty llm request and response (as UI expect these) while not
// applicable for tool_response.

// ToolNameMergedTools is the name of the merged tools.
const ToolNameMergedTools = "(merged tools)"

// TraceMergedToolCalls traces the invocation of a merged tool call.
// Calling this function is not needed for telemetry purposes. This is provided
// for preventing trace-query requests typically sent by web UIs.
func TraceMergedToolCalls(span trace.Span, rspEvent *event.Event) {
	_ = "STUB: not implemented"
	return
}

// Setting empty llm request and response (as UI expect these) while not
// applicable for tool_response.

func resolveInvocationAgentIdentity(invoke *agent.Invocation) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Invocation does not carry a canonical agent ID today, so use
// Invocation.AgentName as the fallback for gen_ai.agent.id.

// TraceBeforeInvokeAgent traces the before invocation of an agent.
func TraceBeforeInvokeAgent(span trace.Span, invoke *agent.Invocation, agentDescription, instructions string, genConfig *model.GenerationConfig) {
	_ = "STUB: not implemented"
	return
}

func traceBeforeInvokeAgentInvocation(span trace.Span, invoke *agent.Invocation) {
	_ = "STUB: not implemented"
	return
}

func setInvokeAgentInputMessageAttributes(span trace.Span, msg model.Message) {
	_ = "STUB: not implemented"
	return
}

func beforeInvokeAgentAttributes(invoke *agent.Invocation) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func setInvokeAgentGenerationConfigAttributes(span trace.Span, genConfig *model.GenerationConfig) {
	_ = "STUB: not implemented"
	return
}

// TokenUsage is token usage information.
type TokenUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// TraceAfterInvokeAgent traces the after invocation of an agent.
func TraceAfterInvokeAgent(
	span trace.Span,
	rspEvent *event.Event,
	tokenUsage *TokenUsage,
	timeToFirstToken time.Duration,
	errorTypeFallback string,
) {
	_ = "STUB: not implemented"
	return
}

// TraceChatAttributes contains TraceChat inputs other than span.
//
// It is used to keep TraceChat signatures stable as parameters evolve.
type TraceChatAttributes struct {
	Invocation       *agent.Invocation
	Request          *model.Request
	Response         *model.Response
	EventID          string
	TimeToFirstToken time.Duration
	TaskType         string
}

// NewSummarizeTaskType creates a task type for summarize.
func NewSummarizeTaskType(name string) string { _ = "STUB: not implemented"; return "" }

// TraceChat traces the invocation of an LLM call.
func TraceChat(span trace.Span, attributes *TraceChatAttributes) { _ = "STUB: not implemented"; return }

// Add invocation attributes

// Add request attributes

// Add response attributes

// Set all attributes at once

// Handle response error status

// buildInvocationAttributes extracts attributes from the invocation.
func buildInvocationAttributes(invoke *agent.Invocation) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// buildRequestAttributes builds request-related attributes.
func buildRequestAttributes(req *model.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Add generation config attributes

// Add stream attribute only when it's true

// Add request body

// Add tool definitions as best-effort structured array (JSON string fallback)

// Add messages

// buildResponseAttributes builds response-related attributes.
func buildResponseAttributes(rsp *model.Response, errorTypeFallback string) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Add error type if present

// Add usage attributes

// Prompt cache tokens (if provided by the model provider)

// OpenAI: cached_tokens

// Anthropic: cache_read_tokens

// Anthropic: cache_creation_tokens

// Add choices attributes

// Extract finish reasons

// Add response body

func responseErrorAttributes(respErr *model.ResponseError, fallback string) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// NewGRPCConn creates a new gRPC connection to the OpenTelemetry Collector.
func NewGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	// It connects the OpenTelemetry Collector through gRPC connection.
	// You can customize the endpoint using SetConfig() or environment variables.
	return nil, nil
}

// Note the use of insecure transport here. TLS is recommended in production.
