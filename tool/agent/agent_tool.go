//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package agent provides agent tool implementations for the agent system.
package agent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Tool wraps an agent as a tool that can be called within a larger application.
// The agent's input schema is used to define the tool's input parameters, and
// the agent's output is returned as the tool's result.
type Tool struct {
	agent                  agent.Agent
	skipSummarization      bool
	streamInner            bool
	innerTextMode          InnerTextMode
	structuredStreamErrors bool
	historyScope           HistoryScope
	responseMode           ResponseMode
	name                   string
	description            string
	inputSchema            *tool.Schema
	outputSchema           *tool.Schema
}

// Option is a function that configures an AgentTool.
type Option func(*agentToolOptions)

// agentToolOptions holds the configuration options for AgentTool.
type agentToolOptions struct {
	skipSummarization      bool
	streamInner            bool
	innerTextMode          InnerTextMode
	structuredStreamErrors bool
	historyScope           HistoryScope
	responseMode           ResponseMode
	description            *string
}

// InnerTextMode controls whether forwarded inner assistant text is visible
// in the parent flow when StreamInner is enabled.
type InnerTextMode = tool.InnerTextMode

const (
	// InnerTextModeDefault preserves the default behavior.
	InnerTextModeDefault = tool.InnerTextModeDefault

	// InnerTextModeInclude forwards inner assistant text to the parent flow.
	InnerTextModeInclude = tool.InnerTextModeInclude

	// InnerTextModeExclude suppresses forwarded inner assistant text while
	// still aggregating that text into the final tool response.
	InnerTextModeExclude = tool.InnerTextModeExclude
)

// ResponseMode controls which child assistant text AgentTool returns as the tool
// result. It does not change session event mirroring or inner streaming
// behavior.
type ResponseMode int

const (
	// ResponseModeDefault preserves the legacy assistant-content
	// concatenation behavior.
	ResponseModeDefault ResponseMode = iota

	// ResponseModeFinalOnly returns only the last complete assistant message
	// emitted by the child agent. If none is emitted, the tool result is
	// an empty string.
	ResponseModeFinalOnly
)

// WithResponseMode sets how AgentTool builds the tool result from child agent
// events.
func WithResponseMode(mode ResponseMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipSummarization sets whether to skip summarization of the agent output.
func WithSkipSummarization(skip bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamInner controls whether the AgentTool should forward inner agent
// streaming events up to the parent flow. When false, the flow will treat the
// tool as callable-only (no inner streaming in the parent transcript).
func WithStreamInner(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInnerTextMode controls whether forwarded inner assistant text is
// visible in the parent flow when StreamInner is enabled.
func WithInnerTextMode(mode InnerTextMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStructuredStreamErrors controls whether AgentTool opts into structured
// error chunks when it is executed through the framework as a streamable tool.
func WithStructuredStreamErrors(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDescription sets the description exposed by the agent tool declaration.
func WithDescription(description string) Option { _ = "STUB: not implemented"; return *new(Option) }

// HistoryScope controls whether and how AgentTool inherits parent history.
//   - HistoryScopeIsolated: keep child events isolated; do not inherit parent history.
//   - HistoryScopeParentBranch: inherit parent branch history by using a hierarchical
//     filter key "parent/child-uuid" so that content processors see parent events via
//     prefix matching while keeping child events in a separate sub-branch.
type HistoryScope int

// HistoryScopeIsolated: keep child events isolated; do not inherit parent history.
// HistoryScopeParentBranch: inherit parent branch history by using a hierarchical
// filter key "parent/child-uuid" so that content processors see parent events via
// prefix matching while keeping child events in a separate sub-branch.
const (
	HistoryScopeIsolated HistoryScope = iota
	HistoryScopeParentBranch
)

// WithHistoryScope sets the history inheritance behavior for AgentTool.
func WithHistoryScope(scope HistoryScope) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewTool creates a new Tool that wraps the given agent.
//
// Note: The tool name is derived from the agent's info (agent.Info().Name).
// The agent name must comply with LLM API requirements for compatibility.
// Some APIs (e.g., Kimi, DeepSeek) enforce strict naming patterns:
// - Must match pattern: ^[a-zA-Z0-9_-]+$
// - Cannot contain Chinese characters, parentheses, or special symbols
//
// Best practice: Use ^[a-zA-Z0-9_-]+ only to ensure maximum compatibility.
func NewTool(agent agent.Agent, opts ...Option) *Tool {
	_ = "STUB: not implemented"
	// Default to allowing summarization so the parent agent can perform its
	// normal post-tool reasoning unless opt-out is requested.
	return nil
}

// Use the agent's input schema if available, otherwise fall back to default.

// Convert the agent's input schema to tool.Schema format.

// Generate default input schema for the agent tool.

// Call executes the agent tool with the provided JSON arguments.
func (at *Tool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Prefer to reuse parent invocation + session so the child can see parent
// history according to the configured history scope.

// Fallback: isolated in-memory run when parent invocation is not available.

// callWithParentInvocation executes the agent using parent invocation context.
// This allows the child agent to inherit parent history based on the configured
// history scope.
func (at *Tool) callWithParentInvocation(
	ctx context.Context,
	parentInv *agent.Invocation,
	message model.Message,
) (string, error) {
	_ = "STUB: not implemented"
	// If the parent invocation does not have a session, fall back to isolated mode.
	return "", nil
}

// Flush all events emitted before this tool call so that the snapshot sees all events.

// Build child filter key based on history scope.

// Run the agent and collect response.

func (at *Tool) surfaceRootNodeIDForParentInvocation(
	parentInv *agent.Invocation,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (at *Tool) childInvocationOptions(
	parentInv *agent.Invocation,
	message model.Message,
	childKey string,
) []agent.InvocationOptions {
	_ = "STUB: not implemented"
	return nil
}

// wrapWithCompletion consumes events, notifies completion when required, and forwards to a new channel.
func (at *Tool) wrapWithCompletion(ctx context.Context, inv *agent.Invocation, src <-chan *event.Event) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// wrapWithCallSemantics consumes events from a child agent invocation that is
// executed without a Runner. It mirrors persisted events into the shared
// Session so multi-step tool calling can work, and notifies completion when
// required.
func (at *Tool) wrapWithCallSemantics(
	ctx context.Context,
	inv *agent.Invocation,
	src <-chan *event.Event,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func ensureInvocationEventFields(inv *agent.Invocation, evt *event.Event) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) updatePendingVisibleCompletionForSession(
	ctx context.Context,
	inv *agent.Invocation,
	pending *event.Event,
	evt *event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (at *Tool) appendPendingVisibleCompletionState(
	ctx context.Context,
	inv *agent.Invocation,
	pending *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) replacePendingVisibleCompletionForSession(
	ctx context.Context,
	inv *agent.Invocation,
	pending *event.Event,
	next *event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (at *Tool) wrapWithStreamSemantics(
	ctx context.Context,
	inv *agent.Invocation,
	src <-chan *event.Event,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func shouldDeferStreamCompletion(
	ctx context.Context,
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (at *Tool) ensureUserMessageForCall(
	ctx context.Context,
	inv *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) appendEvent(
	ctx context.Context,
	inv *agent.Invocation,
	evt *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func sessionHasEventID(inv *agent.Invocation, eventID string) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldMirrorEventToSession(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func persistableSessionEvent(evt *event.Event) *event.Event { _ = "STUB: not implemented"; return nil }

func shouldDelayVisibleCompletionSessionMirror(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func visibleCompletionStateOnlySessionEvent(evt *event.Event) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func shouldSuppressGraphExecutorBarrierEvent(
	inv *agent.Invocation,
	evt *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isGraphCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func isGraphCompletionSnapshotEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func assistantMessageContent(evt *event.Event) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type pendingFinalResultChunk struct {
	Result     any
	StateDelta map[string][]byte
}

func graphCompletionFinalChunk(evt *event.Event) (pendingFinalResultChunk, bool) {
	_ = "STUB: not implemented"
	return *new(pendingFinalResultChunk), false
}

func completionResponseIDFromStateDelta(delta map[string][]byte) string {
	_ = "STUB: not implemented"
	return ""
}

func cloneStateDelta(delta map[string][]byte) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

// callWithIsolatedRunner executes the agent in an isolated environment using
// an in-memory session service. This is used as a fallback when no parent
// invocation context is available.
func (at *Tool) callWithIsolatedRunner(
	ctx context.Context,
	message model.Message,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// buildChildFilterKey constructs a child filter key based on the history scope
// configuration. For HistoryScopeParentBranch, it creates a hierarchical key
// that allows the child to inherit parent history.
func (at *Tool) buildChildFilterKey(parentInv *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

// collectResponse collects and concatenates assistant messages from the event
// channel, returning the complete response text.
func (at *Tool) collectResponse(inv *agent.Invocation, evCh <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// collectFinalResponse returns the last complete child assistant message. If no
// complete assistant message is emitted, it returns an empty string and nil
// error.
func collectFinalResponse(evCh <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// collectLegacyResponse preserves the pre-#1365 concatenation semantics for the
// default callable path. It applies a narrow fix for issue #1640 by skipping a
// trailing graph-completion snapshot event that re-emits the same non-partial
// assistant content that was already collected. Divergent snapshot content is
// still concatenated to keep default output bytes stable for existing callers;
// broader alignment with the snapshot-aware collector is tracked as a separate
// semantic-change request.
func collectLegacyResponse(evCh <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func shouldRewriteCallableCompletion(inv *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeResponseMode(mode ResponseMode) ResponseMode {
	_ = "STUB: not implemented"
	return *new(ResponseMode)
}

// StreamableCall executes the agent tool with streaming support and returns a stream reader.
// It runs the wrapped agent and forwards its streaming text output as chunks.
// The returned chunks' Content are plain strings representing incremental text.
func (at *Tool) StreamableCall(ctx context.Context, jsonArgs []byte) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type streamCompletionState struct {
	pendingCompletionChunk     *pendingFinalResultChunk
	pendingVisibleCompletion   *event.Event
	pendingStreamVisibleEvent  *event.Event
	sawGraphCompletionSnapshot bool
	lastAssistantResponseID    string
	lastAssistantContent       string
	finalOnlyResult            string
	overrideResult             string
}

func (at *Tool) streamableCallContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (at *Tool) runStreamableCall(
	ctx context.Context,
	jsonArgs []byte,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) streamFromParentInvocation(
	ctx context.Context,
	parentInv *agent.Invocation,
	message model.Message,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) forwardSubInvocationStream(
	ctx context.Context,
	subInv *agent.Invocation,
	wrapped <-chan *event.Event,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) updateFinalOnlyStreamResult(
	ev *event.Event,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) capturePendingVisibleCompletion(
	ctx context.Context,
	inv *agent.Invocation,
	ev *event.Event,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) flushPendingVisibleCompletionForSession(
	ctx context.Context,
	inv *agent.Invocation,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) completeSuppressedBarrierEvent(
	ctx context.Context,
	inv *agent.Invocation,
	evt *event.Event,
	pending **event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) completeSuppressedBarrierStreamEvent(
	ctx context.Context,
	inv *agent.Invocation,
	evt *event.Event,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) capturePendingCompletionChunk(
	ev *event.Event,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) updateStreamCompletionState(
	ev *event.Event,
	state *streamCompletionState,
) {
	_ = "STUB: not implemented"
	return
}

func visibleCompletionSessionEvent(evt *event.Event, author string) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func visibleCompletionStreamEvent(
	evt *event.Event,
	author string,
) (*event.Event, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (at *Tool) emitPendingVisibleCompletionEvent(
	state *streamCompletionState,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) emitPendingCompletionChunk(
	state *streamCompletionState,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) emitFinalOnlyResultChunk(
	state *streamCompletionState,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func (at *Tool) streamFromFallbackRunner(
	ctx context.Context,
	message model.Message,
	writer *tool.StreamWriter,
) {
	_ = "STUB: not implemented"
	return
}

func sendStreamableCallError(
	ctx context.Context,
	writer *tool.StreamWriter,
	format string,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func streamableCallErrorEvent(ctx context.Context, err error) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (at *Tool) fallbackRunnerRunOptions(ctx context.Context) []agent.RunOption {
	_ = "STUB: not implemented"
	return nil
}

// SkipSummarization exposes whether the AgentTool prefers skipping
// outer-agent summarization after its tool.response.
func (at *Tool) SkipSummarization() bool { _ = "STUB: not implemented"; return false }

// StructuredStreamErrors reports that AgentTool expects structured error chunks.
func (at *Tool) StructuredStreamErrors() bool { _ = "STUB: not implemented"; return false }

// TRPCAgentGoStructuredStreamErrorsOptIn provides an explicit framework hook
// for structured stream error semantics.
func (at *Tool) TRPCAgentGoStructuredStreamErrorsOptIn() bool {
	_ = "STUB: not implemented"
	return false
}

// StreamInner exposes whether this AgentTool prefers the flow to treat it as
// streamable (forwarding inner deltas) versus callable-only.
func (at *Tool) StreamInner() bool { _ = "STUB: not implemented"; return false }

// InnerTextMode exposes how forwarded inner assistant text should be handled
// when StreamInner is enabled.
func (at *Tool) InnerTextMode() InnerTextMode {
	_ = "STUB: not implemented"
	return *new(InnerTextMode)
}

// Declaration returns the tool's declaration information.
//
// Note: The tool name must comply with LLM API requirements.
// Some APIs (e.g., Kimi, DeepSeek) enforce strict naming patterns:
// - Must match pattern: ^[a-zA-Z0-9_-]+$
// - Cannot contain Chinese characters, parentheses, or special symbols
//
// Best practice: Use ^[a-zA-Z0-9_-]+ only to ensure maximum compatibility.
func (at *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// convertMapToToolSchema converts a map[string]any schema to tool.Schema format.
// This function handles the conversion from the agent's input schema format to the tool schema format.
func convertMapToToolSchema(schema map[string]any) *tool.Schema {
	_ = "STUB: not implemented"
	return nil
}
