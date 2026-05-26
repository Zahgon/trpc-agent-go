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
	"context"
	"flag"
	"fmt"
	"time"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/a2aagent"
	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
	"trpc.group/trpc-go/trpc-agent-go/tool/function"
)

var (
	modelName  = flag.String("model", getEnvOrDefault("MODEL_NAME", "deepseek-v4-flash"), "Model to use")
	host       = flag.String("host", "127.0.0.1:8888", "Host to use")
	streaming  = flag.Bool("streaming", true, "Streaming to use")
	serverMode = flag.String("server-mode", "agent", "A2A server build mode: agent or runner-card")
	remoteOnly = flag.Bool("remote-only", false, "Only output remote agent responses")
	debugMode  = flag.Bool("debug", true, "Enable debug mode to print session events after each turn")
)

// ANSI color codes for terminal output
const (
	colorReset = "\033[0m"
	colorCyan  = "\033[36m" // Cyan for reasoning/thinking content
)

const (
	optionalStateKey = "meta"
	appName          = "a2aagent-demo"
)

func main() {
	flag.Parse()

	// runRemoteAgent will start a a2a server that build with a remote agent
	if err := runA2AServerByAgent("agent_remote_joker", "I am a remote agent, I can tell a joke", *host); err != nil {
		log.Fatalf("Failed to start a2a server: %v", err)
	}

	httpURL := fmt.Sprintf("http://%s", *host)
	a2aAgent := buildA2AAgent(httpURL)

	// Build a different local agent
	localAgent := buildAgent("agent_local_joker", "I am a local agent, I can tell a joke",
		llmagent.WithTools([]tool.Tool{
			function.NewFunctionTool(
				getCurrentTime,
				function.WithName("getCurrentTime"),
				function.WithDescription("This is tool that can get current time")),
		}))
	fmt.Printf("Debug Mode: %t\n", *debugMode)
	startChat(localAgent, a2aAgent)
}

func startChat(localAgent agent.Agent, a2aAgent *a2aagent.A2AAgent) {
	_ = "STUB: not implemented"
	return
}

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Use different userIDs and sessionIDs for remote and local agents

// Add spacing between turns

func printDebugSessions(
	ctx context.Context,
	remoteSessionService session.Service,
	remoteUserID string,
	remoteSessionID string,
	localSessionService session.Service,
	localUserID string,
	localSessionID string,
) {
	_ = "STUB: not implemented"
	return
}

func printSessionEvents(
	ctx context.Context,
	svc session.Service,
	appName string,
	userID string,
	sessionID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func buildDebugEventDetail(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func processMessage(
	remoteRunner runner.Runner,
	localRunner runner.Runner,
	remoteUserID string,
	remoteSessionID *string,
	localUserID string,
	localSessionID *string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Example: Pass custom HTTP headers to A2A agent using WithA2ARequestOptions
// This allows you to add authentication tokens, tracing IDs, or other custom headers

// Only run local agent if remote-only flag is not set

func startNewSession(prefix string) string { _ = "STUB: not implemented"; return "" }

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

func runA2AServerByAgent(agentName, desc, host string) error { _ = "STUB: not implemented"; return nil }

// Create in-memory memory service for demonstration.

// Create in-memory session service for the runner.

// Example: Use WithProcessMessageHook to inspect/modify incoming A2A messages.
// This can read custom metadata injected by the client's BuildMessageHook.

// In runner-only mode, the public agent identity must be supplied
// explicitly via WithAgentCard.

func ensureHostAvailable(host string) error { _ = "STUB: not implemented"; return nil }

func waitForAgentCardReady(host string, serverErrCh <-chan error, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func buildAgent(agentName, desc string, extraOptions ...llmagent.Option) agent.Agent {
	_ = "STUB: not implemented"
	// Create OpenAI model.
	return *new(agent.Agent)
}

// Create LLM agent.

func buildA2AAgent(httpURL string) *a2aagent.A2AAgent { _ = "STUB: not implemented"; return nil }

// optional: specify the state key that transferred to the remote agent by metadata

// Example: Use WithBuildMessageHook to inject custom metadata into A2A messages.
// The hook wraps the default message converter as middleware, allowing you to
// modify the message before/after conversion, or completely replace the conversion logic.
// The custom metadata will be received by the server's ProcessMessageHook.

// Call the default converter. transferState keys are injected by the outer wrapper.

// Inject custom metadata that will be visible in the server's ProcessMessageHook

// processResponse handles both streaming and non-streaming responses with tool call visualization.
func processResponse(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Check if this is the final event.

// handleEvent processes a single event from the event channel.
func handleEvent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle tool calls (return early to avoid processing tool call content as assistant response)

// Handle tool responses (return early to avoid processing tool response content as assistant response)

// Handle content.

// handleToolCalls detects and displays tool calls.
func handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// trpc-agent-go only puts tool calls in Message.ToolCalls, never in Delta.ToolCalls
// even in streaming mode, tool calls are aggregated and sent in final response

// handleToolResponses detects and displays tool responses.
func handleToolResponses(event *event.Event) bool { _ = "STUB: not implemented"; return false }

// Tool responses are always in Message (never in Delta), even in streaming mode
// This follows trpc-agent-go convention

// handleContent processes and displays content.
func handleContent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

// Display reasoning content first (in cyan color)

// extractContent extracts content and reasoning content based on streaming mode.
func extractContent(choice model.Choice) (content string, reasoningContent string) {
	_ = "STUB: not implemented"
	return "", ""
}

// displayReasoningContent prints reasoning content in cyan color.
func displayReasoningContent(reasoningContent string) { _ = "STUB: not implemented"; return }

// displayContent prints content to console.
func displayContent(
	content string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 {
	_ = "STUB: not implemented"

	// getCurrentTime returns current time information.
	return nil
}

func getCurrentTime(_ context.Context, args timeArgs) (timeResult, error) {
	_ = "STUB: not implemented"
	return *new(timeResult), nil
}

// Handle timezone conversion.

// Simplified EST.

// Simplified PST.

// Simplified CST.

// timeArgs represents arguments for the time tool.
type timeArgs struct {
	Timezone string `json:"timezone" jsonschema:"description=Timezone or leave empty for local"`
}

// timeResult represents the current time information.
type timeResult struct {
	Timezone string `json:"timezone"`
	Time     string `json:"time"`
	Date     string `json:"date"`
	Weekday  string `json:"weekday"`
}

func getEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }
