//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package service

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/registry"
	metricregistry "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/usersimulation"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Options holds the options for the evaluation service.
type Options struct {
	EvalSetManager                    evalset.Manager                  // EvalSetManager is used to store and retrieve eval set.
	EvalResultManager                 evalresult.Manager               // EvalResultManager is used to store and retrieve eval results.
	Registry                          registry.Registry                // Registry is used to store and retrieve evaluator.
	MetricRegistry                    metricregistry.Registry          // MetricRegistry resolves runtime metric extensions.
	SessionIDSupplier                 func(ctx context.Context) string // SessionIDSupplier is used to generate session IDs.
	ExpectedRunner                    runner.Runner                    // ExpectedRunner is used to generate dynamic expected outputs.
	UserSimulator                     usersimulation.Simulator         // UserSimulator drives conversationScenario inference.
	Callbacks                         *Callbacks                       // Callbacks holds evaluation callbacks.
	RunOptions                        []agent.RunOption                // RunOptions configures runner.Run calls during inference.
	EvalCaseParallelism               int                              // EvalCaseParallelism controls concurrent eval case processing.
	EvalCaseParallelInferenceEnabled  bool                             // EvalCaseParallelInferenceEnabled toggles parallel inference across eval cases.
	EvalCaseParallelEvaluationEnabled bool                             // EvalCaseParallelEvaluationEnabled toggles parallel evaluation across eval cases.
}

// Option defines a function type for configuring the evaluation service.
type Option func(*Options)

// NewOptions creates a new Options with the default values.
func NewOptions(opt ...Option) *Options { _ = "STUB: not implemented"; return nil }

// WithEvalSetManager sets the eval set manager.
// InMemory eval set manager is used by default.
func WithEvalSetManager(m evalset.Manager) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEvalResultManager sets the eval result manager.
// InMemory eval result manager is used by default.
func WithEvalResultManager(m evalresult.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRegistry sets the evaluator registry.
// Default evaluator registry is used by default.
func WithRegistry(r registry.Registry) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetricRegistry sets the metric runtime registry.
func WithMetricRegistry(r metricregistry.Registry) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionIDSupplier sets the function used to generate session IDs.
// UUID generator is used by default.
func WithSessionIDSupplier(s func(ctx context.Context) string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExpectedRunner sets the runner used to generate dynamic expected outputs.
func WithExpectedRunner(r runner.Runner) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserSimulator sets the simulator used for conversation scenarios.
func WithUserSimulator(sim usersimulation.Simulator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCallbacks sets the evaluation lifecycle callbacks.
func WithCallbacks(c *Callbacks) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunOptions appends agent.RunOption values that will be applied to every runner.Run call during inference.
func WithRunOptions(opt ...agent.RunOption) Option { _ = "STUB: not implemented"; return *new(Option) }

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
