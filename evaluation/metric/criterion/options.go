//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package criterion

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/finalresponse"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/tooltrajectory"
)

// options aggregates configurable parts of Criterion.
type options struct {
	// ToolTrajectory sets the default tool trajectory criterion.
	toolTrajectory *tooltrajectory.ToolTrajectoryCriterion
	// finalResponse sets the final response criterion.
	finalResponse *finalresponse.FinalResponseCriterion
	// llmJudge sets the LLM judge criterion.
	llmJudge *llm.LLMCriterion
}

// newOptions creates a Options with the provided options.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option is a function that configures Criterion.
type Option func(*options)

// WithToolTrajectory sets the tool trajectory criterion.
func WithToolTrajectory(toolTrajectory *tooltrajectory.ToolTrajectoryCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFinalResponse sets the final response criterion.
func WithFinalResponse(finalResponse *finalresponse.FinalResponseCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLLMJudge sets the LLM judge criterion.
func WithLLMJudge(llmJudge *llm.LLMCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
