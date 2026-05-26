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
	"trpc.group/trpc-go/trpc-agent-go/internal/skillprofile"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	skillsOverviewHeader = "Available skills:"

	skillsCapabilityHeader = "Skill tool availability:"

	skillsToolingGuidanceHeader = "Tooling and workspace guidance:"

	skillRootsHeader = "Skill roots:"
	skillDirLabel    = "Skill dir: "
	skillFileLabel   = "Skill file: "

	// SkillLoadModeOnce injects loaded skill content for the next model
	// request, then offloads it from session state.
	SkillLoadModeOnce = "once"
	// SkillLoadModeTurn keeps loaded skill content available for all model
	// requests within the current invocation, and offloads it when the next
	// invocation begins.
	SkillLoadModeTurn = "turn"
	// SkillLoadModeSession keeps loaded skill content available across
	// invocations until cleared or the session expires.
	SkillLoadModeSession = "session"

	defaultSkillLoadMode = SkillLoadModeTurn
)

type skillsRequestProcessorOptions struct {
	capabilityGuidance *string
	protocolGuidance   *string
	toolingGuidance    *string
	loadMode           string
	toolResultMode     bool
	maxLoadedSkills    int
	toolProfile        string
	toolFlags          skillprofile.Flags
	toolFlagsResolver  func(*agent.Invocation) skillprofile.Flags
	hasToolFlags       bool
	execToolsDisabled  bool
	repoResolver       func(*agent.Invocation) skill.Repository
	directoryHints     bool
	filePathHints      bool
}

// SkillsRequestProcessorOption configures SkillsRequestProcessor.
type SkillsRequestProcessorOption func(*skillsRequestProcessorOptions)

