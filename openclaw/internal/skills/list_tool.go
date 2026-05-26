//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skills

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const skillListToolName = "skill_list"

const (
	skillListModeAll      = "all"
	skillListModeEnabled  = "enabled"
	skillListModeDisabled = "disabled"
)

type listInput struct {
	Mode string `json:"mode,omitempty"`
}

type listOutput struct {
	Total    int          `json:"total"`
	Enabled  int          `json:"enabled"`
	Disabled int          `json:"disabled"`
	Skills   []skillEntry `json:"skills,omitempty"`
}

type skillEntry struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
	Reason      string `json:"reason,omitempty"`

	Emoji    string `json:"emoji,omitempty"`
	Homepage string `json:"homepage,omitempty"`

	Requires *skillRequires `json:"requires,omitempty"`
}

type skillRequires struct {
	OS      []string `json:"os,omitempty"`
	Bins    []string `json:"bins,omitempty"`
	AnyBins []string `json:"any_bins,omitempty"`
	Env     []string `json:"env,omitempty"`
	Config  []string `json:"config,omitempty"`
}

// ListTool lists all discovered skills (enabled + disabled) with a small
// eligibility summary. It is intended for discovery and debugging.
type ListTool struct {
	repo *Repository
}

func NewListTool(repo *Repository) *ListTool { _ = "STUB: not implemented"; return nil }

func (t *ListTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *ListTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func normalizeListMode(raw string) string { _ = "STUB: not implemented"; return "" }

func (t *ListTool) buildOutput(mode string) listOutput {
	_ = "STUB: not implemented"
	return *new(listOutput)
}

func (t *ListTool) entryForSummary(s skill.Summary) skillEntry {
	_ = "STUB: not implemented"
	return *new(skillEntry)
}

func normalizeSkillRequires(meta openClawMetadata) *skillRequires {
	_ = "STUB: not implemented"
	return nil
}

var _ tool.Tool = (*ListTool)(nil)
var _ tool.CallableTool = (*ListTool)(nil)
