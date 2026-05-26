//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package langfuse fetches text prompts from the Langfuse prompt management
// API and maps them to [prompt.Text] with double-curly variable syntax.
package langfuse

import (
	"context"
	"net/http"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/prompt"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/langfuse/config"
)

// ClientOption configures a [Client].
type ClientOption func(*clientConfig)

type clientConfig struct {
	httpClient *http.Client
}

const defaultHTTPTimeout = 30 * time.Second

// Client fetches text prompts from the Langfuse REST API.
type Client struct {
	baseURL    string
	publicKey  string
	secretKey  string
	httpClient *http.Client
}

// NewClient creates a Langfuse prompt client from a shared ConnectionConfig.
func NewClient(cfg config.ConnectionConfig, opts ...ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

// FetchOption configures a prompt fetch request.
type FetchOption func(*fetchConfig)

type fetchConfig struct {
	label   string // "" means unset
	version int    // 0 means unset
}

// WithLabel fetches the prompt version carrying the given label.
// Langfuse resolves prompts by either label or version, so this clears any
// previously selected version. The default label is "production" when neither
// label nor version is specified.
func WithLabel(label string) FetchOption { _ = "STUB: not implemented"; return *new(FetchOption) }

// WithVersion fetches a specific prompt version number.
// Langfuse resolves prompts by either version or label, so this clears any
// previously selected label.
func WithVersion(version int) FetchOption { _ = "STUB: not implemented"; return *new(FetchOption) }

// TextPromptResult holds a fetched text prompt and its associated metadata.
type TextPromptResult struct {
	Text    prompt.Text
	Config  map[string]any
	Version int
	Labels  []string
}

// FetchTextPrompt fetches a text prompt by name from the Langfuse API.
// Prompt names containing folder paths are URL-escaped as a single path segment.
// Only prompts with type "text" are accepted; chat prompts return an error.
func (c *Client) FetchTextPrompt(ctx context.Context, name string, opts ...FetchOption) (TextPromptResult, error) {
	_ = "STUB: not implemented"
	return *new(TextPromptResult), nil
}

// apiPromptResponse mirrors the relevant fields of the Langfuse
// GET /api/public/v2/prompts/{name} response.
type apiPromptResponse struct {
	Name    string         `json:"name"`
	Version int            `json:"version"`
	Type    string         `json:"type"`
	Prompt  any            `json:"prompt"`
	Config  map[string]any `json:"config"`
	Labels  []string       `json:"labels"`
}

// --- Source factory with caching ---

// SourceOption configures a cached [prompt.Source] created via [Client.TextPromptSource].
type SourceOption func(*sourceConfig)

type sourceConfig struct {
	cacheTTL time.Duration
}

// WithCacheTTL sets the time-to-live for cached prompt results.
// The default TTL is 60 seconds.
func WithCacheTTL(ttl time.Duration) SourceOption {
	_ = "STUB: not implemented"
	return *new(SourceOption)
}

// TextPromptSource returns a [prompt.Source] that fetches the named text prompt
// with the given fetch options. Results are cached with a default TTL of 60s.
// On fetch failure, a valid cached value is returned if available. Caller
// cancellation or deadline expiry is returned directly and does not use stale
// cache.
func (c *Client) TextPromptSource(name string, opts ...FetchOption) prompt.Source {
	_ = "STUB: not implemented"
	return *new(prompt.Source)
}

// TextPromptSourceWithOptions is like [Client.TextPromptSource] but accepts
// additional [SourceOption] values to configure caching behavior.
func (c *Client) TextPromptSourceWithOptions(name string, fetchOpts []FetchOption, sourceOpts ...SourceOption) prompt.Source {
	_ = "STUB: not implemented"
	return *new(prompt.Source)
}

type cachedSource struct {
	client   *Client
	name     string
	fetchOpt []FetchOption
	ttl      time.Duration

	mu        sync.RWMutex
	cached    prompt.Text
	fetchedAt time.Time
	valid     bool
}

func (s *cachedSource) FetchPrompt(ctx context.Context) (prompt.Text, error) {
	_ = "STUB: not implemented"
	return *new(prompt.Text), nil
}
