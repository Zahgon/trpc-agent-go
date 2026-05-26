//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent. All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates preparing one_shot_messages_by_node from a single
// upstream node.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName   = "oneshot-by-node-preprocess"
	agentName = "oneshot-by-node-preprocess-agent"

	nodeStart = "start"
	nodePrep  = "preprocess"
	nodeLLM1  = "llm1"
	nodeLLM2  = "llm2"

	defaultUserID = "user"
)

var (
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
	userInput = flag.String(
		"user_input",
		"fallback user input",
		"Fallback user input for the run",
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

func preprocess(ctx context.Context, state graph.State) (any, error) {
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

type echoModel struct{}

func (m *echoModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *echoModel) GenerateContent(
	ctx context.Context,
	req *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func firstMessageByRole(msgs []model.Message, role model.Role) string {
	_ = "STUB: not implemented"
	return ""
}

func lastMessageByRole(msgs []model.Message, role model.Role) string {
	_ = "STUB: not implemented"
	return ""
}
