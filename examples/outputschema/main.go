//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates output schema functionality using LLMAgent
// with structured output validation. This example shows how to constrain
// agent responses to specific JSON schemas for consistent data formats.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	maxTokens   = 800
	temperature = 0.3
)

func main() {
	// Parse command line flags.
	modelName := flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	flag.Parse()

	fmt.Printf("📋 Output Schema Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
	fmt.Println("💡 Example queries to try:")
	fmt.Println("   • What's the weather like in Beijing today?")
	fmt.Println("   • Tell me about the weather in Shanghai")
	fmt.Println("   • How's the weather in Guangzhou?")
	fmt.Println("   • Weather forecast for Shenzhen")
	fmt.Println("   • What's the climate like in Chengdu?")
	fmt.Println()
	fmt.Println("🔄 How it works:")
	fmt.Println("   1. Agent analyzes your weather query")
	fmt.Println("   2. Returns structured JSON with temperature, conditions, etc.")
	fmt.Println("   3. Output is validated against a predefined schema")
	fmt.Println()

	// Create and run the chat.
	chat := &outputSchemaChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Output schema chat failed: %v", err)
	}
}

// outputSchemaChat manages the output schema conversation.
type outputSchemaChat struct {
	modelName      string
	runner         runner.Runner
	sessionService session.Service
	userID         string
	sessionID      string
}

// run starts the interactive chat session.
func (c *outputSchemaChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner with output schema agent.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with output schema agent.
func (c *outputSchemaChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create session service.

// Create generation config.

// Define output schema for weather information.

// Create Weather Agent with output schema.

// Create runner with the weather agent and session service.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *outputSchemaChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange through the agent.
func (c *outputSchemaChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the weather agent through the runner.

// Process streaming response.

// processStreamingResponse handles the streaming response from the agent.
func (c *outputSchemaChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is the final runner completion event.

// handleEvent processes a single event from the agent.
func (c *outputSchemaChat) handleEvent(event *event.Event) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle streaming content.

// Helper functions.

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
