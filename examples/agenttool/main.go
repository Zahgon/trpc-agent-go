//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use agent tools to wrap agents as tools
// within a larger application.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	agenttool "trpc.group/trpc-go/trpc-agent-go/tool/agent"
)

const (
	defaultModelName     = "deepseek-v4-flash"
	defaultInnerTextMode = string(agenttool.InnerTextModeInclude)
	responseModeDefault  = "default"
	responseModeFinal    = "final-only"
)

var (
	modelName = flag.String(
		"model",
		defaultModelName,
		"Name of the model to use",
	)
	debugAuthors = flag.Bool(
		"debug",
		false,
		"Print event author names with streamed text",
	)
	showTool = flag.Bool(
		"show-tool",
		false,
		"Show tool outputs (tool.response) in the transcript",
	)
	showInner = flag.Bool(
		"show-inner",
		true,
		"Show inner agent transcript forwarded by agent tool",
	)
	innerTextMode = flag.String(
		"inner-text",
		defaultInnerTextMode,
		"Inner text mode: include or exclude",
	)
	toolResponseMode = flag.String(
		"response-mode",
		responseModeDefault,
		"AgentTool response mode: default or final-only",
	)
)

func main() {
	// Parse command line flags.
	flag.Parse()

	mode, err := parseInnerTextMode(*innerTextMode)
	if err != nil {
		log.Fatalf("invalid -inner-text: %v", err)
	}
	responseMode, err := parseResponseMode(*toolResponseMode)
	if err != nil {
		log.Fatalf("invalid -response-mode: %v", err)
	}

	fmt.Printf("🚀 Agent Tool Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Show inner: %t\n", *showInner)
	fmt.Printf("Inner text mode: %s\n", mode)
	fmt.Printf("Response mode: %s\n", responseModeName(responseMode))
	fmt.Printf("Show tool: %t\n", *showTool)
	fmt.Printf("Available tools: current_time, math-specialist(agent_tool)\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &agentToolChat{
		modelName:     *modelName,
		debugAuthors:  *debugAuthors,
		showTool:      *showTool,
		showInner:     *showInner,
		innerTextMode: mode,
		responseMode:  responseMode,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// agentToolChat manages the conversation with agent tools.
type agentToolChat struct {
	modelName     string
	runner        runner.Runner
	userID        string
	sessionID     string
	debugAuthors  bool
	agentName     string
	streaming     bool
	showTool      bool
	showInner     bool
	innerTextMode agenttool.InnerTextMode
	responseMode  agenttool.ResponseMode
}

// run starts the interactive chat session.
func (c *agentToolChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and tools including agent tools.
func (c *agentToolChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create tools.

// Create a specialized agent for math operations.

// Create tools.

// Create agent tool that wraps the math specialist agent.
// SkipSummarization surfaces the child tool result directly.
// The run still finishes on runner.completion.

// Create LLM agent with tools including the agent tool.

// Enable streaming

// Remember streaming mode for printing logic.

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *agentToolChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle special commands.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *agentToolChat) processMessage(
	ctx context.Context,
	userMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response.

// processStreamingResponse handles streaming with tool call visualization.
func (c *agentToolChat) processStreamingResponse(
	eventChan <-chan *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

// handleEvent processes one event and returns true when it was handled.
func (c *agentToolChat) handleEvent(
	ev *event.Event,
	assistantStarted *bool,
	fullContent *strings.Builder,
) bool {
	_ = "STUB: not implemented"
	// Handle errors
	return false
}

// Handle tool calls

// Handle inner agent streaming

// Handle outer assistant streaming

// Handle tool responses

// handleToolCalls processes tool call events.
func (c *agentToolChat) handleToolCalls(
	ev *event.Event,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleInnerAgentStreaming processes inner agent streaming events.
func (c *agentToolChat) handleInnerAgentStreaming(ev *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// handleAssistantStreaming processes outer assistant streaming events.
func (c *agentToolChat) handleAssistantStreaming(
	ev *event.Event,
	assistantStarted *bool,
	fullContent *strings.Builder,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses processes tool response events.
func (c *agentToolChat) handleToolResponses(ev *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// Partial tool delta; only show if inner streaming is hidden.

// Final tool message - show detailed response

// Tool execution completed

// startNewSession creates a new session.
func (c *agentToolChat) startNewSession() { _ = "STUB: not implemented"; return }

func parseInnerTextMode(mode string) (agenttool.InnerTextMode, error) {
	_ = "STUB: not implemented"
	return *new(agenttool.InnerTextMode), nil
}

func parseResponseMode(mode string) (agenttool.ResponseMode, error) {
	_ = "STUB: not implemented"
	return *new(agenttool.ResponseMode), nil
}

func responseModeName(mode agenttool.ResponseMode) string { _ = "STUB: not implemented"; return "" }
