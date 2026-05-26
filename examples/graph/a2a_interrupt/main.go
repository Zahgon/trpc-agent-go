//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates interrupt/resume across a parent graph and a
// remote graph reached through A2A.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	defaultModelName = "deepseek-v4-flash"

	remoteAgentName = "remote_interrupt_graph"
	parentAgentName = "parent_interrupt_graph"

	remoteNodeRiskSignals    = "remote_risk_signals"
	remoteNodeCaptureCase    = "remote_capture_case_brief"
	remoteNodePrepareVerdict = "remote_prepare_risk_verdict"
	remoteNodeRiskVerdict    = "remote_risk_verdict"
	remoteNodeAsk            = "remote_ask_approval"
	remoteNodeFinalize       = "remote_finalize"

	parentNodeIntake   = "parent_intake"
	parentNodeDecision = "parent_decision_draft"
	parentNodeFinalize = "parent_finalize"

	parentStateKeyCaseBrief     = "case_brief"
	parentStateKeyApproved      = "approved_from_remote"
	parentStateKeyRemoteSummary = "remote_summary"
	parentStateKeyDecisionDraft = "decision_draft"
	parentStateKeyFinalMessage  = "parent_final_message"

	remoteStateKeyCaseBrief    = "remote_case_brief"
	remoteStateKeyRiskSignals  = "remote_risk_signals"
	remoteStateKeyVerdictInput = "remote_risk_verdict_input"
	remoteStateKeyRiskVerdict  = "remote_risk_verdict"
	remoteStateKeyApproved     = "remote_approved"
	remoteStateKeySummary      = "remote_summary"

	defaultLineageID = "demo-a2a-interrupt"
	defaultParentNS  = "parent-a2a-interrupt"
	defaultRemoteNS  = "remote-a2a-interrupt"
	defaultTimeout   = 2 * time.Minute
	defaultInput     = "Payment request: transfer USD 250,000 to a new beneficiary in a high-risk region within 30 minutes."
	pollInterval     = 20 * time.Millisecond
	pollTimeout      = 5 * time.Second
)

var (
	modelName = flag.String("model", getEnvOrDefault("MODEL_NAME", defaultModelName), "OpenAI-compatible model name")
	host      = flag.String("host", "", "A2A host, for example 127.0.0.1:28883 (default: random local port)")
	streaming = flag.Bool("streaming", true, "Use A2A streaming between parent and remote graph")
	timeout   = flag.Duration("timeout", defaultTimeout, "Overall timeout")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Message streaming keeps graph metadata (state_delta)
// on a direct Message event path, which is more robust for interrupt propagation.

func buildRemoteGraphAgent(modelInstance model.Model) (*graphagent.GraphAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildParentGraphAgent(
	subAgent agent.Agent,
	modelInstance model.Model,
) (*graphagent.GraphAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parent_intake leaves the case brief in last_response; use it as the
// child invocation input so remote A2A message content is not empty.

type graphRunResult struct {
	completion         *event.Event
	sawPregelInterrupt bool
	traces             []demoTrace
}

type demoTrace struct {
	scope   string
	node    string
	summary string
}

func runGraphAgent(
	ctx context.Context,
	agt agent.Agent,
	invocationID string,
	userInput string,
	runtimeState graph.State,
) (*graphRunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isGraphCompletionEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

func isPregelInterruptEvent(ev *event.Event) bool { _ = "STUB: not implemented"; return false }

// Pregel interrupt metadata uses "interruptValue" when an interrupt is raised.

func extractSubgraphInterruptInfo(
	values map[string]any,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func decodeJSONString(stateDelta map[string][]byte, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findLatestInterruptedCheckpoint(
	ctx context.Context,
	manager *graph.CheckpointManager,
	lineageID string,
	namespace string,
) (*graph.CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func emitDemoTrace(
	ctx context.Context,
	state graph.State,
	scope string,
	node string,
	summary string,
) {
	_ = "STUB: not implemented"
	return
}

func stringValue(values map[string]any, key string) string { _ = "STUB: not implemented"; return "" }

func decodeDemoTrace(ev *event.Event) (demoTrace, bool) {
	_ = "STUB: not implemented"
	return *new(demoTrace), false
}

func printTraceTranscript(title string, traces []demoTrace) { _ = "STUB: not implemented"; return }

func traceDisplayName(scope, node string) string { _ = "STUB: not implemented"; return "" }

func indentBlockWrapped(s, prefix string, width int) string { _ = "STUB: not implemented"; return "" }

func wrapText(s string, width int) []string { _ = "STUB: not implemented"; return nil }

func buildRemoteReviewSummary(approved bool) string { _ = "STUB: not implemented"; return "" }

func buildRiskSignalsText() string { _ = "STUB: not implemented"; return "" }

func buildRiskVerdictText() string { _ = "STUB: not implemented"; return "" }

type transferRiskAssessment struct {
	level   string
	signals []string
	reason  string
}

func demoRiskAssessment() transferRiskAssessment {
	_ = "STUB: not implemented"
	// Demo simplification: always return a fixed high-risk assessment so the
	// interrupt path and resume behavior are deterministic in examples.
	return *new(transferRiskAssessment)
}

func findLastUserMessage(state graph.State) string { _ = "STUB: not implemented"; return "" }

func printSection(title string) { _ = "STUB: not implemented"; return }

func printKeyValueCard(rows [][2]string) { _ = "STUB: not implemented"; return }

func printParagraph(text string, indent int) { _ = "STUB: not implemented"; return }

func getEnvOrDefault(key, fallback string) string { _ = "STUB: not implemented"; return "" }

func compactGenerationConfig() model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func resolveHost(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func waitForServer(ctx context.Context, url string, serverErr <-chan error) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:noctx
