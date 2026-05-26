//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package hallucination scores sentence-level hallucination judgments.
package hallucination

import (
	"context"
	"regexp"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	positiveVerdict    = "yes"
	labelSupported     = "supported"
	labelUnsupported   = "unsupported"
	labelContradictory = "contradictory"
	labelDisputed      = "disputed"
	labelNotApplicable = "not_applicable"
)

var sentenceBlockRegex = regexp.MustCompile(
	`(?ms)ID:\s*(.*?)\s*` +
		`Reason:\s*(.*?)\s*` +
		`Label:\s*(.*?)\s*` +
		`Verdict:\s*(.*?)(?:\n\s*\n|\z)`,
)

type hallucinationResponseScorer struct {
}

// New returns a response scorer for hallucination judgments.
func New() responsescorer.ResponseScorer {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer)
}

// ScoreBasedOnResponse scores hallucination judgments by averaging sentence verdicts.
func (e *hallucinationResponseScorer) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeLabel(label string) string { _ = "STUB: not implemented"; return "" }

func scoreForLabel(label, verdict string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
