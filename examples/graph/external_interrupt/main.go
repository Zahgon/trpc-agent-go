//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates external graph interrupts ("pause button") and
// resumable checkpoints.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	defaultModelName = "deepseek-v4-flash"
	defaultDemoMode  = "both"
	defaultUserText  = "Say one short sentence about graphs."
	defaultEngine    = "bsp"

	envOpenAIAPIKey = "OPENAI_API_KEY"
	envModelName    = "MODEL_NAME"
)

const (
	demoPlanned = "planned"
	demoForced  = "forced"
	demoBoth    = "both"
)

const (
	engineBSP = "bsp"
	engineDAG = "dag"
)

const (
	nodePrepare   = "prepare"
	nodeCallModel = "call_model"
	nodeFinalize  = "finalize"

	nodeSlow = "slow"
	nodeDone = "done"
)

const (
	stateKeyResult = "result"
	stateKeySlowOK = "slow_ok"
)

const (
	waitStartedTimeout  = 2 * time.Second
	prepareSleep        = 300 * time.Millisecond
	slowWorkDuration    = 300 * time.Millisecond
	forcedInterruptWait = 50 * time.Millisecond
)

var (
	demoMode = flag.String(
		"demo",
		defaultDemoMode,
		"Demo to run: planned|forced|both",
	)
	engine = flag.String(
		"engine",
		defaultEngine,
		"Execution engine: bsp|dag",
	)
	modelName = flag.String(
		"model",
		defaultModelFromEnv(),
		"Model name used by planned demo",
	)
	userText = flag.String(
		"text",
		defaultUserText,
		"User text for the planned demo",
	)
)

type interruptMeta struct {
	NodeID         string          `json:"nodeID,omitempty"`
	InterruptKey   string          `json:"interruptKey,omitempty"`
	LineageID      string          `json:"lineageId,omitempty"`
	CheckpointID   string          `json:"checkpointId,omitempty"`
	InterruptValue json.RawMessage `json:"interruptValue,omitempty"`
}

func main() {
	flag.Parse()

	mode := strings.ToLower(strings.TrimSpace(*demoMode))
	if mode == "" {
		mode = defaultDemoMode
	}

	switch mode {
	case demoPlanned:
		runPlannedDemo()
	case demoForced:
		runForcedDemo()
	case demoBoth:
		runPlannedDemo()
		runForcedDemo()
	default:
		log.Fatalf("unknown -demo value: %q", *demoMode)
	}
}

func runPlannedDemo() { _ = "STUB: not implemented"; return }

func runForcedDemo() { _ = "STUB: not implemented"; return }

func defaultModelFromEnv() string { _ = "STUB: not implemented"; return "" }

func buildPlannedGraph(
	started chan<- struct{},
	modelName string,
	userText string,
) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func modelInstruction() string { _ = "STUB: not implemented"; return "" }

func generationConfig() model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func buildForcedGraph(started chan<- struct{}) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func waitOrCancel(ctx context.Context, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func newInvocation(lineageID string) *agent.Invocation { _ = "STUB: not implemented"; return nil }

func waitForStartedOrExit(started <-chan struct{}) { _ = "STUB: not implemented"; return }

func drainEvents(
	ch <-chan *event.Event,
) (*interruptMeta, *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func extractInterruptMeta(evt *event.Event) *interruptMeta { _ = "STUB: not implemented"; return nil }

func decodeExternalPayload(
	meta *interruptMeta,
) (graph.ExternalInterruptPayload, bool) {
	_ = "STUB: not implemented"
	return *new(graph.ExternalInterruptPayload), false
}

func decodeStringState(done *event.Event, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func decodeBoolState(done *event.Event, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func assistantTextFromState(st graph.State) string { _ = "STUB: not implemented"; return "" }

func textFromParts(parts []model.ContentPart) string { _ = "STUB: not implemented"; return "" }

func shorten(s string, max int) string { _ = "STUB: not implemented"; return "" }

func parseEngineOrExit() graph.ExecutionEngine {
	_ = "STUB: not implemented"
	return *new(graph.ExecutionEngine)
}
