//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package search

import (
	"context"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi"
)

// searchRequest represents the input for the Google search tool.
type searchRequest struct {
	Query  string `json:"query" jsonschema:"description=The search query to execute on Google"`
	Size   int    `json:"size" jsonschema:"description=The number of results to return"`
	Offset int    `json:"offset" jsonschema:"description=The offset of the results to return"`
	Lang   string `json:"lang" jsonschema:"description=The language of the results to return(en/ja/zh-CN/etc)"`
}

// search executes a search query on Google and returns the results.
func (t *ToolSet) search(ctx context.Context, req searchRequest) (result, error) {
	_ = "STUB: not implemented"
	return *new(result), nil
}

func getQuery(reqs []*customsearch.SearchQueriesRequest) string {
	_ = "STUB: not implemented"
	return ""
}

func getDescFromPageMap(pageMap googleapi.RawMessage) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// result represents the search results from Google Search.
type result struct {
	Query string                  `json:"query,omitempty"`
	Items []*simplifiedSearchItem `json:"items"`
}

// simplifiedSearchItem represents a simplified search item.
type simplifiedSearchItem struct {
	Link    string `json:"link"`
	Title   string `json:"title,omitempty"`
	Snippet string `json:"snippet,omitempty"`
	Desc    string `json:"desc,omitempty"`
}
