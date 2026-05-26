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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/aggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/backwarder"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/optimizer"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Engine orchestrates a complete PromptIter lifecycle across evaluation, optimization, acceptance, and stop decisions.
type Engine interface {
	// Describe returns the current structure snapshot used to build traces and profiles.
	Describe(ctx context.Context) (*astructure.Snapshot, error)
	// Run executes multi-round optimization using training and validation feedback loops.
	Run(ctx context.Context, request *RunRequest, opts ...Option) (*RunResult, error)
}

// RunRequest carries the inputs required to start PromptIter optimization.
type RunRequest struct {
	// Train identifies evaluation data used to generate gradients.
	Train []EvalSetInput `json:"train"`
	// Validation identifies evaluation data used for patch acceptance checks.
	Validation []EvalSetInput `json:"validation"`
	// InitialProfile is the baseline profile for round one optimization.
	InitialProfile *promptiter.Profile
	// Teacher executes trace generation requests for evaluation.
	Teacher runner.Runner
	// Judge evaluates generated outputs and returns scoring details.
	Judge runner.Runner
	// EvaluationOptions configures how training and validation runs are executed.
	EvaluationOptions EvaluationOptions
	// BackwardOptions configures backward-stage execution.
	BackwardOptions BackwardOptions
	// AggregationOptions configures aggregation-stage execution.
	AggregationOptions AggregationOptions
	// OptimizerOptions configures optimizer-stage execution.
	OptimizerOptions OptimizerOptions
	// AcceptancePolicy controls minimum quality gain required to accept patching.
	AcceptancePolicy AcceptancePolicy
	// StopPolicy controls termination conditions between rounds.
	StopPolicy StopPolicy
	// MaxRounds is the hard cap for outer optimization iterations.
	MaxRounds int
	// TargetSurfaceIDs limits this run to optimizing only the listed surfaces.
	TargetSurfaceIDs []string
}

// EvalSetInput identifies one evaluation set and optional case filters for a PromptIter run.
type EvalSetInput struct {
	// EvalSetID identifies the eval set to execute.
	EvalSetID string `json:"evalSetId"`
	// EvalCaseIDs limits this eval set to specific eval cases when set.
	EvalCaseIDs []string `json:"evalCaseIds,omitempty"`
	// LossHints stores operator-provided loss reasons keyed by failed eval case and metric.
	LossHints []LossHint `json:"lossHints,omitempty"`
}

// LossHint carries operator-provided context for one failed metric on one eval case.
type LossHint struct {
	// EvalCaseID identifies the eval case to receive this hint.
	EvalCaseID string `json:"evalCaseId"`
	// MetricName identifies the failed metric to receive this hint.
	MetricName string `json:"metricName"`
	// Severity indicates how urgently this hint should influence optimization when set.
	Severity promptiter.LossSeverity `json:"severity,omitempty"`
	// Reason stores the manual loss text used by gradient computation.
	Reason string `json:"reason"`
}

// RunStatus identifies the lifecycle state of one PromptIter run view.
type RunStatus string

const (
	// RunStatusQueued indicates that the run has been created but has not started execution.
	RunStatusQueued RunStatus = "queued"
	// RunStatusRunning indicates that the run is actively executing.
	RunStatusRunning RunStatus = "running"
	// RunStatusSucceeded indicates that the run finished successfully.
	RunStatusSucceeded RunStatus = "succeeded"
	// RunStatusFailed indicates that the run finished with an error.
	RunStatusFailed RunStatus = "failed"
	// RunStatusCanceled indicates that the run was canceled before completion.
	RunStatusCanceled RunStatus = "canceled"
)

