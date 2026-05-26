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
	"time"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/a2aagent"
	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	modelName = flag.String("model", getEnvOrDefault("MODEL_NAME", "deepseek-v4-flash"), "Model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming output")
)

const (
	appName            = "a2aagent-customdatapart-demo"
	customEventTag     = "demo.custom_data"
	customDataPartType = "custom_data"
	customDataPartKind = "custom_part_kind"
	customEventExtKey  = "trpc.a2a.custom_payload"
	customEventHint    = "Custom data part data"
	colorReset         = "\033[0m"
	colorCyan          = "\033[36m"
)

type customPayload struct {
	TraceID string `json:"trace_id"`
	Source  string `json:"source"`
	Hint    string `json:"hint"`
}

func main() {
	flag.Parse()

	httpURL, err := runA2AServer()
	if err != nil {
		log.Fatalf("failed to start a2a server: %v", err)
	}
	a2aAgent := buildA2AAgent(httpURL)
	startChat(a2aAgent)
}

func startChat(a2aAgent *a2aagent.A2AAgent) { _ = "STUB: not implemented"; return }

func runA2AServer() (string, error) { _ = "STUB: not implemented"; return "", nil }

type customDataPartWrapper struct {
	base agent.Agent
	name string
}

func wrapAgentWithCustomDataPart(base agent.Agent, name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func (w *customDataPartWrapper) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always forward the original event stream unchanged so the wrapped
// agent keeps its normal behavior.

// Emit the custom event only after the original response stream finishes,
// so the UI prints the structured payload after the normal assistant text.

func (w *customDataPartWrapper) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (w *customDataPartWrapper) Info() agent.Info {
	_ = "STUB: not implemented"
	return *new(agent.Info)
}

func (w *customDataPartWrapper) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

func (w *customDataPartWrapper) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newCustomDataPartEvent(invocationID, author, hint string) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func customDataPartMapper(ctx context.Context, evt *event.Event) ([]protocol.Part, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildAgent(agentName, desc string, extraOptions ...llmagent.Option) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func buildA2AAgent(httpURL string) *a2aagent.A2AAgent { _ = "STUB: not implemented"; return nil }

// Restore the custom DataPart back into event.Extensions so downstream
// graph logic and UI code can consume the structured payload directly.

// Rehydrate the wire-format DataPart payload back into event.Extensions
// so the rest of the local pipeline can consume typed structured data.

func processResponse(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

func handleEvent(evt *event.Event, assistantStarted *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func printMappedCustomHint(evt *event.Event) error { _ = "STUB: not implemented"; return nil }

func printAssistantContent(evt *event.Event, assistantStarted *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func extractContent(choice model.Choice) (string, string) { _ = "STUB: not implemented"; return "", "" }

func eventContent(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func allocateDemoHost() (string, error) { _ = "STUB: not implemented"; return "", nil }

func waitForAgentCardReady(host string, serverErrCh <-chan error, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }

func markAsCustomDataPart(part *protocol.DataPart) { _ = "STUB: not implemented"; return }

func customDataPartPayload(part *protocol.DataPart) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func getEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }
