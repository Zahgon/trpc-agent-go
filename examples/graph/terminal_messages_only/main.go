//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"fmt"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	parentAgentName = "parent"
	draftAgentName  = "draft"
	finalAgentName  = "final"

	nodeEmit = "emit"

	userInputText = "Show me the final answer only."

	draftChunk = "draft: collecting context"
	draftFinal = "draft: collecting context"
	finalChunk = "final: user-visible answer"
	finalFinal = "final: user-visible answer"
)

func main() {
	parentAgent, err := buildParentAgent()
	if err != nil {
		log.Fatalf("build parent agent failed: %v", err)
	}

	if err := runDemo(parentAgent, false); err != nil {
		log.Fatalf("run default demo failed: %v", err)
	}
	fmt.Println()
	if err := runDemo(parentAgent, true); err != nil {
		log.Fatalf("run terminal-only demo failed: %v", err)
	}
}

func buildParentAgent() (*graphagent.GraphAgent, error) { _ = "STUB: not implemented"; return nil, nil }

func buildChildAgent(
	name string,
	chunk string,
	final string,
) (*graphagent.GraphAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func emitAssistantMessage(
	state graph.State,
	agentName string,
	chunk string,
	final string,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func runDemo(
	parentAgent *graphagent.GraphAgent,
	terminalOnly bool,
) error {
	_ = "STUB: not implemented"
	return nil
}
