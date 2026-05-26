//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package agent

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/internal/tracecapture"
)

// WithExecutionTraceEnabled toggles execution trace recording for this run.
func WithExecutionTraceEnabled(enabled bool) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}

// WithInvocationEntryPredecessorStepIDs sets entry predecessor step ids for a child invocation.
func WithInvocationEntryPredecessorStepIDs(stepIDs []string) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationTraceNodeID sets the static root node id for an invocation.
func WithInvocationTraceNodeID(nodeID string) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// executionTraceEnabled reports whether this invocation has execution trace enabled.
func executionTraceEnabled(inv *Invocation) bool { _ = "STUB: not implemented"; return false }

// InvocationTraceNodeID returns the invocation root node id used by execution trace.
func InvocationTraceNodeID(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

// SetInvocationSurfaceRootNodeID stores one invocation's mounted surface root node id.
func SetInvocationSurfaceRootNodeID(inv *Invocation, nodeID string) {
	_ = "STUB: not implemented"
	return
}

// ClearInvocationSurfaceRootNodeID removes one invocation's mounted surface root node id.
func ClearInvocationSurfaceRootNodeID(inv *Invocation) { _ = "STUB: not implemented"; return }

// InvocationSurfaceRootNodeID returns the effective surface root node id for one invocation.
func InvocationSurfaceRootNodeID(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

// SetInvocationTeamMemberTraceRoot stores one invocation's mounted team member trace root.
func SetInvocationTeamMemberTraceRoot(inv *Invocation, rootNodeID string) {
	_ = "STUB: not implemented"
	return
}

// ClearInvocationTeamMemberTraceRoot removes one invocation's mounted team member trace root.
func ClearInvocationTeamMemberTraceRoot(inv *Invocation) { _ = "STUB: not implemented"; return }

// InvocationTeamMemberTraceRoot returns one invocation's mounted team member trace root when present.
func InvocationTeamMemberTraceRoot(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

// StartExecutionTraceStep records a newly started real step.
func StartExecutionTraceStep(
	inv *Invocation,
	nodeID string,
	input *trace.Snapshot,
	predecessors []string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// FinishExecutionTraceStep finalizes a previously started step.
func FinishExecutionTraceStep(
	inv *Invocation,
	stepID string,
	output *trace.Snapshot,
	stepErr error,
) {
	_ = "STUB: not implemented"
	return
}

// SetExecutionTraceStepAppliedSurfaceIDs records one step's applied surfaces from the invocation agent when supported.
func SetExecutionTraceStepAppliedSurfaceIDs(inv *Invocation, stepID string) {
	_ = "STUB: not implemented"
	return
}

// NextExecutionTracePredecessors returns the predecessor set for the next real step or child invocation.
func NextExecutionTracePredecessors(inv *Invocation) []string {
	_ = "STUB: not implemented"
	return nil
}

// BuildExecutionTrace builds the final trace for the root invocation.
func BuildExecutionTrace(
	inv *Invocation,
	status trace.TraceStatus,
) *trace.Trace {
	_ = "STUB: not implemented"
	return nil
}

func (inv *Invocation) initializeExecutionTrace() { _ = "STUB: not implemented"; return }

func (inv *Invocation) ensureTraceCaptureMetadata() { _ = "STUB: not implemented"; return }

func (inv *Invocation) ensureTraceNodeID() { _ = "STUB: not implemented"; return }

func (inv *Invocation) ensureTraceNodeIDLocked() { _ = "STUB: not implemented"; return }

func (inv *Invocation) executionTraceCapture() *tracecapture.Capture {
	_ = "STUB: not implemented"
	return nil
}

func (inv *Invocation) executionTraceFields() (*tracecapture.Capture, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func cloneStringSlice(values []string) []string { _ = "STUB: not implemented"; return nil }

func escapeTraceLocalName(name string) string { _ = "STUB: not implemented"; return "" }
