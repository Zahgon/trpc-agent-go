//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use the OpenAPI Tool.
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName    = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	modelBaseURL = flag.String("base_url", "", "Base URL for the model API")
	modelToken   = flag.String("api_token", "", "Authentication token for the model API")
	openAPISpec  = flag.String("openapi_spec", "./petstore3.yaml", "Path to the OpenAPI specification file")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Chat with LLMAgent\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &llmAgentChat{
		modelName:   *modelName,
		baseURL:     *modelBaseURL,
		Token:       *modelToken,
		openAPISpec: *openAPISpec,
	}
	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// llmAgentChat manages the conversation.
type llmAgentChat struct {
	modelName   string
	baseURL     string
	Token       string
	openAPISpec string

	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the interactive chat session.
func (c *llmAgentChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the agent.

// Start interactive chat.

// setup creates the LLMAgent.
func (c *llmAgentChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create a model instance.
	return nil
}

// Create generation config.

// Create an LLMAgent with configuration.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *llmAgentChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle special commands.

// Process the user message.

// Add spacing between turns.

// processMessage handles a single message exchange.
func (c *llmAgentChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent.

// Process response.

// processResponse handles the streaming response.
func (c *llmAgentChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is the final event.

// handleEvent processes a single event from the event channel.
func (c *llmAgentChat) handleEvent(event *event.Event, fullContent *strings.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle content.

// extractContent extracts content based on streaming mode.
func (c *llmAgentChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	// In streaming mode, use delta content for real-time display.
	// In non-streaming mode, use full message content.
	return ""
}

// intPtr returns a pointer to the given int value.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to the given float64 value.
	return nil
}

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
