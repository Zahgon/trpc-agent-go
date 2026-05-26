//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graph

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
)

type traceTaskRegistryEntry struct {
	task *traceTaskMetadata
}

type traceTaskMetadata struct {
	mu                   sync.Mutex
	owner                *ExecutionContext
	taskID               string
	nodeID               string
	predecessorStepIDs   []string
	preRunInputSnapshot  *atrace.Snapshot
	childTerminalStepIDs []string
	wrapperStepID        string
	postChildStepID      string
	claimed              bool
	fallbackToWrapper    bool
}

func newTraceTaskMetadata(
	owner *ExecutionContext,
	taskID string,
	nodeID string,
	predecessorStepIDs []string,
	inputSnapshot *atrace.Snapshot,
) *traceTaskMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) registerAgentNodeTraceTask(execCtx *ExecutionContext, task *traceTaskMetadata) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Executor) unregisterAgentNodeTraceTask(execCtx *ExecutionContext, nodeID string, task *traceTaskMetadata) {
	_ = "STUB: not implemented"
	return
}

func claimAgentNodeTraceTask(state State) *traceTaskMetadata { _ = "STUB: not implemented"; return nil }

func (e *ExecutionContext) claimAgentNodeTraceTask(nodeID string) *traceTaskMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (m *traceTaskMetadata) childEntryPredecessorStepIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

func (m *traceTaskMetadata) setChildTerminalStepIDs(stepIDs []string) {
	_ = "STUB: not implemented"
	return
}

func (m *traceTaskMetadata) markFallbackToWrapper() { _ = "STUB: not implemented"; return }

func (m *traceTaskMetadata) shouldFallbackToWrapper() bool { _ = "STUB: not implemented"; return false }

func (m *traceTaskMetadata) materializeWrapper(invocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *traceTaskMetadata) materializePostChildStep(
	invocation *agent.Invocation,
	predecessorStepIDs []string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *traceTaskMetadata) snapshot() traceTaskMetadataSnapshot {
	_ = "STUB: not implemented"
	return *new(traceTaskMetadataSnapshot)
}

type traceTaskMetadataSnapshot struct {
	childTerminalStepIDs []string
	claimed              bool
	fallbackToWrapper    bool
}

func traceNodeIDForAgentNode(invocation *agent.Invocation, nodeID string) string {
	_ = "STUB: not implemented"
	return ""
}
