//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates tool execution timing using ToolCallbacks with OpenTelemetry integration.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	"go.opentelemetry.io/otel/metric"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", false, "Enable streaming mode for responses")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Println("🚀 Tool Timer with Telemetry Example")
	fmt.Println("This example demonstrates how to use ToolCallbacks to measure tool execution time and report to OpenTelemetry.")
	fmt.Println(strings.Repeat("=", 70))

	// Initialize OpenTelemetry.
	meter, err := initTelemetry()
	if err != nil {
		log.Fatalf("Failed to initialize telemetry: %v", err)
	}

	// Create the example.
	example := &toolTimerExample{meter: meter}

	// Setup and run.
	if err := example.run(); err != nil {
		log.Fatalf("Example failed: %v", err)
	}
}

// toolTimerExample demonstrates tool execution timing with telemetry integration.
type toolTimerExample struct {
	meter     metric.Meter
	runner    runner.Runner
	userID    string
	sessionID string
	// Add telemetry metrics.
	agentDurationHistogram metric.Float64Histogram
	toolDurationHistogram  metric.Float64Histogram
	modelDurationHistogram metric.Float64Histogram
	agentCounter           metric.Int64Counter
	toolCounter            metric.Int64Counter
	modelCounter           metric.Int64Counter
}

// run executes the tool timer example.
func (e *toolTimerExample) run() error { _ = "STUB: not implemented"; return nil }

// Initialize telemetry metrics.

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Run the example.

// setup creates the runner with LLM agent, tools and telemetry metrics.
func (e *toolTimerExample) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model using flag.
	return nil
}

// Create tools.

// Create callbacks for timing.

// Create LLM agent with tools and callbacks.

// Create runner.

// Setup identifiers.

// createTools creates the tools for the agent.
func (e *toolTimerExample) createTools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// runExample executes the interactive chat session for tool timer example.
func (e *toolTimerExample) runExample(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle special commands.

// Process the user message.

// Add spacing after error.

// Wait briefly for AfterAgentCallback to complete its output.
// This ensures timing information appears before the next prompt.

// processMessage handles a single message exchange.
func (e *toolTimerExample) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process response.

// startNewSession creates a new session ID.
func (e *toolTimerExample) startNewSession() { _ = "STUB: not implemented"; return }

// processResponse handles the response from the agent.
func (e *toolTimerExample) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Handle tool calls.

// Handle tool responses.

// Handle content.

// Check if this is the final event.
