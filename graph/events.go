//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package graph provides graph-based workflow execution.
package graph

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph/internal/channel"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Event authors for graph-related events.
const (
	// AuthorGraphNode is the author for individual node execution events.
	AuthorGraphNode = "graph-node"
	// AuthorGraphPregel is the author for Pregel-specific events.
	AuthorGraphPregel = "graph-pregel"
)

// Event object types for graph-related events.
const (
	// ObjectTypeGraphExecution is the object type for graph execution events.
	ObjectTypeGraphExecution = "graph.execution"
	// ObjectTypeGraphBarrier is the object type for graph-level barrier events.
	ObjectTypeGraphBarrier = "graph.barrier"
	// ObjectTypeGraphNodeBarrier is the object type for node-level barrier events.
	ObjectTypeGraphNodeBarrier = "graph.node.barrier"
	// ObjectTypeGraphNodeExecution is the object type for node execution events.
	ObjectTypeGraphNodeExecution = "graph.node.execution"
	// ObjectTypeGraphNodeStart is the object type for node start events.
	ObjectTypeGraphNodeStart = "graph.node.start"
	// ObjectTypeGraphNodeComplete is the object type for node completion events.
	ObjectTypeGraphNodeComplete = "graph.node.complete"
	// ObjectTypeGraphNodeError is the object type for node error events.
	ObjectTypeGraphNodeError = "graph.node.error"
	// ObjectTypeGraphNodeCustom is the object type for node custom events emitted by NodeFunc.
	ObjectTypeGraphNodeCustom = "graph.node.custom"
	// ObjectTypeGraphPregelStep is the object type for Pregel step events.
	ObjectTypeGraphPregelStep = "graph.pregel.step"
	// ObjectTypeGraphPregelPlanning is the object type for Pregel planning events.
	ObjectTypeGraphPregelPlanning = "graph.pregel.planning"
	// ObjectTypeGraphPregelExecution is the object type for Pregel execution events.
	ObjectTypeGraphPregelExecution = "graph.pregel.execution"
	// ObjectTypeGraphPregelUpdate is the object type for Pregel update events.
	ObjectTypeGraphPregelUpdate = "graph.pregel.update"
	// ObjectTypeGraphChannelUpdate is the object type for channel update events.
	ObjectTypeGraphChannelUpdate = "graph.channel.update"
	// ObjectTypeGraphStateUpdate is the object type for state update events.
	ObjectTypeGraphStateUpdate = "graph.state.update"
	// ObjectTypeGraphCheckpoint is the object type for checkpoint events.
	ObjectTypeGraphCheckpoint = "graph.checkpoint"
	// ObjectTypeGraphCheckpointCreated is the object type for checkpoint creation events.
	ObjectTypeGraphCheckpointCreated = "graph.checkpoint.created"
	// ObjectTypeGraphCheckpointCommitted is the object type for checkpoint commit events.
	ObjectTypeGraphCheckpointCommitted = "graph.checkpoint.committed"
	// ObjectTypeGraphCheckpointInterrupt is the object type for checkpoint interrupt events.
	ObjectTypeGraphCheckpointInterrupt = "graph.checkpoint.interrupt"
)

// Metadata keys for storing event metadata in StateDelta.
const (
	// MetadataKeyNode is the key for node execution metadata.
	MetadataKeyNode = "_node_metadata"
	// MetadataKeyNodeEmitter is the key for node lifecycle emitter metadata.
	MetadataKeyNodeEmitter = "_node_emitter"
	// MetadataKeyPregel is the key for Pregel step metadata.
	MetadataKeyPregel = "_pregel_metadata"
	// MetadataKeyChannel is the key for channel update metadata.
	MetadataKeyChannel = "_channel_metadata"
	// MetadataKeyState is the key for state update metadata.
	MetadataKeyState = "_state_metadata"
	// MetadataKeyCompletion is the key for completion metadata.
	MetadataKeyCompletion = "_completion_metadata"
	// MetadataKeyTool is the key for tool execution metadata.
	MetadataKeyTool = "_tool_metadata"
	// MetadataKeyModel is the key for model execution metadata.
	MetadataKeyModel = "_model_metadata"
	// MetadataKeyCheckpoint is the key for checkpoint metadata.
	MetadataKeyCheckpoint = "_checkpoint_metadata"
	// MetadataKeyCacheHit is a synthetic key set on node completion events when
	// a cache hit occurs for the node's output.
	MetadataKeyCacheHit = "_cache_hit"
	// MetadataKeyNodeCustom is the key for node custom event metadata.
	MetadataKeyNodeCustom = "_node_custom_metadata"
)

// NodeType represents the type of a graph node.
type NodeType string

// Node type constants.
const (
	NodeTypeFunction NodeType = "function"
	NodeTypeLLM      NodeType = "llm"
	NodeTypeTool     NodeType = "tool"
	NodeTypeAgent    NodeType = "agent"
	NodeTypeJoin     NodeType = "join"
	NodeTypeRouter   NodeType = "router"
)

