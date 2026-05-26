//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates a parent GraphAgent calling a remote GraphAgent
// through an A2A sub-agent, then reading the remote agent's final state back
// from the parent graph state.
//
// Unlike the minimal smoke-test version, this example uses a real LLM node in
// the remote graph so it matches the shape of other graph examples more
// closely. The remote graph writes both a natural-language reply and a
// structured payload into graph state; the parent graph receives those fields
// through A2A state_delta transport and confirms the handoff succeeded.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	defaultModelName = "deepseek-v4-flash"

	parentAgentName = "parent_graph"
	remoteAgentName = "remote_graph"

	remoteNodeInput    = "stash_remote_input"
	remoteNodeModel    = "remote_reply"
	remoteNodeCapture  = "capture_remote_state"
	parentNodeFinalize = "finalize"

	remoteStateKeyOriginalInput = "remote_original_input"
	remoteStateKeyValue         = "remote_child_value"
	remoteStateKeyPayload       = "remote_child_payload"

	parentStateKeyValue      = "value_from_remote"
	parentStateKeyEcho       = "echo_from_remote"
	parentStateKeyPayload    = "remote_state_payload"
	parentStateKeyRawDeltaOK = "raw_state_delta_present"

	defaultInput       = "Please explain why state handoff through the remote agent matters."
	defaultRunTimeout  = 90 * time.Second
	serverPollInterval = 20 * time.Millisecond
	serverPollTimeout  = 5 * time.Second

	remoteTransportValue = "a2a"
)

var (
	modelName      = flag.String("model", getEnvOrDefault("MODEL_NAME", defaultModelName), "OpenAI-compatible model name")
	baseURL        = flag.String("base-url", os.Getenv("OPENAI_BASE_URL"), "OpenAI-compatible base URL")
	apiKey         = flag.String("api-key", os.Getenv("OPENAI_API_KEY"), "API key")
	input          = flag.String("input", defaultInput, "Input sent into the parent graph")
	host           = flag.String("host", "", "Host for the in-process A2A server, for example 127.0.0.1:28883")
	streaming      = flag.Bool("streaming", true, "Use A2A streaming between parent and remote graph")
	modelStreaming = flag.Bool("model-streaming", false, "Use streaming when the remote graph calls the model")
	timeout        = flag.Duration("timeout", defaultRunTimeout, "Overall timeout for the example run")
	verboseEvents  = flag.Bool("verbose-events", false, "Print every event observed from the parent graph run")
)

func main() {
	flag.Parse()
	setupLogging()

	fmt.Printf("Graph A2A Agent Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 56))
	if *apiKey == "" && os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("Hint: provide -api-key/-base-url or set OPENAI_API_KEY/OPENAI_BASE_URL.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	if err := run(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Fprintf(
				os.Stderr,
				"error: %v\nhint: try a longer -timeout or disable remote model streaming with -model-streaming=false\n",
				err,
			)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Poll the agent-card endpoint until the server is ready, instead of
// using a fixed sleep that may be too short on slow machines.

// waitForServer polls the given URL until it returns HTTP 200, the server
// goroutine reports an error, or the context is cancelled.
func waitForServer(ctx context.Context, url string, serverErr <-chan error) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:noctx

func buildRemoteGraphAgent(modelName, baseURL, apiKey string, modelStreaming bool) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func stashRemoteInput(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func buildRemoteStateCaptureNode(modelName string) graph.NodeFunc {
	_ = "STUB: not implemented"
	return *new(graph.NodeFunc)
}

func buildParentGraphAgent(remote agent.Agent) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func mapRemoteFinalState(_ graph.State, result graph.SubgraphResult) graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}

func finalizeParentState(_ context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
