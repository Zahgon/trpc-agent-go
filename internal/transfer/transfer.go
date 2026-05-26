//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package transfer contains internal transfer extension contracts.
package transfer

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
)

type transferMessageContextKey struct{}

const syntheticCompletionExtensionKey = "trpc_agent.transfer.synthetic_completion"

// ContextWithTransferMessage returns a context carrying the raw transfer message.
func ContextWithTransferMessage(ctx context.Context, message string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TransferMessageFromContext returns the raw transfer message carried by ctx.
func TransferMessageFromContext(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// MarkSyntheticCompletionEvent marks a completion event synthesized by transfer.
func MarkSyntheticCompletionEvent(evt *event.Event) { _ = "STUB: not implemented"; return }

// IsSyntheticCompletionEvent reports whether transfer synthesized the event.
func IsSyntheticCompletionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// InvocationCustomizer customizes a transfer target invocation before it runs.
type InvocationCustomizer interface {
	CustomizeTransferInvocation(
		ctx context.Context,
		source *agent.Invocation,
		target *agent.Invocation,
	) error
}

// CompletionObserver observes when a transfer target invocation completes.
type CompletionObserver interface {
	OnTransferComplete(
		ctx context.Context,
		source *agent.Invocation,
		target *agent.Invocation,
		targetEvent *event.Event,
	)
}

// TerminalErrorObserver observes when a transfer target invocation terminates
// with an error.
type TerminalErrorObserver interface {
	OnTransferTerminalError(
		ctx context.Context,
		source *agent.Invocation,
		target *agent.Invocation,
		targetEvent *event.Event,
	)
}
