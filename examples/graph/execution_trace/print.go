//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
)

type normalizedStep struct {
	Label        string
	NodeID       string
	Predecessors []string
	Input        string
	Output       string
	Error        string
}

func printExecutionTrace(executionTrace *atrace.Trace, staticNodeIDs []string) {
	_ = "STUB: not implemented"
	return
}

func normalizeSteps(executionTrace *atrace.Trace) ([]normalizedStep, map[string]int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func summarizeSnapshot(snapshot *atrace.Snapshot) string { _ = "STUB: not implemented"; return "" }

func formatStringList(values []any) string { _ = "STUB: not implemented"; return "" }

func collectTraceEdges(steps []normalizedStep) []string { _ = "STUB: not implemented"; return nil }

func collectSkippedNodeIDs(staticNodeIDs []string, countsByNode map[string]int) []string {
	_ = "STUB: not implemented"
	return nil
}

func collectRepeatedNodeIDs(countsByNode map[string]int) []string {
	_ = "STUB: not implemented"
	return nil
}

func finalStepLabels(steps []normalizedStep) []string { _ = "STUB: not implemented"; return nil }
