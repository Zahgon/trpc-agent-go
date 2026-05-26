//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeinterpreter

// SandboxError is the generic error returned by the SDK for unexpected
// server responses.
type SandboxError struct {
	Message    string
	StatusCode int
}

func (e *SandboxError) Error() string { _ = "STUB: not implemented"; return "" }

// NotFoundError is returned when a resource (context, sandbox, file) is missing.
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// TimeoutError is returned when a request or execution times out.
type TimeoutError struct {
	Message string
}

func (e *TimeoutError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidArgumentError is returned when input parameters are invalid
// (e.g. providing both `context` and `language`).
type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string { _ = "STUB: not implemented"; return "" }

// AuthenticationError is returned when the supplied API key is invalid or
// missing.
type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string { _ = "STUB: not implemented"; return "" }

// RateLimitError is returned when the caller has exceeded the API's rate limit.
type RateLimitError struct {
	Message string
}

func (e *RateLimitError) Error() string { _ = "STUB: not implemented"; return "" }

// formatRequestTimeoutError wraps an error with a friendlier timeout message.
func formatRequestTimeoutError() error { _ = "STUB: not implemented"; return nil }

// formatExecutionTimeoutError wraps an error with a friendlier timeout message
// for code execution.
func formatExecutionTimeoutError() error { _ = "STUB: not implemented"; return nil }
