//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package hedge

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const defaultDelay = 100 * time.Millisecond

type options struct {
	candidates    []model.Model
	name          string
	contextWindow int
	delay         time.Duration
	delays        []time.Duration
}

func newOptions(opt ...Option) options { _ = "STUB: not implemented"; return *new(options) }

// Option configures a hedge model.
type Option func(*options)

// WithCandidates appends hedge candidates in launch order.
// Multiple calls accumulate candidates instead of replacing them.
func WithCandidates(candidates ...model.Model) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithName sets a stable logical model name for the hedge wrapper.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextWindow sets the model context window size in tokens for this
// hedge wrapper. If unset, the wrapper reports a context window only when all
// candidate models report the same positive context window.
func WithContextWindow(tokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDelay sets a fixed interval between successive hedge launches.
func WithDelay(delay time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDelays sets absolute launch offsets for candidates[1:].
func WithDelays(delays ...time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
