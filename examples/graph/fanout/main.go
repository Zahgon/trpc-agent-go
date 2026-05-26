//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates parallel fan-out execution using the graph package.
// This example shows how to build and execute graphs with parallel task distribution,
// LLM nodes, and function nodes using []*Command for dynamic fan-out.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	// Default model name for deepseek-v4-flash.
	defaultModelName = "deepseek-v4-flash"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
)

func main() {
	// Parse command line flags.
	flag.Parse()
	fmt.Printf("🚀 Parallel Fan-out Workflow Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))
	// Create and run the workflow.
	workflow := &fanoutWorkflow{
		modelName: *modelName,
	}
	if err := workflow.run(); err != nil {
		log.Fatalf("Workflow failed: %v", err)
	}
}

// fanoutWorkflow manages the parallel fan-out workflow.
type fanoutWorkflow struct {
	modelName string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the parallel fan-out workflow.
func (w *fanoutWorkflow) run() error { _ = "STUB: not implemented"; return nil }

// Setup the workflow.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// setup creates the graph agent and runner.
func (w *fanoutWorkflow) setup() error {
	_ = "STUB: not implemented"
	// Create the parallel fan-out graph.
	return nil
}

// Create GraphAgent from the compiled graph.

// Create session service.

// Create runner with the graph agent.

// Setup identifiers.

// createFanoutGraph creates a parallel fan-out workflow graph.
func (w *fanoutWorkflow) createFanoutGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Create extended state schema for messages and metadata.
	return nil, nil
}

// Create model instance.

// Create analysis tools.

// Create node callbacks for monitoring and performance tracking.

// Create stateGraph with schema and callbacks.

// Build the workflow graph.

// Add input analysis node.

// Add LLM task planning node.

// Add fan-out node that returns []*Command.

// Add LLM worker node for processing individual tasks.

// Add aggregation node.

// Set up the workflow routing.

// Add workflow edges - this is the key fix for the tool execution flow.

// This edge allows the LLM to continue after tool execution

// Build and return the graph.

// Node function implementations.

func (w *fanoutWorkflow) analyzeInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Get input from GraphAgent's state fields.
	return *new(any), nil
}

// Basic input analysis.

// Return state with input analysis results.

func (w *fanoutWorkflow) createFanoutTasks(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Get the number of tasks from the LLM's response.
	return *new(any), nil
}

// Robustly extract the first integer anywhere in the output (handles **2**, etc.)

// sensible default

// Ensure reasonable bounds.

// Generate commands for parallel execution.

// Build a dedicated user input for the worker stage so the LLM
// incorporates task parameters instead of continuing prior context only.

func (w *fanoutWorkflow) aggregateResults(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Extract results from the parallel tasks.
	return *new(any), nil
}

// Extract execution metadata.

// Create final aggregated output.

// formatTaskResults formats the task results into a readable string.
func (w *fanoutWorkflow) formatTaskResults(results []string) string {
	_ = "STUB: not implemented"
	return ""
}

// formatExecutionStats formats the execution history into a readable string.
func (w *fanoutWorkflow) formatExecutionStats(history []map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// Calculate average execution time.

// Tool function implementations.

func (w *fanoutWorkflow) analyzeTaskComplexity(ctx context.Context, args taskAnalysisArgs) (taskAnalysisResult, error) {
	_ = "STUB: not implemented"

	// Simple complexity analysis.
	return *new(taskAnalysisResult), nil
}

// startInteractiveMode starts the interactive parallel fan-out mode.
func (w *fanoutWorkflow) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Allow larger inputs than the default ~64KB token limit

// Process the content through parallel fan-out.

// Add spacing between content.

// processContent processes a single content through the parallel fan-out workflow.
func (w *fanoutWorkflow) processContent(ctx context.Context, content string) error {
	_ = "STUB: not implemented"
	// Create user message.
	return nil
}

// Run the workflow through the runner.

// Set runtime state for each run.

// Process streaming response.

// processStreamingResponse handles the streaming workflow response.
func (w *fanoutWorkflow) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// After completion, replay final results sequentially to improve readability.

// fakeStreamFinalResults replays final results sequentially with small delays
// to mimic streaming, improving readability after parallel fan-out completes.
func (w *fanoutWorkflow) fakeStreamFinalResults(e *event.Event, delay time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Prefer "final_results", fallback to "results"

// Split header (first line) and body (rest)

func (w *fanoutWorkflow) handleErrorEvent(e *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *fanoutWorkflow) handleGraphNodeEvent(e *event.Event, maxPreviewLen int) {
	_ = "STUB: not implemented"
	return
}

func (w *fanoutWorkflow) processNodeDelta(data []byte, maxPreviewLen int) {
	_ = "STUB: not implemented"
	return
}

func (w *fanoutWorkflow) processToolDelta(data []byte) { _ = "STUB: not implemented"; return }

func (w *fanoutWorkflow) processModelDelta(data []byte, maxPreviewLen int) {
	_ = "STUB: not implemented"
	return
}

func (w *fanoutWorkflow) handleStreamingChoices(e *event.Event, started *bool) {
	_ = "STUB: not implemented"
	return
}

// Only stream planning node output to avoid interleaved parallel outputs.

func (w *fanoutWorkflow) trackStageProgress(e *event.Event, stageCount *int) {
	_ = "STUB: not implemented"
	// Keep stage tracking silent to avoid noisy logs in this example.
	return
}

// Reducers for schema-managed aggregation
func appendMapSliceReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

func intSumReducer(existing, update any) any { _ = "STUB: not implemented"; return *new(any) }

// showHelp displays available commands.
func (w *fanoutWorkflow) showHelp() { _ = "STUB: not implemented"; return }

// formatJSON formats JSON strings for better readability.
func formatJSON(jsonStr string) string { _ = "STUB: not implemented"; return "" }

// Try to pretty print the JSON.

// Fallback to original string if not valid JSON.

// truncateString truncates a string to the specified length and adds ellipsis if needed.
func truncateString(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }

// Type definitions for tool functions.

type taskAnalysisArgs struct {
	Text string `json:"text" description:"Text to analyze for task complexity"`
}

type taskAnalysisResult struct {
	Complexity     string `json:"complexity"`
	WordCount      int    `json:"word_count"`
	SentenceCount  int    `json:"sentence_count"`
	SuggestedTasks int    `json:"suggested_tasks"`
}
