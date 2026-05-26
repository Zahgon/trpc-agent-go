//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates user context propagation and authorization using Invocation State.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	userID    = flag.String("user-id", "alice", "User ID for authentication")
	role      = flag.String("role", "user", "User role (admin, user, guest)")
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
)

func main() {
	// Parse command line flags.
	flag.Parse()

	fmt.Println("🔐 User Context and Authorization Example")
	fmt.Println("This example demonstrates how to use Invocation State for user context propagation and tool authorization.")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	// Create the example.
	example := &userContextExample{
		userID: *userID,
		role:   *role,
	}

	// Setup and run.
	if err := example.run(); err != nil {
		log.Fatalf("Example failed: %v", err)
	}
}

// userContextExample demonstrates user context propagation and authorization.
type userContextExample struct {
	runner    runner.Runner
	userID    string
	role      string
	sessionID string
}

// run executes the user context example.
func (e *userContextExample) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Run the example.

// setup creates the runner with LLM agent and tools.
func (e *userContextExample) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create tools with different permission requirements.

// Create callbacks for authorization and audit.

// Create LLM agent.

// Create runner.

// Setup session.

// createTools creates the tools for the agent.
func (e *userContextExample) createTools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// runExample executes the interactive chat session.
func (e *userContextExample) runExample(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle special commands.

// Process the user message.

// handleCommand handles special commands.
func (e *userContextExample) handleCommand(cmd string) bool {
	_ = "STUB: not implemented"
	return false
}

// processMessage handles a single message exchange.
func (e *userContextExample) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process response.

// processResponse handles the response from the agent.
func (e *userContextExample) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Handle tool calls.

// Handle content.

// Check if this is the final event.

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
