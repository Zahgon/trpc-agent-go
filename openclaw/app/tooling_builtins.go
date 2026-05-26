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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
	"trpc.group/trpc-go/trpc-agent-go/tool/mcp"

	openapitool "trpc.group/trpc-go/trpc-agent-go/tool/openapi"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

const (
	toolProviderBrowser    = "browser"
	toolProviderDuckDuckGo = "duckduckgo"
	toolProviderWebFetch   = "webfetch_http"

	toolSetProviderMCP     = "mcp"
	toolSetProviderFile    = "file"
	toolSetProviderOpenAPI = "openapi"
	toolSetProviderGoogle  = "google"
	toolSetProviderWiki    = "wikipedia"
	toolSetProviderArxiv   = "arxivsearch"
	toolSetProviderEmail   = "email"

	defaultHTTPTimeout = 30 * time.Second

	envGoogleAPIKey   = "GOOGLE_API_KEY"
	envGoogleEngineID = "GOOGLE_SEARCH_ENGINE_ID"

	mcpTransportStdio      = "stdio"
	mcpTransportSSE        = "sse"
	mcpTransportStreamable = "streamable"
)

func init() {
	must(registry.RegisterToolProvider(
		toolProviderBrowser,
		newBrowserTools,
	))
	must(registry.RegisterToolProvider(
		toolProviderDuckDuckGo,
		newDuckDuckGoTools,
	))
	must(registry.RegisterToolProvider(
		toolProviderWebFetch,
		newHTTPWebFetchTools,
	))

	must(registry.RegisterToolSetProvider(
		toolSetProviderMCP,
		newMCPToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderFile,
		newFileToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderOpenAPI,
		newOpenAPIToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderGoogle,
		newGoogleToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderWiki,
		newWikipediaToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderArxiv,
		newArxivToolSet,
	))
	must(registry.RegisterToolSetProvider(
		toolSetProviderEmail,
		newEmailToolSet,
	))
}

type httpToolConfig struct {
	BaseURL   string        `yaml:"base_url,omitempty"`
	UserAgent string        `yaml:"user_agent,omitempty"`
	Timeout   time.Duration `yaml:"timeout,omitempty"`
}

func newBrowserTools(
	_ registry.ToolProviderDeps,
	spec registry.PluginSpec,
) ([]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDuckDuckGoTools(
	_ registry.ToolProviderDeps,
	spec registry.PluginSpec,
) ([]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type httpWebFetchConfig struct {
	AllowedDomains []string      `yaml:"allowed_domains,omitempty"`
	BlockedDomains []string      `yaml:"blocked_domains,omitempty"`
	AllowAll       bool          `yaml:"allow_all_domains,omitempty"`
	Timeout        time.Duration `yaml:"timeout,omitempty"`

	MaxContentLength      int `yaml:"max_content_length,omitempty"`
	MaxTotalContentLength int `yaml:"max_total_content_length,omitempty"`
}

func newHTTPWebFetchTools(
	_ registry.ToolProviderDeps,
	spec registry.PluginSpec,
) ([]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mcpFilterConfig struct {
	Mode  string   `yaml:"mode,omitempty"`
	Names []string `yaml:"names,omitempty"`
}

type mcpReconnectConfig struct {
	Enabled     bool `yaml:"enabled,omitempty"`
	MaxAttempts int  `yaml:"max_attempts,omitempty"`
}

type mcpToolSetConfig struct {
	Transport string            `yaml:"transport,omitempty"`
	ServerURL string            `yaml:"server_url,omitempty"`
	Headers   map[string]string `yaml:"headers,omitempty"`
	Command   string            `yaml:"command,omitempty"`
	Args      []string          `yaml:"args,omitempty"`
	Timeout   time.Duration     `yaml:"timeout,omitempty"`

	ToolFilter *mcpFilterConfig    `yaml:"tool_filter,omitempty"`
	Reconnect  *mcpReconnectConfig `yaml:"reconnect,omitempty"`
}

func newMCPToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

func validateMCPConnection(cfg mcp.ConnectionConfig) error { _ = "STUB: not implemented"; return nil }

func buildMCPToolFilter(cfg *mcpFilterConfig) (tool.FilterFunc, error) {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc), nil
}

type fileToolSetConfig struct {
	BaseDir       string `yaml:"base_dir,omitempty"`
	ReadOnly      *bool  `yaml:"read_only,omitempty"`
	EnableSave    *bool  `yaml:"enable_save,omitempty"`
	EnableReplace *bool  `yaml:"enable_replace,omitempty"`

	EnableRead          *bool `yaml:"enable_read,omitempty"`
	EnableReadMultiple  *bool `yaml:"enable_read_multiple,omitempty"`
	EnableList          *bool `yaml:"enable_list,omitempty"`
	EnableSearchFile    *bool `yaml:"enable_search_file,omitempty"`
	EnableSearchContent *bool `yaml:"enable_search_content,omitempty"`

	MaxFileSize int64 `yaml:"max_file_size,omitempty"`
}

func newFileToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

type openAPISpecConfig struct {
	File   string `yaml:"file,omitempty"`
	URL    string `yaml:"url,omitempty"`
	Inline string `yaml:"inline,omitempty"`
}

type openAPIToolSetConfig struct {
	Spec              *openAPISpecConfig `yaml:"spec,omitempty"`
	AllowExternalRefs bool               `yaml:"allow_external_refs,omitempty"`
	UserAgent         string             `yaml:"user_agent,omitempty"`
	Timeout           time.Duration      `yaml:"timeout,omitempty"`
}

func newOpenAPIToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

func openAPILoader(
	cfg openAPISpecConfig,
	allowExternalRefs bool,
) (openapitool.Loader, error) {
	_ = "STUB: not implemented"
	return *new(openapitool.Loader), nil
}

type googleToolSetConfig struct {
	APIKey   string        `yaml:"api_key,omitempty"`
	EngineID string        `yaml:"engine_id,omitempty"`
	BaseURL  string        `yaml:"base_url,omitempty"`
	Size     int           `yaml:"size,omitempty"`
	Offset   int           `yaml:"offset,omitempty"`
	Lang     string        `yaml:"lang,omitempty"`
	Timeout  time.Duration `yaml:"timeout,omitempty"`
}

func newGoogleToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

type wikipediaToolSetConfig struct {
	Language   string        `yaml:"language,omitempty"`
	MaxResults int           `yaml:"max_results,omitempty"`
	UserAgent  string        `yaml:"user_agent,omitempty"`
	Timeout    time.Duration `yaml:"timeout,omitempty"`
}

func newWikipediaToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

type arxivToolSetConfig struct {
	BaseURL      string        `yaml:"base_url,omitempty"`
	PageSize     int           `yaml:"page_size,omitempty"`
	DelaySeconds time.Duration `yaml:"delay_seconds,omitempty"`
	NumRetries   int           `yaml:"num_retries,omitempty"`
}

func newArxivToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

func newEmailToolSet(
	_ registry.ToolSetProviderDeps,
	spec registry.PluginSpec,
) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

type toolSetNameOverride struct {
	name string
	tool tool.ToolSet
}

func (t toolSetNameOverride) Tools(ctx context.Context) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func (t toolSetNameOverride) Close() error { _ = "STUB: not implemented"; return nil }

func (t toolSetNameOverride) Name() string { _ = "STUB: not implemented"; return "" }

func overrideToolSetName(ts tool.ToolSet, name string) tool.ToolSet {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet)
}
