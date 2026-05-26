//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates token usage tracking for each conversation turn
// using the Runner with interactive command line interface.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Printf("🚀 Token Usage Tracker Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Special commands: /stats, /new, /exit\n")
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat.
	chat := &tokenTrackerChat{
		modelName: *modelName,
		streaming: *streaming,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// tokenTrackerChat manages the conversation with token usage tracking.
type tokenTrackerChat struct {
	modelName string
	streaming bool
	runner    runner.Runner
	userID    string
	sessionID string

	// Token usage tracking
	sessionUsage *SessionTokenUsage
	turnCount    int
}

// SessionTokenUsage tracks token usage for the entire session.
type SessionTokenUsage struct {
	TotalPromptTokens     int
	TotalPromptCached     int
	TotalPromptCacheRead  int
	TotalPromptCacheWrite int
	TotalCompletionTokens int
	TotalTokens           int
	TurnCount             int
	UsageHistory          []TurnUsage
}

// TurnUsage represents token usage for a single turn.
type TurnUsage struct {
	TurnNumber        int
	PromptTokens      int
	PromptCached      int
	PromptCacheRead   int
	PromptCacheWrite  int
	CompletionTokens  int
	TotalTokens       int
	Model             string
	InvocationID      string
	Timestamp         time.Time
	UserMessage       string
	AssistantResponse string
}

// run starts the interactive chat session.
func (c *tokenTrackerChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and tools.
func (c *tokenTrackerChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create LLM agent without tools for simplicity.

// Create session service.

// Create runner.

// Setup identifiers and token tracking.

// startChat runs the interactive conversation loop.
func (c *tokenTrackerChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle special commands.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange with token tracking.
func (c *tokenTrackerChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process response and track token usage.

// processResponse handles both streaming and non-streaming responses with token tracking.
func (c *tokenTrackerChat) processResponse(eventChan <-chan *event.Event, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Track token usage from the event.

// Update token usage (use the latest values from the event).

// Handle content for display.

// Check if this is the final event.

// Show turn-specific token usage.

// addTurnUsage adds token usage for a single turn to the session tracking.
func (c *tokenTrackerChat) addTurnUsage(usage TurnUsage) { _ = "STUB: not implemented"; return }

// showStats displays current session token usage statistics.
func (c *tokenTrackerChat) showStats() { _ = "STUB: not implemented"; return }

// showFinalStats displays final statistics when exiting.
func (c *tokenTrackerChat) showFinalStats() { _ = "STUB: not implemented"; return }

// startNewSession creates a new session and resets token tracking.
func (c *tokenTrackerChat) startNewSession() { _ = "STUB: not implemented"; return }

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
