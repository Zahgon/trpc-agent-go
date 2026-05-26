//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides an SSE (Server-Sent Events) MCP server example.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

func main() {
	port := flag.Int("port", 8080, "Listen port")
	flag.Parse()

	// Create MCP SSE server.
	server := mcp.NewSSEServer("SSE Example Server", "1.0.0")

	// Register recipe tool.
	recipeTool := mcp.NewTool("sse_recipe",
		mcp.WithDescription("Chinese recipe query tool"),
		mcp.WithString("dish", mcp.Description("Dish name")),
	)
	server.RegisterTool(recipeTool, handleRecipe)

	// Register health tip tool.
	healthTipTool := mcp.NewTool("sse_health_tip",
		mcp.WithDescription("Health tip tool"),
		mcp.WithString("category", mcp.Description("Category")),
	)
	server.RegisterTool(healthTipTool, handleHealthTip)

	// Handle signals.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	// Start server.
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Start SSE server, port: %d\n", *port)
	fmt.Printf("Available tools: sse_recipe, sse_health_tip\n") // Available tools: sse_recipe, sse_health_tip.

	go server.Start(addr)
	<-ctx.Done()
	server.Shutdown(context.Background())
}

// Handle recipe tool.
func handleRecipe(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	// Extract dish parameter.
	return nil, nil
}

// Return a simplified but real recipe.

// Handle health tip tool.
func handleHealthTip(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	// Extract category parameter.
	return nil, nil
}

// Return a simplified tip.
