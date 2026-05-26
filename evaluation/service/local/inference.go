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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/service"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/service/internal/inference"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Inference runs the agent for the requested eval cases and returns the inference results for each case.
func (s *local) Inference(ctx context.Context, req *service.InferenceRequest, opt ...service.Option) (results []*service.InferenceResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) runBeforeInferenceSetCallbacks(ctx context.Context, callbacks *service.Callbacks, req *service.InferenceRequest) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *local) runAfterInferenceSetCallbacks(
	ctx context.Context,
	callbacks *service.Callbacks,
	req *service.InferenceRequest,
	results []*service.InferenceResult,
	err error,
	startTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) runBeforeInferenceCaseCallbacks(
	ctx context.Context,
	callbacks *service.Callbacks,
	req *service.InferenceRequest,
	evalCaseID string,
	sessionID string,
) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (s *local) runAfterInferenceCaseCallbacks(
	ctx context.Context,
	callbacks *service.Callbacks,
	req *service.InferenceRequest,
	evalCaseID string,
	result *service.InferenceResult,
	err error,
	startTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) validateInferenceRequest(req *service.InferenceRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) loadInferenceEvalCases(ctx context.Context, req *service.InferenceRequest, mgr evalset.Manager) ([]*evalset.EvalCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) inferEvalCases(ctx context.Context, req *service.InferenceRequest, evalCases []*evalset.EvalCase, opts *service.Options) ([]*service.InferenceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) inferEvalCasesSerial(ctx context.Context, req *service.InferenceRequest, evalCases []*evalset.EvalCase, opts *service.Options) ([]*service.InferenceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) inferEvalCasesParallel(ctx context.Context, req *service.InferenceRequest, evalCases []*evalset.EvalCase, opts *service.Options) ([]*service.InferenceResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *local) inferenceEvalCase(ctx context.Context, req *service.InferenceRequest, evalCase *evalset.EvalCase, opts *service.Options) (result *service.InferenceResult) {
	_ = "STUB: not implemented"
	return nil
}

func (s *local) inferCaseConversations(
	ctx context.Context,
	evalCase *evalset.EvalCase,
	sessionID string,
	runOptions []agent.RunOption,
	opts *service.Options,
) (*inference.Result, []*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *local) inferTraceConversation(
	ctx context.Context,
	evalCase *evalset.EvalCase,
	sessionID string,
	opts *service.Options,
) (*inference.Result, []*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *local) inferStaticConversation(
	ctx context.Context,
	evalCase *evalset.EvalCase,
	sessionID string,
	runOptions []agent.RunOption,
	opts *service.Options,
) (*inference.Result, []*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *local) inferScenarioConversation(
	ctx context.Context,
	evalCase *evalset.EvalCase,
	sessionID string,
	runOptions []agent.RunOption,
	opts *service.Options,
) (*inference.Result, []*evalset.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func expectedRunnerSessionID(sessionID string) string { _ = "STUB: not implemented"; return "" }

func seedMessagesFromPointers(messages []*model.Message) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newInferenceResult(appName, evalSetID, sessionID string, evalCase *evalset.EvalCase) *service.InferenceResult {
	_ = "STUB: not implemented"
	return nil
}

func newFailedInferenceResult(result *service.InferenceResult, err error) *service.InferenceResult {
	_ = "STUB: not implemented"
	return nil
}
