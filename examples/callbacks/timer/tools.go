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

// Tool implementations.

// calculator performs basic operations.
func (e *toolTimerExample) calculator(ctx context.Context, args *calculatorArgs) (*calculatorResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Data structures.

// calculatorArgs represents arguments for the calculator tool.
type calculatorArgs struct {
	Operation string  `json:"operation" description:"The operation to perform"`
	A         float64 `json:"a" description:"First number"`
	B         float64 `json:"b" description:"Second number"`
}

// calculatorResult represents the result of a calculation.
type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
