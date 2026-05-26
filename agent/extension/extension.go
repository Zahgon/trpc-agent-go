//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package extension

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Extension is an agent-scoped capability bundle. See the package
// documentation for the full rationale and lifecycle.
type Extension interface {
	// Name identifies the extension for logging, error wrapping
	// and duplicate detection. Must be non-empty and stable for
	// the lifetime of the agent that hosts the extension.
	Name() string

	// Register is called exactly once during agent construction.
	// The Registry passed in accumulates callbacks and tools; the
	// extension MUST NOT keep a reference to it beyond the call.
	Register(r *Registry)
}

// Contributions is the aggregated result of running Register on
// every extension passed to a consuming agent.
//
// The concrete storage is intentionally opaque. Consuming agents
// should use the accessor methods below instead of depending on
// the internal representation, which leaves this package room to
// evolve how extension contributions are stored without changing
// the public Extension / Registry authoring contract.
type Contributions struct {
	agentCallbacks *agent.Callbacks
	modelCallbacks *model.Callbacks
	toolCallbacks  *tool.Callbacks
	tools          []tool.Tool
}

// Collect runs Register on every extension in install order and
// returns the aggregated Contributions. It enforces two invariants that
// every consuming agent would otherwise have to re-implement:
//
//   - no nil extensions
//   - no two extensions sharing the same Name (case-sensitive)
//
// Both violations are returned as a non-nil error and the partially-
// built Contributions is discarded; consuming agents should treat any
// non-nil error from Collect as fatal-during-construction.
// A panic from Extension.Register is also converted into an error
// that includes the extension name, index and stack trace; later
// extensions are not registered after such a panic.
//
// When extensions is empty the function returns (nil, nil) so the
// nil-Contributions short-circuit path on the consumer side stays simple.
func Collect(extensions []Extension) (*Contributions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func safeRegister(
	e Extension,
	r *Registry,
	name string,
	index int,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AgentCallbacks returns an independent copy of the contributed
// agent callback chain, or nil when no extension contributed agent
// callbacks.
func (c *Contributions) AgentCallbacks() *agent.Callbacks { _ = "STUB: not implemented"; return nil }

// ModelCallbacks returns an independent copy of the contributed
// model callback chain, or nil when no extension contributed model
// callbacks.
func (c *Contributions) ModelCallbacks() *model.Callbacks { _ = "STUB: not implemented"; return nil }

// ToolCallbacks returns an independent copy of the contributed tool
// callback chain, or nil when no extension contributed tool
// callbacks.
func (c *Contributions) ToolCallbacks() *tool.Callbacks { _ = "STUB: not implemented"; return nil }

// Tools returns a shallow copy of the contributed tools in install
// order. Tool values themselves are not cloned; consuming agents
// should treat tool.Tool implementations as immutable after
// construction.
func (c *Contributions) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// IsEmpty reports whether c carries no contributions. Convenience
// helper for consumers that want to skip the merge pipeline
// entirely when no extension actually populated anything.
func (c *Contributions) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// hasAgentContent / hasModelContent / hasToolContent are tiny
// content predicates. They live next to Contributions because consumers
// frequently need to ask "did any extension populate this slot?"
// before deciding whether to allocate a merged callback chain on
// the user-side; centralising the answer keeps the per-agent
// merge code free of "is this empty Callbacks worth merging?"
// boilerplate.
func hasAgentContent(c *agent.Callbacks) bool { _ = "STUB: not implemented"; return false }

func hasModelContent(c *model.Callbacks) bool { _ = "STUB: not implemented"; return false }

func hasToolContent(c *tool.Callbacks) bool { _ = "STUB: not implemented"; return false }
