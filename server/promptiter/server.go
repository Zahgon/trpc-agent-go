//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package promptiter

import (
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/manager"
)

const (
	headerAllow                = "Allow"
	headerContentType          = "Content-Type"
	headerAccessControlOrigin  = "Access-Control-Allow-Origin"
	headerAccessControlMethods = "Access-Control-Allow-Methods"
	headerAccessControlHeaders = "Access-Control-Allow-Headers"
	contentTypeJSON            = "application/json"
)

// Server exposes a control-plane HTTP API for a single PromptIter target app.
type Server struct {
	appName                string
	basePath               string
	structurePath          string
	runsPath               string
	asyncRunsPath          string
	timeout                time.Duration
	engine                 engine.Engine
	manager                manager.Manager
	responseResultSlimming engine.RunResultSlimming
	handler                http.Handler
}

// New creates a new PromptIter HTTP server.
func New(opts ...Option) (*Server, error) { _ = "STUB: not implemented"; return nil, nil }

// Handler returns the HTTP handler exposed by the PromptIter server.
func (s *Server) Handler() http.Handler {
	_ = "STUB: not implemented"

	// BasePath returns the base path exposed by the PromptIter server.
	return *new(http.Handler)
}

func (s *Server) BasePath() string {
	_ = "STUB: not implemented"

	// StructurePath returns the structure endpoint path.
	return ""
}

func (s *Server) StructurePath() string { _ = "STUB: not implemented"; return "" }

// RunsPath returns the runs endpoint path.
func (s *Server) RunsPath() string {
	_ = "STUB: not implemented"

	// AsyncRunsPath returns the asynchronous runs endpoint path.
	return ""
}

func (s *Server) AsyncRunsPath() string { _ = "STUB: not implemented"; return "" }

// Close closes the PromptIter server.
func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) setupHandler() { _ = "STUB: not implemented"; return }

func normalizeBasePath(path string) string { _ = "STUB: not implemented"; return "" }

func joinURLPath(basePath, child string) (string, error) { _ = "STUB: not implemented"; return "", nil }
