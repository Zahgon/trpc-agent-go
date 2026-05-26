//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package admin

import (
	"html/template"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/cron"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/octool"
	ocskills "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/skills"
)

const (
	routeIndex      = "/"
	routeOverview   = "/overview"
	routeSkillsPage = "/skills"
	routePrompts    = "/prompts"
	routeIdentity   = "/identity"
	routePersonas   = "/personas"
	routeChats      = "/chats"
	routeMemory     = "/memory"
	routeAutomation = "/automation"
	routeSessions   = "/sessions"
	routeDebug      = "/debug"
	routeBrowser    = "/browser"

	routeStatusJSON        = "/api/status"
	routeSkillsJSON        = "/api/skills/status"
	routeSkillsRefresh     = "/api/skills/refresh"
	routeSkillToggle       = "/api/skills/toggle"
	routeMemoryFilesJSON   = "/api/memory/files"
	routeMemoryFileAPI     = "/api/memory/file"
	routeMemoryFile        = "/memory/file"
	routeJobsJSON          = "/api/cron/jobs"
	routeJobRun            = "/api/cron/jobs/run"
	routeJobRemove         = "/api/cron/jobs/remove"
	routeJobsClear         = "/api/cron/jobs/clear"
	routeExecSessionsJSON  = "/api/exec/sessions"
	routeUploadsJSON       = "/api/uploads"
	routeUploadSessions    = "/api/uploads/sessions"
	routeUploadFile        = "/uploads/file"
	routeDebugSessionsJSON = "/api/debug/sessions"
	routeDebugTracesJSON   = "/api/debug/traces"
	routeDebugFile         = "/debug/file"
	routePageStateJSON     = "/api/page/state"

	queryNotice    = "notice"
	queryError     = "error"
	queryChatID    = "chat_id"
	querySessionID = "session_id"
	queryChannel   = "channel"
	queryUserID    = "user_id"
	queryKind      = "kind"
	queryMimeType  = "mime_type"
	querySource    = "source"
	queryTrace     = "trace"
	queryName      = "name"
	queryPath      = "path"
	queryDownload  = "download"
	queryCursor    = "cursor"
	queryView      = "view"
	formJobID      = "job_id"
	formSkillKey   = "skill_key"
	formSkillName  = "skill_name"
	formEnabled    = "enabled"
	formReturnTo   = "return_to"
	formReturnPath = "return_path"

	refreshSeconds = 15

	debugBySessionDir   = "by-session"
	debugMetaFileName   = "meta.json"
	debugEventsFileName = "events.jsonl"
	debugResultFileName = "result.json"
	debugEventsMIMEType = "application/x-ndjson; charset=utf-8"

	maxDebugSessionRows = 12
	maxDebugTraceRows   = 18
	maxJobOutputRunes   = 120
	browserProbeTimeout = 1500 * time.Millisecond

	formatTimeLayout = "2006-01-02 15:04:05 MST"

	adminBrandName     = "TRPC-CLAW"
	adminBrandTitle    = "TRPC-CLAW admin"
	adminRuntimePrefix = "trpc-claw"

	pageSummaryPrompts = "" +
		"Edit the main prompt blocks, inspect the assembled " +
		"prompt previews, and keep file-level edits in one place."
	pageSummaryIdentity = "" +
		"Set the default name, keep current-chat names readable, " +
		"and leave the runtime product as a separate read-only fact."
	pageSummaryPersonas = "" +
		"Manage the default persona and any file-backed " +
		"persona definitions exposed by this runtime."
	pageSummaryChats = "" +
		"Inspect each chat's current state, recent transcript, " +
		"and the safest next step for names and persona."
)

type adminView string

const (
	viewOverview   adminView = "overview"
	viewSkills     adminView = "skills"
	viewPrompts    adminView = "prompts"
	viewIdentity   adminView = "identity"
	viewPersonas   adminView = "personas"
	viewChats      adminView = "chats"
	viewMemory     adminView = "memory"
	viewAutomation adminView = "automation"
	viewSessions   adminView = "sessions"
	viewDebug      adminView = "debug"
	viewBrowser    adminView = "browser"
)

type Routes struct {
	HealthPath   string
	MessagesPath string
	StatusPath   string
	CancelPath   string
}

type Config struct {
	AppName    string
	InstanceID string
	StartedAt  time.Time
	Hostname   string
	PID        int
	GoVersion  string

	AgentType      string
	ModelMode      string
	ModelName      string
	SessionBackend string
	MemoryBackend  string

	GatewayAddr   string
	GatewayURL    string
	AdminAddr     string
	AdminURL      string
	AdminAutoPort bool
	Langfuse      LangfuseStatus

	StateDir string
	DebugDir string

	Channels         []string
	GatewayRoutes    Routes
	Skills           SkillsStatusProvider
	Prompts          PromptsProvider
	Identity         IdentityProvider
	Personas         PersonasProvider
	Chats            ChatsProvider
	MemoryFiles      MemoryFileStore
	MemoryUserLabels MemoryUserLabelResolver
	Browser          BrowserConfig

	Cron *cron.Service
	Exec *octool.Manager
}

type BrowserConfig struct {
	Providers []BrowserProvider            `json:"providers,omitempty"`
	Managed   BrowserManagedStatusProvider `json:"-"`
}

type BrowserProvider struct {
	Name             string           `json:"name,omitempty"`
	DefaultProfile   string           `json:"default_profile,omitempty"`
	EvaluateEnabled  bool             `json:"evaluate_enabled"`
	HostServerURL    string           `json:"host_server_url,omitempty"`
	SandboxServerURL string           `json:"sandbox_server_url,omitempty"`
	AllowLoopback    bool             `json:"allow_loopback"`
	AllowPrivateNet  bool             `json:"allow_private_networks"`
	AllowFileURLs    bool             `json:"allow_file_urls"`
	Profiles         []BrowserProfile `json:"profiles,omitempty"`
	Nodes            []BrowserNode    `json:"nodes,omitempty"`
}

type BrowserProfile struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	Transport        string `json:"transport,omitempty"`
	ServerURL        string `json:"server_url,omitempty"`
	BrowserServerURL string `json:"browser_server_url,omitempty"`
}

type BrowserNode struct {
	ID        string `json:"id,omitempty"`
	ServerURL string `json:"server_url,omitempty"`
}

type BrowserManagedStatusProvider interface {
	BrowserManagedStatus() BrowserManagedService
}

type SkillsStatusProvider interface {
	SkillsStatus() (ocskills.StatusReport, error)
	SkillsConfigPath() string
	SkillsRefreshable() bool
	RefreshSkills() error
	SetSkillEnabled(configKey string, enabled bool) error
}

type BrowserManagedService struct {
	Enabled         bool       `json:"enabled"`
	Managed         bool       `json:"managed"`
	State           string     `json:"state,omitempty"`
	URL             string     `json:"url,omitempty"`
	PID             int        `json:"pid,omitempty"`
	WorkDir         string     `json:"work_dir,omitempty"`
	Command         string     `json:"command,omitempty"`
	LogPath         string     `json:"log_path,omitempty"`
	LogRelativePath string     `json:"log_relative_path,omitempty"`
	LogURL          string     `json:"log_url,omitempty"`
	StartedAt       *time.Time `json:"started_at,omitempty"`
	StoppedAt       *time.Time `json:"stopped_at,omitempty"`
	ExitCode        *int       `json:"exit_code,omitempty"`
	LastError       string     `json:"last_error,omitempty"`
	RecentLogs      []string   `json:"recent_logs,omitempty"`
}

type Service struct {
	cfg               Config
	runtimeConfig     RuntimeConfigProvider
	runtimeLifecycle  RuntimeLifecycleProvider
	now               func() time.Time
	browserHTTPClient *http.Client
}

type Option func(*Service)

func WithClock(fn func() time.Time) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBrowserHTTPClient(client *http.Client) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRuntimeConfigProvider(provider RuntimeConfigProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRuntimeLifecycleProvider(
	provider RuntimeLifecycleProvider,
) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(cfg Config, opts ...Option) *Service { _ = "STUB: not implemented"; return nil }

func (s *Service) Handler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

type snapshot struct {
	GeneratedAt time.Time `json:"generated_at"`

	AppName    string    `json:"app_name,omitempty"`
	InstanceID string    `json:"instance_id,omitempty"`
	StartedAt  time.Time `json:"started_at,omitempty"`
	Hostname   string    `json:"hostname,omitempty"`
	PID        int       `json:"pid,omitempty"`
	GoVersion  string    `json:"go_version,omitempty"`
	Uptime     string    `json:"uptime,omitempty"`

	AgentType      string `json:"agent_type,omitempty"`
	ModelMode      string `json:"model_mode,omitempty"`
	ModelName      string `json:"model_name,omitempty"`
	SessionBackend string `json:"session_backend,omitempty"`
	MemoryBackend  string `json:"memory_backend,omitempty"`

	GatewayAddr   string         `json:"gateway_addr,omitempty"`
	GatewayLabel  string         `json:"gateway_label,omitempty"`
	GatewayURL    string         `json:"gateway_url,omitempty"`
	AdminAddr     string         `json:"admin_addr,omitempty"`
	AdminURL      string         `json:"admin_url,omitempty"`
	AdminAutoPort bool           `json:"admin_auto_port"`
	Langfuse      LangfuseStatus `json:"langfuse"`

	StateDir string `json:"state_dir,omitempty"`
	DebugDir string `json:"debug_dir,omitempty"`

	Channels []string      `json:"channels,omitempty"`
	Routes   Routes        `json:"routes,omitempty"`
	Browser  browserStatus `json:"browser"`
	Skills   skillsStatus  `json:"skills"`
	Memory   memoryStatus  `json:"memory"`
	Exec     execStatus    `json:"exec"`
	Uploads  uploadsStatus `json:"uploads"`
	Cron     cronStatus    `json:"cron"`
	Debug    debugStatus   `json:"debug"`
}

type browserStatus struct {
	Enabled       bool                  `json:"enabled"`
	ProviderCount int                   `json:"provider_count"`
	ProfileCount  int                   `json:"profile_count"`
	NodeCount     int                   `json:"node_count"`
	Managed       BrowserManagedService `json:"managed,omitempty"`
	Providers     []browserProviderView `json:"providers,omitempty"`
}

type browserProviderView struct {
	Name             string               `json:"name,omitempty"`
	DefaultProfile   string               `json:"default_profile,omitempty"`
	EvaluateEnabled  bool                 `json:"evaluate_enabled"`
	HostServerURL    string               `json:"host_server_url,omitempty"`
	SandboxServerURL string               `json:"sandbox_server_url,omitempty"`
	AllowLoopback    bool                 `json:"allow_loopback"`
	AllowPrivateNet  bool                 `json:"allow_private_networks"`
	AllowFileURLs    bool                 `json:"allow_file_urls"`
	Host             browserEndpointView  `json:"host,omitempty"`
	Sandbox          browserEndpointView  `json:"sandbox,omitempty"`
	Profiles         []browserProfileView `json:"profiles,omitempty"`
	Nodes            []browserNodeView    `json:"nodes,omitempty"`
}

type browserProfileView struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	Transport        string `json:"transport,omitempty"`
	ServerURL        string `json:"server_url,omitempty"`
	BrowserServerURL string `json:"browser_server_url,omitempty"`
}

type browserEndpointView struct {
	URL       string               `json:"url,omitempty"`
	Reachable bool                 `json:"reachable"`
	Error     string               `json:"error,omitempty"`
	Profiles  []browserRemoteProbe `json:"profiles,omitempty"`
}

type browserRemoteProbe struct {
	Name   string `json:"name,omitempty"`
	State  string `json:"state,omitempty"`
	Driver string `json:"driver,omitempty"`
	Tabs   int    `json:"tabs,omitempty"`
}

type browserNodeView struct {
	ID        string              `json:"id,omitempty"`
	ServerURL string              `json:"server_url,omitempty"`
	Status    browserEndpointView `json:"status,omitempty"`
}

type cronStatus struct {
	Enabled     bool      `json:"enabled"`
	JobCount    int       `json:"job_count"`
	RunningJobs int       `json:"running_jobs"`
	Channels    []string  `json:"channels,omitempty"`
	Jobs        []jobView `json:"jobs,omitempty"`
}

type jobView struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	UserID   string `json:"user_id,omitempty"`

	Channel string `json:"channel,omitempty"`
	Target  string `json:"target,omitempty"`

	MessagePreview string `json:"message_preview,omitempty"`
	LastOutput     string `json:"last_output,omitempty"`

	Enabled    bool       `json:"enabled"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	LastRunAt  *time.Time `json:"last_run_at,omitempty"`
	NextRunAt  *time.Time `json:"next_run_at,omitempty"`
	LastStatus string     `json:"last_status,omitempty"`
	LastError  string     `json:"last_error,omitempty"`
}

type skillsStatus struct {
	Enabled         bool              `json:"enabled"`
	Error           string            `json:"error,omitempty"`
	Writable        bool              `json:"writable"`
	Refreshable     bool              `json:"refreshable"`
	ConfigPath      string            `json:"config_path,omitempty"`
	TotalCount      int               `json:"total_count"`
	ReadyCount      int               `json:"ready_count"`
	NeedsSetupCount int               `json:"needs_setup_count"`
	DisabledCount   int               `json:"disabled_count"`
	BundledCount    int               `json:"bundled_count"`
	Groups          []skillsGroupView `json:"groups,omitempty"`
}

type skillsGroupView struct {
	ID     string      `json:"id,omitempty"`
	Label  string      `json:"label,omitempty"`
	Skills []skillView `json:"skills,omitempty"`
}

type skillView struct {
	Name               string                `json:"name,omitempty"`
	Description        string                `json:"description,omitempty"`
	SkillKey           string                `json:"skill_key,omitempty"`
	ConfigKey          string                `json:"config_key,omitempty"`
	FilePath           string                `json:"file_path,omitempty"`
	Source             string                `json:"source,omitempty"`
	Reason             string                `json:"reason,omitempty"`
	Emoji              string                `json:"emoji,omitempty"`
	Homepage           string                `json:"homepage,omitempty"`
	PrimaryEnv         string                `json:"primary_env,omitempty"`
	Status             string                `json:"status,omitempty"`
	SearchText         string                `json:"search_text,omitempty"`
	Bundled            bool                  `json:"bundled"`
	Always             bool                  `json:"always"`
	Disabled           bool                  `json:"disabled"`
	Eligible           bool                  `json:"eligible"`
	BlockedByAllowlist bool                  `json:"blocked_by_allowlist"`
	Requirements       skillRequirementsView `json:"requirements,omitempty"`
	Missing            skillRequirementsView `json:"missing,omitempty"`
	Install            []skillInstallView    `json:"install,omitempty"`
}

type skillRequirementsView struct {
	OS      []string `json:"os,omitempty"`
	Bins    []string `json:"bins,omitempty"`
	AnyBins []string `json:"any_bins,omitempty"`
	Env     []string `json:"env,omitempty"`
	Config  []string `json:"config,omitempty"`
}

type skillInstallView struct {
	ID    string   `json:"id,omitempty"`
	Kind  string   `json:"kind,omitempty"`
	Label string   `json:"label,omitempty"`
	Bins  []string `json:"bins,omitempty"`
}

type pageData struct {
	Snapshot          snapshot
	Config            RuntimeConfigStatus
	ConfigPending     int
	ConfigCanRestart  bool
	RuntimeControl    RuntimeLifecyclePageStatus
	Prompts           PromptsStatus
	Identity          IdentityStatus
	Personas          PersonasStatus
	Chats             ChatsStatus
	ChatHistoryPath   string
	SelectedChat      *ChatView
	SelectedChatError string
	Notice            string
	Error             string
	PageRefresh       pageRefreshData
	View              adminView
	PageTitle         string
	PageSummary       string
	NavSections       []adminNavSection
}

type pageRefreshData struct {
	CurrentPath     string
	StatePath       string
	Token           string
	UpdatedAt       time.Time
	IntervalSeconds int
	Watch           bool
}

type pageStateStatus struct {
	Token     string    `json:"token,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type pageRefreshInput struct {
	Snapshot          snapshot
	Config            RuntimeConfigStatus
	RuntimeControl    RuntimeLifecyclePageStatus
	Prompts           PromptsStatus
	Identity          IdentityStatus
	Personas          PersonasStatus
	Chats             ChatsStatus
	SelectedChat      *ChatView
	SelectedChatError string
}

type adminNavSection struct {
	Label string
	Items []adminNavItem
}

type adminNavItem struct {
	Label  string
	Path   string
	Active bool
}

func (s *Service) Snapshot() snapshot { _ = "STUB: not implemented"; return *new(snapshot) }

func (s *Service) baseSnapshot() snapshot { _ = "STUB: not implemented"; return *new(snapshot) }

func (s *Service) snapshotForView(view adminView) snapshot {
	_ = "STUB: not implemented"
	return *new(snapshot)
}

func (s *Service) cronStatus() cronStatus { _ = "STUB: not implemented"; return *new(cronStatus) }

func (s *Service) skillsStatus() skillsStatus { _ = "STUB: not implemented"; return *new(skillsStatus) }

func skillViewFromStatus(entry ocskills.StatusEntry) skillView {
	_ = "STUB: not implemented"
	return *new(skillView)
}

func skillRequirementsViewFromStatus(
	req ocskills.StatusRequirements,
) skillRequirementsView {
	_ = "STUB: not implemented"
	return *new(skillRequirementsView)
}

func skillInstallViewsFromStatus(
	options []ocskills.StatusInstallOption,
) []skillInstallView {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) browserStatus() browserStatus {
	_ = "STUB: not implemented"
	return *new(browserStatus)
}

func (s *Service) probeBrowserEndpoint(
	rawURL string,
	cache map[string]browserEndpointView,
) browserEndpointView {
	_ = "STUB: not implemented"
	return *new(browserEndpointView)
}

