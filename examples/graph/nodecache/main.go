//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates node-level caching in a graph using Runner + GraphAgent
// with interactive input and streaming outputs.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/graph"
)

var (
	ttlSeconds = flag.Int("ttl", 60, "Cache TTL (seconds) for the compute node")
)

func main() {
	flag.Parse()
	fmt.Printf("🚀 Node Cache Example (Runner + GraphAgent)\n")
	fmt.Printf("Cache TTL: %ds\n", *ttlSeconds)
	fmt.Println(strings.Repeat("=", 50))

	if err := run(); err != nil {
		log.Fatalf("example failed: %v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// Build a simple graph with a cached compute node.

// Create GraphAgent.

// Runner with in-memory (RAM) session service.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive loop: user enters a number; we run the graph with that input.

// Runner.Run signature: <userID, sessionID, message, options...>.

// To demonstrate cache hits deterministically, use a new session per run
// so previous outputs do not pollute the cache key via state. The cache
// itself is graph-level so it will still hit across sessions.

// Stream events

func buildGraph(ttl time.Duration) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Schema: input n:int, output out:int
	return nil, nil
}

// Slow compute node: simulates expensive work, doubles input.

// Simulate expensive computation

// Build graph with cache backend and policy

// Use a node-level policy with TTL and a simple field-based cache key.

func decodeStateDelta(delta map[string][]byte) graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}
