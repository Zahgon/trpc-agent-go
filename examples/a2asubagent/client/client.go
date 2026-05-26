//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides an A2A (Agent-to-Agent) client example.
package main

import (
	"flag"
	"fmt"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/a2aagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model to use")
)

var agentURLS = []string{
	"http://localhost:8087/",
	"http://localhost:8088/",
}

func main() {
	flag.Parse()

	// Create remote agents as sub-agents
	remoteAgents := make([]agent.Agent, 0)
	for _, url := range agentURLS {
		agent, err := a2aagent.New(a2aagent.WithAgentCardURL(url))
		if err != nil {
			log.Printf("Failed to create agent from %s: %v", url, err)
			continue
		}
		remoteAgents = append(remoteAgents, agent)
		fmt.Printf("Connected to sub-agent: %s\n", agent.Info().Name)
	}

	if len(remoteAgents) == 0 {
		log.Fatal("No remote agents available")
	}

	// Create coordinator agent with sub-agents
	coordinatorAgent := buildCoordinatorAgent(remoteAgents)

	// Start chat
	startChat(coordinatorAgent)
}

func buildCoordinatorAgent(subAgents []agent.Agent) agent.Agent {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return *new(agent.Agent)
}

// Create LLM agent with sub-agents.

func startChat(coordinatorAgent agent.Agent) {
	_ = "STUB: not implemented"
	// Display coordinator agent info
	return
}

// Create runner for coordinator agent

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Add spacing between turns

func processMessage(coordinatorRunner runner.Runner, userID string, sessionID *string) error {
	_ = "STUB: not implemented"
	return nil
}

// Process with coordinator agent only

func startNewSession() string { _ = "STUB: not implemented"; return "" }

// processResponse handles streaming responses from coordinator agent.
func processResponse(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Final newline

// handleTransferEvent processes a single event from the transfer system.
func handleTransferEvent(
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

// handleToolCalls detects and displays tool calls.
func handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses processes tool response completion.
func handleToolResponses(event *event.Event) { _ = "STUB: not implemented"; return }

// handleContent processes streaming content.
func handleContent(
	event *event.Event,
	fullContent *string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	currentAgent *string,
) {
	_ = "STUB: not implemented"
	return
}

// extractContent extracts content based on streaming mode.
func extractContent(choice model.Choice) string { _ = "STUB: not implemented"; return "" }

// displayContent handles content display logic.
func displayContent(
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

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 {
	_ = "STUB: not implemented"

	// handleTransfer processes agent transfer events.
	return nil
}

func handleTransfer(event *event.Event, currentAgent *string, assistantStarted *bool) bool {
	_ = "STUB: not implemented"
	return false
}

// displayToolCalls shows tool call information.
func displayToolCalls(event *event.Event) { _ = "STUB: not implemented"; return }

// displayAgentHeader shows the agent header when starting content.
func displayAgentHeader(event *event.Event, currentAgent *string) {
	_ = "STUB: not implemented"
	return
}

// Helper functions for display formatting.
func getAgentIcon(agentName string) string { _ = "STUB: not implemented"; return "" }

func getAgentDisplayName(agentName string) string { _ = "STUB: not implemented"; return "" }

func getAgentFromTransfer(event *event.Event) string {
	_ = "STUB: not implemented"
	// Parse the transfer event to determine target agent.
	return ""
}

func isTransferTool(toolCall model.ToolCall) bool { _ = "STUB: not implemented"; return false }

func isTransferResponse(event *event.Event) bool { _ = "STUB: not implemented"; return false }
