//
//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package codeexecutor provides an interface and utilities for executing
// code blocks and running programs in workspaces.
package codeexecutor

import (
	"context"
)

// CodeExecutor executes code blocks via a friendly front-door API.
type CodeExecutor interface {
	// ExecuteCode executes the code blocks provided in the input and
	// returns the result.
	ExecuteCode(context.Context, CodeExecutionInput) (CodeExecutionResult, error)
	// CodeBlockDelimiter returns the delimiters used for code blocks.
	CodeBlockDelimiter() CodeBlockDelimiter
}

// CodeExecutionInput is the input for code execution.
type CodeExecutionInput struct {
	CodeBlocks  []CodeBlock `json:"code_blocks"`
	ExecutionID string      `json:"execution_id,omitempty"`
}

// CodeExecutionResult is the result of code execution including files.
type CodeExecutionResult struct {
	Output      string `json:"output"`
	OutputFiles []File `json:"output_files,omitempty"`
}

// String formats a human-readable result.
func (r CodeExecutionResult) String() string { _ = "STUB: not implemented"; return "" }

// File represents a file generated during code execution.
type File struct {
	Name      string `json:"name"`
	Content   string `json:"content,omitempty"`
	MIMEType  string `json:"mime_type"`
	SizeBytes int64  `json:"size_bytes,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`
}

// CodeBlock represents a single block of code to be executed.
type CodeBlock struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// CodeBlockDelimiter defines the start and end delimiters for code blocks.
type CodeBlockDelimiter struct {
	Start string
	End   string
}

// ExtractCodeBlock extracts fenced code blocks using the given delimiter.
func ExtractCodeBlock(
	input string,
	delimiter CodeBlockDelimiter,
) []CodeBlock {
	_ = "STUB: not implemented"
	return nil
}
