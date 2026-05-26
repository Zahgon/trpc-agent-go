//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package mem0 provides an ingest-first integration with mem0.ai.
package mem0

import (
	"net/http"
	"time"
)

const (
	defaultHost             = "https://api.mem0.ai"
	defaultTimeout          = 10 * time.Second
	defaultAsyncMemoryNum   = 1
	defaultMemoryQueueSize  = 10
	defaultMemoryJobTimeout = 30 * time.Second
)

type serviceOpts struct {
	host   string
	apiKey string

	orgID     string
	projectID string

	asyncMode bool
	version   string

	timeout time.Duration
	client  *http.Client

	loadToolEnabled bool

	asyncMemoryNum   int
	memoryQueueSize  int
	memoryJobTimeout time.Duration
}

func (o serviceOpts) clone() serviceOpts { _ = "STUB: not implemented"; return *new(serviceOpts) }

var defaultOptions = serviceOpts{
	host:             defaultHost,
	asyncMode:        true,
	version:          "v2",
	timeout:          defaultTimeout,
	asyncMemoryNum:   defaultAsyncMemoryNum,
	memoryQueueSize:  defaultMemoryQueueSize,
	memoryJobTimeout: defaultMemoryJobTimeout,
}

// ServiceOpt configures a mem0 service.
type ServiceOpt func(*serviceOpts)

// WithHost sets the mem0 API host or base URL.
func WithHost(host string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithAPIKey sets the mem0 API key used for all requests.
func WithAPIKey(apiKey string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithOrgProject sets optional mem0 organization and project identifiers.
func WithOrgProject(orgID, projectID string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncMode controls whether mem0 ingest requests are async.
func WithAsyncMode(async bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithVersion sets the mem0 ingestion API version for create requests.
func WithVersion(version string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithTimeout sets the HTTP timeout for mem0 requests.
func WithTimeout(timeout time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithHTTPClient injects a custom HTTP client for mem0 requests.
func WithHTTPClient(c *http.Client) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithLoadToolEnabled controls whether memory_load is exposed in Tools().
func WithLoadToolEnabled(enabled bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncMemoryNum sets the number of async mem0 ingestion workers.
func WithAsyncMemoryNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryQueueSize sets the queue size for async mem0 ingestion jobs.
func WithMemoryQueueSize(size int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryJobTimeout sets the timeout applied to each ingest job. This
// governs both queued async worker jobs and the synchronous fallback path
// when the queue is full.
func WithMemoryJobTimeout(timeout time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}
