//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package langfuse

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

// Start starts telemetry with Langfuse integration using the function option pattern.
func Start(ctx context.Context, opts ...Option) (clean func(context.Context) error, err error) {
	_ = "STUB: not implemented"
	// Start with default config from environment
	return nil, nil
}

// Apply user-provided options

// Apply truncation config early so callers can rely on it even if Start returns an error.

// Add insecure option only when explicitly configured

func start(ctx context.Context, opts ...otlptracehttp.Option) (clean func(context.Context) error, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encodeAuth encodes the public and secret keys for basic authentication.
func encodeAuth(pk, sk string) string { _ = "STUB: not implemented"; return "" }
