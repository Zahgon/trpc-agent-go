//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package guardrail provides the top-level guardrail plugin facade.
package guardrail

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/approval"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/promptinjection"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/unsafeintent"
)

// Plugin is the top-level guardrail plugin facade.
type Plugin struct {
	name            string
	approval        *approval.Plugin
	promptInjection *promptinjection.Plugin
	unsafeIntent    *unsafeintent.Plugin
}

// New creates a new guardrail plugin.
func New(options ...Option) (*Plugin, error) { _ = "STUB: not implemented"; return nil, nil }

// Name implements plugin.Plugin.
func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (p *Plugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

// Close implements plugin.Closer when sub-capabilities need cleanup.
func (p *Plugin) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