// RunResult stores the state and historical trace of one PromptIter execution.
type RunResult struct {
	// AppName identifies the PromptIter target app that owns this run.
	AppName string
	// ID uniquely identifies this run when the caller uses manager-backed execution.
	ID string
	// Status stores the lifecycle state of the run.
	Status RunStatus
	// CurrentRound stores the latest round started by the run.
	CurrentRound int
	// Structure is the snapshot used for all rounds in this request.
	Structure *astructure.Snapshot
	// BaselineValidation stores the accepted baseline validation result before optimization rounds.
	BaselineValidation *EvaluationResult
	// AcceptedProfile is the profile that passed acceptance and can be published.
	AcceptedProfile *promptiter.Profile
	// Rounds stores intermediate results of every optimization round.
	Rounds []RoundResult
	// ErrorMessage stores the terminal run error when the run failed or was canceled.
	ErrorMessage string
}

// RoundResult captures all observable state for one optimization round.
type RoundResult struct {
	// Round is the one-based index of this optimization cycle.
	Round int
	// InputProfile is the profile evaluated at the start of this round.
	InputProfile *promptiter.Profile
	// Train is the train-set result used for gradient generation.
	Train *EvaluationResult
	// Losses stores terminal losses extracted from train traces.
	Losses []promptiter.CaseLoss
	// Backward stores backward outputs grouped by sample.
	Backward *BackwardResult
	// Aggregation stores gradient merges that remove duplicated surface signals.
	Aggregation *AggregationResult
	// Patches stores optimizer suggestions before acceptance and commit.
	Patches *promptiter.PatchSet
	// OutputProfile is the candidate profile created from generated patches.
	OutputProfile *promptiter.Profile
	// Validation is the validation result used for acceptance.
	Validation *EvaluationResult
	// Acceptance is the acceptance output for this round.
	Acceptance *AcceptanceDecision
	// Stop indicates whether the round triggered an early stop condition.
	Stop *StopDecision
}

// engine is the default Engine implementation.
type engine struct {
	// targetAgent exports the current PromptIter structure for the optimization target.
	targetAgent agent.Agent
	// agentEvaluator executes train and validation evaluations through the shared evaluation framework.
	agentEvaluator evaluation.AgentEvaluator
	// backwarder computes sample-level gradient packets from terminal losses.
	backwarder backwarder.Backwarder
	// aggregator merges sample gradients into per-surface aggregated gradient.
	aggregator aggregator.Aggregator
	// optimizer translates aggregated gradients into patch candidates.
	optimizer optimizer.Optimizer
}

// New creates an Engine implementation with injected collaborators.
func New(ctx context.Context,
	targetAgent agent.Agent,
	agentEvaluator evaluation.AgentEvaluator,
	backwarder backwarder.Backwarder,
	aggregator aggregator.Aggregator,
	optimizer optimizer.Optimizer) (Engine, error) {
	_ = "STUB: not implemented"
	return *new(Engine), nil
}

// Describe returns the structure snapshot used for the current optimization session.
func (e *engine) Describe(ctx context.Context) (*astructure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run executes all optimization stages in sequence for each configured round.
func (e *engine) Run(ctx context.Context, request *RunRequest, opts ...Option) (*RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *engine) run(
	ctx context.Context,
	request *RunRequest,
	observer Observer,
) (*RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *engine) validateRunRequest(request *RunRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engine) describeStructure(ctx context.Context) (*astructure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *engine) executeRound(
	ctx context.Context,
	request *RunRequest,
	structure *structureState,
	targetSurfaceSet targetSurfaceSet,
	observer Observer,
	evaluationOptions EvaluationOptions,
	acceptedProfile *promptiter.Profile,
	acceptedValidationScore float64,
	roundNumber int,
) (*RoundResult, float64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (e *engine) newEvaluationRequest(
	inputs []EvalSetInput,
	profile *promptiter.Profile,
	teacher runner.Runner,
	judge runner.Runner,
	options EvaluationOptions,
) *EvaluationRequest {
	_ = "STUB: not implemented"
	return nil
}

func validateEvalSetInputs(role string, inputs []EvalSetInput) error {
	_ = "STUB: not implemented"
	return nil
}

func isValidLossHintSeverity(severity promptiter.LossSeverity) bool {
	_ = "STUB: not implemented"
	return false
}

func runIndexedParallel(
	ctx context.Context,
	count int,
	parallelism int,
	fn func(context.Context, int) error,
) error {
	_ = "STUB: not implemented"
	return nil
}
