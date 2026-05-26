//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package plugin provides runner-scoped extensions.
package plugin

import (
	"context"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	errNilPlugin = errors.New("plugin is nil")
	errEmptyName = errors.New("plugin name is empty")
)

// Plugin registers hooks into a Runner.
//
// Plugins are registered once on a Runner and applied automatically to all
// invocations created by that Runner.
type Plugin interface {
	// Name returns a stable, unique name for this plugin instance.
	Name() string

	// Register wires plugin callbacks into the provided Registry.
	Register(r *Registry)
}

// Closer is implemented by plugins that need to release resources.
type Closer interface {
	Close(ctx context.Context) error
}

// EventHook is invoked for each event passing through the Runner.
type EventHook func(
	ctx context.Context,
	invocation *agent.Invocation,
	e *event.Event,
) (*event.Event, error)

// Registry exposes hook registration points for a single plugin.
type Registry struct {
	name string
	mgr  *Manager
}

// BeforeAgent registers a before-agent callback.
func (r *Registry) BeforeAgent(cb agent.BeforeAgentCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterAgent registers an after-agent callback.
func (r *Registry) AfterAgent(cb agent.AfterAgentCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// BeforeModel registers a before-model callback.
func (r *Registry) BeforeModel(cb model.BeforeModelCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterModel registers an after-model callback.
func (r *Registry) AfterModel(cb model.AfterModelCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// BeforeTool registers a before-tool callback.
func (r *Registry) BeforeTool(cb tool.BeforeToolCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// AfterTool registers an after-tool callback.
func (r *Registry) AfterTool(cb tool.AfterToolCallbackStructured) {
	_ = "STUB: not implemented"
	return
}

// OnEvent registers an event hook.
func (r *Registry) OnEvent(hook EventHook) { _ = "STUB: not implemented"; return }

// Manager composes multiple plugins into callback sets.
//
// Manager implements agent.PluginManager.
type Manager struct {
	plugins        []Plugin
	agentCallbacks *agent.Callbacks
	modelCallbacks *model.Callbacks
	toolCallbacks  *tool.Callbacks
	eventHooks     []namedEventHook
}

type namedEventHook struct {
	name string
	hook EventHook
}

// NewManager builds a Manager and registers all plugin hooks.
func NewManager(plugins ...Plugin) (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

// MustNewManager panics if plugin registration fails.
func MustNewManager(plugins ...Plugin) *Manager { _ = "STUB: not implemented"; return nil }

// AgentCallbacks implements agent.PluginManager.
func (m *Manager) AgentCallbacks() *agent.Callbacks { _ = "STUB: not implemented"; return nil }

// ModelCallbacks implements agent.PluginManager.
func (m *Manager) ModelCallbacks() *model.Callbacks { _ = "STUB: not implemented"; return nil }

// ToolCallbacks implements agent.PluginManager.
func (m *Manager) ToolCallbacks() *tool.Callbacks { _ = "STUB: not implemented"; return nil }

// OnEvent implements agent.PluginManager.
func (m *Manager) OnEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	e *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close implements agent.PluginManager.
func (m *Manager) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
