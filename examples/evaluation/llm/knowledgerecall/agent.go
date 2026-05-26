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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func newQAAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newSearchTool() tool.Tool {
	_ = "STUB: not implemented"
	return *

	// 1. Create embedder.
	new(tool.Tool)
}

// 2. Create vector store.

// 3. Create knowledge sources.

// 4. Create Knowledge.

// 5. Load documents.

// 6. Create search tool.

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

func calculate(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

func intPtr(v int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
