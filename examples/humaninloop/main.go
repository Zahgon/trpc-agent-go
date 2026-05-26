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
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
)

// lastPendingTicketID remembers the most recent pending approval ticket id.
var lastPendingTicketID string

func main() {
	flag.Parse()

	fmt.Printf("🚀 Human-in-the-Loop (HIL) Reimbursement Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Println(strings.Repeat("=", 50))

	r, err := setupRunner(*modelName, *streaming)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer r.Close()

	printHelp()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("👤 You: ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if !dispatchInput(context.Background(), r, input) {
			return
		}
	}
}

// setupRunner creates the runner with in-memory session.
func setupRunner(model string, stream bool) (runner.Runner, error) {
	_ = "STUB: not implemented"
	return *new(runner.Runner), nil
}

// printHelp prints quick usage hints for the demo.
func printHelp() { _ = "STUB: not implemented"; return }

// dispatchInput routes a single user input. Returns false to exit.
func dispatchInput(ctx context.Context, r runner.Runner, input string) bool {
	_ = "STUB: not implemented"
	return false
}

const userID, sessionID = "user-123", "session-123"

func processStreamingResponse(ctx context.Context, r runner.Runner, message model.Message) (*model.ToolCall, *askForApprovalOutput) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle errors.

// Detect and display tool calls.

// If a long-running call is detected and we haven't sent approval yet, simulate it.

// Detect tool responses.

// If we got a pending approval, simulate external approval automatically (if not already sent).

// Process streaming content.

// Check if this is the final e.
// Don't break on tool response events (Done=true but not final assistant response).

// handleToolCalls processes tool call events and returns updated state
func handleToolCalls(e *event.Event, longRunningFunctionCall *model.ToolCall, toolCallsDetected bool, assistantStarted bool) (*model.ToolCall, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

// Create a local copy to avoid implicit memory aliasing in Go <= 1.21.

// Remember the tool call ID as a fallback ticket for shorthand approval.

// handleToolResponses processes tool response events
func handleToolResponses(e *event.Event, longRunningFunctionCall *model.ToolCall) *askForApprovalOutput {
	_ = "STUB: not implemented"
	return nil
}

// Remember the last pending ticket id to enable shorthand approvals.

// processStreamingContent processes streaming content events
func processStreamingContent(e *event.Event, toolCallsDetected bool, assistantStarted bool, fullContent string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Handle streaming delta content.

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
