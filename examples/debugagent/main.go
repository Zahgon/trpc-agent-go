//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates interactive chat using file operation tools.
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
	baseDir := flag.String("base-dir", ".", "Base directory for file operations")
	flag.Parse()

	fmt.Printf("🐞 Debug Agent Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Base Directory: %s\n", *baseDir)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: save_file, read_file, read_multiple_files, list_file, search_file, search_content, replace_content\n")
	fmt.Println(strings.Repeat("=", 50))
	// Create and run the chat.
	chat := &debugChat{
		modelName: *modelName,
		baseDir:   *baseDir,
	}
	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// debugChat manages the conversation with file operation tools.
type debugChat struct {
	modelName string
	baseDir   string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the interactive chat session.
func (c *debugChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and file operation tools.
func (c *debugChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create file operation tools.

// Create LLM agent with file operation tools.

// Enable streaming

// Execution-first instruction: allow direct bash or write+run script; never write without running

// Create LLM agent with file operation tools and code executor.

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *debugChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Print welcome message with examples.

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *debugChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processStreamingResponse handles the streaming response with file tool visualization.
func (c *debugChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Detect tool responses.

// Process streaming content.

// Handle streaming delta content.

// Only break on runner completion to ensure we consume all postprocessing events

// intPtr returns a pointer to the given int.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to the given float64.
	return nil
}

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
