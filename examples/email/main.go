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
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Send Email  Chat Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: send_email\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &emailChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %s", err.Error())
	}
}

type emailChat struct {
	modelName string
	runner    runner.Runner
	userID    string
	sessionID string
	// current not support streaming
	streaming bool
}

// run starts the interactive chat session.
func (c *emailChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Start interactive chat.

// setup creates the runner with LLM agent and send email tool.
func (c *emailChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create email tool.
// For basic usage:

// Create LLM agent with email tool.

// Enable streaming

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *emailChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Print welcome message with examples.

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *emailChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processResponse handles the response with email tool visualization.
func (c *emailChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Detect tool responses.

// Process content from choices.

// Handle content based on streaming mode.

// Streaming mode: use delta content.

// Non-streaming mode: use full message content.

// Check if this is the final event.

// intPtr returns a pointer to the given int.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to the given float64.
	return nil
}

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
