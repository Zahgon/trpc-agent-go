//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package codeexec provides a code execution tool that allows LLM to execute code.
package codeexec

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Option configures the code execution tool.
type Option func(*config)

type config struct {
	name        string
	description string
	languages   []string
}

// WithName sets the tool name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDescription sets the tool description.
func WithDescription(desc string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLanguages sets the supported languages (default: python, bash).
func WithLanguages(langs ...string) Option {
	_ = "STUB: not implemented"
	return *

	// Defensive copy to avoid caller mutation.
	new(Option)
}

func defaultConfig() config { _ = "STUB: not implemented"; return *new(config) }

func applyOptions(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// NewTool creates a new code execution tool with the given CodeExecutor and options.
//
// This follows the common pattern in this repo: return a `tool.CallableTool` interface and
// keep the concrete implementation unexported.
func NewTool(exec codeexecutor.CodeExecutor, opts ...Option) tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}

type executeCodeTool struct {
	executor codeexecutor.CodeExecutor
	cfg      config
}

// Declaration returns the tool's declaration.
func (t *executeCodeTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// unmarshalCodeBlocks flexibly decodes code_blocks from JSON, handling common LLM
// quirks: the value may be a normal array, a single object (instead of an array),
// or a double-encoded JSON string containing either of the above.
func unmarshalCodeBlocks(raw json.RawMessage) ([]codeexecutor.CodeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the LLM double-encoded the array as a JSON string, unwrap and re-parse.

// Single object — wrap into a slice.

// Call executes the code and returns the result.
func (t *executeCodeTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Best-effort validation. We return it as structured tool output (instead of Go error)
// so the model can correct itself.

func (t *executeCodeTool) isSupportedLanguage(language string) bool {
	_ = "STUB: not implemented"
	return false
}
