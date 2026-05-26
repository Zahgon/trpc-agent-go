//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package mcpbroker provides opt-in MCP discovery/call broker tools
// (mcp_list_servers, mcp_list_tools, mcp_inspect_tools, mcp_call) that an LLM
// agent can use to discover and invoke remote MCP servers at runtime.
//
// # Untrusted remote content
//
// The broker is a passthrough: it forwards remote MCP server output - tool names
// and descriptions, tool result Content, error messages - verbatim to the model.
// Remote content is therefore untrusted from the host's perspective and may be
// crafted to attempt prompt injection, tool-name spoofing, or data exfiltration
// via the model's tool-use loop. The broker intentionally does not sanitise this
// content because removing it would also remove the information the agent
// legitimately needs to reason about the remote server.
//
// Hosts that connect to MCP servers outside their trust domain - especially when
// ad-hoc HTTP is enabled via WithAllowAdHocHTTP - should:
//
//   - Constrain which servers or URLs the agent can reach via WithServers,
//     WithAllowAdHocHTTP, and WithClientOptionsProvider.
//   - Treat remote tool output Content as untrusted input when composing
//     downstream prompts and tool chains.
//   - Use WithErrorInterceptor to redact internal-topology details from MCP
//     server errors that would otherwise reach the model.
package mcpbroker

import (
	"context"
	"errors"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
	mcpcfg "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
	tmcp "trpc.group/trpc-go/trpc-mcp-go"
)

// ErrClientOptionsProviderPanicked is the sentinel returned (wrapped via fmt.Errorf %w) when a
// ClientOptionsProvider panics. The broker recovers the panic so it never unwinds through
// trpc-mcp-go internals; callers can detect this case with errors.Is for observability or to
// fall back to default options. The recovered panic value is included in the wrapped error
// message via %v.
var ErrClientOptionsProviderPanicked = errors.New("ClientOptionsProvider panicked")

const (
	listServersToolName  = "mcp_list_servers"
	listToolsToolName    = "mcp_list_tools"
	inspectToolsToolName = "mcp_inspect_tools"
	callToolToolName     = "mcp_call"
)

const (
	originCode = "code"
)

const (
	targetTypeStdio = "stdio"
	targetTypeHTTP  = "http"
)

// Exported target type strings match ClientOptionsRequest.TargetType after resolution.
const (
	// TargetTypeHTTP indicates SSE or streamable HTTP MCP transport.
	TargetTypeHTTP = targetTypeHTTP
	// TargetTypeStdio indicates stdio MCP transport.
	TargetTypeStdio = targetTypeStdio
)

// Origin values for ClientOptionsRequest.Origin.
const (
	// OriginCode indicates a named server from WithServers.
	OriginCode = originCode
	// OriginAdhoc indicates an ad-hoc URL selector (requires WithAllowAdHocHTTP).
	OriginAdhoc = "adhoc"
)

var defaultSensitiveHeaderDenylist = []string{
	"authorization",
	"proxy-authorization",
	"cookie",
	"set-cookie",
	"x-api-key",
}

// Broker provides a small set of MCP management tools for agents.
type Broker struct {
	options brokerOptions
}

// Option configures a Broker.
type Option func(*brokerOptions)

type brokerOptions struct {
	servers                     map[string]mcpcfg.ConnectionConfig
	allowAdHocHTTP              bool
	adhocHTTPTimeout            time.Duration
	adhocSensitiveHeaderDenyset map[string]struct{}
	httpHeaderInjector          HTTPHeaderInjector
	clientOptionsProvider       ClientOptionsProvider
	errorInterceptor            ErrorInterceptor
}

