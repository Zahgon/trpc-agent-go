//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a React-style graph agent.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

const (
	nodePlanner      = "planner"
	nodeReasoning    = "reasoning"
	nodeTool         = "tool"
	nodeFinalAnswer  = "finalanswer"
	nodeFormatOutput = "formatoutput"
)

const (
	plannerInstruction = `You are the planning stage of a React-style agent.
Create a concise numbered plan that explains how to solve the user's question.
Do not give the answer inside the plan.`

	reasoningInstruction = `You are the reasoning stage of a React-style agent.
Follow the planner output and think step by step.
- Clearly describe your next thought in natural language.
- If a tool call is required, explain why and then output exactly one JSON tool call.
- If no tool is needed, write two short sentences: first describe the reasoning, then state the conclusion.
Do not address the user directly.`

	finalInstruction = `You provide the final answer to the user.
Use the conversation context to craft a helpful response in English.`
)

func main() {
	modelName := flag.String("model", "deepseek-v4-flash", "LLM model name")
	question := flag.String("question", "", "User question; leave empty to type interactively")
	flag.Parse()

	prompt := strings.TrimSpace(*question)
	if prompt == "" {
		prompt = readQuestionFromStdin()
	}

	ctx := context.Background()
	graphDef, err := buildGraph(*modelName)
	if err != nil {
		log.Fatalf("failed to build graph: %v", err)
	}

	agent, err := graphagent.New(
		"react-graph",
		graphDef,
		graphagent.WithDescription("Planner → Reasoning → Tool → FinalAnswer → FormatOutput example"),
	)
	if err != nil {
		log.Fatalf("failed to create graph agent: %v", err)
	}

	sessionSvc := inmemory.NewSessionService()
	r := runner.NewRunner("graph-react-example", agent, runner.WithSessionService(sessionSvc))
	defer r.Close()

	userID := "demo-user"
	sessionID := fmt.Sprintf("react-example-%d", time.Now().Unix())

	message := model.NewUserMessage(prompt)
	eventChan, err := r.Run(ctx, userID, sessionID, message)
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	if err := streamEvents(eventChan); err != nil {
		log.Fatalf("stream failed: %v", err)
	}
}

func readQuestionFromStdin() string { _ = "STUB: not implemented"; return "" }

func buildGraph(modelName string) (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

func streamEvents(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

func maybeStartNode(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func extractNodeMetadata(evt *event.Event) (graph.NodeExecutionMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(graph.NodeExecutionMetadata), false
}

func extractDisplayText(evt *event.Event) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func collectToolCalls(evt *event.Event) []string { _ = "STUB: not implemented"; return nil }

func toolResultText(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func formatOutput(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func readNodeResponse(state graph.State, nodeID string) string {
	_ = "STUB: not implemented"
	return ""
}

func nodeLabel(nodeID string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func calculator(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

type calculatorArgs struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

func extractFormatPayload(evt *event.Event) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
