//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package recorder records runner event streams into reusable evalset assets.
package recorder

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

const (
	defaultPluginName = "evalset_recorder"
)

type options struct {
	name               string
	asyncWriteEnabled  bool
	writeTimeout       time.Duration
	evalSetIDResolver  EvalSetIDResolver
	evalCaseIDResolver EvalCaseIDResolver
	traceModeEnabled   bool
}

func newOptions(opts ...Option) (options, error) {
	_ = "STUB: not implemented"
	return *new(options), nil
}

// Option configures the recorder.
type Option func(*options)

// WithName sets the plugin name used by the runner plugin manager.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAsyncWriteEnabled controls whether the recorder persists turns asynchronously.
func WithAsyncWriteEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWriteTimeout sets the maximum duration for a single persistence attempt.
// If timeout is 0, the recorder uses the original context without adding a deadline.
func WithWriteTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTraceModeEnabled controls whether recorder output is persisted as trace-mode actual conversations.
func WithTraceModeEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// EvalSetIDResolver derives the target EvalSetID for the current invocation.
type EvalSetIDResolver func(ctx context.Context, inv *agent.Invocation) (string, error)

// WithEvalSetIDResolver sets the resolver for EvalSetID.
func WithEvalSetIDResolver(resolver EvalSetIDResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// EvalCaseIDResolver derives the target EvalCaseID (stored in EvalCase.EvalID) for the current invocation.
type EvalCaseIDResolver func(ctx context.Context, inv *agent.Invocation) (string, error)

// WithEvalCaseIDResolver sets the resolver for EvalCaseID (EvalCase.EvalID).
func WithEvalCaseIDResolver(resolver EvalCaseIDResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
