//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	"time"

	"gopkg.in/yaml.v3"

	"trpc.group/trpc-go/trpc-agent-go/internal/skillprofile"
	"trpc.group/trpc-go/trpc-agent-go/model"
	ocskills "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/skills"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/runtimeprofile"
)

const (
	openClawConfigEnvName = "OPENCLAW_CONFIG"

	defaultConfigRootDir = ".trpc-agent-go-github"
	defaultConfigAppDir  = "openclaw"
	defaultConfigFile    = "openclaw.yaml"
	defaultAdminAddr     = "127.0.0.1:19789"
	defaultAdminAutoPort = true

	sessionBackendInMemory   = "inmemory"
	sessionBackendRedis      = "redis"
	sessionBackendSQLite     = "sqlite"
	sessionBackendMySQL      = "mysql"
	sessionBackendPostgres   = "postgres"
	sessionBackendClickHouse = "clickhouse"

	memoryBackendFile      = "file"
	memoryBackendInMemory  = "inmemory"
	memoryBackendRedis     = "redis"
	memoryBackendSQLite    = "sqlite"
	memoryBackendSQLiteVec = "sqlitevec"
	memoryBackendMySQL     = "mysql"
	memoryBackendPostgres  = "postgres"
	memoryBackendPGVector  = "pgvector"

	summaryPolicyAny = "any"
	summaryPolicyAll = "all"

	summaryModeAuto   = "auto"
	summaryModeManual = "manual"

	defaultSessionSummaryEventThreshold = 20
	defaultSkillsLoadMode               = "turn"
	defaultSkillsToolProfile            = skillprofile.KnowledgeOnly
	defaultSkillsWatchDebounce          = 250 * time.Millisecond

	flagAddSessionSummary                             = "add-session-summary"
	flagEnableContextCompaction                       = "enable-context-compaction"
	flagContextCompactionOversizedToolResultMaxTokens = "context-compaction-oversized-tool-result-max-tokens"
	flagMaxHistoryRuns                                = "max-history-runs"
	flagPreloadMemory                                 = "preload-memory"

	flagAgentInstruction       = "agent-instruction"
	flagAgentInstructionFiles  = "agent-instruction-files"
	flagAgentInstructionDir    = "agent-instruction-dir"
	flagAgentSystemPrompt      = "agent-system-prompt"
	flagAgentSystemPromptFiles = "agent-system-prompt-files"
	flagAgentSystemPromptDir   = "agent-system-prompt-dir"

	flagAgentRalphLoopEnabled           = "agent-ralph-loop"
	flagAgentRalphLoopMaxIterations     = "agent-ralph-max-iterations"
	flagAgentRalphLoopCompletionPromise = "agent-ralph-completion-promise"
	flagAgentRalphLoopPromiseTagOpen    = "agent-ralph-promise-tag-open"
	flagAgentRalphLoopPromiseTagClose   = "agent-ralph-promise-tag-close"
	flagAgentRalphLoopVerifyCommand     = "agent-ralph-verify-command"
	flagAgentRalphLoopVerifyWorkDir     = "agent-ralph-verify-workdir"
	flagAgentRalphLoopVerifyTimeout     = "agent-ralph-verify-timeout"
	flagAgentRalphLoopVerifyEnv         = "agent-ralph-verify-env"

	flagEnableParallelTools = "enable-parallel-tools"

	flagSkillsAllowBundled  = "skills-allow-bundled"
	flagSkillsWatch         = "skills-watch"
	flagSkillsWatchBundled  = "skills-watch-bundled"
	flagSkillsWatchDebounce = "skills-watch-debounce"
	flagSkillsToolProfile   = "skills-tool-profile"
	flagSkillsLoadMode      = "skills-load-mode"
	flagSkillsMaxLoaded     = "skills-max-loaded"
	flagSkillsToolResults   = "skills-loaded-content-in-tool-results"
	flagSkillsSkipFallback  = "skills-skip-fallback-on-session-summary"

	flagDebugRecorder     = "debug-recorder"
	flagDebugRecorderDir  = "debug-recorder-dir"
	flagDebugRecorderMode = "debug-recorder-mode"

	flagAdminEnabled  = "admin-enabled"
	flagAdminAddr     = "admin-addr"
	flagAdminAutoPort = "admin-auto-port"

	flagA2AEnabled        = "a2a"
	flagA2AHost           = "a2a-host"
	flagA2AUserIDHeader   = "a2a-user-id-header"
	flagA2AStreaming      = "a2a-streaming"
	flagA2AAdvertiseTools = "a2a-advertise-tools"
	flagA2AName           = "a2a-name"
	flagA2ADescription    = "a2a-description"
)

type runOptions struct {
	ConfigPath string

	AppName  string
	HTTPAddr string

	AdminEnabled  bool
	AdminAddr     string
	AdminAutoPort bool

	LangfuseEnabled                      bool
	LangfuseRequired                     bool
	LangfuseUIBaseURL                    string
	LangfuseTraceURLTemplate             string
	LangfuseObservationLeafValueMaxBytes *int

	A2AEnabled        bool
	A2AHost           string
	A2AUserIDHeader   string
	A2AStreaming      bool
	A2AAdvertiseTools bool
	A2AName           string
	A2ADescription    string

	AddSessionSummary                             bool
	EnableContextCompaction                       bool
	ContextCompactionOversizedToolResultMaxTokens int
	MaxHistoryRuns                                int
	PreloadMemory                                 int

	AgentInstruction       string
	AgentInstructionFiles  string
	AgentInstructionDir    string
	AgentSystemPrompt      string
	AgentSystemPromptFiles string
	AgentSystemPromptDir   string

	AgentType string

	RalphLoopEnabled           bool
	RalphLoopMaxIterations     int
	RalphLoopCompletionPromise string
	RalphLoopPromiseTagOpen    string
	RalphLoopPromiseTagClose   string
	RalphLoopVerifyCommand     string
	RalphLoopVerifyWorkDir     string
	RalphLoopVerifyTimeout     time.Duration
	RalphLoopVerifyEnv         string

	ClaudeBin          string
	ClaudeOutputFormat string
	ClaudeExtraArgs    string
	ClaudeEnv          string
	ClaudeWorkDir      string

	ModelMode           string
	OpenAIModel         string
	OpenAIVariant       string
	OpenAIBaseURL       string
	GenerationConfig    *model.GenerationConfig
	ModelConfig         *yaml.Node
	KnowledgesConfig    []knowledgeEntry
	SkillsRoot          string
	SkillsExtraDir      string
	SkillsDebug         bool
	SkillsAllowBundled  string
	SkillConfigs        map[string]ocskills.SkillConfig
	SkillsWatch         bool
	SkillsWatchBundled  bool
	SkillsWatchDebounce time.Duration
	SkillsToolProfile   string
	SkillsLoadMode      string
	SkillsMaxLoaded     int
	SkillsToolResults   bool
	SkillsSkipFallback  bool
	SkillsToolingGuide  *string
	StateDir            string

	DebugRecorderEnabled bool
	DebugRecorderDir     string
	DebugRecorderMode    string

	AllowUsers      string
	RequireMention  bool
	Mention         string
	RuntimeProfiles *runtimeprofile.Config

	Channels []pluginSpec

	SessionBackend       string
	SessionRedisURL      string
	SessionRedisInstance string
	SessionRedisKeyPref  string
	SessionConfig        *yaml.Node

	MemoryBackend       string
	MemoryRedisURL      string
	MemoryRedisInstance string
	MemoryRedisKeyPref  string
	MemoryLimit         int
	MemoryConfig        *yaml.Node

	MemoryAutoEnabled          bool
	MemoryAutoPolicy           string
	MemoryAutoMessageThreshold int
	MemoryAutoTimeInterval     time.Duration

	SessionSummaryEnabled             bool
	SessionSummaryMode                string
	SessionSummaryPolicy              string
	SessionSummaryEventCount          int
	SessionSummaryTokenCount          int
	SessionSummaryIdleThreshold       time.Duration
	SessionSummaryMaxWords            int
	SessionSummaryApproxRunesPerToken float64

	EnableLocalExec      bool
	EnableOpenClawTools  bool
	OpenClawToolingGuide *string
	EnableParallelTools  bool

	enableOpenClawToolsExplicit bool

	ToolProviders []pluginSpec
	ToolSets      []pluginSpec

	RefreshToolSetsOnRun bool
}

func parseRunOptions(args []string) (runOptions, error) {
	_ = "STUB: not implemented"
	return *new(runOptions), nil
}

func unexpectedArgsError(args []string) error { _ = "STUB: not implemented"; return nil }

func resolveConfigPath(raw string) string { _ = "STUB: not implemented"; return "" }

func defaultConfigPathIfExists() string { _ = "STUB: not implemented"; return "" }

type fileConfig struct {
	AppName  *string `yaml:"app_name,omitempty"`
	StateDir *string `yaml:"state_dir,omitempty"`

	DebugRecorder *debugRecorderConfig `yaml:"debug_recorder,omitempty"`

	HTTP            *httpConfig            `yaml:"http,omitempty"`
	Admin           *adminConfig           `yaml:"admin,omitempty"`
	Observability   *observabilityConfig   `yaml:"observability,omitempty"`
	A2A             *a2aConfig             `yaml:"a2a,omitempty"`
	Agent           *agentRunConfig        `yaml:"agent,omitempty"`
	Model           *modelConfig           `yaml:"model,omitempty"`
	Knowledges      *knowledgesConfig      `yaml:"knowledges,omitempty"`
	Gateway         *gatewayConfig         `yaml:"gateway,omitempty"`
	RuntimeProfiles *runtimeprofile.Config `yaml:"runtime_profiles,omitempty"`
	Channels        []filePluginSpec       `yaml:"channels,omitempty"`
	Skills          *skillsConfig          `yaml:"skills,omitempty"`
	Tools           *toolsConfig           `yaml:"tools,omitempty"`

	Session *sessionConfig `yaml:"session,omitempty"`
	Memory  *memoryConfig  `yaml:"memory,omitempty"`
}

type httpConfig struct {
	Addr *string `yaml:"addr,omitempty"`
}

type adminConfig struct {
	Enabled  *bool   `yaml:"enabled,omitempty"`
	Addr     *string `yaml:"addr,omitempty"`
	AutoPort *bool   `yaml:"auto_port,omitempty"`
}

type observabilityConfig struct {
	Langfuse *langfuseConfig `yaml:"langfuse,omitempty"`
}

type langfuseConfig struct {
	Enabled                      *bool   `yaml:"enabled,omitempty"`
	Required                     *bool   `yaml:"required,omitempty"`
	UIBaseURL                    *string `yaml:"ui_base_url,omitempty"`
	TraceURLTemplate             *string `yaml:"trace_url_template,omitempty"`
	ObservationLeafValueMaxBytes *int    `yaml:"observation_leaf_value_max_bytes,omitempty"`
}

type a2aConfig struct {
	Enabled        *bool   `yaml:"enabled,omitempty"`
	Host           *string `yaml:"host,omitempty"`
	UserIDHeader   *string `yaml:"user_id_header,omitempty"`
	Streaming      *bool   `yaml:"streaming,omitempty"`
	AdvertiseTools *bool   `yaml:"advertise_tools,omitempty"`
	Name           *string `yaml:"name,omitempty"`
	Description    *string `yaml:"description,omitempty"`
}

type debugRecorderConfig struct {
	Enabled *bool   `yaml:"enabled,omitempty"`
	Dir     *string `yaml:"dir,omitempty"`
	Mode    *string `yaml:"mode,omitempty"`
}

type agentRunConfig struct {
	Type *string `yaml:"type,omitempty"`

	AddSessionSummary                             *bool `yaml:"add_session_summary,omitempty"`
	EnableContextCompaction                       *bool `yaml:"enable_context_compaction,omitempty"`
	ContextCompactionOversizedToolResultMaxTokens *int  `yaml:"context_compaction_oversized_tool_result_max_tokens,omitempty"`
	MaxHistoryRuns                                *int  `yaml:"max_history_runs,omitempty"`
	PreloadMemory                                 *int  `yaml:"preload_memory,omitempty"`

	Instruction      *string  `yaml:"instruction,omitempty"`
	InstructionFiles []string `yaml:"instruction_files,omitempty"`
	InstructionDir   *string  `yaml:"instruction_dir,omitempty"`

	SystemPrompt      *string  `yaml:"system_prompt,omitempty"`
	SystemPromptFiles []string `yaml:"system_prompt_files,omitempty"`
	SystemPromptDir   *string  `yaml:"system_prompt_dir,omitempty"`

	RalphLoop *ralphLoopConfig `yaml:"ralph_loop,omitempty"`

	ClaudeBin          *string  `yaml:"claude_bin,omitempty"`
	ClaudeOutputFormat *string  `yaml:"claude_output_format,omitempty"`
	ClaudeExtraArgs    []string `yaml:"claude_extra_args,omitempty"`
	ClaudeEnv          []string `yaml:"claude_env,omitempty"`
	ClaudeWorkDir      *string  `yaml:"claude_work_dir,omitempty"`
}

type ralphLoopConfig struct {
	Enabled           *bool   `yaml:"enabled,omitempty"`
	MaxIterations     *int    `yaml:"max_iterations,omitempty"`
	CompletionPromise *string `yaml:"completion_promise,omitempty"`
	PromiseTagOpen    *string `yaml:"promise_tag_open,omitempty"`
	PromiseTagClose   *string `yaml:"promise_tag_close,omitempty"`

	Verify *ralphLoopVerifyConfig `yaml:"verify,omitempty"`
}

type ralphLoopVerifyConfig struct {
	Command *string  `yaml:"command,omitempty"`
	WorkDir *string  `yaml:"work_dir,omitempty"`
	Timeout *string  `yaml:"timeout,omitempty"`
	Env     []string `yaml:"env,omitempty"`
}

type modelConfig struct {
	Mode             *string               `yaml:"mode,omitempty"`
	Name             *string               `yaml:"name,omitempty"`
	BaseURL          *string               `yaml:"base_url,omitempty"`
	OpenAIVariant    *string               `yaml:"openai_variant,omitempty"`
	GenerationConfig *generationConfigYAML `yaml:"generation_config,omitempty"`
	Config           *rawYAMLNode          `yaml:"config,omitempty"`
}

type generationConfigYAML struct {
	MaxTokens        *int     `yaml:"max_tokens,omitempty"`
	Temperature      *float64 `yaml:"temperature,omitempty"`
	TopP             *float64 `yaml:"top_p,omitempty"`
	Stream           *bool    `yaml:"stream,omitempty"`
	Stop             []string `yaml:"stop,omitempty"`
	PresencePenalty  *float64 `yaml:"presence_penalty,omitempty"`
	FrequencyPenalty *float64 `yaml:"frequency_penalty,omitempty"`
	ReasoningEffort  *string  `yaml:"reasoning_effort,omitempty"`
	ThinkingEnabled  *bool    `yaml:"thinking_enabled,omitempty"`
	ThinkingTokens   *int     `yaml:"thinking_tokens,omitempty"`
}

type gatewayConfig struct {
	AllowUsers      []string `yaml:"allow_users,omitempty"`
	RequireMention  *bool    `yaml:"require_mention,omitempty"`
	MentionPatterns []string `yaml:"mention_patterns,omitempty"`
}

