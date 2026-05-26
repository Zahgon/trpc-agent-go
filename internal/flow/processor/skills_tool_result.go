//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	skillsLoadedContextHeader = "Loaded skill context:"

	skillToolLoad       = "skill_load"
	skillToolSelectDocs = "skill_select_docs"
)

type skillsToolResultProcessorOptions struct {
	loadMode                     string
	skipFallbackOnSessionSummary bool
	repoResolver                 func(*agent.Invocation) skill.Repository
	directoryHints               bool
	filePathHints                bool
}

// SkillsToolResultRequestProcessorOption configures
// SkillsToolResultRequestProcessor.
type SkillsToolResultRequestProcessorOption func(
	*skillsToolResultProcessorOptions,
)

// WithSkillsToolResultLoadMode sets how long loaded skill bodies/docs
// remain available in prompt materialization.
//
// Supported modes:
//   - SkillLoadModeTurn (default)
//   - SkillLoadModeOnce
//   - SkillLoadModeSession (legacy)
func WithSkillsToolResultLoadMode(
	mode string,
) SkillsToolResultRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsToolResultRequestProcessorOption)
}

// WithSkipSkillsFallbackOnSessionSummary controls whether the processor
// skips the "Loaded skill context" system-message fallback when a session
// summary is present in the request.
//
// Default: true.
func WithSkipSkillsFallbackOnSessionSummary(
	skip bool,
) SkillsToolResultRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsToolResultRequestProcessorOption)
}

// WithSkillsToolResultRepositoryResolver sets an invocation-aware repository resolver.
func WithSkillsToolResultRepositoryResolver(
	resolver func(*agent.Invocation) skill.Repository,
) SkillsToolResultRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsToolResultRequestProcessorOption)
}

// WithSkillsToolResultDirectoryHints exposes skill directory locators in
// loaded skill materialization.
func WithSkillsToolResultDirectoryHints(
	enable bool,
) SkillsToolResultRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsToolResultRequestProcessorOption)
}

// WithSkillsToolResultFilePathHints exposes SKILL.md file paths in loaded
// skill materialization.
func WithSkillsToolResultFilePathHints(
	enable bool,
) SkillsToolResultRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsToolResultRequestProcessorOption)
}

// SkillsToolResultRequestProcessor materializes loaded skill content
// into tool result messages (skill_load / skill_select_docs) when
// possible.
//
// If no matching tool result message exists (for example, when history
// is suppressed but state persists), it falls back to a dedicated system
// message containing the loaded skill bodies/docs.
//
// If a session summary is present in the request and the corresponding
// option is enabled, the fallback system message is skipped only when the
// loaded skill content is already represented elsewhere in the prompt.
type SkillsToolResultRequestProcessor struct {
	repo         skill.Repository
	repoResolver func(*agent.Invocation) skill.Repository
	loadMode     string

	skipFallbackOnSessionSummary bool
	directoryHints               bool
	filePathHints                bool
}

// NewSkillsToolResultRequestProcessor creates a processor instance.
func NewSkillsToolResultRequestProcessor(
	repo skill.Repository,
	opts ...SkillsToolResultRequestProcessorOption,
) *SkillsToolResultRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements flow.RequestProcessor.
func (p *SkillsToolResultRequestProcessor) ProcessRequest(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// SupportsContextCompactionRebuild reports whether loaded skill materialization
// can be safely replayed during the sync-summary rebuild path.
func (p *SkillsToolResultRequestProcessor) SupportsContextCompactionRebuild(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

// RebuildRequestForContextCompaction reapplies loaded skill materialization
// without mutating session state during the sync-summary rebuild path.
func (p *SkillsToolResultRequestProcessor) RebuildRequestForContextCompaction(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (p *SkillsToolResultRequestProcessor) applyLoadedSkillContext(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
	repo skill.Repository,
) []string {
	_ = "STUB: not implemented"
	return nil
}

// stable prompt order

func (p *SkillsToolResultRequestProcessor) repositoryForInvocation(
	inv *agent.Invocation,
) skill.Repository {
	_ = "STUB: not implemented"
	return *new(skill.Repository)
}

func hasSessionSummary(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

func (p *SkillsToolResultRequestProcessor) getLoadedSkills(
	inv *agent.Invocation,
) []string {
	_ = "STUB: not implemented"
	return nil
}

type toolCallIndex map[string]model.ToolCall

func indexToolCalls(msgs []model.Message) toolCallIndex {
	_ = "STUB: not implemented"
	return *new(toolCallIndex)
}

func lastSkillToolMsgIndex(
	msgs []model.Message,
	calls toolCallIndex,
) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

type skillNameInput struct {
	Skill string `json:"skill"`
}

func skillNameFromToolMessage(
	m model.Message,
	calls toolCallIndex,
) string {
	_ = "STUB: not implemented"
	return ""
}

// Fallback: parse the short tool output ("loaded: <name>") if
// available.

const loadedPrefix = "loaded:"

func parseLoadedSkillFromText(content string) string { _ = "STUB: not implemented"; return "" }

func isLoadedToolStub(toolOutput string, skillName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *SkillsToolResultRequestProcessor) buildToolResultContent(
	ctx context.Context,
	inv *agent.Invocation,
	repo skill.Repository,
	skillName string,
	toolOutput string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *SkillsToolResultRequestProcessor) getDocsSelection(
	ctx context.Context,
	inv *agent.Invocation,
	repo skill.Repository,
	name string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func buildDocsText(sk *skill.Skill, wanted []string) string { _ = "STUB: not implemented"; return "" }

func (p *SkillsToolResultRequestProcessor) buildFallbackSystemContent(
	ctx context.Context,
	inv *agent.Invocation,
	repo skill.Repository,
	loaded []string,
	materialized map[string]struct{},
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *SkillsToolResultRequestProcessor) appendSkillPathHints(
	b *strings.Builder,
	ctx context.Context,
	repo skill.Repository,
	name string,
) {
	_ = "STUB: not implemented"
	return
}

func (p *SkillsToolResultRequestProcessor) upsertLoadedContextMessage(
	req *model.Request,
	content string,
) {
	_ = "STUB: not implemented"
	return
}

func (p *SkillsToolResultRequestProcessor) removeLoadedContextMessage(
	req *model.Request,
) {
	_ = "STUB: not implemented"
	return
}

func findLoadedContextMessageIndex(msgs []model.Message) int { _ = "STUB: not implemented"; return 0 }

func insertAfterLastSystemMessage(
	req *model.Request,
	msg model.Message,
) {
	_ = "STUB: not implemented"
	return
}

func (p *SkillsToolResultRequestProcessor) maybeOffloadLoadedSkills(
	ctx context.Context,
	inv *agent.Invocation,
	loaded []string,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}
