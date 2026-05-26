//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package local

import (
	"context"
	"sync"

	"github.com/panjf2000/ants/v2"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/service"
)

type evalCaseInferenceParam struct {
	idx      int
	ctx      context.Context
	req      *service.InferenceRequest
	evalCase *evalset.EvalCase
	opts     *service.Options
	svc      *local
	results  []*service.InferenceResult
	wg       *sync.WaitGroup
}

func (p *evalCaseInferenceParam) reset() { _ = "STUB: not implemented"; return }

var evalCaseInferenceParamPool = &sync.Pool{
	New: func() any { return new(evalCaseInferenceParam) },
}

func createEvalCaseInferencePool(size int) (*ants.PoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) ensureEvalCaseInferencePool(size int) (*ants.PoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type evalCaseEvaluationParam struct {
	idx             int
	ctx             context.Context
	req             *service.EvaluateRequest
	inferenceResult *service.InferenceResult
	opts            *service.Options
	svc             *local
	results         []*evalresult.EvalCaseResult
	errs            []error
	wg              *sync.WaitGroup
}

func (p *evalCaseEvaluationParam) reset() { _ = "STUB: not implemented"; return }

var evalCaseEvaluationParamPool = &sync.Pool{
	New: func() any { return new(evalCaseEvaluationParam) },
}

func createEvalCaseEvaluationPool(size int) (*ants.PoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) ensureEvalCaseEvaluationPool(size int) (*ants.PoolWithFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
