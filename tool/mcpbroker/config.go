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
	mcpcfg "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

type transportKind string

const (
	transportStdio      transportKind = "stdio"
	transportSSE        transportKind = "sse"
	transportStreamable transportKind = "streamable"
)

func normalizeNamedServer(name string, cfg mcpcfg.ConnectionConfig, origin string) (namedServer, error) {
	_ = "STUB: not implemented"
	return *new(namedServer), nil
}

func normalizeConnectionConfig(cfg mcpcfg.ConnectionConfig, adHoc bool) (mcpcfg.ConnectionConfig, transportKind, error) {
	_ = "STUB: not implemented"
	return *new(mcpcfg.ConnectionConfig), *new(transportKind), nil
}

func normalizeTransport(raw string, hasCommand, hasURL, adHoc bool) (string, transportKind, error) {
	_ = "STUB: not implemented"
	return "", *new(transportKind), nil
}

// cloneConnectionConfig returns a defensive copy of cfg. Scalar fields are duplicated by Go
// value semantics; the Headers map and Args slice are deep-cloned via cloneStringMap and
// cloneStringSlice. Mutations to the returned config (including its Headers / Args) therefore
// do not propagate back to the source. Callers that hand the clone to external callbacks (such
// as ClientOptionsProvider) rely on this contract.
func cloneConnectionConfig(cfg mcpcfg.ConnectionConfig) mcpcfg.ConnectionConfig {
	_ = "STUB: not implemented"
	return *new(mcpcfg.ConnectionConfig)
}

func cloneStringMap(input map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func cloneStringSlice(input []string) []string { _ = "STUB: not implemented"; return nil }
