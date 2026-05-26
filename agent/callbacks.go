//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package agent provides the core agent functionality.
package agent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// ErrorTypeAgentCallbackError is used for errors from agent callbacks (before/after hooks).
const ErrorTypeAgentCallbackError = "agent_callback_error"

const (
	callbackPanicErrFmt = "%s: %v"
	callbackPanicLogFmt = "%s (invocation: %s, agent: %s): " +
		"%v\n%s"

	beforeAgentCallbackPanic = "before agent callback panic"
	afterAgentCallbackPanic  = "after agent callback panic"
)

// BeforeAgentCallback is called before the agent runs.
// Returns (customResponse, error).
// - customResponse: if not nil, this response will be returned to user and agent execution will be skipped.
// - error: if not nil, agent execution will be stopped with this error.
// Deprecated: Use BeforeAgentCallbackStructured instead for better type safety and context passing.
type BeforeAgentCallback = func(ctx context.Context, invocation *Invocation) (*model.Response, error)

// AfterAgentCallback is called after the agent runs.
// Returns (customResponse, error).
// - customResponse: if not nil, this response will be used instead of the actual agent response.
// - error: if not nil, this error will be returned.
// Deprecated: Use AfterAgentCallbackStructured instead for better type safety and context passing.
type AfterAgentCallback = func(ctx context.Context, invocation *Invocation, runErr error) (*model.Response, error)

// BeforeAgentArgs contains all parameters for before agent callback.
type BeforeAgentArgs struct {
	// Invocation is the invocation context.
	Invocation *Invocation
}

// BeforeAgentResult contains the return value for before agent callback.
type BeforeAgentResult struct {
	// Context if not nil, will be used by the framework for subsequent operations.
	Context context.Context
	// CustomResponse if not nil, will skip agent execution and return this response.
	CustomResponse *model.Response
}

// BeforeAgentCallbackStructured is called before the agent runs.
// Returns (result, error).
// - result: contains optional custom response and context for subsequent operations.
//   - CustomResponse: if not nil, this response will be returned to user and agent execution will be skipped.
//   - Context: if not nil, will be used by the framework for subsequent operations.
//
// - error: if not nil, agent execution will be stopped with this error.
type BeforeAgentCallbackStructured = func(
	ctx context.Context,
	args *BeforeAgentArgs,
) (*BeforeAgentResult, error)

// AfterAgentArgs contains all parameters for after agent callback.
type AfterAgentArgs struct {
	// Invocation is the invocation context.
	Invocation *Invocation
	// FullResponseEvent is the final response event from agent execution (may be nil).
	FullResponseEvent *event.Event
	// Error is the error occurred during agent execution (may be nil).
	Error error
}

// AfterAgentResult contains the return value for after agent callback.
type AfterAgentResult struct {
	// Context if not nil, will be used by the framework for subsequent operations.
	Context context.Context
	// CustomResponse if not nil, will replace the original response.
	CustomResponse *model.Response
}

// AfterAgentCallbackStructured is called after the agent runs.
// Returns (result, error).
// - result: contains optional custom response and context for subsequent operations.
//   - CustomResponse: if not nil, this response will be used instead of the actual agent response.
//   - Context: if not nil, will be used by the framework for subsequent operations.
//
// - error: if not nil, this error will be returned.
type AfterAgentCallbackStructured = func(
	ctx context.Context,
	args *AfterAgentArgs,
) (*AfterAgentResult, error)

// Callbacks holds callbacks for agent operations.
// Internally stores the new structured callback types.
type Callbacks struct {
	// BeforeAgent is a list of callbacks called before the agent runs.
	BeforeAgent []BeforeAgentCallbackStructured
	// AfterAgent is a list of callbacks called after the agent runs.
	AfterAgent []AfterAgentCallbackStructured
	// continueOnError controls whether to continue executing callbacks when an error occurs.
	// Default: false (stop on first error)
	continueOnError bool
	// continueOnResponse controls whether to continue executing callbacks when a CustomResponse is returned.
	// Default: false (stop on first CustomResponse)
	continueOnResponse bool
}

