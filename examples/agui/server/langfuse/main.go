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
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	aguirunner "trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/translator"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/langfuse"
)

const (
	agentName = "agui-agent"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model to use")
	isStream  = flag.Bool("stream", true, "Whether to stream the response")
	address   = flag.String("address", "127.0.0.1:8080", "Listen address")
	path      = flag.String("path", "/agui", "HTTP path")
)

func main() {
	// Start trace with Langfuse integration using environment variables.
	clean, err := langfuse.Start(context.Background())
	if err != nil {
		log.Fatalf("Failed to start trace telemetry: %v", err)
	}
	defer func() {
		if err := clean(context.Background()); err != nil {
			log.Fatalf("Failed to clean up trace telemetry: %v", err)
		}
	}()
	// Parse command line arguments.
	flag.Parse()
	// Build agent and runner.
	agent := newAgent()
	runner := runner.NewRunner(agent.Info().Name, agent)
	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer runner.Close()
	// Build AG-UI server.
	callbacks := translator.NewCallbacks().RegisterAfterTranslate(langfuseCallback())
	server, err := agui.New(runner,
		agui.WithPath(*path),
		agui.WithAGUIRunnerOptions(
			aguirunner.WithUserIDResolver(userIDResolver),
			aguirunner.WithTranslateCallbacks(callbacks),
			aguirunner.WithRunOptionResolver(runOptionResolver),
		),
	)
	if err != nil {
		log.Fatalf("failed to create AG-UI server: %v", err)
	}
	// Start AG-UI server.
	log.Infof("AG-UI: serving agent %q on http://%s%s", agent.Info().Name, *address, *path)
	if err := http.ListenAndServe(*address, server.Handler()); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}

// runOptionResolver resolves the run options for the agent.
func runOptionResolver(ctx context.Context, input *adapter.RunAgentInput) ([]agent.RunOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// langfuseCallback is a callback that sends the output to Langfuse.
func langfuseCallback() translator.AfterTranslateCallback {
	_ = "STUB: not implemented"
	// Store the output for each trace ID.
	return *new(translator.AfterTranslateCallback)
}

// Get the output for a given trace ID, default to empty string.

// Return the callback that sends the output to Langfuse.

// Reset the output.

// Report the output.

// Aggregate the output.

// userIDResolver resolves the user ID from the AG-UI input.
func userIDResolver(ctx context.Context, input *adapter.RunAgentInput) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// newAgent creates a new agent.
		nil
}

func newAgent() agent.Agent { _ = "STUB: not implemented"; return *new(agent.Agent) }

func calculator(ctx context.Context, args calculatorArgs) (calculatorResult, error) {
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

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
