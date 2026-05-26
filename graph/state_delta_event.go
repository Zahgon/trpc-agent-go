//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graph

import (
	"context"
)

const defaultStateDeltaEventType = "state_delta"

// StateDeltaEventOptions configures custom state-delta events.
type StateDeltaEventOptions struct {
	EventType string
	Message   string
	Payload   any
}

// StateDeltaEventOption configures a custom state-delta event.
type StateDeltaEventOption func(*StateDeltaEventOptions)

// WithStateDeltaEventType sets the custom event type.
func WithStateDeltaEventType(
	eventType string,
) StateDeltaEventOption {
	_ = "STUB: not implemented"
	return *new(StateDeltaEventOption)
}

// WithStateDeltaEventMessage sets the custom event message.
func WithStateDeltaEventMessage(
	message string,
) StateDeltaEventOption {
	_ = "STUB: not implemented"
	return *new(StateDeltaEventOption)
}

// WithStateDeltaEventPayload sets the custom event payload.
func WithStateDeltaEventPayload(
	payload any,
) StateDeltaEventOption {
	_ = "STUB: not implemented"
	return *new(StateDeltaEventOption)
}

// EmitCustomStateDelta emits a node custom event carrying the given delta.
//
// This is useful when a callback or node needs to expose business state on an
// error path before the graph can emit its final graph.execution snapshot.
func EmitCustomStateDelta(
	ctx context.Context,
	state State,
	delta State,
	opts ...StateDeltaEventOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func marshalStateDelta(
	delta State,
) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeStateDeltaMaps(
	dst map[string][]byte,
	src map[string][]byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}
