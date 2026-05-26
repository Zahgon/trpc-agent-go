//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package promptinjection

import "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/promptinjection/review"

const defaultPluginName = "promptinjection"

// Option configures the prompt injection plugin.
type Option func(*options)

type options struct {
	name     string
	reviewer review.Reviewer
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithName sets the plugin name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReviewer sets the mandatory reviewer used for prompt injection decisions.
func WithReviewer(reviewer review.Reviewer) Option { _ = "STUB: not implemented"; return *new(Option) }
