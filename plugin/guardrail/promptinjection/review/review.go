//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package review provides reviewer abstractions for prompt injection detection.
package review

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Category is the prompt injection category returned by the reviewer.
type Category string

const (
	// CategorySystemOverride indicates attempts to override higher-priority instructions.
	CategorySystemOverride Category = "system_override"
	// CategoryPolicyBypass indicates attempts to bypass policy or safety controls.
	CategoryPolicyBypass Category = "policy_bypass"
	// CategoryPromptExfiltration indicates attempts to reveal hidden prompts or configuration.
	CategoryPromptExfiltration Category = "prompt_exfiltration"
	// CategoryRoleHijack indicates attempts to impersonate a privileged or higher-priority role.
	CategoryRoleHijack Category = "role_hijack"
	// CategoryToolMisuseInduction indicates attempts to induce unsafe or incorrect tool behavior.
	CategoryToolMisuseInduction Category = "tool_misuse_induction"
)

// TranscriptEntry is a compact transcript line used as review evidence.
type TranscriptEntry struct {
	Role    model.Role
	Content string
}

// Request is the stable prompt injection review request contract.
type Request struct {
	LastUserInput string
	Transcript    []TranscriptEntry
}

// Decision is the stable reviewer output consumed by the prompt injection plugin.
type Decision struct {
	Blocked  bool
	Category Category
	Reason   string
}

// Reviewer evaluates a prompt injection review request and returns a decision.
type Reviewer interface {
	Review(ctx context.Context, req *Request) (*Decision, error)
}

type guardianReviewer struct {
	runner            runner.Runner
	systemPrompt      string
	userIDSupplier    UserIDSupplier
	sessionIDSupplier SessionIDSupplier
}

type decisionPayload struct {
	Blocked  bool     `json:"blocked"`
	Category Category `json:"category"`
	Reason   string   `json:"reason"`
}

// New creates the built-in prompt injection reviewer backed by a runner.
func New(r runner.Runner, options ...Option) (Reviewer, error) {
	_ = "STUB: not implemented"
	return *new(Reviewer), nil
}

func (r *guardianReviewer) Review(ctx context.Context, req *Request) (*Decision, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateDecisionPayload(payload *decisionPayload) error { _ = "STUB: not implemented"; return nil }

func validateCategory(category Category) error { _ = "STUB: not implemented"; return nil }

func collectDecisionPayload(ctx context.Context, events <-chan *event.Event) (*decisionPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
