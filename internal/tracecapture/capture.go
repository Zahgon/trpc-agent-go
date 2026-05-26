//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package tracecapture implements internal execution trace recording.
package tracecapture

import (
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/trace"
)

// StartStepInput contains the metadata needed to start a new step.
type StartStepInput struct {
	InvocationID       string
	ParentInvocationID string
	AgentName          string
	Branch             string
	NodeID             string
	StartedAt          time.Time
	PredecessorStepIDs []string
	AppliedSurfaceIDs  []string
	Input              *trace.Snapshot
}

// Capture records all steps for a single root runner.Run execution.
type Capture struct {
	mu                       sync.Mutex
	rootAgentName            string
	rootInvocationID         string
	sessionID                string
	startedAt                time.Time
	nextStepSeq              int
	steps                    []trace.Step
	stepIndexByID            map[string]int
	terminalByInvocation     map[string]map[string]struct{}
	childInvocationsByParent map[string]map[string]struct{}
}

// New creates a new capture for a root invocation.
func New(rootAgentName, rootInvocationID, sessionID string, startedAt time.Time) *Capture {
	_ = "STUB: not implemented"
	return nil
}

// SetRootAgentName updates the root agent name when it becomes available later.
func (c *Capture) SetRootAgentName(name string) { _ = "STUB: not implemented"; return }

// SetSessionID updates the session id when it becomes available later.
func (c *Capture) SetSessionID(sessionID string) { _ = "STUB: not implemented"; return }

// RegisterInvocation records a parent-child invocation relationship.
func (c *Capture) RegisterInvocation(parentInvocationID string, invocationID string) {
	_ = "STUB: not implemented"
	return
}

// StartStep records a new started step and returns its allocated step id.
func (c *Capture) StartStep(in StartStepInput) string { _ = "STUB: not implemented"; return "" }

// FinishStep updates a previously started step.
func (c *Capture) FinishStep(
	stepID string,
	output *trace.Snapshot,
	errText string,
	endedAt time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// SetStepAppliedSurfaceIDs updates the applied surface ids of one recorded step.
func (c *Capture) SetStepAppliedSurfaceIDs(stepID string, surfaceIDs []string) {
	_ = "STUB: not implemented"
	return
}

// PredecessorsForInvocation returns the current invocation predecessors for the next real step.
func (c *Capture) PredecessorsForInvocation(invocationID string, entryPredecessors []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// TerminalStepIDs returns the current terminal steps for an invocation.
func (c *Capture) TerminalStepIDs(invocationID string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Build materializes the final public trace.
func (c *Capture) Build(status trace.TraceStatus, endedAt time.Time) *trace.Trace {
	_ = "STUB: not implemented"
	return nil
}

func (c *Capture) removeTerminalStepIDsLocked(stepIDs []string) { _ = "STUB: not implemented"; return }

func (c *Capture) sortedStepIDsLocked(stepSet map[string]struct{}) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Capture) effectiveTerminalStepIDsLocked(
	invocationID string,
	visited map[string]struct{},
) []string {
	_ = "STUB: not implemented"
	return nil
}

func cloneStep(step trace.Step) trace.Step { _ = "STUB: not implemented"; return *new(trace.Step) }

func cloneSnapshot(snapshot *trace.Snapshot) *trace.Snapshot { _ = "STUB: not implemented"; return nil }
