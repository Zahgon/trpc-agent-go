//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import (
	"context"

	agenttool "trpc.group/trpc-go/trpc-agent-go/tool"
)

const calculatorName = "calculator"

func newCalculatorTool() agenttool.Tool { _ = "STUB: not implemented"; return *new(agenttool.Tool) }

func calculator(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

type calculatorArgs struct {
	Operation string `json:"operation" description:"add, subtract, multiply, or divide"`
	A         int    `json:"a" description:"The first integer."`
	B         int    `json:"b" description:"The second integer."`
}

type calculatorResult struct {
	Result int `json:"result"`
}
