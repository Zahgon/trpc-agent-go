//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package source provides shared helpers for AG-UI source metadata.
package source

import (
	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	agentevent "trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	// ExtensionKey stores an AG-UI source metadata override on an agent event.
	ExtensionKey = "server.agui.source_metadata.v1"
)

// Metadata is the compact source metadata exposed to AG-UI consumers.
type Metadata struct {
	EventID            string `json:"eventId,omitempty"`
	Author             string `json:"author,omitempty"`
	InvocationID       string `json:"invocationId,omitempty"`
	ParentInvocationID string `json:"parentInvocationId,omitempty"`
	Branch             string `json:"branch,omitempty"`
}

// SnapshotMetadata indexes source metadata for messages snapshot payloads.
type SnapshotMetadata struct {
	Messages  map[string]Metadata `json:"messages,omitempty"`
	ToolCalls map[string]Metadata `json:"toolCalls,omitempty"`
}

// IsZero reports whether the metadata is empty.
func (m Metadata) IsZero() bool { _ = "STUB: not implemented"; return false }

// IsZero reports whether the snapshot metadata is empty.
func (m SnapshotMetadata) IsZero() bool { _ = "STUB: not implemented"; return false }

// FromEvent resolves source metadata from an agent event.
//
// If the event carries an explicit override in Extensions, the override takes
// precedence. A zero-value override intentionally suppresses metadata export.
func FromEvent(ev *agentevent.Event) (Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(Metadata), false
}

// SetEventOverride stores an explicit source metadata override on the event.
func SetEventOverride(
	ev *agentevent.Event,
	metadata Metadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

// FromRawEvent extracts source metadata from an AG-UI rawEvent payload.
func FromRawEvent(raw any) (Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(Metadata), false
}

// BuildSnapshotMetadata derives message and tool-call source indexes from
// persisted AG-UI track events.
func BuildSnapshotMetadata(events []session.TrackEvent) SnapshotMetadata {
	_ = "STUB: not implemented"
	return *new(SnapshotMetadata)
}

func fromEventOverride(ev *agentevent.Event) (Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(Metadata), false
}

func stringFromMap(values map[string]any, key string) string { _ = "STUB: not implemented"; return "" }

func recordSnapshotMetadata(
	messages map[string]Metadata,
	toolCalls map[string]Metadata,
	event aguievents.Event,
	metadata Metadata,
) {
	_ = "STUB: not implemented"
	return
}