// String returns the string representation of the node type.
func (nt NodeType) String() string {
	_ = "STUB: not implemented"

	// NodeEventEmitter identifies which graph component emitted a node lifecycle
	// event. This is auxiliary routing metadata for downstream translators.
	return ""
}

type NodeEventEmitter string

// Node lifecycle emitter constants.
const (
	NodeEventEmitterExecutor    NodeEventEmitter = "executor"
	NodeEventEmitterAgentHelper NodeEventEmitter = "agent_helper"
)

// ExecutionPhase represents the phase of node execution.
type ExecutionPhase string

// Execution phase constants.
const (
	ExecutionPhaseStart    ExecutionPhase = "start"
	ExecutionPhaseComplete ExecutionPhase = "complete"
	ExecutionPhaseError    ExecutionPhase = "error"
)

// String returns the string representation of the execution phase.
func (ep ExecutionPhase) String() string {
	_ = "STUB: not implemented"

	// ToolExecutionPhase represents the phase of tool execution.
	return ""
}

type ToolExecutionPhase string

// Tool execution phase constants.
const (
	ToolExecutionPhaseStart    ToolExecutionPhase = "start"
	ToolExecutionPhaseComplete ToolExecutionPhase = "complete"
	ToolExecutionPhaseError    ToolExecutionPhase = "error"
)

// String returns the string representation of the tool execution phase.
func (tep ToolExecutionPhase) String() string {
	_ = "STUB: not implemented"

	// ModelExecutionPhase represents the phase of model execution.
	return ""
}

type ModelExecutionPhase string

// Model execution phase constants.
const (
	ModelExecutionPhaseStart    ModelExecutionPhase = "start"
	ModelExecutionPhaseComplete ModelExecutionPhase = "complete"
	ModelExecutionPhaseError    ModelExecutionPhase = "error"
)

// String returns the string representation of the model execution phase.
func (mep ModelExecutionPhase) String() string {
	_ = "STUB: not implemented"

	// PregelPhase represents the phase of Pregel execution.
	return ""
}

type PregelPhase string

// Pregel phase constants.
const (
	PregelPhasePlanning  PregelPhase = "planning"
	PregelPhaseExecution PregelPhase = "execution"
	PregelPhaseUpdate    PregelPhase = "update"
	PregelPhaseComplete  PregelPhase = "complete"
	PregelPhaseError     PregelPhase = "error"
)

// String returns the string representation of the Pregel phase.
func (pp PregelPhase) String() string {
	_ = "STUB: not implemented"

	// NodeExecutionMetadata contains metadata about node execution.
	return ""
}

type NodeExecutionMetadata struct {
	// NodeID is the unique identifier of the node.
	NodeID string `json:"nodeId"`
	// NodeType is the type of the node.
	NodeType NodeType `json:"nodeType"`
	// Phase is the execution phase.
	Phase ExecutionPhase `json:"phase"`
	// StartTime is when the execution started.
	StartTime time.Time `json:"startTime,omitempty"`
	// EndTime is when the execution completed.
	EndTime time.Time `json:"endTime,omitempty"`
	// Duration is the execution duration.
	Duration time.Duration `json:"duration,omitempty"`
	// InputKeys are the keys of input state.
	InputKeys []string `json:"inputKeys,omitempty"`
	// OutputKeys are the keys of output state.
	OutputKeys []string `json:"outputKeys,omitempty"`
	// Error is the error message if execution failed.
	Error string `json:"error,omitempty"`
	// ToolCalls contains tool call information for tool nodes.
	ToolCalls []model.ToolCall `json:"toolCalls,omitempty"`
	// ModelName contains the model name for LLM nodes.
	ModelName string `json:"modelName,omitempty"`
	// ModelInput contains the input sent to LLM nodes.
	ModelInput string `json:"modelInput,omitempty"`
	// StepNumber is the Pregel step number.
	StepNumber int `json:"stepNumber,omitempty"`
	// Attempt is the 1-based attempt number for this node execution.
	Attempt int `json:"attempt,omitempty"`
	// MaxAttempts is the maximum allowed attempts when retrying is enabled.
	MaxAttempts int `json:"maxAttempts,omitempty"`
	// NextDelay is the planned delay before the next retry attempt.
	NextDelay time.Duration `json:"nextDelay,omitempty"`
	// Retrying indicates whether a retry will be performed after this error.
	Retrying bool `json:"retrying,omitempty"`
}

// ToolExecutionMetadata contains metadata about tool execution.
type ToolExecutionMetadata struct {
	// ToolName is the name of the tool being executed.
	ToolName string `json:"toolName"`
	// ToolID is the unique identifier of the tool call.
	ToolID string `json:"toolId"`
	// ResponseID is the response/message ID that issued this tool call.
	ResponseID string `json:"responseId,omitempty"`
	// Phase is the execution phase.
	Phase ToolExecutionPhase `json:"phase"`
	// StartTime is when the execution started.
	StartTime time.Time `json:"startTime,omitempty"`
	// EndTime is when the execution completed.
	EndTime time.Time `json:"endTime,omitempty"`
	// Duration is the execution duration.
	Duration time.Duration `json:"duration,omitempty"`
	// Input contains the tool input arguments.
	Input string `json:"input,omitempty"`
	// Output contains the tool output result.
	Output string `json:"output,omitempty"`
	// Error is the error message if execution failed.
	Error string `json:"error,omitempty"`
	// InvocationID is the invocation ID.
	InvocationID string `json:"invocationId,omitempty"`
}

