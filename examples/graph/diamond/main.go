//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a diamond pattern workflow that exposes
// the need for per-node version tracking (versions_seen).
// Without proper versions_seen implementation, the aggregator node
// will execute multiple times instead of once.
package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	// Node names.
	nodeSplitter   = "splitter"
	nodeAnalyzer1  = "analyzer1"
	nodeAnalyzer2  = "analyzer2"
	nodeAggregator = "aggregator"
	nodeFinal      = "final"

	// State keys.
	stateKeyInput      = "input"
	stateKeyAnalysis1  = "analysis1_data"
	stateKeyAnalysis2  = "analysis2_data"
	stateKeyResults    = "results"
	stateKeyExecCounts = "execution_counts"

	// Default values.
	defaultUserID  = "demo-user"
	defaultAppName = "diamond-workflow"
)

// diamondWorkflow manages the diamond pattern workflow.
type diamondWorkflow struct {
	graph      *graph.Graph
	graphAgent agent.Agent
	saver      graph.CheckpointSaver
	runner     runner.Runner

	// Track execution counts to expose the issue.
	executionCounts map[string]int
	executionMutex  sync.Mutex
}

func main() {
	fmt.Println("🔷 Diamond Pattern Workflow Example")
	fmt.Println("Demonstrates per-node version tracking and correct result aggregation.")
	fmt.Println("On resume, versions_seen prevents redundant node executions.")

	// Create and initialize workflow.
	workflow := &diamondWorkflow{
		executionCounts: make(map[string]int),
	}

	if err := workflow.initialize(); err != nil {
		log.Fatalf("Failed to initialize workflow: %v", err)
	}

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer workflow.runner.Close()

	// Run interactive mode.
	if err := workflow.runInteractive(); err != nil {
		log.Fatalf("Interactive mode failed: %v", err)
	}
}

// initialize sets up the workflow components.
func (w *diamondWorkflow) initialize() error {
	_ = "STUB: not implemented"
	// Create the graph.
	return nil
}

// Create checkpoint saver.

// Create graph agent with checkpointing.

// Create session service.

// Create runner.

// createGraph creates the diamond pattern graph.
func (w *diamondWorkflow) createGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Create state schema.
	return nil, nil
}

// Add state fields.

// Append results from both analyzers.

// Create state graph.

// Add nodes.

// Add edges for diamond pattern.
// Splitter fans out to both analyzers.

// Both analyzers converge to aggregator.

// Aggregator routes to final only when both results are ready (barrier).
// Otherwise it does not route further (to End), and will be triggered
// again when the second analyzer completes.

// Node implementations.

func (w *diamondWorkflow) splitterNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Split work to both analysis channels.

func (w *diamondWorkflow) analyzer1Node(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate some processing time.

func (w *diamondWorkflow) analyzer2Node(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Different processing time to ensure analyzers finish at different times.

func (w *diamondWorkflow) aggregatorNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// THIS IS THE KEY ISSUE EXPOSURE.

// Log what we're aggregating.

func (w *diamondWorkflow) finalNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Read live execution counts to reflect final node's own execution.

// Analyze execution counts.

// Helper methods.

func (w *diamondWorkflow) recordExecution(nodeName string) int { _ = "STUB: not implemented"; return 0 }

func (w *diamondWorkflow) getExecutionCounts() map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func (w *diamondWorkflow) resetExecutionCounts() { _ = "STUB: not implemented"; return }

// runInteractive runs the workflow in interactive mode.
func (w *diamondWorkflow) runInteractive() error { _ = "STUB: not implemented"; return nil }

// best-effort parse

// runWorkflow executes the workflow with the given input.
func (w *diamondWorkflow) runWorkflow(ctx context.Context, inputData string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create initial message.

// Create runtime state with input.

// Run the workflow.

// Process events.

// Silently consume events for cleaner output.

// runWorkflowWithLineage executes the workflow with a fixed lineage_id.
func (w *diamondWorkflow) runWorkflowWithLineage(ctx context.Context, inputData, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// resumeWorkflow resumes from latest or specific checkpoint of a lineage.
func (w *diamondWorkflow) resumeWorkflow(ctx context.Context, lineageID, checkpointID string) error {
	_ = "STUB: not implemented"
	return nil
}

// For resume, the input message content is not used to reconstruct state;
// state is restored from checkpoint by the executor.

// listCheckpoints lists recent checkpoints for a lineage.
func (w *diamondWorkflow) listCheckpoints(ctx context.Context, lineageID string, limit int) error {
	_ = "STUB: not implemented"
	return nil
}
