//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates multi-turn chat using the Runner with streaming output, session management,
// tool calling, and shows how to use AgentCallbacks, ModelCallbacks, and ToolCallbacks.
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
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Multi-turn Chat with Runner + Tools + Callbacks\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Available tools: calculator, current_time\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &multiTurnChatWithCallbacks{
		modelName: *modelName,
		streaming: *streaming,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// multiTurnChatWithCallbacks manages the chat with callbacks.
type multiTurnChatWithCallbacks struct {
	modelName string
	streaming bool
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the interactive chat session.
func (c *multiTurnChatWithCallbacks) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and tools.
func (c *multiTurnChatWithCallbacks) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create tools.

// Create callbacks.

// Create LLM agent with tools and callbacks.

// Create runner with in-memory session service.

// Setup identifiers.

// createTools creates the tools for the agent.
func (c *multiTurnChatWithCallbacks) createTools() []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// startChat runs the interactive conversation loop.
func (c *multiTurnChatWithCallbacks) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle special commands.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *multiTurnChatWithCallbacks) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process response.

// processResponse handles both streaming and non-streaming responses with tool call visualization.
func (c *multiTurnChatWithCallbacks) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is the final event.
// Don't break on tool response events (Done=true but not final assistant response).

// handleEvent processes a single event from the event channel.
func (c *multiTurnChatWithCallbacks) handleEvent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle tool calls.

// Handle tool responses.

// Handle content.

// handleToolCalls processes tool call events and returns true if handled.
func (c *multiTurnChatWithCallbacks) handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses processes tool response events and returns true if handled.
func (c *multiTurnChatWithCallbacks) handleToolResponses(event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// handleContent processes content events for both streaming and non-streaming modes.
func (c *multiTurnChatWithCallbacks) handleContent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

// extractContent extracts content based on streaming mode.
func (c *multiTurnChatWithCallbacks) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"

	// Streaming mode: use delta content.
	return ""
}

// Non-streaming mode: use full message content.

// displayContent displays the content with proper formatting.
func (c *multiTurnChatWithCallbacks) displayContent(
	content string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

// startNewSession creates a new session ID.
func (c *multiTurnChatWithCallbacks) startNewSession() { _ = "STUB: not implemented"; return }
