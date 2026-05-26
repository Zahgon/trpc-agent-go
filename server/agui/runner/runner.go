//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package runner wraps a trpc-agent-go runner and translates it to AG-UI events.
package runner

import (
	"context"
	"errors"
	"sync"
	"time"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/types"
	"go.opentelemetry.io/otel/trace"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	trunner "trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/internal/track"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/translator"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	// ErrRunAlreadyExists is returned when a run with the same key is already running.
	ErrRunAlreadyExists = errors.New("agui: run already exists")
	// ErrRunNotFound is returned when a run key cannot be found.
	ErrRunNotFound = errors.New("agui: run not found")
	// errExplicitCancel marks a run that was terminated by the AG-UI cancel API.
	errExplicitCancel = errors.New("agui: explicit cancel")
)

const (
	toolResultInputEventAuthor = "agui.runner"
	errResolveExternalTools    = "resolve external tools: %w"
)

// Runner executes AG-UI runs and emits AG-UI events.
type Runner interface {
	// Run starts processing one AG-UI run request and returns a channel of AG-UI events.
	Run(ctx context.Context, runAgentInput *adapter.RunAgentInput) (<-chan aguievents.Event, error)
}

// New wraps a trpc-agent-go runner with AG-UI specific translation logic.
func New(r trunner.Runner, opt ...Option) Runner { _ = "STUB: not implemented"; return *new(Runner) }

// runner is the default implementation of the Runner.
type runner struct {
	appName                                   string
	appNameResolver                           AppNameResolver
	runner                                    trunner.Runner
	translatorFactory                         TranslatorFactory
	graphNodeLifecycleActivityEnabled         bool
	graphNodeInterruptActivityEnabled         bool
	graphNodeInterruptActivityTopLevelOnly    bool
	reasoningContentEnabled                   bool
	eventSourceMetadataEnabled                bool
	userIDResolver                            UserIDResolver
	translateCallbacks                        *translator.Callbacks
	runAgentInputHook                         RunAgentInputHook
	stateResolver                             StateResolver
	runOptionResolver                         RunOptionResolver
	tracker                                   track.Tracker
	runningMu                                 sync.Mutex
	running                                   map[session.Key]*sessionContext
	startSpan                                 StartSpan
	flushInterval                             time.Duration
	postRunFinalizationTimeout                time.Duration
	timeout                                   time.Duration
	cancelOnContextDoneEnabled                bool
	messagesSnapshotFollowEnabled             bool
	messagesSnapshotFollowMaxDuration         time.Duration
	messagesSnapshotRunLifecycleEventsEnabled bool
	toolResultInputTranslationEnabled         bool
	toolCallDeltaStreamingEnabled             bool
	streamingToolResultActivityEnabled        bool
}

type sessionContext struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
}

type runInput struct {
	key             session.Key
	threadID        string
	runID           string
	userID          string
	messages        *runAgentMessages
	runOption       []agent.RunOption
	translator      translator.Translator
	enableTrack     bool
	span            trace.Span
	resume          *resumeInfo
	terminalEmitted bool
}

type runAgentMessages struct {
	inputMessage *model.Message
	inputID      string
	userMessage  *types.Message
	toolMessages []toolResultInputMessage
}

type toolResultInputMessage struct {
	message   model.Message
	messageID string
}

type resumeInfo struct {
	lineageID    string
	checkpointID string
	resumeMap    map[string]any
	resumeSet    bool
	resumeValue  any
}

func inputMessagesFromRunAgentInput(input *adapter.RunAgentInput) (*runAgentMessages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolMessagesFromRunAgentInput(messages []types.Message) (*runAgentMessages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func withToolResultMessageRewriter(toolMessages []toolResultInputMessage) agent.RunOption {
	_ = "STUB: not implemented"
	return *new(agent.RunOption)
}

func toolResultModelMessages(toolMessages []toolResultInputMessage) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func mergeToolResultRewriteMessages(
	rewritten []model.Message,
	toolResults []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// Run starts processing one AG-UI run request and returns a channel of AG-UI events.
func (r *runner) Run(ctx context.Context, runAgentInput *adapter.RunAgentInput) (<-chan aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) run(ctx context.Context, cancel context.CancelCauseFunc, key session.Key, input *runInput, events chan<- aguievents.Event) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) flushTrack(ctx context.Context, key session.Key) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) emitPostRunTerminalEvent(ctx context.Context, events chan<- aguievents.Event, input *runInput) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) newPostRunContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (r *runner) emitPostRunFinalization(ctx context.Context, events chan<- aguievents.Event, input *runInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) emitPostRunFinalizationEvents(ctx context.Context, events chan<- aguievents.Event, input *runInput) error {
	_ = "STUB: not implemented"
	return nil
}

func parseResumeInfo(opt []agent.RunOption) *resumeInfo { _ = "STUB: not implemented"; return nil }

func newGraphInterruptResumeEvent(info *resumeInfo) *aguievents.ActivityDeltaEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) emitToolResultEvents(ctx context.Context, events chan<- aguievents.Event, input *runInput) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) emitToolResultEvent(
	ctx context.Context,
	events chan<- aguievents.Event,
	input *runInput,
	toolMessage toolResultInputMessage,
) bool {
	_ = "STUB: not implemented"
	return false
}

// newToolResultInputEvent normalizes a tool-result input into an internal event for translation.
func newToolResultInputEvent(messageID string, msg *model.Message) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) handleAgentEvent(ctx context.Context, events chan<- aguievents.Event, input *runInput, event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) applyRunAgentInputHook(ctx context.Context,
	input *adapter.RunAgentInput) (*adapter.RunAgentInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) resolveAppName(ctx context.Context, input *adapter.RunAgentInput) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *runner) handleBeforeTranslate(ctx context.Context, event *event.Event) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *runner) handleAfterTranslate(ctx context.Context, event aguievents.Event) (aguievents.Event, error) {
	_ = "STUB: not implemented"
	return *new(aguievents.Event), nil
}

func (r *runner) emitEvent(ctx context.Context, events chan<- aguievents.Event, event aguievents.Event,
	input *runInput) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) shouldTrackEvent(event aguievents.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) recordUserMessage(ctx context.Context, key session.Key, message *types.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) newExecutionContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelCauseFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelCauseFunc)
}

func (r *runner) register(key session.Key, ctx context.Context, cancel context.CancelCauseFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runner) unregister(key session.Key) { _ = "STUB: not implemented"; return }

func (r *runner) recordTrackEvent(ctx context.Context, key session.Key, event aguievents.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func isExplicitRunCancel(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func contextDoneMessage(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
