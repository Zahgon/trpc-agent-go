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

	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const listDocsToolName = "skill_list_docs"

type listDocsInput struct {
	Skill string `json:"skill"`
}

// ListDocsTool lists available docs for a skill.
type ListDocsTool struct {
	repo skill.Repository
}

// NewListDocsTool creates a ListDocsTool.
func NewListDocsTool(repo skill.Repository) *ListDocsTool { _ = "STUB: not implemented"; return nil }

// Declaration implements tool.Tool.
func (t *ListDocsTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call returns the list of doc filenames.
func (t *ListDocsTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var _ tool.Tool = (*ListDocsTool)(nil)
var _ tool.CallableTool = (*ListDocsTool)(nil)
