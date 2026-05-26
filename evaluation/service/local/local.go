//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package local provides a local implementation of service.Service.
package local

import (
	"context"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion"
	criterionllm "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	metricregistry "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/service"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/usersimulation"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const reasonSeparator = ";"

// local is a local implementation of service.Service.
type local struct {
	runner                            runner.Runner
	expectedRunner                    runner.Runner
	evalSetManager                    evalset.Manager
	evalResultManager                 evalresult.Manager
	registry                          registry.Registry
	metricRegistry                    metricregistry.Registry
	sessionIDSupplier                 func(ctx context.Context) string
	userSimulator                     usersimulation.Simulator
	callbacks                         *service.Callbacks
	runOptions                        []agent.RunOption
	evalCaseParallelism               int
	evalCaseParallelInferenceEnabled  bool
	evalCaseParallelEvaluationEnabled bool
	evalCaseInferencePoolsMu          sync.Mutex
	evalCaseInferencePools            map[int]*ants.PoolWithFunc
	evalCaseEvaluationPoolsMu         sync.Mutex
	evalCaseEvaluationPools           map[int]*ants.PoolWithFunc
}

// New returns a new local evaluation service.
// If no service.Option is provided, the service will use the default options.
func New(runner runner.Runner, opt ...service.Option) (service.Service, error) {
	_ = "STUB: not implemented"
	return *new(service.Service), nil
}

// Close closes the eval service and releases owned resources.
func (s *local) Close() error { _ = "STUB: not implemented"; return nil }

func (s *local) runBeforeEvaluateSetCallbacks(ctx context.Context, callbacks *service.Callbacks, req *service.EvaluateRequest) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *local) runAfterEvaluateSetCallbacks(ctx context.Context, callbacks *service.Callbacks, req *service.EvaluateRequest, result *service.EvalSetRunResult, err error, startTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) runBeforeEvaluateCaseCallbacks(ctx context.Context, callbacks *service.Callbacks, req *service.EvaluateRequest, evalCaseID string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *local) runAfterEvaluateCaseCallbacks(
	ctx context.Context,
	callbacks *service.Callbacks,
	req *service.EvaluateRequest,
	inferenceResult *service.InferenceResult,
	result *evalresult.EvalCaseResult,
	err error,
	startTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Evaluate runs the evaluation on the inference results and returns the eval set run result.
func (s *local) Evaluate(ctx context.Context, req *service.EvaluateRequest, opt ...service.Option) (runResult *service.EvalSetRunResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) resolveMetricExtensions(
	evaluateConfig *service.EvaluateConfig,
	metricRegistry metricregistry.Registry,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) evaluateCaseResults(ctx context.Context, req *service.EvaluateRequest, opts *service.Options) ([]*evalresult.EvalCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) evaluateCaseResultsParallel(ctx context.Context, req *service.EvaluateRequest, opts *service.Options) ([]*evalresult.EvalCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) evaluateCaseResultsSerial(ctx context.Context, req *service.EvaluateRequest, opts *service.Options) ([]*evalresult.EvalCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) evaluateCase(ctx context.Context, req *service.EvaluateRequest, inferenceResult *service.InferenceResult, opts *service.Options) (result *evalresult.EvalCaseResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) failedEvalCaseResult(evalSetID string, inferenceResult *service.InferenceResult, errorMessage string) *evalresult.EvalCaseResult {
	_ = "STUB: not implemented"
	return nil
}

// evaluatePerCase runs the evaluation on the inference result and returns the case evaluation result.
func (s *local) evaluatePerCase(ctx context.Context, inferenceResult *service.InferenceResult,
	evaluateConfig *service.EvaluateConfig, opts *service.Options) (*evalresult.EvalCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overallMetricResults collects the metric results for the entire eval case.

// Iterate through every configured metric and run the evaluation.

// Skip metrics whose evaluator or artifacts are intentionally absent.

// Record the metric outcome for the corresponding invocation.

// Summarize the overall metric results and return the final eval status.

func lookupMetricEvaluator(reg registry.Registry, evalMetric *metric.EvalMetric) (evaluator.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(evaluator.Evaluator), nil
}

func buildCaseRubricIndex(
	evalCaseID string,
	evalMetrics []*metric.EvalMetric,
	caseRubrics []*evalset.EvalCaseRubric,
) (map[string][]*criterionllm.Rubric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCaseEffectiveMetric(evalCaseID string, evalMetric *metric.EvalMetric, caseRubrics []*criterionllm.Rubric) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendCaseRubricsToMetric(evalCaseID string, evalMetric *metric.EvalMetric, caseRubrics []*criterionllm.Rubric) error {
	_ = "STUB: not implemented"
	return nil
}

type caseEvaluationInputs struct {
	actuals   []*evalset.Invocation
	expecteds []*evalset.Invocation
	userID    string
}

func (s *local) prepareCaseEvaluationInputs(
	ctx context.Context,
	inferenceResult *service.InferenceResult,
	evalCase *evalset.EvalCase,
	opts *service.Options,
) (*caseEvaluationInputs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) inferExpectedInferences(
	ctx context.Context,
	evalCase *evalset.EvalCase,
	inputs []*evalset.Invocation,
	sessionID string,
	opts *service.Options,
) ([]*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func attachContextMessages(invocations []*evalset.Invocation, contextMessages []*model.Message) {
	_ = "STUB: not implemented"
	return
}

// In trace mode, Conversation can represent either expected outputs or recorded actual traces for backward compatibility.
// If ActualConversation is provided, Conversation is treated as expecteds aligned by turn.
// If ActualConversation is omitted, Conversation is treated as the actual trace and expecteds are reduced to user-input placeholders.
// If Conversation is omitted but ActualConversation is provided, expecteds are built from ActualConversation as user-input placeholders,
// which represents trace evaluation without expected outputs.
func buildExpectedsForEval(evalCase *evalset.EvalCase) ([]*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func materializeResultCriterion(ctx context.Context, evalMetric *metric.EvalMetric,
	actuals, expecteds []*evalset.Invocation) (*criterion.Criterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func materializeOverallCriterion(ctx context.Context, evalMetric *metric.EvalMetric,
	actuals, expecteds []*evalset.Invocation) (*criterion.Criterion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func materializedPrompt(messages []model.Message) string { _ = "STUB: not implemented"; return "" }

func resolveEvaluatorName(evalMetric *metric.EvalMetric) string {
	_ = "STUB: not implemented"
	return ""
}

// userInputOnlyInvocationsForEval builds placeholder invocations that only preserve user inputs.
// This whitelist prevents trace outputs from being treated as reference answers and stays correct when Invocation gains new fields.
func userInputOnlyInvocationsForEval(conversation []*evalset.Invocation) []*evalset.Invocation {
	_ = "STUB: not implemented"
	return nil
}
