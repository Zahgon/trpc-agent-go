//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a simple two‑node graph: Retrieve → LLM.
// The retrieval node writes ephemeral data into the session's temp namespace,
// and the LLM node uses placeholders to inject that data into its instruction.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
)

const (
	appName          = "graph-retrieval-placeholder"
	defaultModelName = "deepseek-v4-flash"
)

func main() {
	modelName := flag.String("model", defaultModelName, "Model name to use")
	flag.Parse()

	fmt.Println("🔎 Graph Retrieval → LLM (Placeholder Injection)")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println("Type 'exit' to quit")
	fmt.Println(strings.Repeat("=", 60))

	if err := run(context.Background(), *modelName); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, modelName string) error {
	_ = "STUB: not implemented"
	// Session service
	return nil
}

// Create a session with no prefilled state (we'll write temp: keys per turn).

// Build graph: retrieve → llm.

// 1) Retrieval node: simulate recall and write into session.State as temp keys.

// Get the current user input from graph state.

// Simulate a small retrieval based on input.
// In real code, call your vector store or search tool here.

// Write ephemeral keys into the session temp namespace so LLM placeholders can read them.

// No graph state changes are strictly necessary for placeholder injection.
// Return an empty update to keep semantics clear.

// 2) LLM node: instruction uses placeholders from session state (temp namespace).

// Wire: retrieve → answer

// Agent + Runner

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive loop

// fakeRetrieve simulates retrieval results for demonstration purposes.
func fakeRetrieve(query string) []string { _ = "STUB: not implemented"; return nil }

// stream prints streaming events nicely.
func stream(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }
