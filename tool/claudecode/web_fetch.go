//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package claudecode

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func newWebFetchTool(options WebFetchOptions) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func fetchURL(
	ctx context.Context,
	client *http.Client,
	rawURL string,
	options WebFetchOptions,
) (string, int, string, []byte, string, error) {
	_ = "STUB: not implemented"
	return "", 0, "", nil, "", nil
}

func resolveRedirectURL(baseURL string, location string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func processFetchedContent(
	ctx context.Context,
	options WebFetchOptions,
	in webFetchInput,
	content string,
	contentType string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func webFetchDescription() string { _ = "STUB: not implemented"; return "" }

func trimFetchResult(content string, limit int) string { _ = "STUB: not implemented"; return "" }
