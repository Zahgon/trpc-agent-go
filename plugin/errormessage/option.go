//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package errormessage

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
)

const (
	defaultPluginName   = "error_message_rewriter"
	defaultFinishReason = "error"
)

// Resolver returns the assistant-visible content to attach to an error event.
//
// Returning ok=false (or an empty string) leaves the event untouched, which
// falls back to Runner's built-in generic fallback message.
type Resolver func(
	ctx context.Context,
	inv *agent.Invocation,
	e *event.Event,
) (content string, ok bool)

// Option configures the error message plugin.
type Option func(*options)

type options struct {
	name         string
	resolver     Resolver
	finishReason string
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithName sets the plugin name. The name must be unique within a Runner.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContent sets a static message that replaces the framework's default
// fallback content for all error events.
//
// It is shorthand for a Resolver that returns (content, content != ""), so
// passing an empty string is a no-op and the event is left untouched, which
// means Runner's built-in fallback message still applies.
func WithContent(content string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithResolver registers a Resolver that computes the assistant-visible
// content for each error event. It is called only when the event is an error
// event with no existing valid content, so returning a message is always a
// safe, non-destructive operation.
func WithResolver(resolver Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFinishReason overrides the FinishReason attached to the synthesised
// assistant choice. It is used only when the choice does not carry a
// FinishReason yet. The default is "error".
func WithFinishReason(reason string) Option { _ = "STUB: not implemented"; return *new(Option) }
