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

	"trpc.group/trpc-go/trpc-agent-go/tool"
	tmcp "trpc.group/trpc-go/trpc-mcp-go"
)

type listServersInput struct{}

type listServersOutput struct {
	Servers []listServersServer `json:"servers"`
}

type listServersServer struct {
	Name        string `json:"name"`
	Transport   string `json:"transport"`
	Description string `json:"description,omitempty"`
}

type listToolsInput struct {
	Selector  string            `json:"selector"`
	Transport string            `json:"transport,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
}

type listToolsOutput struct {
	Tools []listedTool `json:"tools"`
}

type listedTool struct {
	Name            string `json:"name"`
	Selector        string `json:"selector,omitempty"`
	Signature       string `json:"signature,omitempty"`
	Description     string `json:"description,omitempty"`
	HasOutputSchema bool   `json:"has_output_schema"`
}

type inspectToolsInput struct {
	Selector            string            `json:"selector"`
	Tools               []string          `json:"tools"`
	Transport           string            `json:"transport,omitempty"`
	Headers             map[string]string `json:"headers,omitempty"`
	IncludeOutputSchema bool              `json:"include_output_schema,omitempty"`
}

type inspectToolsOutput struct {
	Tools []inspectedTool `json:"tools"`
}

type inspectedTool struct {
	Name            string         `json:"name"`
	Selector        string         `json:"selector,omitempty"`
	Description     string         `json:"description,omitempty"`
	HasOutputSchema bool           `json:"has_output_schema"`
	InputSchema     map[string]any `json:"input_schema,omitempty"`
	OutputSchema    map[string]any `json:"output_schema,omitempty"`
}

type callToolInput struct {
	Selector  string            `json:"selector"`
	Transport string            `json:"transport,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Arguments map[string]any    `json:"arguments,omitempty"`
}

type callToolOutput struct {
	Content           []any `json:"content,omitempty"`
	StructuredContent any   `json:"structured_content,omitempty"`
	IsError           bool  `json:"is_error,omitempty"`
}

func newBrokerTools(b *Broker) []tool.Tool { _ = "STUB: not implemented"; return nil }

func (b *Broker) listTools(ctx context.Context, input listToolsInput) (listToolsOutput, error) {
	_ = "STUB: not implemented"
	return *new(listToolsOutput), nil
}

func (b *Broker) inspectTools(ctx context.Context, input inspectToolsInput) (inspectToolsOutput, error) {
	_ = "STUB: not implemented"
	return *new(inspectToolsOutput), nil
}

func (b *Broker) callTool(ctx context.Context, input callToolInput) (callToolOutput, error) {
	_ = "STUB: not implemented"
	return *new(callToolOutput), nil
}

func validateCallToolArguments(
	ctx context.Context,
	client tmcp.Connector,
	toolName string,
	arguments map[string]any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func emptyObjectSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func listToolsInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func inspectToolsInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func callToolInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func listServersOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func listToolsOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func inspectToolsOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func callToolOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func joinToolSelector(prefix string, toolName string) string { _ = "STUB: not implemented"; return "" }

func renderToolSignature(mcpTool tmcp.Tool) string { _ = "STUB: not implemented"; return "" }

func schemaTypeName(schema any) string { _ = "STUB: not implemented"; return "" }

func firstSchemaType(value any) (string, bool) { _ = "STUB: not implemented"; return "", false }

func schemaToMap(schema any) map[string]any { _ = "STUB: not implemented"; return nil }

func selectToolsForInspection(mcpTools []tmcp.Tool, requested []string) ([]tmcp.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
