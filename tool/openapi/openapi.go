//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package openapi provides a toolset for a given openapi API specification.
package openapi

import (
	"context"
	"net/http"
	"time"

	openapi "github.com/getkin/kin-openapi/openapi3"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// defaultUserAgent is the default user agent for HTTP requests.
	defaultUserAgent = "trpc-agent-go-openapi/1.0"
	// defaultOpenAPIToolSetName is the default name for the OpenAPI tool set.
	defaultOpenAPIToolSetName = "openapi"
	// defaultTimeout is the default timeout for HTTP requests.
	defaultTimeout = 30 * time.Second
)

// Option is a functional option for configuring the OpenAPI tool.
type Option func(*config)

// config holds the configuration for the OpenAPI tool.
type config struct {
	name      string
	userAgent string

	specLoader Loader
	httpClient *http.Client
}

// WithSpecLoader sets the spec loader to use.
func WithSpecLoader(loader Loader) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserAgent sets the user agent for HTTP requests.
func WithUserAgent(userAgent string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient sets the HTTP client to use.
func WithHTTPClient(httpClient *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the name of the openAPIToolSet.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// openAPIToolSet is a set of tools.
type openAPIToolSet struct {
	spec *docProcessor
	name string

	config *config
	tools  []tool.Tool
}

// Tools implements the ToolSet interface.
func (ts *openAPIToolSet) Tools(ctx context.Context) []tool.Tool {
	_ = "STUB: not implemented"

	// Close implements the ToolSet interface.
	return nil
}

func (ts *openAPIToolSet) Close() error {
	_ = "STUB: not implemented"
	// No resources to clean up for file tools.
	return nil
}

// Name implements the ToolSet interface.
func (ts *openAPIToolSet) Name() string { _ = "STUB: not implemented"; return "" }

// NewToolSet creates a new OpenAPI tool set with the provided options.
func NewToolSet(ctx context.Context, opts ...Option) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

func loadSpec(ctx context.Context, loader Loader) (*openapi.T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
