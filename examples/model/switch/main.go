//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates model switching with the runner.
//
// This example shows two ways to switch models:
//
// 1. Agent-level switching (/switch command):
//   - Changes the agent's default model permanently.
//   - Affects all subsequent requests until changed again.
//   - Use agent.SetModelByName() or agent.SetModel().
//   - Suitable for: changing default behavior, user preferences.
//
// 2. Per-request switching (/model command):
//   - Overrides model for a single request only.
//   - Agent's default model remains unchanged.
//   - Use agent.WithModelName() or agent.WithModel() in RunOptions.
//   - Suitable for: temporary overrides, A/B testing, special queries.
//
// Example usage:
//
//	/switch deepseek-v4-pro    → All future requests use Pro.
//	/model deepseek-v4-flash   → Next request uses Flash, then back to Pro.
package main

import (
	"context"
	"flag"
	"fmt"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func main() {
	// Flags.
	defaultModel := flag.String("model", "deepseek-v4-flash", "Default model name")
	flag.Parse()

	app := &chatApp{
		defaultModel: *defaultModel,
	}
	ctx := context.Background()

	if err := app.setup(ctx); err != nil {
		fmt.Printf("❌ Setup failed: %v\n", err)
		return
	}

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer app.runner.Close()

	if err := app.startChat(ctx); err != nil {
		fmt.Printf("❌ Chat failed: %v\n", err)
	}
}

// chatApp manages the conversation and model switching.
//
// This struct demonstrates state management for two switching approaches:
//   - agent: holds the default model (changed by Method 1).
//   - nextModelName/usePerRequestSwitch: temporary state for Method 2.
type chatApp struct {
	defaultModel        string
	agent               *llmagent.LLMAgent
	runner              runner.Runner
	models              map[string]model.Model // Registered models for validation.
	sessionID           string                 // Current session ID.
	nextModelName       string                 // Model name for next request (empty = use agent's current model).
	usePerRequestSwitch bool                   // Whether to use per-request switching for next message.
}

// setup initializes models, agent, and runner.
//
// Key setup steps for model switching:
//  1. Pre-register all available models in a map.
//  2. Create agent with WithModels() to enable name-based lookup.
//  3. Set initial default model with WithModel().
//
// This setup enables both switching methods:
//   - Method 1: agent.SetModelByName() works because models are registered.
//   - Method 2: agent.WithModelName() in RunOptions also uses the registry.
func (a *chatApp) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Prepare model map with pre-registered models.
// Pre-registration is required for name-based model switching.

// Get the default model instance.

// Create an agent with pre-registered models.
// WithModels: registers the model map for name-based lookup.
// WithModel: sets the initial default model.

// Store models map for validation in handleModelCommand.

// Create runner.

// Initialize session ID.

// startChat runs the interactive conversation loop.
//
// Comparison of two model switching methods:
//
// ┌─────────────────┬──────────────────────┬──────────────────────┐
// │ Feature         │ Method 1: Agent-level│ Method 2: Per-request│
// ├─────────────────┼──────────────────────┼──────────────────────┤
// │ Scope           │ All future requests  │ Single request only  │
// │ Persistence     │ Until changed again  │ Auto-revert after use│
// │ Thread-safety   │ Yes (atomic)         │ Yes (isolated)       │
// │ API             │ SetModelByName()     │ WithModelName()      │
// │ State location  │ Agent instance       │ RunOptions           │
// │ Typical use     │ User preference      │ A/B testing, fallback│
// └─────────────────┴──────────────────────┴──────────────────────┘
func (a *chatApp) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Switch command: changes agent's default model (affects all subsequent requests).

// Model command: sets model for next request only (per-request
// override).

// New session.

// Exit.

// Normal message.

// processMessage sends a message to the agent via runner.
//
// This function demonstrates how to apply per-request model switching:
//  1. Check if a temporary model override is set.
//  2. If yes, add agent.WithModelName() to RunOptions.
//  3. Clear the temporary state after use.
//
// Alternative approaches for per-request switching:
//   - agent.WithModelName(name): lookup from registered models.
//   - agent.WithModel(instance): use a specific model instance.
//   - agent.WithRuntimeState(map): pass additional runtime parameters.
//
// Note: RunOptions are passed as variadic arguments to runner.Run().
func (a *chatApp) processMessage(ctx context.Context, text string) error {
	_ = "STUB: not implemented"
	// Build run options.
	return nil
}

// Apply per-request model override if specified.
// This demonstrates Method 2: per-request switching.

// Option 1: Switch by model name (recommended if model is pre-registered).

// Option 2: Switch by model instance (use when you need custom config).
// modelInstance := openai.New(a.nextModelName)
// runOpts = append(runOpts, agent.WithModel(modelInstance))

// Reset for next request to ensure this only affects current request.

// Run the agent via runner.
// The runner will use the per-request model if specified, otherwise
// falls back to the agent's default model.

// processResponse prints streaming or non-streaming responses.
func (a *chatApp) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle streaming delta content.

// Handle non-streaming message content.

// If streaming, we already printed it; just add newline.

// Non-streaming: print the complete response.

// handleSwitch switches agent's default model (affects all subsequent
// requests).
//
// Agent-level switching (Method 1):
//   - Permanently changes the agent's default model.
//   - All subsequent requests use the new model.
//   - Thread-safe: can be called concurrently.
//   - State persists until explicitly changed again.
//
// Implementation options:
//
//	a) SetModelByName(name): lookup model from pre-registered models map.
//	b) SetModel(instance): directly set a model instance.
//
// Use cases:
//   - User changes their preferred model.
//   - Application switches to a different model tier.
//   - Adapting to different conversation contexts.
func (a *chatApp) handleSwitch(name string) error {
	_ = "STUB: not implemented"
	// Switch model by name using SetModelByName method.
	// This changes the agent's default model for all subsequent requests.
	return nil
}

// List available models on error.

// Alternative: use SetModel to switch by model instance.
// This is useful when you need to create a new model with specific
// configuration:
//   model := openai.New("deepseek-v4-pro")
//   a.agent.SetModel(model)

// handleModelCommand sets model for next request only (per-request override).
//
// Per-request switching (Method 2):
//   - Temporarily overrides model for a single request.
//   - Agent's default model remains unchanged.
//   - No side effects on concurrent requests.
//   - Automatically reverts after the request completes.
//
// Implementation:
//   - Validate the model name exists in registered models.
//   - Store the model name/instance temporarily.
//   - Pass it via agent.WithModelName() or agent.WithModel() in RunOptions.
//   - Clear the temporary state after use.
//
// Use cases:
//   - Testing different models for comparison.
//   - Using a specialized model for specific query types.
//   - A/B testing without affecting other users.
//   - Fallback to a different model for retry scenarios.
func (a *chatApp) handleModelCommand(name string) error {
	_ = "STUB: not implemented"
	// Validate that the model exists in the agent's registered models.
	// This prevents runtime errors when the request is executed.
	return nil
}

// List available models on error.

// This uses per-request model switching via agent.WithModelName().
// The agent's default model remains unchanged.

// startNewSession creates a new session ID and clears history.
func (a *chatApp) startNewSession() {
	_ = "STUB: not implemented"

	// Generate a new session ID based on timestamp.
	return
}
