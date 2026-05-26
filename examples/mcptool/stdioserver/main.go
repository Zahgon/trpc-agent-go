//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides a simple STDIO MCP server example.
package main

import (
	"context"
	"log"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

func main() {
	server := mcp.NewStdioServer("simple-stdio-server", "1.0.0",
		mcp.WithStdioServerLogger(mcp.GetDefaultLogger()),
	)

	// Register echo tool
	echoTool := mcp.NewTool("echo",
		mcp.WithDescription("Simple echo tool that returns the input message with an optional prefix"),
		mcp.WithString("message", mcp.Required(), mcp.Description("The message to echo")),
		mcp.WithString("prefix", mcp.Description("Optional prefix, default is 'Echo: '")),
	)
	server.RegisterTool(echoTool, handleEcho)

	// Register add tool
	addTool := mcp.NewTool("add",
		mcp.WithDescription("Simple addition tool that adds two numbers"),
		mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
		mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
	)
	server.RegisterTool(addTool, handleAdd)

	log.Printf("Starting Simple STDIO MCP Server...")
	log.Printf("Available tools: echo, add")
	log.Printf("Using simplified implementation")

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// handleEcho handles the echo tool.
func handleEcho(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	// Parse message parameter.
	return nil, nil
}

// Parse prefix parameter.

// handleAdd handles the add tool.
func handleAdd(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	// Parse a parameter.
	return nil, nil
}

// Parse b parameter.
