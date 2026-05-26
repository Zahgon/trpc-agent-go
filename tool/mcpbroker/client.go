//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mcpbroker

import (
	"context"
	"time"

	mcpcfg "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
	tmcp "trpc.group/trpc-go/trpc-mcp-go"
)

var defaultClientInfo = tmcp.Implementation{
	Name:    "trpc-agent-go",
	Version: "1.0.0",
}

type resolvedTarget struct {
	Name       string
	Origin     string
	TargetType string
	Config     mcpcfg.ConnectionConfig
}

type targetInput struct {
	ServerName string            `json:"server_name,omitempty"`
	URL        string            `json:"url,omitempty"`
	Transport  string            `json:"transport,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

const (
	// PhaseListTools is ClientOptionsRequest.Phase / HeaderInjectRequest.Phase for mcp_list_tools.
	PhaseListTools = "list_tools"
	// PhaseInspectTools is ClientOptionsRequest.Phase / HeaderInjectRequest.Phase for mcp_inspect_tools.
	PhaseInspectTools = "inspect_tools"
	// PhaseCallTool is ClientOptionsRequest.Phase / HeaderInjectRequest.Phase for mcp_call.
	PhaseCallTool = "call_tool"
)

type operationMetadata struct {
	Selector string
	BaseURL  string
	ToolName string
	Phase    string
}

func (b *Broker) buildAdHocConfig(input targetInput) (mcpcfg.ConnectionConfig, string, error) {
	_ = "STUB: not implemented"
	return *new(mcpcfg.ConnectionConfig), "", nil
}

func (b *Broker) sanitizeAdHocHeaders(headers map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Broker) withPreparedHTTPHeaders(
	ctx context.Context,
	target resolvedTarget,
	meta operationMetadata,
) (mcpcfg.ConnectionConfig, error) {
	_ = "STUB: not implemented"
	return *new(mcpcfg.ConnectionConfig), nil
}

func mergeHeaders(base map[string]string, extra map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func canonicalizeHeaders(headers map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func interceptHTTPOperationError(
	ctx context.Context,
	b *Broker,
	target resolvedTarget,
	meta operationMetadata,
	err error,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *Broker) resolveClientOptions(
	ctx context.Context,
	selector string,
	target resolvedTarget,
	meta operationMetadata,
	cfg mcpcfg.ConnectionConfig,
) ([]tmcp.ClientOption, []tmcp.StdioClientOption, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// invokeClientOptionsProvider runs the host-supplied provider with a panic guard so a
// misbehaving hook surfaces as an error wrapping ErrClientOptionsProviderPanicked instead
// of crashing the agent. The provider runs on the hot path of every list/inspect/call
// invocation; recover() here is cheap insurance against host bugs that would otherwise
// unwind through trpc-mcp-go internals.
func invokeClientOptionsProvider(
	ctx context.Context,
	fn ClientOptionsProvider,
	req *ClientOptionsRequest,
) (out *ClientOptions, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filterNilClientOptions drops nil entries so trpc-mcp-go's option application loop never
// invokes a nil ClientOption (which would panic). Hosts often build options conditionally,
// so tolerating sparse slices keeps the contract forgiving.
func filterNilClientOptions(opts []tmcp.ClientOption) []tmcp.ClientOption {
	_ = "STUB: not implemented"
	return nil
}

// filterNilStdioClientOptions mirrors filterNilClientOptions for the stdio transport.
func filterNilStdioClientOptions(opts []tmcp.StdioClientOption) []tmcp.StdioClientOption {
	_ = "STUB: not implemented"
	return nil
}

func createClient(cfg mcpcfg.ConnectionConfig, extraHTTP []tmcp.ClientOption, extraStdio []tmcp.StdioClientOption) (tmcp.Connector, error) {
	_ = "STUB: not implemented"
	return *new(tmcp.Connector), nil
}

func httpHeaderOptions(headers map[string]string) []tmcp.ClientOption {
	_ = "STUB: not implemented"
	return nil
}

func withTimeoutContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func withOneShotClient[T any](
	ctx context.Context,
	cfg mcpcfg.ConnectionConfig,
	extraHTTP []tmcp.ClientOption,
	extraStdio []tmcp.StdioClientOption,
	fn func(context.Context, tmcp.Connector) (T, error),
) (T, error) {
	_ = "STUB: not implemented"

	// Per-call deadline: applied once and shared by Initialize and every MCP
	// RPC issued inside fn. This matches the user-facing contract that
	// ConnectionConfig.Timeout / WithAdHocHTTPTimeout bound the total
	// wall-clock time of a single broker operation.
	return *new(T), nil
}
