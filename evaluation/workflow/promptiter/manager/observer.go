//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package manager provides asynchronous PromptIter run lifecycle management on top of the synchronous engine.
package manager

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
)

type observer struct {
	manager *manager
	run     *engine.RunResult
}

func (o *observer) append(ctx context.Context, event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyEvent(event *engine.Event) error { _ = "STUB: not implemented"; return nil }

func (o *observer) applyStructureSnapshot(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyBaselineValidation(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundTrainEvaluation(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundLosses(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundBackward(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundAggregation(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundPatchSet(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundOutputProfile(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundValidation(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *observer) applyRoundCompleted(event *engine.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func invalidEventPayloadError(kind engine.EventKind) error { _ = "STUB: not implemented"; return nil }

func validateEventRound(event *engine.Event) error { _ = "STUB: not implemented"; return nil }

func (o *observer) ensureRound(roundNumber int) *engine.RoundResult {
	_ = "STUB: not implemented"
	return nil
}
