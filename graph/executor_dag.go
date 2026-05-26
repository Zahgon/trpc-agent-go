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

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

type dagTaskResult struct {
	task *Task
	step int
	err  error
}

type dagLoop struct {
	executor   *Executor
	ctx        context.Context
	invocation *agent.Invocation
	execCtx    *ExecutionContext
	ckptCfg    *map[string]any
	extIntr    *externalInterruptWatcher
	report     *stepExecutionReport

	executed int
	nextStep int

	ready    []*Task
	inFlight map[string]*Task
	waiting  map[string][]*Task
	sem      chan struct{}
	done     chan dagTaskResult

	draining     bool
	drainErr     error
	drainMaxStep bool

	pendingIntr     *InterruptError
	pendingIntrStep int
	pendingExtra    map[string]any

	pauseRequested bool
	rerunNodes     []string
}

func newDagLoop(
	executor *Executor,
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	ckptCfg *map[string]any,
	startStep int,
	extIntr *externalInterruptWatcher,
) (*dagLoop, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Executor) runDagLoop(
	ctx context.Context,
	invocation *agent.Invocation,
	execCtx *ExecutionContext,
	checkpointConfig *map[string]any,
	startStep int,
	extInterrupt *externalInterruptWatcher,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (l *dagLoop) run() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *dagLoop) seed() error { _ = "STUB: not implemented"; return nil }

func (l *dagLoop) consumeNextNodesFromState() []string { _ = "STUB: not implemented"; return nil }

func (l *dagLoop) tryFinalize() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (l *dagLoop) queue(t *Task) { _ = "STUB: not implemented"; return }

func (l *dagLoop) startReadyTasks() { _ = "STUB: not implemented"; return }

func (l *dagLoop) consumeReadyHead() bool { _ = "STUB: not implemented"; return false }

func (l *dagLoop) startTask(t *Task) bool { _ = "STUB: not implemented"; return false }

func (l *dagLoop) tryAcquireWorker() bool { _ = "STUB: not implemented"; return false }

func (l *dagLoop) releaseWorker() { _ = "STUB: not implemented"; return }

func (l *dagLoop) takeNextStep() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (l *dagLoop) launchTask(t *Task, step int) { _ = "STUB: not implemented"; return }

func (l *dagLoop) planIfIdle() { _ = "STUB: not implemented"; return }

func (l *dagLoop) queuePlannedTasks() { _ = "STUB: not implemented"; return }

func (l *dagLoop) waitForEvent() { _ = "STUB: not implemented"; return }

func (l *dagLoop) handleTaskResult(res dagTaskResult) { _ = "STUB: not implemented"; return }

func (l *dagLoop) startDraining(err error) { _ = "STUB: not implemented"; return }

func (l *dagLoop) startDrainingMaxStep() { _ = "STUB: not implemented"; return }

func (l *dagLoop) maybeRequestExternalInterrupt() { _ = "STUB: not implemented"; return }

func (l *dagLoop) startRequestedExternalInterrupt() { _ = "STUB: not implemented"; return }

func (l *dagLoop) lastStep() int { _ = "STUB: not implemented"; return 0 }

func (l *dagLoop) handleContextDone() { _ = "STUB: not implemented"; return }

func (l *dagLoop) handleTaskError(res dagTaskResult) { _ = "STUB: not implemented"; return }

func (l *dagLoop) recordRerunNode(res dagTaskResult) { _ = "STUB: not implemented"; return }

func (l *dagLoop) startDrainingInterrupt(
	interrupt *InterruptError,
	step int,
	metaExtra map[string]any,
) {
	_ = "STUB: not implemented"
	return
}

func forcedExternalInterruptExtra() map[string]any { _ = "STUB: not implemented"; return nil }

func (l *dagLoop) recordGraphInterruptInput(t *Task) { _ = "STUB: not implemented"; return }

func (l *dagLoop) stateInputForTask(t *Task) State { _ = "STUB: not implemented"; return *new(State) }

func isForcedExternalInterrupt(intr *InterruptError) bool { _ = "STUB: not implemented"; return false }

func (l *dagLoop) finalizeInterrupt() error { _ = "STUB: not implemented"; return nil }

func (l *dagLoop) snapshotNextNodes() []string { _ = "STUB: not implemented"; return nil }

func (l *dagLoop) maybeStartStaticInterruptBefore(t *Task) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *dagLoop) maybeStartStaticInterruptAfter(t *Task, step int) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *dagLoop) maybeCreateFinalCheckpoint() { _ = "STUB: not implemented"; return }

func (e *Executor) createDagTask(
	execCtx *ExecutionContext,
	nodeID string,
	predecessors []string,
) *Task {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) planDagTasks(execCtx *ExecutionContext) []*Task {
	_ = "STUB: not implemented"
	return nil
}

func hasStaticInterrupts(g *Graph) bool { _ = "STUB: not implemented"; return false }
