//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const graphCallOptionsKey = "graph_call_options"

// NodePath identifies a node inside nested subgraphs.
//
// Each segment is a node ID. For example:
//   - {"child"} targets the "child" node on the current graph.
//   - {"child", "llm"} targets the "llm" node inside the "child" subgraph.
type NodePath []string

// CallOption configures per-invocation graph call options.
//
// CallOption is intentionally closed to external implementations.
type CallOption func(*callOptions)

// WithCallOptions sets graph call options for this run.
//
// Call options are stored on the invocation and will be propagated into
// nested GraphAgent subgraphs automatically.
func WithCallOptions(opts ...CallOption) agent.RunOption {
	_ = "STUB: not implemented"
	return *new(agent.RunOption)
}

// WithCallGenerationConfigPatch applies a GenerationConfigPatch to all LLM
// nodes in the current graph scope.
func WithCallGenerationConfigPatch(p model.GenerationConfigPatch) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

// WithCallResumeStateOverrideKeys preserves caller-provided runtime state values
// for the selected keys when resuming from a checkpoint.
//
// When one of these keys is present in RunOptions.RuntimeState for the current
// graph scope, its value overrides the checkpoint-restored value during resume.
// Keys filtered by the existing resume merge logic are still ignored.
func WithCallResumeStateOverrideKeys(keys ...string) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

// DesignateNode applies options to a specific node in the current graph
// scope.
//
// For agent/subgraph nodes, options are applied to the child invocation and
// therefore affect the nested graph.
func DesignateNode(nodeID string, opts ...CallOption) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

// DesignateNodeWithPath applies options to a node inside nested subgraphs.
func DesignateNodeWithPath(path NodePath, opts ...CallOption) CallOption {
	_ = "STUB: not implemented"
	return *new(CallOption)
}

type callOptions struct {
	generation              model.GenerationConfigPatch
	resumeStateOverrideKeys map[string]struct{}
	nodes                   map[string]*callNodeOptions
}

type callNodeOptions struct {
	generation model.GenerationConfigPatch
	child      *callOptions
}

func newCallOptions(opts ...CallOption) *callOptions { _ = "STUB: not implemented"; return nil }

func (c *callOptions) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (c *callOptions) ensureNode(nodeID string) *callNodeOptions {
	_ = "STUB: not implemented"
	return nil
}

func mergeCallOptions(a, b *callOptions) *callOptions { _ = "STUB: not implemented"; return nil }

func mergeCallNodes(
	a map[string]*callNodeOptions,
	b map[string]*callNodeOptions,
) map[string]*callNodeOptions {
	_ = "STUB: not implemented"
	return nil
}

func cloneCallOptions(in *callOptions) *callOptions { _ = "STUB: not implemented"; return nil }

func cloneCallNodeMap(
	in map[string]*callNodeOptions,
) map[string]*callNodeOptions {
	_ = "STUB: not implemented"
	return nil
}

func cloneCallNodeOptions(in *callNodeOptions) *callNodeOptions {
	_ = "STUB: not implemented"
	return nil
}

func isEmptyGenPatch(p model.GenerationConfigPatch) bool { _ = "STUB: not implemented"; return false }

func cloneGenPatch(
	p model.GenerationConfigPatch,
) model.GenerationConfigPatch {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfigPatch)
}

func mergeGenPatch(
	base model.GenerationConfigPatch,
	override model.GenerationConfigPatch,
) model.GenerationConfigPatch {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfigPatch)
}

func copyCustomAgentConfigs(in map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func newStringSet(keys ...string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func cloneStringSet(in map[string]struct{}) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func mergeStringSet(
	a map[string]struct{},
	b map[string]struct{},
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func graphCallOptionsFromConfigs(cfgs map[string]any) *callOptions {
	_ = "STUB: not implemented"
	return nil
}

func withScopedGraphCallOptions(
	cfgs map[string]any,
	nodeID string,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func scopeCallOptionsForSubgraph(
	parent *callOptions,
	nodeID string,
) *callOptions {
	_ = "STUB: not implemented"
	return nil
}

func generationPatchForNode(
	opts *callOptions,
	nodeID string,
) model.GenerationConfigPatch {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfigPatch)
}
