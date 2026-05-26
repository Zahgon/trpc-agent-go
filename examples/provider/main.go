//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use Provider with LLM agents.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	providerName        = flag.String("provider", "openai", "Name of the provider to use, openai/anthropic/ollama/hunyuan")
	modelName           = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	isStream            = flag.Bool("stream", true, "Whether to stream the response")
	apiKey              = flag.String("api-key", "", "Override the provider API key")
	baseURL             = flag.String("base-url", "", "Override the provider base URL")
	secretID            = flag.String("secret-id", "", "Set secret id for hunyuan")
	secretKey           = flag.String("secret-key", "", "Set secret key for hunyuan")
	channelBufferSize   = flag.Int("channel-buffer", 0, "Override provider channel buffer size")
	enableTokenTailor   = flag.Bool("token-tailor", false, "Enable provider token tailoring")
	maxTailorInputToken = flag.Int("max-input-tokens", 0, "Maximum input tokens when token tailoring is enabled")
)

func main() {
	// Parse command line flags.
	flag.Parse()
	fmt.Printf("🧠 Provider Agent Demo\n")
	fmt.Printf("Provider: %s\n", *providerName)
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Stream: %t\n", *isStream)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: calculator\n")
	fmt.Printf("The agent will perform mathematical calculations\n")
	fmt.Println(strings.Repeat("=", 60))

	// Create and run the chat.
	chat := &providerChat{
		providerName:      *providerName,
		modelName:         *modelName,
		streaming:         *isStream,
		apiKey:            *apiKey,
		baseURL:           *baseURL,
		channelBufferSize: *channelBufferSize,
		tokenTailoring:    *enableTokenTailor,
		maxInputTokens:    *maxTailorInputToken,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// providerChat manages the conversation with provider.
type providerChat struct {
	providerName      string
	modelName         string
	streaming         bool
	apiKey            string
	baseURL           string
	channelBufferSize int
	tokenTailoring    bool
	maxInputTokens    int
	runner            runner.Runner
	userID            string
	sessionID         string
}

// run starts the interactive chat session.
func (c *providerChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and Provider.
func (c *providerChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create LLM agent with tools.

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *providerChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *providerChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response with Provider awareness.

// processStreamingResponse handles the streaming response with Provider visualization.
func (c *providerChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Process text content with Provider awareness.

// Handle tool responses.

// End the response

// calculate performs mathematical calculations.
func (c *providerChat) calculate(_ context.Context, args *calcArgs) (*calcResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type calcArgs struct {
	Operation string  `json:"operation" description:"The operation to perform,enum=add,enum=subtract,enum=multiply,enum=divide,enum=power"`
	A         float64 `json:"a" description:"First number"`
	B         float64 `json:"b" description:"Second number"`
}

type calcResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

// Helper functions.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
