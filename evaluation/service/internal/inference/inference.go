//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package inference runs agent sessions to generate invocations for evaluation.
package inference

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/usersimulation"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Result contains all inference artifacts for one eval case.
type Result struct {
	Invocations     []*evalset.Invocation
	ExecutionTraces []*trace.Trace
}

// Inference executes the agent against the provided invocations.
func Inference(
	ctx context.Context,
	runner runner.Runner,
	invocations []*evalset.Invocation,
	initialSession *evalset.SessionInput,
	sessionID string,
	runOptions []agent.RunOption,
) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accumulate each invocation response.

// InferenceWithConversationScenario executes the agent against a simulated multi-turn conversation.
func InferenceWithConversationScenario(
	ctx context.Context,
	r runner.Runner,
	simulator usersimulation.Simulator,
	evalCaseID string,
	scenario *evalset.ConversationScenario,
	initialSession *evalset.SessionInput,
	sessionID string,
	runOptions []agent.RunOption,
) (result *Result, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// inferenceInvocation executes the agent for a single invocation.
func inferenceInvocation(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
	initialSession *evalset.SessionInput,
	invocation *evalset.Invocation,
	runOptions []agent.RunOption,
) (*evalset.Invocation, *trace.Trace, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Capture the invocation ID, final response, tool uses, and tool responses.

// Capture tool call uses.

// Capture tool call responses.

func eventFinalResponse(evt *event.Event) *model.Message { _ = "STUB: not implemented"; return nil }

// convertTools converts the tool call to tools.
func convertTools(event *event.Event) ([]*evalset.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseToolCallArguments(arguments []byte) any { _ = "STUB: not implemented"; return *new(any) }

// mergeToolResultResponse merges the tool result response into the tools.
func mergeToolResultResponse(event *event.Event, toolIDIdx map[string]int, tools []*evalset.Tool) error {
	_ = "STUB: not implemented"
	return nil
}

func parseToolResultContent(content string) any { _ = "STUB: not implemented"; return *new(any) }
