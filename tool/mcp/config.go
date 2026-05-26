//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mcp

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

// filterMode defines how the filter should behave.
type filterMode string

// transport specifies the transport method: "stdio", "sse", "streamable_http".
type transport string

const (
	// transportStdio is the stdio transport.
	transportStdio transport = "stdio"
	// transportSSE is the Server-Sent Events transport.
	transportSSE transport = "sse"
	// transportStreamable is the streamable HTTP transport.
	transportStreamable transport = "streamable"

	// FilterModeInclude specifies that only listed tools should be included.
	FilterModeInclude filterMode = "include"
	// FilterModeExclude specifies that listed tools should be excluded.
	FilterModeExclude filterMode = "exclude"
)

// Default configurations.
const (
	// defaultMaxReconnectAttempts is the default maximum number of session reconnection attempts.
	defaultMaxReconnectAttempts = 3
	// minReconnectAttempts is the minimum allowed reconnection attempts.
	minReconnectAttempts = 1
	// maxReconnectAttemptsLimit is the maximum allowed reconnection attempts.
	maxReconnectAttemptsLimit = 10
)

var (
	defaultClientInfo = mcp.Implementation{
		Name:    "trpc-agent-go",
		Version: "1.0.0",
	}
)

// ConnectionConfig defines the configuration for connecting to an MCP server.
type ConnectionConfig struct {
	// Transport specifies the transport method: "stdio", "sse", "streamable".
	Transport string `json:"transport"`

	// Streamable/SSE configuration.
	ServerURL string            `json:"server_url,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`

	// STDIO configuration.
	Command string   `json:"command,omitempty"`
	Args    []string `json:"args,omitempty"`

	// Common configuration.
	Timeout time.Duration `json:"timeout,omitempty"`

	// Description provides a capability summary of this MCP server.
	// When set, mcpbroker exposes it in mcp_list_servers so the model can
	// decide which server to explore without listing every tool first.
	Description string `json:"description,omitempty"`

	// Advanced configuration.
	ClientInfo mcp.Implementation `json:"client_info,omitempty"`
}

// SessionReconnectConfig defines configuration for automatic session reconnection.
type SessionReconnectConfig struct {
	// EnableAutoReconnect enables automatic session reconnection when session expires.
	// Default: false
	EnableAutoReconnect bool `json:"enable_auto_reconnect"`
	// MaxReconnectAttempts specifies the maximum number of reconnection attempts
	// within the session lifecycle to prevent infinite reconnection loops.
	// Valid range: 1-10, default: 3
	MaxReconnectAttempts int `json:"max_reconnect_attempts"`
}

// toolSetConfig holds internal configuration for ToolSet.
type toolSetConfig struct {
	connectionConfig       ConnectionConfig
	toolFilterFunc         tool.FilterFunc         // Tool filter function.
	mcpOptions             []mcp.ClientOption      // MCP client options.
	sessionReconnectConfig *SessionReconnectConfig // Session reconnection configuration.
	name                   string                  // ToolSet name for identification and conflict resolution.
}

// ToolSetOption is a function type for configuring ToolSet.
type ToolSetOption func(*toolSetConfig)

// WithToolFilterFunc configures tool filtering.
//
// Example:
//
//	import "trpc.group/trpc-go/trpc-agent-go/tool"
//
//	// Include only specific tools
//	mcp.WithToolFilterFunc(tool.NewIncludeToolNamesFilter("echo", "add"))
//
//	// Exclude specific tools
//	mcp.WithToolFilterFunc(tool.NewExcludeToolNamesFilter("deprecated_tool"))
func WithToolFilterFunc(filterFunc tool.FilterFunc) ToolSetOption {
	_ = "STUB: not implemented"
	return *new(ToolSetOption)
}

// WithMCPOptions sets additional MCP client options.
// This can be used to pass options to the underlying MCP client.
func WithMCPOptions(options ...mcp.ClientOption) ToolSetOption {
	_ = "STUB: not implemented"
	return *new(ToolSetOption)
}

// WithSessionReconnect enables automatic session reconnection with specified max attempts.
// maxAttempts: maximum number of reconnection attempts (valid range: 1-10, recommended: 3)
// When enabled, the session manager will automatically attempt to recreate
// the MCP session when it receives session-expired errors from the transport layer.
func WithSessionReconnect(maxAttempts int) ToolSetOption {
	_ = "STUB: not implemented"
	return *new(ToolSetOption)
}

// Clamp to valid range

// WithSessionReconnectConfig enables automatic session reconnection with custom configuration.
// This provides full control over reconnection behavior for advanced use cases.
func WithSessionReconnectConfig(config SessionReconnectConfig) ToolSetOption {
	_ = "STUB: not implemented"
	return *new(ToolSetOption)
}

// Always enable auto reconnect when using this option

// Clamp to valid range

// WithName sets the name of the MCP toolset for identification and conflict resolution.
// This name will be used when implementing tool name prefixing to avoid conflicts
// between tools from different toolsets.
func WithName(name string) ToolSetOption { _ = "STUB: not implemented"; return *new(ToolSetOption) }

// Per-request dynamic HTTP headers (e.g. user-specific Authorization tokens)
// are supported directly through the upstream MCP client. Pass an
// mcp.WithHTTPBeforeRequest hook to WithMCPOptions and read whatever
// request-scoped values you need from the context:
//
//	import (
//	    tmcp "trpc.group/trpc-go/trpc-mcp-go"
//	    toolmcp "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
//	)
//
//	ts := toolmcp.NewMCPToolSet(cfg,
//	    toolmcp.WithMCPOptions(tmcp.WithHTTPBeforeRequest(
//	        func(ctx context.Context, req *http.Request) error {
//	            for k, v := range headersFromContext(ctx) {
//	                req.Header.Set(k, v)
//	            }
//	            return nil
//	        },
//	    )),
//	)
//
// When this ToolSet is wired into llmagent via WithToolSets, enable
// llmagent.WithRefreshToolSetsOnRun(true) if initialize/listTools must also
// observe request-scoped headers. Otherwise only tools/call is guaranteed to
// use the current run context.

// validateTransport validates the transport string and returns the internal transport type.
func validateTransport(t string) (transport, error) {
	_ = "STUB: not implemented"
	return *new(transport), nil
}
