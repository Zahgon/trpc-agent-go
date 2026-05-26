//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package processor provides content processing logic for agent requests.
// It includes utilities for including, filtering, and rearranging session
// events for LLM requests, as well as helpers for function call/response
// event handling.
package processor

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// SessionSummaryInjectionMode controls how the session summary is injected
// into the model request.
type SessionSummaryInjectionMode string

const (
	// SessionSummaryInjectionSystem injects the summary as a system message
	// (default). The summary is merged into the existing system message or
	// prepended as a new one. This makes the summary part of the preserved
	// head in token tailoring and is not subject to sliding-window trimming.
	SessionSummaryInjectionSystem SessionSummaryInjectionMode = "system"

	// SessionSummaryInjectionUser injects the summary as a user message
	// placed near session history. The processor prefers merging it into the
	// first user history/current message when possible; if none exists and
	// the existing prompt prefix already ends with a user message, it falls
	// back to merging there to avoid introducing an extra adjacent user
	// block. This mode allows the summary to participate in token-budget
	// trimming like any other user-anchored round, enabling a true
	// sliding-window experience.
	SessionSummaryInjectionUser SessionSummaryInjectionMode = "user"
)

// Content inclusion options.
const (
	// BranchFilterModePrefix Prefix matching pattern
	BranchFilterModePrefix = "prefix"
	// BranchFilterModeSubtree includes only events whose FilterKey is the
	// same as the current filter key or is a descendant of it.
	//
	// Unlike BranchFilterModePrefix, it does not include ancestor FilterKeys.
	// This is useful for isolating history across independent scopes
	// (e.g., permission/tenant views) within the same session.
	BranchFilterModeSubtree = "subtree"
	// BranchFilterModeAll include all
	BranchFilterModeAll = "all"
	// BranchFilterModeExact exact match
	BranchFilterModeExact = "exact"

	// TimelineFilterAll includes all historical message records
	// Suitable for scenarios requiring full conversation context
	TimelineFilterAll = "all"
	// TimelineFilterCurrentRequest only includes messages within the current request cycle
	// Filters out previous historical records, keeping only messages related to this request
	TimelineFilterCurrentRequest = "request"
	// TimelineFilterCurrentInvocation only includes messages within the current invocation session
	// Suitable for scenarios requiring isolation between different invocation cycles in long-running sessions
	TimelineFilterCurrentInvocation = "invocation"
)

// Reasoning content mode constants control how reasoning_content is handled in
// multi-turn conversations. This is particularly important for models like
// DeepSeek that output reasoning_content (thinking chain) alongside the final
// content.
const (
	// ReasoningContentModeKeepAll keeps all reasoning_content in history.
	// Use this for debugging or when you need to retain thinking chains.
	ReasoningContentModeKeepAll = "keep_all"

	// ReasoningContentModeDiscardPreviousTurns discards reasoning_content from
	// ordinary previous request turns. Messages within the current request retain
	// their reasoning_content, and previous requests that performed tool calls
	// also retain it because DeepSeek thinking mode requires tool-call reasoning
	// to be replayed in later turns.
	// Reference: https://api-docs.deepseek.com/guides/thinking_mode#tool-calls
	ReasoningContentModeDiscardPreviousTurns = "discard_previous_turns"

	// ReasoningContentModeDiscardAll discards all reasoning_content from history.
	// Use this for maximum bandwidth savings when reasoning history is not needed.
	ReasoningContentModeDiscardAll = "discard_all"
)

// ContentRequestProcessor implements content processing logic for agent requests.
type ContentRequestProcessor struct {
	// BranchFilterMode determines how to include content from session events.
	// Options: "prefix", "all", "exact" (default: "prefix").
	BranchFilterMode string
	// AddContextPrefix controls whether to add "For context:" prefix when converting foreign events.
	// When false, foreign agent events are passed directly without the prefix.
	AddContextPrefix bool
	// AddSessionSummary controls whether to prepend the current branch summary
	// to the request if available.
	AddSessionSummary bool
	// SessionSummaryInjectionMode controls how the session summary is injected
	// into the model request. Default is SessionSummaryInjectionSystem.
	SessionSummaryInjectionMode SessionSummaryInjectionMode
	// MaxHistoryRuns sets the maximum number of history messages when AddSessionSummary is false.
	// When 0 (default), no limit is applied.
	MaxHistoryRuns int
	// PreserveSameBranch keeps events authored within the same invocation branch in
	// their original roles instead of re-labeling them as user context. This
	// allows graph executions to retain authentic assistant/tool transcripts
	// while still enabling cross-agent contextualization when branches differ.
	PreserveSameBranch bool
	// PreserveForeignMessages keeps events authored by other agents in their
	// original roles and order instead of converting them into user-context
	// messages. This is opt-in because some handoff flows rely on the default
	// foreign-event contextualization behavior.
	PreserveForeignMessages bool
	// TimelineFilterMode controls whether to append history messages to the request.
	TimelineFilterMode string
	// ReasoningContentMode controls how reasoning_content is handled in multi-turn
	// conversations. Default is ReasoningContentModeDiscardPreviousTurns, which
	// keeps reasoning needed for tool-call replay while dropping ordinary older
	// reasoning history.
	ReasoningContentMode string
	// PreloadMemory controls framework-side memory preload.
	// When > 0, it acts as an adaptive preload budget:
	//   - If total memories <= N, preload all memories.
	//   - If total memories > N, preload top-N search results.
	//   - If query extraction is empty, the search fails, or the search
	//     returns no matches, fall back to loading up to N memories
	//     directly.
	// When 0 (default), no memories are preloaded (use tools instead).
	// When < 0, all memories are loaded.
	PreloadMemory int
	// PreloadSessionRecall sets the number of recalled
	// session events to inject into the system prompt.
	// When > 0, query-time search runs across other
	// sessions for the current user.
	// When 0 (default), it is disabled.
	PreloadSessionRecall int
	// PreloadSessionRecallMinScore filters low-confidence
	// recall hits before injection.
	PreloadSessionRecallMinScore float64
	// PreloadSessionRecallSearchMode controls the
	// retrieval mode used for query-time session recall.
	// Default is hybrid when unset.
	PreloadSessionRecallSearchMode session.SearchMode
	// SummaryFormatter allows custom formatting of session summary content.
	// When nil (default), uses the default formatSummaryContent function.
	SummaryFormatter func(summary string) string
	// EventMessageProjector rewrites one event-derived message before it
	// is appended to the model request.
	EventMessageProjector EventMessageProjector
	// ContextCompactionConfig controls request-side historical tool-result
	// compaction before messages are sent to the model.
	ContextCompactionConfig ContextCompactionConfig
	fewShotResolver         func(*agent.Invocation) [][]model.Message
}

type contentRequestRuntimeConfig struct {
	includeMode string
}

// EventMessageProjector projects one event-derived message into the
// model-facing request view.
type EventMessageProjector func(
	inv *agent.Invocation,
	evt event.Event,
	msg model.Message,
) model.Message

// ContentOption is a functional option for configuring the ContentRequestProcessor.
type ContentOption func(*ContentRequestProcessor)

// WithBranchFilterMode sets how to include content from session events.
func WithBranchFilterMode(mode string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithTimelineFilterMode sets whether to append history messages to the request.
func WithTimelineFilterMode(mode string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithAddContextPrefix controls whether to add "For context:" prefix when converting foreign events.
func WithAddContextPrefix(addPrefix bool) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithAddSessionSummary controls whether to prepend the current branch summary
// as a system message when available.
func WithAddSessionSummary(add bool) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithSessionSummaryInjectionMode sets the injection mode for session summaries.
//
// Available modes:
//   - SessionSummaryInjectionSystem (default): injects as system message,
//     merged into existing system message or prepended.
//   - SessionSummaryInjectionUser: injects as a user message near history.
//     The processor first tries to merge it into the first user
//     history/current message; if none exists and the existing prompt prefix
//     already ends with user, it falls back to merging there.
func WithSessionSummaryInjectionMode(mode SessionSummaryInjectionMode) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithMaxHistoryRuns sets the maximum number of history messages when AddSessionSummary is false.
// When 0 (default), no limit is applied.
func WithMaxHistoryRuns(maxRuns int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreserveSameBranch toggles preserving original roles for events emitted
// from the same invocation branch. When enabled, messages that originate from
// nodes in the current agent/graph execution keep their assistant/tool roles
// instead of being rewritten as user context.
func WithPreserveSameBranch(preserve bool) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreserveForeignMessages toggles preserving original roles/order for
// events emitted by other agents instead of rewriting them into user context.
func WithPreserveForeignMessages(preserve bool) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithReasoningContentMode sets how reasoning_content is handled in multi-turn
// conversations. This is particularly important for DeepSeek models where
// ordinary previous-turn reasoning_content may be omitted, but tool-call
// reasoning_content must be replayed in later turns.
//
// Available modes:
//   - ReasoningContentModeDiscardPreviousTurns: Discard reasoning_content from
//     ordinary previous requests, keep for the current request and previous
//     requests that performed tool calls (default, recommended).
//   - ReasoningContentModeKeepAll: Keep all reasoning_content.
//   - ReasoningContentModeDiscardAll: Discard all reasoning_content from history.
func WithReasoningContentMode(mode string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreloadMemory sets the framework-side memory preload behavior.
//   - Set to 0 (default) to disable preloading (use tools instead).
//   - Set to N (N > 0) to use adaptive preload with budget N.
//     Small memory sets are preloaded in full. Larger sets use search and
//     fall back to loading up to N memories directly when search cannot
//     provide usable results.
//   - Set to -1 to load all memories.
//     WARNING: Loading all memories may significantly increase token usage
//     and API costs, especially for users with many stored memories.
//     Consider using a positive budget (e.g., 10-50) for production use.
func WithPreloadMemory(limit int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreloadSessionRecall sets the number of recalled
// session events to preload into the system prompt.
func WithPreloadSessionRecall(limit int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreloadSessionRecallMinScore sets the minimum
// search score required for recalled session events.
func WithPreloadSessionRecallMinScore(minScore float64) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithPreloadSessionRecallSearchMode sets the retrieval
// mode used for query-time session recall preload.
// Default is session.SearchModeHybrid.
func WithPreloadSessionRecallSearchMode(
	mode session.SearchMode,
) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithSummaryFormatter sets a custom formatter for session summary content.
func WithSummaryFormatter(formatter func(summary string) string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithEventMessageProjector sets a projector that rewrites one
// event-derived message before it is appended to the request.
func WithEventMessageProjector(
	projector EventMessageProjector,
) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithEnableContextCompaction toggles prompt-side context compaction during
// history projection. Historical oversized tool results can be compacted
// regardless of whether AddSessionSummary is enabled.
func WithEnableContextCompaction(enable bool) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionKeepRecentRequests preserves the latest N completed
// requests in full when context compaction is enabled.
func WithContextCompactionKeepRecentRequests(n int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionToolResultMaxTokens sets the token threshold above which
// historical tool results are replaced with a placeholder.
func WithContextCompactionToolResultMaxTokens(tokens int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionOversizedToolResultMaxTokens sets the token threshold
// above which any tool result (including from the current request) is truncated
// using head+tail preservation. Like Pass 1, this requires
// EnableContextCompaction=true to take effect, so EnableContextCompaction=false
// guarantees the framework will not modify tool results.
func WithContextCompactionOversizedToolResultMaxTokens(tokens int) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionTokenCounter sets the token counter used by context
// compaction for request thresholds and tool-result budgets.
func WithContextCompactionTokenCounter(counter model.TokenCounter) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionSkipRecentFunc sets the function that determines how
// many tail events are protected from historical tool-result compaction.
func WithContextCompactionSkipRecentFunc(
	skipFunc ContextCompactionSkipRecentFunc,
) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionForceCleanToolNames sets tool names whose results should
// always be compacted to a placeholder while context compaction is enabled.
func WithContextCompactionForceCleanToolNames(names ...string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

// WithContextCompactionKeepToolNames sets tool names whose results should be
// left untouched by context compaction.
func WithContextCompactionKeepToolNames(names ...string) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

func toolNameSet(names []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

// WithFewShotResolver sets an invocation-aware few-shot resolver.
func WithFewShotResolver(
	resolver func(*agent.Invocation) [][]model.Message,
) ContentOption {
	_ = "STUB: not implemented"
	return *new(ContentOption)
}

const (
	mergedUserSeparator = "\n\n"
	contextPrefix       = "For context:"

	contentHasSessionSummaryStateKey = "processor:content:has_session_summary"
	// contentHasCompactedToolResultsStateKey indicates that current-turn tool
	// results were compacted to preserve the active ReAct loop after the session
	// summary absorbed earlier invocation history. Historical request compaction
	// must not set this flag, because downstream processors use it as a
	// same-turn signal.
	contentHasCompactedToolResultsStateKey = "processor:content:has_compacted_tool_results"
	compactedToolResultPlaceholder         = "Tool result omitted from raw history; details are captured in the session summary above."
)

const (
	attachedFilesAnnotationPrefix = "Attached files"
	attachedFileNameFallbackFmt   = "upload_%d"
	attachedFilesMaxPreview       = 20
	hostRefPrefix                 = "host://"
	ignoredAttachmentMimeType     = "application/octet-stream"
)

// NewContentRequestProcessor creates a new content request processor.
func NewContentRequestProcessor(opts ...ContentOption) *ContentRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Default only to include
// filtered contents.
// Default to add context prefix.
// Default to rewriting same-branch lineage events to user context so
// that downstream subagents see a single consolidated user message
// stream unless explicitly opted back into preserving roles.

// Default to append history message.

// Default to disable memory preloading (use tools instead).

// Pass 2 is opt-in: callers must explicitly set a positive value
// AND enable context compaction. Defaulting to 0 keeps the
// processor from silently rewriting tool results.

// Apply options.

// ProcessRequest implements the flow.RequestProcessor interface.
// It handles adding messages from the session events to the request.
func (p *ContentRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Append per-filter messages from session events when allowed.

// Send a preprocessing event.

func (p *ContentRequestProcessor) injectFewShotMessages(
	invocation *agent.Invocation,
	req *model.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (p *ContentRequestProcessor) runtimeConfigFromInvocation(
	invocation *agent.Invocation,
) contentRequestRuntimeConfig {
	_ = "STUB: not implemented"
	return *new(contentRequestRuntimeConfig)
}

func (p *ContentRequestProcessor) appendSessionMessages(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	skipHistory bool,
	includeInvocationMessage bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip session summary when include_contents=none, but still get current
// invocation's events (tool calls/results) to maintain ReAct loop context.

// Fetch session summary early so we can insert it after other
// semi-stable system blocks (for example, preloaded memories).

// Preload memories into system prompt if configured.
// PreloadMemory: 0 = disabled, -1 = all, N > 0 = adaptive preload budget.

// User-mode injection is deferred until after history messages are
// collected, so the summary can be merged with the first user
// message in history when applicable.

// Default system-mode: inject as system context message.

// When include_contents=none, only get events from current invocation
// to preserve tool call history within the current ReAct loop.
// This fixes the infinite loop issue where the agent doesn't see its
// own tool calls when running as an isolated subgraph.

// When user-mode summary injection is active, prepend the summary as a
// user message near history. Prefer merging into the first user
// history/current message so the summary stays attached to the live user
// turn. If no such message exists, fall back to a trailing user message in
// req.Messages (for example, injected context) to avoid creating an extra
// adjacent user block.

// injectSystemContextMessage injects summary or memory context into request.
// It merges the content into an existing system message if one exists,
// or prepends as a new system message if none exists.
func (p *ContentRequestProcessor) injectSystemContextMessage(
	req *model.Request,
	msg model.Message,
) {
	_ = "STUB: not implemented"
	return
}

// injectInjectedContextMessages inserts per-run context messages into the request
// before session-derived history is appended.
func (p *ContentRequestProcessor) injectInjectedContextMessages(invocation *agent.Invocation, req *model.Request) {
	_ = "STUB: not implemented"
	return
}

// getSessionSummaryText returns the raw session summary text and its
// UpdatedAt timestamp for the current branch. It does not format or assign
// a role — callers decide how to inject the text into the request.
func (p *ContentRequestProcessor) getSessionSummaryText(inv *agent.Invocation) (string, time.Time) {
	_ = "STUB: not implemented"
	return "", *new(time.Time)
}

// Acquire read lock to protect Summaries access.

// For BranchFilterModeAll, prefer the full-session summary under empty filter key.

// Try exact match first.

// For BranchFilterModePrefix, aggregate summaries with matching prefix.

// getSessionSummaryMessage returns the current-branch session summary as a
// system message if available and non-empty, along with its UpdatedAt timestamp.
func (p *ContentRequestProcessor) getSessionSummaryMessage(inv *agent.Invocation) (*model.Message, time.Time) {
	_ = "STUB: not implemented"
	return nil, *new(time.Time)
}

// prependSummaryUserMessage prepends the session summary as a user message
// before history messages. It checks three merge opportunities in order:
//  1. If history/current contains a user message, merge the summary into the
//     first available one so the summary stays attached to the live user turn.
//  2. If no such history/current user message exists and reqPrefix
//     (req.Messages before history) ends with a user message, merge the
//     summary into that trailing prefix message to avoid an extra adjacent
//     user block.
//  3. Otherwise prepend as an independent user message.
//
// When merging into reqPrefix, the function mutates reqPrefix in place and
// returns messages unchanged. The caller appends messages to req.Messages
// after this call.
func (p *ContentRequestProcessor) prependSummaryUserMessage(
	summaryText string,
	messages []model.Message,
	reqPrefix []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// Case 1: merge into the first available user history/current message.

// Case 2: reqPrefix (existing req.Messages) ends with a user message.
// Merge summary into that message only as a fallback when there is no
// user history/current message to attach the summary to.

// Case 3: prepend as independent user message.

// formatSummaryForUser returns a user-role-friendly summary text.
// It uses the custom SummaryFormatter if set, otherwise applies a neutral
// default suitable for user-channel injection.
func (p *ContentRequestProcessor) formatSummaryForUser(summary string) string {
	_ = "STUB: not implemented"
	return ""
}

// aggregatePrefixSummaries aggregates all summaries whose keys have the given prefix.
func (p *ContentRequestProcessor) aggregatePrefixSummaries(
	summaries map[string]*session.Summary,
	prefix string,
) (string, time.Time) {
	_ = "STUB: not implemented"
	return "", *new(time.Time)
}

// Check if key matches prefix (key starts with "prefix/" or key equals prefix).

// formatSummary applies custom formatter if available, otherwise uses default.
func (p *ContentRequestProcessor) formatSummary(summary string) string {
	_ = "STUB: not implemented"
	return ""
}

// Default format.

// getHistoryMessages gets history messages for the current filter, potentially truncated by MaxHistoryRuns.
// This method is used when AddSessionSummary is false to get recent history messages.
func (p *ContentRequestProcessor) getIncrementMessages(inv *agent.Invocation, since time.Time) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// use error fill message content if message content is empty

// insert invocation message

// Apply compaction to the already timeline-filtered projection. Tool-result
// policy (force-clean/keep) and historical passes must run for scoped modes
// such as request/invocation, not only when TimelineFilterAll is selected.

// Get current request ID for reasoning content filtering.

// Convert events to messages with reasoning content handling.

// Convert foreign events or keep as-is.

// Apply reasoning content stripping based on mode.

// Apply MaxHistoryRuns limit when AddSessionSummary is false.

// compactCurrentInvocationEvent preserves the minimum structured state needed
// for same-turn tool loops after a summary has already absorbed earlier
// invocation history. Assistant tool-call messages are kept intact, while tool
// results are replaced with a small placeholder that points the model at the
// summary for details.
func (p *ContentRequestProcessor) compactCurrentInvocationEvent(
	evt event.Event,
	inv *agent.Invocation,
	filter string,
	isZeroTime bool,
	since time.Time,
) (event.Event, bool) {
	_ = "STUB: not implemented"
	return *new(event.Event), false
}

func compactedCurrentInvocationMessage(
	msg model.Message,
	cfg ContextCompactionConfig,
) (model.Message, bool) {
	_ = "STUB: not implemented"
	return *new(model.Message), false
}

func annotateUserMessagesWithAttachedFiles(
	messages []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func annotateUserMessageWithAttachedFiles(
	msg model.Message,
) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func hasAttachedFilesAnnotation(parts []model.ContentPart) bool {
	_ = "STUB: not implemented"
	return false
}

func buildAttachedFilesAnnotationText(
	parts []model.ContentPart,
) string {
	_ = "STUB: not implemented"
	return ""
}

func fileNamesForAnnotation(
	parts []model.ContentPart,
) ([]string, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func fileLabelForAnnotation(file *model.File, count int) string {
	_ = "STUB: not implemented"
	return ""
}

func fileMimeLabel(file *model.File) string { _ = "STUB: not implemented"; return "" }

func fileNameFromAnnotationRef(fileID string) string { _ = "STUB: not implemented"; return "" }

func annotationRefDisplay(fileID string) string { _ = "STUB: not implemented"; return "" }

func baseNameForAnnotation(raw string) string { _ = "STUB: not implemented"; return "" }

func fileNameFromArtifactRef(fileID string) string { _ = "STUB: not implemented"; return "" }

// applyMaxHistoryRuns trims messages to at most maxRuns entries from the tail.
// If the trim boundary falls on a tool-result message whose corresponding
// tool_use was truncated, the boundary is advanced past any such orphaned
// results to prevent API 400 "unexpected tool_use_id" errors.
func applyMaxHistoryRuns(messages []model.Message, maxRuns int) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// Only scan the truncated prefix when the boundary actually falls on a
// tool-result message; otherwise there's nothing to skip.

// Collect tool-call IDs that will be truncated (before startIdx).

// Skip orphaned tool results whose corresponding call was truncated.

// processReasoningContent applies reasoning content stripping based on the
// configured mode and request boundaries.
func (p *ContentRequestProcessor) processReasoningContent(
	msg model.Message,
	messageRequestID string,
	currentRequestID string,
	requestHasToolCalls bool,
) model.Message {
	_ = "STUB: not implemented"
	// Only process assistant messages with reasoning content.
	return *new(model.Message)
}

// Discard all reasoning_content.

// Keep all reasoning_content: do nothing.

// ReasoningContentModeDiscardPreviousTurns or empty (default):
// Discard reasoning_content from ordinary previous requests.
// Current request messages and requests with tool calls retain their
// reasoning_content for provider replay requirements.

func (p *ContentRequestProcessor) projectEventMessage(
	inv *agent.Invocation,
	evt event.Event,
	msg model.Message,
) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

// getCurrentInvocationMessages gets messages only from the current invocation.
// This is used when include_contents=none to preserve tool call history within
// the current ReAct loop while isolating from parent/other branch history.
func (p *ContentRequestProcessor) getCurrentInvocationMessages(inv *agent.Invocation) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (p *ContentRequestProcessor) collectCurrentInvocationEvents(
	inv *agent.Invocation,
) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

func isCurrentInvocationEligibleEvent(
	evt event.Event,
	invocationID string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeCurrentInvocationEvent(evt event.Event) event.Event {
	_ = "STUB: not implemented"
	return *new(event.Event)
}

func containsInvocationMessage(
	events []event.Event,
	invocationMessage model.Message,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ContentRequestProcessor) projectCurrentInvocationMessages(
	inv *agent.Invocation,
	events []event.Event,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (p *ContentRequestProcessor) projectMessagesForEvent(
	inv *agent.Invocation,
	evt event.Event,
	currentRequestID string,
	toolCallRequestIDs map[string]struct{},
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func requestIDsWithToolCalls(events []event.Event) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func requestHasToolCalls(requestIDs map[string]struct{}, requestID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ContentRequestProcessor) truncateOversizedToolResultMessages(
	messages []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (p *ContentRequestProcessor) insertInvocationMessage(
	events []event.Event, inv *agent.Invocation) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (p *ContentRequestProcessor) mergeUserMessages(
	messages []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// shouldIncludeEvent decides whether an event should be included in the model
// request and whether that event should be treated as the invocation message.
//
// The second return value (isInvocationMessage) is intentionally strict: only
// exact invocation-message matches return true. Mid-turn user-message
// protection may still include an event, but returns false for this flag to
// avoid conflating inclusion with strict message equality.
func (p *ContentRequestProcessor) shouldIncludeEvent(evt event.Event, inv *agent.Invocation, filter string,
	isZeroTime bool, since time.Time) (bool, bool) {
	_ = "STUB: not implemented"
	// Fast reject malformed, partial, or empty-content events.
	return false, false
}

// Exact invocation message match keeps existing semantics.

// Keep the current invocation user message even when summary UpdatedAt
// would otherwise exclude it. This preserves the original request while
// still allowing same-turn tool/assistant history already covered by the
// summary to be compacted out of the next prompt.

// Use strict After so events stamped exactly at summary UpdatedAt are
// treated as already summarized and not re-sent.

// isEventEligibleForInclusion checks basic event validity before expensive
// filtering logic runs.
func isEventEligibleForInclusion(evt event.Event) bool { _ = "STUB: not implemented"; return false }

// isStrictInvocationMessage checks whether the event exactly matches the
// current invocation message, including content equality semantics.
func isStrictInvocationMessage(evt event.Event, inv *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

// isCurrentInvocationUserMessage keeps the current invocation's user message
// even when summary UpdatedAt would exclude it by timestamp.
//
// RequestID + InvocationID matching avoids preserving unrelated user messages
// from other invocations that may share the same request scope.
func isCurrentInvocationUserMessage(evt event.Event, inv *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

// hasCompactedCurrentInvocationToolResults reports whether same-invocation tool
// result events before the active summary cutoff are actually compacted out of
// the raw prompt history.
func (p *ContentRequestProcessor) hasCompactedCurrentInvocationToolResults(
	inv *agent.Invocation,
	since time.Time,
) bool {
	_ = "STUB: not implemented"
	return false
}

func eventHasCompactedCurrentInvocationToolResult(
	evt event.Event,
	cfg ContextCompactionConfig,
) bool {
	_ = "STUB: not implemented"
	return false
}

// passTimelineFilter applies request/invocation timeline constraints.
func (p *ContentRequestProcessor) passTimelineFilter(evt event.Event, inv *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

// passBranchFilter applies branch-scoping constraints.
func (p *ContentRequestProcessor) passBranchFilter(evt event.Event, filter string) bool {
	_ = "STUB: not implemented"
	return false
}

func filterSubtree(eventFilterKey, filterKey string) bool { _ = "STUB: not implemented"; return false }

func invocationMessageEqual(invMsg model.Message, evtMsg model.Message) bool {
	_ = "STUB: not implemented"
	return false
}

// isOtherAgentReply checks whether the event is a reply from another agent.
func (p *ContentRequestProcessor) isOtherAgentReply(
	currentAgentName string,
	currentBranch string,
	evt *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Treat events within the same branch lineage as non-foreign to
// preserve original roles. This includes both descendants and
// ancestors of the current branch.

// convertForeignEvent converts an event authored by another agent as a user-content event.
func (p *ContentRequestProcessor) convertForeignEvent(evt *event.Event) event.Event {
	_ = "STUB: not implemented"
	return *new(event.Event)
}

// Create a new event with user context.

// Build content parts for context.

// When prefix is disabled, pass the content directly.

// When prefix is disabled, pass tool call info directly.

// When prefix is disabled, pass tool result directly.

// Set the converted message.

// rearrangeEventsForLatestFunctionResponse rearranges the events for the latest function_response.
func (p *ContentRequestProcessor) rearrangeLatestFuncResp(
	events []event.Event,
) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Check if latest event is a function response.

// Look for corresponding function call event.

// Collect function response events between call and latest response.

// Build result with rearranged events.

// rearrangeEventsForAsyncFunctionResponsesInHistory rearranges the async function_response events in the history.
func (p *ContentRequestProcessor) rearrangeAsyncFuncRespHist(
	events []event.Event,
) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Create a local copy to avoid implicit memory aliasing.
// This bug is fixed in go 1.22.
// See: https://tip.golang.org/doc/go1.22#language

// Function response should be handled with function call below.

// Merge multiple async function responses.

type pendingToolCallRound struct {
	eventIndex int
	pendingIDs map[string]struct{}
}

type matchedToolResponseEvent struct {
	eventIndex    int
	choiceIndices []int
}

// toolResponseMatchesByCallEvent matches tool-result choices to the nearest
// preceding tool-call round that is still waiting for the result ID.
func toolResponseMatchesByCallEvent(events []event.Event) map[int][]matchedToolResponseEvent {
	_ = "STUB: not implemented"
	return nil
}

// appendToolResponseChoice records one matching choice while coalescing choices
// from the same response event into one match.
func appendToolResponseChoice(
	matches []matchedToolResponseEvent,
	eventIndex int,
	choiceIndex int,
) []matchedToolResponseEvent {
	_ = "STUB: not implemented"
	return nil
}

// filterToolResponseEvent clones a matched response event with only the tool
// result choices that belong to the current tool-call round.
func filterToolResponseEvent(events []event.Event, match matchedToolResponseEvent) event.Event {
	_ = "STUB: not implemented"
	return *new(event.Event)
}

// mergeFunctionResponseEvents merges a list of function_response events into one event.
func (p *ContentRequestProcessor) mergeFunctionResponseEvents(
	functionResponseEvents []event.Event,
) event.Event {
	_ = "STUB: not implemented"
	return *new(event.Event)
}

// Start with the first event as base.

// Collect all tool response messages, preserving each individual ToolID.

func toMap(ids []string) map[string]bool { _ = "STUB: not implemented"; return nil }

// toStringSet converts IDs to a set for membership checks.
func toStringSet(ids []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

// getPreloadMemoryMessage returns preloaded memories as a system message if available.
func (p *ContentRequestProcessor) getPreloadMemoryMessage(
	ctx context.Context,
	inv *agent.Invocation,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

// Validate user key.

// Handle PreloadMemory: 0 = disabled, -1 = all, N > 0 = adaptive budget.

// getAdaptivePreloadMemoryMessage preloads all memories for small memory sets
// and falls back to query-aware search for larger sets.
func (p *ContentRequestProcessor) getAdaptivePreloadMemoryMessage(
	ctx context.Context,
	inv *agent.Invocation,
	userKey memory.UserKey,
	budget int,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

// loadPreloadMemoryMessage loads memories directly and formats them as a
// system message.
func (p *ContentRequestProcessor) loadPreloadMemoryMessage(
	ctx context.Context,
	inv *agent.Invocation,
	userKey memory.UserKey,
	limit int,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

func newPreloadMemoryMessage(memories []*memory.Entry) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

// buildPreloadSearchQuery extracts the current user text used for adaptive
// preload search.
func buildPreloadSearchQuery(msg model.Message) string { _ = "STUB: not implemented"; return "" }

// formatMemoryContent formats memories for system prompt injection.
func formatMemoryContent(memories []*memory.Entry) string { _ = "STUB: not implemented"; return "" }

// Append metadata inline for richer context.

// Do not render topic labels in the preload prompt. The memory_add
// tool expects topics as []string, and showing inline
// "topics=foo, bar" text can lead models to copy a scalar value into
// tool arguments.

func (p *ContentRequestProcessor) getPreloadSessionRecallMessage(
	ctx context.Context,
	inv *agent.Invocation,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

func extractSearchQueryText(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func formatSessionRecallContent(
	results []session.EventSearchResult,
) string {
	_ = "STUB: not implemented"
	return ""
}
