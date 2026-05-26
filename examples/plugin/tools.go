//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolNameCalculator = "calculator"
	toolDescCalculator = "Perform basic math operations."

	opAdd      = "add"
	opSubtract = "subtract"
	opMultiply = "multiply"
	opDivide   = "divide"
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

func newCalculatorTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func calculatorInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func calculate(
	_ context.Context,
	args calculatorArgs,
) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}