type skillsConfig struct {
	Root      *string  `yaml:"root,omitempty"`
	ExtraDirs []string `yaml:"extra_dirs,omitempty"`
	Debug     *bool    `yaml:"debug,omitempty"`

	AllowBundled       []string `yaml:"allow_bundled,omitempty"`
	AllowBundledCamel  []string `yaml:"allowBundled,omitempty"`
	Watch              *bool    `yaml:"watch,omitempty"`
	WatchBundled       *bool    `yaml:"watch_bundled,omitempty"`
	WatchBundledCamel  *bool    `yaml:"watchBundled,omitempty"`
	WatchDebounceMS    *int     `yaml:"watch_debounce_ms,omitempty"`
	WatchDebounceCamel *int     `yaml:"watchDebounceMs,omitempty"`
	ToolProfile        *string  `yaml:"tool_profile,omitempty"`
	ToolProfileCamel   *string  `yaml:"toolProfile,omitempty"`
	LoadMode           *string  `yaml:"load_mode,omitempty"`
	LoadModeCamel      *string  `yaml:"loadMode,omitempty"`
	MaxLoadedSkills    *int     `yaml:"max_loaded_skills,omitempty"`
	MaxLoadedCamel     *int     `yaml:"maxLoadedSkills,omitempty"`

	ToolResults          *bool   `yaml:"loaded_content_in_tool_results,omitempty"`
	ToolResultsCamel     *bool   `yaml:"loadedContentInToolResults,omitempty"`
	SkipSummaryFallback  *bool   `yaml:"skip_fallback_on_session_summary,omitempty"`
	SkipFallbackCamel    *bool   `yaml:"skipFallbackOnSessionSummary,omitempty"`
	ToolingGuidance      *string `yaml:"tooling_guidance,omitempty"`
	ToolingGuidanceCamel *string `yaml:"toolingGuidance,omitempty"`

	Entries map[string]skillEntryConfig `yaml:"entries,omitempty"`
}

type skillEntryConfig struct {
	Enabled     *bool             `yaml:"enabled,omitempty"`
	APIKey      string            `yaml:"api_key,omitempty"`
	APIKeyCamel string            `yaml:"apiKey,omitempty"`
	Env         map[string]string `yaml:"env,omitempty"`
}

type toolsConfig struct {
	EnableLocalExec           *bool   `yaml:"enable_local_exec,omitempty"`
	EnableOpenClawTools       *bool   `yaml:"enable_openclaw_tools,omitempty"`
	OpenClawToolingGuide      *string `yaml:"openclaw_tooling_guidance,omitempty"`
	OpenClawToolingGuideCamel *string `yaml:"openClawToolingGuidance,omitempty"`
	EnableParallelTools       *bool   `yaml:"enable_parallel_tools,omitempty"`
	RefreshToolSetsOnRun      *bool   `yaml:"refresh_toolsets_on_run,omitempty"`

	Providers []filePluginSpec `yaml:"providers,omitempty"`
	ToolSets  []filePluginSpec `yaml:"toolsets,omitempty"`
}

type sessionConfig struct {
	Backend *string        `yaml:"backend,omitempty"`
	Redis   *redisConfig   `yaml:"redis,omitempty"`
	Summary *summaryConfig `yaml:"summary,omitempty"`
	Config  *rawYAMLNode   `yaml:"config,omitempty"`
}

type memoryConfig struct {
	Backend *string      `yaml:"backend,omitempty"`
	Redis   *redisConfig `yaml:"redis,omitempty"`
	Limit   *int         `yaml:"limit,omitempty"`
	Auto    *memoryAuto  `yaml:"auto,omitempty"`
	Config  *rawYAMLNode `yaml:"config,omitempty"`
}

type knowledgesConfig struct {
	Providers []knowledgeProviderConfig `yaml:"providers,omitempty"`

	// Entries is the deprecated field name (pre-v0.0.4). Kept here so
	// that KnownFields(true) does not reject it with a confusing
	// "field entries not found" error; instead we return a clear
	// migration message in fileConfig.apply.
	Entries []rawYAMLNode `yaml:"entries,omitempty"`
}

type knowledgeProviderConfig struct {
	Type        string       `yaml:"type,omitempty"`
	Name        string       `yaml:"name,omitempty"`
	Description string       `yaml:"description,omitempty"`
	MaxResults  *int         `yaml:"max_results,omitempty"`
	Config      *rawYAMLNode `yaml:"config,omitempty"`
}

