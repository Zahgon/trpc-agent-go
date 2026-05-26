//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package steer provides invocation-scoped queued user-message storage.
// Runner uses it to accept steer messages for an active run, while llmflow
// drains and persists them only at safe loop boundaries.
package steer

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// StateKeyQueuedUserMessages is the invocation state key used by Attach.
const StateKeyQueuedUserMessages = "__queued_user_messages__"

// Queue stores queued user messages in FIFO order.
type Queue struct {
	mu       sync.Mutex
	messages []model.Message
	closed   bool
}

// NewQueue creates an empty queue.
func NewQueue() *Queue {
	_ = "STUB: not implemented"

	// Enqueue appends one message unless the queue has been closed.
	return nil
}

func (q *Queue) Enqueue(message model.Message) bool { _ = "STUB: not implemented"; return false }

// Drain returns all queued messages in FIFO order.
func (q *Queue) Drain() []model.Message { _ = "STUB: not implemented"; return nil }

// Close rejects future enqueues.
func (q *Queue) Close() { _ = "STUB: not implemented"; return }

// Attach binds a queue to the invocation.
func Attach(inv *agent.Invocation, queue *Queue) { _ = "STUB: not implemented"; return }

// IsAttached reports whether a queue is attached to the invocation.
func IsAttached(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

// Drain removes and returns queued messages from the invocation.
func Drain(inv *agent.Invocation) []model.Message { _ = "STUB: not implemented"; return nil }

// Close rejects future enqueues for the invocation queue.
func Close(inv *agent.Invocation) { _ = "STUB: not implemented"; return }

// Clear closes the queue and removes it from the invocation.
func Clear(inv *agent.Invocation) { _ = "STUB: not implemented"; return }
