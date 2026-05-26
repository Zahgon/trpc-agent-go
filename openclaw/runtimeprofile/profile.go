//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package runtimeprofile defines per-request OpenClaw runtime profiles.
package runtimeprofile

import (
	"context"
	"encoding/json"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	runtimeStateKeyPrefix = "openclaw.profile."

	// ExtensionKey is the request extension key used to select a profile.
	ExtensionKey = "openclaw.runtime_profile"

	// RuntimeStateProfileID is the runtime-state key for the profile id.
	RuntimeStateProfileID = "openclaw.profile.id"
	// RuntimeStateProfileVersion is the runtime-state key for the profile
	// version.
	RuntimeStateProfileVersion = "openclaw.profile.version"
	// RuntimeStateWorkspaceWorkdir is the runtime-state key for the
	// profile workspace workdir.
	RuntimeStateWorkspaceWorkdir = "openclaw.profile.workspace.workdir"
	// RuntimeStateWorkspaceAllowedRoots is the runtime-state key for the
	// profile workspace allowed roots.
	RuntimeStateWorkspaceAllowedRoots = "openclaw.profile.workspace.roots"
	// RuntimeStateCredentialAllowedRefs is the runtime-state key for
	// profile credential references.
	RuntimeStateCredentialAllowedRefs = "openclaw.profile.credentials.refs"
	// RuntimeStateSkillInclude is the runtime-state key for profile skill
	// allowlists.
	RuntimeStateSkillInclude = "openclaw.profile.skills.include"
	// RuntimeStateSkillExclude is the runtime-state key for profile skill
	// denylists.
	RuntimeStateSkillExclude = "openclaw.profile.skills.exclude"
	// RuntimeStateSkillRoots is the runtime-state key for profile skill
	// repository roots.
	RuntimeStateSkillRoots = "openclaw.profile.skills.roots"
	// RuntimeStateKnowledgeIndexes is the runtime-state key for profile
	// knowledge indexes.
	RuntimeStateKnowledgeIndexes = "openclaw.profile.knowledge.indexes"
	// RuntimeStateToolSets is the runtime-state key for profile toolsets.
	RuntimeStateToolSets = "openclaw.profile.tools.toolsets"
	// RuntimeStateToolCredentialRefs is the runtime-state key for
	// tool credential references.
	RuntimeStateToolCredentialRefs = "openclaw.profile.tools.credentials"
	// RuntimeStateIsolationMode is the runtime-state key for the profile
	// isolation mode.
	RuntimeStateIsolationMode = "openclaw.profile.isolation.mode"
	// RuntimeStateIsolationAgentCache is the runtime-state key for
	// per-profile agent cache policy.
	RuntimeStateIsolationAgentCache = "openclaw.profile.isolation.agent_cache"
	// RuntimeStateIsolationToolSetCache is the runtime-state key for
	// per-profile toolset cache policy.
	RuntimeStateIsolationToolSetCache = "openclaw.profile.isolation." +
		"toolset_cache"
	// RuntimeStateIsolationServiceMode is the runtime-state key for
	// per-profile service/process policy.
	RuntimeStateIsolationServiceMode = "openclaw.profile.isolation.service"
)

// ErrProfileNotFound means a resolver could not find the selected profile.
var ErrProfileNotFound = errors.New("runtime profile not found")

// ErrProfileSelectorDenied means selector policy rejected profile resolution.
var ErrProfileSelectorDenied = errors.New("runtime profile selector denied")

// ErrConfigInvalid means a runtime profile config is internally inconsistent.
var ErrConfigInvalid = errors.New("runtime profile config invalid")

// ErrWorkspaceDenied means a path is outside the profile workspace policy.
var ErrWorkspaceDenied = errors.New("runtime profile workspace denied")

// ErrCredentialDenied means a credential ref is outside the profile policy.
var ErrCredentialDenied = errors.New("runtime profile credential denied")

type contextKey struct{}

type requestContextKey struct{}

// Config describes static runtime profiles.
type Config struct {
	Default           string             `yaml:"default,omitempty"`
	Required          bool               `yaml:"required,omitempty"`
	FallbackToDefault bool               `yaml:"fallback_to_default,omitempty"`
	Profiles          map[string]Profile `yaml:"profiles,omitempty"`
	Selectors         []Selector         `yaml:"selectors,omitempty"`
}

// Request describes one profile resolution request.
type Request struct {
	Channel    string
	ProfileID  string
	TenantID   string
	UserID     string
	SessionID  string
	RequestID  string
	Extensions map[string]json.RawMessage
}

// Extension is the normalized runtime-profile request extension.
type Extension struct {
	ProfileID string `json:"profile_id,omitempty"`
	TenantID  string `json:"tenant_id,omitempty"`
}

