//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	webSearchToolName = "web_search"

	ddgFormQueryKey    = "q"
	ddgFormContentType = "application/x-www-form-urlencoded"
	ddgHTMLSearchURL   = "https://html.duckduckgo.com/html/"
	ddgHTTPTimeout     = 30 * time.Second

	defaultMaxResults = 5
	maxSearchResults  = 20

	httpPrefix  = "http://"
	httpsPrefix = "https://"

	duckDuckGoHost = "duckduckgo.com"
)

const ddgUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/120.0.0.0 Safari/537.36"

const (
	ddgLinkPattern = `class="result__a"[^>]*href="([^"]+)"[^>]*>` +
		`([^<]+)</a>`
	ddgSnippetPattern = `class="result__snippet"[^>]*>([^<]+)</a>`
	ddgRedirectPath   = "/l/"
	ddgRedirectKey    = "uddg"
)

type webSearchRequest struct {
	Query string `json:"query" jsonschema:"description=Search query"`
	Limit int    `json:"limit,omitempty" jsonschema:"description=Max results"`
}

type webSearchResponse struct {
	Query   string            `json:"query"`
	Results []webSearchResult `json:"results"`
	Summary string            `json:"summary"`
	Error   string            `json:"error,omitempty"`
}

type webSearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Snippet string `json:"snippet"`
}

func newWebSearchTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func webSearch(
	ctx context.Context,
	req webSearchRequest,
) (webSearchResponse, error) {
	_ = "STUB: not implemented"
	return *new(webSearchResponse), nil
}

func parseDDGHTML(html string, limit int) []webSearchResult { _ = "STUB: not implemented"; return nil }

func sanitizeSearchLimit(limit int) int { _ = "STUB: not implemented"; return 0 }

func normalizeSearchURL(raw string) string { _ = "STUB: not implemented"; return "" }

func isHTTPURL(value string) bool { _ = "STUB: not implemented"; return false }

func isSearchResultURL(value string) bool { _ = "STUB: not implemented"; return false }

func cleanHTML(value string) string { _ = "STUB: not implemented"; return "" }
