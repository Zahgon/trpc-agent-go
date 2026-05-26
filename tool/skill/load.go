//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package skill provides skill-related tools (function calls)
// for loading skills on demand.
package skill

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// stateDeltaProvider is consumed by the flow to attach state delta
// on tool.response events.
type stateDeltaProvider interface {
	StateDelta(toolCallID string, args []byte, resultJSON []byte) map[string][]byte
}

// LoadTool enables loading a skill into session state.
// It produces deltas under prefixes defined by skill package.
type LoadTool struct {
	repo        skill.Repository
	description string
}

const defaultLoadToolDescription = "Load a skill body and optional docs. " +
	"Prefer progressive disclosure: load SKILL.md first, " +
	"then load only needed docs. " +
	"Safe to call multiple times to add or replace docs. " +
	"Do not call this to list skills; names and descriptions " +
	"are already in context. Use when a task needs a skill's " +
	"SKILL.md body and selected docs in context."

type loadToolOptions struct {
	description string
}

// LoadToolOption configures LoadTool.
type LoadToolOption func(*loadToolOptions)

// WithLoadToolDescription overrides the skill_load tool description.
func WithLoadToolDescription(
	description string,
) LoadToolOption {
	_ = "STUB: not implemented"
	return *new(LoadToolOption)
}

// NewLoadTool creates a new LoadTool.
func NewLoadTool(repo skill.Repository) *LoadTool { _ = "STUB: not implemented"; return nil }

// NewLoadToolWithOptions creates a new LoadTool with optional overrides.
func NewLoadToolWithOptions(
	repo skill.Repository,
	opts ...LoadToolOption,
) *LoadTool {
	_ = "STUB: not implemented"
	return nil
}

// loadInput is the schema for skill_load.
type loadInput struct {
	Skill          string   `json:"skill"`
	Docs           []string `json:"docs,omitempty"`
	IncludeAllDocs bool     `json:"include_all_docs,omitempty"`
}

// Declaration implements tool.Tool.
func (t *LoadTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call validates and returns a message for user feedback.
func (t *LoadTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// validate existence

// StateDelta builds delta keys to mark loaded skill and doc selection.
func (t *LoadTool) StateDelta(_ string, args []byte, _ []byte) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// StateDeltaForInvocation writes agent-scoped state for the invocation.
func (t *LoadTool) StateDeltaForInvocation(
	inv *agent.Invocation,
	toolCallID string,
	args []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func (t *LoadTool) stateDelta(
	agentName string,
	args []byte,
) (map[string][]byte, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// Mark as loaded.

// Docs selection

var _ tool.Tool = (*LoadTool)(nil)
var _ tool.CallableTool = (*LoadTool)(nil)
var _ stateDeltaProvider = (*LoadTool)(nil)
