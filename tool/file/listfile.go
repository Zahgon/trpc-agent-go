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

// listFileRequest represents the input for the list file operation.
type listFileRequest struct {
	// Path is a relative directory under base_directory.
	Path string `json:"path" jsonschema:"description=Relative directory path under base_directory or workspace:// directory ref; empty means the base directory"`

	// WithSize returns the size of the files.
	WithSize bool `json:"with_size" jsonschema:"description=Whether to include file sizes in files_with_size"`
}

// listFileResponse represents the output from the list file operation.
type listFileResponse struct {
	BaseDirectory string   `json:"base_directory"`
	Path          string   `json:"path"`
	Files         []string `json:"files"`
	Folders       []string `json:"folders"`
	Message       string   `json:"message"`

	FilesWithSize []fileInfo `json:"files_with_size"`
}

type fileInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// listFile performs the list file operation.
func (f *fileToolSet) listFile(
	ctx context.Context,
	req *listFileRequest,
) (*listFileResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resolve the target path.

// Check if the target path exists.

// If the target is a file, return information about that file.

// If the target is a directory, list its contents.

// Collect files and folders.

// Create a summary message.

func listWorkspaceEntries(
	ctx context.Context,
	dir string,
) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listFileTool returns a callable tool for listing file.
func (f *fileToolSet) listFileTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}
