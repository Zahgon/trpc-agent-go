//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agui

import (
	"time"

	aguirunner "trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/service"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/service/sse"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	defaultBasePath              = "/"
	defaultPath                  = "/"
	defaultServiceFactory        = sse.New
	defaultMessagesSnapshotPath  = "/history"
	defaultCancelPath            = "/cancel"
	defaultMessagesSnapshotState = false
	defaultCancelState           = false
)

// options holds the options for the AG-UI server.
type options struct {
	basePath                string
	path                    string
	serviceFactory          ServiceFactory
	aguiRunnerOptions       []aguirunner.Option
	messagesSnapshotPath    string
	messagesSnapshotEnabled bool
	cancelPath              string
	cancelEnabled           bool
	heartbeatInterval       time.Duration
	appName                 string
	sessionService          session.Service
}

// newOptions creates a new options instance.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option is a function that configures the options.
type Option func(*options)

// WithBasePath sets the base path for service listening, "/" in default.
func WithBasePath(basePath string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPath sets the chat message endpoint path for AG-UI service, "/" in default.
func WithPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCancelPath sets the cancel endpoint path for AG-UI service, "/cancel" in default.
func WithCancelPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCancelEnabled enables the cancel handler.
func WithCancelEnabled(e bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCancelOnContextDoneEnabled controls whether an AG-UI run is canceled when the request context is done.
func WithCancelOnContextDoneEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ServiceFactory is a function that creates AG-UI service.
type ServiceFactory func(runner aguirunner.Runner, opt ...service.Option) service.Service

// WithServiceFactory sets the service factory, sse.New in default.
func WithServiceFactory(f ServiceFactory) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAGUIRunnerOptions sets the AG-UI runner options.
func WithAGUIRunnerOptions(aguiRunnerOpts ...aguirunner.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTimeout sets the maximum execution time for a run, 1h in default.
func WithTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFlushInterval sets how often buffered AG-UI events are flushed for a session.
func WithFlushInterval(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPostRunFinalizationTimeout sets the maximum duration allowed for post-run finalization.
func WithPostRunFinalizationTimeout(d time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHeartbeatInterval sets how often the SSE transport sends heartbeat frames.
func WithHeartbeatInterval(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGraphNodeLifecycleActivityEnabled controls whether the AG-UI server emits graph node lifecycle activity events.
func WithGraphNodeLifecycleActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphNodeInterruptActivityEnabled controls whether the AG-UI server emits graph interrupt activity events.
func WithGraphNodeInterruptActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphNodeInterruptActivityTopLevelOnly controls whether the AG-UI server only emits graph interrupt activity events for the top-level invocation.
func WithGraphNodeInterruptActivityTopLevelOnly(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithReasoningContentEnabled controls whether the AG-UI server emits reasoning content events.
func WithReasoningContentEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEventSourceMetadataEnabled controls whether translated AG-UI events
// include source metadata from the original trpc-agent-go event in rawEvent.
func WithEventSourceMetadataEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToolResultInputTranslationEnabled controls whether echoed tool-result inputs pass through the AG-UI translator.
func WithToolResultInputTranslationEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToolCallDeltaStreamingEnabled controls whether partial tool-call
// arguments are emitted before the final tool call response.
func WithToolCallDeltaStreamingEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamingToolResultActivityEnabled controls whether partial tool-result
// chunks are emitted as activity events while only the final tool result is
// retained on the tool-result path and in message snapshots.
func WithStreamingToolResultActivityEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMessagesSnapshotPath sets the HTTP path for the messages snapshot handler, "/history" in default.
func WithMessagesSnapshotPath(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesSnapshotEnabled enables the MessagesSnapshotHandler.
func WithMessagesSnapshotEnabled(e bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesSnapshotFollowEnabled controls whether the messages snapshot handler tails persisted track events.
func WithMessagesSnapshotFollowEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMessagesSnapshotFollowMaxDuration sets the maximum duration for messages snapshot tailing.
func WithMessagesSnapshotFollowMaxDuration(d time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMessagesSnapshotRunLifecycleEventsEnabled controls whether persisted RUN_* events
// are included as activity messages in MESSAGES_SNAPSHOT.
func WithMessagesSnapshotRunLifecycleEventsEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAppName sets the app name.
func WithAppName(n string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAppNameResolver sets the app name resolver.
func WithAppNameResolver(r aguirunner.AppNameResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionService sets the session service.
func WithSessionService(service session.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
