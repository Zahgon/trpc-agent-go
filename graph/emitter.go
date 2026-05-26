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
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
)

// EventEmitter is the interface for emitting events from within NodeFunc.
// It provides a convenient way for nodes to emit custom events, progress updates,
// and streaming text during execution.
type EventEmitter interface {
	// Emit sends a custom event to the event channel.
	// Returns an error if the event cannot be sent (e.g., channel closed or timeout).
	Emit(evt *event.Event) error

	// EmitCustom sends a custom event with the specified event type and payload.
	// The event will be automatically enriched with node context (NodeID, InvocationID, etc.).
	EmitCustom(eventType string, payload any) error

	// EmitProgress sends a progress event with the specified progress percentage and message.
	// Progress should be a value between 0 and 100.
	EmitProgress(progress float64, message string) error

	// EmitText sends a streaming text event.
	// This is useful for streaming intermediate text output from a node.
	EmitText(text string) error

	// Context returns the context associated with this emitter.
	Context() context.Context
}

// eventEmitter is the default implementation of EventEmitter.
type eventEmitter struct {
	ctx          context.Context
	eventChan    chan<- *event.Event
	nodeID       string
	invocationID string
	stepNumber   int
	branch       string
	timeout      time.Duration
}

// EventEmitterOption is a function that configures an eventEmitter.
type EventEmitterOption func(*eventEmitter)

// WithEmitterContext sets the context for the emitter.
func WithEmitterContext(ctx context.Context) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// WithEmitterNodeID sets the node ID for the emitter.
func WithEmitterNodeID(nodeID string) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// WithEmitterInvocationID sets the invocation ID for the emitter.
func WithEmitterInvocationID(invocationID string) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// WithEmitterStepNumber sets the step number for the emitter.
func WithEmitterStepNumber(stepNumber int) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// WithEmitterBranch sets the branch for the emitter.
func WithEmitterBranch(branch string) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// WithEmitterTimeout sets the timeout for emit operations.
func WithEmitterTimeout(timeout time.Duration) EventEmitterOption {
	_ = "STUB: not implemented"
	return *new(EventEmitterOption)
}

// NewEventEmitter creates a new EventEmitter with the given event channel and options.
// If eventChan is nil, returns a no-op emitter that safely ignores all emit calls.
func NewEventEmitter(eventChan chan<- *event.Event, opts ...EventEmitterOption) EventEmitter {
	_ = "STUB: not implemented"
	return *new(EventEmitter)
}

// Emit sends a custom event to the event channel.
func (e *eventEmitter) Emit(evt *event.Event) error { _ = "STUB: not implemented"; return nil }

// Inject context information if not already set

// EmitCustom sends a custom event with the specified event type and payload.
func (e *eventEmitter) EmitCustom(eventType string, payload any) error {
	_ = "STUB: not implemented"
	return nil
}

// EmitProgress sends a progress event with the specified progress percentage and message.
func (e *eventEmitter) EmitProgress(progress float64, message string) error {
	_ = "STUB: not implemented"
	// Clamp progress to 0-100
	return nil
}

// EmitText sends a streaming text event.
func (e *eventEmitter) EmitText(text string) error { _ = "STUB: not implemented"; return nil }

// Context returns the context associated with this emitter.
func (e *eventEmitter) Context() context.Context {
	_ = "STUB: not implemented"

	// emitWithRecover sends an event to the channel with panic recovery.
	return *new(context.Context)
}

func (e *eventEmitter) emitWithRecover(evt *event.Event) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Don't propagate panic as error

// noopEmitter is a no-op implementation of EventEmitter.
// It safely ignores all emit calls and is used when EventChan is unavailable.
type noopEmitter struct{}

// Emit does nothing and returns nil.
func (n *noopEmitter) Emit(evt *event.Event) error {
	_ = "STUB: not implemented"

	// EmitCustom does nothing and returns nil.
	return nil
}

func (n *noopEmitter) EmitCustom(eventType string, payload any) error {
	_ = "STUB: not implemented"

	// EmitProgress does nothing and returns nil.
	return nil
}

func (n *noopEmitter) EmitProgress(progress float64, message string) error {
	_ = "STUB: not implemented"

	// EmitText does nothing and returns nil.
	return nil
}

func (n *noopEmitter) EmitText(text string) error {
	_ = "STUB: not implemented"

	// Context returns a background context.
	return nil
}

func (n *noopEmitter) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetEventEmitter retrieves an EventEmitter from the given State.
// It extracts the ExecutionContext from the state and creates an EventEmitter
// with the appropriate context information.
// If the state does not contain a valid ExecutionContext or EventChan,
// returns a no-op emitter that safely ignores all emit calls.
func GetEventEmitter(state State) EventEmitter {
	_ = "STUB: not implemented"
	return *new(EventEmitter)
}

// GetEventEmitterWithContext retrieves an EventEmitter from the given State with a custom context.
func GetEventEmitterWithContext(ctx context.Context, state State) EventEmitter {
	_ = "STUB: not implemented"
	return *new(EventEmitter)
}

// Get ExecutionContext from state

// Check if EventChan is available

// Get current node ID from state

// Create EventEmitter with context information
