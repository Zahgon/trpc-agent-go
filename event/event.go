//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package event provides the event system for agent communication.
package event

import (
	"context"
	"encoding/json"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// InitVersion is the initial version of the event format.
	InitVersion int = iota // 0

	// CurrentVersion is the current version of the event format.
	CurrentVersion
)

const (
	// EmitWithoutTimeout is the default timeout for emitting events.
	EmitWithoutTimeout = 0 * time.Second

	// FilterKeyDelimiter is the delimiter for hierarchical event filtering.
	FilterKeyDelimiter = "/"

	// TagDelimiter is the delimiter for event tags.
	TagDelimiter = ";"
)

const (
	// CodeExecutionTag is the tag value for code execution code event.
	CodeExecutionTag = "code_execution_code"

	// CodeExecutionResultTag is the tag value for code execution result event.
	CodeExecutionResultTag = "code_execution_result"

	// TransferTag is the tag for transfer event.
	TransferTag = "transfer"

	// ToolCallArgsExtensionKey stores tool call arguments keyed by tool call ID
	// on tool result events.
	ToolCallArgsExtensionKey = "trpc_agent.tool_call_args"
)

// Event represents an event in conversation between agents and users.
type Event struct {
	// Response is the base struct for all LLM response functionality.
	*model.Response

	// RequestID is the request ID of the event.
	RequestID string `json:"requestID,omitempty"`

	// InvocationID is the invocation ID of the event.
	InvocationID string `json:"invocationId"`

	// ParentInvocationID is the parent invocation ID of the event.
	ParentInvocationID string `json:"parentInvocationId,omitempty"`

	// Author is the author of the event.
	Author string `json:"author"`

	// ID is the unique identifier of the event.
	ID string `json:"id"`

	// Timestamp is the timestamp of the event.
	Timestamp time.Time `json:"timestamp"`

	// Branch records agent execution chain information.
	// In multi-agent mode, this is useful for tracing agent execution trajectories.
	Branch string `json:"branch,omitempty"`

	// Tag Uses tags to annotate events with business-specific labels.
	Tag string `json:"tag,omitempty"`

	// RequiresCompletion indicates if this event needs completion signaling.
	RequiresCompletion bool `json:"requiresCompletion,omitempty"`

	// LongRunningToolIDs is the Set of ids of the long running function calls.
	// Agent client will know from this field about which function call is long running.
	// only valid for function call event
	LongRunningToolIDs map[string]struct{} `json:"longRunningToolIDs,omitempty"`

	// StateDelta contains state changes to be applied to the session.
	StateDelta map[string][]byte `json:"stateDelta,omitempty"`

	// Extensions stores optional event metadata in a namespaced,
	// versioned JSON format.
	Extensions map[string]json.RawMessage `json:"extensions,omitempty"`

	// StructuredOutput carries a typed, in-memory structured output payload.
	// This is not serialized and is meant for immediate consumer access.
	StructuredOutput any `json:"-"`
	// ExecutionTrace carries an in-memory execution trace artifact for this run.
	// This is not serialized and is meant for immediate consumer access.
	ExecutionTrace *trace.Trace `json:"-"`

	// Actions carry flow-level hints that influence how this event is treated
	// by the runner/flow (e.g., skip summarization after a tool response).
	Actions *EventActions `json:"actions,omitempty"`

	// filterKey is identifier for hierarchical event filtering.
	FilterKey string `json:"filterKey,omitempty"`

	// version for handling version compatibility issues.
	Version int `json:"version,omitempty"`
}

// ContainsTag checks if the event contains the specified tag.
func (e *Event) ContainsTag(tag string) bool { _ = "STUB: not implemented"; return false }

// EventActions represents optional actions/hints attached to an event.
// These are used by the flow to adjust control behavior without
// overloading Response fields.
type EventActions struct {
	// SkipSummarization indicates that the flow should not run an
	// additional summarization step after this event. Commonly used
	// for final tool.response events returned by AgentTool.
	SkipSummarization bool `json:"skipSummarization,omitempty"`
}

// Clone creates a deep copy of the event.
func (e *Event) Clone() *Event { _ = "STUB: not implemented"; return nil }

func cloneRawMessage(raw json.RawMessage) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func cloneExecutionTrace(executionTrace *trace.Trace) *trace.Trace {
	_ = "STUB: not implemented"
	return nil
}

func cloneExecutionTraceStep(step trace.Step) trace.Step {
	_ = "STUB: not implemented"
	return *new(trace.Step)
}

func cloneExecutionTraceSnapshot(snapshot *trace.Snapshot) *trace.Snapshot {
	_ = "STUB: not implemented"
	return nil
}

