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

// readFileRequest represents the input for the read file operation.
type readFileRequest struct {
	FileName  string `json:"file_name" jsonschema:"description=Relative file path under base_directory or workspace:// or artifact:// file ref to read"`
	StartLine *int   `json:"start_line,omitempty" jsonschema:"description=Optional 1-based start line to begin reading from"`
	NumLines  *int   `json:"num_lines,omitempty" jsonschema:"description=Optional maximum number of lines to return"`
}

// readFileResponse represents the output from the read file operation.
type readFileResponse struct {
	BaseDirectory string `json:"base_directory"`
	FileName      string `json:"file_name"`
	Contents      string `json:"contents"`
	Message       string `json:"message"`
}

// readFile performs the read file operation.
func (f *fileToolSet) readFile(
	ctx context.Context,
	req *readFileRequest,
) (*readFileResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate the start line and number of lines.

func validateReadFileRequest(req *readFileRequest) error { _ = "STUB: not implemented"; return nil }

const (
	errNotTextFile     = "file is not a UTF-8 text file"
	errNotTextFileTmpl = "file is not a UTF-8 text file (mime: %s)"
)

func validateTextString(content string, mimeType string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTextBytes(data []byte, mimeType string) error { _ = "STUB: not implemented"; return nil }

func notTextFileErr(mimeType string) error { _ = "STUB: not implemented"; return nil }

func (f *fileToolSet) readFileFromRef(
	ctx context.Context,
	req *readFileRequest,
	rsp *readFileResponse,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *fileToolSet) readFileFromDiskOrCache(
	ctx context.Context,
	req *readFileRequest,
	rsp *readFileResponse,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *fileToolSet) readFileFromCache(
	ctx context.Context,
	req *readFileRequest,
	rsp *readFileResponse,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *fileToolSet) sliceReadFile(
	req *readFileRequest,
	content string,
) (string, int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, 0, false, nil
}

func sliceTextByLines(
	text string,
	startLine *int,
	numLines *int,
) (string, int, int, int, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, 0, nil
}

// readFileTool returns a callable tool for reading file.
func (f *fileToolSet) readFileTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}
