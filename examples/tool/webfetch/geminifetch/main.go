//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates interactive chat using Gemini web fetch tool.
// The tool uses Gemini's URL Context feature for server-side web fetching and analysis.
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
	geminiModel := flag.String("gemini-model", "gemini-2.5-flash", "Gemini model for web fetching")
	flag.Parse()

	fmt.Printf("🚀 Gemini Web Fetch Chat Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Gemini Fetch Model: %s\n", *geminiModel)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: gemini_web_fetch\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &geminiWebFetchChat{
		modelName:   *modelName,
		geminiModel: *geminiModel,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// geminiWebFetchChat manages the conversation with Gemini web fetch capability.
type geminiWebFetchChat struct {
	modelName   string
	geminiModel string
	runner      runner.Runner
	userID      string
	sessionID   string
}

// run starts the interactive chat session.
func (c *geminiWebFetchChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Start interactive chat.

// setup creates the runner with LLM agent and Gemini web fetch tool.
func (c *geminiWebFetchChat) setup(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create Gemini web fetch tool.

// Create LLM agent with Gemini web fetch tool.

// Enable streaming

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *geminiWebFetchChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Print welcome message with examples.

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *geminiWebFetchChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processStreamingResponse handles the streaming response with Gemini web fetch tool visualization.
func (c *geminiWebFetchChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Detect tool responses.

// Truncate long tool responses for display

// Process streaming content.

// Handle streaming delta content.

// Check if this is the final event.

// intPtr returns a pointer to the given int.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to the given float64.
	return nil
}

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
