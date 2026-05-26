//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package echotool registers a tiny tool provider plugin.
//
// It is intended as a reference implementation for writing custom tool
// providers. The registered tool simply echoes one string argument.
package echotool

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

const (
	pluginType = "echotool"

	schemaTypeObject = "object"
	schemaTypeString = "string"

	argText = "text"
)

func init() {
	if err := registry.RegisterToolProvider(pluginType, newTools); err != nil {
		panic(err)
	}
}

type providerCfg struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func newTools(
	_ registry.ToolProviderDeps,
	spec registry.PluginSpec,
) ([]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type echoTool struct {
	name string
	desc string
}

func (t echoTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t echoTool) Call(_ context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
