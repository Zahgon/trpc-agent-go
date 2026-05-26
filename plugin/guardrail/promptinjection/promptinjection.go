//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package promptinjection provides a runner-scoped prompt injection guardrail plugin.
package promptinjection

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	promptreview "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/promptinjection/review"
)

// Plugin is the prompt injection guardrail implementation.
type Plugin struct {
	name         string
	reviewer     promptreview.Reviewer
	tokenCounter model.TokenCounter
}

// New creates a new prompt injection plugin.
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

func promptInjectionDenyMessage(decision *promptreview.Decision) string {
	_ = "STUB: not implemented"
	return ""
}