// CallbacksOption configures Callbacks behavior.
type CallbacksOption func(*Callbacks)

// WithContinueOnError sets whether to continue executing callbacks when an error occurs.
func WithContinueOnError(continueOnError bool) CallbacksOption {
	_ = "STUB: not implemented"
	return *new(CallbacksOption)
}

// WithContinueOnResponse sets whether to continue executing callbacks when a CustomResponse is returned.
func WithContinueOnResponse(continueOnResponse bool) CallbacksOption {
	_ = "STUB: not implemented"
	return *new(CallbacksOption)
}

// NewCallbacks creates a new Callbacks instance for agent.
func NewCallbacks(opts ...CallbacksOption) *Callbacks { _ = "STUB: not implemented"; return nil }

// Clone returns an independent copy of c, including callback lists
// and execution options.
func (c *Callbacks) Clone() *Callbacks { _ = "STUB: not implemented"; return nil }

// RegisterBeforeAgent registers a before agent callback.
// Supports both old and new callback function signatures.
// Old signatures are automatically wrapped into new signatures.
func (c *Callbacks) RegisterBeforeAgent(cb any) *Callbacks { _ = "STUB: not implemented"; return nil }

// Call old signature

// Return empty result to indicate callback was executed.

// RegisterAfterAgent registers an after agent callback.
// Supports both old and new callback function signatures.
// Old signatures are automatically wrapped into new signatures.
func (c *Callbacks) RegisterAfterAgent(cb any) *Callbacks { _ = "STUB: not implemented"; return nil }

// Call old signature

// Return empty result to indicate callback was executed.

// handleCallbackError processes callback error and returns whether to continue.
func (c *Callbacks) handleCallbackError(err error, firstErr *error) (shouldStop bool) {
	_ = "STUB: not implemented"
	return false
}

// processCallbackResult processes callback result and updates context.
// Returns whether to stop execution immediately.
func (c *Callbacks) processCallbackResult(
	result *BeforeAgentResult,
	ctx *context.Context,
	lastResult **BeforeAgentResult,
) (shouldStop bool) {
	_ = "STUB: not implemented"
	return false
}

func recoverAgentCallbackPanic(
	ctx context.Context,
	stage string,
	inv *Invocation,
	errp *error,
) {
	_ = "STUB: not implemented"
	return
}

func (c *Callbacks) runBeforeAgentCallback(
	ctx context.Context,
	cb BeforeAgentCallbackStructured,
	args *BeforeAgentArgs,
) (result *BeforeAgentResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// finalizeBeforeAgentResult determines the final return value for before agent callbacks.
func (c *Callbacks) finalizeBeforeAgentResult(
	lastResult *BeforeAgentResult,
	firstErr error,
) (*BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunBeforeAgent runs all before agent callbacks in order.
// This method uses the new structured callback interface.
// If a callback returns a non-nil Context in the result, it will be used for subsequent callbacks.
func (c *Callbacks) RunBeforeAgent(
	ctx context.Context,
	args *BeforeAgentArgs,
) (*BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processAfterCallbackResult processes after callback result and updates context.
// Returns whether to stop execution immediately.
func (c *Callbacks) processAfterCallbackResult(
	result *AfterAgentResult,
	ctx *context.Context,
	lastResult **AfterAgentResult,
) (shouldStop bool) {
	_ = "STUB: not implemented"
	return false
}

// finalizeAfterAgentResult determines the final return value for after agent callbacks.
func (c *Callbacks) finalizeAfterAgentResult(
	lastResult *AfterAgentResult,
	firstErr error,
) (*AfterAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Callbacks) runAfterAgentCallback(
	ctx context.Context,
	cb AfterAgentCallbackStructured,
	args *AfterAgentArgs,
) (result *AfterAgentResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterAgent runs all after agent callbacks in order.
// This method uses the new structured callback interface.
// If a callback returns a non-nil Context in the result, it will be used for subsequent callbacks.
func (c *Callbacks) RunAfterAgent(
	ctx context.Context,
	args *AfterAgentArgs,
) (*AfterAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
