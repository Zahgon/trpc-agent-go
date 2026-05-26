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
	"net/http"

	a2a "trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	defaultA2AUserIDHeader = "X-User-ID"

	a2aVersionKey = "version"
)

// A2ASurface provides the HTTP handler and routes served by OpenClaw A2A.
type A2ASurface struct {
	Handler       http.Handler
	BasePath      string
	URL           string
	AgentCardPath string
	UserIDHeader  string
}

func newA2ASurface(
	ag agent.Agent,
	procRunner runner.Runner,
	opts runOptions,
) (A2ASurface, error) {
	_ = "STUB: not implemented"
	return *new(A2ASurface), nil
}

func buildOpenClawA2ACard(
	ag agent.Agent,
	opts runOptions,
	host string,
) (a2a.AgentCard, error) {
	_ = "STUB: not implemented"
	return *new(a2a.AgentCard), nil
}

func extractA2ABasePath(host string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func buildRuntimeHTTPHandler(
	gatewayHandler http.Handler,
	surface A2ASurface,
) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func mountA2ASurface(
	mux *http.ServeMux,
	surface A2ASurface,
) error {
	_ = "STUB: not implemented"
	return nil
}

func a2aStartupLines(surface A2ASurface) []startupLogLine { _ = "STUB: not implemented"; return nil }
