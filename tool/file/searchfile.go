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

// searchFileRequest represents the input for the search file operation.
type searchFileRequest struct {
	// Path is a relative directory under base_directory.
	Path string `json:"path" jsonschema:"description=Relative directory path under base_directory or workspace:// directory ref; empty means the base directory"`
	// Pattern is a glob to match file names.
	Pattern string `json:"pattern" jsonschema:"description=Glob pattern to match files or folders such as *.go or **/*.md"`
	// CaseSensitive controls glob case matching.
	CaseSensitive bool `json:"case_sensitive" jsonschema:"description=Whether glob matching should be case-sensitive"`
}

// searchFileResponse represents the output from the search file operation.
type searchFileResponse struct {
	BaseDirectory string   `json:"base_directory"`
	Path          string   `json:"path"`
	Pattern       string   `json:"pattern"`
	Files         []string `json:"files"`
	Folders       []string `json:"folders"`
	Message       string   `json:"message"`
}

// searchFile performs the search file operation.
func (f *fileToolSet) searchFile(
	ctx context.Context,
	req *searchFileRequest,
) (*searchFileResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate pattern

// Resolve and validate the target path.

// Check if the target path exists.

// Check if the target path is a file.

// Find files matching the pattern.

// Separate files and folders.

// Skip entries that can't be stat.

// searchFileTool returns a callable tool for searching file.
func (f *fileToolSet) searchFileTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}