// ModelExecutionMetadata contains metadata about model execution.
type ModelExecutionMetadata struct {
	// ModelName is the name of the model being executed.
	ModelName string `json:"modelName"`
	// NodeID is the unique identifier of the node.
	NodeID string `json:"nodeId"`
	// ResponseID is the response/message ID for this model output.
	ResponseID string `json:"responseId,omitempty"`
	// Phase is the execution phase.
	Phase ModelExecutionPhase `json:"phase"`
	// StartTime is when the execution started.
	StartTime time.Time `json:"startTime,omitempty"`
	// EndTime is when the execution completed.
	EndTime time.Time `json:"endTime,omitempty"`
	// Duration is the execution duration.
	Duration time.Duration `json:"duration,omitempty"`
	// Input contains the model input (messages or prompt).
	Input string `json:"input,omitempty"`
	// Output contains the final model output result.
	Output string `json:"output,omitempty"`
	// Error is the error message if execution failed.
	Error string `json:"error,omitempty"`
	// InvocationID is the invocation ID.
	InvocationID string `json:"invocationId,omitempty"`
	// StepNumber is the Pregel step number.
	StepNumber int `json:"stepNumber,omitempty"`
}

// PregelStepMetadata contains metadata about Pregel step execution.
type PregelStepMetadata struct {
	// StepNumber is the step number.
	StepNumber int `json:"stepNumber"`
	// Phase is the Pregel phase.
	Phase PregelPhase `json:"phase"`
	// TaskCount is the number of tasks in this step.
	TaskCount int `json:"taskCount"`
	// UpdatedChannels are the channels updated in this step.
	UpdatedChannels []string `json:"updatedChannels,omitempty"`
	// ActiveNodes are the nodes active in this step.
	ActiveNodes []string `json:"activeNodes,omitempty"`
	// StartTime is when the step started.
	StartTime time.Time `json:"startTime,omitempty"`
	// EndTime is when the step completed.
	EndTime time.Time `json:"endTime,omitempty"`
	// Duration is the step duration.
	Duration time.Duration `json:"duration,omitempty"`
	// Error is the error message if step failed.
	Error string `json:"error,omitempty"`
	// NodeID is the ID of the node where interrupt occurred.
	NodeID string `json:"nodeID,omitempty"`
	// InterruptKey is the key that was passed to Interrupt().
	InterruptKey string `json:"interruptKey,omitempty"`
	// InterruptValue is the value passed to interrupt().
	InterruptValue any `json:"interruptValue,omitempty"`
	// LineageID is the lineage ID of the checkpoint that the interrupt belongs to.
	LineageID string `json:"lineageId,omitempty"`
	// CheckpointID is the checkpoint ID that can be used for resuming.
	CheckpointID string `json:"checkpointId,omitempty"`
	// CheckpointNS is the checkpoint namespace that can be used for resuming.
	CheckpointNS string `json:"checkpointNs,omitempty"`
}

// ChannelUpdateMetadata contains metadata about channel updates.
type ChannelUpdateMetadata struct {
	// ChannelName is the name of the channel.
	ChannelName string `json:"channelName"`
	// ChannelType is the type of the channel.
	ChannelType channel.Behavior `json:"channelType"`
	// ValueCount is the number of values in the channel.
	ValueCount int `json:"valueCount"`
	// Available indicates if the channel is available.
	Available bool `json:"available"`
	// TriggeredNodes are the nodes triggered by this channel.
	TriggeredNodes []string `json:"triggeredNodes,omitempty"`
}

// StateUpdateMetadata contains metadata about state updates.
type StateUpdateMetadata struct {
	// UpdatedKeys are the keys that were updated.
	UpdatedKeys []string `json:"updatedKeys"`
	// RemovedKeys are the keys that were removed.
	RemovedKeys []string `json:"removedKeys,omitempty"`
	// StateSize is the total size of the state.
	StateSize int `json:"stateSize"`
}

// JSONMetadata represents the JSON structure for metadata stored in StateDelta.
type JSONMetadata struct {
	// Node metadata for node execution events.
	Node *NodeExecutionMetadata `json:"node,omitempty"`
	// Pregel metadata for Pregel step events.
	Pregel *PregelStepMetadata `json:"pregel,omitempty"`
	// Channel metadata for channel update events.
	Channel *ChannelUpdateMetadata `json:"channel,omitempty"`
	// State metadata for state update events.
	State *StateUpdateMetadata `json:"state,omitempty"`
	// Completion metadata for completion events.
	Completion *CompletionMetadata `json:"completion,omitempty"`
	// Tool metadata for tool execution events.
	Tool *ToolExecutionMetadata `json:"tool,omitempty"`
	// Model metadata for model execution events.
	Model *ModelExecutionMetadata `json:"model,omitempty"`
}

