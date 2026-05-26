//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates per-run tool filtering functionality.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func main() {
	// Parse command line arguments
	modelName := flag.String("model", "deepseek-v4-flash", "Model name to use")
	filterMode := flag.String("filter", "", "Filter mode: exclude-demo, include-demo, per-agent, or empty for no filter")
	flag.Parse()

	fmt.Printf("🚀 Multi-Agent Tool Filtering Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	if *filterMode != "" {
		fmt.Printf("Filter Mode: %s\n", *filterMode)
	}
	fmt.Printf("Enter 'exit' to end the conversation\n")
	fmt.Printf("Available agents: math-agent (calculator, text_tool), time-agent (time_tool, text_tool)\n")
	fmt.Println(strings.Repeat("=", 60))

	// Create and run chat system
	chat := &toolFilterDemo{
		modelName:  *modelName,
		filterMode: *filterMode,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Demo failed: %v", err)
	}
}

// toolFilterDemo manages the multi-agent tool filtering demonstration
type toolFilterDemo struct {
	modelName  string
	runner     runner.Runner
	userID     string
	sessionID  string
	filterMode string
}

// run starts the demo
func (c *toolFilterDemo) run() error { _ = "STUB: not implemented"; return nil }

// Setup runner

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat

// setup creates a runner with multiple sub-agents
func (c *toolFilterDemo) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model with request callback to show tools
	return nil
}

// Print tools that will be sent to the model

// Create math-agent with calculator and text tools

// Create time-agent with time and text tools

// Create coordinator agent with sub-agents

// Create runner

// Set identifiers

// startChat runs the interactive conversation loop
func (c *toolFilterDemo) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle exit command

// Process user message

// processMessage processes a single message exchange
func (c *toolFilterDemo) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Build run options with tool filtering

// Apply tool filtering based on filter mode

// Demo: Exclude specific tools using NewExcludeToolNamesFilter

// Demo: Include only specific tools using NewIncludeToolNamesFilter

// Demo: Custom per-agent filtering using FilterFunc
// This demonstrates how users can implement agent-specific filtering

// Run agent through runner

// Process streaming response

// createPerAgentFilter creates a custom filter that applies different rules per agent.
// This demonstrates how users can implement agent-specific filtering logic.
//
// Key insight: The FilterFunc receives context.Context, which contains the invocation
// information. You can use agent.InvocationFromContext(ctx) to get the current agent name
// and apply agent-specific filtering rules.
func (c *toolFilterDemo) createPerAgentFilter() tool.FilterFunc {
	_ = "STUB: not implemented"
	// Define per-agent allowed tools
	return *new(tool.FilterFunc)
}

// text_tool is NOT in this list, so it will be filtered out

// text_tool is NOT in this list, so it will be filtered out

// Get the current agent name from invocation context
// The context contains invocation information during tool filtering

// If no invocation context, allow all tools (fallback)

// Check if this tool is allowed for the current agent

// If agent not in the map, allow all tools (fallback)

// Return true only if the tool is in the agent's allowed list

// processStreamingResponse processes streaming response
func (c *toolFilterDemo) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors

// Detect and display tool calls

// Process streaming content

// Check if this is the final event

// Calculator tool
type calculatorRequest struct {
	Expression string `json:"expression" jsonschema:"description=Mathematical expression to calculate,required"`
}

type calculatorResponse struct {
	Result  float64 `json:"result"`
	Message string  `json:"message"`
}

func createCalculatorTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}

func calculateExpression(_ context.Context, req calculatorRequest) (calculatorResponse, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResponse), nil
}

func evaluateBasicExpression(expr string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Simple evaluation for demo (supports +, -, *, /)

// Time tool
type timeRequest struct {
	Operation string `json:"operation" jsonschema:"description=Operation: current, date, or timestamp,required"`
}

type timeResponse struct {
	Result string `json:"result"`
}

func createTimeTool() tool.CallableTool { _ = "STUB: not implemented"; return *new(tool.CallableTool) }

func getTimeInfo(_ context.Context, req timeRequest) (timeResponse, error) {
	_ = "STUB: not implemented"
	return *new(timeResponse), nil
}

// Text tool
type textRequest struct {
	Text      string `json:"text" jsonschema:"description=Text to process,required"`
	Operation string `json:"operation" jsonschema:"description=Operation: uppercase or lowercase,required"`
}

type textResponse struct {
	Result string `json:"result"`
}

func createTextTool() tool.CallableTool { _ = "STUB: not implemented"; return *new(tool.CallableTool) }

func processText(_ context.Context, req textRequest) (textResponse, error) {
	_ = "STUB: not implemented"
	return *new(textResponse), nil
}

// Helper functions
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }

var _ = math.Sqrt // Avoid unused import error
