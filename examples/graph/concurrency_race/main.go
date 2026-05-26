//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main provides a minimal graph example to expose the current
// concurrency issue when reusing a shared GraphAgent/Executor across
// concurrent invocations.
//
// The graph is intentionally simple:
//
//	start -> worker
//
// Node "worker" increments a per-run counter field. When the executor is
// reused concurrently and channels are shared at the Graph level, some
// runs will skip the "worker" node entirely, leaving the counter at 0.
// After we refactor channel state into per-execution context, all runs
// should consistently see counter == 1.
package main

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	stateKeyCounter = "counter"

	defaultUserID      = "user-concurrency-demo"
	defaultConcurrency = 32
	defaultRounds      = 64
)

func main() {
	ctx := context.Background()

	r, err := newConcurrentRunner()
	if err != nil {
		panic(fmt.Errorf("failed to create runner: %w", err))
	}
	defer r.Close()

	fmt.Printf("🚀 concurrency_race example: %d goroutines × %d rounds\n",
		defaultConcurrency, defaultRounds)

	failures := runConcurrentInvocations(ctx, r, defaultConcurrency, defaultRounds)
	if len(failures) == 0 {
		fmt.Println("✅ No missing worker executions observed (try increasing rounds/concurrency if needed).")
		return
	}

	fmt.Printf("❌ Detected %d runs where worker node did not execute:\n", len(failures))
	for _, f := range failures {
		fmt.Println("   -", f)
	}
}

// newConcurrentRunner builds a minimal graph, wraps it in a GraphAgent and
// Runner, and returns the Runner instance for shared reuse across goroutines.
func newConcurrentRunner() (runner.Runner, error) {
	_ = "STUB: not implemented"
	return *new(runner.Runner), nil
}

// buildTestGraph constructs:
//
//	start -> worker
//
// The worker node increments stateKeyCounter by 1. With correct per-execution
// channel isolation, every run should see counter == 1 at completion.
func buildTestGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// No-op; the edge start -> worker is what matters for the branch channel.

// Static edge start -> worker. The underlying implementation will create
// a shared branch channel "branch:to:worker", which is where the race
// manifests when multiple executions share the same Graph channels.

// runConcurrentInvocations executes many runs concurrently against the same
// Runner. It returns a slice of human-readable failure descriptions where
// the worker node did not increment the counter.
func runConcurrentInvocations(
	ctx context.Context,
	r runner.Runner,
	concurrency, rounds int,
) []string {
	_ = "STUB: not implemented"
	return nil
}

// runSingle executes a single run and returns the final counter value observed
// from the graph.execution completion event.
func runSingle(
	ctx context.Context,
	r runner.Runner,
	userID, sessionID string,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
