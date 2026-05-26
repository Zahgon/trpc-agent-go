//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates per-node callbacks functionality in the graph package.
// This example shows how to use both global and per-node callbacks for fine-grained
// control over node execution behavior.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

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
	interactive = flag.Bool("interactive", false,
		"Run in interactive mode")
)

func main() {
	// Parse command line flags.
	flag.Parse()
	fmt.Printf("🚀 Per-Node Callbacks Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the workflow.
	workflow := &perNodeCallbacksWorkflow{
		modelName: *modelName,
	}
	if err := workflow.run(); err != nil {
		log.Fatalf("Workflow failed: %v", err)
	}
}

// perNodeCallbacksWorkflow demonstrates per-node callback functionality.
type perNodeCallbacksWorkflow struct {
	modelName string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the per-node callbacks workflow.
func (w *perNodeCallbacksWorkflow) run() error { _ = "STUB: not implemented"; return nil }

// Setup the workflow.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// setup creates the graph agent and runner.
func (w *perNodeCallbacksWorkflow) setup() error {
	_ = "STUB: not implemented"
	// Create the workflow graph.
	return nil
}

// Create GraphAgent from the compiled graph.

// Create session service.

// Create runner.

// Generate session ID.

// createWorkflowGraph creates the workflow graph with per-node callbacks.
func (w *perNodeCallbacksWorkflow) createWorkflowGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Define state schema.
	return nil, nil
}

// Create global callbacks for logging and monitoring.

// Create the workflow graph with global callbacks.

// Step 1: Process input with custom pre-callback that modifies input.

// Modify input to add a prefix.

// Step 2: Transform result with error handling callback.

// Check if step1 result exists.

// Set a fallback result on error.

// Step 3: Final processing with conditional callback.

// Check input length and potentially skip processing.

// Return a custom state update to skip node execution
// and mark a final result directly.

// Add a timestamp to the result.

// Set up the workflow edges.

// Node functions for the workflow.

// processStep1 processes the input and returns a result.
func (w *perNodeCallbacksWorkflow) processStep1(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate some processing time.

// processStep2 transforms the step1 result.
func (w *perNodeCallbacksWorkflow) processStep2(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate potential error for demonstration.

// processStep3 performs final processing.
func (w *perNodeCallbacksWorkflow) processStep3(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate some processing time.

// runDefaultExamples runs predefined examples.
func (w *perNodeCallbacksWorkflow) runDefaultExamples(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Add delay between examples.

// startInteractiveMode starts the interactive mode.
func (w *perNodeCallbacksWorkflow) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Process the input.

// Add spacing between inputs.

// processInput processes a single input through the workflow.
func (w *perNodeCallbacksWorkflow) processInput(ctx context.Context, input string) error {
	_ = "STUB: not implemented"
	// Create user message.
	return nil
}

// Run the workflow through the runner.

// Set runtime state for each run.

// Process streaming response.

// processStreamingResponse handles the streaming workflow response.
func (w *perNodeCallbacksWorkflow) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Track node execution events via metadata regardless of author.

// Handle errors after metadata so per-node error callbacks
// can be observed via metadata.

// Process streaming content from LLM nodes.

// Handle streaming delta content.

// Add newline when streaming is complete.

// Track workflow stages.

// Handle completion and final response.

// Check for final response in the completion event.

// Check for final result in state delta.

// showHelp displays help information.
func (w *perNodeCallbacksWorkflow) showHelp() { _ = "STUB: not implemented"; return }
