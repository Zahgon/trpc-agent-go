//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

func newRemoteEvalAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newJudgeAgent(modelName string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

type calculatorInput struct {
	Operation string
	A         float64
	B         float64
}

type calculatorToolResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

func calculate(_ context.Context, args calculatorInput) (calculatorToolResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorToolResult), nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
