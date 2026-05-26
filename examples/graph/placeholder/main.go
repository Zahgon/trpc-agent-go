//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates placeholder usage in a Graph (StateGraph + GraphAgent)
// workflow with session state integration. It mirrors the capabilities from
// examples/placeholder but implemented as a graph with a single LLM node.
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
	defaultModelName = "deepseek-v4-flash"
	appName          = "graph-placeholder-demo"
)

type demo struct {
	modelName      string
	sessionService session.Service
	runner         runner.Runner
	userID         string
	sessionID      string
}

func main() {
	model := flag.String("model", defaultModelName, "Model name to use")
	flag.Parse()

	fmt.Println("🔗 Graph Placeholder Demo")
	fmt.Printf("Model: %s\n", *model)
	fmt.Println("Type 'exit' to end the session")
	fmt.Println("Commands: /set-user-topics <topics>, /set-app-banner <text>, /show-state")
	fmt.Println(strings.Repeat("=", 60))

	d := &demo{modelName: *model}
	if err := d.run(context.Background()); err != nil {
		log.Fatalf("demo failed: %v", err)
	}
}

func (d *demo) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

func (d *demo) initialize(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Session service.
	return nil
}

// Create initial session state with placeholders.

// Unprefixed (readonly) placeholder example

// Prefixed placeholders (mutable via service)

// Build graph with a single LLM node using placeholders in instruction.

// GraphAgent + Runner

func (d *demo) loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *demo) handleSetUserTopics(ctx context.Context, topics string) {
	_ = "STUB: not implemented"
	return
}

func (d *demo) handleSetAppBanner(ctx context.Context, banner string) {
	_ = "STUB: not implemented"
	return
}

func (d *demo) handleShowState(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *demo) ask(ctx context.Context, text string) error { _ = "STUB: not implemented"; return nil }

func (d *demo) stream(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Stream tokens for model responses.

// Stop line on completion.
