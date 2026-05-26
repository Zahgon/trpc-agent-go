//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package agui provides the ability to communicate with the front end through the AG-UI protocol.
package agui

import (
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/service"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// Server provides AG-UI server.
type Server struct {
	basePath       string          // basePath is the base path for the service.
	appName        string          // appName is the static app name configured for the AG-UI runner.
	path           string          // path is the chat message endpoint path.
	handler        http.Handler    // handler serves chat and optional history routes.
	sessionService session.Service // sessionService backs stored conversations for snapshots.
}

// New creates a AG-UI server instance.
func New(runner runner.Runner, opt ...Option) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newService creates a new service instance.
func newService(runner runner.Runner, opts *options) (service.Service, error) {
	_ = "STUB: not implemented"
	return *new(service.Service), nil
}

// joinURLPath joins the base path and the path into a URL path.
func joinURLPath(basePath, path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Handler returns the http.Handler serving AG-UI requests.
func (s *Server) Handler() http.Handler {
	_ = "STUB: not implemented"

	// Path returns the chat message endpoint path joined with BasePath.
	return *new(http.Handler)
}

func (s *Server) Path() string {
	_ = "STUB: not implemented"

	// BasePath returns the base URL path prefix shared by chat and history endpoints.
	return ""
}

func (s *Server) BasePath() string { _ = "STUB: not implemented"; return "" }
