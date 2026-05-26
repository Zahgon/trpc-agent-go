//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a one-shot user input workflow using the graph
// package with GraphAgent and Runner. It shows how user input is consumed
// exactly once, then cleared from state by the LLM node execution.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	defaultModelName = "deepseek-v4-flash"
	appName          = "userinputonce"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
	inputFlag = flag.String("input", "",
		"User input to process. If empty, read from stdin once")
)

func main() {
	flag.Parse()
	fmt.Printf("🚀 One-shot User Input Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	content := *inputFlag
	if strings.TrimSpace(content) == "" {
		var err error
		content, err = readSingleLine()
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}
	}

	if err := runOnce(content); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func readSingleLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func runOnce(content string) error { _ = "STUB: not implemented"; return nil }

// Build graph with a single LLM node. After execution, user input is
// cleared by the LLM node per graph/state_graph.go behavior.

// Create agent and runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Create user message and run once.

func processEvents(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Print streaming tokens from model.

// When finished, show final response if present.

func truncate(s string, max int) string { _ = "STUB: not implemented"; return "" }

// verifyCleared is a function node that logs whether user_input has been
// cleared after the LLM node. This demonstrates the one-shot behavior in a
// concrete way without inspecting internal state directly.
func verifyCleared(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
