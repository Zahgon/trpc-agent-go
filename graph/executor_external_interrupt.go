//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"errors"
	"sync"
)

const (
	// ExternalInterruptKey identifies interrupts requested via
	// WithGraphInterrupt.
	ExternalInterruptKey = "external_interrupt"

	// CheckpointMetaKeyGraphInterruptInputs stores per-node input snapshots that
	// should be restored when resuming after a forced external interrupt.
	CheckpointMetaKeyGraphInterruptInputs = "graph_interrupt_inputs"
)

var errGraphInterruptTimeout = errors.New("graph interrupt timeout")

type externalInterruptWatcher struct {
	state *graphInterruptState

	stopOnce sync.Once
	stopCh   chan struct{}

	cancel context.CancelCauseFunc
}

func newExternalInterruptWatcher(
	parent context.Context,
	state *graphInterruptState,
) (context.Context, *externalInterruptWatcher) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (w *externalInterruptWatcher) listen() { _ = "STUB: not implemented"; return }

func (w *externalInterruptWatcher) stop() { _ = "STUB: not implemented"; return }

func (w *externalInterruptWatcher) requested() bool { _ = "STUB: not implemented"; return false }

func (w *externalInterruptWatcher) forced(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

type stepExecutionReport struct {
	mu sync.Mutex

	completed map[*Task]bool
	inputs    map[*Task]State
	fields    map[string]StateField
}

func newStepExecutionReport(
	fields map[string]StateField,
) *stepExecutionReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *stepExecutionReport) recordInput(task *Task, input State) {
	_ = "STUB: not implemented"
	return
}

func (r *stepExecutionReport) markCompleted(task *Task) { _ = "STUB: not implemented"; return }

func (r *stepExecutionReport) isCompleted(task *Task) bool { _ = "STUB: not implemented"; return false }

func (r *stepExecutionReport) inputFor(task *Task) (State, bool) {
	_ = "STUB: not implemented"
	return *new(State), false
}

// ExternalInterruptPayload is stored in InterruptError.Value for external
// interrupts requested via WithGraphInterrupt.
type ExternalInterruptPayload struct {
	Key    string `json:"key"`
	Forced bool   `json:"forced"`
}

func newExternalInterruptError(forced bool) *InterruptError { _ = "STUB: not implemented"; return nil }
