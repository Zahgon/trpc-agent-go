//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agent

import (
	"context"
)

// InvocationContext carries the invocation information.
type InvocationContext struct {
	context.Context
}
type invocationKey struct{}

// NewInvocationContext creates a new InvocationContext.
func NewInvocationContext(ctx context.Context, invocation *Invocation) *InvocationContext {
	_ = "STUB: not implemented"
	return nil
}

// InvocationFromContext returns the invocation from the context.
func InvocationFromContext(ctx context.Context) (*Invocation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// EnsureInvocation ensures ctx contains a non-nil invocation and returns the updated context and invocation.
func EnsureInvocation(ctx context.Context) (context.Context, *Invocation) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// GetStateValueFromContext retrieves a typed value from the invocation state
// stored in the context.
//
// Returns the typed value and true if the invocation exists, the key exists,
// and the type matches, or the zero value and false otherwise.
//
// Example:
//
//	if startTime, ok := GetStateValueFromContext[time.Time](ctx, "agent:start_time"); ok {
//	    duration := time.Since(startTime)
//	}
//	if requestID, ok := GetStateValueFromContext[string](ctx, "middleware:request_id"); ok {
//	    log.Printf("Request ID: %s", requestID)
//	}
func GetStateValueFromContext[T any](ctx context.Context, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// GetRuntimeStateValueFromContext retrieves a typed value from the runtime state
// stored in the invocation's RunOptions within the context.
//
// Returns the typed value and true if the invocation exists, the key exists in
// RuntimeState, and the type matches, or the zero value and false otherwise.
//
// Example:
//
//	if userID, ok := GetRuntimeStateValueFromContext[string](ctx, "user_id"); ok {
//	    log.Printf("User ID: %s", userID)
//	}
//	if roomID, ok := GetRuntimeStateValueFromContext[int](ctx, "room_id"); ok {
//	    log.Printf("Room ID: %d", roomID)
//	}
func GetRuntimeStateValueFromContext[T any](ctx context.Context, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// CheckContextCancelled check context cancelled
func CheckContextCancelled(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
