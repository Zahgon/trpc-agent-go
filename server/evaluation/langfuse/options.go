//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package langfuse

import (
	"context"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

const (
	defaultPath        = "/langfuse/remote-experiment"
	defaultUserID      = "langfuse-remote-user"
	defaultEnvironment = "development"
	defaultTimeout     = time.Hour
)

// Option configures the Langfuse remote experiment handler.
type Option func(*options)

type options struct {
	path           string
	baseURL        string
	publicKey      string
	secretKey      string
	caseBuilder    CaseBuilder
	traceTags      []string
	userIDSupplier UserIDSupplier
	environment    string
	timeout        time.Duration
	httpClient     *http.Client
	runOptions     []agent.RunOption
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithPath sets the handler route path.
func WithPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets the Langfuse public API base URL.
func WithBaseURL(baseURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPublicKey sets the Langfuse public API key.
func WithPublicKey(publicKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSecretKey sets the Langfuse secret API key.
func WithSecretKey(secretKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCaseBuilder sets the dataset item to case conversion function.
func WithCaseBuilder(caseBuilder CaseBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTraceTags sets the default trace tags used when the payload does not override them.
func WithTraceTags(tags ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// UserIDSupplier returns the default user ID used by one remote experiment run.
type UserIDSupplier func(ctx context.Context) string

// WithUserIDSupplier sets the user ID supplier used when the case spec does not provide one.
func WithUserIDSupplier(supplier UserIDSupplier) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnvironment sets the default Langfuse environment attached to traces and scores.
func WithEnvironment(environment string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout sets the maximum execution time for one remote experiment request.
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient sets the HTTP client used for Langfuse public API calls.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunOptions appends agent run options applied to every remote experiment case.
func WithRunOptions(runOptions ...agent.RunOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
