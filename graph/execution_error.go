//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graph

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// StateKeyExecutionErrors is the default state key used to store
	// collected execution errors.
	StateKeyExecutionErrors = "execution_errors"

	// ExecutionErrorEventType is the default custom event type used when
	// emitting fallback state for fatal execution errors.
	ExecutionErrorEventType = "execution_error"
)

// ExecutionErrorSeverity describes whether a recorded error was recovered.
type ExecutionErrorSeverity string

const (
	// ExecutionErrorSeverityRecoverable marks errors that were recorded and
	// converted into a replacement node result.
	ExecutionErrorSeverityRecoverable ExecutionErrorSeverity = "recoverable"

	// ExecutionErrorSeverityFatal marks errors that still terminated the graph.
	ExecutionErrorSeverityFatal ExecutionErrorSeverity = "fatal"
)

// ExecutionError captures structured business-visible error details.
type ExecutionError struct {
	Severity   ExecutionErrorSeverity `json:"severity"`
	NodeID     string                 `json:"nodeId,omitempty"`
	NodeName   string                 `json:"nodeName,omitempty"`
	NodeType   NodeType               `json:"nodeType,omitempty"`
	StepNumber int                    `json:"stepNumber,omitempty"`
	Timestamp  time.Time              `json:"timestamp"`
	Error      *model.ResponseError   `json:"error,omitempty"`
}

// ExecutionErrorPolicy describes how a node error should be handled.
type ExecutionErrorPolicy struct {
	Recover bool

	// Replacement optionally overrides the replacement result used when
	// Recover is true. Prefer State or *Command so the collector can merge the
	// execution_errors update automatically.
	Replacement any

	// ResponseError optionally overrides the structured error fields written
	// into the collected record.
	ResponseError *model.ResponseError
}

// RecoverableExecutionError marks an error as recoverable for the default
// collector policy.
type RecoverableExecutionError interface {
	error
	Recoverable() bool
}

type recoverableExecutionError struct {
	cause error
}

func (e recoverableExecutionError) Error() string { _ = "STUB: not implemented"; return "" }

func (e recoverableExecutionError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e recoverableExecutionError) Recoverable() bool {
	_ = "STUB: not implemented"

	// ExecutionErrorPolicyFunc decides how a node error should be recorded.
	return false
}

type ExecutionErrorPolicyFunc func(
	ctx context.Context,
	callbackCtx *NodeCallbackContext,
	state State,
	result any,
	err error,
) ExecutionErrorPolicy

// DefaultExecutionErrorPolicy is the framework's default recovery policy.
func DefaultExecutionErrorPolicy(
	_ context.Context,
	_ *NodeCallbackContext,
	_ State,
	_ any,
	err error,
) ExecutionErrorPolicy {
	_ = "STUB: not implemented"
	return *new(ExecutionErrorPolicy)
}

// ExecutionErrorCollectorOption configures an ExecutionErrorCollector.
type ExecutionErrorCollectorOption func(*ExecutionErrorCollector)

// ExecutionErrorCollector provides reusable graph error collection helpers.
type ExecutionErrorCollector struct {
	stateKey  string
	eventType string
	policy    ExecutionErrorPolicyFunc
}

// NewExecutionErrorCollector creates a new collector.
func NewExecutionErrorCollector(
	opts ...ExecutionErrorCollectorOption,
) *ExecutionErrorCollector {
	_ = "STUB: not implemented"
	return nil
}

// WithExecutionErrorStateKey overrides the state key used by the collector.
func WithExecutionErrorStateKey(
	key string,
) ExecutionErrorCollectorOption {
	_ = "STUB: not implemented"
	return *new(ExecutionErrorCollectorOption)
}

// WithExecutionErrorEventType overrides the custom event type used on fatal
// fallback state emission.
func WithExecutionErrorEventType(
	eventType string,
) ExecutionErrorCollectorOption {
	_ = "STUB: not implemented"
	return *new(ExecutionErrorCollectorOption)
}

