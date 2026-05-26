//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates streamable http and stdio mcp tools usage.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

func main() {
	// Parse command line flags.
	modelName := flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	flag.Parse()

	fmt.Printf("🚀 MCP tools usage (STDIO, Streamable HTTP, and SSE)\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: calculator, current_time, echo, add, get_weather, get_news, sse_echo, sse_info\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &multiTurnChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// multiTurnChat manages the conversation.
type multiTurnChat struct {
	modelName  string
	runner     runner.Runner
	userID     string
	sessionID  string
	mcpToolSet []*mcp.ToolSet
}

// run starts the interactive chat session.
func (c *multiTurnChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and tools.
func (c *multiTurnChat) setup(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create tools.

// Create Stdio MCP tools.

// Create Streamable MCP tools.

// Use ServerURL instead of URL

// WithSimpleRetry(3): Uses default settings with 3 retry attempts
// - MaxRetries: 3 (range: 0-10)
// - InitialBackoff: 500ms (default, range: 1ms-30s)
// - BackoffFactor: 2.0 (default, range: 1.0-10.0)
// - MaxBackoff: 8s (default, range: up to 5 minutes)
// Retry sequence: 500ms -> 1s -> 2s (total max delay: ~3.5s)

// other mcp options.
// tmcp.WithHTTPHeaders(http.Header{
// 	"User-Agent": []string{"trpc-agent-go/1.0.0"},
// }),

// Create SSE MCP tools with session reconnection.

// SSE server URL.

// Enable session reconnection for automatic recovery when server restarts (max 3 attempts)

// WithRetry: Custom retry configuration for fine-tuned control.
// Retry sequence: 1s -> 1.5s -> 2.25s -> 3.375s -> 5.0625s (capped at 15s)

// Maximum retry attempts (range: 0-10, default: 2)
// Initial delay before first retry (range: 1ms-30s, default: 500ms)
// Exponential backoff multiplier (range: 1.0-10.0, default: 2.0)
// Maximum delay cap (range: up to 5 minutes, default: 8s)

// Create LLM agent with tools.

// Enable streaming

// Create runner.

// Setup identifiers.

// Store toolsets for proper cleanup.

// startChat runs the interactive conversation loop.
func (c *multiTurnChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle exit command.

// Process the user message.

// Add spacing between turns

// Close the toolset when done.

// processMessage handles a single message exchange.
func (c *multiTurnChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processStreamingResponse handles the streaming response with tool call visualization.
func (c *multiTurnChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Detect tool responses.

// Process streaming content.

// Handle streaming delta content.

// Check if this is the final event.
// Don't break on tool response events (Done=true but not final assistant response).

// CallableTool implementations.

// calculate performs basic mathematical operations.
func (c *multiTurnChat) calculate(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

// Handle division by zero

// getCurrentTime returns current time information.
func (c *multiTurnChat) getCurrentTime(_ context.Context, args timeArgs) (timeResult, error) {
	_ = "STUB: not implemented"
	return *new(timeResult), nil
}

// Handle timezone conversion.

// Simplified EST

// Simplified PST

// Simplified CST

// calculatorArgs represents arguments for the calculator tool.
type calculatorArgs struct {
	Operation string  `json:"operation" description:"The operation to perform"`
	A         float64 `json:"a" description:"First number"`
	B         float64 `json:"b" description:"Second number"`
}

// calculatorResult represents the result of a calculation.
type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

// timeArgs represents arguments for the time tool.
type timeArgs struct {
	Timezone string `json:"timezone" description:"Timezone or leave empty for local"`
}

// timeResult represents the current time information.
type timeResult struct {
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
	Date     string `json:"date"`
	Weekday  string `json:"weekday"`
}

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
