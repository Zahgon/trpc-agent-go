//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"reflect"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// StateKeyUserInput is the key of the user input.
	// It is consumed once and then cleared after successful LLM execution.
	StateKeyUserInput = "user_input"
	// StateKeyOneShotMessages is the key for one-shot messages that override
	// the current round input completely. It is consumed once and then cleared.
	StateKeyOneShotMessages = "one_shot_messages"
	// StateKeyOneShotMessagesByNode stores one-shot messages scoped to a target
	// node ID. The value is a map[nodeID][]model.Message, allowing parallel
	// branches to prepare one-shot inputs for different LLM nodes without
	// clobbering a shared global key.
	StateKeyOneShotMessagesByNode = "one_shot_messages_by_node"
	// StateKeyLastResponse is the key of the last response.
	StateKeyLastResponse = "last_response"
	// StateKeyLastToolResponse stores the last tool output as a JSON string.
	// It is set by Tools nodes after successful tool execution.
	StateKeyLastToolResponse = "last_tool_response"
	// StateKeyNodeResponses is the key of the node responses.
	StateKeyNodeResponses = "node_responses"
	// StateKeySession is the key of the session.
	StateKeySession = "session"
	// StateKeyMessages is the key of the messages.
	// Typically it is used and updated by the LLM node.
	StateKeyMessages = "messages"
	// StateKeyLastResponseID stores the ID of the last model response.
	StateKeyLastResponseID = "last_response_id"
	// StateKeyMetadata is the key of the metadata.
	StateKeyMetadata = "metadata"
	// StateKeyExecContext is the key of the execution context.
	StateKeyExecContext = "exec_context"
	// StateKeyToolCallbacks is the key of the tool callbacks.
	StateKeyToolCallbacks = "tool_callbacks"
	// StateKeyModelCallbacks is the key of the model callbacks.
	StateKeyModelCallbacks = "model_callbacks"
	// StateKeyAgentCallbacks is the key of the agent callbacks.
	StateKeyAgentCallbacks = "agent_callbacks"
	// StateKeyCurrentNodeID is the key for storing the current node ID in the state.
	StateKeyCurrentNodeID = "current_node_id"
	// StateKeyParentAgent is the key for storing the parent GraphAgent that owns sub-agents.
	StateKeyParentAgent = "parent_agent"
)

const currentTraceStepIDStateKey = "__current_trace_step_id__"

// State represents the state that flows through the graph.
// This is the shared data structure that flows between nodes.
type State map[string]any

