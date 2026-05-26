//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const calculatorDescription = "Run a basic arithmetic calculation."

type calculatorInput struct {
	A         float64 `json:"a" jsonschema:"description=Left operand,required"`
	B         float64 `json:"b" jsonschema:"description=Right operand,required"`
	Operation string  `json:"operation" jsonschema:"description=Operation: add, subtract, multiply, or divide.,required"`
}

type calculatorOutput struct {
	Expression string  `json:"expression" jsonschema:"description=Expression that was calculated"`
	Result     float64 `json:"result" jsonschema:"description=Calculation result"`
}

func calculatorTools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func calculate(ctx context.Context, input calculatorInput) (calculatorOutput, error) {
	_ = "STUB: not implemented"
	return *new(calculatorOutput), nil
}

func calculateResult(a, b float64, operation string) (float64, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}
