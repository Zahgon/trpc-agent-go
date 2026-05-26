//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main is the main package for the graph activity AG-UI server.
package main

import (
	"context"
	"flag"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/graph/checkpoint/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
	aguiadapter "trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	aguirunner "trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	sessioninmemory "trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

const (
	nodePrepare            = "prepare"
	nodeRecipeCalcLLM      = "recipe_calc_llm"
	nodeExecuteTools       = "execute_tools"
	nodeConfirm            = "confirm"
	nodeDraftMessageLLM    = "draft_message_llm"
	nodePolishMessageAgent = "polish_message_agent"
	nodeFinish             = "finish"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "OpenAI-compatible model name.")
	isStream  = flag.Bool("stream", true, "Whether to stream the response.")
	address   = flag.String("address", "127.0.0.1:8080", "Listen address.")
	path      = flag.String("path", "/agui", "HTTP path.")
)

func main() {
	flag.Parse()

	modelInstance := openai.New(*modelName)
	generationConfig := model.GenerationConfig{
		MaxTokens:   intPtr(512),
		Temperature: floatPtr(0.2),
		Stream:      *isStream,
	}

	g, err := buildGraph(modelInstance, generationConfig)
	if err != nil {
		log.Fatalf("build graph failed: %v", err)
	}

	checkpointSaver := inmemory.NewSaver()
	subAgent := llmagent.New(
		nodePolishMessageAgent,
		llmagent.WithDescription("Polishes the draft recipe message into a clear, friendly message."),
		llmagent.WithModel(modelInstance),
		llmagent.WithGenerationConfig(generationConfig),
		llmagent.WithInstruction("You are a helpful cooking assistant. Rewrite the input as a clear, friendly message."),
	)
	ga, err := graphagent.New(
		"agui-graph-activity",
		g,
		graphagent.WithDescription("AG-UI server that emits activity events for graph node starts."),
		graphagent.WithInitialState(graph.State{}),
		graphagent.WithSubAgents([]agent.Agent{subAgent}),
		graphagent.WithCheckpointSaver(checkpointSaver),
	)
	if err != nil {
		log.Fatalf("create graph agent failed: %v", err)
	}

	sessionService := sessioninmemory.NewSessionService()
	r := runner.NewRunner(ga.Info().Name, ga, runner.WithSessionService(sessionService))
	defer r.Close()

	server, err := agui.New(
		r,
		agui.WithPath(*path),
		agui.WithGraphNodeLifecycleActivityEnabled(true),
		agui.WithGraphNodeInterruptActivityEnabled(true),
		agui.WithAGUIRunnerOptions(
			aguirunner.WithStateResolver(resolveRuntimeState),
		),
		agui.WithMessagesSnapshotEnabled(true),
		agui.WithSessionService(sessionService),
		agui.WithAppName(ga.Info().Name),
	)
	if err != nil {
		log.Fatalf("create AG-UI server failed: %v", err)
	}

	log.Infof("AG-UI: serving agent %q on http://%s%s", ga.Info().Name, *address, *path)
	if err := http.ListenAndServe(*address, server.Handler()); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}

func buildGraph(modelInstance model.Model, generationConfig model.GenerationConfig) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func confirmNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func finishNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type calculatorArgs struct {
	Operation string  `json:"operation" description:"Operation: add, subtract, multiply, divide."`
	A         float64 `json:"a" description:"First number."`
	B         float64 `json:"b" description:"Second number."`
}

type calculatorResult struct {
	Result float64 `json:"result"`
}

func calculator(ctx context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

func resolveRuntimeState(_ context.Context, input *aguiadapter.RunAgentInput) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
