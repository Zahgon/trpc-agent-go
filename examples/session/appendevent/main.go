//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to directly append events to session without
// invoking the model. This is useful for scenarios like:
// - Pre-loading conversation history
// - Inserting system messages or context
// - Recording user actions or metadata
// - Building conversation context from external sources
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
)

func main() {
	flag.Parse()

	fmt.Printf("🚀 AppendEvent Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Println(strings.Repeat("=", 50))

	chat := &appendEventChat{
		modelName: *modelName,
		streaming: *streaming,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

type appendEventChat struct {
	modelName  string
	streaming  bool
	runner     runner.Runner
	sessionSvc session.Service
	userID     string
	sessionID  string
	sessionIDs []string
}

// run starts the interactive chat session.
func (c *appendEventChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner and session service.

// Ensure runner resources are cleaned up.

// Start interactive chat.

func (c *appendEventChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Initialize session service (in-memory).

// startChat runs the interactive conversation loop.
func (c *appendEventChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle commands or process as normal message.

// Normal chat message - process through runner.

// printCommands displays available commands.
func (c *appendEventChat) printCommands() { _ = "STUB: not implemented"; return }

// handleCommand processes user commands.
// Returns: (handled, shouldExit, error)
func (c *appendEventChat) handleCommand(ctx context.Context, userInput string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// handleAppendSystem handles /append-system command.
func (c *appendEventChat) handleAppendSystem(ctx context.Context, userInput string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// handleAppendAssistant handles /append-assistant command.
func (c *appendEventChat) handleAppendAssistant(ctx context.Context, userInput string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// handleAppendUser handles /append command.
func (c *appendEventChat) handleAppendUser(ctx context.Context, userInput string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// handleUseSession handles /use command.
func (c *appendEventChat) handleUseSession(userInput string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// appendUserMessage appends a user message directly to session.
func (c *appendEventChat) appendUserMessage(ctx context.Context, content string) error {
	_ = "STUB: not implemented"
	return nil
}

// appendSystemMessage appends a system message directly to session.
func (c *appendEventChat) appendSystemMessage(ctx context.Context, content string) error {
	_ = "STUB: not implemented"
	return nil
}

// appendAssistantMessage appends an assistant message directly to session.
func (c *appendEventChat) appendAssistantMessage(ctx context.Context, content string) error {
	_ = "STUB: not implemented"
	return nil
}

// appendMessageToSession appends a message as an event to the session.
//
// Note: An Event can represent both user requests and model responses.
// - User messages: Created when users send messages (author: "user")
// - Assistant messages: Created when model generates responses (author: agent name)
// - System messages: Created for system instructions (author: "system")
//
// Required fields for creating an event:
//   - invocationID: Unique identifier for this invocation (required)
//   - author: Event author, e.g., "user", "system", or agent name (required)
//   - response: *model.Response with at least Choices containing Message (required)
//
// Auto-generated fields (by event.NewResponseEvent):
//   - ID: Auto-generated UUID
//   - Timestamp: Auto-set to current time
//   - Version: Auto-set to CurrentVersion
//
// For persistence to session, Response must satisfy:
//   - Response != nil
//   - !IsPartial (or has StateDelta)
//   - IsValidContent() returns true (Choices with Message.Content, Message.ContentParts, or tool calls).
//
// Optional but recommended fields:
//   - RequestID: For request tracking (set manually)
//   - FilterKey: For event filtering (auto-set by framework)
func (c *appendEventChat) appendMessageToSession(
	ctx context.Context,
	message model.Message,
	author string,
) error {
	_ = "STUB: not implemented"
	// Get or create session.
	return nil
}

// Create event from message.
// Required: invocationID, author, response with Choices

// Required: unique invocation identifier
// Required: event author

// Required: Response with Choices
// Recommended: false for non-final events

// Required: choice index
// Required: message with Content or ContentParts

// Optional Response fields:
// Object:    "",  // Optional: object type
// Created:   0,   // Optional: creation timestamp
// Model:     "",  // Optional: model name
// IsPartial: false, // Required: must be false for persistence

// Optional: Set RequestID for tracking

// Persist event to session.

// listEvents displays all events in the current session.
func (c *appendEventChat) listEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

// processMessage handles a single message exchange.
func (c *appendEventChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.
// The runner will automatically load all previously appended events
// from the session and include them in the conversation context.

// Process response.
