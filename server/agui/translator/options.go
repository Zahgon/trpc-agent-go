//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package translator

// options configures which AG-UI events the translator emits.
type options struct {
	graphNodeLifecycleActivityEnabled      bool // graphNodeLifecycleActivityEnabled enables graph node lifecycle activity events.
	graphNodeInterruptActivityEnabled      bool // graphNodeInterruptActivityEnabled enables graph interrupt activity events.
	graphNodeInterruptActivityTopLevelOnly bool // graphNodeInterruptActivityTopLevelOnly drops nested graph interrupt activity events.
	reasoningContentEnabled                bool // reasoningContentEnabled controls whether reasoning content events are emitted.
	eventSourceMetadataEnabled             bool // eventSourceMetadataEnabled attaches trpc-agent-go source metadata to translated AG-UI events.
	toolCallDeltaStreamingEnabled          bool // toolCallDeltaStreamingEnabled streams partial tool-call arguments.
	streamingToolResultActivityEnabled     bool // streamingToolResultActivityEnabled rewrites partial tool results as activity events.
}

// Option is a function that configures the options.
type Option func(*options)

// newOptions creates a new options instance.
func newOptions(opt ...Option) options { _ = "STUB: not implemented"; return *new(options) }

// WithGraphNodeLifecycleActivityEnabled controls whether the translator emits
// ACTIVITY_DELTA events with activityType "graph.node.lifecycle".
func WithGraphNodeLifecycleActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphNodeInterruptActivityEnabled controls whether the translator emits
// ACTIVITY_DELTA events with activityType "graph.node.interrupt".
func WithGraphNodeInterruptActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphNodeInterruptActivityTopLevelOnly controls whether the translator only emits
// graph interrupt activity events for the top-level invocation.
func WithGraphNodeInterruptActivityTopLevelOnly(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithReasoningContentEnabled controls whether the translator emits REASONING_* events.
func WithReasoningContentEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEventSourceMetadataEnabled controls whether the translator attaches
// source metadata from the original trpc-agent-go event to each translated
// AG-UI event via the AG-UI event's rawEvent field.
func WithEventSourceMetadataEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToolCallDeltaStreamingEnabled controls whether the translator emits
// streamed tool call events from chat completion tool-call deltas.
func WithToolCallDeltaStreamingEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamingToolResultActivityEnabled controls whether the translator rewrites
// partial tool-result chunks into activity events and leaves only the final
// tool result on the tool-result path.
func WithStreamingToolResultActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
