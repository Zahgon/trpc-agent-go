//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package agent

import "context"

type graphCompletionCaptureKey struct{}

// WithGraphCompletionCapture keeps terminal graph completion events available
// to internal graph consumers even when caller-visible forwarding is disabled.
func WithGraphCompletionCapture(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithoutGraphCompletionCapture clears any inherited capture flag for the
// current visible stream while preserving the rest of the context.
func WithoutGraphCompletionCapture(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ShouldCaptureGraphCompletion reports whether the current context keeps
// terminal graph completion events available for internal consumers.
func ShouldCaptureGraphCompletion(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func graphCompletionCaptureValue(ctx context.Context) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// PreserveGraphCompletionCapture copies the graph completion capture setting
// from base into next when next does not provide an explicit override.
func PreserveGraphCompletionCapture(
	base context.Context,
	next context.Context,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
