//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.

// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a customer support workflow using GraphAgent with A2A sub-agents.
// This example shows how to build and execute graphs that coordinate with remote agents
// through the A2A protocol for specialized tasks.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	// Default model name for deepseek-v4-flash.
	defaultModelName = "deepseek-v4-flash"
	// Default A2A server host.
	defaultA2AHost = "0.0.0.0:8888"
)

// State keys for customer support workflow.
const (
	stateKeyCustomerQuery    = "customer_query"
	stateKeyQueryType        = "query_type"
	stateKeyPriority         = "priority"
	stateKeyTechnicalDetails = "technical_details"
	stateKeyResponse         = "response"
	stateKeyFinalAnswer      = "final_answer"
)

// Query types.
const (
	queryTypeTechnical = "technical"
	queryTypeBilling   = "billing"
	queryTypeGeneral   = "general"
)

// Priority levels.
const (
	priorityHigh   = "high"
	priorityMedium = "medium"
	priorityLow    = "low"
)

// Node names for the workflow graph.
const (
	nodeAnalyzeQuery     = "analyze_query"
	nodeTechnicalSupport = "technical_support"
	nodeBillingSupport   = "billing_support"
	nodeGeneralSupport   = "general_support"
	nodeFormatResponse   = "format_response"
)

// Agent names.
const (
	agentTechnicalSupport = "technical-support-agent"
	a2aStateKeyMetadata   = "meta"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
	a2aHost = flag.String("a2a-host", defaultA2AHost,
		"A2A server host to connect to")
	interactive = flag.Bool("interactive", false,
		"Run in interactive mode")
)

