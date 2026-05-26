//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates an interactive, multi‑turn chat using Graph + GraphAgent + Runner.
// It highlights that conversation history persists via the session service across runs,
// and shows tool use with streaming outputs.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model name to use")
)

func main() {
	flag.Parse()
	fmt.Printf("🤖 Graph Multi‑turn Chat (tools + streaming)\n")
	fmt.Printf("Model: %s\n\n", *modelName)
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("💡 Hint: OPENAI_API_KEY is not set. Configure your provider API key/base URL if required.")
	}

	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	_ = "STUB: not implemented"
	// Build graph: simple chat with optional tools.
	return nil
}

// Define a simple calculator tool: only two numbers and an operator.

// Instruction encourages tool use when appropriate.

// If chat produces tool calls → tools; otherwise → End (finish this round)

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive loop using same session for history.

func streamPrint(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// No streaming occurred; emit the message once.

func renderToolCalls(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func renderToolResponses(e *event.Event) bool { _ = "STUB: not implemented"; return false }

func displayToolName(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func prettifyJSON(raw string) string { _ = "STUB: not implemented"; return "" }

func end(_ context.Context, _ graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type calcArgs struct {
	A  float64 `json:"a" description:"First number"`
	B  float64 `json:"b" description:"Second number"`
	Op string  `json:"op" description:"Operator: one of +, -, *, /, ^"`
}

type calcResult struct {
	Result float64 `json:"result"`
}

func calc(_ context.Context, in calcArgs) (calcResult, error) {
	_ = "STUB: not implemented"
	return *new(calcResult), nil
}
