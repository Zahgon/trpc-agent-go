//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package engine implements PromptIter orchestration and runtime flow for a generation round.
package engine

import (
	"context"

	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/backwarder"
)

// BackwardOptions configures backward-stage execution behavior.
type BackwardOptions struct {
	// CaseParallelismEnabled enables concurrent backward processing across eval cases.
	CaseParallelismEnabled bool
	// CaseParallelism caps concurrent backward processing across eval cases when CaseParallelismEnabled is true. Zero uses GOMAXPROCS.
	CaseParallelism int
}

// CaseBackwardResult stores all step gradients produced for one eval case.
type CaseBackwardResult struct {
	// EvalSetID identifies the source evaluation set.
	EvalSetID string
	// EvalCaseID identifies the source evaluation case.
	EvalCaseID string
	// StepGradients stores gradients per step in topological order.
	StepGradients []promptiter.StepGradient
}

// BackwardResult stores aggregated backward outputs for every case in this round.
type BackwardResult struct {
	// Cases stores per-case backward outputs for downstream aggregation.
	Cases []CaseBackwardResult
}

func (e *engine) backward(
	ctx context.Context,
	structure *structureState,
	profile *promptiter.Profile,
	train *EvaluationResult,
	losses []promptiter.CaseLoss,
	targetSurfaceSet targetSurfaceSet,
	options BackwardOptions,
) (*BackwardResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type caseResultKey struct {
	evalSetID  string
	evalCaseID string
}

func indexCaseResults(result *EvaluationResult) map[caseResultKey]CaseResult {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) backwardCase(
	ctx context.Context,
	structure *structureState,
	overrideIndex map[string]promptiter.SurfaceOverride,
	evalCase CaseResult,
	caseLoss promptiter.CaseLoss,
	targetSurfaceSet targetSurfaceSet,
) (*CaseBackwardResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type indexedTraceStep struct {
	step  atrace.Step
	order int
}

func indexTraceSteps(
	structure *structureState,
	trace *atrace.Trace,
) (map[string]indexedTraceStep, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeIncomingPackets(packets []backwarder.GradientPacket) []backwarder.GradientPacket {
	_ = "STUB: not implemented"
	return nil
}

func buildBackwardRequest(
	structure *structureState,
	overrideIndex map[string]promptiter.SurfaceOverride,
	traceIndex map[string]indexedTraceStep,
	evalCase CaseResult,
	step atrace.Step,
	incoming []backwarder.GradientPacket,
	targetSurfaceSet targetSurfaceSet,
) (*backwarder.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func allowedGradientSurfaceIDsOrNil(
	targetSurfaceSet targetSurfaceSet,
	allowedGradientSurfaceIDs []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func cloneTraceSnapshot(snapshot *atrace.Snapshot) *atrace.Snapshot {
	_ = "STUB: not implemented"
	return nil
}
