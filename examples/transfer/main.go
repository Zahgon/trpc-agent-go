//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates agent transfer functionality with sub-agents.
// This example shows how to create a main agent with multiple specialized
// sub-agents and use the transfer_to_agent tool to delegate tasks.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	alog "trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func main() {
	// Parse command line flags.
	modelName := flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	debug := flag.Bool("debug", false, "Enable debug logging and verbose event traces")
	endInvocation := flag.Bool("end-invocation", false, "Enable end parent invocation after transfer")
	flag.Parse()

	fmt.Printf("🔄 Agent Transfer Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available sub-agents: math-agent, weather-agent, research-agent, time-agent\n")
	fmt.Printf("Use natural language to request tasks - the coordinator will transfer to appropriate agents\n")
	fmt.Println(strings.Repeat("=", 70))

	// Enable debug logging if requested.
	if *debug {
		alog.SetLevel(alog.LevelDebug)
		fmt.Println("🪵 Debug logging enabled (internal flow logs at DEBUG level)")
	}

	// Create and run the chat.
	chat := &transferChat{
		modelName:                  *modelName,
		debug:                      *debug,
		endInvocationAfterTransfer: *endInvocation,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// transferChat manages the conversation with agent transfer functionality.
type transferChat struct {
	modelName                  string
	runner                     runner.Runner
	userID                     string
	sessionID                  string
	debug                      bool
	endInvocationAfterTransfer bool
}

// run starts the interactive chat session.
func (c *transferChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with main agent and sub-agents.
func (c *transferChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create sub-agents.

// Create coordinator agent with sub-agents.

// Create runner.

// Setup identifiers.

// createCoordinatorAgent creates the main coordinator agent with sub-agents.
func (c *transferChat) createCoordinatorAgent(modelInstance model.Model, subAgents []agent.Agent) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// startChat runs the interactive conversation loop.
func (c *transferChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *transferChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response with transfer awareness.

// processStreamingResponse handles the streaming response with transfer visualization.
func (c *transferChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Verbose event trace for debugging transfer ordering/author

// Safety check: after transfer, no parent/coordinator events should appear

// Final newline

// handleTransferEvent processes a single event from the transfer system.
func (c *transferChat) handleTransferEvent(
	event *event.Event,
	fullContent *string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	currentAgent *string,
) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle agent transfers.

// Handle tool calls.

// Handle content.

// Handle tool responses.

// handleTransfer processes agent transfer events.
func (c *transferChat) handleTransfer(event *event.Event, currentAgent *string, assistantStarted *bool) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolCalls detects and displays tool calls.
func (c *transferChat) handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// displayToolCalls shows tool call information.
func (c *transferChat) displayToolCalls(event *event.Event) { _ = "STUB: not implemented"; return }

// handleContent processes streaming content.
func (c *transferChat) handleContent(
	event *event.Event,
	fullContent *string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	currentAgent *string,
) {
	_ = "STUB: not implemented"
	return
}

// extractContent extracts content from the choice.
func (c *transferChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	// Only use delta content to avoid duplication in streaming responses
	return ""
}

// displayContent handles content display logic.
func (c *transferChat) displayContent(
	event *event.Event,
	content string,
	fullContent *string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	currentAgent *string,
) {
	_ = "STUB: not implemented"
	return
}

// displayAgentHeader shows the agent header when starting content.
func (c *transferChat) displayAgentHeader(event *event.Event, currentAgent *string) {
	_ = "STUB: not implemented"
	return
}

// handleToolResponses processes tool response completion.
func (c *transferChat) handleToolResponses(event *event.Event) { _ = "STUB: not implemented"; return }

// Helper functions for display formatting.
func (c *transferChat) getAgentIcon(agentName string) string { _ = "STUB: not implemented"; return "" }

func (c *transferChat) getAgentDisplayName(agentName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *transferChat) getAgentFromTransfer(event *event.Event) string {
	_ = "STUB: not implemented"
	// Parse the transfer event to determine target agent.
	return ""
}

func (c *transferChat) isTransferTool(toolCall model.ToolCall) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *transferChat) isTransferResponse(event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// Helper functions.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
