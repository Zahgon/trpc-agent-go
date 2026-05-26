//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates node-level Retry/Backoff with an unstable function node.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	failCount = flag.Int("fail", 2, "Number of initial failures for the unstable node")
	latency   = flag.Duration("latency", 200*time.Millisecond, "Simulated latency per attempt")
	verbose   = flag.Bool("verbose", false, "Enable verbose event logging")
)

func main() {
	flag.Parse()
	fmt.Println("🔁 Graph Retry/Backoff Example")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	_ = "STUB: not implemented"
	// Build graph
	return nil
}

// Create GraphAgent

// Create runner

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive loop

func createGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// Unstable node: fails N-1 times then succeeds.
// Use a retry policy that matches any error for demo purposes.

// LLM answer node: summarizes fetched data and answers the user.

// Wiring

// attemptTracker stores per-invocation attempt counts for the unstable node.
var attemptTracker sync.Map // key: invocationID+":"+nodeID -> int

func unstableAPINode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Simulate external call that may fail a few times
	return *new(any), nil
}

// Simulated latency

// on success, cleanup tracker

// Success: return fetched data stored in state

func handleStreaming(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Show retry metadata when verbose

// Stream model deltas
