//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates how business code can route real summary model
// calls to different summarizers using request-scoped ctx values.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model name to use for both chat and summary generation")
	waitSec   = flag.Int("wait-sec", 12, "Max wait time in seconds for async summary generation")

	billingInput = flag.String(
		"billing-input",
		"I am a VIP customer. Invoice INV-8842 contains a duplicate charge of $129. Draft a short billing escalation note with the facts only.",
		"First turn user input used before the sync summary path",
	)
	supportInput = flag.String(
		"support-input",
		"Switch context. I reset MFA and now I cannot log in on mobile. Give me concise next support steps.",
		"Second turn user input used before the async summary path",
	)
	finalInput = flag.String(
		"final-input",
		"What support actions are still pending for the mobile login issue?",
		"Third turn user input used after async summary to show isolated summary injection",
	)
)

type summaryRequest struct {
	Tenant string
	Scene  string
}

type summaryMode string

const (
	summaryModeSync  summaryMode = "sync"
	summaryModeAsync summaryMode = "async"
)

type summaryRequestKey struct{}
type summaryModeKey struct{}

func main() {
	flag.Parse()

	d := &contextAwareDemo{
		modelName: *modelName,
		wait:      time.Duration(*waitSec) * time.Second,
	}
	if err := d.run(context.Background(), *billingInput, *supportInput, *finalInput); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

type contextAwareDemo struct {
	modelName string
	wait      time.Duration

	runner         runner.Runner
	sessionService session.Service
	app            string
	userID         string
	sessionID      string

	agentReqSeq int64
}

func (d *contextAwareDemo) run(
	ctx context.Context,
	billingInput string,
	supportInput string,
	finalInput string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *contextAwareDemo) setup() error { _ = "STUB: not implemented"; return nil }

func (d *contextAwareDemo) runTurn(
	ctx context.Context,
	req summaryRequest,
	input string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *contextAwareDemo) beforeAgentModel(
	_ context.Context,
	args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *contextAwareDemo) createSummaryWithRequest(
	ctx context.Context,
	req summaryRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *contextAwareDemo) enqueueSummaryWithRequest(
	ctx context.Context,
	req summaryRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *contextAwareDemo) waitSummary(
	ctx context.Context,
	req summaryRequest,
	wantContains string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *contextAwareDemo) readSummary(
	ctx context.Context,
	req summaryRequest,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *contextAwareDemo) fetchSession(ctx context.Context) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *contextAwareDemo) filterKey(req summaryRequest) string {
	_ = "STUB: not implemented"
	return ""
}

// WithSummaryRequest stores business request metadata on ctx.
func WithSummaryRequest(ctx context.Context, req summaryRequest) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SummaryRequestFromContext reads business request metadata from ctx.
func SummaryRequestFromContext(ctx context.Context) (summaryRequest, bool) {
	_ = "STUB: not implemented"
	return *new(summaryRequest), false
}

// WithSummaryMode stores a business-defined sync/async marker on ctx.
func WithSummaryMode(ctx context.Context, mode summaryMode) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SummaryModeFromContext reads the business-defined sync/async marker.
func SummaryModeFromContext(ctx context.Context) (summaryMode, bool) {
	_ = "STUB: not implemented"
	return *new(summaryMode), false
}

type routingSummarizer struct {
	billingSync  summary.SessionSummarizer
	billingAsync summary.SessionSummarizer
	supportSync  summary.SessionSummarizer
	supportAsync summary.SessionSummarizer
}

var _ summary.ContextAwareSummarizer = (*routingSummarizer)(nil)

func newRoutingSummarizer(modelName string) *routingSummarizer {
	_ = "STUB: not implemented"
	return nil
}

func (r *routingSummarizer) ShouldSummarize(sess *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *routingSummarizer) ShouldSummarizeWithContext(
	ctx context.Context,
	sess *session.Session,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *routingSummarizer) Summarize(
	ctx context.Context,
	sess *session.Session,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *routingSummarizer) SetPrompt(prompt string) { _ = "STUB: not implemented"; return }

func (r *routingSummarizer) SetModel(m model.Model) { _ = "STUB: not implemented"; return }

func (r *routingSummarizer) Metadata() map[string]any { _ = "STUB: not implemented"; return nil }

func (r *routingSummarizer) route(ctx context.Context) summary.SessionSummarizer {
	_ = "STUB: not implemented"
	return *new(summary.SessionSummarizer)
}

func (r *routingSummarizer) all() []summary.SessionSummarizer {
	_ = "STUB: not implemented"
	return nil
}

func newRouteSummarizer(modelName string, routeName string, prompt string) summary.SessionSummarizer {
	_ = "STUB: not implemented"
	return *new(summary.SessionSummarizer)
}

func billingPrompt(mode string) string { _ = "STUB: not implemented"; return "" }

func supportPrompt(mode string) string { _ = "STUB: not implemented"; return "" }

func isSessionSummaryMessage(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func preview(s string, max int) string { _ = "STUB: not implemented"; return "" }

func defaultString(s, fallback string) string { _ = "STUB: not implemented"; return "" }

func normalizeFilterComponent(s string) string { _ = "STUB: not implemented"; return "" }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
