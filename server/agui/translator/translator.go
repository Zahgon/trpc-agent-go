//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package translator translates trpc-agent-go events to AG-UI events.
package translator

import (
	"context"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	agentevent "trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

// Translator translates trpc-agent-go events to AG-UI events.
type Translator interface {
	// Translate translates a trpc-agent-go event to AG-UI events.
	Translate(ctx context.Context, event *agentevent.Event) ([]aguievents.Event, error)
}

// TranslatorFactory is a function that creates a translator for an AG-UI run.
type Factory func(ctx context.Context, input *adapter.RunAgentInput, opts ...Option) (Translator, error)

// NewFactory creates a default translator factory for AG-UI.
// The returned factory constructs the default translator with the provided input and options.
func NewFactory(baseOpts ...Option) Factory { _ = "STUB: not implemented"; return *new(Factory) }

// PostRunFinalizingTranslator extends Translator with post-run finalization events.
type PostRunFinalizingTranslator interface {
	Translator
	// PostRunFinalizationEvents returns AG-UI events needed to finalize open protocol streams after a run ends.
	PostRunFinalizationEvents(ctx context.Context) ([]aguievents.Event, error)
}

// New creates a new event translator.
func New(ctx context.Context, threadID, runID string, opts ...Option) (Translator, error) {
	_ = "STUB: not implemented"
	return *new(Translator), nil
}

// translator is the default implementation of the Translator.
type translator struct {
	threadID                               string
	runID                                  string
	lastMessageID                          string
	receivingMessage                       bool
	lastReasoningMessageID                 string
	receivingReasoning                     bool
	seenResponseIDs                        map[string]struct{}
	seenToolCallIDs                        map[string]struct{}
	toolCallDeltas                         map[toolCallDeltaKey]*toolCallDeltaState
	toolCallDeltasByID                     map[string]*toolCallDeltaState
	graphNodeLifecycleActivityEnabled      bool
	graphNodeInterruptActivityEnabled      bool
	graphNodeInterruptActivityTopLevelOnly bool
	reasoningContentEnabled                bool
	eventSourceMetadataEnabled             bool
	toolCallDeltaStreamingEnabled          bool
	streamingToolResultActivityEnabled     bool
	streamingToolResultContent             map[string]string
}

const skillRunArtifactsStateKey = skill.StateKeyArtifacts

// Translate translates one trpc-agent-go event into zero or more AG-UI events.
func (t *translator) Translate(ctx context.Context, event *agentevent.Event) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GraphAgent emits model/tool metadata via StateDelta instead of raw tool_calls.

// Handle node custom events (progress, text, custom).

// PostRunFinalizationEvents closes any active reasoning or text streams after a run ends.
func (t *translator) PostRunFinalizationEvents(context.Context) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *translator) finalizeEvents(
	src *agentevent.Event,
	events []aguievents.Event,
) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// A zero-value override intentionally suppresses rawEvent export.

type artifactRef struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Ref     string `json:"ref"`
}

type skillRunArtifactsDelta struct {
	ToolCallID string        `json:"tool_call_id"`
	Artifacts  []artifactRef `json:"artifacts"`
}

func (t *translator) toolArtifactsEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

const (
	graphNodeLifecycleActivityType = "graph.node.lifecycle"
	graphNodePatchPath             = "/node"
	graphNodeInterruptActivityType = "graph.node.interrupt"
	graphNodeInterruptPatchPath    = "/interrupt"
)

type graphNodePatchValue struct {
	NodeID string `json:"nodeId"`
	Phase  string `json:"phase"`
	Error  string `json:"error,omitempty"`
}

type graphNodeInterruptPatchValue struct {
	NodeID       string `json:"nodeId"`
	Key          string `json:"key,omitempty"`
	Prompt       any    `json:"prompt"`
	CheckpointID string `json:"checkpointId,omitempty"`
	LineageID    string `json:"lineageId,omitempty"`
}

func (t *translator) graphNodeActivityEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// Backward-compatible fallback for older cores that do not emit explicit
// lifecycle source metadata.

func (t *translator) graphNodeInterruptActivityEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// reasoningEvents translates reasoning_content emitted by models (e.g. DeepSeek, Claude Thinking)
// into AG-UI REASONING_* events.
func (t *translator) reasoningEvents(rsp *model.Response) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Different message ID means a new reasoning message.

// Streaming response.

// For streaming response, don't need to emit final completion event.
// It means the response is ended.

// textMessageEvent translates a text message trpc-agent-go event to AG-UI events.
func (t *translator) textMessageEvent(rsp *model.Response) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Different message ID means a new message.

// Streaming response.

// Streaming chunk.

// For streaming response, don't need to emit final completion event.
// It means the response is ended.

// toolCallEvent translates a tool call trpc-agent-go event to AG-UI events.
func (t *translator) toolCallEvent(rsp *model.Response) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toolResultEvent translates a tool result trpc-agent-go event to AG-UI events.
func (t *translator) toolResultEvent(rsp *model.Response, messageID string) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *translator) toolResultActivityEvents(rsp *model.Response) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *translator) toolResultActivityEvent(toolCallID, chunk string) (aguievents.Event, bool) {
	_ = "STUB: not implemented"
	return *new(aguievents.Event), false
}

func (t *translator) clearToolResultActivityState(rsp *model.Response) {
	_ = "STUB: not implemented"
	return
}

// formatToolCallArguments formats a tool call arguments event to a string.
func formatToolCallArguments(arguments []byte) string { _ = "STUB: not implemented"; return "" }

// graphModelEvents converts graph model metadata (from StateDelta) into text events.
func (t *translator) graphModelEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// graphToolEvents converts graph tool metadata (from StateDelta) into tool call events.
func (t *translator) graphToolEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

func (t *translator) recordResponseID(id string) { _ = "STUB: not implemented"; return }

func (t *translator) hasSeenResponseID(id string) bool { _ = "STUB: not implemented"; return false }

func (t *translator) recordToolCallID(id string) { _ = "STUB: not implemented"; return }

func (t *translator) hasSeenToolCallID(id string) bool { _ = "STUB: not implemented"; return false }

// graphNodeCustomEvents converts graph node custom metadata (from StateDelta) into AG-UI events.
// It handles three types of node custom events:
//   - Custom events: Converted to AG-UI Custom events with full payload
//   - Progress events: Converted to AG-UI Custom events with progress information
//   - Text events: Converted to TextMessageContent events if in message context,
//     otherwise converted to AG-UI Custom events
func (t *translator) graphNodeCustomEvents(evt *agentevent.Event) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// handleProgressEvent converts a progress event to AG-UI Custom events.
func (t *translator) handleProgressEvent(meta graph.NodeCustomEventMetadata) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// handleTextEvent converts a text event to AG-UI events.
// If currently receiving a message, it emits a TextMessageContent event;
// otherwise, it emits a Custom event.
func (t *translator) handleTextEvent(meta graph.NodeCustomEventMetadata) []aguievents.Event {
	_ = "STUB: not implemented"
	// If we're currently in a message context and the text is from the same
	// message context, emit as TextMessageContent for seamless streaming.
	return nil
}

// Otherwise emit as Custom event with text content.

// handleCustomEvent converts a generic custom event to AG-UI Custom events.
func (t *translator) handleCustomEvent(meta graph.NodeCustomEventMetadata) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}