// Prompt defines prompt overrides for one run.
type Prompt struct {
	Instruction  string `yaml:"instruction,omitempty"`
	SystemPrompt string `yaml:"system_prompt,omitempty"`
}

// ToolPolicy defines name-based tool visibility and execution policy.
type ToolPolicy struct {
	Include          []string          `yaml:"include,omitempty"`
	Exclude          []string          `yaml:"exclude,omitempty"`
	ExecutionInclude []string          `yaml:"execution_include,omitempty"`
	ExecutionExclude []string          `yaml:"execution_exclude,omitempty"`
	ToolSets         []string          `yaml:"toolsets,omitempty"`
	CredentialRefs   map[string]string `yaml:"credential_refs,omitempty"`
}

// KnowledgePolicy defines per-run knowledge query policy.
type KnowledgePolicy struct {
	Indexes []string       `yaml:"indexes,omitempty"`
	Filter  map[string]any `yaml:"filter,omitempty"`
}

// WorkspacePolicy defines profile-scoped filesystem boundaries.
type WorkspacePolicy struct {
	Workdir      string   `yaml:"workdir,omitempty"`
	AllowedRoots []string `yaml:"allowed_roots,omitempty"`
}

// CredentialPolicy defines profile-scoped credential references.
type CredentialPolicy struct {
	AllowedRefs []string `yaml:"allowed_refs,omitempty"`
}

// SkillPolicy defines profile-scoped skill visibility and repositories.
type SkillPolicy struct {
	Include []string `yaml:"include,omitempty"`
	Exclude []string `yaml:"exclude,omitempty"`
	Roots   []string `yaml:"roots,omitempty"`
}

// IsolationMode describes how strongly a profile should be isolated.
type IsolationMode string

const (
	// IsolationModeShared uses the process-level agent and toolsets.
	IsolationModeShared IsolationMode = "shared"
	// IsolationModeProfileCache allows profile-keyed agent/toolset caches.
	IsolationModeProfileCache IsolationMode = "profile_cache"
	// IsolationModeService reserves a separate service/process boundary.
	IsolationModeService IsolationMode = "service"
)

// IsolationPolicy defines optional hard-isolation contracts.
type IsolationPolicy struct {
	Mode         IsolationMode `yaml:"mode,omitempty"`
	AgentCache   bool          `yaml:"agent_cache,omitempty"`
	ToolSetCache bool          `yaml:"toolset_cache,omitempty"`
	ServiceMode  string        `yaml:"service_mode,omitempty"`
}

// Profile is a resolved per-request runtime profile.
type Profile struct {
	ID          string           `yaml:"id,omitempty"`
	Version     string           `yaml:"version,omitempty"`
	AppName     string           `yaml:"app_name,omitempty"`
	AgentName   string           `yaml:"agent_name,omitempty"`
	ModelName   string           `yaml:"model_name,omitempty"`
	Prompt      Prompt           `yaml:"prompt,omitempty"`
	Tools       ToolPolicy       `yaml:"tools,omitempty"`
	Knowledge   KnowledgePolicy  `yaml:"knowledge,omitempty"`
	Workspace   WorkspacePolicy  `yaml:"workspace,omitempty"`
	Credentials CredentialPolicy `yaml:"credentials,omitempty"`
	Skills      SkillPolicy      `yaml:"skills,omitempty"`
	Isolation   IsolationPolicy  `yaml:"isolation,omitempty"`
	State       map[string]any   `yaml:"runtime_state,omitempty"`
	ExtraModel  map[string]any   `yaml:"model_request_extra,omitempty"`
}

// ValidateConfig validates profile keys, aliases, defaults, and isolation
// modes without assuming any OpenClaw application-specific agent names.
func ValidateConfig(cfg Config) error { _ = "STUB: not implemented"; return nil }

