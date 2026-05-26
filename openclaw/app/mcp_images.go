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
	mcpContentTypeImage = "image"

	mcpMimeTypePNG  = "image/png"
	mcpMimeTypeJPG  = "image/jpg"
	mcpMimeTypeJPEG = "image/jpeg"
	mcpMimeTypeWebP = "image/webp"
	mcpMimeTypeGIF  = "image/gif"

	mcpImageDetailAuto = "auto"

	mcpImagesUserContent = "MCP tool returned image(s)."
)

type mcpContentItem struct {
	Type     string `json:"type,omitempty"`
	Data     string `json:"data,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

type mcpImage struct {
	Data   []byte
	Format string
}

func mcpImageResultMessages(
	ctx context.Context,
	in *tool.ToolResultMessagesInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func extractMCPImages(ctx context.Context, result any) []mcpImage {
	_ = "STUB: not implemented"
	return nil
}

func unwrapMCPResultContent(result any) any { _ = "STUB: not implemented"; return *new(any) }

func mcpImageFormatFromMime(mime string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
