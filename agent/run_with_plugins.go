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

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

func pluginAgentCallbacks(inv *Invocation) *Callbacks { _ = "STUB: not implemented"; return nil }

// RunWithPlugins runs an agent with Runner-provided plugins applied.
//
// This wrapper is used by the Runner and by internal multi-agent helpers
// (e.g., chain, parallel, cycle, transfer, graph agent-nodes) so plugins
// consistently apply even when agents are invoked indirectly.
//
// Callback semantics (what callers should know):
//
//  1. BeforeAgent fires ONCE per invocation passed in. For multi-agent
//     containers (chain, parallel, cycle, graph agent-node), the callback
//     fires once for EACH sub-agent invocation — not just once for the root
//     run. Hooks that expect "one call per Runner turn" must check
//     `args.Invocation.Agent` or similar to distinguish.
//
//  2. BeforeAgent.CustomResponse SHORT-CIRCUITS the sub-agent: `ag.Run` is
//     not called and no other events are produced besides the synthetic
//     response event. Code paths that rely on the sub-agent emitting
//     terminal state (for example, graph's `GraphCompletionEvent` used by
//     agent-nodes to populate `SubgraphResult.FinalState`) will NOT see
//     those events — callers / output mappers must handle a nil
//     `FinalState` gracefully when short-circuit is possible.
//
//  3. BeforeAgent may return a derived Context; both `ag.Run` and the
//     background AfterAgent goroutine observe it via `CloneContext`.
//
//  4. AfterAgent fires after the sub-agent's event stream closes and
//     receives the last non-partial response event in `args.FullResponseEvent`
//     (nil if there wasn't one). If `args.Error` is non-nil, it was
//     derived from the sub-agent's final `model.ResponseError`.
//
//  5. AfterAgent.CustomResponse APPENDS an extra response event to the
//     forwarded stream. In consumers that track "last response" (the graph
//     agent-node's `StateKeyLastResponse` and analogous downstream state),
//     this appended event becomes the new last response. Use this
//     intentionally when you want to override the sub-agent's output.
//
//  6. AfterAgent returning an error appends a single error event of type
//     `ErrorTypeAgentCallbackError` instead of overriding the response.
func RunWithPlugins(
	ctx context.Context,
	invocation *Invocation,
	ag Agent,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func singleResponseEventChan(
	ctx context.Context,
	invocation *Invocation,
	rsp *model.Response,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func invocationID(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

func agentName(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

func wrapAfterAgentCallbacks(
	ctx context.Context,
	invocation *Invocation,
	callbacks *Callbacks,
	src <-chan *event.Event,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func agentErrorFromEvent(fullRespEvent *event.Event) error { _ = "STUB: not implemented"; return nil }

func afterAgentEvent(
	invocation *Invocation,
	result *AfterAgentResult,
	callbackErr error,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}
