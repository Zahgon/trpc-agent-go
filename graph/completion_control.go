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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// WithGraphCompletionCapture keeps terminal graph completion events available
// to internal graph consumers even when caller-visible forwarding is disabled.
func WithGraphCompletionCapture(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithoutGraphCompletionCapture clears any inherited capture flag for the
// current visible stream while preserving the rest of the context.
func WithoutGraphCompletionCapture(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ShouldCaptureGraphCompletion reports whether the current context keeps
// terminal graph completion events available for internal consumers.
func ShouldCaptureGraphCompletion(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldCaptureGraphCompletion(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// IsGraphCompletionEvent reports whether the event is a terminal
// graph.execution event.
func IsGraphCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func isGraphCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// IsVisibleGraphCompletionEvent reports whether the event is a caller-visible
// response rewritten from a terminal graph completion event.
func IsVisibleGraphCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func isVisibleGraphCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// VisibleGraphCompletionEvent rewrites a terminal graph completion event into a
// caller-visible response event while preserving the final state delta.
func VisibleGraphCompletionEvent(evt *event.Event) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// VisibleGraphCompletionEventForAuthor rewrites a terminal graph completion
// event into a caller-visible response event while restoring the caller-visible
// author when one is provided.
func VisibleGraphCompletionEventForAuthor(
	evt *event.Event,
	author string,
) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// RecordAssistantResponseID stores stable dedup identifiers when the event
// contains a non-partial assistant message so visible completion snapshots can
// avoid re-emitting the same final answer text.
func RecordAssistantResponseID(
	emitted map[string]struct{},
	evt *event.Event,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// VisibleGraphCompletionEventWithDedup rewrites a terminal graph completion
// event into a caller-visible response event and clears duplicated final
// choices when the corresponding assistant response was already emitted.
func VisibleGraphCompletionEventWithDedup(
	evt *event.Event,
	emittedAssistantResponseIDs map[string]struct{},
) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// VisibleGraphCompletionEventWithDedupForAuthor rewrites a terminal graph
// completion event into a caller-visible response event and clears duplicated
// final choices when the corresponding assistant response was already emitted.
func VisibleGraphCompletionEventWithDedupForAuthor(
	evt *event.Event,
	emittedAssistantResponseIDs map[string]struct{},
	author string,
) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// VisibleGraphCompletionEventsForForwarding returns the caller-visible event to
// emit and the full completion snapshot to preserve for callbacks.
func VisibleGraphCompletionEventsForForwarding(
	evt *event.Event,
	emittedAssistantResponseIDs map[string]struct{},
) (*event.Event, *event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// VisibleGraphCompletionEventsForForwardingWithAuthor returns the caller-visible
// event to emit and the full completion snapshot to preserve for callbacks.
func VisibleGraphCompletionEventsForForwardingWithAuthor(
	evt *event.Event,
	emittedAssistantResponseIDs map[string]struct{},
	author string,
) (*event.Event, *event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// ShouldSuppressGraphCompletionEvent reports whether the caller-visible stream
// should hide the terminal graph completion event for this invocation.
func ShouldSuppressGraphCompletionEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	evt *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldClearVisibleGraphCompletionChoices(
	evt *event.Event,
	emittedAssistantResponseIDs map[string]struct{},
) bool {
	_ = "STUB: not implemented"
	return false
}

func visibleGraphCompletionNeedsFullResponseSnapshot(
	raw *event.Event,
	visible *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func assistantChoiceSignature(choices []model.Choice) string { _ = "STUB: not implemented"; return "" }