// CompletionMetadata contains metadata about graph completion.
type CompletionMetadata struct {
	// TotalSteps is the total number of steps executed.
	TotalSteps int `json:"totalSteps"`
	// TotalDuration is the total execution duration.
	TotalDuration time.Duration `json:"totalDuration"`
	// FinalStateKeys is the number of final-state keys actually serialized
	// into the completion event StateDelta.
	FinalStateKeys int `json:"finalStateKeys"`
	// FinalResponseID carries the stable identity of the terminal assistant
	// response when the graph can provide one. When available, this is
	// typically the underlying model response ID.
	FinalResponseID string `json:"finalResponseID,omitempty"`
	// SnapshotOnly marks terminal completion snapshots that should remain
	// available to callers but must not be replayed as conversational history.
	SnapshotOnly bool `json:"snapshotOnly,omitempty"`
}

// NodeCustomEventCategory represents the category of node custom events.
type NodeCustomEventCategory string

// Node custom event category constants.
const (
	// NodeCustomEventCategoryCustom is the category for general custom events.
	NodeCustomEventCategoryCustom NodeCustomEventCategory = "custom"
	// NodeCustomEventCategoryProgress is the category for progress events.
	NodeCustomEventCategoryProgress NodeCustomEventCategory = "progress"
	// NodeCustomEventCategoryText is the category for streaming text events.
	NodeCustomEventCategoryText NodeCustomEventCategory = "text"
)

// String returns the string representation of the node custom event category.
func (c NodeCustomEventCategory) String() string {
	_ = "STUB: not implemented"

	// NodeCustomEventMetadata contains metadata about node custom events.
	return ""
}

type NodeCustomEventMetadata struct {
	// EventType is the user-defined event type.
	EventType string `json:"eventType"`
	// Category is the category of the custom event (custom, progress, text).
	Category NodeCustomEventCategory `json:"category"`
	// NodeID is the ID of the node that emitted the event.
	NodeID string `json:"nodeId"`
	// InvocationID is the invocation ID of the current execution.
	InvocationID string `json:"invocationId"`
	// StepNumber is the Pregel step number when the event was emitted.
	StepNumber int `json:"stepNumber,omitempty"`
	// Timestamp is when the event was emitted.
	Timestamp time.Time `json:"timestamp"`
	// Payload is the custom payload data.
	Payload any `json:"payload,omitempty"`
	// Progress is the progress percentage (0-100) for progress events.
	Progress float64 `json:"progress,omitempty"`
	// Message is the message for progress events or text content for text events.
	Message string `json:"message,omitempty"`
}

// EventOption is a function that configures a graph event.
type EventOption func(*event.Event)

