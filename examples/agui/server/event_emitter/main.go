//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use EventEmitter to emit custom events from NodeFunc.
// This example shows:
// - Emitting custom events with payload
// - Emitting progress events during long-running operations
// - Emitting streaming text events
// - AGUI Server integration to receive these events as AG-UI protocol events
package main

import (
	"context"
	"flag"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
)

const (
	nodeStart    = "start"
	nodeProcess  = "process"
	nodeAnalyze  = "analyze"
	nodeComplete = "complete"
)

var (
	address = flag.String("address", "127.0.0.1:8080", "Listen address")
	path    = flag.String("path", "/agui", "HTTP path")
)

func main() {
	flag.Parse()

	// Build the graph with event emitter demonstration
	g, err := buildGraph()
	if err != nil {
		log.Fatalf("Failed to build graph: %v", err)
	}

	// Create GraphAgent
	ga, err := graphagent.New(
		"event-emitter-demo",
		g,
		graphagent.WithDescription("Demonstration of Node EventEmitter functionality"),
		graphagent.WithInitialState(graph.State{}),
	)
	if err != nil {
		log.Fatalf("Failed to create graph agent: %v", err)
	}

	// Create runner
	r := runner.NewRunner(ga.Info().Name, ga)
	defer r.Close()

	// Create AG-UI server
	server, err := agui.New(r, agui.WithPath(*path))
	if err != nil {
		log.Fatalf("Failed to create AG-UI server: %v", err)
	}

	log.Infof("🚀 Starting AG-UI server with EventEmitter demo at http://%s%s", *address, *path)
	log.Info("📝 This example demonstrates:")
	log.Info("   - Custom events with payload (workflow.started, workflow.completed)")
	log.Info("   - Progress events (node.progress)")
	log.Info("   - Streaming text events (node.text)")
	log.Info("")
	log.Info("💡 Run the client example to test:")
	log.Info("   go run ./client/event_emitter")

	if err = http.ListenAndServe(*address, server.Handler()); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}

// buildGraph creates a graph that demonstrates EventEmitter usage in NodeFunc.
func buildGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// Node 1: Start - emit custom event with initial status

// Node 2: Process - emit progress events during processing

// Node 3: Analyze - emit streaming text events

// Node 4: Complete - emit final custom event

// Set up edges

// startNode demonstrates emitting a custom event with payload.
func startNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Get EventEmitter from state

// Get user input from messages

// Emit custom event: workflow started

// Note: Do not return error when event sending fails,
// to prevent client disconnection from affecting Agent workflow

// processNode demonstrates emitting progress events during a long-running operation.
func processNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate a long-running process with progress updates

// Check context cancellation

// Calculate progress percentage

// Emit progress event

// Simulate work

// analyzeNode demonstrates emitting streaming text events.
func analyzeNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simulate streaming analysis output

// Check context cancellation

// Emit streaming text event

// trim newline for log

// Simulate streaming delay

// completeNode demonstrates emitting a final custom event with results.
func completeNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Emit custom event: workflow completed

// Simulated duration

// Also emit a final progress event to indicate 100% complete

// Ensure we implement agent.Agent interface requirements
var _ agent.Agent = (*graphagent.GraphAgent)(nil)
