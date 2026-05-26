//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package guardrail

import (
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/approval"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/promptinjection"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/unsafeintent"
)

const defaultPluginName = "guardrail"

// Option configures the guardrail plugin.
type Option func(*options)

type options struct {
	name            string
	approval        *approval.Plugin
	promptInjection *promptinjection.Plugin
	unsafeIntent    *unsafeintent.Plugin
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithName sets the plugin name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithApproval attaches the approval capability.
func WithApproval(approvalPlugin *approval.Plugin) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPromptInjection attaches the prompt injection capability.
func WithPromptInjection(promptInjectionPlugin *promptinjection.Plugin) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUnsafeIntent attaches the unsafe intent capability.
func WithUnsafeIntent(unsafeIntentPlugin *unsafeintent.Plugin) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
