//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
)

// Tool implementations (simulated file operations).

// readFile simulates reading a file.
func (e *userContextExample) readFile(ctx context.Context, args *fileArgs) (*fileResult, error) {
	_ = "STUB: not implemented"
	// Simulate file reading.
	return nil, nil
}

// writeFile simulates writing to a file.
func (e *userContextExample) writeFile(ctx context.Context, args *writeFileArgs) (*fileResult, error) {
	_ = "STUB: not implemented"
	// Simulate file writing.
	return nil, nil
}

// deleteFile simulates deleting a file.
func (e *userContextExample) deleteFile(ctx context.Context, args *fileArgs) (*fileResult, error) {
	_ = "STUB: not implemented"
	// Simulate file deletion.
	return nil, nil
}

// listFiles simulates listing files.
func (e *userContextExample) listFiles(ctx context.Context, args *listFilesArgs) (*listFilesResult, error) {
	_ = "STUB: not implemented"
	// Simulate file listing.
	return nil, nil
}

// Data structures.

// fileArgs represents arguments for file operations.
type fileArgs struct {
	Filename string `json:"filename" description:"The name of the file"`
}

// writeFileArgs represents arguments for writing a file.
type writeFileArgs struct {
	Filename string `json:"filename" description:"The name of the file"`
	Content  string `json:"content" description:"The content to write"`
}

// listFilesArgs represents arguments for listing files.
type listFilesArgs struct {
	Directory string `json:"directory" description:"The directory to list (default: current)"`
}

// fileResult represents the result of a file operation.
type fileResult struct {
	Filename  string `json:"filename"`
	Operation string `json:"operation"`
	Content   string `json:"content"`
	Success   bool   `json:"success"`
}

// listFilesResult represents the result of listing files.
type listFilesResult struct {
	Directory string   `json:"directory"`
	Files     []string `json:"files"`
	Count     int      `json:"count"`
}
