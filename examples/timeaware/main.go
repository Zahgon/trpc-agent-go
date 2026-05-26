//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates multi-turn chat using the Runner with streaming
// output and time awareness.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName  = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming  = flag.Bool("streaming", true, "Enable streaming mode for responses")
	addTime    = flag.Bool("add-time", true, "Add current time to the system prompt")
	timezone   = flag.String("timezone", "UTC", "Timezone for time display (e.g., UTC, EST, PST)")
	timeFormat = flag.String("time-format", "2006-01-02 15:04:05 UTC", "Time format for display (Go time format)")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Multi-turn Chat with Time Aware\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Add Current Time: %t\n", *addTime)
	fmt.Printf("Timezone: %s\n", *timezone)
	fmt.Printf("Time Format: %s\n", *timeFormat)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &multiTurnChat{
		modelName:  *modelName,
		streaming:  *streaming,
		addTime:    *addTime,
		timezone:   *timezone,
		timeFormat: *timeFormat,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// multiTurnChat manages the conversation.
type multiTurnChat struct {
	modelName  string
	streaming  bool
	runner     runner.Runner
	userID     string
	addTime    bool
	timezone   string
	timeFormat string
}

// run starts the interactive chat session.
func (c *multiTurnChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent.
func (c *multiTurnChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create LLM agent.

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *multiTurnChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *multiTurnChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process response.

// processResponse handles both streaming and non-streaming responses.
func (c *multiTurnChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is the final event.

// handleEvent processes a single event from the event channel.
func (c *multiTurnChat) handleEvent(event *event.Event, fullContent *string) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle content.

// extractContent extracts content based on streaming mode.
func (c *multiTurnChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"

	// Streaming mode: use delta content.
	return ""
}

// Non-streaming mode: use full message content.

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