// WithExecutionErrorPolicy sets a custom error handling policy.
func WithExecutionErrorPolicy(
	policy ExecutionErrorPolicyFunc,
) ExecutionErrorCollectorOption {
	_ = "STUB: not implemented"
	return *new(ExecutionErrorCollectorOption)
}

// WithRecoverableExecutionErrors extends the default recovery policy with an
// additional recoverable-error predicate.
func WithRecoverableExecutionErrors(
	shouldRecover func(error) bool,
) ExecutionErrorCollectorOption {
	_ = "STUB: not implemented"
	return *new(ExecutionErrorCollectorOption)
}

// MarkRecoverable wraps err so the default collector policy treats it as
// recoverable.
func MarkRecoverable(err error) error { _ = "STUB: not implemented"; return nil }

// NewRecoverableError returns a recoverable error with the provided message.
func NewRecoverableError(message string) error { _ = "STUB: not implemented"; return nil }

// IsRecoverableExecutionError reports whether err matches the default
// recoverable-error contract.
func IsRecoverableExecutionError(err error) bool { _ = "STUB: not implemented"; return false }

// StateKey returns the state key used by the collector.
func (c *ExecutionErrorCollector) StateKey() string {
	_ = "STUB: not implemented"

	// StateField returns a StateField suitable for collecting execution errors.
	return ""
}

func (c *ExecutionErrorCollector) StateField() StateField {
	_ = "STUB: not implemented"
	return *new(StateField)
}

// AddField registers the collector's state field onto a schema.
func (c *ExecutionErrorCollector) AddField(
	schema *StateSchema,
) *StateSchema {
	_ = "STUB: not implemented"
	return nil
}

// NodeCallbacks returns callbacks that collect execution errors on node
// failure.
func (c *ExecutionErrorCollector) NodeCallbacks() *NodeCallbacks {
	_ = "STUB: not implemented"
	return nil
}

// SubgraphStateUpdate extracts collected execution errors from a child agent
// result so the parent graph can merge them into its own state.
func (c *ExecutionErrorCollector) SubgraphStateUpdate(
	result SubgraphResult,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

// SubgraphOutputMapper returns a mapper that merges child execution errors
// into the parent graph state.
func (c *ExecutionErrorCollector) SubgraphOutputMapper() SubgraphOutputMapper {
	_ = "STUB: not implemented"
	return *new(SubgraphOutputMapper)
}

// NewExecutionError creates a structured record from a node callback context.
func NewExecutionError(
	callbackCtx *NodeCallbackContext,
	err error,
	severity ExecutionErrorSeverity,
) ExecutionError {
	_ = "STUB: not implemented"
	return *new(ExecutionError)
}

// ExecutionErrorSliceReducer appends execution error slices.
func ExecutionErrorSliceReducer(existing, update any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// DecodeExecutionErrors unmarshals a serialized execution error slice.
func DecodeExecutionErrors(
	raw []byte,
) ([]ExecutionError, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecutionErrorsFromStateDelta extracts execution errors from an event state
// delta using the provided state key.
func ExecutionErrorsFromStateDelta(
	stateDelta map[string][]byte,
	key string,
) ([]ExecutionError, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ExecutionErrorCollector) afterNode(
	ctx context.Context,
	callbackCtx *NodeCallbackContext,
	state State,
	result any,
	nodeErr error,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func mergeExecutionErrorReplacement(
	replacement any,
	update State,
) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func mergeStateForExecutionError(
	dst State,
	update State,
) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func cloneExecutionErrors(
	executionErrors []ExecutionError,
) []ExecutionError {
	_ = "STUB: not implemented"
	return nil
}

func cloneResponseError(
	err *model.ResponseError,
) *model.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func executionErrorMessage(
	record ExecutionError,
) string {
	_ = "STUB: not implemented"
	return ""
}
