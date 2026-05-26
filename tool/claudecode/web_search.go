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

	"golang.org/x/net/html"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type codeSearchBackend interface {
	search(context.Context, webSearchInput) ([]webSearchHit, error)
}

type duckDuckGoSearchBackend struct {
	client    *http.Client
	baseURL   string
	userAgent string
	size      int
	offset    int
}

type googleSearchBackend struct {
	client  *http.Client
	options *WebSearchOptions
}

func newWebSearchTool(options *WebSearchOptions) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func newSearchBackend(options *WebSearchOptions) (codeSearchBackend, error) {
	_ = "STUB: not implemented"
	return *new(codeSearchBackend), nil
}

func webSearchDescription() string { _ = "STUB: not implemented"; return "" }

func (b *duckDuckGoSearchBackend) search(
	ctx context.Context,
	in webSearchInput,
) ([]webSearchHit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDuckDuckGoHTML(body []byte, in webSearchInput, offset int, limit int) []webSearchHit {
	_ = "STUB: not implemented"
	return nil
}

func normalizeDuckDuckGoResultURL(rawURL string) string { _ = "STUB: not implemented"; return "" }

func htmlHasClass(node *html.Node, className string) bool { _ = "STUB: not implemented"; return false }

func htmlNodeText(node *html.Node) string { _ = "STUB: not implemented"; return "" }

func (b *googleSearchBackend) search(
	ctx context.Context,
	in webSearchInput,
) ([]webSearchHit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applySearchWindow(hits []webSearchHit, offset int, limit int) []webSearchHit {
	_ = "STUB: not implemented"
	return nil
}

func webSearchSize(options *WebSearchOptions) int { _ = "STUB: not implemented"; return 0 }

func webSearchOffset(options *WebSearchOptions) int { _ = "STUB: not implemented"; return 0 }
