//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package review provides reviewer abstractions for unsafe intent detection.
package review

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Category is the unsafe intent category returned by the reviewer.
type Category string

const (
	// CategoryCyberAbuse indicates cyber abuse, malware, intrusion, or exploitation intent.
	CategoryCyberAbuse Category = "cyber_abuse"
	// CategoryCredentialTheft indicates credential theft or account takeover intent.
	CategoryCredentialTheft Category = "credential_theft" // #nosec G101 - This is a classification label, not a credential.
	// CategoryFraudDeception indicates fraud, scam, or deception intent.
	CategoryFraudDeception Category = "fraud_deception"
	// CategoryPrivacyAbuse indicates privacy invasion, stalking, or data abuse intent.
	CategoryPrivacyAbuse Category = "privacy_abuse"
	// CategoryPhysicalHarm indicates violent or physical harm intent.
	CategoryPhysicalHarm Category = "physical_harm"
	// CategorySelfHarm indicates self-harm or suicide assistance intent.
	CategorySelfHarm Category = "self_harm"
	// CategorySexualAbuse indicates sexual abuse or exploitation intent.
	CategorySexualAbuse Category = "sexual_abuse"
	// CategoryOtherUnsafeIntent indicates other clearly unsafe or disallowed intent.
	CategoryOtherUnsafeIntent Category = "other_unsafe_intent"
)

// TranscriptEntry is a compact transcript line used as review evidence.
type TranscriptEntry struct {
	Role    model.Role
	Content string
}

// Request is the stable unsafe intent review request contract.
type Request struct {
	LastUserInput string
	Transcript    []TranscriptEntry
}

// Decision is the stable reviewer output consumed by the unsafe intent plugin.
type Decision struct {
	Blocked  bool
	Category Category
	Reason   string
}

// Reviewer evaluates an unsafe intent review request and returns a decision.
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

// New creates the built-in unsafe intent reviewer backed by a runner.
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