// WithNodeMetadata adds node execution metadata to the event.
func WithNodeMetadata(metadata NodeExecutionMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// NodeEventEmitterFromStateDelta extracts the node lifecycle emitter marker
// from an event state delta when one is present.
func NodeEventEmitterFromStateDelta(stateDelta map[string][]byte) NodeEventEmitter {
	_ = "STUB: not implemented"
	return *new(NodeEventEmitter)
}

// SetNodeEventEmitterInStateDelta stores a node lifecycle emitter marker in
// state delta metadata. Empty emitters are ignored.
func SetNodeEventEmitterInStateDelta(stateDelta map[string][]byte, emitter NodeEventEmitter) {
	_ = "STUB: not implemented"
	return
}

// WithToolMetadata adds tool execution metadata to the event.
func WithToolMetadata(metadata ToolExecutionMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// WithModelMetadata adds model execution metadata to the event.
func WithModelMetadata(metadata ModelExecutionMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// WithPregelMetadata adds Pregel step metadata to the event.
func WithPregelMetadata(metadata PregelStepMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// WithChannelMetadata adds channel update metadata to the event.
func WithChannelMetadata(metadata ChannelUpdateMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// WithStateMetadata adds state update metadata to the event.
func WithStateMetadata(metadata StateUpdateMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// WithNodeCustomMetadata adds node custom event metadata to the event.
func WithNodeCustomMetadata(metadata NodeCustomEventMetadata) EventOption {
	_ = "STUB: not implemented"
	return *new(EventOption)
}

// Store metadata in StateDelta as JSON.

// Marshal metadata to JSON.

// NewGraphEvent creates a new graph-related event.
func NewGraphEvent(invocationID, author, objectType string, opts ...EventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// formatNodeAuthor returns nodeID if non-empty; otherwise returns fallback.
func formatNodeAuthor(nodeID, fallbackAuthor string) string { _ = "STUB: not implemented"; return "" }

// NodeEventOptions contains options for creating node events.
type NodeEventOptions struct {
	InvocationID  string
	NodeID        string
	NodeType      NodeType
	Emitter       NodeEventEmitter
	StepNumber    int
	StartTime     time.Time
	EndTime       time.Time
	InputKeys     []string
	OutputKeys    []string
	ToolCalls     []model.ToolCall
	ModelName     string
	ModelInput    string
	Error         string
	ResponseError *model.ResponseError
	// Retry metadata (optional)
	Attempt     int
	MaxAttempts int
	NextDelay   time.Duration
	Retrying    bool
}

// NodeEventOption is a function that configures node event options.
type NodeEventOption func(*NodeEventOptions)

// ToolEventOptions contains options for creating tool events.
type ToolEventOptions struct {
	InvocationID string
	ToolName     string
	ToolID       string
	ResponseID   string
	Phase        ToolExecutionPhase
	StartTime    time.Time
	EndTime      time.Time
	Input        string
	Output       string
	Error        error
	// NodeID is optional. When provided, author becomes node-scoped.
	NodeID string
	// IncludeResponse controls whether NewToolExecutionEvent should attach a
	// tool.response payload in addition to metadata.
	IncludeResponse bool
}

// ToolEventOption is a function that configures tool event options.
type ToolEventOption func(*ToolEventOptions)

// WithToolEventNodeID sets the node ID for tool events.
func WithToolEventNodeID(nodeID string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// ModelEventOptions contains options for creating model events.
type ModelEventOptions struct {
	InvocationID string
	ModelName    string
	NodeID       string
	ResponseID   string
	Phase        ModelExecutionPhase
	StartTime    time.Time
	EndTime      time.Time
	Input        string
	Output       string
	Error        error
	StepNumber   int
}

// ModelEventOption is a function that configures model event options.
type ModelEventOption func(*ModelEventOptions)

// WithNodeEventInvocationID sets the invocation ID for node events.
func WithNodeEventInvocationID(invocationID string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventNodeID sets the node ID for node events.
func WithNodeEventNodeID(nodeID string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventNodeType sets the node type for node events.
func WithNodeEventNodeType(nodeType NodeType) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventEmitter sets the lifecycle emitter/source for node events.
func WithNodeEventEmitter(emitter NodeEventEmitter) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventStepNumber sets the step number for node events.
func WithNodeEventStepNumber(stepNumber int) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventStartTime sets the start time for node events.
func WithNodeEventStartTime(startTime time.Time) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventEndTime sets the end time for node events.
func WithNodeEventEndTime(endTime time.Time) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventInputKeys sets the input keys for node events.
func WithNodeEventInputKeys(inputKeys []string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventOutputKeys sets the output keys for node events.
func WithNodeEventOutputKeys(outputKeys []string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventToolCalls sets the tool calls for node events.
func WithNodeEventToolCalls(toolCalls []model.ToolCall) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventModelName sets the model name for node events.
func WithNodeEventModelName(modelName string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventModelInput sets the model input for node events.
func WithNodeEventModelInput(modelInput string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventError sets the error message for node events.
func WithNodeEventError(errMsg string) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventResponseError sets ResponseError for node events.
func WithNodeEventResponseError(err *model.ResponseError) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventAttempt sets the current attempt number (1-based).
func WithNodeEventAttempt(attempt int) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventMaxAttempts sets the maximum attempts.
func WithNodeEventMaxAttempts(maxAttempts int) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventNextDelay sets the planned next delay before retry.
func WithNodeEventNextDelay(delay time.Duration) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithNodeEventRetrying indicates whether a retry will be performed.
func WithNodeEventRetrying(retrying bool) NodeEventOption {
	_ = "STUB: not implemented"
	return *new(NodeEventOption)
}

// WithToolEventInvocationID sets the invocation ID for tool events.
func WithToolEventInvocationID(invocationID string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventToolName sets the tool name for tool events.
func WithToolEventToolName(toolName string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventToolID sets the tool ID for tool events.
func WithToolEventToolID(toolID string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventResponseID sets the parent response ID for tool events.
func WithToolEventResponseID(responseID string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventPhase sets the phase for tool events.
func WithToolEventPhase(phase ToolExecutionPhase) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventStartTime sets the start time for tool events.
func WithToolEventStartTime(startTime time.Time) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventEndTime sets the end time for tool events.
func WithToolEventEndTime(endTime time.Time) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventInput sets the input for tool events.
func WithToolEventInput(input string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventOutput sets the output for tool events.
func WithToolEventOutput(output string) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventError sets the error for tool events.
func WithToolEventError(err error) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithToolEventIncludeResponse controls whether the resulting event should
// embed a tool.response payload in addition to metadata.
func WithToolEventIncludeResponse(include bool) ToolEventOption {
	_ = "STUB: not implemented"
	return *new(ToolEventOption)
}

// WithModelEventResponseID sets the response ID for model events.
func WithModelEventResponseID(responseID string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventInvocationID sets the invocation ID for model events.
func WithModelEventInvocationID(invocationID string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventModelName sets the model name for model events.
func WithModelEventModelName(modelName string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventNodeID sets the node ID for model events.
func WithModelEventNodeID(nodeID string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventPhase sets the phase for model events.
func WithModelEventPhase(phase ModelExecutionPhase) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventStartTime sets the start time for model events.
func WithModelEventStartTime(startTime time.Time) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventEndTime sets the end time for model events.
func WithModelEventEndTime(endTime time.Time) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventInput sets the input for model events.
func WithModelEventInput(input string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventOutput sets the output for model events.
func WithModelEventOutput(output string) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventError sets the error for model events.
func WithModelEventError(err error) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// WithModelEventStepNumber sets the step number for model events.
func WithModelEventStepNumber(stepNumber int) ModelEventOption {
	_ = "STUB: not implemented"
	return *new(ModelEventOption)
}

// PregelEventOptions contains options for creating Pregel events.
type PregelEventOptions struct {
	InvocationID    string
	StepNumber      int
	Phase           PregelPhase
	TaskCount       int
	UpdatedChannels []string
	ActiveNodes     []string
	StartTime       time.Time
	EndTime         time.Time
	Error           string
	ResponseError   *model.ResponseError
	NodeID          string
	InterruptKey    string
	InterruptValue  any
	LineageID       string
	CheckpointID    string
	CheckpointNS    string
}

// PregelEventOption is a function that configures Pregel event options.
type PregelEventOption func(*PregelEventOptions)

// WithPregelEventInvocationID sets the invocation ID for Pregel events.
func WithPregelEventInvocationID(invocationID string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventStepNumber sets the step number for Pregel events.
func WithPregelEventStepNumber(stepNumber int) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventPhase sets the phase for Pregel events.
func WithPregelEventPhase(phase PregelPhase) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventTaskCount sets the task count for Pregel events.
func WithPregelEventTaskCount(taskCount int) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventUpdatedChannels sets the updated channels for Pregel events.
func WithPregelEventUpdatedChannels(updatedChannels []string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventActiveNodes sets the active nodes for Pregel events.
func WithPregelEventActiveNodes(activeNodes []string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventStartTime sets the start time for Pregel events.
func WithPregelEventStartTime(startTime time.Time) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventEndTime sets the end time for Pregel events.
func WithPregelEventEndTime(endTime time.Time) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventError sets the error message for Pregel events.
func WithPregelEventError(errMsg string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventResponseError sets the structured ResponseError for Pregel
// events.
func WithPregelEventResponseError(err *model.ResponseError) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventNodeID sets the node ID for Pregel events.
func WithPregelEventNodeID(nodeID string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventInterruptKey sets the interrupt key for Pregel events.
func WithPregelEventInterruptKey(key string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventInterruptValue sets the interrupt value for Pregel events.
func WithPregelEventInterruptValue(value any) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventLineageID sets the lineage ID for Pregel events.
func WithPregelEventLineageID(lineageID string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventCheckpointID sets the checkpoint ID for Pregel events.
func WithPregelEventCheckpointID(checkpointID string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// WithPregelEventCheckpointNS sets the checkpoint namespace for Pregel events.
func WithPregelEventCheckpointNS(checkpointNS string) PregelEventOption {
	_ = "STUB: not implemented"
	return *new(PregelEventOption)
}

// ChannelEventOptions contains options for creating channel events.
type ChannelEventOptions struct {
	InvocationID   string
	ChannelName    string
	ChannelType    channel.Behavior
	ValueCount     int
	Available      bool
	TriggeredNodes []string
}

// ChannelEventOption is a function that configures channel event options.
type ChannelEventOption func(*ChannelEventOptions)

// WithChannelEventInvocationID sets the invocation ID for channel events.
func WithChannelEventInvocationID(invocationID string) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// WithChannelEventChannelName sets the channel name for channel events.
func WithChannelEventChannelName(channelName string) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// WithChannelEventChannelType sets the channel type for channel events.
func WithChannelEventChannelType(channelType channel.Behavior) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// WithChannelEventValueCount sets the value count for channel events.
func WithChannelEventValueCount(valueCount int) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// WithChannelEventAvailable sets the availability for channel events.
func WithChannelEventAvailable(available bool) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// WithChannelEventTriggeredNodes sets the triggered nodes for channel events.
func WithChannelEventTriggeredNodes(triggeredNodes []string) ChannelEventOption {
	_ = "STUB: not implemented"
	return *new(ChannelEventOption)
}

// StateEventOptions contains options for creating state events.
type StateEventOptions struct {
	InvocationID string
	UpdatedKeys  []string
	RemovedKeys  []string
	StateSize    int
}

// StateEventOption is a function that configures state event options.
type StateEventOption func(*StateEventOptions)

// WithStateEventInvocationID sets the invocation ID for state events.
func WithStateEventInvocationID(invocationID string) StateEventOption {
	_ = "STUB: not implemented"
	return *new(StateEventOption)
}

// WithStateEventUpdatedKeys sets the updated keys for state events.
func WithStateEventUpdatedKeys(updatedKeys []string) StateEventOption {
	_ = "STUB: not implemented"
	return *new(StateEventOption)
}

// WithStateEventRemovedKeys sets the removed keys for state events.
func WithStateEventRemovedKeys(removedKeys []string) StateEventOption {
	_ = "STUB: not implemented"
	return *new(StateEventOption)
}

// WithStateEventStateSize sets the state size for state events.
func WithStateEventStateSize(stateSize int) StateEventOption {
	_ = "STUB: not implemented"
	return *new(StateEventOption)
}

// CompletionEventOptions contains options for creating completion events.
type CompletionEventOptions struct {
	InvocationID    string
	FinalState      State
	FinalResponseID string
	TotalSteps      int
	TotalDuration   time.Duration
	SnapshotOnly    bool
}

// CompletionEventOption is a function that configures completion event options.
type CompletionEventOption func(*CompletionEventOptions)

// WithCompletionEventInvocationID sets the invocation ID for completion events.
func WithCompletionEventInvocationID(invocationID string) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// WithCompletionEventFinalState sets the final state for completion events.
func WithCompletionEventFinalState(finalState State) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// WithCompletionEventFinalResponseID sets the terminal response ID for
// completion events when one is available.
func WithCompletionEventFinalResponseID(responseID string) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// WithCompletionEventTotalSteps sets the total steps for completion events.
func WithCompletionEventTotalSteps(totalSteps int) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// WithCompletionEventTotalDuration sets the total duration for completion events.
func WithCompletionEventTotalDuration(totalDuration time.Duration) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// WithCompletionEventSnapshotOnly marks completion output as terminal
// snapshot content that should not be replayed as conversational history.
func WithCompletionEventSnapshotOnly(snapshotOnly bool) CompletionEventOption {
	_ = "STUB: not implemented"
	return *new(CompletionEventOption)
}

// NewNodeStartEvent creates a new node start event.
func NewNodeStartEvent(opts ...NodeEventOption) *event.Event { _ = "STUB: not implemented"; return nil }

// NewNodeCompleteEvent creates a new node completion event.
func NewNodeCompleteEvent(opts ...NodeEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewNodeErrorEvent creates a new node error event.
func NewNodeErrorEvent(opts ...NodeEventOption) *event.Event { _ = "STUB: not implemented"; return nil }

// NewToolExecutionEvent creates a new tool execution event.
func NewToolExecutionEvent(opts ...ToolEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewModelExecutionEvent creates a new model execution event.
func NewModelExecutionEvent(opts ...ModelEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewPregelStepEvent creates a new Pregel step event.
func NewPregelStepEvent(opts ...PregelEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewPregelErrorEvent creates a new Pregel error event.
func NewPregelErrorEvent(opts ...PregelEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Build base graph event with metadata.

// Mirror error to Event.Error for easier consumption by clients that
// only check event.Error, while keeping object as graph.pregel.step
// for compatibility with existing consumers.

// NewPregelInterruptEvent creates a new Pregel interrupt event.
func NewPregelInterruptEvent(opts ...PregelEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewChannelUpdateEvent creates a new channel update event.
func NewChannelUpdateEvent(opts ...ChannelEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewStateUpdateEvent creates a new state update event.
func NewStateUpdateEvent(opts ...StateEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewGraphCompletionEvent creates a new graph completion event.
func NewGraphCompletionEvent(opts ...CompletionEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Extract final response from state if available

// Always initialize StateDelta to a non-nil map to ensure consumers can rely on it.

// Also include a serialized snapshot of the final state itself so downstream
// consumers (including tests) can reconstruct state without additional logic.

// Add completion metadata to StateDelta after serializing the final state so
// FinalStateKeys reflects the keys that were actually written.

// ensureStateDelta initializes StateDelta if nil.
func ensureStateDelta(e *event.Event) { _ = "STUB: not implemented"; return }

// extractFinalResponse fetches the last response text from state.
func extractFinalResponse(state State) string { _ = "STUB: not implemented"; return "" }

// buildFinalChoices constructs the terminal assistant message choice.
func buildFinalChoices(text string) []model.Choice { _ = "STUB: not implemented"; return nil }

// addCompletionMetadata attaches completion metadata to StateDelta.
func addCompletionMetadata(e *event.Event, options *CompletionEventOptions, finalStateKeys int) {
	_ = "STUB: not implemented"
	return
}

// FinalResponseIDFromStateDelta returns the stable terminal response identity
// carried by a graph completion snapshot when one is available.
func FinalResponseIDFromStateDelta(stateDelta map[string][]byte) string {
	_ = "STUB: not implemented"
	return ""
}

// CompletionSnapshotOnlyFromStateDelta reports whether a completion snapshot
// should be excluded from conversational history replay.
func CompletionSnapshotOnlyFromStateDelta(stateDelta map[string][]byte) bool {
	_ = "STUB: not implemented"
	return false
}

// SetCompletionSnapshotOnlyInStateDelta updates completion metadata to mark the
// snapshot as replay-ineligible conversational history.
func SetCompletionSnapshotOnlyInStateDelta(stateDelta map[string][]byte, snapshotOnly bool) {
	_ = "STUB: not implemented"
	return
}

// serializeFinalState writes serializable final state keys into StateDelta and
// returns the number of keys that were actually written.
func serializeFinalState(e *event.Event, state State) int { _ = "STUB: not implemented"; return 0 }

// Skip internal/ephemeral keys that are not JSON-serializable or can race
// due to concurrent updates (e.g., execution context and callbacks).

// Marshal a deep-copied snapshot to avoid racing on shared references.

// Special case: when users put JSON bytes into graph state (e.g.,
// json.Marshal output), encoding/json would base64 it if we marshal the
// []byte again. If it's already valid JSON, keep it as-is so downstream
// consumers can json.Unmarshal it directly.

// extractStateKeys extracts all keys from a state map.
func extractStateKeys(state State) []string { _ = "STUB: not implemented"; return nil }

// Create a copy of the state to avoid concurrent access issues

// CheckpointEventOptions contains options for creating checkpoint events.
type CheckpointEventOptions struct {
	InvocationID   string
	CheckpointID   string
	Source         string
	Step           int
	Duration       time.Duration
	Bytes          int64
	WritesCount    int
	ResumeReplay   bool
	InterruptValue any
}

// CheckpointEventOption is a function that configures checkpoint event options.
type CheckpointEventOption func(*CheckpointEventOptions)

// WithCheckpointEventInvocationID sets the invocation ID.
func WithCheckpointEventInvocationID(invocationID string) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventCheckpointID sets the checkpoint ID.
func WithCheckpointEventCheckpointID(checkpointID string) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventSource sets the checkpoint source.
func WithCheckpointEventSource(source string) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventStep sets the step number.
func WithCheckpointEventStep(step int) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventDuration sets the duration.
func WithCheckpointEventDuration(duration time.Duration) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventBytes sets the bytes written.
func WithCheckpointEventBytes(bytes int64) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventWritesCount sets the writes count.
func WithCheckpointEventWritesCount(count int) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventResumeReplay sets the resume replay flag.
func WithCheckpointEventResumeReplay(replay bool) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// WithCheckpointEventInterruptValue sets the interrupt value.
func WithCheckpointEventInterruptValue(value any) CheckpointEventOption {
	_ = "STUB: not implemented"
	return *new(CheckpointEventOption)
}

// NewCheckpointCreatedEvent creates a new checkpoint created event.
func NewCheckpointCreatedEvent(opts ...CheckpointEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewCheckpointCommittedEvent creates a new checkpoint committed event.
func NewCheckpointCommittedEvent(opts ...CheckpointEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewCheckpointInterruptEvent creates a new checkpoint interrupt event.
func NewCheckpointInterruptEvent(opts ...CheckpointEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NodeCustomEventOptions contains options for creating node custom events.
type NodeCustomEventOptions struct {
	InvocationID string
	NodeID       string
	EventType    string
	Category     NodeCustomEventCategory
	StepNumber   int
	Payload      any
	Progress     float64
	Message      string
	Branch       string
}

// NodeCustomEventOption is a function that configures node custom event options.
type NodeCustomEventOption func(*NodeCustomEventOptions)

// WithNodeCustomEventInvocationID sets the invocation ID for node custom events.
func WithNodeCustomEventInvocationID(invocationID string) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventNodeID sets the node ID for node custom events.
func WithNodeCustomEventNodeID(nodeID string) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventEventType sets the event type for node custom events.
func WithNodeCustomEventEventType(eventType string) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventCategory sets the category for node custom events.
func WithNodeCustomEventCategory(category NodeCustomEventCategory) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventStepNumber sets the step number for node custom events.
func WithNodeCustomEventStepNumber(stepNumber int) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventPayload sets the payload for node custom events.
func WithNodeCustomEventPayload(payload any) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventProgress sets the progress for node custom events.
func WithNodeCustomEventProgress(progress float64) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventMessage sets the message for node custom events.
func WithNodeCustomEventMessage(message string) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// WithNodeCustomEventBranch sets the branch for node custom events.
func WithNodeCustomEventBranch(branch string) NodeCustomEventOption {
	_ = "STUB: not implemented"
	return *new(NodeCustomEventOption)
}

// NewNodeCustomEvent creates a new node custom event.
// This function is used for creating general custom events emitted by NodeFunc.
func NewNodeCustomEvent(opts ...NodeCustomEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// NewNodeProgressEvent creates a new progress event for node execution.
// Progress should be a value between 0 and 100.
func NewNodeProgressEvent(opts ...NodeCustomEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Ensure category and event type are set for progress events

// Clamp progress to 0-100

// NewNodeTextEvent creates a new streaming text event for node execution.
// This is useful for streaming intermediate text output from a node.
func NewNodeTextEvent(opts ...NodeCustomEventOption) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Ensure category and event type are set for text events
