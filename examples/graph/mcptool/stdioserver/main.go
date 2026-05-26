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

// WeatherRequest represents the input structure for weather tool.
type WeatherRequest struct {
	Location string `json:"location" jsonschema:"required,description=City name or location"`
	Units    string `json:"units,omitempty" jsonschema:"description=Temperature units,enum=celsius,enum=fahrenheit,default=celsius"`
}

// WeatherResponse represents the output structure for weather tool.
type WeatherResponse struct {
	Location    string `json:"location" jsonschema:"required,description=Requested location"`
	Temperature int    `json:"temperature" jsonschema:"required,description=Current temperature in degrees"`
	Condition   string `json:"condition" jsonschema:"required,description=Weather condition"`
	Humidity    int    `json:"humidity" jsonschema:"required,description=Humidity percentage"`
	WindSpeed   int    `json:"windSpeed" jsonschema:"required,description=Wind speed in km/h"`
	Units       string `json:"units" jsonschema:"required,description=Temperature units"`
}

func main() {
	server := mcp.NewStdioServer("graph-use-mcp-stdio-server", "1.0.0",
		mcp.WithStdioServerLogger(mcp.GetDefaultLogger()),
	)

	// Register echo tool.
	echoTool := mcp.NewTool("echo",
		mcp.WithDescription("Simple echo tool that returns the input message with an optional prefix"),
		mcp.WithString("message", mcp.Required(), mcp.Description("The message to echo")),
		mcp.WithString("prefix", mcp.Description("Optional prefix, default is 'Echo: '")),
	)
	server.RegisterTool(echoTool, handleEcho)

	// Register add tool.
	addTool := mcp.NewTool("add",
		mcp.WithDescription("Simple addition tool that adds two numbers"),
		mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
		mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
	)
	server.RegisterTool(addTool, handleAdd)

	// Register a weather tool with struct‑first API and OutputSchema.
	weatherTool := mcp.NewTool(
		"get_weather",
		mcp.WithDescription("Get current weather for a location with structured output"),
		mcp.WithInputStruct[WeatherRequest](),
		mcp.WithOutputStruct[WeatherResponse](),
	)

	server.RegisterTool(weatherTool, mcp.NewTypedToolHandler(
		func(ctx context.Context, req *mcp.CallToolRequest, input WeatherRequest) (WeatherResponse, error) {
			units := input.Units
			if units == "" {
				units = "celsius"
			}

			resp := WeatherResponse{
				Location:    input.Location,
				Temperature: 22,
				Condition:   "Sunny",
				Humidity:    45,
				WindSpeed:   10,
				Units:       units,
			}

			return resp, nil
		},
	))

	log.Printf("Starting graph-use-mcp STDIO MCP Server...")
	log.Printf("Available tools: echo, add, get_weather")

	// Start server.
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