func profileIDAliases(
	profiles map[string]Profile,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addProfileIDAlias(
	aliases map[string]string,
	alias string,
	effectiveID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIsolationPolicy(policy IsolationPolicy) error { _ = "STUB: not implemented"; return nil }

// HasProfile returns true when profile carries any runtime contract.
func HasProfile(profile Profile) bool { _ = "STUB: not implemented"; return false }

// Resolver resolves one request to one runtime profile.
type Resolver interface {
	Resolve(ctx context.Context, req Request) (Profile, error)
}

// ResolverFunc adapts a function to Resolver.
type ResolverFunc func(ctx context.Context, req Request) (Profile, error)

// Resolve implements Resolver.
func (f ResolverFunc) Resolve(
	ctx context.Context,
	req Request,
) (Profile, error) {
	_ = "STUB: not implemented"
	return *new(Profile), nil
}

// WithRequest stores the profile selection request on the context.
func WithRequest(ctx context.Context, req Request) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// RequestFromContext returns the profile selection request stored on ctx.
func RequestFromContext(ctx context.Context) (Request, bool) {
	_ = "STUB: not implemented"
	return *new(Request), false
}

// WithProfile stores the resolved profile on the context.
func WithProfile(ctx context.Context, profile Profile) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ProfileFromContext returns the resolved profile stored on ctx.
func ProfileFromContext(ctx context.Context) (Profile, bool) {
	_ = "STUB: not implemented"
	return *new(Profile), false
}

func cloneRequest(req Request) Request { _ = "STUB: not implemented"; return *new(Request) }

// AppNameFromContext returns the profile app name or the provided fallback.
func AppNameFromContext(ctx context.Context, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

// TraceFields returns non-secret profile metadata suitable for logs/traces.
func TraceFields(profile Profile) map[string]any { _ = "STUB: not implemented"; return nil }

// MapResolver resolves profiles from an in-memory map.
type MapResolver struct {
	defaultID         string
	fallbackToDefault bool
	profiles          map[string]profileCacheKey
	cache             map[profileCacheKey]Profile
}

type profileCacheKey struct {
	id      string
	version string
}

// NewMapResolver creates a resolver backed by static profiles.
func NewMapResolver(cfg Config) *MapResolver { _ = "STUB: not implemented"; return nil }

// Resolve implements Resolver.
func (r *MapResolver) Resolve(
	ctx context.Context,
	req Request,
) (Profile, error) {
	_ = "STUB: not implemented"
	return *new(Profile), nil
}

// ProfileIDs implements Catalog.
func (r *MapResolver) ProfileIDs(context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppNames implements Catalog.
func (r *MapResolver) AppNames(context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtensionFromRequestExtensions reads the runtime-profile request extension.
func ExtensionFromRequestExtensions(
	extensions map[string]json.RawMessage,
) (Extension, bool, error) {
	_ = "STUB: not implemented"
	return *new(Extension), false, nil
}

// RunOptions converts a profile to agent run options.
func RunOptions(profile Profile) []agent.RunOption { _ = "STUB: not implemented"; return nil }

// RuntimeAppName returns the app name OpenClaw should use for a profile run.
func RuntimeAppName(profile Profile) string { _ = "STUB: not implemented"; return "" }

func requiresProfileRuntimeIsolation(policy IsolationPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func runtimeState(profile Profile) map[string]any { _ = "STUB: not implemented"; return nil }

func userRuntimeState(values map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func addRuntimeStateString(
	state map[string]any,
	key string,
	value string,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func addRuntimeStateStrings(
	state map[string]any,
	key string,
	values []string,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func addTraceString(fields map[string]any, key string, value string) {
	_ = "STUB: not implemented"
	return
}

func addTraceStrings(fields map[string]any, key string, values []string) {
	_ = "STUB: not implemented"
	return
}

func hasPrompt(prompt Prompt) bool { _ = "STUB: not implemented"; return false }

func hasToolPolicy(policy ToolPolicy) bool { _ = "STUB: not implemented"; return false }

func hasKnowledgePolicy(policy KnowledgePolicy) bool { _ = "STUB: not implemented"; return false }

func hasWorkspacePolicy(policy WorkspacePolicy) bool { _ = "STUB: not implemented"; return false }

func hasCredentialPolicy(policy CredentialPolicy) bool { _ = "STUB: not implemented"; return false }

func hasSkillPolicy(policy SkillPolicy) bool { _ = "STUB: not implemented"; return false }

func hasIsolationPolicy(policy IsolationPolicy) bool { _ = "STUB: not implemented"; return false }

func toolVisibilityFilter(
	toolPolicy ToolPolicy,
	knowledgePolicy KnowledgePolicy,
	credentialPolicy CredentialPolicy,
) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func toolPolicyFilter(policy ToolPolicy) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func knowledgeIndexFilter(indexes []string) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func toolCredentialFilter(
	refs map[string]string,
	allowedRefs []string,
) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func credentialRefForTool(
	tl tool.Tool,
	refs map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

type knowledgeIndexNamer interface {
	KnowledgeIndexName() string
}

func sourceKnowledgeIndex(tl tool.Tool) string { _ = "STUB: not implemented"; return "" }

func toolNamesFilter(include []string, exclude []string) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func toolSetNamesFilter(names []string) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

type toolSetNamer interface {
	ToolSetName() string
}

func sourceToolSetName(tl tool.Tool) string { _ = "STUB: not implemented"; return "" }

func allToolFilters(filters ...tool.FilterFunc) tool.FilterFunc {
	_ = "STUB: not implemented"
	return *new(tool.FilterFunc)
}

func nameSet(names []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func cleanStringMap(values map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func allowedCredentialRefs(
	refs map[string]string,
	allowedRefs []string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func cleanStrings(values []string) []string { _ = "STUB: not implemented"; return nil }
