//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates per-node named ends (multi-ends) in the graph package.
// A decision node returns symbolic branches (e.g., "approve", "reject") which
// are resolved via node-local ends to concrete destinations.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

const (
	nodeStart    = "start"
	nodeDecide   = "decide"
	nodeApproved = "approved"
	nodeRejected = "rejected"
	nodeFinal    = "final"

	keyDecision = "decision"
	keyPath     = "path"
	keyResult   = "result"
)

var (
	choice = flag.String("choice", "approve", "Branch choice: approve|reject")
)

func main() {
	flag.Parse()
	fmt.Println("🚀 Multi-Ends Branching Example")

	// Build graph
	g, err := buildGraph()
	if err != nil {
		log.Fatalf("failed to build graph: %v", err)
	}

	// Create a GraphAgent
	ga, err := graphagent.New(
		"multiends-demo",
		g,
		graphagent.WithDescription("Demonstration of per-node named ends (multi-ends)"),
		graphagent.WithInitialState(graph.State{}),
	)
	if err != nil {
		log.Fatalf("failed to create graph agent: %v", err)
	}

	// Create session service and runner
	sess := inmemory.NewSessionService()
	r := runner.NewRunner("multiends-app", ga, runner.WithSessionService(sess))

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer r.Close()

	// Run a single turn using the provided choice as user input
	if err := runOnce(r, *choice); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func buildGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Define a simple schema
	return nil, nil
}

// Add nodes

// Topology

// startNode reads StateKeyUserInput (provided by the runner) and writes it into
// a decision key for the decide node to consume.
func startNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// default choice

// decideNode returns a symbolic branch using Command.GoTo.
func decideNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Unknown choice: default to reject

func approvedNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func rejectedNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func finalNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func runOnce(r runner.Runner, userInput string) error { _ = "STUB: not implemented"; return nil }

// Print error and continue to drain the channel

// Extract final state snapshot from the terminal event
