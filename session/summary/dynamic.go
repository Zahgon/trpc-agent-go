//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package summary

import (
	"context"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const metadataDynamicSummarizerType = "dynamic"

var errNoDynamicSummarizerResolved = errors.New("no dynamic summarizer resolved")

// NewDynamicSummarizer creates a summarizer that resolves the actual
// summarizer from the current request context and session at summary time.
//
// It is useful when the session service should be long-lived while the summary
// model, prompt, or checks need to vary per request. Returning nil from resolve
// skips automatic summary checks. Resolver errors also make the automatic
// summary gate return false, while Summarize propagates resolver errors.
// Calling Summarize directly, or forcing a summary without a resolved
// summarizer, returns an error.
func NewDynamicSummarizer(
	resolve func(context.Context, *session.Session) (SessionSummarizer, error),
) SessionSummarizer {
	_ = "STUB: not implemented"
	return *new(SessionSummarizer)
}

type dynamicSummarizer struct {
	resolve func(context.Context, *session.Session) (SessionSummarizer, error)
}

var _ ContextAwareSummarizer = (*dynamicSummarizer)(nil)

func (d *dynamicSummarizer) resolveSummarizer(
	ctx context.Context,
	sess *session.Session,
) (SessionSummarizer, error) {
	_ = "STUB: not implemented"
	return *new(SessionSummarizer), nil
}

// ShouldSummarize checks if the session should be summarized.
func (d *dynamicSummarizer) ShouldSummarize(sess *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

// ShouldSummarizeWithContext resolves the current summarizer and evaluates its
// summary gate with the provided context.
func (d *dynamicSummarizer) ShouldSummarizeWithContext(
	ctx context.Context,
	sess *session.Session,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Summarize resolves the current summarizer and delegates summary generation
// to it.
func (d *dynamicSummarizer) Summarize(
	ctx context.Context,
	sess *session.Session,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetPrompt is a no-op for dynamic summarizers. Configure request-scoped
// prompts in the resolver instead.
func (d *dynamicSummarizer) SetPrompt(string) {
	_ = "STUB: not implemented"

	// SetModel is a no-op for dynamic summarizers. Configure request-scoped models
	// in the resolver instead.
	return
}

func (d *dynamicSummarizer) SetModel(model.Model) {
	_ = "STUB: not implemented"

	// Metadata returns metadata about the dynamic summarizer.
	return
}

func (d *dynamicSummarizer) Metadata() map[string]any { _ = "STUB: not implemented"; return nil }
