//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package codeinterpreter provides a CodeInterpreter implementation for E2B.
package codeinterpreter

import (
	"context"
	"time"
)

// RunCodeOpts holds options for Sandbox.RunCode.
//
// Either Language or Context may be supplied — not both. When both are empty
// the default Python context is used.
type RunCodeOpts struct {
	// Language to use. Must not be combined with Context.
	Language RunCodeLanguage
	// Context (pre-created kernel) to run the code in.
	Context *Context
	// OnStdout is called for every stdout chunk.
	OnStdout OnStdoutFunc
	// OnStderr is called for every stderr chunk.
	OnStderr OnStderrFunc
	// OnResult is called for every Result (display call or final result).
	OnResult OnResultFunc
	// OnError is called when the kernel reports an error for this cell.
	OnError OnErrorFunc
	// Envs are extra environment variables exposed to the running code.
	Envs map[string]string
	// Timeout is the maximum execution time for this cell (default: 300s).
	// Pass -1 to disable.
	Timeout time.Duration
	// RequestTimeout is the HTTP-level request timeout.
	RequestTimeout time.Duration
}

// CreateCodeContextOpts holds options for Sandbox.CreateCodeContext.
type CreateCodeContextOpts struct {
	// Cwd is the working directory for the context (default /home/user).
	Cwd string
	// Language of the new context (default python).
	Language RunCodeLanguage
	// RequestTimeout overrides the default HTTP request timeout.
	RequestTimeout time.Duration
}

// RunCode executes the supplied code in the sandbox and returns the full
// Execution result.
//
// Streaming output is forwarded to the callbacks on opts in real time.
func (s *Sandbox) RunCode(ctx context.Context, code string, opts *RunCodeOpts) (*Execution, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build request body

// Build request with its own context to honour Timeout/RequestTimeout.

// parseOutputLine decodes a single NDJSON line emitted by /execute and
// dispatches it to the correct handler.
func parseOutputLine(execution *Execution, line string, opts *RunCodeOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore un-parseable lines — they shouldn't happen but we don't
// want to blow up a long running execution over them.

// The "type" / "is_main_result" fields are not relevant as keys on the
// resulting Result object; newResultFromRaw handles the split.

func getInt64(m map[string]any, key string) int64 { _ = "STUB: not implemented"; return 0 }

// CreateCodeContext creates a fresh kernel in which subsequent code can be run.
func (s *Sandbox) CreateCodeContext(ctx context.Context, opts *CreateCodeContextOpts) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListCodeContexts lists the contexts currently available in the sandbox.
func (s *Sandbox) ListCodeContexts(ctx context.Context) ([]*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RestartCodeContext restarts the given context. The parameter can either be a
// *Context or a context-id string.
func (s *Sandbox) RestartCodeContext(ctx context.Context, c any) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveCodeContext removes the context. The parameter can either be a
// *Context or a context-id string.
func (s *Sandbox) RemoveCodeContext(ctx context.Context, c any) error {
	_ = "STUB: not implemented"
	return nil
}

// doContextRequest is a helper for creating a Context via POST /contexts.
func (s *Sandbox) doContextRequest(ctx context.Context, method, path string, body map[string]any, reqTimeout time.Duration) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// jupyterRequest performs an HTTP request against the sandbox jupyter server
// and returns the raw body, status code and any transport-level error.
func (s *Sandbox) jupyterRequest(ctx context.Context, method, path string, body any, reqTimeout time.Duration) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// contextID returns the string ID from either a *Context or a plain string.
func contextID(c any) (string, error) { _ = "STUB: not implemented"; return "", nil }
