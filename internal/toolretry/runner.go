//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package toolretry executes retryable single tool calls.
package toolretry

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// CallFunc executes one raw tool-call attempt.
type CallFunc func(context.Context, []byte) (any, error)

// ResultErrorFunc classifies whether a raw tool result represents a result-level failure.
type ResultErrorFunc func(any) bool

// TerminalErrorFunc reports whether an error must not be retried.
type TerminalErrorFunc func(error) bool

// ExecuteInput contains the inputs required to execute a retryable tool call.
type ExecuteInput struct {
	ToolName        string
	ToolCallID      string
	Arguments       []byte
	Policy          *tool.RetryPolicy
	Call            CallFunc
	ResultError     ResultErrorFunc
	IsTerminalError TerminalErrorFunc
}

// Result contains the final outcome of the tool-call runner.
type Result struct {
	Result any
	Error  error
}

// Execute runs a tool call with the configured retry policy.
func Execute(ctx context.Context, input ExecuteInput) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

type attemptOutcome struct {
	RawResult   any
	RawError    error
	ResultError bool
}

func resolveMaxAttempts(policy *tool.RetryPolicy) int { _ = "STUB: not implemented"; return 0 }

func contextResult(ctx context.Context) (Result, bool) {
	_ = "STUB: not implemented"
	return *new(Result), false
}

func executeAttempt(ctx context.Context, input ExecuteInput) attemptOutcome {
	_ = "STUB: not implemented"
	return *new(attemptOutcome)
}

func isSuccessfulAttempt(outcome attemptOutcome) bool { _ = "STUB: not implemented"; return false }

func shouldReturnAttempt(
	policy *tool.RetryPolicy,
	attempt int,
	maxAttempts int,
	isTerminalError TerminalErrorFunc,
	outcome attemptOutcome,
) bool {
	_ = "STUB: not implemented"
	return false
}

func finalizeAttempt(outcome attemptOutcome) Result { _ = "STUB: not implemented"; return *new(Result) }

func evaluateRetry(
	ctx context.Context,
	input ExecuteInput,
	policy *tool.RetryPolicy,
	attempt int,
	maxAttempts int,
	outcome attemptOutcome,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func sleepWithPolicy(
	ctx context.Context,
	policy tool.RetryPolicy,
	attempt int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func computeDelay(
	policy tool.RetryPolicy,
	attempt int,
) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func joinPolicyEvaluationError(
	rawErr error,
	policyErr error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func cloneBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }
