//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a minimal multi-turn chat powered by Runner.
// It focuses on core control flow with an in-memory session backend so the
// example stays self-contained and easy to run.
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
)

var (
	modelName      = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming      = flag.Bool("streaming", true, "Enable streaming mode for responses")
	enableParallel = flag.Bool("enable-parallel", false, "Enable parallel tool execution (default: false, serial execution)")
	variant        = flag.String("variant", "openai", "Name of the variant to use when calling the OpenAI provider")
)

const (
	appName   = "runner-quickstart"
	agentName = "chat-assistant"
)

func main() {
	flag.Parse()

	fmt.Printf("🚀 Runner quickstart: multi-turn chat with tools\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Parallel tools: %t\n", *enableParallel)
	fmt.Printf("Session backend: in-memory (simple demo)\n")
	fmt.Printf("Type '/exit' to end the conversation\n")
	fmt.Printf("Available tools: calculator, current_time\n")
	fmt.Println(strings.Repeat("=", 50))

	chat := &multiTurnChat{
		modelName: *modelName,
		streaming: *streaming,
		variant:   *variant,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("chat failed: %v", err)
	}
}

// multiTurnChat manages the conversation loop for the demo.
type multiTurnChat struct {
	modelName string
	streaming bool
	runner    runner.Runner
	userID    string
	sessionID string
	variant   string
}

func (c *multiTurnChat) run() error { _ = "STUB: not implemented"; return nil }

// setup builds the runner with a model, tools, and the in-memory session store.
func (c *multiTurnChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *multiTurnChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *multiTurnChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *multiTurnChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *multiTurnChat) handleEvent(
	evt *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *multiTurnChat) handleToolCalls(
	evt *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChat) handleToolResponses(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChat) handleContent(
	evt *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *multiTurnChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *multiTurnChat) displayContent(
	content string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}
