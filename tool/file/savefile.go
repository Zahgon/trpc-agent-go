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

// saveFileRequest represents the input for the save file operation.
type saveFileRequest struct {
	// FileName is a path relative to base_directory.
	FileName string `json:"file_name" jsonschema:"description=Relative file path under base_directory to write"`
	// Contents is the file content to write.
	Contents string `json:"contents" jsonschema:"description=Text content to write into the file"`
	// Overwrite controls whether an existing file is replaced.
	Overwrite bool `json:"overwrite" jsonschema:"description=Whether to replace the file if it already exists"`
}

// saveFileResponse represents the output from the save file operation.
type saveFileResponse struct {
	BaseDirectory string `json:"base_directory"`
	FileName      string `json:"file_name"`
	Message       string `json:"message"`
}

// saveFile performs the save file operation.
func (f *fileToolSet) saveFile(
	_ context.Context,
	req *saveFileRequest,
) (*saveFileResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resolve and validate the file path.

// Create parent directories if they don't exist.

// Check if file exists and overwrite is disabled.

// Write the file.

// saveFileTool returns a callable tool for saving file.
func (f *fileToolSet) saveFileTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}
