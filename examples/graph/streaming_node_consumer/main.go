//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates node-to-node streaming inside a GraphAgent.
//
// Goal:
//   - An upstream LLM node streams deltas.
//   - A downstream consumer node reads those deltas in real time.
//
// Key APIs:
//   - graph.WithStreamOutput(streamName): make an LLM/Agent node publish
//     deltas
//   - graph.OpenStreamReader(ctx, streamName): read the stream in another
//     node
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	sessioninmemory "trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

const (
	appName        = "streaming-node-consumer-demo"
	graphAgentName = "streaming-node-consumer"

	userID    = "user"
	sessionID = "session"

	nodeSetup   = "setup"
	nodeLLM     = "llm"
	nodeConsume = "consume"
	nodeFinish  = "finish"

	envOpenAIKey   = "OPENAI_API_KEY"
	envModelName   = "MODEL_NAME"
	envOpenAIModel = "OPENAI_MODEL"

	defaultModel  = "gpt-5"
	defaultLines  = 6
	defaultPrompt = "Write a 6-line welcome script for a podcast."

	streamNameLLM = "llm:deltas"

	stateKeyParsedLines = "parsed_lines"
)

var (
	promptFlag = flag.String(
		"prompt",
		defaultPrompt,
		"User prompt for the LLM",
	)
	modelFlag = flag.String(
		"model",
		defaultModelName(),
		"LLM model name",
	)
	linesFlag = flag.Int(
		"lines",
		defaultLines,
		"Expected number of lines in the response",
	)
	printLLMFlag = flag.Bool(
		"print-llm",
		false,
		"Print raw LLM deltas from event stream",
	)
)

func main() {
	flag.Parse()

	if *linesFlag <= 0 {
		log.Fatalf("invalid -lines: %d", *linesFlag)
	}

	if os.Getenv(envOpenAIKey) == "" {
		fmt.Printf("%s is not set.\n", envOpenAIKey)
		fmt.Println("Export it if your gateway requires it.")
	}

	g, err := buildGraph(*modelFlag, *linesFlag)
	if err != nil {
		log.Fatalf("build graph: %v", err)
	}

	ga, err := graphagent.New(graphAgentName, g)
	if err != nil {
		log.Fatalf("create graph agent: %v", err)
	}

	sessSvc := sessioninmemory.NewSessionService()
	r := runner.NewRunner(appName, ga, runner.WithSessionService(sessSvc))
	defer r.Close()

	eventCh, err := r.Run(
		context.Background(),
		userID,
		sessionID,
		model.NewUserMessage(*promptFlag),
		agent.WithStreamMode(
			agent.StreamModeMessages,
			agent.StreamModeCustom,
		),
	)
	if err != nil {
		log.Fatalf("runner run: %v", err)
	}

	for e := range eventCh {
		if e == nil || e.Response == nil {
			continue
		}
		switch e.Object {
		case model.ObjectTypeChatCompletionChunk:
			if !*printLLMFlag {
				continue
			}
			if len(e.Choices) > 0 {
				fmt.Print(e.Choices[0].Delta.Content)
			}

		case graph.ObjectTypeGraphNodeCustom:
			msg, ok := parseNodeCustomText(e)
			if !ok {
				continue
			}
			fmt.Printf("[consume] %s\n", msg)

		case model.ObjectTypeRunnerCompletion:
			if e.IsRunnerCompletion() {
				fmt.Println("done")
			}
		}
	}
}

func defaultModelName() string { _ = "STUB: not implemented"; return "" }

func buildGraph(modelName string, lines int) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scriptInstruction(lines int) string { _ = "STUB: not implemented"; return "" }

func setupNode(ctx context.Context, _ graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func consumeNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func finishNode(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func parseNodeCustomText(e *event.Event) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
