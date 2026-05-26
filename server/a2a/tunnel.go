//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package a2a

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
)

const defaultBatchSize = 5
const defaultFlushInterval = 200 * time.Millisecond

// eventTunnel is the event tunnel.
// It provides a way to tunnel events from agent to a2a server.
// And aggregate events into batches to improve performance.
type eventTunnel struct {
	batchSize     int
	flushInterval time.Duration

	batch []*event.Event
	// produce runs in a dedicated goroutine and must observe ctx cancellation
	// around blocking work so the producer can stop when Run returns.
	produce func(context.Context) (*event.Event, bool)
	consume func([]*event.Event) (bool, error)

	ctx    context.Context
	cancel context.CancelFunc
}

// newEventTunnel creates a new event tunnel.
// The produce callback is invoked from a dedicated goroutine. It must return
// promptly when the provided context is done to avoid leaking that goroutine.
func newEventTunnel(
	batchSize int,
	flushInterval time.Duration,
	produce func(context.Context) (*event.Event, bool),
	consume func([]*event.Event) (bool, error),
) *eventTunnel {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the event tunnel.
func (t *eventTunnel) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *eventTunnel) startProducing() <-chan *event.Event { _ = "STUB: not implemented"; return nil }

func (t *eventTunnel) handleProducedEvent(produced *event.Event) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *eventTunnel) flushPendingBatch(reason string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *eventTunnel) finalizeRun() error { _ = "STUB: not implemented"; return nil }

func (t *eventTunnel) flushBatch() (bool, error) { _ = "STUB: not implemented"; return false, nil }
