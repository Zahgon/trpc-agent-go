//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates interactive chat using DuckDuckGo Instant Answer API.
// The tool provides factual, encyclopedic information but not real-time data.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func main() {
	// Parse command line flags.
	modelName := flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	flag.Parse()

	fmt.Printf("🚀 DuckDuckGo Search Chat Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: duckduckgo_search\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &searchChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// searchChat manages the conversation with DuckDuckGo search.
type searchChat struct {
	modelName string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the interactive chat session.
func (c *searchChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and DuckDuckGo search tool.
func (c *searchChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create DuckDuckGo search tool.
// For basic usage:

// Create LLM agent with DuckDuckGo search tool.

// Enable streaming

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *searchChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Print welcome message with examples.

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *searchChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processStreamingResponse handles the streaming response with search tool visualization.
func (c *searchChat) processStreamingResponse(eventChan <-chan *event.Event) error {
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

// intPtr returns a pointer to the given int.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to the given float64.
	return nil
}

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