// ClientOptionsProvider supplies extra trpc-mcp-go client options after the broker has resolved
// the target and merged injected HTTP headers, but before the underlying MCP client is created.
// Returning (nil, nil) applies no extra options.
//
// # Trust boundary
//
// The provider is host-trusted code registered via WithClientOptionsProvider; it is not
// model-controlled or end-user-controlled input. The broker therefore does NOT subject the
// returned options to the ad-hoc header denylist or any other sanitisation. By design, the
// provider can override Authorization or any other header previously set by
// WithHTTPHeaderInjector - hosts rely on this for service-account auth, request signing, and
// audit hooks. The flip side is that a provider that accidentally shadows the injector's
// Authorization is a host bug, not a framework vulnerability.
//
// # Failure modes
//
//   - If the provider returns an error, the broker aborts before establishing the MCP connection.
//   - If the provider panics, the broker recovers and surfaces an error wrapping
//     ErrClientOptionsProviderPanicked (use errors.Is to detect this case) instead of unwinding
//     through trpc-mcp-go internals.
//   - nil entries inside the returned ClientOptions.HTTP / .Stdio slices are silently filtered out
//     so a conditional option builder that yields nil does not crash trpc-mcp-go.
//
// # Caveats
//
// HTTP-layer retries (e.g. tmcp.WithRetry) configured here are not aware of MCP-level side
// effects. Retrying tools/call automatically can cause duplicated side effects; hosts that need
// retry semantics for tool calls should implement them above the broker, gated by idempotency.
type ClientOptionsProvider func(context.Context, *ClientOptionsRequest) (*ClientOptions, error)

// ClientOptionsRequest is the broker-side context passed to ClientOptionsProvider before creating
// the underlying MCP client. Config is a defensive clone of the final ConnectionConfig for this
// call (see cloneConnectionConfig); mutations to Config (including its Headers map and Args slice)
// do not propagate back to the broker or the actual request.
type ClientOptionsRequest struct {
	Selector   string
	ServerName string
	Origin     string // OriginCode or OriginAdhoc
	TargetType string // TargetTypeHTTP or TargetTypeStdio

	Transport string
	BaseURL   string
	ToolName  string
	Phase     string // PhaseListTools | PhaseInspectTools | PhaseCallTool

	Config mcpcfg.ConnectionConfig
}

// ClientOptions carries optional HTTP and stdio client options. HTTP applies to SSE and streamable
// transports; Stdio applies to stdio transports.
//
// Default broker headers are applied first; options here follow and may override or extend
// behavior intentionally (see trpc-mcp-go ClientOption / StdioClientOption). For example, a host
// can use tmcp.WithHTTPBeforeRequest to rewrite an Authorization header set earlier by
// WithHTTPHeaderInjector. nil entries in HTTP / Stdio are filtered out before being applied.
type ClientOptions struct {
	HTTP  []tmcp.ClientOption
	Stdio []tmcp.StdioClientOption
}

// HTTPHeaderInjector resolves per-run HTTP headers for MCP requests. The broker
// calls it after target resolution and before WithClientOptionsProvider, and
// merges the returned headers into the outgoing request.
//
// # Trust BaseURL with care for ad-hoc targets
//
// When HeaderInjectRequest.IsAdHoc is true, BaseURL originates from a
// model-supplied URL and may have been chosen by an attacker via prompt
// injection. Returning a host secret (OAuth token, session cookie, vendor API
// key) here without first cross-checking that BaseURL.Host belongs to a trusted
// destination can exfiltrate that secret to whatever endpoint the model picks.
// The recommended pattern is:
//
//   - For named servers (IsAdHoc == false): inject the secret bound to that
//     server name.
//   - For ad-hoc targets (IsAdHoc == true): either skip injection, or only
//     inject when the parsed BaseURL.Host is on a host-defined allowlist.
type HTTPHeaderInjector func(context.Context, *HeaderInjectRequest) (map[string]string, error)

// HeaderInjectRequest describes an MCP HTTP request before execution. IsAdHoc
// distinguishes trusted named-server targets from model-supplied ad-hoc URLs;
// see HTTPHeaderInjector for the implications when injecting secrets.
type HeaderInjectRequest struct {
	Selector  string
	BaseURL   string
	ToolName  string
	Phase     string
	Transport string
	IsAdHoc   bool
}

