//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"context"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/memoryfile"
)

const (
	memoryToolFileName = "MEMORY.md"

	memoryToolReadFileFS       = "fs_read_file"
	memoryToolSaveFileFS       = "fs_save_file"
	memoryToolReplaceContentFS = "fs_replace_content"
)

var errMemorySaveFileExists = errors.New(
	"memory file exists and overwrite=false",
)

type memoryToolTarget struct {
	AppName string
	UserID  string
	Path    string
}

type memoryReadFileRequest struct {
	FileName  string `json:"file_name"`
	StartLine *int   `json:"start_line,omitempty"`
	NumLines  *int   `json:"num_lines,omitempty"`
}

type memoryReadFileResponse struct {
	BaseDirectory string `json:"base_directory"`
	FileName      string `json:"file_name"`
	Contents      string `json:"contents"`
	Message       string `json:"message"`
}

type memorySaveFileRequest struct {
	FileName  string `json:"file_name"`
	Contents  string `json:"contents"`
	Overwrite bool   `json:"overwrite"`
}

type memorySaveFileResponse struct {
	BaseDirectory string `json:"base_directory"`
	FileName      string `json:"file_name"`
	Message       string `json:"message"`
}

type memoryReplaceContentRequest struct {
	FileName        string `json:"file_name"`
	OldString       string `json:"old_string"`
	NewString       string `json:"new_string"`
	NumReplacements int    `json:"num_replacements,omitempty"`
}

type memoryReplaceContentResponse struct {
	BaseDirectory string `json:"base_directory"`
	FileName      string `json:"file_name"`
	Message       string `json:"message"`
}

func registerMemoryFileToolCallback(
	callbacks *tool.Callbacks,
	store *memoryfile.Store,
	stateDir string,
) {
	_ = "STUB: not implemented"
	return
}

func newMemoryFileToolCallback(
	store *memoryfile.Store,
	stateDir string,
) tool.BeforeToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.BeforeToolCallbackStructured)
}

func normalizeMemoryToolName(name string) string { _ = "STUB: not implemented"; return "" }

func memoryToolTargetFromContext(
	ctx context.Context,
	store *memoryfile.Store,
) (memoryToolTarget, bool, error) {
	_ = "STUB: not implemented"
	return *new(memoryToolTarget), false, nil
}

func handleMemoryReadFileTool(
	target memoryToolTarget,
	baseDir string,
	args []byte,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleMemorySaveFileTool(
	ctx context.Context,
	store *memoryfile.Store,
	stateDir string,
	target memoryToolTarget,
	args []byte,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleMemoryReplaceContentTool(
	ctx context.Context,
	store *memoryfile.Store,
	stateDir string,
	target memoryToolTarget,
	args []byte,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func memoryToolResult(result any) *tool.BeforeToolResult { _ = "STUB: not implemented"; return nil }

func isMemoryFileAlias(fileName string) bool { _ = "STUB: not implemented"; return false }

func nextMemorySaveContents(
	existing string,
	incoming string,
	overwrite bool,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func appendMemorySnippet(existing string, incomingTrimmed string) string {
	_ = "STUB: not implemented"
	return ""
}

func looksLikeMemoryAppendSnippet(text string) bool { _ = "STUB: not implemented"; return false }

func sliceMemoryTextByLines(
	text string,
	startLine *int,
	numLines *int,
) (string, int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return "", 0, 0, 0, false, nil
}
