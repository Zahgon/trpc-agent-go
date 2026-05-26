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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// SimpleIntentAgent is a minimal custom Agent implementation that demonstrates
// how to embed business logic flow without using Graph.
//
// Flow:
//  1. Run a quick intent classification using the LLM: "chitchat" or "task".
//  2. If chitchat, directly answer the user.
//  3. If task, provide a step plan (or route to downstream logic/tools in real apps).
type SimpleIntentAgent struct {
	name        string
	description string
	model       model.Model
}

func NewSimpleIntentAgent(name, description string, m model.Model) *SimpleIntentAgent {
	_ = "STUB: not implemented"
	return nil
}

// Info implements agent.Agent.
func (a *SimpleIntentAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// Tools implements agent.Agent. No tools in this minimal example.
func (a *SimpleIntentAgent) Tools() []tool.Tool {
	_ = "STUB: not implemented"

	// SubAgents implements agent.Agent. This example has no sub-agents.
	return nil
}

func (a *SimpleIntentAgent) SubAgents() []agent.Agent {
	_ = "STUB: not implemented"

	// FindSubAgent implements agent.Agent.
	return nil
}

func (a *SimpleIntentAgent) FindSubAgent(string) agent.Agent {
	_ = "STUB: not implemented"

	// Run implements agent.Agent.
	return *new(agent.Agent)
}

func (a *SimpleIntentAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Align invocation metadata for downstream consumers (consistency with LLMAgent behavior).

// 1) Intent classification (not forwarded to user).

// 2) Branch on intent and stream the chosen LLM response back to the user.

// Fallback if classifier uncertain.

// classifyIntent asks the LLM to classify the user input.
func (a *SimpleIntentAgent) classifyIntent(ctx context.Context, inv *agent.Invocation) string {
	_ = "STUB: not implemented"
	// System prompt keeps output constrained to a single label.
	return ""
}

// On error, default to chitchat to remain robust.

// Accumulate both deltas and final message content for non-streaming providers.

// replyChitChat streams a direct conversational answer.
func (a *SimpleIntentAgent) replyChitChat(ctx context.Context, inv *agent.Invocation, out chan<- *event.Event) {
	_ = "STUB: not implemented"
	return
}

// replyTaskPlan streams a simple step plan for task-like requests.
func (a *SimpleIntentAgent) replyTaskPlan(ctx context.Context, inv *agent.Invocation, out chan<- *event.Event) {
	_ = "STUB: not implemented"
	return
}

// sanitizeIntent normalizes the LLM output for intent classification.
func sanitizeIntent(s string) string { _ = "STUB: not implemented"; return "" }

// Keep just the first line in case the model added explanations.

// Trim common punctuation / quotes around a single token.

// Collapse spaces.
