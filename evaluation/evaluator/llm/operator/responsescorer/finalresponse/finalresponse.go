//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package finalresponse converts judge feedback into validity scores for final responses.
package finalresponse

import (
	"context"
	"regexp"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const labelValid string = "valid" // labelValid marks a valid agent response.

// finalResponseBlockRegex extracts the reasoning and validity label from the judge response.
var finalResponseBlockRegex = regexp.MustCompile(
	`(?ms)reasoning:\s*(.*?)\s*` + // 1: reasoning text
		`is_the_agent_response_valid:\s*(.*?)\s*$`, // 2: validity label
)

type finalResponseResponseScorer struct {
}

// New returns a response scorer for final responses.
func New() responsescorer.ResponseScorer {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer)
}

// ScoreBasedOnResponse converts judge feedback to a numeric score.
func (e *finalResponseResponseScorer) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	_ *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractReasoningAndLabel parses judge output in text form.
func extractReasoningAndLabel(content string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
