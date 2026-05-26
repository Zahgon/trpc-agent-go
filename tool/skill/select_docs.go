//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package skill provides skill-related tools (function calls).
package skill

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	selectDocsToolName = "skill_select_docs"
	modeAdd            = "add"
	modeReplace        = "replace"
	modeClear          = "clear"
)

type selectDocsInput struct {
	Skill          string   `json:"skill"`
	Docs           []string `json:"docs,omitempty"`
	IncludeAllDocs bool     `json:"include_all_docs,omitempty"`
	Mode           string   `json:"mode,omitempty"`
}

type selectDocsOutput struct {
	Skill          string   `json:"skill"`
	Selected       []string `json:"selected_docs,omitempty"`
	IncludeAllDocs bool     `json:"include_all_docs,omitempty"`
	Mode           string   `json:"mode,omitempty"`
}

// SelectDocsTool updates doc selection for a loaded skill.
type SelectDocsTool struct {
	repo skill.Repository
}

// NewSelectDocsTool creates a SelectDocsTool.
func NewSelectDocsTool(repo skill.Repository) *SelectDocsTool {
	_ = "STUB: not implemented"
	return nil
}

// Declaration implements tool.Tool.
func (t *SelectDocsTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call computes the final selection (may read session state for add).
func (t *SelectDocsTool) Call(
	ctx context.Context, args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// StateDelta writes the selection using the tool result JSON.
func (t *SelectDocsTool) StateDelta(
	_ string, _ []byte, resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// StateDeltaForInvocation writes agent-scoped state for the invocation.
func (t *SelectDocsTool) StateDeltaForInvocation(
	inv *agent.Invocation,
	toolCallID string,
	args []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func (t *SelectDocsTool) stateDelta(
	agentName string,
	resultJSON []byte,
) (map[string][]byte, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// Ensure empty slice encodes to [] rather than null.

var _ tool.Tool = (*SelectDocsTool)(nil)
var _ tool.CallableTool = (*SelectDocsTool)(nil)

// parseSelectArgs validates and normalizes the input.
func (t *SelectDocsTool) parseSelectArgs(
	ctx context.Context,
	args []byte,
) (selectDocsInput, error) {
	_ = "STUB: not implemented"
	return *new(selectDocsInput), nil
}

// previousSelection reads any prior selection from session state.
func (t *SelectDocsTool) previousSelection(
	ctx context.Context, skillName string,
) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *SelectDocsTool) outClear(
	in selectDocsInput,
) selectDocsOutput {
	_ = "STUB: not implemented"
	return *new(selectDocsOutput)
}

func (t *SelectDocsTool) outAdd(
	prev []string, in selectDocsInput,
) selectDocsOutput {
	_ = "STUB: not implemented"
	return *new(selectDocsOutput)
}

func (t *SelectDocsTool) outReplace(
	in selectDocsInput,
) selectDocsOutput {
	_ = "STUB: not implemented"
	return *new(selectDocsOutput)
}
