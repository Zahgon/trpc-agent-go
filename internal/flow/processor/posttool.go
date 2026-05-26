//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// DefaultPostToolPrompt is the default dynamic prompt injected after tool
// results to guide the model's response behavior. Inspired by CrewAI's
// post_tool_reasoning mechanism.
const DefaultPostToolPrompt = "[Tool Prompt] Analyze the tool result. " +
	"If requirements are met, provide the Final Answer. " +
	"Otherwise, call the next tool. \n" +
	"Final Answer Requirement:\n" +
	"Only output the exact answer the user needs. " +
	"Answer as if you already know the information. Do NOT expose any internal process, tool usage, or source of information. " +
	"Do not use phrases that reference tools, searches, or retrieved data, like \"according to xxx, based on xxx\". Keep your answer concise and to the point.\n"

// PostToolRequestProcessor inspects the built req.Messages for pending
// tool-result messages. When the current request still ends at a tool result,
// it appends the dynamic prompt to the existing system message to steer the
// model's next response. This avoids inserting a user message between tool
// results and the assistant turn, which is not part of the standard OpenAI
// message ordering contract.
//
// Inspired by CrewAI's post_tool_reasoning mechanism, but injected via
// system prompt instead of user message for better cross-model compatibility.
//
// This processor MUST be registered after ContentRequestProcessor so that
// req.Messages already contains the full conversation history including
// any tool results.
type PostToolRequestProcessor struct {
	// prompt is the dynamic prompt text appended to the system message
	// when tool results are detected. When empty, no prompt is injected.
	prompt string
}

// PostToolOption is a functional option for PostToolRequestProcessor.
type PostToolOption func(*PostToolRequestProcessor)

// WithPostToolPrompt overrides the default dynamic prompt text.
// Set to empty string to disable prompt injection entirely.
func WithPostToolPrompt(prompt string) PostToolOption {
	_ = "STUB: not implemented"
	return *new(PostToolOption)
}

// NewPostToolRequestProcessor creates a new PostToolRequestProcessor.
func NewPostToolRequestProcessor(opts ...PostToolOption) *PostToolRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements flow.RequestProcessor.
// It checks whether req.Messages still has pending tool results from the
// active tool loop. If so, it appends the dynamic prompt to the system
// message.
func (p *PostToolRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// No system message exists; prepend one.

// SupportsContextCompactionRebuild reports that post-tool prompting can be
// safely replayed during the sync-summary rebuild path.
func (p *PostToolRequestProcessor) SupportsContextCompactionRebuild(
	_ *agent.Invocation,
) bool {
	_ = "STUB: not implemented"

	// RebuildRequestForContextCompaction re-applies post-tool prompting during the
	// safe sync-summary rebuild path without replaying the full processor chain.
	return false
}

func (p *PostToolRequestProcessor) RebuildRequestForContextCompaction(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
) {
	_ = "STUB: not implemented"
	return
}

// hasPendingToolResultMessages returns true when the latest non-system message
// in the request is a tool result. Historical tool results that are followed
// by assistant or user messages do not count as pending.
func hasPendingToolResultMessages(msgs []model.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func hasCompactedToolResultMessages(inv *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}
