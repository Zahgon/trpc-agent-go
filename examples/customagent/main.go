//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates implementing a custom Agent by hand without Graph.
// It performs a simple intent classification and branches the logic:
// - chitchat: reply conversationally
// - task: provide a short actionable plan
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
)

func main() {
	flag.Parse()

	fmt.Printf("🚀 Custom Agent (intent-branching)\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	// Build model and custom agent.
	m := openai.New(*modelName)
	ag := NewSimpleIntentAgent(
		"biz-agent",
		"A custom agent demonstrating business flow branching by intent",
		m,
	)

	// Use Runner for session + event handling.
	r := runner.NewRunner("customagent-app", ag)

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer r.Close()

	ctx := context.Background()

	chat := &interactiveChat{
		runner:    r,
		modelName: *modelName,
		userID:    "user",
		sessionID: fmt.Sprintf("custom-session-%d", time.Now().Unix()),
	}

	if err := chat.start(ctx); err != nil {
		log.Fatalf("chat failed: %v", err)
	}
}

type interactiveChat struct {
	runner    runner.Runner
	modelName string
	userID    string
	sessionID string
}

func (c *interactiveChat) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *interactiveChat) handle(ctx context.Context, text string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *interactiveChat) startNewSession() { _ = "STUB: not implemented"; return }

func printContent(evt *event.Event) { _ = "STUB: not implemented"; return }

// Default streaming: print only delta to avoid duplicating final content.

func isToolLike(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// Minimal check: tool calls or tool role messages.
