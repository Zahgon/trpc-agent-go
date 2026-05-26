//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package evaluation

import (
	"context"
	"net/http"
	"time"

	coreevaluation "trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

const (
	headerAllow                = "Allow"
	headerContentType          = "Content-Type"
	headerAccessControlOrigin  = "Access-Control-Allow-Origin"
	headerAccessControlMethods = "Access-Control-Allow-Methods"
	headerAccessControlHeaders = "Access-Control-Allow-Headers"

	contentTypeJSON = "application/json"
)

// Server provides an HTTP API server for online evaluation flows.
type Server struct {
	appName           string
	basePath          string
	setsPath          string
	metricsPath       string
	runsPath          string
	resultsPath       string
	timeout           time.Duration
	agentEvaluator    coreevaluation.AgentEvaluator
	evalSetManager    evalset.Manager
	metricManager     metric.Manager
	evalResultManager evalresult.Manager
	routeRegistrars   []RouteRegistrar
	handler           http.Handler
}

// New creates a new evaluation server.
func New(opts ...Option) (*Server, error) { _ = "STUB: not implemented"; return nil, nil }

// Handler returns the HTTP handler exposed by the evaluation server.
func (s *Server) Handler() http.Handler {
	_ = "STUB: not implemented"

	// BasePath returns the base path exposed by the evaluation server.
	return *new(http.Handler)
}

func (s *Server) BasePath() string {
	_ = "STUB: not implemented"

	// SetsPath returns the sets collection endpoint path.
	return ""
}

func (s *Server) SetsPath() string {
	_ = "STUB: not implemented"

	// MetricsPath returns the metrics collection endpoint path.
	return ""
}

func (s *Server) MetricsPath() string { _ = "STUB: not implemented"; return "" }

// RunsPath returns the runs collection endpoint path.
func (s *Server) RunsPath() string {
	_ = "STUB: not implemented"

	// ResultsPath returns the results collection endpoint path.
	return ""
}

func (s *Server) ResultsPath() string { _ = "STUB: not implemented"; return "" }

// Close closes the evaluation server.
func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) setupHandler() error { _ = "STUB: not implemented"; return nil }

// Register collection and item routes for evaluation sets.

// Register collection and item routes for evaluation metrics.

// Register collection and item routes for evaluation runs.

// Register collection and item routes for evaluation results.

func normalizeBasePath(path string) string { _ = "STUB: not implemented"; return "" }

func joinURLPath(basePath, child string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Server) handleCORS(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (s *Server) redirectTrailingSlashToCanonicalPath(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) respondJSON(w http.ResponseWriter, r *http.Request, statusCode int, payload any) {
	_ = "STUB: not implemented"
	return
}

func statusCodeFromError(err error) int { _ = "STUB: not implemented"; return 0 }

func (s *Server) respondStatusError(w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) logResponseWriteError(r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

func newExecutionContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (s *Server) decodeJSONRequestBody(w http.ResponseWriter, r *http.Request, dst any) error {
	_ = "STUB: not implemented"
	return nil
}

func errorMessageFromError(err error) string { _ = "STUB: not implemented"; return "" }