// ErrorInterceptor lets business code classify and transform MCP HTTP errors.
type ErrorInterceptor func(context.Context, *BrokerErrorRequest) (*BrokerErrorDecision, error)

// BrokerErrorRequest describes an MCP HTTP operation error.
type BrokerErrorRequest struct {
	Selector  string
	BaseURL   string
	ToolName  string
	Phase     string
	Transport string
	IsAdHoc   bool
	Err       error
}

// BrokerErrorDecision controls how an intercepted error is surfaced.
type BrokerErrorDecision struct {
	WrapError error
	Handled   bool
}

// New creates a new MCP broker.
func New(opts ...Option) *Broker { _ = "STUB: not implemented"; return nil }

// WithServers adds named MCP server configurations provided by code.
func WithServers(servers map[string]mcpcfg.ConnectionConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAllowAdHocHTTP controls whether ad-hoc HTTP MCP targets are allowed.
//
// When enabled, mcp_list_tools / mcp_inspect_tools / mcp_call accept model-supplied
// URLs (selectors that look like an http:// or https:// URL) instead of named
// servers from WithServers. This is a powerful surface and a risky one: the broker
// only enforces the http/https scheme and the ad-hoc header denylist (see
// WithAdHocSensitiveHeaderDenylist); it does NOT validate destination identity,
// host policy, or reachability.
//
// # SSRF risk
//
// Without further controls, a sufficiently-instructed model can ask the broker to
// connect to internal addresses (localhost, link-local, cloud metadata endpoints,
// intranet services), exfiltrate data to attacker-controlled hosts, or probe
// internal networks. URL allow/deny policy is the host's responsibility.
//
// # Recommended pattern
//
// Use WithClientOptionsProvider together with tmcp.WithHTTPBeforeRequest to enforce
// a per-deployment URL policy. The provider receives
// ClientOptionsRequest.Origin == OriginAdhoc for model-supplied URLs, so a host can
// reject or rewrite them before the connection is established.
//
// # Hang risk and bounding
//
// trpc-mcp-go's underlying http.Client has no built-in request timeout, so an
// ad-hoc MCP call against a stalled or hostile remote can hang for as long as the
// caller-supplied context permits. When that context has no deadline (common in
// long-running agent runs or background jobs), the call can hang indefinitely and
// occupy a goroutine plus a TCP connection until the OS gives up. Hosts that
// enable ad-hoc HTTP should bound it explicitly via WithAdHocHTTPTimeout, or via
// WithClientOptionsProvider for finer-grained per-call policy.
func WithAllowAdHocHTTP(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAdHocHTTPTimeout sets the upper bound on the total wall-clock time the
// broker spends on a single ad-hoc HTTP MCP operation (mcp_list_tools,
// mcp_inspect_tools, or mcp_call). The bound covers Initialize and every MCP
// RPC the operation issues; the broker wraps the entire operation in
// context.WithTimeout once, taking whichever deadline is tighter between this
// value and the caller-supplied context. An ad-hoc target is one resolved
// from an http/https URL selector rather than a name configured via
// WithServers.
//
// # Why this option is opt-in
//
// trpc-mcp-go's underlying http.Client has no built-in request timeout, so an
// ad-hoc MCP call against a stalled or hostile remote can hang indefinitely
// when the caller-supplied context has no deadline. Because the appropriate
// bound is workload-specific (short tool calls might warrant a few seconds;
// analytical tools may need minutes), the broker exposes this option instead
// of imposing a framework-wide default that would be wrong for many
// deployments.
//
// # Values
//
//   - d > 0: apply d as the per-operation deadline upper bound for ad-hoc
//     HTTP calls. The clock starts before Initialize and stops after the
//     last MCP RPC needed for the operation.
//   - d <= 0: no broker-level deadline (default). The operation then relies
//     entirely on the caller-supplied context deadline; without one, a
//     stalled remote can hang the operation indefinitely. Only choose this
//     when an upstream deadline source is guaranteed.
//
// # Scope
//
// This option only affects ad-hoc HTTP operations. Named servers configured
// via WithServers continue to use their own ConnectionConfig.Timeout, which
// is not modified by this option.
func WithAdHocHTTPTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAdHocSensitiveHeaderDenylist adds case-insensitive header names that the
// broker must reject when the model supplies them via the "headers" parameter of
// mcp_list_tools / mcp_inspect_tools / mcp_call for an ad-hoc HTTP MCP target.
// Comparison is performed against the lowercase name; entries here extend (never
// shrink) the built-in default.
//
// # Default coverage
//
// The built-in default covers the protocol-standard auth headers - Authorization,
// Proxy-Authorization, Cookie, Set-Cookie - plus X-API-Key, a de-facto industry
// standard across cloud providers and API gateways. These are intentionally
// narrow: every host will agree the model should not be able to spell them.
//
// # Vendor-specific headers are the host's responsibility
//
// Many ecosystems use non-standard token-bearing headers - OpenStack uses
// X-Auth-Token, AWS uses X-Amz-Security-Token, GitLab uses Private-Token, and
// many internal systems invent their own. The broker deliberately does NOT add
// these to the default list. Doing so would silently reject legitimate model
// traffic in deployments that do not use those vendors, and would create a false
// sense of safety in deployments that do (an attacker can always rename the
// exfiltration header). Hosts whose MCP servers - or whose neighbouring services
// reachable on the same network - rely on such headers should pass them in via
// this option.
func WithAdHocSensitiveHeaderDenylist(headers []string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHTTPHeaderInjector injects per-run HTTP headers derived from context and target metadata.
func WithHTTPHeaderInjector(fn HTTPHeaderInjector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClientOptionsProvider registers a hook that returns extra trpc-mcp-go client options after
// resolveTarget and WithHTTPHeaderInjector. Use this for URL policy, auditing, custom HTTP
// clients, or stdio options without adding parallel WithX hooks for each concern. See
// ClientOptionsProvider for the trust boundary, failure modes, and the HTTP-retry caveat.
func WithClientOptionsProvider(fn ClientOptionsProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithErrorInterceptor intercepts HTTP MCP execution errors and may translate them for model consumption.
func WithErrorInterceptor(fn ErrorInterceptor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Tools returns the broker's MCP management tools.
func (b *Broker) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

type namedServer struct {
	Name       string
	Origin     string
	TargetType string
	Config     mcpcfg.ConnectionConfig
}

func (b *Broker) resolveNamedServers() ([]namedServer, map[string]namedServer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (b *Broker) resolveTarget(input targetInput) (resolvedTarget, error) {
	_ = "STUB: not implemented"
	return *new(resolvedTarget), nil
}

func (b *Broker) listServers(ctx context.Context, _ listServersInput) (listServersOutput, error) {
	_ = "STUB: not implemented"
	return *new(listServersOutput), nil
}

func (b *Broker) resolveListSelector(
	selector string,
	transport string,
	headers map[string]string,
) (resolvedTarget, string, error) {
	_ = "STUB: not implemented"
	return *new(resolvedTarget), "", nil
}

func (b *Broker) resolveCallSelector(
	selector string,
	transport string,
	headers map[string]string,
) (resolvedTarget, string, string, error) {
	_ = "STUB: not implemented"
	return *new(resolvedTarget), "", "", nil
}

func (b *Broker) splitNamedToolSelector(selector string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func splitHTTPToolSelector(selector string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func shouldUseFragmentHTTPToolSelector(selector string) bool {
	_ = "STUB: not implemented"
	return false
}

func looksLikeHTTPSelector(selector string) bool { _ = "STUB: not implemented"; return false }
