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
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// createMathAgent creates a specialized math agent.
func (c *transferChat) createMathAgent(modelInstance model.Model) agent.Agent {
	_ = "STUB: not implemented"
	// Math calculation tool.
	return *new(agent.Agent)
}

// Lower temperature for more precise calculations.

// calculate performs mathematical operations.
func (c *transferChat) calculate(_ context.Context, args calcArgs) (calcResult, error) {
	_ = "STUB: not implemented"
	return *new(calcResult), nil
}

// Data structures for math tool.
type calcArgs struct {
	Operation string  `json:"operation" jsonschema:"description=The operation to perform,enum=add,enum=subtract,enum=multiply,enum=divide,enum=power,required"`
	A         float64 `json:"a" jsonschema:"description=First number operand,required"`
	B         float64 `json:"b" jsonschema:"description=Second number operand,required"`
}

type calcResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
	Error     string  `json:"error,omitempty"`
}