func main() {
	// Parse command line flags.
	flag.Parse()
	fmt.Printf("🚀 Customer Support Workflow with A2A Sub-Agent Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("A2A Host: %s\n", *a2aHost)
	fmt.Println(strings.Repeat("=", 60))

	// Create and run the workflow.
	workflow := &customerSupportWorkflow{
		modelName: *modelName,
		a2aHost:   *a2aHost,
	}
	if err := workflow.run(); err != nil {
		log.Fatalf("Workflow failed: %v", err)
	}
}

// customerSupportWorkflow manages the customer support workflow with A2A sub-agents.
type customerSupportWorkflow struct {
	modelName string
	a2aHost   string
	runner    runner.Runner
	userID    string
	sessionID string
}

// run starts the customer support workflow.
func (w *customerSupportWorkflow) run() error { _ = "STUB: not implemented"; return nil }

// Setup logging.

// Start A2A server in background.

// Wait for server to start.

// Setup the workflow.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// setupLogging configures logging for the workflow.
func (w *customerSupportWorkflow) setupLogging() { _ = "STUB: not implemented"; return }

type hookProcessor struct {
	next taskmanager.MessageProcessor
}

func (h *hookProcessor) ProcessMessage(
	ctx context.Context,
	message protocol.Message,
	options taskmanager.ProcessOptions,
	handler taskmanager.TaskHandler,
) (*taskmanager.MessageProcessingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// startA2AServer starts the A2A server with a specialized remote agent.
func (w *customerSupportWorkflow) startA2AServer() { _ = "STUB: not implemented"; return }

// Enable A2A streaming to demonstrate live token streaming end-to-end.

// Start server in a separate goroutine to avoid blocking

// Redirect A2A server logs to avoid mixing with our output

// Give the server a moment to start

// buildRemoteAgent creates a specialized remote agent for technical support.
func (w *customerSupportWorkflow) buildRemoteAgent() agent.Agent {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return *new(agent.Agent)
}

// Create specialized tools for technical support.

// Create LLM agent with tools for technical support.

// Enable streaming so A2A can forward live deltas

// setup creates the graph agent and runner.
func (w *customerSupportWorkflow) setup() error {
	_ = "STUB: not implemented"
	// Create the customer support workflow graph.
	return nil
}

// Create A2A agent for remote technical support.

// Important: use the same name as the agent node ID so the graph can
// resolve and invoke this sub-agent.

// Create GraphAgent from the compiled graph with A2A sub-agent.

// Create session service.

// Create runner.

// Generate session ID.

// createCustomerSupportGraph creates the customer support workflow graph.
func (w *customerSupportWorkflow) createCustomerSupportGraph() (*graph.Graph, error) {
	_ = "STUB: not implemented"
	// Define state schema for customer support workflow.
	return nil, nil
}

// Create the workflow graph.

// Initial analysis node.

// Technical support using built-in agent node that routes to sub-agent.

// Billing support node.

// General support node.

// Final response formatting.

// Route based on query type.

// Set up the workflow edges.

// runDefaultExamples runs predefined customer support examples.
func (w *customerSupportWorkflow) runDefaultExamples(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Response content will stream/live-print below; avoid preface duplication.

// Add a small delay between queries to prevent race conditions

// startInteractiveMode starts interactive mode for user input.
func (w *customerSupportWorkflow) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Response content will stream/live-print below; avoid preface duplication.

// processQuery processes a single customer query.
func (w *customerSupportWorkflow) processQuery(ctx context.Context, query string) error {
	_ = "STUB: not implemented"
	// Create user message.
	return nil
}

// Run the workflow.

// Process streaming response.

// processStreamingResponse handles the streaming workflow response.
func (w *customerSupportWorkflow) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// streaming seen from A2A agent

// Track node execution events.

// Handle errors after metadata so node-level error
// metadata remains visible.

// Process streaming content from LLM nodes and A2A agents.

// Handle streaming delta content.

// Add newline when streaming is complete (when choice is done).

// Add newline after streaming completes
// Reset for next response

// Check for A2A agent responses in state updates.
// Only label as A2A when the event author is the A2A agent to avoid
// mislabeling final formatted responses. If we've already streamed from A2A,
// skip this to avoid duplication.

// Try to decode as structured model.Response first

// Track workflow stages (concise to avoid duplicate long content).

// Handle completion and final response.

// If A2A streaming happened, avoid repeating the same content; show a concise completion mark.

// Print final response once: prefer formatted final_answer, otherwise fallback to completion response.

// showHelp displays help information.
func (w *customerSupportWorkflow) showHelp() { _ = "STUB: not implemented"; return }

// Node functions for the workflow.

// analyzeCustomerQuery analyzes the customer query to determine type and priority.
func (w *customerSupportWorkflow) analyzeCustomerQuery(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Simple keyword-based analysis for demonstration.

// Determine query type.

// Determine priority.

// routeByQueryType routes the workflow based on query type.
func (w *customerSupportWorkflow) routeByQueryType(ctx context.Context, state graph.State) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// handleBillingQuery handles billing-related queries.
func (w *customerSupportWorkflow) handleBillingQuery(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Return state update with response

// handleGeneralQuery handles general customer service queries.
func (w *customerSupportWorkflow) handleGeneralQuery(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Return state update with response

// formatFinalResponse formats the final response for the customer.
func (w *customerSupportWorkflow) formatFinalResponse(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Get the response from the previous node
	return *new(any), nil
}

// If no response from previous node, try to get from agent/LLM response.

// Return state update with final response. Also set graph.StateKeyLastResponse
// so completion events carry the final content for consumers that only read
// the standard last_response field.

// Helper functions.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 {
	_ = "STUB: not implemented"

	// stateBuilder helps build state updates in a type-safe way.
	return nil
}

type stateBuilder struct {
	state graph.State
}

// newStateBuilder creates a new state builder.
func newStateBuilder() *stateBuilder { _ = "STUB: not implemented"; return nil }

// SetCustomerQuery sets the customer query.
func (sb *stateBuilder) SetCustomerQuery(query string) *stateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SetQueryType sets the query type.
func (sb *stateBuilder) SetQueryType(queryType string) *stateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SetPriority sets the priority.
func (sb *stateBuilder) SetPriority(priority string) *stateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SetResponse sets the response.
func (sb *stateBuilder) SetResponse(response string) *stateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SetFinalAnswer sets the final answer.
func (sb *stateBuilder) SetFinalAnswer(answer string) *stateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns the built state.
func (sb *stateBuilder) Build() graph.State {
	_ = "STUB: not implemented"

	// truncateString truncates a string to the specified length and adds ellipsis if needed.
	return *new(graph.State)
}

func truncateString(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }
