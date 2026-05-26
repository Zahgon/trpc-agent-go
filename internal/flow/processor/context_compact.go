//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// DefaultContextCompactionKeepRecentRequests preserves the latest N
	// completed requests in full when request-side context compaction is enabled.
	DefaultContextCompactionKeepRecentRequests = 1
	// DefaultContextCompactionToolResultMaxTokens is the default token
	// threshold above which historical tool results are replaced with a
	// placeholder.
	DefaultContextCompactionToolResultMaxTokens = 1024

	// DefaultContextCompactionOversizedToolResultMaxTokens is the recommended
	// token threshold above which ANY tool result (including current request)
	// is truncated to head+tail when Pass 2 is opted into.
	//
	// NOTE: this constant is only the suggested value to pass to
	// WithContextCompactionOversizedToolResultMaxTokens; it is NOT applied
	// automatically. Pass 2 only runs when both EnableContextCompaction is
	// true and the threshold is greater than 0. The default for the option
	// itself is 0 (disabled) so that EnableContextCompaction=false truly
	// means "framework will not modify tool results".
	DefaultContextCompactionOversizedToolResultMaxTokens = 8192

	historicalToolResultPlaceholder = "Historical tool result omitted to save context."
	policyToolResultPlaceholder     = "Tool result omitted by context compaction policy."
)

// ContextCompactionConfig controls request-side history compaction applied
// while projecting session events into a model request.
type ContextCompactionConfig struct {
	Enabled             bool
	KeepRecentRequests  int
	ToolResultMaxTokens int
	// OversizedToolResultMaxTokens is the token threshold above which any tool
	// result (including current-request results) is truncated using head+tail
	// preservation. Like Pass 1, this also requires Enabled=true; it will not
	// fire when context compaction is turned off, even if a positive threshold
	// is configured. 0 disables it regardless of Enabled.
	OversizedToolResultMaxTokens int
	// TokenCounter estimates request and tool-result size for compaction decisions.
	// When nil, SimpleTokenCounter is used.
	TokenCounter model.TokenCounter
	// SkipRecentFunc returns how many tail events should be treated as recent
	// and protected from historical tool-result compaction.
	SkipRecentFunc ContextCompactionSkipRecentFunc

	toolResultCompactionRules toolResultCompactionRules
}

// ContextCompactionSkipRecentFunc determines how many recent events should be
// protected from historical tool-result compaction.
type ContextCompactionSkipRecentFunc func(events []event.Event) int

// ContextCompactionStats reports how much prompt history was compacted during
// request projection.
type ContextCompactionStats struct {
	ToolResultsCompacted int
	EstimatedTokensSaved int
}

type toolResultCompactionRules struct {
	forceCleanToolNames map[string]struct{}
	keepToolNames       map[string]struct{}
}

func normalizeContextCompactionConfig(
	cfg ContextCompactionConfig,
) ContextCompactionConfig {
	_ = "STUB: not implemented"
	return *new(ContextCompactionConfig)
}

func normalizeToolNameSet(in map[string]struct{}) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func compactIncrementEvents(
	ctx context.Context,
	events []event.Event,
	currentRequestID string,
	currentInvocationID string,
	cfg ContextCompactionConfig,
) ([]event.Event, ContextCompactionStats) {
	_ = "STUB: not implemented"
	return nil, *new(ContextCompactionStats)
}

// Pass 0: named tool results → full placeholder replacement.
// This is explicit user policy and applies before recency protection.

// Pass 1: historical tool results → full placeholder replacement.
// Gated on Enabled (requires context compaction to be on).

// Pass 2: oversized tool results (including current request) → head+tail
// truncation. Gated on EnableContextCompaction together with Pass 1, so
// the framework never silently rewrites tool results when context
// compaction is disabled.

func (cfg ContextCompactionConfig) hasForceCleanToolResults() bool {
	_ = "STUB: not implemented"
	return false
}

