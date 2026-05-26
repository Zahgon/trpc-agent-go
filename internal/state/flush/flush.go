//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package flush provides internal utilities for coordinating session flush requests between the runner and dependent
// components (e.g., AgentTool).It stores a flush function on the Invocation state so that callers can request a flush.
package flush

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// StateKeyFlushSession is the invocation state key used by flush.Attach.
const StateKeyFlushSession = "__flush_session__"

// FlushRequest represents a single session flush request.
// The runner will close ACK when it has appended all events that were enqueued before this request was processed.
type FlushRequest struct {
	ACK chan struct{} // ACK is closed by the runner when the flush is complete.
}

// Flusher defines a flush function used to synchronize session events.
type flusher func(context.Context) error

// flusherHolder holds a shared flusher so Clear can invalidate it for all clones.
type flusherHolder struct {
	mu sync.RWMutex
	fn flusher
}

func (h *flusherHolder) set(fn flusher) { _ = "STUB: not implemented"; return }

func (h *flusherHolder) get() flusher { _ = "STUB: not implemented"; return *new(flusher) }

func (h *flusherHolder) clear() { _ = "STUB: not implemented"; return }

// IsAttached reports whether a flush function has been attached to the invocation.
func IsAttached(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

// Attach binds a flush function to the given invocation and wires it to the provided flush channel.
// When Invoke is called, the function will enqueue a FlushRequest on ch and wait for ACK to be closed by the runner.
func Attach(ctx context.Context, inv *agent.Invocation, ch chan *FlushRequest) {
	_ = "STUB: not implemented"
	return
}

// Enqueue the flush request on the flush channel.

// Wait for the ACK to be closed by the runner.

// Reuse existing holder if present; otherwise create one.

// Invoke executes the flush function stored on the invocation state if present.
func Invoke(ctx context.Context, inv *agent.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

// Clear removes any flush function stored on the invocation state.
// This is intended to be called by the runner when the event loop finishes.
func Clear(inv *agent.Invocation) { _ = "STUB: not implemented"; return }
