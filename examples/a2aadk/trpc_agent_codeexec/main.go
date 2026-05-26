// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.

// Package main demonstrates handling code execution events from ADK A2A server
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/a2aagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

var (
	agentURL = flag.String("url", "http://localhost:8082", "ADK A2A agent server URL with code execution")
)

func main() {
	flag.Parse()

	fmt.Println("========================================")
	fmt.Println("trpc-agent-go A2A Code Execution Demo")
	fmt.Println("========================================")
	fmt.Printf("Connecting to ADK A2A server: %s\n", *agentURL)
	fmt.Println("========================================")
	fmt.Println()

	// Create A2A agent client
	a2aAgent, err := a2aagent.New(
		a2aagent.WithAgentCardURL(*agentURL),
	)
	if err != nil {
		log.Fatalf("Failed to create A2A agent: %v", err)
	}

	// Display agent card info
	card := a2aAgent.GetAgentCard()
	fmt.Printf("Connected to agent:\n")
	fmt.Printf("  Name: %s\n", card.Name)
	fmt.Printf("  Description: %s\n", card.Description)
	fmt.Printf("  URL: %s\n", card.URL)
	fmt.Println()

	// Create session service and runner
	sessionService := inmemory.NewSessionService()
	agentRunner := runner.NewRunner("test", a2aAgent, runner.WithSessionService(sessionService))
	defer agentRunner.Close()

	ctx := context.Background()
	userID := "test_user"
	// Use unique session ID to avoid polluted session history on ADK server
	sessionID := fmt.Sprintf("session_%d", time.Now().UnixNano())

	fmt.Printf("Client Session Info:\n")
	fmt.Printf("  User ID: %s\n", userID)
	fmt.Printf("  Session ID: %s\n", sessionID)
	fmt.Println()

	// Test 1: Simple code execution
	fmt.Println("Test 1: Simple Python Code Execution")
	fmt.Println("=====================================")
	testQuery(ctx, agentRunner, userID, sessionID,
		"Calculate the sum of numbers from 1 to 10 using Python code")
	fmt.Println()

	// Test 2: Data analysis
	fmt.Println("Test 2: Data Analysis with Code")
	fmt.Println("================================")
	testQuery(ctx, agentRunner, userID, sessionID,
		"Analyze this data: [5, 12, 8, 15, 7, 9, 11]. Calculate mean, median, and standard deviation using Python.")
	fmt.Println()

	// Test 3: Plot generation (if matplotlib is available)
	fmt.Println("Test 3: Code with Multiple Steps")
	fmt.Println("=================================")
	testQuery(ctx, agentRunner, userID, sessionID,
		"Create a list of fibonacci numbers up to the 10th term using Python")
	fmt.Println()
}

func testQuery(ctx context.Context, agentRunner runner.Runner, userID, sessionID, query string) {
	_ = "STUB: not implemented"
	return
}

// processCodeExecutionResponse processes events and displays code execution information
func processCodeExecutionResponse(events <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle code execution events

// Handle code execution result events

// Capture assistant content from intermediate (non-final) events only

// Print content when we receive the final response event

// handleCodeExecution processes code execution events
func handleCodeExecution(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// Use Delta for streaming response

// handleCodeExecutionResult processes code execution result events.
// Requires: ObjectType == codeexecution && Tag == code_execution_result
func handleCodeExecutionResult(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// Check ObjectType first, then Tag for result

// Use Delta for streaming response

// captureFinalContent extracts assistant text from message or delta
func captureFinalContent(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

// Only capture assistant messages, skip tool responses and code execution

// Skip code execution events (both code and result have the same ObjectType)
