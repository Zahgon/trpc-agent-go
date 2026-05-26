//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package a2aagent

import (
	"encoding/json"

	"trpc.group/trpc-go/trpc-a2a-go/client"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// StreamingRespHandler handles the streaming response content
// return the content will be added to the final aggregated content
type StreamingRespHandler func(resp *model.Response) (string, error)

// ConvertToA2AMessageFunc is the function signature for converting an invocation to an A2A protocol message.
type ConvertToA2AMessageFunc func(isStream bool, agentName string, invocation *agent.Invocation) (*protocol.Message, error)

// BuildMessageHook wraps the A2A message conversion with additional functionality.
// The hook receives the next converter function and returns a new converter function.
// Users can modify the invocation before calling next, modify the message after calling next,
// or completely replace the conversion logic by not calling next.
//
// This follows the same middleware pattern as server-side ProcessMessageHook.
type BuildMessageHook func(next ConvertToA2AMessageFunc) ConvertToA2AMessageFunc

// A2ADataPartToolResponse is a public tool response payload used by custom
// DataPart mappers.
type A2ADataPartToolResponse struct {
	ID      string
	Name    string
	Content string
}

// A2ADataPartMappingResult is the mapper-visible result holder used to enrich
// conversion output without depending on the converter's internal parseResult.
//
// Mappers receive a snapshot initialized from the current parse state. Changes
// are applied only when the mapper returns matched=true.
type A2ADataPartMappingResult struct {
	textContent            string
	reasoningContent       string
	toolCalls              []model.ToolCall
	toolResponses          []A2ADataPartToolResponse
	codeExecution          string
	codeExecutionResult    string
	eventExtensions        map[string]json.RawMessage
	textContentSet         bool
	reasoningContentSet    bool
	codeExecutionSet       bool
	codeExecutionResultSet bool
}

// GetTextContent returns the current text content snapshot.
func (r *A2ADataPartMappingResult) GetTextContent() string { _ = "STUB: not implemented"; return "" }

// SetTextContent overwrites text content when the mapper matches.
func (r *A2ADataPartMappingResult) SetTextContent(text string) { _ = "STUB: not implemented"; return }

// GetReasoningContent returns the current reasoning content snapshot.
func (r *A2ADataPartMappingResult) GetReasoningContent() string {
	_ = "STUB: not implemented"
	return ""
}

// SetReasoningContent overwrites reasoning content when the mapper matches.
func (r *A2ADataPartMappingResult) SetReasoningContent(text string) {
	_ = "STUB: not implemented"
	return
}

// AppendToolCall appends a tool call when the mapper matches.
func (r *A2ADataPartMappingResult) AppendToolCall(call model.ToolCall) {
	_ = "STUB: not implemented"
	return
}

// AppendToolResponse appends a tool response when the mapper matches.
func (r *A2ADataPartMappingResult) AppendToolResponse(resp A2ADataPartToolResponse) {
	_ = "STUB: not implemented"
	return
}

// GetCodeExecution returns the current executable code snapshot.
func (r *A2ADataPartMappingResult) GetCodeExecution() string { _ = "STUB: not implemented"; return "" }

// SetCodeExecution overwrites executable code when the mapper matches.
func (r *A2ADataPartMappingResult) SetCodeExecution(code string) { _ = "STUB: not implemented"; return }

// GetCodeExecutionResult returns the current code execution result snapshot.
func (r *A2ADataPartMappingResult) GetCodeExecutionResult() string {
	_ = "STUB: not implemented"
	return ""
}

// SetCodeExecutionResult overwrites code execution result when the mapper matches.
func (r *A2ADataPartMappingResult) SetCodeExecutionResult(result string) {
	_ = "STUB: not implemented"
	return
}

// SetEventExtension stores one serialized event extension when the mapper matches.
//
// This is useful for preserving custom A2A DataPart payloads through graph and
// server pipelines without forcing them into Message.Content.
func (r *A2ADataPartMappingResult) SetEventExtension(key string, value any) error {
	_ = "STUB: not implemented"
	return nil
}

func cloneA2AExtensionRawMessage(raw json.RawMessage) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func cloneA2AExtensions(
	extensions map[string]json.RawMessage,
) map[string]json.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

// A2ADataPartMapper maps an inbound A2A DataPart into the default parser result.
//
// Built-in DataPart handling (function call/response, code execution) runs
// first. Mappers are invoked only when the DataPart is not consumed by the
// built-ins. Returning matched=true means this mapper consumed the part.
// Returning matched=false leaves the part ignored by the default converter.
type A2ADataPartMapper func(part *protocol.DataPart, result *A2ADataPartMappingResult) (
	matched bool,
	err error,
)

// Option configures the A2AAgent
type Option func(*A2AAgent)

// WithName sets the name of agent
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDescription sets the agent description
func WithDescription(description string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAgentCardURL set the agent card URL
func WithAgentCardURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAgentCard set the agent card
func WithAgentCard(agentCard *server.AgentCard) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCustomEventConverter adds a custom A2A event converter to the A2AAgent.
func WithCustomEventConverter(converter A2AEventConverter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithA2ADataPartMapper registers a lightweight inbound DataPart mapper on the
// default A2A event converter.
//
// If WithCustomEventConverter provides a custom converter, this mapper is
// ignored.
func WithA2ADataPartMapper(mapper A2ADataPartMapper) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCustomA2AConverter adds a custom A2A message converter to the A2AAgent.
// This converter will be used to convert invocations to A2A protocol messages.
func WithCustomA2AConverter(converter InvocationA2AConverter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithA2AClientExtraOptions adds extra options to the A2A client.
func WithA2AClientExtraOptions(opts ...client.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamingChannelBufSize set the buf size of streaming protocol
func WithStreamingChannelBufSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamingRespHandler sets a handler function to process streaming responses.
func WithStreamingRespHandler(handler StreamingRespHandler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTransferStateKey sets the keys in session state to transfer to the A2A agent message by metadata.
//
// Supported patterns:
//   - "*"         — transfer all keys from RuntimeState
//   - "prefix*"   — transfer keys with the given prefix (e.g. "user.*" or "user*")
//   - "*suffix"   — transfer keys with the given suffix (e.g. "*.id" or "*id")
//   - "exact_key" — transfer only the exact key
func WithTransferStateKey(key ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBuildMessageHook sets a hook to customize the A2A message conversion.
// The hook wraps the default converter (including transferStateKey processing) as a middleware,
// following the same pattern as server-side WithProcessMessageHook.
//
// Example - modify message after conversion:
//
//	a2aagent.WithBuildMessageHook(func(next a2aagent.ConvertToA2AMessageFunc) a2aagent.ConvertToA2AMessageFunc {
//	    return func(isStream bool, agentName string, inv *agent.Invocation) (*protocol.Message, error) {
//	        msg, err := next(isStream, agentName, inv)
//	        if err != nil {
//	            return nil, err
//	        }
//	        // inject custom metadata
//	        if msg.Metadata == nil {
//	            msg.Metadata = make(map[string]any)
//	        }
//	        msg.Metadata["custom_key"] = "custom_value"
//	        return msg, nil
//	    }
//	})
func WithBuildMessageHook(hook BuildMessageHook) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUserIDHeader sets the HTTP header name to send UserID to the A2A server.
// If not set, defaults to "X-User-ID".
// The UserID will be extracted from invocation.Session.UserID and sent via the specified header.
func WithUserIDHeader(header string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableStreaming explicitly controls whether to use streaming protocol.
// If not set (nil), the agent will use the streaming capability from the agent card.
// This option overrides the agent card's capability setting.
func WithEnableStreaming(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }
