//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	nodeInternalCallLLM   = "internal_call_llm"
	nodeInternalTool      = "internal_tool"
	nodeExternalCallLLM   = "external_call_llm"
	nodeExternalInterrupt = "external_interrupt"
	nodeAnswerLLM         = "answer_llm"
)

const internalCallInstruction = `You are a verification assistant for a mixed internal/external graph demo.

First gather server-side context. Before answering, call internal_lookup and internal_profile exactly once based on the user request, and emit both tool calls in the same assistant message.`

const externalCallInstruction = `You are a verification assistant for a mixed internal/external graph demo.

The server-side internal tool results are now available. Before answering, call external_search and external_approval exactly once based on the user request and internal results, and emit both tool calls in the same assistant message.`

const answerInstruction = `You are a helpful assistant. Answer using all available internal and external tool results.`

func buildGraph(modelInstance model.Model, generationConfig model.GenerationConfig) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func routeAfterToolCall(toolNode, fallbackNode string) func(context.Context, graph.State) (string, error) {
	_ = "STUB: not implemented"
	return nil
}
