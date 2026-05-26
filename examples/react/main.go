//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates React planning with LLM agents using structured
// planning instructions, tool calling, and response processing.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func main() {
	// Parse command line flags.
	modelName := flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	flag.Parse()

	fmt.Printf("🧠 React Planning Agent Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Printf("Available tools: search, calculator, weather\n")
	fmt.Printf("The agent will use React planning to structure its responses\n")
	fmt.Println(strings.Repeat("=", 60))

	// Create and run the chat.
	chat := &reactPlanningChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// reactPlanningChat manages the conversation with React planning.
type reactPlanningChat struct {
	modelName string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the interactive chat session.
func (c *reactPlanningChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

// setup creates the runner with LLM agent and React planner.
func (c *reactPlanningChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return nil
}

// Create tools for demonstration.

// Create React planner.

// Create LLM agent with React planner and tools.

// Enable streaming

// Create runner.

// Setup identifiers.

// startChat runs the interactive conversation loop.
func (c *reactPlanningChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle exit command.

// Process the user message.

// Add spacing between turns

// processMessage handles a single message exchange.
func (c *reactPlanningChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run the agent through the runner.

// Process streaming response with React planning awareness.

// processStreamingResponse handles the streaming response with React planning visualization.
func (c *reactPlanningChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls.

// Process text content with React planning awareness.

// Handle tool responses.

// End the response

// Tool implementations for demonstration.

// search simulates a search tool.
func (c *reactPlanningChat) search(_ context.Context, args searchArgs) (searchResult, error) {
	_ = "STUB: not implemented"
	return *new(searchResult), nil
}

// calculate performs mathematical calculations.
func (c *reactPlanningChat) calculate(_ context.Context, args calcArgs) (calcResult, error) {
	_ = "STUB: not implemented"
	return *new(calcResult), nil
}

// getWeather simulates weather information retrieval.
func (c *reactPlanningChat) getWeather(_ context.Context, args weatherArgs) (weatherResult, error) {
	_ = "STUB: not implemented"
	return *new(weatherResult), nil
}

// Tool argument and result types.

type searchArgs struct {
	Query string `json:"query" description:"The search query"`
}

type searchResult struct {
	Query   string   `json:"query"`
	Results []string `json:"results"`
	Count   int      `json:"count"`
}

type calcArgs struct {
	Operation string  `json:"operation" description:"The operation to perform,enum=add,enum=subtract,enum=multiply,enum=divide,enum=power"`
	A         float64 `json:"a" description:"First number"`
	B         float64 `json:"b" description:"Second number"`
}

type calcResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}

type weatherArgs struct {
	Location string `json:"location" description:"The location to get weather for"`
}

type weatherResult struct {
	Location       string  `json:"location"`
	Temperature    float64 `json:"temperature"`
	Condition      string  `json:"condition"`
	Humidity       int     `json:"humidity"`
	Recommendation string  `json:"recommendation"`
}

// Helper functions.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
