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
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// --- HTTP tool ----------------------------------------------------------
//
// simulatedHTTPTool models any tool that ultimately talks to an external
// HTTP service (MCP SSE/Streamable, webhooks, custom gateway clients). It
// never actually dials the network; instead it builds an *http.Request and
// prints whatever headers the request would carry, so the effect of
// identity.HeadersFromContext is observable.

type httpToolArgs struct {
	Path string `json:"path"`
}

func newHTTPTool() tool.CallableTool { _ = "STUB: not implemented"; return *new(tool.CallableTool) }

// httpToolImpl reads identity headers from context and attaches them to a
// fake request, mimicking what a real mcp.WithHTTPBeforeRequest hook would
// do before sending the request over the wire.
func httpToolImpl(ctx context.Context, args httpToolArgs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatRequest(req *http.Request) string { _ = "STUB: not implemented"; return "" }

// --- Command tool --------------------------------------------------------
//
// simulatedCommandTool models any tool that spawns a child process
// (skill_run, workspace_exec, or a custom bin wrapper). It does not
// actually fork; it just prints the env slice it would hand to exec.Cmd,
// so callers can see how identity.EnvVarsFromContext flows through.

type commandToolArgs struct {
	Command string `json:"command"`
}

func newCommandTool() tool.CallableTool { _ = "STUB: not implemented"; return *new(tool.CallableTool) }

// commandToolImpl reads identity env vars from context and merges them
// with a small static base. A real tool would hand `env` to exec.Cmd.Env,
// which is exactly what codeexecutor.NewEnvInjectingCodeExecutor does for
// skill_run / workspace_exec so each tool implementation does not have to.
func commandToolImpl(ctx context.Context, args commandToolArgs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// joinParts renders a tool result as a single compact string. Tools would
// normally return structured data here; we keep it as a string so the demo
// output prints cleanly.
func joinParts(parts []string) string {
	_ = "STUB: not implemented"
	// Drop zero-length entries defensively.
	return ""
}
