//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graphagent

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

type terminalMessageFilter struct {
	enabled                bool
	llmNodeAuthors         map[string]struct{}
	terminalLLMAuthors     map[string]struct{}
	agentScopePrefixes     map[string]struct{}
	terminalScopePrefixes  map[string]struct{}
	terminalResponseIDs    map[string]struct{}
	nonTerminalResponseIDs map[string]struct{}
}

func newTerminalMessageFilter(
	invocation *agent.Invocation,
	g *graph.Graph,
) *terminalMessageFilter {
	_ = "STUB: not implemented"
	return nil
}

func terminalMessageRootFilterKey(
	invocation *agent.Invocation,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *terminalMessageFilter) Allows(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *terminalMessageFilter) observe(evt *event.Event) { _ = "STUB: not implemented"; return }

func (f *terminalMessageFilter) recordResponseID(
	responseID string,
	terminal bool,
) {
	_ = "STUB: not implemented"
	return
}

func (f *terminalMessageFilter) allowsVisibleGraphCompletion(
	evt *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *terminalMessageFilter) matchAgentScopePrefix(
	filterKey string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Overlapping scopes such as "a" and "a/b" are ambiguous for an event
// keyed as "a/b/...". Filter only when the scope owner is unique.

func terminalAgentScopePrefix(rootFilterKey string, node *graph.Node) string {
	_ = "STUB: not implemented"
	return ""
}

func isTerminalMessageNode(g *graph.Graph, node *graph.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// Without an explicit PathMap, runtime may still fall back to a
// concrete node ID, so we cannot prove the node is terminal.

func shouldFilterTerminalMessageEvent(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func completionResponseIDFromStateDelta(evt *event.Event) string {
	_ = "STUB: not implemented"
	return ""
}

func matchesFilterPrefix(filterKey string, prefix string) bool {
	_ = "STUB: not implemented"
	return false
}
