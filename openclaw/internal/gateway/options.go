//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package gateway

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/memoryfile"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/persona"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

// SessionIDFunc builds a session ID for the inbound message.
type SessionIDFunc func(InboundMessage) (string, error)

// RunOptionInput describes one gateway run before invoking the runner.
type RunOptionInput struct {
	Inbound    InboundMessage
	UserID     string
	SessionID  string
	RequestID  string
	Message    model.Message
	Trace      *debugrecorder.Trace
	Extensions map[string]json.RawMessage
}

// RunOptionResolver decorates context and options for one gateway run.
type RunOptionResolver func(
	ctx context.Context,
	input RunOptionInput,
) (context.Context, []agent.RunOption, error)

type options struct {
	basePath      string
	messagesPath  string
	streamPath    string
	streamPathSet bool
	statusPath    string
	cancelPath    string
	healthPath    string

	maxBodyBytes int64
	maxPartBytes int64

	partFetcher          partFetcher
	allowPrivatePartURLs bool
	allowedPartPatterns  []string
	audioTranscriber     audioTranscriber
	appName              string

	sessionIDFunc SessionIDFunc

	allowUsers      map[string]struct{}
	requireMention  bool
	mentionPatterns []string

	runOptionResolver RunOptionResolver

	recorder *debugrecorder.Recorder
	uploads  *uploads.Store

	personaStore    *persona.Store
	memoryFileStore *memoryfile.Store
}

// Option is a function that configures a gateway server.
type Option func(*options)

func newOptions(opts ...Option) options { _ = "STUB: not implemented"; return *new(options) }

// WithBasePath sets the base path for all gateway endpoints except health.
func WithBasePath(basePath string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesPath sets the relative path for the messages endpoint.
func WithMessagesPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesStreamPath sets the relative path for the streaming
// messages endpoint.
func WithMessagesStreamPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStatusPath sets the relative path for the status endpoint.
func WithStatusPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCancelPath sets the relative path for the cancel endpoint.
func WithCancelPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHealthPath sets the health check endpoint path.
func WithHealthPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxBodyBytes sets the maximum bytes to read from an HTTP body.
func WithMaxBodyBytes(max int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxContentPartBytes sets the maximum bytes to fetch for one
// content part.
func WithMaxContentPartBytes(max int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentPartFetcher sets a custom content-part fetcher.
func WithContentPartFetcher(fetcher partFetcher) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAllowPrivateContentPartURLs allows content parts to fetch URLs that
// resolve to loopback or private network addresses.
func WithAllowPrivateContentPartURLs(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAllowedContentPartDomains restricts content-part URL fetches to the
// provided domains or URL patterns.
//
// Each entry is either:
//   - "example.com" (allows all paths), or
//   - "example.com/path" (allows /path/...).
//
// The host match is case-insensitive and allows subdomains.
func WithAllowedContentPartDomains(domains ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPersonaStore sets the preset persona store used for per-chat
// system-message injection.
func WithPersonaStore(store *persona.Store) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMemoryFileStore sets the file-based memory store used for per-run
// context injection.
func WithMemoryFileStore(store *memoryfile.Store) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAppName sets the app name used by file-based memory injection.
func WithAppName(appName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSessionIDFunc sets a custom session ID function.
func WithSessionIDFunc(fn SessionIDFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAllowUsers sets a user allowlist.
//
// When set, only the listed user IDs are allowed to send messages.
// If called with no arguments, the allowlist becomes empty and all users are
// denied.
func WithAllowUsers(users ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequireMentionInThreads enables mention gating for thread messages.
func WithRequireMentionInThreads(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMentionPatterns sets patterns used for mention gating.
func WithMentionPatterns(patterns ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDebugRecorder enables file-based debug recording for gateway requests.
//
// When set, the gateway creates a per-request trace when no trace is present
// in the request context.
func WithDebugRecorder(rec *debugrecorder.Recorder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRunOptionResolver decorates one gateway run before runner.Run.
func WithRunOptionResolver(resolver RunOptionResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUploadStore persists inbound file parts to stable host paths.
func WithUploadStore(store *uploads.Store) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAudioTranscriber overrides inbound audio transcription.
func WithAudioTranscriber(transcriber audioTranscriber) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
