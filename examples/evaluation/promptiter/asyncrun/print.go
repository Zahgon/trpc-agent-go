//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
)

func printSummary(
	result *promptiterengine.RunResult,
	dataDir string,
	outputDir string,
	initialInstruction string,
	targetSurfaceID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func initialValidationScore(result *promptiterengine.RunResult) float64 {
	_ = "STUB: not implemented"
	return 0
}

func finalAcceptedValidationScore(result *promptiterengine.RunResult) float64 {
	_ = "STUB: not implemented"
	return 0
}

func evaluationResultScore(result *promptiterengine.EvaluationResult) float64 {
	_ = "STUB: not implemented"
	return 0
}
