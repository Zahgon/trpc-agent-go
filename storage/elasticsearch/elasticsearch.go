//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package elasticsearch provides Elasticsearch client interface, implementation and options.
package elasticsearch

import (
	esv7 "github.com/elastic/go-elasticsearch/v7"
	esv8 "github.com/elastic/go-elasticsearch/v8"
	esv9 "github.com/elastic/go-elasticsearch/v9"

	ielasticsearch "trpc.group/trpc-go/trpc-agent-go/internal/storage/elasticsearch"
)

// defaultClientBuilder selects implementation by Version and builds a client.
func defaultClientBuilder(builderOpts ...ClientBuilderOpt) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// newClientV7 builds a v7 client from generic builder options.
func newClientV7(o *ClientBuilderOpts) (*esv7.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newClientV8 builds a v8 client from generic builder options.
func newClientV8(o *ClientBuilderOpts) (*esv8.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newClientV9 builds a v9 client from generic builder options.
func newClientV9(o *ClientBuilderOpts) (*esv9.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WrapSDKClient wraps a generic Elasticsearch SDK client with our storage interface.
//
// WARNING: This function is for INTERNAL USE ONLY!
// Do NOT call this function directly from external packages.
// This is an internal implementation detail that may change without notice.
// Use the public API provided by the parent storage/elasticsearch package instead.
//
// This function is only exported to allow access from other internal packages
// within the same module (knowledge/vectorstore/elasticsearch, etc.).
func WrapSDKClient(client any) (ielasticsearch.Client, error) {
	_ = "STUB: not implemented"
	return *new(ielasticsearch.Client), nil
}

// Already wrapped (useful for testing with mock clients)
// This case is placed last to ensure concrete SDK types are matched first
