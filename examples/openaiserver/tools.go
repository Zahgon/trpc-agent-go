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
)

// Constants for supported calculator operations.
const (
	opAdd      = "add"
	opSubtract = "subtract"
	opMultiply = "multiply"
	opDivide   = "divide"
)

// calculatorArgs holds the input for the calculator tool.
type calculatorArgs struct {
	Operation string  `json:"operation" jsonschema:"description=The operation to perform,enum=add,enum=subtract,enum=multiply,enum=divide,required"`
	A         float64 `json:"a" jsonschema:"description=First number operand,required"`
	B         float64 `json:"b" jsonschema:"description=Second number operand,required"`
}

// calculatorResult holds the output for the calculator tool.
type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

// timeArgs holds the input for the time tool.
type timeArgs struct {
	Timezone string `json:"timezone" jsonschema:"description=Timezone or leave empty for local,required"`
}

// timeResult holds the output for the time tool.
type timeResult struct {
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
	Date     string `json:"date"`
	Weekday  string `json:"weekday"`
}

// Calculator tool implementation.
// calculate performs the requested mathematical operation.
// It supports add, subtract, multiply, and divide operations.
func calculate(ctx context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"

	// Select operation based on input.
	return *new(calculatorResult), nil
}

// Time tool implementation.
// getCurrentTime returns the current time for the specified timezone.
// If the timezone is invalid or empty, it defaults to local time.
func getCurrentTime(ctx context.Context, args timeArgs) (timeResult, error) {
	_ = "STUB: not implemented"
	return *new(timeResult), nil
}

// Attempt to load the specified timezone.

// This example demonstrates how to integrate tRPC agent orchestration
// with LLM-based tools, providing a simple HTTP server compatible with
// OpenAI API for manual testing. It is intended as a reference for
// developers looking to build custom LLM agents with tool support in Go.
//
// The calculator tool supports basic arithmetic operations, while the
// time tool provides current time information for a given timezone.
//
// The code is structured for clarity and ease of extension.
