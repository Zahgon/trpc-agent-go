//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package callback

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/service"
)

func wrapCallbackError(point string, idx int, name string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func callCallbackWithRecovery[Args any, Result any, CallbackFn ~func(context.Context, *Args) (*Result, error)](
	ctx context.Context,
	point string,
	idx int,
	name string,
	callback CallbackFn,
	args *Args,
) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func callCallbackWithRecoveryInto[Args any, Result any, CallbackFn ~func(context.Context, *Args) (*Result, error)](
	ctx context.Context,
	point string,
	idx int,
	name string,
	callback CallbackFn,
	args *Args,
	resultp **Result,
	errp *error,
) {
	_ = "STUB: not implemented"
	return
}

func runCallbacks[Args any, Result any, CallbackFn ~func(context.Context, *Args) (*Result, error)](
	ctx *context.Context,
	callbacks []service.NamedCallback[CallbackFn],
	args *Args,
	point string,
	getContext func(*Result) context.Context,
) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunBeforeInferenceSet runs all before inference set callbacks in order.
func RunBeforeInferenceSet(ctx context.Context, callbacks *service.Callbacks, args *service.BeforeInferenceSetArgs) (*service.BeforeInferenceSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterInferenceSet runs all after inference set callbacks in order.
func RunAfterInferenceSet(ctx context.Context, callbacks *service.Callbacks, args *service.AfterInferenceSetArgs) (*service.AfterInferenceSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunBeforeInferenceCase runs all before inference case callbacks in order.
func RunBeforeInferenceCase(ctx context.Context, callbacks *service.Callbacks, args *service.BeforeInferenceCaseArgs) (*service.BeforeInferenceCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterInferenceCase runs all after inference case callbacks in order.
func RunAfterInferenceCase(ctx context.Context, callbacks *service.Callbacks, args *service.AfterInferenceCaseArgs) (*service.AfterInferenceCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunBeforeEvaluateSet runs all before evaluate set callbacks in order.
func RunBeforeEvaluateSet(ctx context.Context, callbacks *service.Callbacks, args *service.BeforeEvaluateSetArgs) (*service.BeforeEvaluateSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterEvaluateSet runs all after evaluate set callbacks in order.
func RunAfterEvaluateSet(ctx context.Context, callbacks *service.Callbacks, args *service.AfterEvaluateSetArgs) (*service.AfterEvaluateSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunBeforeEvaluateCase runs all before evaluate case callbacks in order.
func RunBeforeEvaluateCase(ctx context.Context, callbacks *service.Callbacks, args *service.BeforeEvaluateCaseArgs) (*service.BeforeEvaluateCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterEvaluateCase runs all after evaluate case callbacks in order.
func RunAfterEvaluateCase(ctx context.Context, callbacks *service.Callbacks, args *service.AfterEvaluateCaseArgs) (*service.AfterEvaluateCaseResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
