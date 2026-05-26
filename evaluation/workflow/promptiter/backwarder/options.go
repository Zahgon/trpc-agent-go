//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package backwarder computes backward propagation outputs from trace and gradient data.
package backwarder

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// options stores optional backwarder behavior toggles.
type options struct {
	runOptions        []agent.RunOption
	messageBuilder    MessageBuilder
	userIDSupplier    UserIDSupplier
	sessionIDSupplier SessionIDSupplier
}

// Option mutates backwarder options during construction.
type Option func(*options)

// newOptions applies all backwarder options and returns a configured options set.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithRunOptions appends runner options for backward invocations.
func WithRunOptions(runOptions ...agent.RunOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMessageBuilder overrides how backward requests are encoded for the runner.
func WithMessageBuilder(builder MessageBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// UserIDSupplier provides a user ID for one backward runner invocation.
type UserIDSupplier func(ctx context.Context) string

func defaultUserIDSupplier() UserIDSupplier { _ = "STUB: not implemented"; return *new(UserIDSupplier) }

// WithUserIDSupplier overrides how backward runner user IDs are generated.
func WithUserIDSupplier(supplier UserIDSupplier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SessionIDSupplier provides a session ID for one backward runner invocation.
type SessionIDSupplier func(ctx context.Context) string

func defaultSessionIDSupplier() SessionIDSupplier {
	_ = "STUB: not implemented"
	return *new(SessionIDSupplier)
}

// WithSessionIDSupplier overrides how backward runner session IDs are generated.
func WithSessionIDSupplier(supplier SessionIDSupplier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
