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

// Registry is the per-extension view handed to Extension.Register.
// It looks like a small fan-out API but is actually a thin facade
// over the Contributions-wide callback containers — every Register* call
// appends into the same *agent.Callbacks / *model.Callbacks /
// *tool.Callbacks that Collect built once. Tools, on the other
// hand, are accumulated per-Registry first and merged into the
// Contributions by Collect after Register returns; this preserves the
// install-order tool layout used by LLMAgent's earlier-wins
// dedup rule.
//
// All callback methods accept the *Structured* callback variants
// (BeforeAgentCallbackStructured, etc.). Errors returned by an
// extension's callback are wrapped with the extension's Name so
// observability code can trace which extension produced the
// failure — mirroring plugin.Registry's wrapping convention.
//
// Registry is NOT safe for concurrent mutation. Extensions are
// expected to run Register synchronously inside the agent
// constructor; the Registry is discarded once Register returns.
type Registry struct {
	name           string
	agentCallbacks *agent.Callbacks
	modelCallbacks *model.Callbacks
	toolCallbacks  *tool.Callbacks
	tools          []tool.Tool
}

// newRegistry is the package-internal Registry constructor used by
// Collect. The callback containers are shared across all Registry
// instances Collect creates for a single Contributions value, so the merged
// callback order across extensions matches install order naturally
// without an extra merge step.
func newRegistry(
	name string,
	ac *agent.Callbacks,
	mc *model.Callbacks,
	tc *tool.Callbacks,
) *Registry {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name the extension declared at construction
// time. Exposed so extensions whose tools or callbacks need to
// embed their owning extension's identifier (for logging, metrics,
// state-key prefixes) can read it back without storing it twice.
func (r *Registry) Name() string { _ = "STUB: not implemented"; return "" }

// Tools appends framework-managed tools contributed by this
// extension. nil entries are silently dropped (no panic) so
// extensions can build a slice with conditional inclusion without
// guarding every append at the call site.
//
// Tool name dedup is the CONSUMING agent's responsibility, not
// Registry's. LLMAgent's policy is earlier-wins (see
// appendExtensionTools); other consumers may differ.
func (r *Registry) Tools(tools ...tool.Tool) { _ = "STUB: not implemented"; return }

// BeforeAgent registers a before-agent callback. The callback is
// wrapped so any non-nil error it returns is prefixed with the
// extension name.
func (r *Registry) BeforeAgent(cb agent.BeforeAgentCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterAgent registers an after-agent callback. Error wrapping
// behaviour matches BeforeAgent.
func (r *Registry) AfterAgent(cb agent.AfterAgentCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// BeforeModel registers a before-model callback. Error wrapping
// behaviour matches BeforeAgent.
func (r *Registry) BeforeModel(cb model.BeforeModelCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterModel registers an after-model callback. Error wrapping
// behaviour matches BeforeAgent.
func (r *Registry) AfterModel(cb model.AfterModelCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// BeforeTool registers a before-tool callback. Error wrapping
// behaviour matches BeforeAgent.
func (r *Registry) BeforeTool(cb tool.BeforeToolCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterTool registers an after-tool callback. Error wrapping
// behaviour matches BeforeAgent.
func (r *Registry) AfterTool(cb tool.AfterToolCallbackStructured) {
	_ = "STUB: not implemented"
	return
}
