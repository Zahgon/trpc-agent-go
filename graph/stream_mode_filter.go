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
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
)

type streamModeMask uint8

const (
	streamModeMaskMessages streamModeMask = 1 << iota
	streamModeMaskUpdates
	streamModeMaskCheckpoints
	streamModeMaskTasks
	streamModeMaskCustom
)

// StreamModeFilter decides whether an event should be forwarded to callers.
//
// It categorizes events into coarse groups.
// It only affects event forwarding; events are still processed internally.
type StreamModeFilter struct {
	enabled bool
	mask    streamModeMask
}

// NewStreamModeFilter builds a filter from run-level StreamMode selection.
func NewStreamModeFilter(
	enabled bool,
	modes []agent.StreamMode,
) StreamModeFilter {
	_ = "STUB: not implemented"
	return *new(StreamModeFilter)
}

func streamModeMaskFrom(modes []agent.StreamMode) streamModeMask {
	_ = "STUB: not implemented"
	return *new(streamModeMask)
}

// Allows reports whether event should be forwarded to callers.
func (f StreamModeFilter) Allows(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeErrorEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeMessageEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeUpdateEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeCheckpointEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeTaskEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func isStreamModeCustomEvent(e *event.Event) bool { _ = "STUB: not implemented"; return false }
