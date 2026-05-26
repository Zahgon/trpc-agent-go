//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package engine implements PromptIter orchestration and runtime flow for a generation round.
package engine

type options struct {
	observer Observer
}

// Option configures one advanced PromptIter run behavior.
type Option func(*options)

// WithObserver appends one runtime observer to the run.
func WithObserver(observer Observer) Option { _ = "STUB: not implemented"; return *new(Option) }

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }
