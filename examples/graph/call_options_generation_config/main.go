//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates per-run GenerationConfig overrides using
// graph call options.
package main

import (
	"context"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	appName = "call-options-generation-config"

	parentAgentName = "parent"
	childAgentName  = "child_agent"

	nodeParentLLM = "parent_llm"
	nodeChildLLM  = "llm"

	userID    = "user"
	userInput = "hello"

	parentTempDefault = 0.8
	childTempDefault  = 0.7

	callTempGlobal = 0.2
	callTempChild  = 0.0

	callMaxTokensParent = 111
	callMaxTokensChild  = 222

	assistantOK = "ok"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func buildChildGraph(m model.Model) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildParentGraph(m model.Model) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCallOptions() agent.RunOption { _ = "STUB: not implemented"; return *new(agent.RunOption) }

func runOnce(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
	runOpts ...agent.RunOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func newSessionID() string { _ = "STUB: not implemented"; return "" }

type printModel struct {
	name string
}

func (m *printModel) GenerateContent(
	ctx context.Context,
	req *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *printModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func fmtInt(p *int) string { _ = "STUB: not implemented"; return "" }

func fmtFloat(p *float64) string { _ = "STUB: not implemented"; return "" }
