//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package review

import (
	"context"
)

const (
	reviewerUserIDPrefix    = "promptinjection-reviewer-user:"
	reviewerSessionIDPrefix = "promptinjection-reviewer-session:"
)

// Option configures the built-in prompt injection reviewer.
type Option func(*options)

// UserIDSupplier returns the user ID used for the internal reviewer run.
type UserIDSupplier func(ctx context.Context, req *Request) (string, error)

// SessionIDSupplier returns the session ID used for the internal reviewer run.
type SessionIDSupplier func(ctx context.Context, req *Request) (string, error)

type options struct {
	systemPrompt      string
	userIDSupplier    UserIDSupplier
	sessionIDSupplier SessionIDSupplier
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithSystemPrompt overrides the built-in reviewer system prompt.
func WithSystemPrompt(prompt string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserIDSupplier overrides the user ID supplier for reviewer runs.
func WithUserIDSupplier(supplier UserIDSupplier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionIDSupplier overrides the session ID supplier for reviewer runs.
func WithSessionIDSupplier(supplier SessionIDSupplier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func defaultUserIDSupplier(ctx context.Context, req *Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func defaultSessionIDSupplier(ctx context.Context, req *Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
