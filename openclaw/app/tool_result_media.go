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

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolResultMediaLineFile = "MEDIA:"
	toolResultMediaLineDir  = "MEDIA_DIR:"

	maxToolResultImages = 6

	maxToolResultImageBytes int64 = 12 << 20

	toolResultImagesTraceKind = "tool.result.images"
)

const (
	toolResultSingleImageText = "Generated image attached for " +
		"direct inspection: %s."
	toolResultMultiImageText = "Generated images attached for " +
		"direct inspection: %s."
)

type toolResultMediaPayload struct {
	Output     string   `json:"output,omitempty"`
	MediaFiles []string `json:"media_files,omitempty"`
	MediaDirs  []string `json:"media_dirs,omitempty"`
}

type toolResultImage struct {
	Name   string
	Data   []byte
	Format string
}

func openClawToolResultMessages(
	ctx context.Context,
	in *tool.ToolResultMessagesInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func toolResultImageMessages(
	ctx context.Context,
	in *tool.ToolResultMessagesInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func loadToolResultImages(result any) []toolResultImage { _ = "STUB: not implemented"; return nil }

func collectToolResultImagePaths(result any) []string { _ = "STUB: not implemented"; return nil }

func parseToolResultMediaPayload(
	result any,
) (toolResultMediaPayload, bool) {
	_ = "STUB: not implemented"
	return *new(toolResultMediaPayload), false
}

func appendToolResultImagePath(
	out []string,
	seen map[string]struct{},
	raw string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendToolResultDirImages(
	out []string,
	seen map[string]struct{},
	root string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendUniqueToolResultPath(
	out []string,
	seen map[string]struct{},
	path string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func toolResultOutputPaths(output string) []string { _ = "STUB: not implemented"; return nil }

func toolResultPathFromLine(line string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func resolveToolResultPath(raw string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func toolResultImageFormat(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func toolResultImageMessageText(images []toolResultImage) string {
	_ = "STUB: not implemented"
	return ""
}

func recordToolResultImages(
	ctx context.Context,
	toolName string,
	images []toolResultImage,
) {
	_ = "STUB: not implemented"
	return
}
