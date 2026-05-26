//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package identity

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// PluginName is the name used when registering the plugin with a Runner.
	PluginName = "identity"

	// stateKey is the invocation-state key where the resolved Identity is
	// stored. It is intentionally unexported: external callers should read
	// identity via FromContext (available inside tool callbacks and tool
	// implementations), not by peeking at invocation state directly.
	stateKey = "identity:current"
)

// Plugin is a Runner plugin that transparently resolves and injects user
// identity into every tool call.
//
// Resolved identity is stored in Invocation state during BeforeAgent and then
// attached to each tool-call context during BeforeTool. Sensitive fields such
// as Identity.Headers and Identity.EnvVars stay in context so HTTP clients and
// command executors can consume them at request or execution time without
// copying secrets into model-visible tool arguments.
//
// Consumer contract:
//
//   - HTTP-based tools (MCP SSE/Streamable, webhooks): install a per-request
//     hook that reads identity headers from context and writes them onto the
//     outgoing request. For MCP toolsets that means passing
//     mcp.WithHTTPBeforeRequest to tool/mcp.WithMCPOptions (see Usage below).
//   - Command-executing tools (skill_run, workspace_exec, interactive
//     sessions): wrap the executor with
//     codeexecutor.NewEnvInjectingCodeExecutor(exec, identity.EnvVarsFromContext).
//     Without this wrapper, Identity.EnvVars is resolved into context but
//     never reaches exec.Cmd.Env, so the env-injection half of the plugin is
//     a no-op. identity.EnvVarsFromContext has the same signature as
//     codeexecutor.RunEnvProvider, so no adapter is required.
//   - Custom tools implemented by the user: read identity via
//     identity.FromContext(ctx) inside the tool's Call implementation.
//
// Usage:
//
//	import (
//	    tmcp "trpc.group/trpc-go/trpc-mcp-go"
//	    "trpc.group/trpc-go/trpc-agent-go/codeexecutor"
//	    toolmcp "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
//	)
//
//	provider := identity.ProviderFunc(func(ctx context.Context, uid, sid string) (*identity.Identity, error) {
//	    return &identity.Identity{
//	        UserID: uid,
//	        Token:  generateToken(uid),
//	        Headers: map[string]string{"Authorization": "Bearer " + generateToken(uid)},
//	        EnvVars: map[string]string{"USER_ACCESS_TOKEN": generateToken(uid)},
//	    }, nil
//	})
//	// 1. Register the plugin so identity is resolved once per invocation.
//	runner := runner.NewRunner(
//	    "app",
//	    myAgent,
//	    runner.WithPlugins(identity.NewPlugin(provider)),
//	)
//	// 2. Wrap the executor so skill_run / workspace_exec receive EnvVars.
//	exec = codeexecutor.NewEnvInjectingCodeExecutor(exec, identity.EnvVarsFromContext)
//	// 3. For MCP HTTP transports, install a per-request hook that pulls
//	// identity headers from context and sets them on every outgoing request.
//	ts := toolmcp.NewMCPToolSet(cfg,
//	    toolmcp.WithMCPOptions(tmcp.WithHTTPBeforeRequest(
//	        func(ctx context.Context, req *http.Request) error {
//	            headers, err := identity.HeadersFromContext(ctx)
//	            if err != nil {
//	                return err
//	            }
//	            for k, v := range headers {
//	                req.Header.Set(k, v)
//	            }
//	            return nil
//	        },
//	    )),
//	)
type Plugin struct {
	name     string
	provider Provider
	opts     pluginOptions
}

// compile-time check.
var _ plugin.Plugin = (*Plugin)(nil)

type pluginOptions struct {
	argInjection bool // inject _identity key into all tool args. Default false.
}

// Option configures a Plugin.
type Option func(*pluginOptions)

// WithArgInjection enables injection of identity-related fields into generic
// tool JSON arguments (adds a reserved "_identity" key). If tool arguments
// already contain "_identity", injection is skipped to avoid clobbering
// caller-provided data. Default: false.
func WithArgInjection(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewPlugin creates a new identity Plugin with the given Provider.
//
// By default, NewPlugin only propagates identity through context. Tool
// arguments are left unchanged unless WithArgInjection(true) is enabled.
func NewPlugin(provider Provider, opts ...Option) *Plugin { _ = "STUB: not implemented"; return nil }

// NewNamedPlugin creates an identity Plugin with a custom name.
func NewNamedPlugin(name string, provider Provider, opts ...Option) *Plugin {
	_ = "STUB: not implemented"
	return nil
}

// Name implements plugin.Plugin.
func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (p *Plugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

// beforeAgent resolves the user identity and stores it in Invocation state.
func (p *Plugin) beforeAgent(
	ctx context.Context,
	args *agent.BeforeAgentArgs,
) (*agent.BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// beforeTool injects the resolved identity into the tool call context and,
// optionally, into tool arguments.
func (p *Plugin) beforeTool(
	ctx context.Context,
	args *tool.BeforeToolArgs,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tryInjectIdentityArg adds an "_identity" key to generic tool arguments.
func (p *Plugin) tryInjectIdentityArg(rawArgs []byte, id *Identity) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func decodeObjectArgs(rawArgs []byte) (map[string]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
