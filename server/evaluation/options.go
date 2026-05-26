//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package evaluation provides a reusable HTTP API server for online evaluation workflows.
package evaluation

import (
	"time"

	coreevaluation "trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

const (
	defaultBasePath    = "/evaluation"
	defaultSetsPath    = "/sets"
	defaultMetricsPath = "/metrics"
	defaultRunsPath    = "/runs"
	defaultResultsPath = "/results"
)

// Option configures the evaluation server.
type Option func(*options)

type options struct {
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
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithAppName sets the app name used by the evaluation server.
func WithAppName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBasePath sets the base path used by the evaluation server.
func WithBasePath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSetsPath sets the sets collection path relative to BasePath.
func WithSetsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetricsPath sets the metrics collection path relative to BasePath.
func WithMetricsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunsPath sets the runs collection path relative to BasePath.
func WithRunsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithResultsPath sets the results collection path relative to BasePath.
func WithResultsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout sets the maximum execution time for an online evaluation run.
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAgentEvaluator sets the agent evaluator used by the evaluation server.
func WithAgentEvaluator(agentEvaluator coreevaluation.AgentEvaluator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvalSetManager sets the eval set manager used by the evaluation server.
// When omitted, set routes are not registered.
func WithEvalSetManager(manager evalset.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMetricManager sets the metric manager used by the evaluation server.
// When omitted, metric routes are not registered.
func WithMetricManager(manager metric.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEvalResultManager sets the eval result manager used by the evaluation server.
// When omitted, result routes are not registered.
func WithEvalResultManager(manager evalresult.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRouteRegistrar appends a custom route registrar to the evaluation server.
func WithRouteRegistrar(registrar RouteRegistrar) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
