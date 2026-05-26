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

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
)

const (
	agentName    = "toolcallid-assistant"
	toolNameCalc = "calculator"
	opAdd        = "add"
	opSubtract   = "subtract"
	opMultiply   = "multiply"
	opDivide     = "divide"
)

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

func newAgent(modelName string, variant string, streaming bool) *llmagent.LLMAgent {
	_ = "STUB: not implemented"
	return nil
}

func calculate(ctx context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
