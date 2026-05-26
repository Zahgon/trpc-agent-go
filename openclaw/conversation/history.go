//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversation

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// HistoryOptions controls speaker-aware history projection.
type HistoryOptions struct {
	AddSessionSummary bool
	MaxHistoryRuns    int
	LabelOverrides    map[string]string
}

// TurnOptions controls speaker-aware turn projection.
type TurnOptions struct {
	Limit          int
	IncludeSystem  bool
	LabelOverrides map[string]string
}

// Turn is one speaker-aware conversation turn.
type Turn struct {
	Role      string    `json:"role,omitempty"`
	Speaker   string    `json:"speaker,omitempty"`
	ActorID   string    `json:"actor_id,omitempty"`
	QuoteText string    `json:"quote_text,omitempty"`
	Text      string    `json:"text,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// BuildInjectedContextMessages projects visible conversation history from
// persisted session events.
func BuildInjectedContextMessages(
	sess *session.Session,
	opts HistoryOptions,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// BuildTurns projects visible conversation turns from persisted session
// events.
func BuildTurns(
	sess *session.Session,
	opts TurnOptions,
) []Turn {
	_ = "STUB: not implemented"
	return nil
}

// FormatTurns renders projected turns as plain text.
func FormatTurns(turns []Turn) string { _ = "STUB: not implemented"; return "" }

// BuildSummaryText renders conversation events as plain text for summary.
func BuildSummaryText(events []event.Event) string { _ = "STUB: not implemented"; return "" }

func buildVisibleHistory(
	events []event.Event,
	since time.Time,
	labelOverrides map[string]string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func buildTurns(
	events []event.Event,
	opts TurnOptions,
) []Turn {
	_ = "STUB: not implemented"
	return nil
}

func buildSummaryLines(events []event.Event) []string { _ = "STUB: not implemented"; return nil }

func includeEvent(evt event.Event, since time.Time) bool { _ = "STUB: not implemented"; return false }

func turnsFromEvent(
	evt event.Event,
	includeSystem bool,
	labelOverrides map[string]string,
) []Turn {
	_ = "STUB: not implemented"
	return nil
}

func visibleMessagesFromEvent(
	evt event.Event,
	labelOverrides map[string]string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func visibleUserMessages(
	evt event.Event,
	labelOverrides map[string]string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// ProjectEventMessage projects one event-derived user message into the
// model-facing request view while keeping persisted session events
// structured.
func ProjectEventMessage(
	inv *agent.Invocation,
	evt event.Event,
	msg model.Message,
) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func visibleAssistantMessages(evt event.Event) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func visibleSystemMessages(evt event.Event) []model.Message { _ = "STUB: not implemented"; return nil }

func userTurnsFromEvent(
	evt event.Event,
	labelOverrides map[string]string,
) []Turn {
	_ = "STUB: not implemented"
	return nil
}

func assistantTurnsFromEvent(evt event.Event) []Turn { _ = "STUB: not implemented"; return nil }

func systemTurnsFromEvent(evt event.Event) []Turn { _ = "STUB: not implemented"; return nil }

func summaryLinesFromEvent(
	evt event.Event,
	labelOverrides map[string]string,
) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func renderUserMessage(
	msg model.Message,
	annotation Annotation,
	labelOverrides map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func projectedUserContentText(
	msg model.Message,
	annotation Annotation,
	labelOverrides map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func projectionMetadata(
	inv *agent.Invocation,
	evt event.Event,
) (Annotation, map[string]string, bool) {
	_ = "STUB: not implemented"
	return *new(Annotation), nil, false
}

func runtimeState(inv *agent.Invocation) map[string]any { _ = "STUB: not implemented"; return nil }

func hasProjectionMetadata(annotation Annotation) bool { _ = "STUB: not implemented"; return false }

// isSyntheticProjectionEvent matches the zero-value event.Event{} used by
// ContentRequestProcessor when projecting the current invocation message.
// If event.Event grows new non-zero-default fields, or callers populate any
// field on this synthetic event, update this heuristic accordingly.
func isSyntheticProjectionEvent(evt event.Event) bool { _ = "STUB: not implemented"; return false }

func nonTextContentParts(
	parts []model.ContentPart,
) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func renderAssistantMessage(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func messageText(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func sessionSummary(
	sess *session.Session,
) (string, time.Time, bool) {
	_ = "STUB: not implemented"
	return "", *new(time.Time), false
}

func speakerLabel(
	annotation Annotation,
	labelOverrides map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func formatTurn(turn Turn) string { _ = "STUB: not implemented"; return "" }

func formatSummary(summaryText string) string { _ = "STUB: not implemented"; return "" }