func (cfg ContextCompactionConfig) keepToolResult(msg model.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func (cfg ContextCompactionConfig) forceCleanToolResult(msg model.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func applyForceCleanToolResultPass(
	ctx context.Context,
	events []event.Event,
	cfg ContextCompactionConfig,
) ([]event.Event, ContextCompactionStats) {
	_ = "STUB: not implemented"
	return nil, *new(ContextCompactionStats)
}

func applyHistoricalToolResultPass(
	ctx context.Context,
	events []event.Event,
	currentKey string,
	keepRecentRequests int,
	maxTokens int,
	cfg ContextCompactionConfig,
) ([]event.Event, ContextCompactionStats) {
	_ = "STUB: not implemented"
	return nil, *new(ContextCompactionStats)
}

func compactHistoricalToolResultEvent(
	ctx context.Context,
	evt event.Event,
	protectedRequestIDs map[string]struct{},
	maxTokens int,
	cfg ContextCompactionConfig,
) (event.Event, bool, int, int) {
	_ = "STUB: not implemented"
	return *new(event.Event), false, 0, 0
}

func applyOversizedToolResultPass(
	ctx context.Context,
	events []event.Event,
	maxTokens int,
	cfg ContextCompactionConfig,
) ([]event.Event, ContextCompactionStats) {
	_ = "STUB: not implemented"
	return nil, *new(ContextCompactionStats)
}

func rewriteToolResultEventMessages(
	ctx context.Context,
	evt event.Event,
	maxTokens int,
	rewrite func(context.Context, model.Message, int) (model.Message, bool, int),
) (event.Event, bool, int, int) {
	_ = "STUB: not implemented"
	return *new(event.Event), false, 0, 0
}

func mergeContextCompactionStats(
	base ContextCompactionStats,
	delta ContextCompactionStats,
) ContextCompactionStats {
	_ = "STUB: not implemented"
	return *new(ContextCompactionStats)
}

func collectProtectedRequestIDs(
	events []event.Event,
	currentKey string,
	keepRecentRequests int,
	skipRecentFunc ContextCompactionSkipRecentFunc,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func protectRecentEvents(
	protected map[string]struct{},
	events []event.Event,
	skipRecentFunc ContextCompactionSkipRecentFunc,
) {
	_ = "STUB: not implemented"
	return
}

func collectCompletedCompactionUnitKeys(events []event.Event) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func compactionUnitKey(requestID, invocationID string) string { _ = "STUB: not implemented"; return "" }

// truncateOversizedToolResultMessage applies head+tail truncation to any tool
// result whose estimated token count exceeds maxTokens. Unlike the historical
// placeholder compaction, this preserves the beginning and end of the content
// so the model can still see key information. Inspired by Codex's
// truncate_middle_chars and Claude Code's per-tool maxResultSizeChars.
//
// TODO: text ContentParts are preserved as-is; truncating individual text parts
// inside ContentParts is deferred until multimodal tool results are common.
func truncateOversizedToolResultMessage(
	ctx context.Context,
	msg model.Message,
	maxTokens int,
) (model.Message, bool, int) {
	_ = "STUB: not implemented"
	return *new(model.Message), false, 0
}

func truncateOversizedToolResultMessageWithCounter(
	ctx context.Context,
	msg model.Message,
	maxTokens int,
	counter model.TokenCounter,
) (model.Message, bool, int) {
	_ = "STUB: not implemented"
	return *new(model.Message), false, 0
}

func truncateMiddleToTokenBudget(
	ctx context.Context,
	msg model.Message,
	maxTokens int,
	counter model.TokenCounter,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// truncateMiddle keeps the first half and last half of the content (by
// character count) up to maxChars total, inserting a marker in the middle
// showing how much was removed. This preserves the beginning (usually
// contains key structure/headers) and end (usually contains conclusions)
// of the tool output.
func truncateMiddle(s string, maxChars int) string { _ = "STUB: not implemented"; return "" }

func cleanToolResultMessageWithCounter(
	ctx context.Context,
	msg model.Message,
	counter model.TokenCounter,
) (model.Message, bool, int) {
	_ = "STUB: not implemented"
	return *new(model.Message), false, 0
}

// Force-clean is policy-driven, not threshold-driven. Even if token
// counting fails, still replace the payload with policyToolResultPlaceholder;
// savedTokens falls back to 0 because the exact savings are unknown.

func compactHistoricalToolResultMessage(
	ctx context.Context,
	msg model.Message,
	maxTokens int,
) (model.Message, bool, int) {
	_ = "STUB: not implemented"
	return *new(model.Message), false, 0
}

func compactHistoricalToolResultMessageWithCounter(
	ctx context.Context,
	msg model.Message,
	maxTokens int,
	counter model.TokenCounter,
) (model.Message, bool, int) {
	_ = "STUB: not implemented"
	return *new(model.Message), false, 0
}