// WithSkillLoadMode sets how long loaded skill bodies/docs remain
// available in the system prompt.
//
// Supported modes:
//   - SkillLoadModeTurn (default)
//   - SkillLoadModeOnce
//   - SkillLoadModeSession (legacy)
func WithSkillLoadMode(mode string) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsToolingGuidance overrides the tooling/workspace guidance
// block appended to the skills overview.
//
// Behavior:
//   - Not configured: use the built-in default guidance.
//   - Configured with empty string: omit both the tooling/workspace
//     guidance block and the capability disclosure block.
//   - Configured with non-empty string: append the provided text.
func WithSkillsToolingGuidance(
	guidance string,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsCapabilityGuidance overrides the capability disclosure block
// appended to the skills overview.
//
// Behavior:
//   - Not configured: use the built-in default disclosure.
//   - Configured with empty string: omit the capability block.
//   - Configured with non-empty string: append the provided text.
func WithSkillsCapabilityGuidance(
	guidance string,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsProtocolGuidance overrides the full skill protocol block
// appended after the skills overview.
//
// Behavior:
//   - Not configured: use the built-in capability/tooling guidance flow.
//   - Configured with empty string: omit all built-in skill guidance.
//   - Configured with non-empty string: append the provided text and skip
//     the built-in capability/tooling guidance blocks.
func WithSkillsProtocolGuidance(
	guidance string,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsLoadedContentInToolResults enables an alternative injection
// mode where loaded SKILL.md bodies and selected docs are materialized
// into the corresponding tool result messages
// (skill_load / skill_select_docs) instead of being appended to the
// system prompt.
//
// This keeps the system prompt more stable for prompt caching while
// preserving the progressive disclosure behavior.
func WithSkillsLoadedContentInToolResults(
	enable bool,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillToolProfile configures the registered skill tool profile so the
// processor can emit mode-appropriate guidance.
func WithSkillToolProfile(profile string) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillToolFlags configures the exact built-in skill tool capabilities so
// the processor can emit guidance that matches the final registered tool set.
func WithSkillToolFlags(flags skillprofile.Flags) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillToolFlagsResolver sets an invocation-aware skill tool capability resolver.
func WithSkillToolFlagsResolver(
	resolver func(*agent.Invocation) skillprofile.Flags,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillExecToolsDisabled tells the processor that skill_exec and its
// companion session tools were not registered (e.g. because the executor
// does not support interactive sessions).  The processor omits the
// corresponding guidance lines so the model is never taught to use tools
// it cannot call.
func WithSkillExecToolsDisabled() SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsRepositoryResolver sets an invocation-aware repository resolver.
func WithSkillsRepositoryResolver(
	resolver func(*agent.Invocation) skill.Repository,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithMaxLoadedSkills caps how many skills remain "loaded" in session
// state.
//
// When max <= 0, no cap is applied (default behavior).
//
// When max > 0, the processor keeps at most max most-recently touched
// skills and offloads the rest by clearing their state keys.
func WithMaxLoadedSkills(max int) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsDirectoryHints exposes skill directory locators in the skills
// overview and in loaded skill materialization.
func WithSkillsDirectoryHints(
	enable bool,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// WithSkillsFilePathHints exposes SKILL.md file locators in the skills
// overview and in loaded skill materialization.
func WithSkillsFilePathHints(
	enable bool,
) SkillsRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(SkillsRequestProcessorOption)
}

// SkillsRequestProcessor injects skill overviews and loaded contents.
//
// Behavior:
//   - Overview: injects names + descriptions (cheap).
//   - Loaded skills: inject full SKILL.md body.
//   - Docs: inject doc texts selected via state keys.
//
// State keys used (per agent, ephemeral):
//   - skill.LoadedKey(agentName, skillName) -> "1"
//   - skill.DocsKey(agentName, skillName) ->
//     "*" or JSON array of file names.
type SkillsRequestProcessor struct {
	repo               skill.Repository
	repoResolver       func(*agent.Invocation) skill.Repository
	capabilityGuidance *string
	protocolGuidance   *string
	toolingGuidance    *string
	loadMode           string
	toolResultMode     bool
	maxLoadedSkills    int
	toolFlags          skillprofile.Flags
	toolFlagsResolver  func(*agent.Invocation) skillprofile.Flags
	directoryHints     bool
	filePathHints      bool
}

const (
	skillsTurnInitStateKey = "processor:skills:turn_init"
)

// NewSkillsRequestProcessor creates a processor instance.
func NewSkillsRequestProcessor(
	repo skill.Repository,
	opts ...SkillsRequestProcessorOption,
) *SkillsRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

func normalizeSkillLoadMode(mode string) string { _ = "STUB: not implemented"; return "" }

// ProcessRequest implements flow.RequestProcessor.
func (p *SkillsRequestProcessor) ProcessRequest(
	ctx context.Context, inv *agent.Invocation, req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// 1) Always inject overview (names + descriptions) into system
//    message. Merge into existing system message if present.

// Loaded skill bodies/docs are materialized into tool results by a
// post-content request processor.

// 2) Loaded skills full content (merge into existing system message).
// stable prompt order

// Docs

// Summary line to make selected docs explicit.

// Send a preprocessing trace event even when only overview is
// injected, for consistent trace semantics.

func (p *SkillsRequestProcessor) repositoryForInvocation(
	inv *agent.Invocation,
) skill.Repository {
	_ = "STUB: not implemented"
	return *new(skill.Repository)
}

func (p *SkillsRequestProcessor) maybeCapLoadedSkills(
	ctx context.Context,
	inv *agent.Invocation,
	loaded []string,
	ch chan<- *event.Event,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func keepMostRecentSkills(
	inv *agent.Invocation,
	loaded []string,
	max int,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func loadedSkillSet(loaded []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func loadedSkillOrder(
	inv *agent.Invocation,
	loaded []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func loadedSkillOrderFromState(
	inv *agent.Invocation,
	loadedSet map[string]struct{},
) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillsToOrderFromEvents(
	order []string,
	events []event.Event,
	agentName string,
	loadedSet map[string]struct{},
) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillsToOrderFromToolResponseEvent(
	ev event.Event,
	agentName string,
	loadedSet map[string]struct{},
	order []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func fillLoadedSkillOrderAlphabetically(
	order []string,
	loadedSet map[string]struct{},
) []string {
	_ = "STUB: not implemented"
	return nil
}

func skillNameFromToolResponse(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func (p *SkillsRequestProcessor) maybeClearSkillStateForTurn(
	ctx context.Context,
	inv *agent.Invocation,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func clearSkillState(inv *agent.Invocation) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func (p *SkillsRequestProcessor) maybeOffloadLoadedSkills(
	ctx context.Context,
	inv *agent.Invocation,
	loaded []string,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (p *SkillsRequestProcessor) injectOverview(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
	repo skill.Repository,
) {
	_ = "STUB: not implemented"
	return
}

// No system message yet: create one at the front.

func (p *SkillsRequestProcessor) toolFlagsForInvocation(
	inv *agent.Invocation,
) skillprofile.Flags {
	_ = "STUB: not implemented"
	return *new(skillprofile.Flags)
}

func (p *SkillsRequestProcessor) protocolGuidanceText(
	flags skillprofile.Flags,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *SkillsRequestProcessor) toolingGuidanceText(
	flags skillprofile.Flags,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *SkillsRequestProcessor) capabilityGuidanceText(
	flags skillprofile.Flags,
) string {
	_ = "STUB: not implemented"
	return ""
}

func defaultToolingAndWorkspaceGuidance(flags skillprofile.Flags) string {
	_ = "STUB: not implemented"
	return ""
}

func defaultCatalogOnlyGuidance() string { _ = "STUB: not implemented"; return "" }

func defaultKnowledgeOnlyGuidance(flags skillprofile.Flags) string {
	_ = "STUB: not implemented"
	return ""
}

func defaultDocHelpersOnlyGuidance(flags skillprofile.Flags) string {
	_ = "STUB: not implemented"
	return ""
}

func defaultFullToolingAndWorkspaceGuidance(flags skillprofile.Flags) string {
	_ = "STUB: not implemented"
	return ""
}

type skillRootAlias struct {
	alias string
	root  string
}

func buildSkillRootsText(repo skill.Repository) string { _ = "STUB: not implemented"; return "" }

func skillRootAliases(repo skill.Repository) []skillRootAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *SkillsRequestProcessor) skillOverviewSuffix(
	ctx context.Context,
	repo skill.Repository,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func skillDirectoryLocator(
	ctx context.Context,
	repo skill.Repository,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func skillFileLocator(
	ctx context.Context,
	repo skill.Repository,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func skillDirectoryText(
	ctx context.Context,
	repo skill.Repository,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func skillFileText(
	ctx context.Context,
	repo skill.Repository,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// canonicalPathForRel expands symlinks in a path so filepath.Rel agrees across
// aliases such as /var vs /private/var on macOS.
func canonicalPathForRel(p string) string { _ = "STUB: not implemented"; return "" }

func relativeSkillPath(root string, path string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (p *SkillsRequestProcessor) appendSkillPathHints(
	b *strings.Builder,
	ctx context.Context,
	repo skill.Repository,
	name string,
) {
	_ = "STUB: not implemented"
	return
}

func appendKnowledgeGuidance(
	b *strings.Builder,
	flags skillprofile.Flags,
) {
	_ = "STUB: not implemented"
	return
}

func normalizeGuidance(guidance string) string { _ = "STUB: not implemented"; return "" }

func (p *SkillsRequestProcessor) getLoadedSkills(
	inv *agent.Invocation,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *SkillsRequestProcessor) getDocsSelection(
	ctx context.Context,
	inv *agent.Invocation,
	name string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

// Select all doc files present.

func (p *SkillsRequestProcessor) buildDocsText(
	sk *skill.Skill, wanted []string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// Build a map for quick lookup of requested docs.

// Separate docs with a marker title.

// mergeIntoSystem appends content into the existing system message when
// available; otherwise, it creates a new system message at the front.
func (p *SkillsRequestProcessor) mergeIntoSystem(
	req *model.Request, content string,
) {
	_ = "STUB: not implemented"
	return
}
