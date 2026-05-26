//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package search provides a Google Search tool set.
package search

import (
	"context"

	"google.golang.org/api/customsearch/v1"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type config struct {
	apiKey   string
	engineID string
	baseURL  string
	size     int
	offset   int
	lang     string
}

// Option is a function that configures the Google search tool.
type Option func(*config)

// WithAPIKey sets the API key for the Google search tool.
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEngineID sets the search engine ID for the Google search tool.
func WithEngineID(id string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets the base URL for the Google search tool.
func WithBaseURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSize sets the size for the Google search tool.
// Number of search results to return.
// Default is 5.
func WithSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOffset sets the offset for the Google search tool.
// The index of the first result to return.
// Default is 0.
func WithOffset(offset int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLanguage sets the language for the Google search tool.
// Default is "en".
func WithLanguage(lang string) Option { _ = "STUB: not implemented"; return *new(Option) }

// ToolSet represents a Google search tool
type ToolSet struct {
	srv   *customsearch.Service
	name  string
	cfg   *config
	tools []tool.Tool
}

// Tools implements the ToolSet interface.
func (t *ToolSet) Tools(ctx context.Context) []tool.Tool {
	_ = "STUB: not implemented"

	// Close implements the ToolSet interface.
	return nil
}

func (t *ToolSet) Close() error {
	_ = "STUB: not implemented"

	// Name implements the ToolSet interface.
	return nil
}

func (t *ToolSet) Name() string {
	_ = "STUB: not implemented"

	// NewToolSet creates a new Google Search with the provided options.
	return ""
}

func NewToolSet(ctx context.Context, opts ...Option) (*ToolSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
