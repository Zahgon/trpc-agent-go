//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates reasoning/thinking mode using the Runner with streaming output.
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
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	modelName       = flag.String("model", "deepseek-v4-pro", "Name of the model to use")
	streaming       = flag.Bool("streaming", true, "Enable streaming mode for responses")
	thinkingEnabled = flag.Bool("thinking", true, "Enable reasoning/thinking mode if provider supports it")
	thinkingTokens  = flag.Int("thinking-tokens", 2048, "Max reasoning tokens if provider supports it")
	variant         = flag.String("variant", "openai", "Name of Variant to use when use openai provider, openai / hunyuan / deepseek / qwen")
	debug           = flag.Bool("debug", true, "Print messages sent to model API for debugging")
	reasoningMode   = flag.String("reasoning-mode", "discard_previous",
		"How to handle reasoning_content in history: keep_all, discard_previous, discard_all")
)

func main() {
	flag.Parse()

	fmt.Printf("🧠 Thinking Demo (Reasoning)")
	fmt.Printf("\nModel: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Thinking: %t (tokens=%d)\n", *thinkingEnabled, *thinkingTokens)
	fmt.Printf("Reasoning Mode: %s\n", *reasoningMode)
	fmt.Printf("Debug: %t\n", *debug)
	fmt.Println(strings.Repeat("=", 50))

	chat := &thinkingChat{modelName: *modelName, streaming: *streaming, variant: *variant}
	if err := chat.run(context.Background()); err != nil {
		log.Fatalf("Thinking demo failed: %v", err)
	}
}

type thinkingChat struct {
	modelName string
	streaming bool
	runner    runner.Runner
	userID    string
	sessionID string
	appName   string
	sessSvc   session.Service
	variant   string
}

func (c *thinkingChat) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

func (c *thinkingChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Add debug callback if enabled.

// always use in-memory session for this demo

// Add reasoning content mode based on flag.

func (c *thinkingChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *thinkingChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *thinkingChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Show reasoning content.

// Dim style for reasoning content.

// Dim style for reasoning content.

// Show normal content.

// Insert a newline once between reasoning and normal content in streaming mode.

// Print timing information at the end

// printTimingInfo displays timing information from the final event.
func (c *thinkingChat) printTimingInfo(event *event.Event) { _ = "STUB: not implemented"; return }

// For timing info

// Time to first token

// Reasoning duration

// Token usage

func (c *thinkingChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *thinkingChat) showHistory(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Print reasoning (dim) if present in final message.

// Then print visible content.

func (c *thinkingChat) startNewSession() { _ = "STUB: not implemented"; return }

// resolveReasoningMode converts the flag value to llmagent constant.
func resolveReasoningMode() string { _ = "STUB: not implemented"; return "" }

func intPtr(i int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