// GetStateValue retrieves a typed value from the state.
//
// Returns the typed value and true if the key exists and the type matches,
// or the zero value and false otherwise.
//
// Example:
//
//	if messages, ok := GetStateValue[[]model.Message](state, StateKeyMessages); ok {
//	    // use messages
//	}
//	if userInput, ok := GetStateValue[string](state, StateKeyUserInput); ok {
//	    // use userInput
//	}
func GetStateValue[T any](s State, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// SetOneShotMessagesForNode sets one-shot messages for a specific target node.
// It writes an incremental update to StateKeyOneShotMessagesByNode.
func SetOneShotMessagesForNode(
	nodeID string,
	msgs []model.Message,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// SetOneShotMessagesByNode sets one-shot messages for multiple target nodes.
//
// It writes an incremental update to StateKeyOneShotMessagesByNode:
//
//   - Each entry's message slice is defensively copied.
//   - An empty message slice clears that node's entry.
//   - Empty node IDs are ignored.
//
// This is useful when a single upstream node needs to prepare one-shot inputs
// for many downstream LLM nodes in one return value.
func SetOneShotMessagesByNode(
	byNode map[string][]model.Message,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// ClearOneShotMessagesByNode clears the entire one-shot-by-node map.
func ClearOneShotMessagesByNode() State { _ = "STUB: not implemented"; return *new(State) }

// ClearOneShotMessagesForNode clears one-shot messages for a specific node by
// deleting its entry in StateKeyOneShotMessagesByNode.
func ClearOneShotMessagesForNode(nodeID string) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// GetOneShotMessagesForNode retrieves one-shot messages for a specific node.
// It returns a defensive copy of the message slice.
func GetOneShotMessagesForNode(
	state State,
	nodeID string,
) ([]model.Message, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Clone creates a deep copy of the state.
func (s State) Clone() State { _ = "STUB: not implemented"; return *new(State) }

func (s State) safeClone() State { _ = "STUB: not implemented"; return *new(State) }

// Use jsonSafeCopy so that nested non-serializable
// types (chan, func, sync.Mutex, etc.) are stripped
// and the result is safe for json.Marshal.

func (s State) deepCopy(retainUnsafeKey bool, fields map[string]StateField) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// StateReducer is a function that determines how state updates are merged.
// It takes existing and new values and returns the merged result.
type StateReducer func(existing, update any) any

// StateField defines a field in the state schema with its type and reducer.
type StateField struct {
	Type            reflect.Type
	Reducer         StateReducer
	Default         func() any
	Required        bool
	DisableDeepCopy bool
}

// StateSchema defines the structure and behavior of graph state.
// This defines the structure and behavior of state.
type StateSchema struct {
	mu     sync.RWMutex
	Fields map[string]StateField
}

// NewStateSchema creates a new state schema.
func NewStateSchema() *StateSchema { _ = "STUB: not implemented"; return nil }

// AddField adds a field to the state schema.
func (s *StateSchema) AddField(name string, field StateField) *StateSchema {
	_ = "STUB: not implemented"
	return nil
}

// ApplyUpdate applies a state update using the defined reducers.
func (s *StateSchema) ApplyUpdate(currentState State, update State) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// Ignore internal/ephemeral keys in updates. They are owned by
// the executor and may contain concurrently-mutated maps.

// If no field definition, use default behavior (override) with
// deep copy to avoid sharing mutable references across goroutines.

// Apply reducer with deep-copied update to prevent reference sharing.

// Ensure merged complex values are not shared by taking a deep copy.

// Validate validates a state against the schema.
func (s *StateSchema) Validate(state State) error { _ = "STUB: not implemented"; return nil }

// validateSchema validates the schema struct.
func (s *StateSchema) validateSchema() error { _ = "STUB: not implemented"; return nil }

// Validate that Type and Reducer are not nil.

// Validate that Default is assignable to Type.

// Common reducer functions.

// DefaultReducer overwrites the existing value with the update.
func DefaultReducer(existing, update any) any {
	_ = "STUB: not implemented"
	// For composite types, return a deep copy to avoid shared references.
	return *new(any)
}

// CoverReducer overwrites the existing value with the update.
func CoverReducer(existing, update any) any {
	_ = "STUB: not implemented"

	// AppendReducer appends update to existing slice.
	return *new(any)
}

func AppendReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// Fallback to default behavior if not slices

// StringSliceReducer appends string slices specifically.
func StringSliceReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// Fallback to default behavior if not string slices

// MergeReducer merges update map into existing map.
func MergeReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// Fallback to default behavior; deep copy if composite

// OneShotMessagesByNodeReducer merges targeted one-shot inputs scoped by node
// ID. It treats the update as a partial map[nodeID][]model.Message and applies
// per-entry overrides. A nil update clears the entire map. For entries:
//   - nil / empty message slice: delete the node entry
//   - non-empty message slice: replace the node entry
func OneShotMessagesByNodeReducer(existing, update any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func applyOneShotMessagesByNodeUpdate(
	base map[string][]model.Message,
	update map[string][]model.Message,
) {
	_ = "STUB: not implemented"
	return
}

func applyOneShotMessagesByNodeAnyUpdate(
	base map[string][]model.Message,
	update map[string]any,
) {
	_ = "STUB: not implemented"
	return
}

func normalizeOneShotMessagesByNode(
	m map[string][]model.Message,
) map[string][]model.Message {
	_ = "STUB: not implemented"
	return nil
}

func decodeOneShotMessagesByNodeExisting(v any) map[string][]model.Message {
	_ = "STUB: not implemented"
	return nil
}

func decodeMapStringAny(v any) (map[string]any, bool) { _ = "STUB: not implemented"; return nil, false }

func decodeMessages(v any) ([]model.Message, error) { _ = "STUB: not implemented"; return nil, nil }

// MessageReducer handles message arrays with ID-based updates and MessageOp support.
func MessageReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// no-op

// Fallback to default behavior for unsupported types