type pluginSpec struct {
	Type   string     `yaml:"type,omitempty"`
	Name   string     `yaml:"name,omitempty"`
	Config *yaml.Node `yaml:"config,omitempty"`
}

type rawYAMLNode struct {
	Node *yaml.Node
}

func (r *rawYAMLNode) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

type filePluginSpec struct {
	Type   string       `yaml:"type,omitempty"`
	Name   string       `yaml:"name,omitempty"`
	Config *rawYAMLNode `yaml:"config,omitempty"`
}

type redisConfig struct {
	URL      *string `yaml:"url,omitempty"`
	Instance *string `yaml:"instance,omitempty"`
	KeyPref  *string `yaml:"key_prefix,omitempty"`
}

type summaryConfig struct {
	Enabled             *bool    `yaml:"enabled,omitempty"`
	Mode                *string  `yaml:"mode,omitempty"`
	Policy              *string  `yaml:"policy,omitempty"`
	EventThreshold      *int     `yaml:"event_threshold,omitempty"`
	TokenThreshold      *int     `yaml:"token_threshold,omitempty"`
	IdleThreshold       *string  `yaml:"idle_threshold,omitempty"`
	MaxWords            *int     `yaml:"max_words,omitempty"`
	ApproxRunesPerToken *float64 `yaml:"approx_runes_per_token,omitempty"`
}

type memoryAuto struct {
	Enabled          *bool   `yaml:"enabled,omitempty"`
	Policy           *string `yaml:"policy,omitempty"`
	MessageThreshold *int    `yaml:"message_threshold,omitempty"`
	TimeInterval     *string `yaml:"time_interval,omitempty"`
}

func loadConfigFile(path string) (*fileConfig, error) { _ = "STUB: not implemented"; return nil, nil }

func expandEnvPlaceholders(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func isValidEnvName(name string) bool { _ = "STUB: not implemented"; return false }

func (cfg *fileConfig) apply(
	opts *runOptions,
	set map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func applyLangfuseConfig(
	cfg *langfuseConfig,
	opts *runOptions,
) {
	_ = "STUB: not implemented"
	return
}

func applyRalphLoopConfig(
	cfg *ralphLoopConfig,
	opts *runOptions,
	set map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func convertPluginSpecs(specs []filePluginSpec) []pluginSpec { _ = "STUB: not implemented"; return nil }

func convertSkillConfigs(
	entries map[string]skillEntryConfig,
) map[string]ocskills.SkillConfig {
	_ = "STUB: not implemented"
	return nil
}

func convertKnowledgeConfigs(
	providers []knowledgeProviderConfig,
) ([]knowledgeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applySessionSummary(
	cfg *summaryConfig,
	opts *runOptions,
	set map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func applyMemoryAuto(
	cfg *memoryAuto,
	opts *runOptions,
	set map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func parseDuration(raw string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func flagWasSet(set map[string]struct{}, name string) bool { _ = "STUB: not implemented"; return false }

func skillsOptionExitCode(set map[string]struct{}) int { _ = "STUB: not implemented"; return 0 }

func finalizeRunOptions(opts *runOptions) error { _ = "STUB: not implemented"; return nil }

func normalizeA2AOptions(opts *runOptions) { _ = "STUB: not implemented"; return }

func normalizeA2AHost(raw string) string { _ = "STUB: not implemented"; return "" }

func normalizeSkillsToolProfile(raw string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func normalizeSkillsLoadMode(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func firstBoolPtr(primary, fallback *bool) *bool { _ = "STUB: not implemented"; return nil }

func firstIntPtr(primary, fallback *int) *int { _ = "STUB: not implemented"; return nil }

func firstStringPtr(primary, fallback *string) *string { _ = "STUB: not implemented"; return nil }

func resolveGenerationConfigYAML(
	cfg *generationConfigYAML,
) *model.GenerationConfig {
	_ = "STUB: not implemented"
	return nil
}

func trimStringPtr(v *string) *string { _ = "STUB: not implemented"; return nil }
