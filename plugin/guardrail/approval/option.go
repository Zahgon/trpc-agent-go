//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package approval

import "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/approval/review"

const defaultPluginName = "approval"

// Option configures the approval plugin.
type Option func(*options)

type options struct {
	name              string
	reviewer          review.Reviewer
	defaultToolPolicy ToolPolicy
	toolPolicies      map[string]ToolPolicy
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithName sets the plugin name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReviewer sets the reviewer used for approval-required tool calls.
func WithReviewer(reviewer review.Reviewer) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDefaultToolPolicy sets the default policy used when no explicit tool policy exists.
func WithDefaultToolPolicy(policy ToolPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithToolPolicy sets the policy for a single tool name.
func WithToolPolicy(name string, policy ToolPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
