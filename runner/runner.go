//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package runner provides the core runner functionality.
package runner

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/internal/state/flush"
	"trpc.group/trpc-go/trpc-agent-go/internal/state/steer"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// Author types for events.
const (
	authorUser = "user"

	errMsgEmptyRequestID               = "runner: empty request id"
	errMsgNilCancelFunc                = "runner: nil cancel function"
	interruptedAssistantFinishReason   = "cancelled"
	interruptedAssistantExtensionKey   = "trpc_agent.runner.interrupted_assistant"
	cancelledSessionPersistenceTimeout = time.Second
)

var (
	// ErrRunNotFound indicates that the request ID is not active anymore.
	ErrRunNotFound = errors.New("runner: request id not running")
	// ErrQueuedUserMessageUnsupported indicates that the runner does not
	// support safe-boundary user-message enqueue.
	ErrQueuedUserMessageUnsupported = errors.New(
		"runner: queued user messages unsupported",
	)
	// ErrInvalidQueuedUserMessage indicates that the message cannot be
	// inserted as a user steer message.
	ErrInvalidQueuedUserMessage = errors.New(
		"runner: queued message must be a non-empty user message",
	)
)

// Option is a function that configures a Runner.
type Option func(*Options)

// WithSessionService sets the session service to use.
func WithSessionService(service session.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// AgentFactory creates an agent for a single run.
//
// This enables request-scoped agent construction (for example, building the
// agent with a prompt/model/sandbox that depends on the current request).
type AgentFactory func(
	ctx context.Context,
	ro agent.RunOptions,
) (agent.Agent, error)

// WithMemoryService sets the memory service to use.
func WithMemoryService(service memory.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionIngestor sets the session ingestor that receives completed
// session transcripts for ingestion into an external long-term memory
// platform.
//
// The name is intentionally scoped to "Session" because today the contract
// only operates on completed session transcripts. Keeping the option name
// specific leaves room for additional ingestor flavours (e.g. event-level
// or user-level) without overloading a single option.
func WithSessionIngestor(ingestor session.Ingestor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithArtifactService sets the artifact service to use.
func WithArtifactService(service artifact.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAgent adds an agent to the runner registry for name-based lookup.
func WithAgent(name string, ag agent.Agent) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAgentFactory registers an agent factory for name-based lookup.
//
// When the runner resolves an agent name and no registered agent instance
// exists for that name, it will fall back to this factory and create a new
// agent for the current run.
func WithAgentFactory(name string, factory AgentFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPlugins registers plugins on the runner.
func WithPlugins(plugins ...plugin.Plugin) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAwaitUserReplyRouting enables one-shot next-user-turn routing from
// session state produced by agent.MarkAwaitingUserReply or the
// await_user_reply framework tool.
//
// When enabled, Runner checks session state before each user turn. If the
// previous turn persisted an await-user-reply route, Runner consumes that
// route once and starts the new run from the recorded agent instead of the
// default agent.
//
// Default: false.
func WithAwaitUserReplyRouting(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPersistInterruptedAssistant sets the runner default for whether a
// cancelled streaming run persists already-emitted assistant text as a final
// assistant message.
//
// The default is false to preserve cancellation semantics for callers that
// expect cancelled partial text not to affect later turns. A single run can
// override this default with agent.WithPersistInterruptedAssistant.
func WithPersistInterruptedAssistant(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Runner is the interface for running agents.
type Runner interface {
	Run(
		ctx context.Context,
		userID string,
		sessionID string,
		message model.Message,
		runOpts ...agent.RunOption,
	) (<-chan *event.Event, error)

	// Close closes the runner and releases owned resources.
	// It's safe to call Close multiple times.
	// Only resources created by the runner (not provided by user) will be closed.
	Close() error
}

// ManagedRunner extends Runner with run control APIs.
//
// RequestID is used as the run identifier.
//
// If the caller does not set a request ID via agent.WithRequestID,
// Runner will generate one and inject it into every emitted event.
type ManagedRunner interface {
	Runner

	// Cancel cancels a running invocation by request ID.
	// It returns true if a matching run was found.
	Cancel(requestID string) bool

	// RunStatus returns the current status for a running invocation.
	// It returns false when the request ID is unknown or the run completed.
	RunStatus(requestID string) (RunStatus, bool)
}

// SteerableRunner extends ManagedRunner with safe-boundary user steering.
//
// EnqueueUserMessage never mutates the session immediately. The message is
// queued and then appended by llmflow only after the current tool_call /
// tool_response boundary is complete and before the next model request.
type SteerableRunner interface {
	ManagedRunner

	// EnqueueUserMessage queues a user message for the active request.
	EnqueueUserMessage(requestID string, message model.Message) error
}

// EnqueueUserMessage queues a user message on runners that support steering.
func EnqueueUserMessage(
	r Runner,
	requestID string,
	message model.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RunStatus is a snapshot of a running invocation.
type RunStatus struct {
	RequestID    string
	InvocationID string
	AgentName    string
	SessionKey   session.Key
	StartedAt    time.Time
	LastEventAt  time.Time
	EventCount   int
}

// runner runs agents.
type runner struct {
	appName                            string
	defaultAgentName                   string
	agents                             map[string]agent.Agent
	agentFactories                     map[string]AgentFactory
	sessionService                     session.Service
	memoryService                      memory.Service
	ingestor                           session.Ingestor
	artifactService                    artifact.Service
	pluginManager                      agent.PluginManager
	ralphLoop                          *RalphLoopConfig
	awaitUserReplyRouting              bool
	persistInterruptedAssistantDefault bool

	// Resource management fields.
	ownedSessionService bool      // Indicates if sessionService was created by this runner.
	closeOnce           sync.Once // Ensures Close is called only once.

	runsMu sync.RWMutex
	runs   map[string]*runHandle
}

type runHandle struct {
	cancel context.CancelFunc
	queue  *steer.Queue

	mu     sync.RWMutex
	status RunStatus
}

// Options is the options for the Runner.
type Options struct {
	sessionService                     session.Service
	memoryService                      memory.Service
	ingestor                           session.Ingestor
	artifactService                    artifact.Service
	agents                             map[string]agent.Agent
	agentFactories                     map[string]AgentFactory
	plugins                            []plugin.Plugin
	ralphLoop                          *RalphLoopConfig
	awaitUserReplyRouting              bool
	persistInterruptedAssistantDefault bool
}

// newOptions creates a new Options.
func newOptions(opt ...Option) Options { _ = "STUB: not implemented"; return *new(Options) }

// NewRunner creates a new Runner.
func NewRunner(appName string, ag agent.Agent, opts ...Option) Runner {
	_ = "STUB: not implemented"
	return *new(Runner)
}

// Track if we created the session service.

// Register the default agent for observability defaults.

// Register all runner identities for observability fallback.

// NewRunnerWithAgentFactory creates a Runner whose default agent is created
// on demand for each run.
//
// This is useful when agent configuration depends on the current request
// (prompt, model, sandbox instance, etc.), and you want to avoid
// initializing a heavy agent at service startup.
func NewRunnerWithAgentFactory(
	appName string,
	defaultAgentName string,
	factory AgentFactory,
	opts ...Option,
) Runner {
	_ = "STUB: not implemented"
	return *new(Runner)
}

// Close closes the runner and cleans up owned resources.
// It's safe to call Close multiple times.
// Only resources created by this runner will be closed.
func (r *runner) Close() error { _ = "STUB: not implemented"; return nil }

// Only close resources that we own (created by this runner).

func (r *runner) cancelAllRuns() { _ = "STUB: not implemented"; return }

// Run runs the agent.
func (r *runner) Run(
	ctx context.Context,
	userID string,
	sessionID string,
	message model.Message,
	runOpts ...agent.RunOption,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resolve per-request app name override. When the caller provides an
// AppName via RunOption, it takes precedence over the runner default so
// that a single runner can isolate session/memory data across projects.

// Resolve or create the session for this user and conversation.

// Ensure the invocation can be accessed by downstream components (e.g., tools)
// by embedding it into the context. This is necessary for tools like
// transfer_to_agent that rely on agent.InvocationFromContext(ctx).

// Create flush channel and attach flusher before agent.Run to ensure cloned invocations inherit it.

// Run the agent and get the event channel.

// Attempt to persist the error event so the session reflects the failure.

// Populate content to ensure it is valid for persistence (and viewable by users).

// Process the agent events and emit them to the output channel.

func (r *runner) applyRunnerRunDefaults(ro *agent.RunOptions) { _ = "STUB: not implemented"; return }

// seedSessionHistory persists caller-supplied history messages into an empty
// session so that subsequent turns and tool calls build on the same canonical
// transcript. It is a no-op when no messages are provided or the session
// already contains events.
func (r *runner) seedSessionHistory(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	ag agent.Agent,
	ro agent.RunOptions,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// appendSessionMessages persists messages into the session transcript in the
// provided order.
func (r *runner) appendSessionMessages(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	ag agent.Agent,
	messages []model.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

// appendMessagesAsSessionEvents persists messages into session events in the
// provided order.
func (r *runner) appendMessagesAsSessionEvents(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	ag agent.Agent,
	messages []model.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

// appendIncomingMessage appends the user's incoming message to the session
// when it carries a payload and is not already covered by the seeded history.
func (r *runner) appendIncomingMessage(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	message model.Message,
	ro agent.RunOptions,
	historySeeded bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) Cancel(requestID string) bool { _ = "STUB: not implemented"; return false }

func (r *runner) RunStatus(requestID string) (RunStatus, bool) {
	_ = "STUB: not implemented"
	return *new(RunStatus), false
}

func (r *runner) EnqueueUserMessage(
	requestID string,
	message model.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) newExecutionContext(
	ctx context.Context,
	ro agent.RunOptions,
) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (r *runner) registerRun(
	requestID string,
	status RunStatus,
	cancel context.CancelFunc,
	queue *steer.Queue,
) (*runHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) unregisterRun(requestID string) { _ = "STUB: not implemented"; return }

func (r *runner) lookupRun(requestID string) *runHandle { _ = "STUB: not implemented"; return nil }

func (r *runner) lookupCancel(requestID string) context.CancelFunc {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc)
}

// resolveAgent decides which agent to use for this run.
func (r *runner) selectAgent(
	ctx context.Context,
	ro agent.RunOptions,
) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func (r *runner) selectedRootLookupName(
	ro agent.RunOptions,
	awaitUserReplyRootName string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *runner) wrapSelectedAgent(ag agent.Agent) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// getOrCreateSession returns an existing session or creates a new one.
func (r *runner) getOrCreateSession(
	ctx context.Context, key session.Key,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// eventLoopContext bundles all channels and state required by the event loop.
type eventLoopContext struct {
	sess                               *session.Session
	invocation                         *agent.Invocation
	agentEventCh                       <-chan *event.Event
	flushChan                          chan *flush.FlushRequest
	processedEventCh                   chan *event.Event
	runHandle                          *runHandle
	baselineFinalResponseID            string
	priorAssistantResponseIDs          map[string]struct{}
	finalStateDelta                    map[string][]byte
	finalChoices                       []model.Choice
	fallbackChoices                    []model.Choice
	fallbackResponseID                 string
	fallbackStateDelta                 map[string][]byte
	finalError                         *model.ResponseError
	graphCompletionSeen                bool
	freshAssistantContentProduced      bool
	persistedAssistantResponseIDs      map[string]struct{}
	persistedAssistantChoiceSignatures map[string]struct{}
	emittedAssistantChoiceSignatures   map[string]struct{}
	visibleCompletionResponseIDs       map[string]struct{}
	visibleCompletionChoiceSignatures  map[string]struct{}
	sawTerminalError                   bool
	streamFilter                       graph.StreamModeFilter
	interruptedAssistants              map[string]*interruptedAssistantAccumulator
	interruptedAssistantSequence       int64
	// emittedAssistantResponseIDs tracks response IDs that already produced a
	// non-partial assistant message event during this run.
	//
	// It is used to avoid echoing the same final assistant message again in the
	// runner-completion event when graph final model responses are emitted.
	emittedAssistantResponseIDs map[string]struct{}
}

type interruptedAssistantAccumulator struct {
	sess                               *session.Session
	sequence                           int64
	responseID                         string
	author                             string
	invocationID                       string
	parentInvocationID                 string
	branch                             string
	filterKey                          string
	requestID                          string
	created                            int64
	choiceContent                      map[int]*strings.Builder
	persistedAssistantResponseIDs      map[string]struct{}
	persistedAssistantChoiceSignatures map[string]struct{}
}

// processAgentEvents consumes agent events, persists to session, and emits.
func (r *runner) processAgentEvents(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	agentEventCh <-chan *event.Event,
	flushChan chan *flush.FlushRequest,
	handle *runHandle,
) chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// runEventLoop drives the main event processing loop for a single invocation.
func (r *runner) runEventLoop(ctx context.Context, loop *eventLoopContext) {
	_ = "STUB: not implemented"
	return
}

// Agent event stream completed.

// Disable further flush requests for this invocation.

// Flush channel closed, disable further flush handling.

// Handle the flush request.

// processSingleAgentEvent handles a single agent event.
func (r *runner) processSingleAgentEvent(ctx context.Context, loop *eventLoopContext, agentEvent *event.Event) error {
	_ = "STUB: not implemented"
	return nil

	// Preserve existing behavior: skip nil events without failing the loop.
}

// Capture graph-level completion snapshot for final event.

// Append qualifying events to session and trigger summarization.

// Notify completion if required.

// Emit event to output channel.

func (r *runner) recordPersistedAssistantEvent(
	loop *eventLoopContext,
	agentEvent *event.Event,
	persisted bool,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) recordPersistedInterruptedAssistantSessionEvent(
	loop *eventLoopContext,
	persistSession *session.Session,
	agentEvent *event.Event,
	persisted bool,
) {
	_ = "STUB: not implemented"
	return
}

func recordPersistedAssistantOnAccumulator(
	acc *interruptedAssistantAccumulator,
	agentEvent *event.Event,
	persisted bool,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) recordVisibleCompletionEmission(
	loop *eventLoopContext,
	agentEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) recordRunEvent(loop *eventLoopContext) { _ = "STUB: not implemented"; return }

func (r *runner) applyEventPlugins(
	ctx context.Context,
	invocation *agent.Invocation,
	e *event.Event,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func copyEventInvocationFields(dst *event.Event, src *event.Event) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) markCompletionSnapshotOnly(
	loop *eventLoopContext,
	agentEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) recordEmittedAssistantResponseID(
	loop *eventLoopContext,
	e *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func eventHasAssistantMessageContent(e *event.Event) bool { _ = "STUB: not implemented"; return false }

type interruptedAssistantMetadata struct {
	Reason string `json:"reason,omitempty"`
}

func interruptedAssistantAccumulatorForSession(
	loop *eventLoopContext,
	persistSession *session.Session,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantAccumulatorForEvent(
	loop *eventLoopContext,
	persistSession *session.Session,
	agentEvent *event.Event,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantAccumulatorForLineage(
	loop *eventLoopContext,
	persistSession *session.Session,
	lineageKey string,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func getInterruptedAssistantAccumulatorForEvent(
	loop *eventLoopContext,
	persistSession *session.Session,
	agentEvent *event.Event,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func getInterruptedAssistantAccumulator(
	loop *eventLoopContext,
	persistSession *session.Session,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func defaultInterruptedAssistantAccumulator(
	loop *eventLoopContext,
) *interruptedAssistantAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantSessionKey(sess *session.Session) string {
	_ = "STUB: not implemented"
	return ""
}

func interruptedAssistantAccumulatorKeyPrefix(sess *session.Session) string {
	_ = "STUB: not implemented"
	return ""
}

func interruptedAssistantAccumulatorKey(sess *session.Session, lineageKey string) string {
	_ = "STUB: not implemented"
	return ""
}

func interruptedAssistantLineageKey(
	loop *eventLoopContext,
	agentEvent *event.Event,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *runner) recordInterruptedAssistantDelta(
	loop *eventLoopContext,
	agentEvent *event.Event,
	persistSession *session.Session,
) {
	_ = "STUB: not implemented"
	return
}

func interruptedAssistantHasTextDelta(rsp *model.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func captureInterruptedAssistantEventIdentity(
	acc *interruptedAssistantAccumulator,
	agentEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func fillInterruptedAssistantRequestID(
	acc *interruptedAssistantAccumulator,
	loop *eventLoopContext,
) {
	_ = "STUB: not implemented"
	return
}

func injectInterruptedAssistantEventIdentity(
	inv *agent.Invocation,
	acc *interruptedAssistantAccumulator,
	evt *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// safePersistInterruptedAssistant guards cancellation-time partial persistence
// against panics from session services.
func (r *runner) safePersistInterruptedAssistant(ctx context.Context, loop *eventLoopContext) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) persistInterruptedAssistant(ctx context.Context, loop *eventLoopContext) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) interruptedAssistantEvent(
	ctx context.Context,
	loop *eventLoopContext,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) interruptedAssistantEventForAccumulator(
	ctx context.Context,
	loop *eventLoopContext,
	acc *interruptedAssistantAccumulator,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantChoices(loop *eventLoopContext) []model.Choice {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantChoicesFromAccumulator(
	acc *interruptedAssistantAccumulator,
) []model.Choice {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) interruptedAssistantAlreadyPersisted(
	loop *eventLoopContext,
	choices []model.Choice,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) interruptedAssistantAlreadyPersistedForAccumulator(
	acc *interruptedAssistantAccumulator,
	choices []model.Choice,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldPersistInterruptedAssistant(loop *eventLoopContext) bool {
	_ = "STUB: not implemented"
	return false
}

func contextDoneReason(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func sessionPersistenceContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// safeEmitRunnerCompletion guards emitRunnerCompletion against panics from session services.
func (r *runner) safeEmitRunnerCompletion(ctx context.Context, loop *eventLoopContext) {
	_ = "STUB: not implemented"
	return
}

// handleFlushRequest drains buffered agent events when a flush request arrives and closes the request's ACK channel
// once all events currently buffered in the agent event channel have been processed.
func (r *runner) handleFlushRequest(ctx context.Context, loop *eventLoopContext, req *flush.FlushRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// handleEventPersistence appends qualifying events to the session and triggers
// asynchronous summarization.
func (r *runner) handleEventPersistence(
	ctx context.Context,
	invocation *agent.Invocation,
	sess *session.Session,
	persistSession *session.Session,
	agentEvent *event.Event,
) bool {
	_ = "STUB: not implemented"
	// Ensure error events have content so they are valid for persistence.
	return false
}

// Append event to session if it's complete (not partial).

// Skip user messages, tool call events, and invalid content.
// These should not trigger summarization.

// Trigger summary check after tool results to handle long tool call
// sequences (ReAct loops). The existing ShouldSummarize checker
// (event count / token threshold) decides whether to actually run.
// Also trigger after final assistant text responses as before.
// Skip if the event explicitly opts out of summarization.

// When sync intra-run summary is active for this
// invocation, the flow already summarises between LLM
// iterations. Skip redundant async enqueue for intermediate
// tool-result events but still allow the final assistant
// response to trigger an async job so the session summary
// is up-to-date at turn end.

// Use EnqueueSummaryJob for true asynchronous processing.
// Prefer filter-specific summarization to avoid scanning all filters.

// Do not enqueue full-session summary here. The worker will cascade
// a full-session summarization after a branch update when appropriate.

// Note: Auto memory extraction is triggered once at runner completion,
// not here, to avoid redundant extraction calls.

// shouldPersistEvent determines if an event should be persisted to the session.
// Events are persisted if they contain state deltas or are complete, valid
// responses.
func (r *runner) shouldPersistEvent(agentEvent *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func isGraphCompletionEvent(agentEvent *event.Event) bool { _ = "STUB: not implemented"; return false }

func isGraphCompletionSnapshotEvent(agentEvent *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldSuppressGraphCompletionEvent(
	loop *eventLoopContext,
	agentEvent *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldSuppressGraphExecutorBarrierEvent(
	loop *eventLoopContext,
	agentEvent *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

// captureGraphCompletion captures the final state delta and choices from a
// graph execution completion event.
func (r *runner) captureGraphCompletion(
	agentEvent *event.Event,
) (map[string][]byte, []model.Choice) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) captureCompletionFallback(
	loop *eventLoopContext,
	agentEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// A later visible terminal response supersedes any earlier hidden graph completion snapshot.

// The last non-partial response wins so the completion event reflects
// the terminal outcome seen by the runner.

func (r *runner) captureRoutedCompletionError(
	loop *eventLoopContext,
	agentEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func sameSession(a *session.Session, b *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeCompletionFallbackStateDelta(
	dst map[string][]byte,
	src map[string][]byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func shouldPropagateFallbackStateKey(key string) bool { _ = "STUB: not implemented"; return false }

func mergeStateDelta(
	dst map[string][]byte,
	src map[string][]byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func cloneResponseError(err *model.ResponseError) *model.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func shouldPropagateFallbackState(err *model.ResponseError) bool {
	_ = "STUB: not implemented"
	return false
}

// emitRunnerCompletion creates and emits the final runner completion event,
// optionally propagating graph-level completion data.
func (r *runner) emitRunnerCompletion(ctx context.Context, loop *eventLoopContext) {
	_ = "STUB: not implemented"
	// Resolve per-request app name override for the completion Author.
	return
}

// Create runner completion event.

// Propagate graph-level completion data if available.

// Append runner completion event to session.

// Use a context to deliver runner-completion after cancellation without blocking cleanup indefinitely.

// Enqueue auto memory extraction job if memory service is configured.

// Enqueue external session ingestion if configured.

func graphCompletionSessionStateDelta(stateDelta map[string][]byte) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func shouldDropGraphCompletionSessionStateKey(key string) bool {
	_ = "STUB: not implemented"
	return false
}

func resolveExecutionTraceStatus(loop *eventLoopContext, ctxErr error) trace.TraceStatus {
	_ = "STUB: not implemented"
	return *new(trace.TraceStatus)
}

// propagateGraphCompletion propagates graph-level completion data (state delta
// and final choices) to the runner completion event.
func (r *runner) propagateGraphCompletion(
	runnerCompletionEvent *event.Event,
	finalStateDelta map[string][]byte,
	finalChoices []model.Choice,
	echoFinalChoices bool,
) {
	_ = "STUB: not implemented"
	// Initialize state delta map if needed.
	return
}

// Copy state delta with byte ownership.

// Optionally echo the final text as a non-streaming assistant message
// if graph provided it in its completion.

// Keep only content to avoid carrying tool deltas etc.
// Use JSON marshal/unmarshal to deep-copy minimal fields safely.

func shouldClearRunnerCompletionChoicesInSession(
	loop *eventLoopContext,
	finalChoices []model.Choice,
	finalStateDelta map[string][]byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) completionChoicesForRunner(
	loop *eventLoopContext,
	finalStateDelta map[string][]byte,
) []model.Choice {
	_ = "STUB: not implemented"
	return nil
}

// shouldEchoFinalChoicesInCompletion decides whether Runner should copy the
// graph's final assistant message into the runner-completion event.
//
// Default behavior: Graph Large Language Model (LLM) nodes do not emit the
// final (Done=true) assistant response event. In that mode, Runner must echo
// the final choices so clients can reliably read the final answer.
//
// When GraphEmitFinalModelResponses is enabled, graph LLM nodes emit final
// (Done=true) assistant response events. In that mode, Runner will skip
// echoing final choices if it can match the graph's last response identifier
// (ID) to a response ID that was already emitted, avoiding duplicates.
func (r *runner) shouldEchoFinalChoicesInCompletion(
	loop *eventLoopContext,
	finalChoices []model.Choice,
	finalStateDelta map[string][]byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func visibleCompletionAlreadyEmitted(
	loop *eventLoopContext,
	finalChoices []model.Choice,
	finalStateDelta map[string][]byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldMarkCompletionSnapshotOnly(
	loop *eventLoopContext,
	choices []model.Choice,
	finalStateDelta map[string][]byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isResumeRun(loop *eventLoopContext) bool { _ = "STUB: not implemented"; return false }

func runProducedAssistantContent(loop *eventLoopContext) bool {
	_ = "STUB: not implemented"
	return false
}

func isSnapshotOnlyVisibleGraphCompletion(e *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func cloneChoices(choices []model.Choice) []model.Choice { _ = "STUB: not implemented"; return nil }

func assistantChoiceSignature(choices []model.Choice) string { _ = "STUB: not implemented"; return "" }

func finalResponseIDFromStateDelta(finalStateDelta map[string][]byte) string {
	_ = "STUB: not implemented"
	return ""
}

func baselineFinalResponseID(sess *session.Session, runtimeState map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func collectPriorAssistantResponseIDs(sess *session.Session) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func collectPriorAssistantResponseIDsForLineage(
	loop *eventLoopContext,
	sess *session.Session,
	lineageKey string,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func collectPriorAssistantChoiceSignatures(sess *session.Session) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func collectPriorAssistantChoiceSignaturesForLineage(
	loop *eventLoopContext,
	sess *session.Session,
	lineageKey string,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func interruptedAssistantSignatureKey(
	requestID string,
	invocationID string,
	choices []model.Choice,
) string {
	_ = "STUB: not implemented"
	return ""
}

func baselineFinalResponseIDFromRuntimeState(runtimeState map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func stringValueFromRuntimeState(value any) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func completionMetadataFinalResponseID(value any) string { _ = "STUB: not implemented"; return "" }

func finalResponseTextFromStateDelta(finalStateDelta map[string][]byte) string {
	_ = "STUB: not implemented"
	return ""
}

func assistantChoicePrimaryContent(choices []model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *runner) rewriteUserMessage(
	ctx context.Context,
	appName string,
	userID string,
	sessionID string,
	message model.Message,
	ro agent.RunOptions,
) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) resolveCurrentTurnMessages(
	ctx context.Context,
	appName string,
	userID string,
	sessionID string,
	message model.Message,
	ro agent.RunOptions,
) (model.Message, []model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil, nil
}

func (r *runner) persistCurrentTurnMessages(
	ctx context.Context,
	sess *session.Session,
	invocation *agent.Invocation,
	ag agent.Agent,
	message model.Message,
	persistedCurrentTurnMessages []model.Message,
	ro agent.RunOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// shouldAppendUserMessage checks if the incoming user message should be
// appended to the session.
func shouldAppendUserMessage(message model.Message, seed []model.Message) bool {
	_ = "STUB: not implemented"
	return false
}

// Only a trailing seeded user turn can cover the incoming user message.

func mergeCurrentTurnMessagesIntoSeed(
	seed []model.Message,
	original model.Message,
	currentTurn []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func filterPayloadMessages(messages []model.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func normalizeQueuedUserMessage(
	message model.Message,
) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

// ensureErrorEventContent ensures that error events have valid content.
// This is necessary because some models return error responses without content,
// which would otherwise be discarded by the session service.
func ensureErrorEventContent(e *event.Event) { _ = "STUB: not implemented"; return }

// If content is valid (non-empty), do nothing.

// Ensure Choices slice exists

// Populate content if empty

// Ensure FinishReason is set

// RunWithMessages is a convenience helper that lets callers pass a full
// conversation history ([]model.Message) directly. The messages seed the LLM
// request while the runner continues to merge in newer session events. It
// preserves backward compatibility by delegating to Runner.Run with an empty
// message and a RunOption that carries the conversation history.
func RunWithMessages(
	ctx context.Context,
	r Runner,
	userID string,
	sessionID string,
	messages []model.Message,
	runOpts ...agent.RunOption,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Derive the latest user message for invocation state compatibility
// (e.g., used by GraphAgent to set initial user_input).

// enqueueAutoMemoryJob triggers auto memory extraction if memory service is
// configured.
func (r *runner) enqueueAutoMemoryJob(ctx context.Context, sess *session.Session) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) enqueueSessionIngest(
	ctx context.Context,
	sess *session.Session,
	inv *agent.Invocation,
) {
	_ = "STUB: not implemented"
	return
}

// defaultIngestOptions builds the per-request ingestion options the runner
// passes to Ingestor.IngestSession on each turn. The defaults thread the
// session ID through as run_id and the active invocation's agent name through
// as agent_id, giving downstream backends (e.g. mem0) natural grouping keys
// without requiring callers to construct options manually.
func (r *runner) defaultIngestOptions(
	sess *session.Session,
	inv *agent.Invocation,
) []session.IngestOption {
	_ = "STUB: not implemented"
	return nil
}
