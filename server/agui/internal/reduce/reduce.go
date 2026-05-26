//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package reduce implements the logic to reduce the AG-UI track events into message snapshots.
package reduce

import (
	"strings"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// reducer reduces the AG-UI track events into message snapshots.
type reducer struct {
	appName                   string
	userID                    string
	includeRunLifecycleEvents bool
	texts                     map[string]*textState
	reasonings                map[string]*reasoningState
	lastReasoningChunkID      string
	toolCalls                 map[string]*toolCallState
	messages                  []*aguievents.Message
}

// textPhase is the phase of the text message.
type textPhase int

const (
	textReceiving textPhase = iota
	textEnded
)

// textState is the state of the text message.
type textState struct {
	role    string
	name    string
	content strings.Builder
	phase   textPhase
	index   int
}

type reasoningPhase int

const (
	reasoningReceiving reasoningPhase = iota
	reasoningEnded
)

type reasoningState struct {
	role    string
	name    string
	content strings.Builder
	phase   reasoningPhase
	index   int
	started bool
}

// toolPhase is the phase of the tool call.
type toolPhase int

const (
	toolAwaitingArgs toolPhase = iota
	toolAwaitingResult
	toolCompleted
)

// toolCallState is the state of the tool call.
type toolCallState struct {
	messageID string
	name      string
	content   strings.Builder
	phase     toolPhase
	index     int
}

// Reduce reduces the AG-UI track events into message snapshots.
// In order to fetch the history messages as much as possible, still return the messages even if there is an error.
func Reduce(appName, userID string, events []session.TrackEvent, opt ...Option) ([]aguievents.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In order to fetch the history messages as much as possible, still return the messages even if there is an error.

// new creates a new reducer.
func new(appName, userID string, opts options) *reducer { _ = "STUB: not implemented"; return nil }

// reduce reduces the AG-UI track event into a message snapshot.
func (r *reducer) reduce(trackEvent session.TrackEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) reduceEvent(evt aguievents.Event) error { _ = "STUB: not implemented"; return nil }

func (r *reducer) handleUserMessageCustomEvent(e *aguievents.CustomEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) finalizePartial() { _ = "STUB: not implemented"; return }

func (r *reducer) sanitizeSnapshotMessage(message *aguievents.Message) *aguievents.Message {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleRunStarted(e *aguievents.RunStartedEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleRunFinished(e *aguievents.RunFinishedEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleRunError(e *aguievents.RunErrorEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) appendRunActivity(e aguievents.Event, content map[string]any) {
	_ = "STUB: not implemented"
	return
}

// handleTextStart handles the text message start event.
func (r *reducer) handleTextStart(e *aguievents.TextMessageStartEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleTextContent handles the text message content event.
func (r *reducer) handleTextContent(e *aguievents.TextMessageContentEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleTextEnd handles the text message end event.
func (r *reducer) handleTextEnd(e *aguievents.TextMessageEndEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleTextChunk handles the text message chunk event.
func (r *reducer) handleTextChunk(e *aguievents.TextMessageChunkEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleReasoningMessageStart(e *aguievents.ReasoningMessageStartEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleReasoningContent(e *aguievents.ReasoningMessageContentEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleReasoningEnd(e *aguievents.ReasoningMessageEndEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleReasoningChunk(e *aguievents.ReasoningMessageChunkEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reducer) handleReasoningEncryptedValue(e *aguievents.ReasoningEncryptedValueEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleToolStart handles the tool call start event.
func (r *reducer) handleToolStart(e *aguievents.ToolCallStartEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleToolArgs handles the tool call arguments event.
func (r *reducer) handleToolArgs(e *aguievents.ToolCallArgsEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleToolEnd handles the tool call end event.
func (r *reducer) handleToolEnd(e *aguievents.ToolCallEndEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleToolResult handles the tool call result event.
func (r *reducer) handleToolResult(e *aguievents.ToolCallResultEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// handleActivity handles the activity event.
func (r *reducer) handleActivity(e aguievents.Event) error { _ = "STUB: not implemented"; return nil }
