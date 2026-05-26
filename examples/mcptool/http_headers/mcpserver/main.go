//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides a simple SSE MCP server that logs received HTTP headers.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

func main() {
	fmt.Println("🚀 Starting SSE MCP Server with HTTP Header Logging...")
	fmt.Println("Server will log all received HTTP headers")
	fmt.Println("Listening on http://localhost:3000/mcp")
	fmt.Println(strings.Repeat("=", 50))

	// Create MCP server
	server := mcp.NewServer(
		"header-demo-server",
		"1.0.0",
		mcp.WithServerAddress(":3000"),
		mcp.WithServerPath("/mcp"),
	)

	// Register tools
	registerWeatherTool(server)
	registerEchoTool(server)

	// Wrap the server's HTTP handler to log headers
	originalHandler := server.HTTPHandler()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logHeaders(r)
		originalHandler.ServeHTTP(w, r)
	})

	fmt.Println("✅ Server ready")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func logHeaders(r *http.Request) { _ = "STUB: not implemented"; return }

// Read body content

// Restore the body for subsequent handlers

// Log custom headers

// Log other interesting headers

func registerWeatherTool(server *mcp.Server) { _ = "STUB: not implemented"; return }

func registerEchoTool(server *mcp.Server) { _ = "STUB: not implemented"; return }
