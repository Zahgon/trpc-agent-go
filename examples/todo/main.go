//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the todo_write tool.
//
// After each assistant turn the current checklist is read from the
// session state and pretty-printed to the terminal so you can see how
// the agent plans, updates status, and eventually finishes a multi-step
// task.
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
	"trpc.group/trpc-go/trpc-agent-go/tool/todo"
)

var (
	modelName = flag.String("model", "deepseek-chat", "Name of the model to use (any OpenAI-compatible identifier)")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
	variant   = flag.String("variant", "openai", "Variant passed to the OpenAI provider")
	seed      = flag.String("seed", "", "Optional first user message; if set, the chat auto-starts with it")
)

const (
	appName   = "todo-demo"
	agentName = "todo-assistant"
)

func main() {
	flag.Parse()

	fmt.Println("Todo tool demo: plan + track multi-step work")
	fmt.Printf("Model:     %s (variant=%s)\n", *modelName, *variant)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Println("Commands:  /exit to quit, /list to print the current checklist")
	fmt.Println(strings.Repeat("=", 60))

	c := &chat{modelName: *modelName, streaming: *streaming, variant: *variant}
	if err := c.run(); err != nil {
		log.Fatalf("chat failed: %v", err)
	}
}

// chat wraps the runner, session service and conversation loop.
type chat struct {
	modelName string
	streaming bool
	variant   string

	runner    runner.Runner
	sessSvc   session.Service
	userID    string
	sessionID string
}

func (c *chat) run() error { _ = "STUB: not implemented"; return nil }

func (c *chat) setup() error { _ = "STUB: not implemented"; return nil }

// Build the todo tool. The verification nudge hook demonstrates how
// to attach an extra policy reminder without touching the tool core:
// when the model closes out 3+ tasks all-at-once we remind it to do
// a verification pass before declaring success.

// Keep the final "all completed" state visible in the demo.
// The default behavior (true) clears the list to signal a
// fresh planning session next time.

func (c *chat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *chat) processMessage(ctx context.Context, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

// processResponse renders the stream: tool calls, tool responses and
// assistant text are each formatted distinctly so the checklist updates
// are easy to spot.
func (c *chat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Tool calls - always show the arguments so you can watch the
// model write the checklist.

// Tool responses. We render two views side-by-side to show
// both consumption patterns a real deployment will use:
//
//   1. The raw JSON — what the LLM sees (includes the nudge).
//   2. A decoded todo.Output — what a cloud frontend (AG-UI,
//      Web UI, etc.) consumes directly from the event stream,
//      without any extra session fetch.

// Assistant text (streaming delta or full).

// printTodos reads the current checklist from the session at the end of
// a turn and renders it as ASCII.
//
// Two consumption patterns exist and this demo shows both:
//
//  1. In-stream, structured: decode todo.Output from the tool-result
//     event (see processResponse above). This is what a cloud
//     frontend (AG-UI, WebSocket UI, etc.) will use — no extra fetch.
//  2. End-of-turn, canonical: read the session and call todo.GetTodos.
//     This is what a REST endpoint or an audit job would use when it
//     only needs the current state, not the change stream.
func (c *chat) printTodos() { _ = "STUB: not implemented"; return }

// The tool writes under a branch-scoped key. A single-agent setup
// uses the agent name as the branch (see agent.Invocation.Branch).

// formatTodos is a tiny ASCII pretty-printer local to this demo. Rich
// frontends (AG-UI, web UIs, etc.) should render todo.Item directly in
// their own native style instead of reusing a fixed string like this.
func formatTodos(todos []todo.Item) string { _ = "STUB: not implemented"; return "" }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }
