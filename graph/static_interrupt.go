//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

// StaticInterruptKeyPrefixBefore/After are interrupt key prefixes used by
// static interrupts (debug breakpoints).
const (
	StaticInterruptKeyPrefixBefore = "static_interrupt_before:"
	StaticInterruptKeyPrefixAfter  = "static_interrupt_after:"
)

// StaticInterruptPhase indicates when a static interrupt is triggered.
type StaticInterruptPhase string

// Supported static interrupt phases.
const (
	StaticInterruptPhaseBefore StaticInterruptPhase = "before"
	StaticInterruptPhaseAfter  StaticInterruptPhase = "after"
)

// StaticInterruptPayload is stored in InterruptError.Value for static
// interrupts, providing structured information for debugging UIs.
type StaticInterruptPayload struct {
	Phase       StaticInterruptPhase `json:"phase"`
	Nodes       []string             `json:"nodes"`
	ActiveNodes []string             `json:"activeNodes,omitempty"`
}

func (e *Executor) maybeStaticInterruptBefore(
	execCtx *ExecutionContext,
	tasks []*Task,
	step int,
) *InterruptError {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) maybeStaticInterruptAfter(
	tasks []*Task,
	step int,
) *InterruptError {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) staticInterruptHits(
	tasks []*Task,
	before bool,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Executor) newStaticInterruptError(
	phase StaticInterruptPhase,
	nodes []string,
	activeNodes []string,
	step int,
) *InterruptError {
	_ = "STUB: not implemented"
	return nil
}

func uniqueSortedTaskNodes(tasks []*Task) []string { _ = "STUB: not implemented"; return nil }

func keysOfSet(set map[string]struct{}) []string { _ = "STUB: not implemented"; return nil }

func getStaticInterruptSkips(state State) map[string]any { _ = "STUB: not implemented"; return nil }

func hasSkipsForAll(skips map[string]any, nodeIDs []string) bool {
	_ = "STUB: not implemented"
	return false
}

func setSkips(skips map[string]any, nodeIDs []string) { _ = "STUB: not implemented"; return }

func clearSkips(state State, skips map[string]any, nodeIDs []string) {
	_ = "STUB: not implemented"
	return
}
