//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package file

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// readMultipleFilesRequest represents the input for the read multiple files
// operation.
type readMultipleFilesRequest struct {
	Patterns      []string `json:"patterns" jsonschema:"description=Glob patterns to read such as *.go or workspace://out/*.txt"`
	CaseSensitive bool     `json:"case_sensitive" jsonschema:"description=Whether glob matching should be case-sensitive"`
}

// readMultipleFilesResponse represents the output from the
// read_multiple_files operation.
type readMultipleFilesResponse struct {
	BaseDirectory string            `json:"base_directory"`
	Files         []*fileReadResult `json:"files"`
	Message       string            `json:"message"`
}

// fileReadResult represents the per-file read result.
type fileReadResult struct {
	FileName string `json:"file_name"`
	Contents string `json:"contents"`
	Message  string `json:"message"`
}

// readMultipleFiles performs the read multiple files operation with support
// for glob patterns.
func (f *fileToolSet) readMultipleFiles(
	ctx context.Context,
	req *readMultipleFilesRequest,
) (*readMultipleFilesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readFiles concurrently reads the given relative path files.
func (f *fileToolSet) readFiles(
	ctx context.Context,
	files []string,
) []*fileReadResult {
	_ = "STUB: not implemented"
	return nil
}

// Capture the per-iteration path to avoid data race on the loop variable.

// readMultipleFilesTool returns a callable tool for reading multiple files.
func (f *fileToolSet) readMultipleFilesTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}
