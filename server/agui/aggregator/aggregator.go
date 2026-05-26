//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package aggregator buffers and merges AG-UI events before they are persisted.
package aggregator

import (
	"context"
	"strings"
	"sync"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
)

// Aggregator buffers and merges AG-UI events before they are persisted.
type Aggregator interface {
	// Append ingests one event and returns zero or more aggregated events ready to persist.
	Append(ctx context.Context, event aguievents.Event) ([]aguievents.Event, error)
	// Flush emits any buffered events and clears internal state.
	Flush(ctx context.Context) ([]aguievents.Event, error)
}

// Factory creates a new Aggregator instance.
type Factory func(ctx context.Context, opt ...Option) Aggregator

// New creates a new aggregator with the given options.
func New(ctx context.Context, opt ...Option) Aggregator {
	_ = "STUB: not implemented"
	return *new(Aggregator)
}

// aggregator merges adjacent text, reasoning, and tool-call argument events before persistence.
type aggregator struct {
	mu       sync.Mutex
	enabled  bool            // enabled indicates whether aggregation is active.
	lastID   string          // lastID tracks the buffered message or tool call.
	lastType bufferType      // lastType tracks the event type being buffered.
	buffer   strings.Builder // buffer stores concatenated deltas for the buffered entity.
}

type bufferType int

const (
	bufferTypeUnknown bufferType = iota
	bufferTypeText
	bufferTypeReasoning
	bufferTypeToolArgs
)

// Append aggregates adjacent content events with the same message or tool call ID.
func (a *aggregator) Append(_ context.Context, event aguievents.Event) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Flush flushes any buffered text and reasoning content.
func (a *aggregator) Flush(context.Context) ([]aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleTextContent merges content when message ID matches the buffer; otherwise flushes first.
func (a *aggregator) handleTextContent(event *aguievents.TextMessageContentEvent) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

func (a *aggregator) handleReasoningContent(event *aguievents.ReasoningMessageContentEvent) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

func (a *aggregator) handleToolArgs(event *aguievents.ToolCallArgsEvent) []aguievents.Event {
	_ = "STUB: not implemented"
	return nil
}

// flush emits the buffered content as one event and clears internal state.
func (a *aggregator) flush() []aguievents.Event { _ = "STUB: not implemented"; return nil }
