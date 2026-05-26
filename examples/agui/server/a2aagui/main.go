//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates exposing a remote A2A Agent through an AG-UI server.
//
// Architecture:
//
//	[Browser / AG-UI Client]
//	         | AG-UI (SSE)
//	         v
//	  [AG-UI Server :8080]
//	         | uses A2AAgent as runner backend
//	         v
//	  [A2A Agent Client]
//	         | A2A protocol (HTTP/JSON-RPC)
//	         v
//	  [A2A Server :8888]
//	         |
//	         v
//	  [Local LLM Agent + Calculator Tool]
//
// Usage:
//
//	go run ./examples/agui/a2aagui -model deepseek-v4-flash
//
// Then connect any AG-UI frontend (e.g. CopilotKit) to http://localhost:8080/agui
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/a2aagent"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	aguirunner "trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
)

var (
	modelName    = flag.String("model", getEnvOrDefault("MODEL_NAME", "deepseek-v4-flash"), "LLM model name")
	a2aAddr      = flag.String("a2a-addr", "0.0.0.0:8888", "A2A server listen address")
	aguiAddr     = flag.String("agui-addr", "0.0.0.0:8080", "AG-UI server listen address")
	aguiPath     = flag.String("agui-path", "/agui", "AG-UI endpoint path")
	enableStream = flag.Bool("stream", true, "Enable streaming")
)

const appName = "a2a-agui-demo"

func main() {
	flag.Parse()

	// Step 1: Start an A2A Server hosting a local Agent with tools.
	startA2AServer()
	time.Sleep(500 * time.Millisecond)

	// Step 2: Create an A2AAgent pointing to the remote A2A Server.
	a2aURL := fmt.Sprintf("http://%s", *a2aAddr)
	remoteAgent, err := a2aagent.New(
		a2aagent.WithAgentCardURL(a2aURL),
		a2aagent.WithEnableStreaming(*enableStream),
	)
	if err != nil {
		log.Fatalf("Failed to create A2AAgent: %v", err)
	}
	log.Infof("A2AAgent connected to %s (agent: %s)", a2aURL, remoteAgent.Info().Name)

	// Step 3: Wrap the A2AAgent in a Runner + AG-UI Server.
	sessionService := inmemory.NewSessionService()
	r := runner.NewRunner(appName, remoteAgent, runner.WithSessionService(sessionService))
	defer r.Close()

	aguiServer, err := agui.New(
		r,
		agui.WithPath(*aguiPath),
		agui.WithAppName(appName),
		agui.WithSessionService(sessionService),
		agui.WithAGUIRunnerOptions(aguirunner.WithUserIDResolver(userIDResolver)),
	)
	if err != nil {
		log.Fatalf("Failed to create AG-UI server: %v", err)
	}

	log.Infof("AG-UI server listening on http://%s%s", *aguiAddr, *aguiPath)
	log.Infof("Connect any AG-UI frontend (e.g. CopilotKit) to this endpoint.")
	if err := http.ListenAndServe(*aguiAddr, aguiServer.Handler()); err != nil {
		log.Fatalf("AG-UI server stopped: %v", err)
	}
}

// startA2AServer creates and starts an A2A Server with a local LLM Agent.
func startA2AServer() { _ = "STUB: not implemented"; return }

func userIDResolver(_ context.Context, input *adapter.RunAgentInput) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func calculator(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

type calculatorArgs struct {
	Operation string  `json:"operation" description:"add, subtract, multiply, divide, power"`
	A         float64 `json:"a" description:"First number"`
	B         float64 `json:"b" description:"Second number"`
}

type calculatorResult struct {
	Result float64 `json:"result"`
}

func intPtr(i int) *int                               { _ = "STUB: not implemented"; return nil }
func floatPtr(f float64) *float64                     { _ = "STUB: not implemented"; return nil }
func getEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }
