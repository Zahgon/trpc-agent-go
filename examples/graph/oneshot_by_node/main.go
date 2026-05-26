//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates one_shot_messages_by_node for parallel branches.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName          = "oneshot-by-node"
	defaultModelName = "deepseek-v4-flash"

	nodeStart = "start"
	nodePrepA = "prep_a"
	nodePrepB = "prep_b"
	nodeLLM1  = "llm1"
	nodeLLM2  = "llm2"

	runtimeStateUserIDKey = "user_id"
)

var (
	modelName = flag.String(
		"model",
		defaultModelName,
		"Name of the model to use",
	)
	q1 = flag.String(
		"q1",
		"What is 1+1? Reply with prefix LLM1:",
		"One-shot user prompt for llm1",
	)
	q2 = flag.String(
		"q2",
		"What is 2+2? Reply with prefix LLM2:",
		"One-shot user prompt for llm2",
	)
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := runOnce(ctx); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func runOnce(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func startNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func prepareLLM1(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func prepareLLM2(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func printSessionState(
	ctx context.Context,
	svc session.Service,
	userID string,
	sessionID string,
) error {
	_ = "STUB: not implemented"
	return nil
}
