//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/graph"
)

// createNodeCallbacks creates comprehensive callbacks for monitoring and performance tracking.
func (w *fanoutWorkflow) createNodeCallbacks() *graph.NodeCallbacks {
	_ = "STUB: not implemented"
	return nil
}

// onBeforeNode records timing and execution metadata.
func (w *fanoutWorkflow) onBeforeNode(ctx context.Context, callbackCtx *graph.NodeCallbackContext, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// onAfterNode updates history, adds metadata, and appends task results.
func (w *fanoutWorkflow) onAfterNode(
	ctx context.Context,
	callbackCtx *graph.NodeCallbackContext,
	state graph.State,
	result any,
	nodeErr error,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Append result for process_task when success: include actual model output

// Prefer node-specific response content from the node's result first

// Fallback to previous state when result doesn't contain it (should be rare)

// Enrich state result with execution metadata if applicable.

// Persist execution history so downstream nodes (like aggregator) can read it.

// onNodeError records error info and lightweight classification.
func (w *fanoutWorkflow) onNodeError(
	ctx context.Context,
	callbackCtx *graph.NodeCallbackContext,
	state graph.State,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// helpers
func (w *fanoutWorkflow) computeExecutionTime(state graph.State, nodeID string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (w *fanoutWorkflow) updateLastHistory(state graph.State, dur time.Duration, nodeErr error) {
	_ = "STUB: not implemented"
	return
}

func (w *fanoutWorkflow) maybeWarnSlow(nodeName string, dur time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (w *fanoutWorkflow) buildTaskResultString(state graph.State) string {
	_ = "STUB: not implemented"
	return ""
}
