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
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	// ViolationTagPrefix prefixes the tag carrying the matched word.
	// Tag format: "violation=<word>", separated with event.TagDelimiter when multiple tags exist.
	ViolationTagPrefix = "violation="
)

// ProhibitedWords is the list of prohibited words to filter.
var ProhibitedWords = []string{
	"pirated serial number",
	"crack password",
}

// MarkViolationHook checks events for prohibited words and marks them.
func MarkViolationHook() session.AppendEventHook {
	_ = "STUB: not implemented"
	return *new(session.AppendEventHook)
}

// FilterViolationHook filters out events containing prohibited content on GetSession.
// This prevents violated Q&A pairs from being sent to LLM.
func FilterViolationHook() session.GetSessionHook {
	_ = "STUB: not implemented"
	return *new(session.GetSessionHook)
}

// containsProhibitedWord checks if content contains any prohibited word.
// Returns the matched word or empty string.
func containsProhibitedWord(content string) string { _ = "STUB: not implemented"; return "" }

// filterViolationEvents removes events marked as violation and their paired Q/A.
// If a user message is violated, skip it and the following assistant response.
// If an assistant response is violated, skip it and the preceding user message.
func filterViolationEvents(sess *session.Session) int { _ = "STUB: not implemented"; return 0 }

// First pass: mark indices to skip

// If user message is violated, also skip the next assistant response

// If assistant response is violated, also skip the preceding user message

// Second pass: build filtered list

func getEventContent(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func truncate(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }

func parseViolationTag(tag string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func appendTags(existing string, tags ...string) string { _ = "STUB: not implemented"; return "" }

// Supported consecutive user message strategies.
const (
	strategyMerge       = "merge"
	strategyPlaceholder = "placeholder"
	strategySkip        = "skip"
)

// validConsecutiveStrategies returns the list of valid strategy names.
func validConsecutiveStrategies() []string { _ = "STUB: not implemented"; return nil }

// isValidConsecutiveStrategy checks if the given strategy is valid.
func isValidConsecutiveStrategy(s string) bool { _ = "STUB: not implemented"; return false }

// FixConsecutiveUserMessagesHook fixes consecutive user messages in session history.
// This is a GetSessionHook that runs when session is retrieved, before sending to LLM.
//
// Supported strategies:
//   - "merge": Merge consecutive user messages into one.
//   - "placeholder": Insert placeholder assistant responses between consecutive user messages.
//   - "skip": Keep only the last user message, skip earlier ones.
//
// Using GetSessionHook is simpler than AppendEventHook because:
//  1. No need to access sessionService (no persistence needed, just fix in-memory).
//  2. No recursion concerns.
//  3. Fixes happen at read time, keeping storage unchanged.
func FixConsecutiveUserMessagesHook(strategy string) session.GetSessionHook {
	_ = "STUB: not implemented"
	return *new(session.GetSessionHook)
}

// fixConsecutiveUserMessages modifies session events to fix consecutive user messages.
func fixConsecutiveUserMessages(sess *session.Session, strategy string) {
	_ = "STUB: not implemented"
	return
}

// mergeConsecutiveUserMessages merges consecutive user messages into one.
func mergeConsecutiveUserMessages(events []event.Event) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// If current is user message and previous in result is also user message, merge.

// Merge tags from current event to prevent losing violation tags.

// insertPlaceholdersBetweenUserMessages inserts placeholder assistant responses.
func insertPlaceholdersBetweenUserMessages(events []event.Event) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// If current is user message and previous in result is also user message,
// insert placeholder before current.

// skipEarlierConsecutiveUserMessages keeps only the last user message in consecutive sequence.
func skipEarlierConsecutiveUserMessages(events []event.Event) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// If current is user message and previous in result is also user message,
// replace previous with current (keep the later one).
