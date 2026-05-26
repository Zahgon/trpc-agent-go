//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package evaluation

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	metricregistry "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/service"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/usersimulation"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// defaultNumRuns is the default number of runs.
const defaultNumRuns = 1

// options holds the configuration options for the evaluation.
type options struct {
	evalSetManager                    evalset.Manager
	evalResultManager                 evalresult.Manager
	metricManager                     metric.Manager
	registry                          registry.Registry
	metricRegistry                    metricregistry.Registry
	evalService                       service.Service
	expectedRunner                    runner.Runner
	userSimulator                     usersimulation.Simulator
	callbacks                         *service.Callbacks
	judgeRunner                       runner.Runner
	judgeRunnerNumSamples             *int
	numRuns                           int
	evalCaseIDs                       []string
	numRunsParallelEnabled            *bool
	evalCaseParallelism               *int
	evalCaseParallelInferenceEnabled  *bool
	evalCaseParallelEvaluationEnabled *bool
	runDetailsEnabled                 bool
	runDetailsCollector               *runDetailsCollector
	runOptions                        []agent.RunOption
}

// newOptions creates a new options with the default values.
func newOptions(opt ...Option) *options {
	_ = "STUB: not implemented"
	// Initialize options with default values.
	return nil
}

// Apply user options.

// Option defines a function type for configuring the evaluation.
type Option func(*options)

// WithEvalSetManager sets the eval set manager.
func WithEvalSetManager(m evalset.Manager) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEvalResultManager sets the eval result manager.
func WithEvalResultManager(m evalresult.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMetricManager sets the metric manager.
func WithMetricManager(m metric.Manager) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRegistry sets the evaluator registry.
func WithRegistry(r registry.Registry) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetricRegistry sets the metric runtime registry.
func WithMetricRegistry(r metricregistry.Registry) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvaluationService sets the evaluation service.
func WithEvaluationService(s service.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithUserSimulator sets the simulator used for conversation scenarios.
func WithUserSimulator(sim usersimulation.Simulator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCallbacks sets evaluation callbacks for evaluation service.
func WithCallbacks(c *service.Callbacks) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJudgeRunner injects a judge runner for all LLM judge evaluators.
func WithJudgeRunner(judge runner.Runner) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJudgeRunnerNumSamples sets how many samples to collect from the judge runner.
func WithJudgeRunnerNumSamples(numSamples int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExpectedRunner sets the runner used to generate dynamic expected outputs.
func WithExpectedRunner(r runner.Runner) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNumRuns sets the number of runs.
func WithNumRuns(numRuns int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEvalCaseIDs limits evaluation to the specified eval case IDs.
func WithEvalCaseIDs(evalCaseIDs ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNumRunsParallelEnabled enables or disables parallel execution across evaluation runs.
func WithNumRunsParallelEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvalCaseParallelism sets the maximum number of eval cases processed in parallel.
func WithEvalCaseParallelism(parallelism int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvalCaseParallelInferenceEnabled enables or disables parallel inference across eval cases.
func WithEvalCaseParallelInferenceEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvalCaseParallelEvaluationEnabled enables or disables parallel evaluation across eval cases.
func WithEvalCaseParallelEvaluationEnabled(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRunDetailsEnabled enables or disables per-run inference details in evaluation results.
func WithRunDetailsEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunOptions appends agent.RunOption values that will be applied to every runner.Run call during inference.
func WithRunOptions(opt ...agent.RunOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func (o *options) validate(requireEvalService bool) error { _ = "STUB: not implemented"; return nil }