func (s *Service) handleOverview(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSkillsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handlePromptsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleIdentityPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handlePersonasPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleChatsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleMemoryPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleAutomationPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSessionsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleDebugPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleBrowserPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) renderPage(
	w http.ResponseWriter,
	r *http.Request,
	view adminView,
) {
	_ = "STUB: not implemented"
	return
}

func pendingRestartFields(status RuntimeConfigStatus) int { _ = "STUB: not implemented"; return 0 }

func resolveSelectedChat(
	status ChatsStatus,
	provider ChatsProvider,
	selectedID string,
) (*ChatView, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func validateChatDetail(base ChatView, detail ChatView) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeChatView(
	base ChatView,
	detail ChatView,
) ChatView {
	_ = "STUB: not implemented"
	return *new(ChatView)
}

func adminNavSections(
	active adminView,
	showConfig bool,
	showRuntimeControl bool,
) []adminNavSection {
	_ = "STUB: not implemented"
	return nil
}

func runtimeDiagnosticsNavItems(
	showRuntimeControl bool,
) []adminNavItem {
	_ = "STUB: not implemented"
	return nil
}

func pageTitle(view adminView) string { _ = "STUB: not implemented"; return "" }

func pageSummary(view adminView) string { _ = "STUB: not implemented"; return "" }

func pageRefreshWatch(view adminView) bool { _ = "STUB: not implemented"; return false }

func buildPageRefreshData(
	r *http.Request,
	view adminView,
	input pageRefreshInput,
) pageRefreshData {
	_ = "STUB: not implemented"
	return *new(pageRefreshData)
}

func pageRefreshCurrentPath(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func pageRefreshStatePath(
	r *http.Request,
	view adminView,
) string {
	_ = "STUB: not implemented"
	return ""
}

func pageRefreshToken(
	view adminView,
	input pageRefreshInput,
) string {
	_ = "STUB: not implemented"
	return ""
}

func compactChatView(chat *ChatView) *ChatView { _ = "STUB: not implemented"; return nil }

func refreshTokenForValue(value any) string { _ = "STUB: not implemented"; return "" }

func (s *Service) pageRefreshInput(
	r *http.Request,
	view adminView,
) pageRefreshInput {
	_ = "STUB: not implemented"
	return *new(pageRefreshInput)
}

func (s *Service) handleStatusJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handlePageStateJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSkillsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleMemoryFilesJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleMemoryFileAPI(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleMemoryFileJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSaveMemoryFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleMemoryFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleRefreshSkills(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleToggleSkill(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) redirectWithMessageAt(
	w http.ResponseWriter,
	r *http.Request,
	key string,
	message string,
	fragment string,
) {
	_ = "STUB: not implemented"
	return
}

func redirectPathFromRequest(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func navPath(raw string) string { _ = "STUB: not implemented"; return "" }

func navViewForPath(path string) adminView { _ = "STUB: not implemented"; return *new(adminView) }

func (s *Service) handleJobsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleExecSessionsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleUploadsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleUploadSessionsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleUploadFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func uploadFiltersFromRequest(r *http.Request) uploadFilters {
	_ = "STUB: not implemented"
	return *new(uploadFilters)
}

func (s *Service) handleDebugSessionsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleDebugTracesJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleDebugFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func resolveUploadFile(root string, rel string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Service) handleRunJob(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleRemoveJob(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleClearJobs(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) requireJobAction(
	w http.ResponseWriter,
	r *http.Request,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *Service) requireCronPOST(
	w http.ResponseWriter,
	r *http.Request,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Service) requireSkillTogglePOST(
	w http.ResponseWriter,
	r *http.Request,
) (string, bool, string, string, bool) {
	_ = "STUB: not implemented"
	return "", false, "", "", false
}

func (s *Service) redirectWithMessage(
	w http.ResponseWriter,
	r *http.Request,
	key string,
	message string,
) {
	_ = "STUB: not implemented"
	return
}

func writeJSON(w http.ResponseWriter, code int, value any) { _ = "STUB: not implemented"; return }

func jobViewFromJob(job *cron.Job) jobView { _ = "STUB: not implemented"; return *new(jobView) }

func fallbackJobName(job *cron.Job) string { _ = "STUB: not implemented"; return "" }

const maxRunesPreview = 96

func summarizeText(text string, maxRunes int) string { _ = "STUB: not implemented"; return "" }

func cloneTime(src *time.Time) *time.Time { _ = "STUB: not implemented"; return nil }

func intFromMap(raw any) int { _ = "STUB: not implemented"; return 0 }

func stringSliceFromMap(raw any) []string { _ = "STUB: not implemented"; return nil }

func formatTime(raw any) string { _ = "STUB: not implemented"; return "" }

func formatUptime(startedAt time.Time, now time.Time) string { _ = "STUB: not implemented"; return "" }

func compactPortLabel(addr string) string { _ = "STUB: not implemented"; return "" }

func browserEndpointSummary(view browserEndpointView) string { _ = "STUB: not implemented"; return "" }

func displayAdminAppName(name string) string { _ = "STUB: not implemented"; return "" }

func (s *Service) resolveDebugFile(
	tracePath string,
	name string,
	relPath string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Service) resolveDebugTraceDir(tracePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveDebugRootFile(root string, relPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isAllowedDebugFile(name string) bool { _ = "STUB: not implemented"; return false }

var adminPage = template.Must(
	template.New("admin").Funcs(template.FuncMap{
		"formatTime":                 formatTime,
		"browserEndpointSummary":     browserEndpointSummary,
		"displayAdminAppName":        displayAdminAppName,
		"promptSections":             promptSections,
		"promptBlockCount":           promptBlockCount,
		"hasPromptValue":             hasPromptValue,
		"promptValuesDiffer":         promptValuesDiffer,
		"promptCollapsedSummary":     promptCollapsedSummary,
		"promptInlineEditorTitle":    promptInlineEditorTitle,
		"promptInlineEditorSummary":  promptInlineEditorSummary,
		"promptRuntimeEditorSummary": promptRuntimeEditorSummary,
		"personaStoreTitle":          personaStoreTitle,
		"personaStoreUsageLabels":    personaStoreUsageLabels,
		"personaCustomPersonas":      personaCustomPersonas,
		"personaBuiltInPersonas":     personaBuiltInPersonas,
		"personaStoreBuiltInCount":   personaStoreBuiltInCount,
		"personaStoreCustomCount":    personaStoreCustomCount,
		"personaDisplayName":         personaDisplayName,
		"personaKindLabel":           personaKindLabel,
		"personaSummaryText":         personaSummaryText,
		"chatDisplayLabel":           chatDisplayLabel,
		"chatHistoryAPIPath":         chatHistoryAPIPath,
		"chatHiddenHistory":          chatHiddenHistory,
		"chatHiddenTranscript":       chatHiddenTranscript,
		"chatHiddenTurns":            chatHiddenTurns,
		"chatKnownUsers":             chatKnownUsers,
		"chatHasTranscript":          chatHasTranscript,
		"chatHistorySummary":         chatHistorySummary,
		"chatNameSourceLabel":        chatNameSourceLabel,
		"chatTranscriptLabel":        chatTranscriptLabel,
		"chatTranscriptSummary":      chatTranscriptSummary,
		"chatTurnSpeaker":            chatTurnSpeaker,
		"chatOverrideSample":         chatOverrideSample,
		"chatVisibleHistory":         chatVisibleHistory,
		"chatVisibleTranscript":      chatVisibleTranscript,
		"chatVisibleTurns":           chatVisibleTurns,
		"hasTime":                    hasTime,
	}).Parse(
		adminPageHTML +
			promptsPageTemplateHTML +
			chatsPageTemplateHTML +
			identityPageTemplateHTML +
			personasPageTemplateHTML,
	),
)

const adminPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>TRPC-CLAW admin</title>
  <style>
    :root {
      color-scheme: light;
      --bg: #f3eee7;
      --panel: rgba(255, 252, 247, 0.92);
      --panel-strong: #fffdf8;
      --line: #d7cfc2;
      --ink: #1d1a16;
      --muted: #5f574d;
      --accent: #0f6f61;
      --warn: #9a2f2f;
      --ok: #2d6d3f;
      --shadow: 0 18px 40px rgba(35, 29, 22, 0.08);
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      min-height: 100vh;
      font-family: "Iowan Old Style", "Palatino Linotype", serif;
      color: var(--ink);
      background:
        radial-gradient(circle at top left, #fff8ef, transparent 38%),
        linear-gradient(180deg, #efe7dc 0%, var(--bg) 100%);
    }
    .app-shell {
      display: grid;
      grid-template-columns: 272px minmax(0, 1fr);
      min-height: 100vh;
    }
    .sidebar {
      position: sticky;
      top: 0;
      align-self: start;
      height: 100vh;
      overflow-y: auto;
      overscroll-behavior: contain;
      scrollbar-gutter: stable;
      padding: 24px 18px 22px;
      border-right: 1px solid rgba(215, 207, 194, 0.92);
      background: rgba(255, 250, 244, 0.78);
      backdrop-filter: blur(16px);
    }
    .sidebar-brand {
      display: flex;
      align-items: center;
      gap: 12px;
      margin-bottom: 28px;
    }
    .sidebar-mark {
      width: 42px;
      height: 42px;
      border-radius: 14px;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      background: var(--accent);
      color: white;
      font-weight: 700;
      letter-spacing: 0.04em;
      box-shadow: var(--shadow);
    }
    .sidebar-eyebrow {
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.12em;
    }
    .sidebar-title {
      margin-top: 2px;
      font-size: 26px;
      font-weight: 700;
      line-height: 1.1;
    }
    .sidebar-subtle {
      margin-top: 4px;
      color: var(--muted);
      font-size: 14px;
    }
    main {
      margin: 0;
      width: 100%;
      padding: 32px 28px 40px;
    }
    .page-wrap {
      max-width: 1440px;
    }
    .sidebar-nav {
      display: grid;
      gap: 22px;
    }
    .sidebar-section-title {
      margin: 0 0 10px;
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.12em;
    }
    .sidebar-links {
      display: grid;
      gap: 8px;
    }
    .sidebar-link {
      display: flex;
      align-items: center;
      min-height: 42px;
      padding: 10px 14px;
      border-radius: 14px;
      border: 1px solid transparent;
      color: var(--ink);
      text-decoration: none;
      font-weight: 700;
      transition: background 120ms ease, border-color 120ms ease, color 120ms ease;
    }
    .sidebar-link:hover {
      background: rgba(255, 253, 248, 0.88);
      border-color: rgba(215, 207, 194, 0.88);
    }
    .sidebar-link.active {
      background: rgba(15, 111, 97, 0.1);
      border-color: rgba(15, 111, 97, 0.24);
      color: var(--accent);
      box-shadow: var(--shadow);
    }
    .page-header {
      margin-bottom: 18px;
    }
    .page-toolbar {
      margin-top: 16px;
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      justify-content: space-between;
      gap: 12px 16px;
    }
    .page-toolbar-copy {
      display: grid;
      gap: 4px;
      min-width: 0;
    }
    .page-toolbar-updated {
      color: var(--muted);
      font-size: 13px;
      font-weight: 700;
      letter-spacing: 0.02em;
    }
    .page-toolbar-note {
      color: var(--muted);
      font-size: 14px;
      max-width: 720px;
    }
    .page-refresh-link {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      min-height: 40px;
      padding: 8px 14px;
      border-radius: 999px;
      border: 1px solid var(--line);
      background: rgba(255, 253, 248, 0.92);
      color: var(--ink);
      text-decoration: none;
      font-weight: 700;
      box-shadow: var(--shadow);
    }
    .page-refresh-link:hover {
      border-color: rgba(15, 111, 97, 0.28);
      color: var(--accent);
    }
    .page-refresh-alert {
      margin-top: 16px;
    }
    .page-kicker {
      margin: 0 0 10px;
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.12em;
    }
    h1, h2 { margin: 0 0 14px; }
    h1 { font-size: 36px; }
    h2 { font-size: 22px; }
    p, li, td, th, button, code {
      font-size: 15px;
      line-height: 1.5;
    }
    .subtle {
      color: var(--muted);
      max-width: 860px;
    }
    .notice {
      margin: 18px 0 0;
      padding: 12px 14px;
      border-radius: 14px;
      border: 1px solid var(--line);
      background: var(--panel-strong);
      box-shadow: var(--shadow);
    }
    .notice.ok { border-color: rgba(45, 109, 63, 0.3); }
    .notice.err { border-color: rgba(154, 47, 47, 0.3); }
    .stats,
    .panels {
      display: grid;
      gap: 16px;
      margin-top: 24px;
    }
    .stats { grid-template-columns: repeat(auto-fit, minmax(180px, 1fr)); }
    .panels { grid-template-columns: repeat(auto-fit, minmax(320px, 1fr)); }
    .card {
      border: 1px solid var(--line);
      border-radius: 20px;
      padding: 20px;
      background: var(--panel);
      box-shadow: var(--shadow);
      backdrop-filter: blur(8px);
    }
    .stat-label {
      color: var(--muted);
      text-transform: uppercase;
      letter-spacing: 0.08em;
      font-size: 12px;
    }
    .stat-value {
      display: block;
      margin-top: 8px;
      font-size: 28px;
      font-weight: 700;
    }
    .meta {
      margin: 0;
      display: grid;
      grid-template-columns: minmax(110px, 160px) 1fr;
      gap: 8px 12px;
    }
    .meta dt {
      color: var(--muted);
      font-weight: 700;
      min-width: 0;
    }
    .meta dd {
      margin: 0;
      min-width: 0;
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    a { color: var(--accent); }
    code {
      background: rgba(15, 111, 97, 0.08);
      padding: 2px 6px;
      border-radius: 8px;
      word-break: break-all;
    }
    input[type="text"],
    textarea {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 16px;
      padding: 12px 14px;
      font: inherit;
      background: var(--panel-strong);
      color: var(--ink);
    }
    textarea {
      min-height: 160px;
      resize: vertical;
      white-space: pre-wrap;
      font-family: "SFMono-Regular", "SFMono-Regular", monospace;
      line-height: 1.45;
    }
    table {
      width: 100%;
      border-collapse: collapse;
      margin-top: 12px;
    }
    th, td {
      text-align: left;
      vertical-align: top;
      min-width: 0;
      padding: 12px 10px;
      border-top: 1px solid var(--line);
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    th {
      color: var(--muted);
      font-size: 13px;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
    .actions {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }
    form { margin: 0; }
    button {
      border: 0;
      border-radius: 999px;
      padding: 8px 14px;
      background: var(--accent);
      color: white;
      cursor: pointer;
    }
    button.secondary {
      background: #c9bca9;
      color: var(--ink);
    }
    button.warn { background: var(--warn); }
    .empty {
      margin-top: 14px;
      color: var(--muted);
    }
    .preview-box {
      max-width: 220px;
    }
    .preview-box img,
    .preview-box video {
      display: block;
      max-width: 220px;
      max-height: 140px;
      border-radius: 12px;
      border: 1px solid var(--line);
      background: white;
    }
    .preview-box audio {
      width: 220px;
      max-width: 100%;
    }
    .filter-tabs {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      margin-top: 14px;
    }
    .filter-tab {
      background: #e6ddcf;
      color: var(--ink);
    }
    .filter-tab.active {
      background: var(--accent);
      color: white;
    }
    .skills-controls {
      display: grid;
      grid-template-columns: minmax(0, 1fr) auto;
      gap: 12px 16px;
      align-items: center;
      margin-top: 18px;
    }
    .skills-search-wrap {
      min-width: 0;
    }
    .skills-controls input {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 999px;
      padding: 12px 16px;
      font: inherit;
      background: var(--panel-strong);
      color: var(--ink);
    }
    .skills-toolbar-side {
      display: flex;
      align-items: center;
      justify-content: flex-end;
      gap: 12px;
    }
    .skills-shown {
      color: var(--muted);
      font-weight: 700;
      white-space: nowrap;
    }
    .skills-header {
      display: flex;
      flex-wrap: wrap;
      gap: 10px 18px;
      align-items: flex-end;
      justify-content: space-between;
    }
    .skills-header-copy {
      max-width: 820px;
    }
    .skills-lead {
      margin: 8px 0 0;
      color: var(--muted);
      max-width: 700px;
    }
    .skills-ops-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
      gap: 12px;
      margin-top: 16px;
    }
    .skills-op-card {
      border: 1px solid var(--line);
      border-radius: 16px;
      padding: 14px 16px;
      background: rgba(255, 253, 248, 0.72);
    }
    .skills-op-label {
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.1em;
    }
    .skills-op-value {
      margin-top: 8px;
      font-weight: 700;
      line-height: 1.45;
    }
    .skills-op-value code {
      background: rgba(15, 111, 97, 0.08);
    }
    .skills-op-note {
      margin-top: 8px;
      color: var(--muted);
      font-size: 14px;
    }
    .skills-op-actions {
      margin-top: 12px;
      display: flex;
      align-items: center;
      gap: 10px;
      flex-wrap: wrap;
    }
    .skills-op-actions form {
      margin: 0;
    }
    .skills-op-actions button {
      padding: 7px 12px;
      font-size: 14px;
    }
    .config-sections {
      display: grid;
      gap: 18px;
      margin-top: 16px;
    }
    .config-section-card {
      border: 1px solid var(--line);
      border-radius: 18px;
      padding: 18px;
      background: rgba(255, 253, 248, 0.8);
      box-shadow: 0 10px 24px rgba(35, 29, 22, 0.04);
    }
    .config-field-list {
      display: grid;
      gap: 14px;
      margin-top: 14px;
    }
    .config-field {
      border: 1px solid rgba(215, 207, 194, 0.9);
      border-radius: 16px;
      background: rgba(255, 252, 247, 0.94);
      overflow: hidden;
      transition: box-shadow 140ms ease, border-color 140ms ease;
    }
    .config-field[open] {
      border-color: rgba(15, 111, 97, 0.28);
      box-shadow: 0 14px 28px rgba(35, 29, 22, 0.08);
    }
    .config-field summary {
      list-style: none;
      cursor: pointer;
      padding: 14px 16px;
    }
    .config-field summary::-webkit-details-marker {
      display: none;
    }
    .config-field-detail {
      padding: 0 16px 16px;
      border-top: 1px solid rgba(215, 207, 194, 0.65);
    }
    .config-field-top {
      display: flex;
      align-items: flex-start;
      justify-content: space-between;
      gap: 12px;
    }
    .config-field-title {
      margin: 0;
      font-size: 18px;
    }
    .config-badges {
      display: flex;
      align-items: center;
      gap: 8px;
      flex-wrap: wrap;
    }
    .config-badge {
      border-radius: 999px;
      border: 1px solid rgba(15, 111, 97, 0.18);
      background: rgba(15, 111, 97, 0.08);
      color: var(--accent);
      padding: 4px 10px;
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
    .config-badge.warn {
      border-color: rgba(154, 47, 47, 0.18);
      background: rgba(154, 47, 47, 0.08);
      color: var(--warn);
    }
    .config-meta {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
      gap: 10px 16px;
      margin-top: 12px;
    }
    .config-meta-block {
      border-radius: 14px;
      background: rgba(243, 238, 231, 0.62);
      padding: 10px 12px;
    }
    .config-meta-label {
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
    .config-meta-value {
      margin-top: 6px;
      line-height: 1.5;
      word-break: break-word;
    }
    .config-form {
      display: grid;
      gap: 10px;
      margin-top: 14px;
    }
    .config-form-row {
      display: flex;
      align-items: center;
      gap: 10px;
      flex-wrap: wrap;
    }
    .config-form input,
    .config-form select {
      min-width: min(100%, 320px);
      padding: 10px 12px;
      border-radius: 12px;
      border: 1px solid var(--line);
      background: rgba(255, 255, 255, 0.96);
      color: var(--ink);
      font: inherit;
    }
    .config-form .subtle {
      margin: 0;
    }
    .runtime-meta-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
      gap: 10px 16px;
      margin-top: 12px;
    }
    .runtime-meta-card {
      border-radius: 14px;
      background: rgba(243, 238, 231, 0.62);
      padding: 10px 12px;
    }
    .runtime-meta-label {
      color: var(--muted);
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
    .runtime-meta-value {
      margin-top: 6px;
      line-height: 1.5;
      word-break: break-word;
    }
    .runtime-pending {
      margin-top: 16px;
      padding-top: 16px;
      border-top: 1px solid rgba(215, 207, 194, 0.65);
    }
    .runtime-pending h3,
    .runtime-version-card h3,
    .runtime-changelog-card h3 {
      margin: 0 0 10px;
      font-size: 16px;
    }
    .runtime-summary-list {
      margin: 12px 0 0;
      padding-left: 20px;
    }
    .runtime-summary-list li {
      color: var(--muted);
    }
    .runtime-version-form,
    .runtime-version-view-form {
      display: grid;
      gap: 10px;
      margin-top: 14px;
    }
    .runtime-version-form select,
    .runtime-version-view-form select {
      min-width: min(100%, 320px);
      padding: 10px 12px;
      border-radius: 12px;
      border: 1px solid var(--line);
      background: rgba(255, 255, 255, 0.96);
      color: var(--ink);
      font: inherit;
    }
    .runtime-version-actions {
      display: flex;
      align-items: center;
      gap: 10px;
      flex-wrap: wrap;
    }
    .runtime-changelog {
      margin-top: 12px;
      padding: 14px;
      border-radius: 16px;
      border: 1px solid rgba(215, 207, 194, 0.85);
      background: rgba(255, 253, 248, 0.9);
      white-space: pre-wrap;
      font-family: "SFMono-Regular", "SFMono-Regular", monospace;
      line-height: 1.45;
      max-height: 420px;
      overflow: auto;
    }
    .skills-group {
      margin-top: 18px;
    }
    .skills-group h3 {
      margin: 0 0 10px;
      font-size: 16px;
      color: var(--muted);
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }
    .skill-card {
      border: 1px solid var(--line);
      border-radius: 18px;
      background: var(--panel-strong);
      margin-top: 12px;
      overflow: hidden;
      transition: box-shadow 140ms ease, border-color 140ms ease;
    }
    .skill-card[open] {
      border-color: rgba(15, 111, 97, 0.28);
      box-shadow: 0 14px 28px rgba(35, 29, 22, 0.08);
    }
    .skill-card summary {
      list-style: none;
      cursor: pointer;
      padding: 16px 18px;
    }
    .skill-card summary::-webkit-details-marker {
      display: none;
    }
    .skill-main {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      gap: 18px;
    }
    .skill-copy {
      min-width: 0;
      flex: 1 1 auto;
    }
    .skill-headline {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      gap: 10px;
    }
    .skill-name {
      display: flex;
      align-items: center;
      gap: 8px;
      font-weight: 700;
      font-size: 17px;
    }
    .skill-dot {
      width: 10px;
      height: 10px;
      border-radius: 999px;
      background: var(--line);
      flex: 0 0 auto;
    }
    .skill-dot.ready { background: var(--ok); }
    .skill-dot.needs-setup { background: #c27a20; }
    .skill-dot.disabled { background: var(--muted); }
    .skill-badges {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }
    .skill-badges.inline {
      margin-top: 10px;
    }
    .skill-summary-side {
      display: flex;
      flex-direction: column;
      align-items: flex-end;
      justify-content: flex-start;
      gap: 8px;
      flex: 0 0 auto;
      min-width: 72px;
    }
    .skill-badge {
      border-radius: 999px;
      padding: 4px 9px;
      font-size: 12px;
      border: 1px solid var(--line);
      background: rgba(15, 111, 97, 0.08);
      color: var(--ink);
    }
    .skill-badge.status {
      padding: 6px 12px;
      font-size: 13px;
      font-weight: 700;
      letter-spacing: 0.01em;
    }
    .skill-badge.ready {
      color: var(--ok);
      border-color: rgba(45, 109, 63, 0.25);
      background: rgba(45, 109, 63, 0.08);
    }
    .skill-badge.needs-setup {
      color: #9b5f12;
      border-color: rgba(194, 122, 32, 0.25);
      background: rgba(194, 122, 32, 0.08);
    }
    .skill-badge.disabled {
      color: var(--muted);
      border-color: rgba(95, 87, 77, 0.18);
      background: rgba(95, 87, 77, 0.08);
    }
    .skill-description {
      margin-top: 10px;
      color: #3f3932;
      max-width: 820px;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      overflow: hidden;
    }
    .skill-card[open] .skill-description {
      display: block;
      overflow: visible;
    }
    .skill-reason {
      margin-top: 8px;
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      gap: 8px;
      color: var(--muted);
      font-size: 14px;
      min-width: 0;
    }
    .skill-reason-label {
      display: inline-flex;
      align-items: center;
      min-height: 22px;
      padding: 2px 8px;
      border-radius: 999px;
      border: 1px solid rgba(95, 87, 77, 0.18);
      background: rgba(95, 87, 77, 0.06);
      color: var(--muted);
      font-size: 12px;
      font-weight: 700;
      letter-spacing: 0.01em;
    }
    .skill-reason-label.needs-setup {
      color: #9b5f12;
      border-color: rgba(194, 122, 32, 0.25);
      background: rgba(194, 122, 32, 0.08);
    }
    .skill-reason-label.disabled {
      color: #655b50;
      border-color: rgba(95, 87, 77, 0.22);
      background: rgba(95, 87, 77, 0.1);
    }
    .skill-reason-text {
      min-width: 0;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
    .skill-card[open] .skill-reason-text {
      white-space: normal;
    }
    .skill-toggle-group {
      display: inline-flex;
      align-items: center;
      flex: 0 0 auto;
    }
    .skill-inline-toggle-form {
      margin: 0;
      flex: 0 0 auto;
    }
    .skill-inline-toggle {
      display: inline-flex;
      align-items: center;
      border: 0;
      border-radius: 999px;
      padding: 0;
      background: transparent;
      color: var(--ink);
      cursor: pointer;
      font: inherit;
    }
    .skill-inline-toggle:focus-visible {
      outline: 2px solid rgba(15, 111, 97, 0.42);
      outline-offset: 3px;
    }
    .skill-inline-toggle-track {
      position: relative;
      width: 54px;
      height: 32px;
      border-radius: 999px;
      background: rgba(95, 87, 77, 0.28);
      border: 1px solid rgba(95, 87, 77, 0.18);
      transition: background 120ms ease, border-color 120ms ease;
    }
    .skill-inline-toggle-track::after {
      content: "";
      position: absolute;
      top: 3px;
      left: 3px;
      width: 24px;
      height: 24px;
      border-radius: 999px;
      background: white;
      box-shadow: 0 4px 10px rgba(35, 29, 22, 0.18);
      transition: transform 120ms ease;
    }
    .skill-inline-toggle.enabled .skill-inline-toggle-track {
      background: rgba(45, 109, 63, 0.9);
      border-color: rgba(45, 109, 63, 0.35);
    }
    .skill-inline-toggle.enabled .skill-inline-toggle-track::after {
      transform: translateX(22px);
    }
    .skill-details {
      margin-top: 14px;
      padding: 14px 18px 18px;
      border-top: 1px solid var(--line);
      background: rgba(15, 111, 97, 0.02);
    }
    .skill-details-head {
      display: flex;
      flex-wrap: wrap;
      align-items: flex-start;
      gap: 12px;
      margin-bottom: 14px;
    }
    .skill-details-grid {
      display: grid;
      gap: 12px;
      grid-template-columns: repeat(auto-fit, minmax(230px, 1fr));
    }
    .skill-list {
      margin: 8px 0 0;
      padding-left: 18px;
    }
    .memory-preview {
      max-width: 540px;
      color: #3f3932;
      overflow-wrap: anywhere;
      white-space: pre-wrap;
    }
    .prompt-detail {
      overflow: hidden;
    }
    .prompt-detail[open] {
      border-color: rgba(15, 111, 97, 0.28);
      box-shadow: 0 14px 28px rgba(35, 29, 22, 0.08);
    }
    .prompt-detail summary {
      list-style: none;
      cursor: pointer;
    }
    .prompt-detail summary::-webkit-details-marker {
      display: none;
    }
    .prompt-detail-copy,
    .prompt-detail-hint {
      margin: 8px 0 0;
    }
    .prompt-detail-body {
      margin-top: 14px;
      padding-top: 14px;
      border-top: 1px solid var(--line);
    }
    .memory-scope {
      display: grid;
      gap: 4px;
    }
    .memory-user-label {
      color: var(--muted);
      font-size: 0.92rem;
      font-weight: 700;
    }
    .memory-controls {
      margin: 18px 0 12px;
    }
    .memory-filter-head {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 12px;
      margin-bottom: 10px;
    }
    .memory-filter-head label {
      color: var(--muted);
      font-size: 0.92rem;
      font-weight: 700;
    }
    .memory-filter-grid {
      display: grid;
      grid-template-columns: minmax(0, 1fr);
      gap: 12px;
      min-width: 0;
    }
    .memory-filter-field {
      min-width: 0;
    }
    .memory-filter-field label {
      display: block;
      margin-bottom: 6px;
      color: var(--muted);
      font-size: 0.92rem;
      font-weight: 700;
    }
    .memory-filter-field input {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 999px;
      padding: 12px 16px;
      font: inherit;
      background: var(--panel-strong);
      color: var(--ink);
    }
    .memory-shown {
      color: var(--muted);
      font-weight: 700;
      white-space: nowrap;
    }
    .memory-list {
      display: grid;
      gap: 14px;
      margin-top: 16px;
    }
    .memory-card-head {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      gap: 14px;
      align-items: flex-start;
    }
    .memory-card-meta {
      display: grid;
      gap: 6px;
      color: var(--muted);
      font-size: 0.9rem;
      text-align: right;
      white-space: nowrap;
    }
    .memory-path {
      margin-top: 8px;
      color: var(--muted);
      overflow-wrap: anywhere;
    }
    .memory-card-toolbar {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      gap: 10px 14px;
      align-items: center;
      margin-bottom: 12px;
    }
    .memory-editor-form {
      margin-top: 12px;
    }
    .memory-editor-form textarea {
      width: 100%;
      min-height: 240px;
      margin-top: 8px;
      font: inherit;
      resize: vertical;
    }
    .memory-editor-status {
      margin-top: 8px;
      color: var(--muted);
      font-size: 0.92rem;
    }
    .memory-editor-status.error {
      color: #b6544d;
    }
    .chat-list {
      display: grid;
      gap: 14px;
      margin-top: 16px;
    }
    .chat-card {
      border: 1px solid var(--line);
      border-radius: 18px;
      padding: 16px 18px;
      background: rgba(255, 253, 248, 0.72);
    }
    .chat-card-head {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      gap: 12px;
      flex-wrap: wrap;
    }
    .chat-card-copy {
      min-width: 0;
      flex: 1 1 280px;
    }
    .chat-card-title {
      font-size: 18px;
      font-weight: 700;
      line-height: 1.35;
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    .chat-card-kind {
      margin-top: 6px;
      color: var(--muted);
    }
    .chat-card-link {
      flex: 0 0 auto;
      white-space: nowrap;
    }
    .chat-card-grid {
      display: grid;
      gap: 10px 18px;
      grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
      margin-top: 14px;
    }
    .chat-card-meta {
      min-width: 0;
    }
    .chat-card-label {
      color: var(--muted);
      font-size: 12px;
      font-weight: 700;
      letter-spacing: 0.08em;
      text-transform: uppercase;
    }
    .chat-card-value {
      margin-top: 6px;
      line-height: 1.45;
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    .chat-detail-section + .chat-detail-section {
      margin-top: 24px;
      padding-top: 22px;
      border-top: 1px solid var(--line);
    }
    .chat-detail-head {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      gap: 12px;
      flex-wrap: wrap;
      margin-bottom: 12px;
    }
    .chat-detail-head h3,
    .chat-action-card h4,
    .chat-transcript-title {
      margin: 0;
    }
    .chat-disclosure,
    .chat-disclosure-more {
      border: 1px solid var(--line);
      border-radius: 16px;
      background: rgba(255, 253, 248, 0.72);
      overflow: hidden;
    }
    .chat-disclosure[open],
    .chat-disclosure-more[open] {
      border-color: rgba(15, 111, 97, 0.28);
      box-shadow: 0 14px 28px rgba(35, 29, 22, 0.08);
    }
    .chat-disclosure summary,
    .chat-disclosure-more summary {
      list-style: none;
      cursor: pointer;
      padding: 14px 16px;
    }
    .chat-disclosure summary::-webkit-details-marker,
    .chat-disclosure-more summary::-webkit-details-marker {
      display: none;
    }
    .chat-disclosure-meta {
      margin: 8px 0 0;
    }
    .chat-disclosure-body {
      margin-top: 0;
      padding: 0 16px 16px;
      border-top: 1px solid var(--line);
    }
    .chat-disclosure-body > :first-child {
      margin-top: 14px;
    }
    .chat-disclosure-body > :last-child {
      margin-bottom: 0;
    }
    .chat-disclosure-more {
      margin-top: 12px;
    }
    .chat-history-shell {
      margin-top: 14px;
    }
    .chat-history-status {
      margin: 0;
    }
    .chat-history-toolbar {
      margin: 14px 0 0;
    }
    .chat-history-more {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 16px;
      padding: 12px 14px;
      background: rgba(255, 253, 248, 0.82);
      color: var(--ink);
      font: inherit;
      text-align: left;
      cursor: pointer;
      transition: border-color 120ms ease, box-shadow 120ms ease;
    }
    .chat-history-more:hover,
    .chat-history-more:focus {
      border-color: rgba(15, 111, 97, 0.38);
      box-shadow: 0 12px 24px rgba(35, 29, 22, 0.08);
      outline: none;
    }
    .chat-history-bounded {
      margin: 12px 0 0;
    }
    .chat-timeline {
      display: grid;
      gap: 12px;
      margin-top: 14px;
    }
    .chat-timeline-session {
      padding-top: 12px;
      border-top: 1px solid var(--line);
    }
    .chat-timeline-session:first-child {
      padding-top: 0;
      border-top: 0;
    }
    .chat-timeline-session-head {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      gap: 12px;
      flex-wrap: wrap;
    }
    .chat-timeline-session-label {
      margin: 0;
      font-size: 16px;
      font-weight: 700;
    }
    .chat-timeline-session-meta {
      margin-top: 6px;
    }
    .chat-transcript-list {
      display: grid;
      gap: 14px;
      margin-top: 12px;
    }
    .chat-transcript-card,
    .chat-action-card {
      border: 1px solid var(--line);
      border-radius: 16px;
      padding: 14px 16px;
      background: rgba(255, 253, 248, 0.72);
      min-width: 0;
    }
    .chat-transcript-head {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      gap: 12px;
      flex-wrap: wrap;
    }
    .chat-turn-list {
      display: grid;
      gap: 10px;
      margin-top: 14px;
    }
    .chat-turn {
      border-left: 3px solid rgba(15, 111, 97, 0.2);
      padding-left: 12px;
      min-width: 0;
    }
    .chat-turn-head {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
      gap: 12px;
      flex-wrap: wrap;
    }
    .chat-turn-speaker {
      font-weight: 700;
    }
    .chat-turn-quote {
      margin: 8px 0 0;
      padding-left: 12px;
      border-left: 2px solid var(--line);
      color: var(--muted);
    }
    .chat-turn-text {
      margin-top: 8px;
      white-space: pre-wrap;
      overflow-wrap: anywhere;
      word-break: break-word;
    }
    .chat-action-grid {
      display: grid;
      gap: 12px;
      grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
      margin-top: 12px;
    }
    @media (max-width: 760px) {
      .app-shell {
        grid-template-columns: 1fr;
      }
      .sidebar {
        position: static;
        height: auto;
        overflow: visible;
        border-right: 0;
        border-bottom: 1px solid rgba(215, 207, 194, 0.92);
      }
      main {
        padding: 24px 16px 32px;
      }
      h1 { font-size: 30px; }
      .meta { grid-template-columns: 1fr; }
      .skills-controls {
        grid-template-columns: 1fr;
      }
      .memory-filter-grid {
        grid-template-columns: 1fr;
      }
      .skills-toolbar-side {
        justify-content: space-between;
      }
      .memory-filter-head {
        align-items: flex-start;
        flex-direction: column;
      }
      .skills-header {
        align-items: flex-start;
      }
      .skill-main {
        flex-direction: column;
      }
      .skill-summary-side {
        width: 100%;
        justify-content: flex-start;
        align-items: flex-start;
      }
      table, thead, tbody, th, td, tr { display: block; }
      thead { display: none; }
      td {
        padding: 10px 0;
        border-top: 0;
      }
      tr {
        padding: 14px 0;
        border-top: 1px solid var(--line);
      }
    }
  </style>
</head>
<body>
  <div class="app-shell">
    <aside class="sidebar">
      <div class="sidebar-brand">
        <div class="sidebar-mark">TC</div>
        <div>
          <div class="sidebar-eyebrow">control</div>
          <div class="sidebar-title">TRPC-CLAW</div>
          {{if .Snapshot.AppName}}
          <div class="sidebar-subtle">{{displayAdminAppName .Snapshot.AppName}}</div>
          {{end}}
        </div>
      </div>
      <nav class="sidebar-nav" aria-label="Admin sections">
        {{range .NavSections}}
        <section>
          <div class="sidebar-section-title">{{.Label}}</div>
          <div class="sidebar-links">
            {{range .Items}}
            <a class="sidebar-link{{if .Active}} active{{end}}" href="{{.Path}}">
              {{.Label}}
            </a>
            {{end}}
          </div>
        </section>
        {{end}}
      </nav>
    </aside>
    <main>
      <div class="page-wrap">
        <header class="page-header">
          <p class="page-kicker">TRPC-CLAW admin</p>
          <h1>{{.PageTitle}}</h1>
          <p class="subtle">{{.PageSummary}}</p>
          <div class="page-toolbar">
            <div class="page-toolbar-copy">
              <div class="page-toolbar-updated">
                Updated {{formatTime .PageRefresh.UpdatedAt}}
              </div>
              {{if .PageRefresh.Watch}}
              <div class="page-toolbar-note">
                This page watches for newer runtime state without
                interrupting your reading or editing.
              </div>
              {{end}}
            </div>
            <a class="page-refresh-link" href="{{.PageRefresh.CurrentPath}}">
              Refresh page
            </a>
          </div>
        </header>
        {{if .PageRefresh.Watch}}
        <div
          class="notice page-refresh-alert"
          hidden
          data-page-stale-root
          data-page-state-path="{{.PageRefresh.StatePath}}"
          data-page-state-token="{{.PageRefresh.Token}}"
          data-page-refresh-interval="{{.PageRefresh.IntervalSeconds}}"
        >
          New runtime state is available for this page.
          <a href="{{.PageRefresh.CurrentPath}}">Refresh page</a>
        </div>
        {{end}}
        {{if .Notice}}<div class="notice ok">{{.Notice}}</div>{{end}}
        {{if .Error}}<div class="notice err">{{.Error}}</div>{{end}}

    {{if eq .View "overview"}}
    <section class="stats">
      <article class="card">
        <span class="stat-label">Instance</span>
        <span class="stat-value">{{.Snapshot.InstanceID}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Gateway</span>
        <span class="stat-value">{{if .Snapshot.GatewayLabel}}{{.Snapshot.GatewayLabel}}{{else}}{{.Snapshot.GatewayAddr}}{{end}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Jobs</span>
        <span class="stat-value">{{.Snapshot.Cron.JobCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Skills</span>
        <span class="stat-value">{{.Snapshot.Skills.TotalCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Memory Files</span>
        <span class="stat-value">{{.Snapshot.Memory.FileCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Exec Sessions</span>
        <span class="stat-value">{{.Snapshot.Exec.SessionCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Uploads</span>
        <span class="stat-value">{{.Snapshot.Uploads.FileCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Browser Profiles</span>
        <span class="stat-value">{{.Snapshot.Browser.ProfileCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Debug Sessions</span>
        <span class="stat-value">{{.Snapshot.Debug.SessionCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Recent Traces</span>
        <span class="stat-value">{{.Snapshot.Debug.TraceCount}}</span>
      </article>
      <article class="card">
        <span class="stat-label">Langfuse</span>
        <span class="stat-value">
          {{if .Snapshot.Langfuse.Ready}}ready
          {{else if .Snapshot.Langfuse.Enabled}}error
          {{else}}off{{end}}
        </span>
      </article>
    </section>

    <section class="panels">
      <article class="card">
        <h2>Runtime</h2>
        <dl class="meta">
          <dt>App</dt>
          <dd>{{displayAdminAppName .Snapshot.AppName}}</dd>
          <dt>Agent Type</dt>
          <dd>
            {{if .Snapshot.AgentType}}
              {{.Snapshot.AgentType}}
            {{else}}
              -
            {{end}}
          </dd>
          <dt>Model</dt>
          <dd>
            {{if .Snapshot.ModelName}}
              {{.Snapshot.ModelMode}} / {{.Snapshot.ModelName}}
            {{else if .Snapshot.ModelMode}}
              {{.Snapshot.ModelMode}}
            {{else}}-{{end}}
          </dd>
          <dt>Session Backend</dt>
          <dd>
            {{if .Snapshot.SessionBackend}}
              {{.Snapshot.SessionBackend}}
            {{else}}
              -
            {{end}}
          </dd>
          <dt>Memory Backend</dt>
          <dd>
            {{if .Snapshot.MemoryBackend}}
              {{.Snapshot.MemoryBackend}}
            {{else}}
              -
            {{end}}
          </dd>
          <dt>Host</dt>
          <dd>
            {{if .Snapshot.Hostname}}
              {{.Snapshot.Hostname}}
            {{else}}
              -
            {{end}}
          </dd>
          <dt>PID</dt>
          <dd>{{if .Snapshot.PID}}{{.Snapshot.PID}}{{else}}-{{end}}</dd>
          <dt>Started</dt>
          <dd>{{formatTime .Snapshot.StartedAt}}</dd>
          <dt>Uptime</dt>
          <dd>{{.Snapshot.Uptime}}</dd>
          <dt>Gateway URL</dt>
          <dd>
            <a href="{{.Snapshot.GatewayURL}}">
              {{.Snapshot.GatewayURL}}
            </a>
          </dd>
          <dt>Admin URL</dt>
          <dd><a href="{{.Snapshot.AdminURL}}">{{.Snapshot.AdminURL}}</a></dd>
          <dt>Admin Auto Port</dt>
          <dd>{{.Snapshot.AdminAutoPort}}</dd>
          <dt>State Dir</dt>
          <dd><code>{{.Snapshot.StateDir}}</code></dd>
          <dt>Debug Dir</dt>
          <dd><code>{{.Snapshot.DebugDir}}</code></dd>
          <dt>Generated</dt>
          <dd>{{formatTime .Snapshot.GeneratedAt}}</dd>
        </dl>
      </article>

      <article class="card">
        <h2>Gateway Surface</h2>
        <dl class="meta">
          <dt>Health</dt>
          <dd><code>{{.Snapshot.Routes.HealthPath}}</code></dd>
          <dt>Messages</dt>
          <dd><code>{{.Snapshot.Routes.MessagesPath}}</code></dd>
          <dt>Status</dt>
          <dd><code>{{.Snapshot.Routes.StatusPath}}</code></dd>
          <dt>Cancel</dt>
          <dd><code>{{.Snapshot.Routes.CancelPath}}</code></dd>
          <dt>Channels</dt>
          <dd>
            {{if .Snapshot.Channels}}
              {{range $i, $ch := .Snapshot.Channels}}
                {{if $i}}, {{end}}{{$ch}}
              {{end}}
            {{else}}none{{end}}
          </dd>
          <dt>JSON</dt>
          <dd>
            <a href="/api/status">status</a> ·
            <a href="/api/skills/status">skills</a> ·
            <a href="/api/memory/files">memory</a> ·
            <a href="/api/cron/jobs">jobs</a> ·
            <a href="/api/exec/sessions">exec</a> ·
            <a href="/api/uploads">uploads</a> ·
            <a href="/api/uploads/sessions">upload sessions</a> ·
            <a href="/api/debug/sessions">debug sessions</a> ·
            <a href="/api/debug/traces">debug traces</a>
          </dd>
        </dl>
      </article>

      <article class="card">
        <h2>Automation</h2>
        {{if .Snapshot.Cron.Enabled}}
        <p class="subtle">
          Persisted jobs continue after gateway restarts. Telegram quick
          commands like <code>/jobs</code> and <code>/jobs_clear</code> remain
          useful, but the admin surface is the place for global inspection and
          one-off management.
        </p>
        <div class="actions">
          <form method="post" action="/api/cron/jobs/clear">
            <input type="hidden" name="return_path" value="/overview">
            <button
              class="warn"
              type="submit"
              onclick="return confirm('Clear all scheduled jobs?');"
            >
              Clear All Jobs
            </button>
          </form>
        </div>
        {{else}}
        <p class="empty">Scheduled jobs are not enabled for this runtime.</p>
        {{end}}
      </article>

      <article class="card">
        <h2>Skills Surface</h2>
        <p class="subtle">
          Read-only management view for bundled, local, and external skills.
          It highlights disabled skills, setup gaps, and the config/env
          requirements you still need to satisfy before a skill becomes usable.
        </p>
        <dl class="meta">
          <dt>Total</dt>
          <dd>{{.Snapshot.Skills.TotalCount}}</dd>
          <dt>Ready</dt>
          <dd>{{.Snapshot.Skills.ReadyCount}}</dd>
          <dt>Needs Setup</dt>
          <dd>{{.Snapshot.Skills.NeedsSetupCount}}</dd>
          <dt>Disabled</dt>
          <dd>{{.Snapshot.Skills.DisabledCount}}</dd>
          <dt>Bundled</dt>
          <dd>{{.Snapshot.Skills.BundledCount}}</dd>
          <dt>JSON</dt>
          <dd><a href="/api/skills/status">/api/skills/status</a></dd>
        </dl>
        <p class="subtle" style="margin-top: 12px;">
          Open <a href="/skills">Skills Inventory</a>.
        </p>
      </article>

      <article class="card">
        <h2>Sessions</h2>
        <p class="subtle">
          Exec sessions and persisted uploads live on their own page so the
          overview stays scannable.
        </p>
        <dl class="meta">
          <dt>Exec Sessions</dt>
          <dd>{{.Snapshot.Exec.SessionCount}}</dd>
          <dt>Running Exec</dt>
          <dd>{{.Snapshot.Exec.RunningCount}}</dd>
          <dt>Uploads</dt>
          <dd>{{.Snapshot.Uploads.FileCount}}</dd>
          <dt>Upload Sessions</dt>
          <dd>{{len .Snapshot.Uploads.Sessions}}</dd>
          <dt>Open</dt>
          <dd><a href="/sessions">Sessions</a></dd>
        </dl>
      </article>

      <article class="card">
        <h2>Memory Surface</h2>
        <p class="subtle">
          Durable memory can use structured backends or file-backed
          <code>MEMORY.md</code> scopes. The Memory page inventories the
          file-backed scopes when this runtime uses the file backend.
        </p>
        <dl class="meta">
          <dt>Backend</dt>
          <dd>{{if .Snapshot.Memory.Backend}}{{.Snapshot.Memory.Backend}}{{else}}-{{end}}</dd>
          <dt>File Memory</dt>
          <dd>{{.Snapshot.Memory.FileEnabled}}</dd>
          <dt>Files</dt>
          <dd>{{.Snapshot.Memory.FileCount}}</dd>
          <dt>Total Bytes</dt>
          <dd>{{.Snapshot.Memory.TotalBytes}}</dd>
          <dt>JSON</dt>
          <dd><a href="/api/memory/files">/api/memory/files</a></dd>
          <dt>Open</dt>
          <dd><a href="/memory">Memory</a></dd>
        </dl>
      </article>

      <article class="card">
        <h2>Debug</h2>
        <p class="subtle">
          Trace browsing and Langfuse drill-down live on a separate page.
        </p>
        <dl class="meta">
          <dt>Debug Sessions</dt>
          <dd>{{.Snapshot.Debug.SessionCount}}</dd>
          <dt>Recent Traces</dt>
          <dd>{{.Snapshot.Debug.TraceCount}}</dd>
          <dt>Langfuse</dt>
          <dd>
            {{if .Snapshot.Langfuse.Ready}}ready
            {{else if .Snapshot.Langfuse.Enabled}}starting
            {{else}}off{{end}}
          </dd>
          <dt>Status</dt>
          <dd>
            {{if .Snapshot.Debug.Error}}
              <span class="subtle">{{.Snapshot.Debug.Error}}</span>
            {{else if .Snapshot.Debug.Enabled}}
              ready
            {{else}}
              idle
            {{end}}
          </dd>
          <dt>Open</dt>
          <dd><a href="/debug">Debug</a></dd>
        </dl>
      </article>

      <article class="card">
        <h2>Browser</h2>
        <p class="subtle">
          Browser providers, managed browser-server state, and profile details
          are grouped on their own page.
        </p>
        <dl class="meta">
          <dt>Providers</dt>
          <dd>{{.Snapshot.Browser.ProviderCount}}</dd>
          <dt>Profiles</dt>
          <dd>{{.Snapshot.Browser.ProfileCount}}</dd>
          <dt>Nodes</dt>
          <dd>{{.Snapshot.Browser.NodeCount}}</dd>
          <dt>Status</dt>
          <dd>
            {{if .Snapshot.Browser.Managed.Enabled}}
              {{if .Snapshot.Browser.Managed.State}}
                {{.Snapshot.Browser.Managed.State}}
              {{else}}
                configured
              {{end}}
            {{else if .Snapshot.Browser.Enabled}}
              ready
            {{else}}
              idle
            {{end}}
          </dd>
          <dt>Open</dt>
          <dd><a href="/browser">Browser</a></dd>
        </dl>
      </article>
    </section>
    {{end}}

    {{if eq .View "config"}}
    <section class="card" style="margin-top: 24px;" id="config-admin">
      <div class="skills-header">
        <div class="skills-header-copy">
          <h2>Runtime Config</h2>
          <p class="skills-lead">
            This page writes directly to the main runtime config file and keeps
            unset fields distinct from explicit values.
          </p>
        </div>
      </div>
      <div class="skills-ops-grid">
        <div class="skills-op-card">
          <div class="skills-op-label">Config file</div>
          <div class="skills-op-value">
            {{if .Config.ConfigPath}}
              <code>{{.Config.ConfigPath}}</code>
            {{else}}
              Runtime config path is not available
            {{end}}
          </div>
          <div class="skills-op-note">
            Reset deletes the YAML key so the runtime falls back to inherited
            behavior.
          </div>
        </div>
        <div class="skills-op-card">
          <div class="skills-op-label">Apply model</div>
          <div class="skills-op-value">Restart required by default</div>
          <div class="skills-op-note">
            The current runtime stays unchanged until you restart, so this page
            shows both the saved config and the live runtime side by side.
          </div>
        </div>
        {{if .ConfigCanRestart}}
        <div class="skills-op-card">
          <div class="skills-op-label">Runtime control</div>
          <div class="skills-op-value">Restart and version actions</div>
          <div class="skills-op-note">
            Use the runtime controls page for graceful restarts, forced
            restarts, and target-version actions.
          </div>
          <div class="skills-op-actions">
            <a class="page-refresh-link" href="/runtime-control">
              Open Runtime Control
            </a>
          </div>
        </div>
        {{end}}
      </div>
      {{if gt .ConfigPending 0}}
      <div class="skills-ops-grid">
        <div class="skills-op-card">
          <div class="skills-op-label">Pending restart</div>
          <div class="skills-op-value">
            {{.ConfigPending}} saved field{{if ne .ConfigPending 1}}s{{end}}
            still need a restart before the running runtime will use them.
          </div>
          <div class="skills-op-note">
            Saved config changes are already on disk. The live runtime stays on
            the previous values until the next restart.
          </div>
          {{if .ConfigCanRestart}}
          <div class="skills-op-actions">
            <a class="page-refresh-link" href="/runtime-control">
              Open Runtime Control
            </a>
          </div>
          {{end}}
        </div>
      </div>
      {{end}}
      {{if .Config.Error}}
      <div class="notice err" style="margin-top: 12px;">
        {{.Config.Error}}
      </div>
      {{end}}
      {{if .Config.Sections}}
      <div class="config-sections">
        {{range .Config.Sections}}
        <section class="config-section-card" id="config-section-{{.Key}}">
          <h3>{{.Title}}</h3>
          {{if .Summary}}
          <p class="subtle">{{.Summary}}</p>
          {{end}}
          <div class="config-field-list">
            {{range .Fields}}
            <details class="config-field" id="config-field-{{.Key}}">
              {{$field := .}}
              <summary class="config-field-summary">
                <div class="config-field-top">
                  <div>
                    <h4 class="config-field-title">{{.Title}}</h4>
                    {{if .Summary}}
                    <p class="subtle">{{.Summary}}</p>
                    {{end}}
                  </div>
                  <div class="config-badges">
                    <span class="config-badge">
                      {{if eq .ApplyMode "hot"}}Hot apply{{else if eq .ApplyMode "next_turn"}}Next turn{{else}}Restart{{end}}
                    </span>
                    {{if .PendingRestart}}
                    <span class="config-badge warn">Pending restart</span>
                    {{end}}
                  </div>
                </div>
              </summary>
              <div class="config-field-detail">
                <div class="config-meta">
                  <div class="config-meta-block">
                    <div class="config-meta-label">Configured</div>
                    <div class="config-meta-value">
                      {{if eq .ConfiguredSource "explicit"}}
                        {{if .ConfiguredValue}}<code>{{.ConfiguredValue}}</code>{{else}}<code>(empty)</code>{{end}}
                      {{else}}
                        <span class="subtle">Inherited</span>
                      {{end}}
                    </div>
                    {{if .ConfiguredSourceLabel}}
                    <div class="subtle">{{.ConfiguredSourceLabel}}</div>
                    {{end}}
                  </div>
                  <div class="config-meta-block">
                    <div class="config-meta-label">Current runtime</div>
                    <div class="config-meta-value">
                      {{if .RuntimeValue}}<code>{{.RuntimeValue}}</code>{{else}}<span class="subtle">-</span>{{end}}
                    </div>
                    {{if .RuntimeSourceLabel}}
                    <div class="subtle">{{.RuntimeSourceLabel}}</div>
                    {{end}}
                  </div>
                </div>
                {{if eq .InputType "readonly"}}
                <p class="subtle" style="margin-top: 12px;">
                  Read-only runtime diagnostics.
                </p>
                {{else}}
                <form method="post" action="/api/config/save" class="config-form">
                  <input type="hidden" name="field_key" value="{{.Key}}">
                  <input type="hidden" name="return_to" value="config-field-{{.Key}}">
                  <input type="hidden" name="return_path" value="/config">
                  {{if eq .InputType "select"}}
                  <select name="value">
                    {{range $field.Options}}
                    <option value="{{.Value}}" {{if eq $field.EditorValue .Value}}selected{{end}}>
                      {{.Label}}
                    </option>
                    {{end}}
                  </select>
                  {{else if eq .InputType "number"}}
                  <input
                    type="number"
                    name="value"
                    value="{{.EditorValue}}"
                    {{if .Placeholder}}placeholder="{{.Placeholder}}"{{end}}
                  >
                  {{else}}
                  <input
                    type="text"
                    name="value"
                    value="{{.EditorValue}}"
                    {{if .Placeholder}}placeholder="{{.Placeholder}}"{{end}}
                  >
                  {{end}}
                  <div class="config-form-row">
                    <button type="submit">Save</button>
                    {{if .Resettable}}
                    <button class="secondary" type="submit" formaction="/api/config/reset">
                      Reset
                    </button>
                    {{end}}
                  </div>
                </form>
                {{end}}
              </div>
            </details>
            {{end}}
          </div>
        </section>
        {{end}}
      </div>
      {{else if not .Config.Error}}
      <p class="subtle" style="margin-top: 12px;">
        Runtime config editing is not available for this runtime.
      </p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "runtime_control"}}
    {{if .RuntimeControl.Error}}
    <div class="notice err" style="margin-top: 24px;">
      {{.RuntimeControl.Error}}
    </div>
    {{end}}
    {{if .RuntimeControl.Enabled}}
    <section class="panels" style="margin-top: 24px;">
      <article class="card" id="runtime-control-state">
        <h2>Runtime State</h2>
        <p class="subtle">
          Review the current version, queue depth, and any lifecycle action
          that is already in flight.
        </p>
        <div class="runtime-meta-grid">
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">Current Version</div>
            <div class="runtime-meta-value">
              {{if .RuntimeControl.Status.CurrentVersion}}
                <code>{{.RuntimeControl.Status.CurrentVersion}}</code>
              {{else}}
                <span class="subtle">-</span>
              {{end}}
            </div>
          </div>
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">Latest Version</div>
            <div class="runtime-meta-value">
              {{if .RuntimeControl.Index.LatestVersion}}
                <code>{{.RuntimeControl.Index.LatestVersion}}</code>
              {{else}}
                <span class="subtle">Unavailable</span>
              {{end}}
            </div>
          </div>
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">State</div>
            <div class="runtime-meta-value">
              {{if .RuntimeControl.Status.State}}
                {{.RuntimeControl.Status.State}}
              {{else}}
                <span class="subtle">Unknown</span>
              {{end}}
            </div>
          </div>
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">Exit Code</div>
            <div class="runtime-meta-value">
              {{.RuntimeControl.Status.ExitCode}}
            </div>
          </div>
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">Running Requests</div>
            <div class="runtime-meta-value">
              {{.RuntimeControl.Status.RunningRequests}}
            </div>
          </div>
          <div class="runtime-meta-card">
            <div class="runtime-meta-label">Queued Requests</div>
            <div class="runtime-meta-value">
              {{.RuntimeControl.Status.QueuedRequests}}
            </div>
          </div>
        </div>
        {{if .RuntimeControl.Status.Pending}}
        <div class="runtime-pending">
          <h3>Pending Action</h3>
          <dl class="meta">
            <dt>Kind</dt>
            <dd>{{.RuntimeControl.Status.Pending.Kind}}</dd>
            <dt>Mode</dt>
            <dd>{{.RuntimeControl.Status.Pending.Mode}}</dd>
            {{if .RuntimeControl.Status.Pending.TargetVersion}}
            <dt>Target Version</dt>
            <dd>
              <code>{{.RuntimeControl.Status.Pending.TargetVersion}}</code>
            </dd>
            {{end}}
            {{if .RuntimeControl.Status.Pending.Actor}}
            <dt>Actor</dt>
            <dd>{{.RuntimeControl.Status.Pending.Actor}}</dd>
            {{end}}
            {{if .RuntimeControl.Status.Pending.Source}}
            <dt>Source</dt>
            <dd>{{.RuntimeControl.Status.Pending.Source}}</dd>
            {{end}}
            {{if hasTime .RuntimeControl.Status.Pending.RequestedAt}}
            <dt>Requested</dt>
            <dd>{{formatTime .RuntimeControl.Status.Pending.RequestedAt}}</dd>
            {{end}}
            {{if hasTime .RuntimeControl.Status.UpdatedAt}}
            <dt>Updated</dt>
            <dd>{{formatTime .RuntimeControl.Status.UpdatedAt}}</dd>
            {{end}}
          </dl>
          {{if .RuntimeControl.Status.Pending.Summary}}
          <ul class="runtime-summary-list">
            {{range .RuntimeControl.Status.Pending.Summary}}
            <li>{{.}}</li>
            {{end}}
          </ul>
          {{end}}
        </div>
        {{else}}
        <p class="subtle" style="margin-top: 14px;">
          No restart or version switch is currently pending.
        </p>
        {{end}}
      </article>

      <article class="card" id="runtime-control-quick-actions">
        <h2>Quick Actions</h2>
        <p class="subtle">
          These controls line up with the runtime lifecycle card behavior:
          graceful actions wait for in-flight work, forced actions exit after
          the forced-shutdown window.
        </p>
        <div class="skills-ops-grid">
          <div class="skills-op-card">
            <div class="skills-op-label">Restart</div>
            <div class="skills-op-value">Graceful restart</div>
            <div class="skills-op-note">
              Stop admitting new work, drain active requests, and then restart.
            </div>
            <div class="skills-op-actions">
              <form method="post" action="/api/runtime/control/action">
                <input type="hidden" name="kind" value="restart">
                <input type="hidden" name="mode" value="graceful">
                <input type="hidden" name="return_path" value="/runtime-control">
                <input type="hidden" name="return_to" value="runtime-control-quick-actions">
                <button type="submit">Graceful Restart</button>
              </form>
            </div>
          </div>
          <div class="skills-op-card">
            <div class="skills-op-label">Restart</div>
            <div class="skills-op-value">Force restart</div>
            <div class="skills-op-note">
              Request a restart and exit once the forced shutdown timeout lands.
            </div>
            <div class="skills-op-actions">
              <form method="post" action="/api/runtime/control/action">
                <input type="hidden" name="kind" value="restart">
                <input type="hidden" name="mode" value="force">
                <input type="hidden" name="return_path" value="/runtime-control">
                <input type="hidden" name="return_to" value="runtime-control-quick-actions">
                <button class="warn" type="submit">Force Restart</button>
              </form>
            </div>
          </div>
          <div class="skills-op-card">
            <div class="skills-op-label">Upgrade</div>
            <div class="skills-op-value">Latest version</div>
            <div class="skills-op-note">
              Ask the runtime to switch to the latest published release.
            </div>
            <div class="skills-op-actions">
              <form method="post" action="/api/runtime/control/action">
                <input type="hidden" name="kind" value="upgrade">
                <input type="hidden" name="mode" value="graceful">
                <input type="hidden" name="return_path" value="/runtime-control">
                <input type="hidden" name="return_to" value="runtime-control-quick-actions">
                <button type="submit">Upgrade to Latest</button>
              </form>
            </div>
          </div>
          <div class="skills-op-card">
            <div class="skills-op-label">Upgrade</div>
            <div class="skills-op-value">Force latest upgrade</div>
            <div class="skills-op-note">
              Request the latest version and force the handoff after the drain
              window ends.
            </div>
            <div class="skills-op-actions">
              <form method="post" action="/api/runtime/control/action">
                <input type="hidden" name="kind" value="upgrade">
                <input type="hidden" name="mode" value="force">
                <input type="hidden" name="return_path" value="/runtime-control">
                <input type="hidden" name="return_to" value="runtime-control-quick-actions">
                <button class="warn" type="submit">Force Upgrade to Latest</button>
              </form>
            </div>
          </div>
        </div>
      </article>
    </section>

    <section class="panels" style="margin-top: 16px;">
      <article class="card runtime-version-card" id="runtime-control-version">
        <h2>Target Version</h2>
        <p class="subtle">
          Pick a release to inspect notes or request a direct switch to that
          version.
        </p>
        {{if .RuntimeControl.Index.MinSupportedTarget}}
        <p class="subtle">
          Minimum supported target:
          <code>{{.RuntimeControl.Index.MinSupportedTarget}}</code>
        </p>
        {{end}}
        {{if .RuntimeControl.Index.Versions}}
        <form method="get" action="/runtime-control" class="runtime-version-view-form">
          <select name="version">
            {{range .RuntimeControl.Index.Versions}}
            <option value="{{.Version}}" {{if eq $.RuntimeControl.SelectedVersion .Version}}selected{{end}}>
              {{.Version}}
            </option>
            {{end}}
          </select>
          <div class="runtime-version-actions">
            <button class="secondary" type="submit">View Release Notes</button>
          </div>
        </form>
        <form method="post" action="/api/runtime/control/action" class="runtime-version-form">
          <input type="hidden" name="kind" value="upgrade">
          <input type="hidden" name="return_path" value="/runtime-control">
          <input type="hidden" name="return_to" value="runtime-control-version">
          <select name="target_version">
            {{range .RuntimeControl.Index.Versions}}
            <option value="{{.Version}}" {{if eq $.RuntimeControl.SelectedVersion .Version}}selected{{end}}>
              {{.Version}}
            </option>
            {{end}}
          </select>
          <div class="runtime-version-actions">
            <button type="submit" name="mode" value="graceful">
              Switch Gracefully
            </button>
            <button class="warn" type="submit" name="mode" value="force">
              Switch Forcefully
            </button>
          </div>
        </form>
        {{else}}
        <p class="subtle" style="margin-top: 14px;">
          No published versions are available from the configured release
          source.
        </p>
        {{end}}
      </article>

      <article class="card runtime-changelog-card" id="runtime-control-changelog">
        <h2>Release Notes</h2>
        {{if .RuntimeControl.Changelog.Version}}
        <p class="subtle">
          Notes for <code>{{.RuntimeControl.Changelog.Version}}</code>.
        </p>
        {{else if .RuntimeControl.SelectedVersion}}
        <p class="subtle">
          Notes for <code>{{.RuntimeControl.SelectedVersion}}</code> are not
          available yet.
        </p>
        {{else}}
        <p class="subtle">
          Pick a version to inspect release notes.
        </p>
        {{end}}
        {{if .RuntimeControl.Changelog.Summary}}
        <ul class="runtime-summary-list">
          {{range .RuntimeControl.Changelog.Summary}}
          <li>{{.}}</li>
          {{end}}
        </ul>
        {{end}}
        {{if .RuntimeControl.Changelog.Changelog}}
        <div class="runtime-changelog">
          {{.RuntimeControl.Changelog.Changelog}}
        </div>
        {{else}}
        <p class="subtle" style="margin-top: 14px;">
          The selected version does not currently have release notes.
        </p>
        {{end}}
      </article>
    </section>
    {{else if not .RuntimeControl.Error}}
    <section class="card" style="margin-top: 24px;">
      <h2>Runtime Control</h2>
      <p class="subtle">
        Runtime lifecycle controls are not available for this runtime.
      </p>
    </section>
    {{end}}
    {{end}}

    {{if eq .View "skills"}}
    <section class="card" style="margin-top: 24px;" id="skills-admin" data-skills-root>
      <div class="skills-header">
        <div class="skills-header-copy">
          <h2>Skills Inventory</h2>
          <p class="skills-lead">
            Bundled, local, project, and external skills discovered by this
            runtime.
          </p>
        </div>
      </div>
      <div class="skills-ops-grid">
        <div class="skills-op-card">
          <div class="skills-op-label">
            {{if .Snapshot.Skills.Writable}}Config-backed changes{{else}}Runtime state{{end}}
          </div>
          <div class="skills-op-value">
            {{if .Snapshot.Skills.Writable}}
              Writes to <code>{{.Snapshot.Skills.ConfigPath}}</code>
            {{else}}
              Read-only runtime view
            {{end}}
          </div>
          <div class="skills-op-note">
            {{if .Snapshot.Skills.Refreshable}}
              Enabled changes apply on the next turn.
            {{else if .Snapshot.Skills.Writable}}
              Enabled changes are saved, but runtime updates still require a restart.
            {{else}}
              Enable and disable controls are unavailable for this runtime.
            {{end}}
          </div>
        </div>
        {{if .Snapshot.Skills.Refreshable}}
        <div class="skills-op-card">
          <div class="skills-op-label">Refresh from disk</div>
          <div class="skills-op-value">
            Rescan skill folders and update this inventory.
          </div>
          <div class="skills-op-note">
            Use this after adding or removing skill folders on disk.
          </div>
          <div class="skills-op-actions">
            <form method="post" action="/api/skills/refresh">
              <input type="hidden" name="return_to" value="skills-admin">
              <input type="hidden" name="return_path" value="/skills">
              <button class="secondary" type="submit">Refresh inventory</button>
            </form>
          </div>
        </div>
        {{end}}
      </div>
      {{if .Snapshot.Skills.Error}}
      <div class="notice err" style="margin-top: 12px;">
        {{.Snapshot.Skills.Error}}
      </div>
      {{end}}
      {{if .Snapshot.Skills.Groups}}
      <div class="filter-tabs">
        <button class="filter-tab active" type="button" data-skill-tab="all">
          All {{.Snapshot.Skills.TotalCount}}
        </button>
        <button class="filter-tab" type="button" data-skill-tab="ready">
          Ready {{.Snapshot.Skills.ReadyCount}}
        </button>
        <button class="filter-tab" type="button" data-skill-tab="needs-setup">
          Needs Setup {{.Snapshot.Skills.NeedsSetupCount}}
        </button>
        <button class="filter-tab" type="button" data-skill-tab="disabled">
          Disabled {{.Snapshot.Skills.DisabledCount}}
        </button>
      </div>
      <div class="skills-controls">
        <div class="skills-search-wrap">
          <input
            type="search"
            placeholder="Search skills by name, path, key, env, or reason"
            data-skills-filter
          >
        </div>
        <div class="skills-toolbar-side">
          <span class="skills-shown"><span data-skills-shown>{{.Snapshot.Skills.TotalCount}}</span> shown</span>
        </div>
      </div>

      {{range .Snapshot.Skills.Groups}}
      <div class="skills-group" data-skills-group id="skills-group-{{.ID}}">
        <h3>{{.Label}}</h3>
        {{range .Skills}}
        <details
          class="skill-card"
          id="skill-card-{{.ConfigKey}}"
          data-skill-card
          data-skill-status="{{.Status}}"
          data-skill-search="{{.SearchText}}"
        >
          <summary>
            <div class="skill-main">
              <div class="skill-copy">
                <div class="skill-headline">
                  <div class="skill-name">
                    <span class="skill-dot {{.Status}}"></span>
                    {{if .Emoji}}<span>{{.Emoji}}</span>{{end}}
                    <span>{{.Name}}</span>
                  </div>
                </div>
                <div class="skill-badges inline">
                  {{if .Bundled}}<span class="skill-badge">bundled</span>{{end}}
                  {{if .BlockedByAllowlist}}<span class="skill-badge">allowlist</span>{{end}}
                  {{if .Always}}<span class="skill-badge">always</span>{{end}}
                  {{if .PrimaryEnv}}<span class="skill-badge">{{.PrimaryEnv}}</span>{{end}}
                </div>
                <div class="skill-description">{{.Description}}</div>
                {{if .Reason}}
                <div class="skill-reason">
                  <span class="skill-reason-label {{.Status}}">
                    {{if eq .Status "needs-setup"}}Setup Required{{else if eq .Status "disabled"}}Disabled{{else}}Reason{{end}}
                  </span>
                  <span class="skill-reason-text">{{.Reason}}</span>
                </div>
                {{end}}
              </div>
              <div class="skill-summary-side">
                {{if $.Snapshot.Skills.Writable}}
                <div class="skill-toggle-group">
                  <form
                    method="post"
                    action="/api/skills/toggle"
                    class="skill-inline-toggle-form"
                    data-skill-inline-toggle
                    data-skill-config="{{.ConfigKey}}"
                  >
                    <input type="hidden" name="skill_key" value="{{.ConfigKey}}">
                    <input type="hidden" name="skill_name" value="{{.Name}}">
                    <input type="hidden" name="enabled" value="{{if .Disabled}}true{{else}}false{{end}}">
                    <input type="hidden" name="return_to" value="skill-card-{{.ConfigKey}}">
                    <input type="hidden" name="return_path" value="/skills">
                    <button
                      class="skill-inline-toggle {{if not .Disabled}}enabled{{end}}"
                      type="submit"
                      data-skill-toggle-button
                      data-skill-switch="{{.ConfigKey}}"
                      role="switch"
                      aria-checked="{{if .Disabled}}false{{else}}true{{end}}"
                      aria-label="Enabled for {{.Name}}"
                      title="{{if .Disabled}}Enable{{else}}Disable{{end}} {{.Name}}"
                    >
                      <span class="skill-inline-toggle-track" aria-hidden="true"></span>
                    </button>
                  </form>
                </div>
                {{end}}
              </div>
            </div>
          </summary>
          <div class="skill-details">
            <div class="skill-details-head">
              <div class="subtle">
                {{if and $.Snapshot.Skills.Writable $.Snapshot.Skills.Refreshable}}
                The row-level Enabled switch saves
                <code>skills.entries.{{.ConfigKey}}.enabled</code> and refreshes
                this runtime for the next turn.
                {{else if $.Snapshot.Skills.Writable}}
                The row-level Enabled switch saves
                <code>skills.entries.{{.ConfigKey}}.enabled</code> for this
                skill.
                {{else}}
                Enable/disable controls are unavailable for this runtime.
                {{end}}
              </div>
            </div>
            <div class="skill-details-grid">
              <div>
                <strong>Skill Key</strong>
                <div><code>{{.SkillKey}}</code></div>
              </div>
              <div>
                <strong>Config Key</strong>
                <div><code>{{.ConfigKey}}</code></div>
              </div>
              <div>
                <strong>Source</strong>
                <div>{{if .Source}}{{.Source}}{{else}}unknown{{end}}</div>
              </div>
              <div>
                <strong>Primary Env</strong>
                <div>{{if .PrimaryEnv}}<code>{{.PrimaryEnv}}</code>{{else}}-{{end}}</div>
              </div>
              <div>
                <strong>Path</strong>
                <div><code>{{.FilePath}}</code></div>
              </div>
              <div>
                <strong>Homepage</strong>
                <div>
                  {{if .Homepage}}
                  <a href="{{.Homepage}}" target="_blank" rel="noopener noreferrer">{{.Homepage}}</a>
                  {{else}}-{{end}}
                </div>
              </div>
            </div>

            {{if or .Missing.Bins .Missing.AnyBins .Missing.Env .Missing.Config .Missing.OS}}
            <div style="margin-top: 14px;">
              <strong>Missing Requirements</strong>
              <ul class="skill-list">
                {{if .Missing.Bins}}
                <li>bins:
                  {{range $i, $item := .Missing.Bins}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Missing.AnyBins}}
                <li>one of:
                  {{range $i, $item := .Missing.AnyBins}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Missing.Env}}
                <li>env:
                  {{range $i, $item := .Missing.Env}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Missing.Config}}
                <li>config:
                  {{range $i, $item := .Missing.Config}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Missing.OS}}
                <li>os:
                  {{range $i, $item := .Missing.OS}}{{if $i}}, {{end}}{{$item}}{{end}}
                </li>
                {{end}}
              </ul>
            </div>
            {{end}}

            {{if or .Requirements.Bins .Requirements.AnyBins .Requirements.Env .Requirements.Config .Requirements.OS}}
            <div style="margin-top: 14px;">
              <strong>Declared Requirements</strong>
              <ul class="skill-list">
                {{if .Requirements.Bins}}
                <li>bins:
                  {{range $i, $item := .Requirements.Bins}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Requirements.AnyBins}}
                <li>one of:
                  {{range $i, $item := .Requirements.AnyBins}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Requirements.Env}}
                <li>env:
                  {{range $i, $item := .Requirements.Env}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Requirements.Config}}
                <li>config:
                  {{range $i, $item := .Requirements.Config}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}}
                </li>
                {{end}}
                {{if .Requirements.OS}}
                <li>os:
                  {{range $i, $item := .Requirements.OS}}{{if $i}}, {{end}}{{$item}}{{end}}
                </li>
                {{end}}
              </ul>
            </div>
            {{end}}

            {{if .Install}}
            <div style="margin-top: 14px;">
              <strong>Suggested Installers</strong>
              <ul class="skill-list">
                {{range .Install}}
                <li>
                  <code>{{.Label}}</code>
                  {{if .Bins}}
                  <span class="subtle">
                    (provides {{range $i, $item := .Bins}}{{if $i}}, {{end}}<code>{{$item}}</code>{{end}})
                  </span>
                  {{end}}
                </li>
                {{end}}
              </ul>
            </div>
            {{end}}
          </div>
        </details>
        {{end}}
      </div>
      {{end}}
      {{else if not .Snapshot.Skills.Error}}
      <p class="empty">No skills discovered.</p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "prompts"}}
    {{template "promptsPage" .}}
    {{end}}

    {{if eq .View "identity"}}
    {{template "identityPage" .}}
    {{end}}

    {{if eq .View "personas"}}
    {{template "personasPage" .}}
    {{end}}

    {{if eq .View "chats"}}
    {{template "chatsPage" .}}
    {{end}}

    {{if eq .View "memory"}}
    <section class="panels">
      <article class="card">
        <h2>Memory Backend</h2>
        <dl class="meta">
          <dt>Backend</dt>
          <dd>{{if .Snapshot.Memory.Backend}}{{.Snapshot.Memory.Backend}}{{else}}-{{end}}</dd>
          <dt>Storage Mode</dt>
          <dd>
            {{if eq .Snapshot.Memory.Backend "file"}}
              {{if .Snapshot.Memory.FileEnabled}}
                File-backed <code>MEMORY.md</code>
              {{else}}
                file backend not configured
              {{end}}
            {{else if .Snapshot.Memory.FileEnabled}}
              File-backed <code>MEMORY.md</code>
            {{else if .Snapshot.Memory.Enabled}}
              Structured memory service
            {{else}}
              unavailable
            {{end}}
          </dd>
          <dt>Structured Memory</dt>
          <dd>
            {{if eq .Snapshot.Memory.Backend "file"}}
              not used by file backend
            {{else if and .Snapshot.Memory.Enabled (not .Snapshot.Memory.FileEnabled)}}
              enabled
            {{else}}
              unavailable
            {{end}}
          </dd>
          <dt>File Inventory</dt>
          <dd>
            {{if eq .Snapshot.Memory.Backend "file"}}
              {{if .Snapshot.Memory.FileEnabled}}available{{else}}not configured{{end}}
            {{else}}
              not available
            {{end}}
          </dd>
          <dt>JSON</dt>
          <dd><a href="/api/memory/files">/api/memory/files</a></dd>
        </dl>
      </article>
      <article class="card">
        <h2>File Inventory</h2>
        <dl class="meta">
          <dt>Root</dt>
          <dd>
            {{if .Snapshot.Memory.Root}}
              <code>{{.Snapshot.Memory.Root}}</code>
            {{else}}
              -
            {{end}}
          </dd>
          <dt>Files</dt>
          <dd>{{.Snapshot.Memory.FileCount}}</dd>
          <dt>Total Bytes</dt>
          <dd>{{.Snapshot.Memory.TotalBytes}}</dd>
          <dt>Last Modified</dt>
          <dd>{{formatTime .Snapshot.Memory.LastModified}}</dd>
        </dl>
      </article>
    </section>

    <section class="card" style="margin-top: 24px;">
      <h2>Memory Files</h2>
      <p class="subtle">
        File-backed memory stores one visible <code>MEMORY.md</code> per
        app/user scope. Use this inventory to inspect what durable memory the
        runtime can inject into future turns.
      </p>
      {{if .Snapshot.Memory.Error}}
      <div class="notice err" style="margin-top: 12px;">
        {{.Snapshot.Memory.Error}}
      </div>
      {{end}}
      {{if .Snapshot.Memory.Files}}
      <div class="memory-controls" data-memory-root>
        <div class="memory-filter-head">
          <label for="memory-search">Search memory</label>
          <span class="memory-shown">
            <span data-memory-shown>{{.Snapshot.Memory.FileCount}}</span> shown
          </span>
        </div>
        <div class="memory-filter-grid">
          <div class="memory-filter-field">
            <input
              id="memory-search"
              type="search"
              placeholder="Search app, user, path, or preview"
              data-memory-search
            >
          </div>
        </div>
      </div>
      <div class="memory-list">
        {{range .Snapshot.Memory.Files}}
        <details
          class="card prompt-detail"
          id="{{.CardID}}"
          data-memory-card
          data-memory-path="{{.RelativePath}}"
          data-memory-load-url="{{.LoadURL}}"
          data-memory-search="{{.SearchValue}}"
        >
          <summary>
            <div class="memory-card-head">
              <div>
                <div class="memory-scope">
                  <span>app <code>{{.AppName}}</code></span>
                  <span>
                    user <code>{{.UserID}}</code>
                    {{if .UserLabel}}
                    <span class="memory-user-label">{{.UserLabel}}</span>
                    {{end}}
                  </span>
                </div>
                <div class="memory-path">
                  <code>{{.RelativePath}}</code>
                </div>
                <p class="subtle prompt-detail-hint">
                  {{if .Preview}}
                    {{.Preview}}
                  {{else}}
                    No visible memory content yet.
                  {{end}}
                </p>
              </div>
              <div class="memory-card-meta">
                <span>{{.SizeBytes}} bytes</span>
                <span>{{formatTime .ModifiedAt}}</span>
              </div>
            </div>
          </summary>
          <div class="prompt-detail-body">
            <div class="memory-card-toolbar">
              <div class="subtle">
                Edit the full <code>MEMORY.md</code> file for this scope.
              </div>
              <a
                href="{{.OpenURL}}"
                target="_blank"
                rel="noopener noreferrer"
              >
                Open Raw File
              </a>
            </div>
            <form
              class="memory-editor-form"
              method="post"
              action="/api/memory/file"
              data-memory-form
            >
              <input type="hidden" name="path" value="{{.RelativePath}}">
              <input type="hidden" name="return_path" value="/memory">
              <input type="hidden" name="return_to" value="{{.CardID}}">
              <label for="editor-{{.CardID}}">
                Full file content
              </label>
              <textarea
                id="editor-{{.CardID}}"
                name="content"
                data-memory-editor
                placeholder="Open this scope to load MEMORY.md"
                disabled
              ></textarea>
              <div class="memory-editor-status" data-memory-status>
                Open this scope to load its full file.
              </div>
              <div
                class="notice err"
                style="margin-top: 12px;"
                data-memory-error
                hidden
              ></div>
              <div class="actions" style="margin-top: 12px;">
                <button type="submit" data-memory-save disabled>
                  Save File
                </button>
                <button
                  class="secondary"
                  type="button"
                  data-memory-reset
                  disabled
                >
                  Revert
                </button>
              </div>
            </form>
          </div>
        </details>
        {{end}}
      </div>
      <p class="empty" data-memory-empty hidden>No matching memory files.</p>
      {{else if not .Snapshot.Memory.Error}}
      <p class="empty">
        {{if eq .Snapshot.Memory.Backend "file"}}
          {{if .Snapshot.Memory.FileEnabled}}
            No file-backed memory files discovered yet.
          {{else}}
            File-backed memory store is not configured for this runtime.
          {{end}}
        {{else if .Snapshot.Memory.FileEnabled}}
          No file-backed memory files discovered yet.
        {{else}}
          File-backed memory inventory is only available when the runtime uses
          the <code>file</code> memory backend.
        {{end}}
      </p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "automation"}}
    <section class="panels">
      <article class="card">
        <h2>Automation</h2>
        {{if .Snapshot.Cron.Enabled}}
        <p class="subtle">
          Persisted jobs continue after gateway restarts. Use this page for
          scheduling, one-off runs, and cleanup.
        </p>
        <div class="actions">
          <form method="post" action="/api/cron/jobs/clear">
            <input type="hidden" name="return_path" value="/automation">
            <button
              class="warn"
              type="submit"
              onclick="return confirm('Clear all scheduled jobs?');"
            >
              Clear All Jobs
            </button>
          </form>
        </div>
        {{else}}
        <p class="empty">Scheduled jobs are not enabled for this runtime.</p>
        {{end}}
      </article>
    </section>
    <section class="card" style="margin-top: 24px;">
      <h2>Scheduled Jobs</h2>
      {{if .Snapshot.Cron.Jobs}}
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Schedule</th>
            <th>Delivery</th>
            <th>Status</th>
            <th>Timing</th>
            <th>Task</th>
            <th>Last Output</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Cron.Jobs}}
          <tr>
            <td>
              <strong>{{.Name}}</strong><br>
              <code>{{.ID}}</code><br>
              owner {{.UserID}}
            </td>
            <td>{{.Schedule}}</td>
            <td>
              {{if .Channel}}
                {{.Channel}} → {{.Target}}
              {{else}}
                no delivery target
              {{end}}
            </td>
            <td>
              {{if .LastStatus}}{{.LastStatus}}{{else}}idle{{end}}
              {{if .LastError}}
                <br>
                <span class="subtle">{{.LastError}}</span>
              {{end}}
            </td>
            <td>
              next {{formatTime .NextRunAt}}<br>
              last {{formatTime .LastRunAt}}
            </td>
            <td>{{.MessagePreview}}</td>
            <td>{{if .LastOutput}}{{.LastOutput}}{{else}}-{{end}}</td>
            <td>
              <div class="actions">
                <form method="post" action="/api/cron/jobs/run">
                  <input type="hidden" name="job_id" value="{{.ID}}">
                  <input type="hidden" name="return_path" value="/automation">
                  <button type="submit">Run Now</button>
                </form>
                <form method="post" action="/api/cron/jobs/remove">
                  <input type="hidden" name="job_id" value="{{.ID}}">
                  <input type="hidden" name="return_path" value="/automation">
                  <button
                    class="secondary"
                    type="submit"
                    onclick="return confirm('Remove this job?');"
                  >
                    Remove
                  </button>
                </form>
              </div>
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No scheduled jobs.</p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "sessions"}}
    <section class="panels">
      <article class="card">
        <h2>Exec Surface</h2>
        <dl class="meta">
          <dt>Enabled</dt>
          <dd>{{.Snapshot.Exec.Enabled}}</dd>
          <dt>Sessions</dt>
          <dd>{{.Snapshot.Exec.SessionCount}}</dd>
          <dt>Running</dt>
          <dd>{{.Snapshot.Exec.RunningCount}}</dd>
          <dt>JSON</dt>
          <dd><a href="/api/exec/sessions">/api/exec/sessions</a></dd>
        </dl>
      </article>
      <article class="card">
        <h2>Uploads</h2>
        <dl class="meta">
          <dt>Enabled</dt>
          <dd>{{.Snapshot.Uploads.Enabled}}</dd>
          <dt>Root</dt>
          <dd><code>{{.Snapshot.Uploads.Root}}</code></dd>
          <dt>Files</dt>
          <dd>{{.Snapshot.Uploads.FileCount}}</dd>
          <dt>Total Bytes</dt>
          <dd>{{.Snapshot.Uploads.TotalBytes}}</dd>
        </dl>
      </article>
    </section>
    <section class="card" style="margin-top: 24px;">
      <h2>Exec Sessions</h2>
      {{if .Snapshot.Exec.Sessions}}
      <table>
        <thead>
          <tr>
            <th>Session</th>
            <th>Status</th>
            <th>Command</th>
            <th>Timing</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Exec.Sessions}}
          <tr>
            <td><code>{{.SessionID}}</code></td>
            <td>{{.Status}}{{if .ExitCode}}<br>exit {{.ExitCode}}{{end}}</td>
            <td><code>{{.Command}}</code></td>
            <td>
              started {{.StartedAt}}
              {{if .DoneAt}}<br>done {{.DoneAt}}{{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No exec sessions.</p>
      {{end}}
    </section>

    <section class="card" style="margin-top: 24px;">
      <h2>Upload Sessions</h2>
      {{if .Snapshot.Uploads.Sessions}}
      <table>
        <thead>
          <tr>
            <th>Channel</th>
            <th>User</th>
            <th>Session</th>
            <th>Files</th>
            <th>Total Bytes</th>
            <th>Last Modified</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Uploads.Sessions}}
          <tr>
            <td>{{.Channel}}</td>
            <td><code>{{.UserID}}</code></td>
            <td>
              <a href="/api/uploads?session_id={{urlquery .SessionID}}">
                <code>{{.SessionID}}</code>
              </a>
            </td>
            <td>{{.FileCount}}</td>
            <td>{{.TotalBytes}}</td>
            <td>{{formatTime .LastModified}}</td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No upload sessions indexed yet.</p>
      {{end}}
    </section>

    <section class="card" style="margin-top: 24px;">
      <h2>Recent Uploads</h2>
      {{if .Snapshot.Uploads.Files}}
      <table>
        <thead>
          <tr>
            <th>Channel</th>
            <th>User</th>
            <th>Session</th>
            <th>Name</th>
            <th>Kind</th>
            <th>MIME</th>
            <th>Source</th>
            <th>Preview</th>
            <th>Relative Path</th>
            <th>Size</th>
            <th>Modified</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Uploads.Files}}
          <tr>
            <td>{{.Channel}}</td>
            <td><code>{{.UserID}}</code></td>
            <td>
              <a href="/api/uploads?session_id={{urlquery .SessionID}}">
                <code>{{.SessionID}}</code>
              </a>
            </td>
            <td>
              <a href="{{.OpenURL}}" target="_blank"
                rel="noopener noreferrer">{{.Name}}</a>
              <br>
              <a href="{{.DownloadURL}}">download</a>
            </td>
            <td>
              <a href="/api/uploads?kind={{urlquery .Kind}}">
                {{.Kind}}
              </a>
            </td>
            <td>
              {{if .MimeType}}
              <a href="/api/uploads?mime_type={{urlquery .MimeType}}">
                <code>{{.MimeType}}</code>
              </a>
              {{else}}
              <span class="subtle">-</span>
              {{end}}
            </td>
            <td>
              {{if .Source}}
              <a href="/api/uploads?source={{urlquery .Source}}">
                {{.Source}}
              </a>
              {{else}}
              <span class="subtle">-</span>
              {{end}}
            </td>
            <td>
              <div class="preview-box">
                {{if eq .Kind "image"}}
                <a href="{{.OpenURL}}" target="_blank"
                  rel="noopener noreferrer">
                  <img src="{{.OpenURL}}" alt="{{.Name}}">
                </a>
                {{else if eq .Kind "audio"}}
                <audio controls preload="none" src="{{.OpenURL}}">
                  Your browser does not support audio preview.
                </audio>
                {{else if eq .Kind "video"}}
                <video controls preload="metadata" muted src="{{.OpenURL}}">
                  Your browser does not support video preview.
                </video>
                {{else if eq .Kind "pdf"}}
                <a href="{{.OpenURL}}" target="_blank"
                  rel="noopener noreferrer">open preview</a>
                {{else}}
                <span class="subtle">n/a</span>
                {{end}}
              </div>
            </td>
            <td><code>{{.RelativePath}}</code></td>
            <td>{{.SizeBytes}}</td>
            <td>{{formatTime .ModifiedAt}}</td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No uploads indexed yet.</p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "debug"}}
    <section class="panels">
      <article class="card">
        <h2>Debug Index</h2>
        <dl class="meta">
          <dt>Indexed Dir</dt>
          <dd><code>{{.Snapshot.Debug.BySessionDir}}</code></dd>
          <dt>Session Count</dt>
          <dd>{{.Snapshot.Debug.SessionCount}}</dd>
          <dt>Trace Count</dt>
          <dd>{{.Snapshot.Debug.TraceCount}}</dd>
          <dt>Status</dt>
          <dd>
            {{if .Snapshot.Debug.Error}}
              <span class="subtle">{{.Snapshot.Debug.Error}}</span>
            {{else if .Snapshot.Debug.Enabled}}
              ready
            {{else}}
              idle
            {{end}}
          </dd>
        </dl>
      </article>
      <article class="card">
        <h2>Langfuse</h2>
        <dl class="meta">
          <dt>Enabled</dt>
          <dd>{{.Snapshot.Langfuse.Enabled}}</dd>
          <dt>Ready</dt>
          <dd>{{.Snapshot.Langfuse.Ready}}</dd>
          <dt>Status</dt>
          <dd>
            {{if .Snapshot.Langfuse.Error}}
              <span class="subtle">{{.Snapshot.Langfuse.Error}}</span>
            {{else if .Snapshot.Langfuse.Ready}}
              ready
            {{else if .Snapshot.Langfuse.Enabled}}
              starting
            {{else}}
              idle
            {{end}}
          </dd>
        </dl>
      </article>
    </section>
    <section class="card" style="margin-top: 24px;">
      <h2>Debug Sessions</h2>
      {{if .Snapshot.Debug.Sessions}}
      <table>
        <thead>
          <tr>
            <th>Session</th>
            <th>Trace Count</th>
            <th>Last Seen</th>
            <th>Latest Trace</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Debug.Sessions}}
          <tr>
            <td><code>{{.SessionID}}</code></td>
            <td>{{.TraceCount}}</td>
            <td>
              {{formatTime .LastTraceAt}}<br>
              {{if .Channel}}{{.Channel}}{{end}}
              {{if .RequestID}}
                <br>
                <span class="subtle">{{.RequestID}}</span>
              {{end}}
              {{if .TraceID}}
                <br>
                <span class="subtle">trace {{.TraceID}}</span>
              {{end}}
            </td>
            <td>
              {{if .LangfuseURL}}
                <a href="{{.LangfuseURL}}" target="_blank"
                  rel="noopener noreferrer">langfuse</a> ·
              {{end}}
              {{if .MetaURL}}
                <a href="{{.MetaURL}}" target="_blank">meta</a>
              {{end}}
              {{if .EventsURL}}
                · <a href="{{.EventsURL}}" target="_blank">events</a>
              {{end}}
              {{if .ResultURL}}
                · <a href="{{.ResultURL}}" target="_blank">result</a>
              {{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No debug sessions indexed yet.</p>
      {{end}}
    </section>

    <section class="card" style="margin-top: 24px;">
      <h2>Recent Traces</h2>
      {{if .Snapshot.Debug.RecentTraces}}
      <table>
        <thead>
          <tr>
            <th>Session</th>
            <th>Started</th>
            <th>Request</th>
            <th>Files</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Debug.RecentTraces}}
          <tr>
            <td><code>{{.SessionID}}</code></td>
            <td>{{formatTime .StartedAt}}</td>
            <td>
              {{if .Channel}}{{.Channel}}{{else}}-{{end}}
              {{if .RequestID}}
                <br>
                <span class="subtle">{{.RequestID}}</span>
              {{end}}
              {{if .MessageID}}
                <br>
                <span class="subtle">msg {{.MessageID}}</span>
              {{end}}
              {{if .TraceID}}
                <br>
                <span class="subtle">trace {{.TraceID}}</span>
              {{end}}
            </td>
            <td>
              {{if .LangfuseURL}}
                <a href="{{.LangfuseURL}}" target="_blank"
                  rel="noopener noreferrer">langfuse</a> ·
              {{end}}
              {{if .MetaURL}}
                <a href="{{.MetaURL}}" target="_blank">meta</a>
              {{end}}
              {{if .EventsURL}}
                · <a href="{{.EventsURL}}" target="_blank">events</a>
              {{end}}
              {{if .ResultURL}}
                · <a href="{{.ResultURL}}" target="_blank">result</a>
              {{end}}
              {{if .TracePath}}
                <br>
                <span class="subtle">
                  <code>{{.TracePath}}</code>
                </span>
              {{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">No recent traces.</p>
      {{end}}
    </section>
    {{end}}

    {{if eq .View "browser"}}
    <section class="card" style="margin-top: 24px;">
      <h2>Browser Surface</h2>
      <p class="subtle">
        Native browser tool wiring, including host browser-server routing,
        sandbox targets, node targets, and profile inventory.
      </p>
      <dl class="meta">
        <dt>Enabled</dt>
        <dd>{{.Snapshot.Browser.Enabled}}</dd>
        <dt>Providers</dt>
        <dd>{{.Snapshot.Browser.ProviderCount}}</dd>
        <dt>Profiles</dt>
        <dd>{{.Snapshot.Browser.ProfileCount}}</dd>
        <dt>Nodes</dt>
        <dd>{{.Snapshot.Browser.NodeCount}}</dd>
        <dt>Status</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.Enabled}}
            {{if .Snapshot.Browser.Managed.State}}
              {{.Snapshot.Browser.Managed.State}}
            {{else}}
              configured
            {{end}}
          {{else if .Snapshot.Browser.Enabled}}
            ready
          {{else}}
            idle
          {{end}}
        </dd>
      </dl>
      {{if .Snapshot.Browser.Managed.Enabled}}
      <h3 style="margin: 16px 0 8px;">Local browser-server</h3>
      <dl class="meta">
        <dt>Managed</dt>
        <dd>{{.Snapshot.Browser.Managed.Managed}}</dd>
        <dt>URL</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.URL}}
            <code>{{.Snapshot.Browser.Managed.URL}}</code>
          {{else}}
            -
          {{end}}
        </dd>
        <dt>PID</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.PID}}
            {{.Snapshot.Browser.Managed.PID}}
          {{else}}
            -
          {{end}}
        </dd>
        <dt>Work Dir</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.WorkDir}}
            <code>{{.Snapshot.Browser.Managed.WorkDir}}</code>
          {{else}}
            -
          {{end}}
        </dd>
        <dt>Command</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.Command}}
            <code>{{.Snapshot.Browser.Managed.Command}}</code>
          {{else}}
            -
          {{end}}
        </dd>
        <dt>Log</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.LogURL}}
            <a
              href="{{.Snapshot.Browser.Managed.LogURL}}"
              target="_blank"
              rel="noopener noreferrer"
            >
              open log
            </a>
            <br><code>{{.Snapshot.Browser.Managed.LogPath}}</code>
          {{else if .Snapshot.Browser.Managed.LogPath}}
            <code>{{.Snapshot.Browser.Managed.LogPath}}</code>
          {{else}}
            -
          {{end}}
        </dd>
        <dt>Started</dt>
        <dd>{{formatTime .Snapshot.Browser.Managed.StartedAt}}</dd>
        <dt>Stopped</dt>
        <dd>{{formatTime .Snapshot.Browser.Managed.StoppedAt}}</dd>
        <dt>Exit</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.ExitCode}}
            {{.Snapshot.Browser.Managed.ExitCode}}
          {{else}}
            -
          {{end}}
        </dd>
        <dt>Error</dt>
        <dd>
          {{if .Snapshot.Browser.Managed.LastError}}
            {{.Snapshot.Browser.Managed.LastError}}
          {{else}}
            -
          {{end}}
        </dd>
      </dl>
      {{if .Snapshot.Browser.Managed.RecentLogs}}
      <pre
        style="margin-top: 12px; white-space: pre-wrap;"
      >{{range .Snapshot.Browser.Managed.RecentLogs}}
{{.}}
{{end}}</pre>
      {{end}}
      {{end}}
      {{if .Snapshot.Browser.Providers}}
      <table>
        <thead>
          <tr>
            <th>Provider</th>
            <th>Default</th>
            <th>Host</th>
            <th>Sandbox</th>
            <th>Guards</th>
            <th>Profiles</th>
            <th>Nodes</th>
          </tr>
        </thead>
        <tbody>
          {{range .Snapshot.Browser.Providers}}
          <tr>
            <td>{{if .Name}}{{.Name}}{{else}}browser{{end}}</td>
            <td>
              {{if .DefaultProfile}}
                {{.DefaultProfile}}
              {{else}}
                -
              {{end}}
            </td>
            <td>
              {{if .Host.URL}}
                <code>{{.Host.URL}}</code><br>
                <span class="subtle">{{browserEndpointSummary .Host}}</span>
              {{else}}-{{end}}
            </td>
            <td>
              {{if .Sandbox.URL}}
                <code>{{.Sandbox.URL}}</code><br>
                <span class="subtle">
                  {{browserEndpointSummary .Sandbox}}
                </span>
              {{else}}-{{end}}
            </td>
            <td>
              loopback={{.AllowLoopback}},
              private={{.AllowPrivateNet}},
              file={{.AllowFileURLs}}
            </td>
            <td>
              {{if .Profiles}}
                {{range $i, $profile := .Profiles}}
                  {{if $i}}, {{end}}{{$profile.Name}}
                {{end}}
              {{else}}-{{end}}
            </td>
            <td>
              {{if .Nodes}}
                {{range $i, $node := .Nodes}}
                  {{if $i}}<br>{{end}}{{$node.ID}}
                  {{if $node.Status.URL}}
                    <br>
                    <span class="subtle">
                      {{browserEndpointSummary $node.Status}}
                    </span>
                  {{end}}
                {{end}}
              {{else}}-{{end}}
            </td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <p class="empty">Browser tool is not configured for this runtime.</p>
      {{end}}
    </section>
    {{end}}
  <script>
    const resolveRequestURL = (reference) => {
      const trimmed = typeof reference === "string"
        ? reference.trim()
        : "";
      if (!trimmed) {
        return null;
      }
      try {
        return new URL(trimmed, window.location.href);
      } catch (_) {
        return null;
      }
    };

    (function () {
      const root = document.querySelector("[data-page-stale-root]");
      if (!root) return;

      const stateURL = resolveRequestURL(
        root.getAttribute("data-page-state-path") || ""
      );
      const initialToken =
        root.getAttribute("data-page-state-token") || "";
      const intervalValue = Number(
        root.getAttribute("data-page-refresh-interval") || "0"
      );
      if (!stateURL || !initialToken || intervalValue <= 0) {
        return;
      }

      let currentToken = initialToken;
      let stopped = false;

      const poll = async () => {
        if (stopped || document.hidden) {
          return;
        }
        try {
          const response = await fetch(stateURL.toString(), {
            headers: { Accept: "application/json" },
          });
          if (!response.ok) {
            return;
          }
          const payload = await response.json();
          const nextToken =
            payload && typeof payload.token === "string"
              ? payload.token
              : "";
          if (!nextToken || nextToken === currentToken) {
            return;
          }
          root.hidden = false;
          stopped = true;
        } catch (_) {}
      };

      window.setInterval(poll, intervalValue * 1000);
    })();

    (function () {
      const root = document.querySelector("[data-skills-root]");
      if (!root) return;

      const search = root.querySelector("[data-skills-filter]");
      const shown = root.querySelector("[data-skills-shown]");
      const tabs = Array.from(root.querySelectorAll("[data-skill-tab]"));
      const cards = Array.from(root.querySelectorAll("[data-skill-card]"));
      const groups = Array.from(root.querySelectorAll("[data-skills-group]"));
      const inlineToggleForms = Array.from(root.querySelectorAll("[data-skill-inline-toggle]"));
      const inlineToggleButtons = Array.from(root.querySelectorAll("[data-skill-toggle-button]"));
      const scrollRestoreKey = "openclaw-admin-skills-scroll";
      let active = "all";

      const saveScrollRestore = (form) => {
        if (!form || !window.sessionStorage) return;
        const button = form.querySelector("[data-skill-toggle-button]");
        const card = form.closest("[data-skill-card]");
        const returnField = form.querySelector('input[name="return_to"]');
        const payload = {
          path: window.location.pathname,
          configKey: form.getAttribute("data-skill-config") || "",
          cardId: card ? card.id : "",
          viewportTop: button ? button.getBoundingClientRect().top : 0,
          scrollY: window.scrollY || window.pageYOffset || 0,
          active,
          searchValue: search ? search.value : "",
          ts: Date.now()
        };
        try {
          window.sessionStorage.setItem(
            scrollRestoreKey,
            JSON.stringify(payload)
          );
        } catch (_) {}
        if (returnField) {
          returnField.value = "";
        }
      };

      const restoreScrollPosition = () => {
        if (!window.sessionStorage) return;
        let raw = "";
        try {
          raw = window.sessionStorage.getItem(scrollRestoreKey) || "";
        } catch (_) {
          return;
        }
        if (!raw) return;

        let payload = null;
        try {
          payload = JSON.parse(raw);
        } catch (_) {
          payload = null;
        }
        try {
          window.sessionStorage.removeItem(scrollRestoreKey);
        } catch (_) {}
        if (!payload || payload.path !== window.location.pathname) {
          return;
        }
        if (typeof payload.searchValue === "string" && search) {
          search.value = payload.searchValue;
        }
        if (typeof payload.active === "string" && payload.active) {
          active = payload.active;
        }
        refresh();

        const apply = () => {
          const savedY = Number(payload.scrollY);
          if (Number.isFinite(savedY)) {
            window.scrollTo(0, savedY);
          }

          let target = null;
          if (payload.configKey) {
            target = root.querySelector(
              '[data-skill-switch="' + payload.configKey + '"]'
            );
          }
          if (!target && payload.cardId) {
            target = document.getElementById(payload.cardId);
          }
          const savedTop = Number(payload.viewportTop);
          if (!target || !Number.isFinite(savedTop)) {
            return;
          }
          const rect = target.getBoundingClientRect();
          if (!rect || rect.height === 0) {
            return;
          }
          window.scrollBy(0, rect.top - savedTop);
        };

        window.requestAnimationFrame(() => {
          window.requestAnimationFrame(apply);
        });
      };

      const matches = (card) => {
        const status = card.getAttribute("data-skill-status") || "";
        if (active !== "all" && status !== active) return false;
        const needle = (search && search.value ? search.value : "").trim().toLowerCase();
        if (!needle) return true;
        const haystack = (card.getAttribute("data-skill-search") || "").toLowerCase();
        return haystack.indexOf(needle) >= 0;
      };

      const refresh = () => {
        let visibleCount = 0;
        cards.forEach((card) => {
          const visible = matches(card);
          card.hidden = !visible;
          if (visible) visibleCount += 1;
        });
        groups.forEach((group) => {
          const visibleCards = group.querySelectorAll("[data-skill-card]:not([hidden])");
          group.hidden = visibleCards.length === 0;
        });
        if (shown) shown.textContent = String(visibleCount);
        tabs.forEach((tab) => {
          tab.classList.toggle("active", (tab.getAttribute("data-skill-tab") || "") === active);
        });
      };

      tabs.forEach((tab) => {
        tab.addEventListener("click", () => {
          active = tab.getAttribute("data-skill-tab") || "all";
          refresh();
        });
      });
      if (search) {
        search.addEventListener("input", refresh);
      }
      inlineToggleForms.forEach((form) => {
        form.addEventListener("click", (event) => {
          event.stopPropagation();
        });
        form.addEventListener("keydown", (event) => {
          event.stopPropagation();
        });
        form.addEventListener("submit", () => {
          saveScrollRestore(form);
        });
      });
      inlineToggleButtons.forEach((button) => {
        button.addEventListener("click", (event) => {
          event.preventDefault();
          event.stopPropagation();
          const form = button.form;
          if (!form) return;
          if (typeof form.requestSubmit === "function") {
            form.requestSubmit();
            return;
          }
          form.submit();
        });
      });
      refresh();
      restoreScrollPosition();
    })();

    (function () {
      const revealHashDetails = () => {
        if (!window.location.hash) {
          return;
        }
        const target = document.querySelector(window.location.hash);
        if (!target) {
          return;
        }
        const disclosure = target.matches("details")
          ? target
          : target.closest("details");
        if (!disclosure) {
          return;
        }
        disclosure.open = true;
      };

      window.addEventListener("hashchange", revealHashDetails);
      revealHashDetails();
    })();

    (function () {
      const roots = Array.from(
        document.querySelectorAll("[data-chat-history-root]")
      );
      if (!roots.length) return;

      const itemKindSession = "session";
      const itemKindTurn = "turn";
      const fallbackHistoryError = "Unable to load chat history right now.";
      const fallbackHistoryEmpty =
        "No recent chat transcript is available in this runtime " +
        "for this chat right now.";
      const historyLoadingText = "Loading recent messages...";
      const historyMoreText = "Load older messages";
      const historyBoundedText =
        "This admin view only loads the most recent tracked " +
        "session lines for this chat.";

      const formatDateTime = (value) => {
        if (!value) return "";
        const parsed = new Date(value);
        if (Number.isNaN(parsed.getTime())) return value;
        return parsed.toLocaleString(undefined, {
          year: "numeric",
          month: "2-digit",
          day: "2-digit",
          hour: "2-digit",
          minute: "2-digit",
          second: "2-digit",
          hour12: false,
          timeZoneName: "short",
        });
      };

      const speakerLabel = (item) => {
        if (item && typeof item.speaker === "string" && item.speaker.trim()) {
          return item.speaker.trim();
        }
        const role = item && typeof item.role === "string"
          ? item.role.trim()
          : "";
        switch (role) {
          case "user":
            return "User";
          case "assistant":
            return "Assistant";
          case "system":
            return "System";
          default:
            return "Turn";
        }
      };

      const setText = (node, text) => {
        if (!node) return;
        node.textContent = text;
      };

      const createDiv = (className, text) => {
        const node = document.createElement("div");
        if (className) {
          node.className = className;
        }
        if (typeof text === "string") {
          node.textContent = text;
        }
        return node;
      };

      const updateMeta = (root, page) => {
        const meta = root.querySelector("[data-chat-history-meta]");
        if (!meta || !page) return;
        const loaded = Number(root.getAttribute("data-chat-history-loaded") || "0");
        const total = Number(page.turn_count || 0);
        const sessions = Number(page.session_line_count || 0);
        if (total <= 0) {
          meta.textContent = fallbackHistoryEmpty;
          return;
        }
        const parts = [];
        parts.push(
          "Showing " + String(loaded) + " of " + String(total) +
          " messages"
        );
        if (sessions > 0) {
          parts.push("from " + String(sessions) + " recent session lines");
        }
        meta.textContent = parts.join(" ");
      };

      const updateBoundedNote = (root, bounded) => {
        const note = root.querySelector("[data-chat-history-bounded]");
        if (!note) return;
        note.hidden = !bounded;
        if (bounded) {
          note.textContent = historyBoundedText;
        }
      };

      const updateMoreButton = (root, nextCursor) => {
        const toolbar = root.querySelector("[data-chat-history-toolbar]");
        const button = root.querySelector("[data-chat-history-more]");
        if (!toolbar || !button) return;
        const hasMore = typeof nextCursor === "string" && nextCursor !== "";
        toolbar.hidden = !hasMore;
        button.hidden = !hasMore;
        button.disabled = false;
        button.textContent = historyMoreText;
        if (hasMore) {
          button.setAttribute("data-chat-history-next", nextCursor);
        } else {
          button.removeAttribute("data-chat-history-next");
        }
      };

      const clearMessages = (root) => {
        const error = root.querySelector("[data-chat-history-error]");
        const empty = root.querySelector("[data-chat-history-empty]");
        if (error) {
          error.hidden = true;
          error.textContent = "";
        }
        if (empty) {
          empty.hidden = true;
          empty.textContent = "";
        }
      };

      const showError = (root, message) => {
        const error = root.querySelector("[data-chat-history-error]");
        const empty = root.querySelector("[data-chat-history-empty]");
        if (empty) {
          empty.hidden = true;
          empty.textContent = "";
        }
        if (!error) return;
        error.hidden = false;
        error.textContent = message || fallbackHistoryError;
      };

      const showEmpty = (root, message) => {
        const empty = root.querySelector("[data-chat-history-empty]");
        const error = root.querySelector("[data-chat-history-error]");
        if (error) {
          error.hidden = true;
          error.textContent = "";
        }
        if (!empty) return;
        empty.hidden = false;
        empty.textContent = message || fallbackHistoryEmpty;
      };

      const renderSessionItem = (item) => {
        const wrapper = createDiv("chat-timeline-session");
        const head = createDiv("chat-timeline-session-head");
        const copy = document.createElement("div");
        const title = createDiv(
          "chat-timeline-session-label",
          item.session_label || "Recent session"
        );
        const meta = createDiv("subtle chat-timeline-session-meta");
        const code = document.createElement("code");
        code.textContent = item.session_id || "";
        meta.appendChild(code);
        copy.appendChild(title);
        copy.appendChild(meta);
        head.appendChild(copy);
        if (item.last_activity) {
          head.appendChild(
            createDiv("subtle", formatDateTime(item.last_activity))
          );
        }
        wrapper.appendChild(head);
        return wrapper;
      };

      const renderTurnItem = (item) => {
        const article = document.createElement("article");
        article.className = "chat-turn";
        const head = createDiv("chat-turn-head");
        head.appendChild(
          createDiv("chat-turn-speaker", speakerLabel(item))
        );
        if (item.timestamp) {
          head.appendChild(
            createDiv("subtle", formatDateTime(item.timestamp))
          );
        }
        article.appendChild(head);
        if (typeof item.quote_text === "string" && item.quote_text.trim()) {
          article.appendChild(
            createDiv("chat-turn-quote", item.quote_text)
          );
        }
        article.appendChild(
          createDiv("chat-turn-text", item.text || "")
        );
        return article;
      };

      const renderHistoryItem = (item) => {
        if (!item || typeof item.kind !== "string") {
          return null;
        }
        if (item.kind === itemKindSession) {
          return renderSessionItem(item);
        }
        if (item.kind === itemKindTurn) {
          return renderTurnItem(item);
        }
        return null;
      };

      const updateLoading = (root, loading) => {
        root.setAttribute(
          "data-chat-history-loading",
          loading ? "true" : "false"
        );
        const button = root.querySelector("[data-chat-history-more]");
        if (!button) return;
        button.disabled = loading;
        if (loading) {
          button.textContent = historyLoadingText;
        } else if (button.hidden !== true) {
          button.textContent = historyMoreText;
        }
      };

      const fetchHistory = async (root, cursor) => {
        const chatID = root.getAttribute("data-chat-id") || "";
        const url = resolveRequestURL(
          root.getAttribute("data-chat-history-path") || ""
        );
        if (!url) {
          throw new Error(fallbackHistoryError);
        }
        url.searchParams.set("chat_id", chatID);
        if (cursor) {
          url.searchParams.set("cursor", cursor);
        }
        const response = await fetch(url.toString(), {
          headers: { Accept: "application/json" },
        });
        if (!response.ok) {
          const text = (await response.text()).trim();
          throw new Error(text || fallbackHistoryError);
        }
        return response.json();
      };

      const loadHistory = async (root, cursor) => {
        const items = root.querySelector("[data-chat-history-items]");
        if (!items) return;
        if (root.getAttribute("data-chat-history-loading") === "true") {
          return;
        }
        const prepend = typeof cursor === "string" && cursor !== "";
        const anchor = prepend ? items.firstElementChild : null;
        const anchorTop = anchor ? anchor.getBoundingClientRect().top : 0;
        updateLoading(root, true);
        clearMessages(root);
        try {
          const page = await fetchHistory(root, cursor);
          const fragment = document.createDocumentFragment();
          const pageItems = Array.isArray(page.items) ? page.items : [];
          pageItems.forEach((item) => {
            const node = renderHistoryItem(item);
            if (node) {
              fragment.appendChild(node);
            }
          });
          if (!prepend) {
            items.innerHTML = "";
          }
          if (prepend) {
            items.prepend(fragment);
          } else {
            items.appendChild(fragment);
          }

          const loaded = Number(
            root.getAttribute("data-chat-history-loaded") || "0"
          ) + Number(page.returned_turn_count || 0);
          root.setAttribute(
            "data-chat-history-loaded",
            String(loaded)
          );
          updateMeta(root, page);
          updateBoundedNote(root, Boolean(page.bounded));
          updateMoreButton(root, page.next_cursor || "");

          if (!pageItems.length && Number(page.turn_count || 0) === 0) {
            showEmpty(root, fallbackHistoryEmpty);
          }
          if (prepend && anchor) {
            window.scrollBy(
              0,
              anchor.getBoundingClientRect().top - anchorTop
            );
          }
          root.setAttribute("data-chat-history-loaded-once", "true");
        } catch (err) {
          showError(
            root,
            err && typeof err.message === "string"
              ? err.message
              : fallbackHistoryError
          );
        } finally {
          updateLoading(root, false);
        }
      };

      roots.forEach((root) => {
        const disclosure = root.closest("details");
        const button = root.querySelector("[data-chat-history-more]");
        if (button) {
          button.addEventListener("click", () => {
            const nextCursor =
              button.getAttribute("data-chat-history-next") || "";
            if (!nextCursor) return;
            loadHistory(root, nextCursor);
          });
        }
        const ensureLoaded = () => {
          if (
            root.getAttribute("data-chat-history-loaded-once") ===
            "true"
          ) {
            return;
          }
          loadHistory(root, "");
        };
        if (disclosure) {
          disclosure.addEventListener("toggle", () => {
            if (disclosure.open) {
              ensureLoaded();
            }
          });
          if (disclosure.open) {
            ensureLoaded();
          }
          return;
        }
        ensureLoaded();
      });
    })();

    (function () {
      const root = document.querySelector("[data-memory-root]");
      if (!root) return;

      const search = root.querySelector("[data-memory-search]");
      const shown = root.querySelector("[data-memory-shown]");
      const empty = document.querySelector("[data-memory-empty]");
      const cards = Array.from(
        document.querySelectorAll("[data-memory-card]")
      );

      const readControls = (card) => ({
        editor: card.querySelector("[data-memory-editor]"),
        save: card.querySelector("[data-memory-save]"),
        reset: card.querySelector("[data-memory-reset]"),
        error: card.querySelector("[data-memory-error]"),
        status: card.querySelector("[data-memory-status]"),
      });

      const setStatus = (card, message, isError) => {
        const status = readControls(card).status;
        if (!status) return;
        status.textContent = message;
        status.classList.toggle("error", Boolean(isError));
      };

      const clearError = (card) => {
        const error = readControls(card).error;
        if (!error) return;
        error.hidden = true;
        error.textContent = "";
      };

      const showError = (card, message) => {
        const error = readControls(card).error;
        if (!error) return;
        error.hidden = false;
        error.textContent = message;
        setStatus(card, message, true);
      };

      const updateDirtyState = (card) => {
        const controls = readControls(card);
        const editor = controls.editor;
        if (!editor) return;
        const original = editor.dataset.original || "";
        const loaded = card.dataset.memoryLoaded === "true";
        const dirty = !editor.disabled && editor.value !== original;
        if (controls.save) {
          controls.save.disabled = editor.disabled || !dirty;
        }
        if (controls.reset) {
          controls.reset.disabled = editor.disabled || !dirty;
        }
        if (editor.disabled) {
          setStatus(
            card,
            "Open this scope to load its full file.",
            false
          );
          return;
        }
        if (dirty) {
          setStatus(card, "Unsaved changes.", false);
          return;
        }
        if (loaded) {
          setStatus(
            card,
            "Loaded from disk. Edit the file to enable save.",
            false
          );
        }
      };

      const loadCard = async (card) => {
        if (
          card.dataset.memoryLoaded === "true" ||
          card.dataset.memoryLoading === "true"
        ) {
          updateDirtyState(card);
          return;
        }
        const controls = readControls(card);
        const editor = controls.editor;
        card.dataset.memoryLoading = "true";
        clearError(card);
        setStatus(card, "Loading full file...", false);
        if (editor) {
          editor.disabled = true;
        }
        if (controls.save) {
          controls.save.disabled = true;
        }
        if (controls.reset) {
          controls.reset.disabled = true;
        }
        try {
          const url = resolveRequestURL(
            card.getAttribute("data-memory-load-url") || ""
          );
          if (!url) {
            throw new Error("memory file endpoint is unavailable");
          }
          const response = await fetch(url.toString(), {
            headers: { Accept: "application/json" },
          });
          if (!response.ok) {
            const text = (await response.text()).trim();
            throw new Error(text || "failed to load memory file");
          }
          const payload = await response.json();
          const content =
            typeof payload.content === "string" ? payload.content : "";
          if (editor) {
            editor.value = content;
            editor.dataset.original = content;
            editor.disabled = false;
          }
          card.dataset.memoryLoaded = "true";
          updateDirtyState(card);
        } catch (err) {
          showError(
            card,
            err && typeof err.message === "string"
              ? err.message
              : "failed to load memory file"
          );
        } finally {
          delete card.dataset.memoryLoading;
        }
      };

      const matches = (card) => {
        const needle =
          (search && search.value ? search.value : "")
            .trim()
            .toLowerCase();
        if (!needle) return true;
        const haystack = (
          card.getAttribute("data-memory-search") || ""
        ).toLowerCase();
        if (haystack.indexOf(needle) === -1) return false;
        return true;
      };

      const refresh = () => {
        let visibleCount = 0;
        cards.forEach((card) => {
          const visible = matches(card);
          card.hidden = !visible;
          if (visible) visibleCount += 1;
        });
        if (shown) shown.textContent = String(visibleCount);
        if (empty) empty.hidden = visibleCount !== 0;
      };

      cards.forEach((card) => {
        const disclosure = card.matches("details")
          ? card
          : card.closest("details");
        const controls = readControls(card);
        if (controls.editor) {
          controls.editor.addEventListener("input", () => {
            updateDirtyState(card);
          });
        }
        if (controls.reset) {
          controls.reset.addEventListener("click", () => {
            if (!controls.editor) return;
            controls.editor.value =
              controls.editor.dataset.original || "";
            updateDirtyState(card);
          });
        }
        if (disclosure) {
          disclosure.addEventListener("toggle", () => {
            if (disclosure.open) {
              loadCard(card);
            }
          });
          if (disclosure.open) {
            loadCard(card);
          }
        }
      });

      if (search) {
        search.addEventListener("input", refresh);
      }
      refresh();
    })();
  </script>
      </div>
    </main>
  </div>
  <script>
    (function() {
      const pendingScrollKey = "openclaw.admin.pendingScroll";
      const pendingScrollMaxAgeMS = 30000;
      const sidebar = document.querySelector(".sidebar");
      const viewportPadding = 16;

      function readPendingScroll() {
        try {
          const raw = window.sessionStorage.getItem(pendingScrollKey);
          if (!raw) return null;
          window.sessionStorage.removeItem(pendingScrollKey);
          const value = JSON.parse(raw);
          if (!value || typeof value !== "object") return null;
          if (Date.now() - value.savedAt > pendingScrollMaxAgeMS) {
            return null;
          }
          if (
            typeof value.targetPath === "string" &&
            value.targetPath !== window.location.pathname
          ) {
            return null;
          }
          return value;
        } catch (err) {
          return null;
        }
      }

      function savePendingScroll(targetPath) {
        try {
          window.sessionStorage.setItem(pendingScrollKey, JSON.stringify({
            savedAt: Date.now(),
            targetPath: targetPath,
            sidebarTop: sidebar ? sidebar.scrollTop : 0
          }));
        } catch (err) {}
      }

      function revealActiveLink() {
        const activeLink = sidebar &&
          sidebar.querySelector(".sidebar-link.active");
        if (!activeLink) {
          return;
        }
        const activeRect = activeLink.getBoundingClientRect();
        if (sidebar.scrollHeight <= sidebar.clientHeight) {
          const topEdge = viewportPadding;
          const bottomEdge = window.innerHeight - viewportPadding;
          if (activeRect.top < topEdge) {
            window.scrollBy(0, activeRect.top - topEdge);
            return;
          }
          if (activeRect.bottom > bottomEdge) {
            window.scrollBy(0, activeRect.bottom - bottomEdge);
          }
          return;
        }
        const sidebarRect = sidebar.getBoundingClientRect();
        const topEdge = sidebarRect.top + viewportPadding;
        const bottomEdge = sidebarRect.bottom - viewportPadding;
        if (activeRect.top < topEdge) {
          sidebar.scrollTop -= topEdge - activeRect.top;
          return;
        }
        if (activeRect.bottom > bottomEdge) {
          sidebar.scrollTop += activeRect.bottom - bottomEdge;
        }
      }

      document.querySelectorAll(".sidebar-link").forEach(function(link) {
        link.addEventListener("click", function(evt) {
          if (
            evt.defaultPrevented ||
            evt.button !== 0 ||
            evt.metaKey ||
            evt.ctrlKey ||
            evt.shiftKey ||
            evt.altKey
          ) {
            return;
          }
          const targetURL = new URL(link.href, window.location.href);
          savePendingScroll(targetURL.pathname);
        });
      });

      const pendingScroll = readPendingScroll();
      if (pendingScroll) {
        if (sidebar && Number.isFinite(pendingScroll.sidebarTop)) {
          sidebar.scrollTop = pendingScroll.sidebarTop;
        }
      } else {
        revealActiveLink();
      }
    })();
  </script>
</body>
</html>`
