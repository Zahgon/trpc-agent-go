//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package unsafeintent provides a runner-scoped unsafe intent guardrail plugin.
package unsafeintent

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	unsafereview "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/unsafeintent/review"
)

// Plugin is the unsafe intent guardrail implementation.
type Plugin struct {
	name         string
	reviewer     unsafereview.Reviewer
	tokenCounter model.TokenCounter
}

// New creates a new unsafe intent plugin.
func New(options ...Option) (*Plugin, error) { _ = "STUB: not implemented"; return nil, nil }

// Name implements plugin.Plugin.
func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (p *Plugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

func (p *Plugin) beforeModel() model.BeforeModelCallbackStructured {
	_ = "STUB: not implemented"
	return *new(model.BeforeModelCallbackStructured)
}

func (p *Plugin) blockedResponse(content string) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func unsafeIntentDenyMessage(decision *unsafereview.Decision) string {
	_ = "STUB: not implemented"
	return ""
}
