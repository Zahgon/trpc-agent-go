//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main implements a tiny STDIO MCP server used by the Claude Code CLI example.
package main

import (
	"context"
	"log"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

const calculatorToolName = "calculator"

// calculatorArgs is the argument schema for the calculator MCP tool.
type calculatorArgs struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

// calculatorResult is the JSON payload returned by the calculator MCP tool.
type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

func main() {
	server := mcp.NewStdioServer("cli-example-calculator", "1.0.0")
	server.RegisterTool(
		mcp.NewTool(
			calculatorToolName,
			mcp.WithDescription("Perform arithmetic operations including add, subtract, multiply, and divide."),
			mcp.WithString("operation", mcp.Required(), mcp.Description("Operation: add, subtract, multiply, or divide.")),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First operand.")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second operand.")),
		),
		handleCalculator,
	)
	if err := server.Start(); err != nil {
		log.Fatalf("mcp server failed: %v", err)
	}
}

// handleCalculator evaluates the calculator tool request and returns a JSON string result.
func handleCalculator(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseCalculatorArgs parses the calculator tool call arguments into a strongly typed struct.
func parseCalculatorArgs(req *mcp.CallToolRequest) (calculatorArgs, error) {
	_ = "STUB: not implemented"
	return *new(calculatorArgs), nil
}

// compute applies the requested arithmetic operation and returns a structured JSON result.
func compute(args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}