// Filter checks if the event matches the specified filter key.
func (e *Event) Filter(filterKey string) bool { _ = "STUB: not implemented"; return false }

// New creates a new Event with generated ID and timestamp.
func New(invocationID, author string, opts ...Option) *Event { _ = "STUB: not implemented"; return nil }

// NewErrorEvent creates a new error Event with the specified error details.
// This provides a clean way to create error events without manual field assignment.
func NewErrorEvent(invocationID, author, errorType, errorMessage string,
	opts ...Option) *Event {
	_ = "STUB: not implemented"
	return nil
}

// NewResponseEvent creates a new Event from a model Response.
func NewResponseEvent(invocationID, author string, response *model.Response,
	opts ...Option) *Event {
	_ = "STUB: not implemented"
	return nil
}

// DefaultEmitTimeoutErr is the default error returned when a wait notice times out.
var DefaultEmitTimeoutErr = NewEmitEventTimeoutError("emit event timeout.")

// EmitEventTimeoutError represents an error that signals the emit event timeout.
type EmitEventTimeoutError struct {
	// Message contains the stop reason
	Message string
}

// Error implements the error interface.
func (e *EmitEventTimeoutError) Error() string {
	_ = "STUB: not implemented"

	// AsEmitEventTimeoutError checks if an error is a EmitEventTimeoutError using errors.As.
	return ""
}

func AsEmitEventTimeoutError(err error) (*EmitEventTimeoutError, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// NewEmitEventTimeoutError creates a new EmitEventTimeoutError with the given message.
func NewEmitEventTimeoutError(message string) *EmitEventTimeoutError {
	_ = "STUB: not implemented"
	return nil
}

// IsRunnerCompletion reports whether this event is the terminal completion
// event emitted by Runner. It is the most reliable signal that the entire
// run has finished (regardless of the specific Agent implementation), and the
// recommended condition to stop consuming the event stream.
func (e *Event) IsRunnerCompletion() bool { _ = "STUB: not implemented"; return false }

// IsError reports whether this event carries any error signal.
// It is broader than IsTerminalError and also matches non-terminal graph
// observability events such as graph.node.error.
func (e *Event) IsError() bool { _ = "STUB: not implemented"; return false }

// IsTerminalError reports whether this event represents a terminal failure for
// the overall run.
func (e *Event) IsTerminalError() bool { _ = "STUB: not implemented"; return false }

// EmitEvent sends an event to the channel without timeout.
func EmitEvent(ctx context.Context, ch chan<- *Event, e *Event) error {
	_ = "STUB: not implemented"
	return nil
}

// snapshotEvent returns a string representation of e if trace logging is
// enabled, or an empty string otherwise. The snapshot must be taken while the
// caller still holds exclusive ownership of *e — before ch <- e — because
// once the send completes the receiver may mutate the struct concurrently.
func snapshotEvent(e *Event) string { _ = "STUB: not implemented"; return "" }

func redactedEventForLogging(e *Event) Event { _ = "STUB: not implemented"; return *new(Event) }

func tryEmitReadyEvent(ctx context.Context, ch chan<- *Event, e *Event) (bool, error) {
	_ = "STUB: not implemented"
	// Snapshot before send: once ch <- e returns, the receiver owns *e and
	// may mutate it concurrently (runner.copyEventInvocationFields). Reading
	// *e after the send for logging is a data race.
	return false, nil
}

// EmitEventWithTimeout sends an event to the channel with optional timeout.
func EmitEventWithTimeout(ctx context.Context, ch chan<- *Event,
	e *Event, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// If the context is already cancelled, prefer returning immediately
// rather than attempting to send. This avoids a racy select where both
// the send and the ctx.Done() cases are ready, which could otherwise
// result in emitting an event after cancellation.

// Fall back to a blocking send. Snapshot before send — same race as above.

// Snapshot before send — same race as above.

// MarshalJSON implements json.Marshaler and produces a format that
// preserves legacy flattened fields while also embedding minimal
// response metadata (ID/timestamp) under the dedicated "response" key.
func (e Event) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler by accepting both legacy flattened
// payloads and the new nested-response representation, preferring the nested
// response when present.
func (e *Event) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// First parse the flat structure.
	return nil
}

// Then try to read nested metadata.

// Tolerate nested part failure, it does not affect the overall failure, preserve the flat fields.

// eventNoMethods is the alias of Event for avoiding recursive calls of custom MarshalJSON/UnmarshalJSON.
type eventNoMethods Event

// responseMeta is the minimal response metadata for JSON nested.
type responseMeta struct {
	ID        string    `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// jsonEvent is the final JSON structure to be output/read,
// including flat event fields and nested response metadata.
type jsonEvent struct {
	*eventNoMethods
	Response *responseMeta `json:"response,omitempty"`
}
