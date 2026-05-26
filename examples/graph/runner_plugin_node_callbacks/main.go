//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to apply graph NodeCallbacks from a
// runner-scoped plugin.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	appName     = "runner-plugin-node-callbacks"
	agentName   = "demo-graph"
	userID      = "demo-user"
	sessionPref = "demo-session"

	nodeUpper  = "upper"
	nodeAnswer = "answer"

	stateKeyUpper = "upper_text"

	answerPrefix = "Uppercase: "
	emptyInput   = "(empty)"

	pluginName = "inject_node_callbacks"
)

var input = flag.String(
	"input",
	"hello plugin hooks",
	"User input for the graph",
)

func main() {
	flag.Parse()

	g, err := buildGraph()
	if err != nil {
		log.Fatalf("build graph: %v", err)
	}

	ga, err := graphagent.New(agentName, g)
	if err != nil {
		log.Fatalf("create graph agent: %v", err)
	}

	cb := newNodeLoggerCallbacks()
	p := &nodeCallbacksPlugin{callbacks: cb}

	r := runner.NewRunner(
		appName,
		ga,
		runner.WithPlugins(p),
	)
	defer r.Close()

	ctx := context.Background()
	sessionID := fmt.Sprintf("%s-%d", sessionPref, time.Now().Unix())

	events, err := r.Run(
		ctx,
		userID,
		sessionID,
		model.NewUserMessage(*input),
		agent.WithStreamMode(agent.StreamModeTasks),
	)
	if err != nil {
		log.Fatalf("run: %v", err)
	}

	completion := waitRunnerCompletion(events)
	if err := printCompletion(completion); err != nil {
		log.Fatalf("print result: %v", err)
	}
}

func buildGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

func upperNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func answerNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func newNodeLoggerCallbacks() *graph.NodeCallbacks { _ = "STUB: not implemented"; return nil }

func beforeNodeLog(
	ctx context.Context,
	cbCtx *graph.NodeCallbackContext,
	state graph.State,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func afterNodeLog(
	ctx context.Context,
	cbCtx *graph.NodeCallbackContext,
	state graph.State,
	result any,
	nodeErr error,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func onNodeErrorLog(
	ctx context.Context,
	cbCtx *graph.NodeCallbackContext,
	state graph.State,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

type nodeCallbacksPlugin struct {
	callbacks *graph.NodeCallbacks
}

func (p *nodeCallbacksPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *nodeCallbacksPlugin) Register(reg *plugin.Registry) { _ = "STUB: not implemented"; return }

func (p *nodeCallbacksPlugin) beforeAgent(
	ctx context.Context,
	args *agent.BeforeAgentArgs,
) (*agent.BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func waitRunnerCompletion(events <-chan *event.Event) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func printCompletion(e *event.Event) error { _ = "STUB: not implemented"; return nil }

func assistantText(e *event.Event) (string, bool) { _ = "STUB: not implemented"; return "", false }

func stateString(delta map[string][]byte, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
