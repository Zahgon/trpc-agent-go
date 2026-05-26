//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package review provides approval reviewer abstractions and the built-in guardian reviewer.
package review

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Reviewer evaluates a review request and returns an approval decision.
type Reviewer interface {
	Review(ctx context.Context, req *Request) (*Decision, error)
}

// Request is the stable approval request contract passed to reviewers.
type Request struct {
	Action     Action
	Transcript []TranscriptEntry
}

// Action is the exact tool action being reviewed.
type Action struct {
	ToolName        string
	ToolDescription string
	Arguments       json.RawMessage
}

// TranscriptEntry is a compact transcript line used as approval evidence.
type TranscriptEntry struct {
	Role    model.Role
	Content string
}

// Decision is the stable reviewer output consumed by the approval plugin.
type Decision struct {
	Approved  bool
	RiskScore int
	RiskLevel string
	Reason    string
}

type guardianReviewer struct {
	runner            runner.Runner
	systemPrompt      string
	riskThreshold     int
	userIDSupplier    UserIDSupplier
	sessionIDSupplier SessionIDSupplier
}

type decisionPayload struct {
	RiskScore int    `json:"risk_score"`
	RiskLevel string `json:"risk_level"`
	Reason    string `json:"reason"`
}

// New creates the built-in guardian reviewer backed by a runner.
func New(r runner.Runner, options ...Option) (Reviewer, error) {
	_ = "STUB: not implemented"
	return *new(Reviewer), nil
}

func (r *guardianReviewer) Review(ctx context.Context, req *Request) (*Decision, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectDecisionPayload(ctx context.Context, events <-chan *event.Event) (*decisionPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
