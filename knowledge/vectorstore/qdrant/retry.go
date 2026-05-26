//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package qdrant

import (
	"context"
	"time"
)

// retryConfig holds the configuration for retry operations.
type retryConfig struct {
	maxRetries     int
	baseRetryDelay time.Duration
	maxRetryDelay  time.Duration
}

// isTransientError checks if the error is a transient gRPC error that can be retried.
func isTransientError(err error) bool { _ = "STUB: not implemented"; return false }

// retry executes the operation with exponential backoff for transient errors.
func retry[T any](ctx context.Context, cfg retryConfig, op func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// retryVoid executes a void operation with exponential backoff for transient errors.
func retryVoid(ctx context.Context, cfg retryConfig, op func() error) error {
	_ = "STUB: not implemented"
	return nil
}
