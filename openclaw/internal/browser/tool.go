//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package browser

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	actionStatus     = "status"
	actionStart      = "start"
	actionStop       = "stop"
	actionProfiles   = "profiles"
	actionTabs       = "tabs"
	actionOpen       = "open"
	actionFocus      = "focus"
	actionClose      = "close"
	actionSnapshot   = "snapshot"
	actionScreenshot = "screenshot"
	actionNavigate   = "navigate"
	actionConsole    = "console"
	actionCookies    = "cookies"
	actionStorage    = "storage"
	actionPDF        = "pdf"
	actionDownload   = "download"
	actionWaitDL     = "waitDownload"
	actionUpload     = "upload"
	actionDialog     = "dialog"
	actionOffline    = "offline"
	actionHeaders    = "headers"
	actionCreds      = "credentials"
	actionGeo        = "geolocation"
	actionMedia      = "media"
	actionTimezone   = "timezone"
	actionLocale     = "locale"
	actionDevice     = "device"
	actionAct        = "act"
)

const (
	actClick          = "click"
	actType           = "type"
	actPress          = "press"
	actHover          = "hover"
	actScrollIntoView = "scrollIntoView"
	actDrag           = "drag"
	actSelect         = "select"
	actFill           = "fill"
	actResize         = "resize"
	actWait           = "wait"
	actEvaluate       = "evaluate"
	actClose          = "close"
)

const (
	targetHost    = "host"
	targetSandbox = "sandbox"
	targetNode    = "node"
)

const (
	tabActionList   = "list"
	tabActionNew    = "create"
	tabActionSelect = "select"
	tabActionClose  = "close"
)

const (
	stateOpGet   = "get"
	stateOpSet   = "set"
	stateOpClear = "clear"
)

var supportedActions = []string{
	actionStatus,
	actionStart,
	actionStop,
	actionProfiles,
	actionTabs,
	actionOpen,
	actionFocus,
	actionClose,
	actionSnapshot,
	actionScreenshot,
	actionNavigate,
	actionConsole,
	actionCookies,
	actionStorage,
	actionPDF,
	actionDownload,
	actionWaitDL,
	actionUpload,
	actionDialog,
	actionOffline,
	actionHeaders,
	actionCreds,
	actionGeo,
	actionMedia,
	actionTimezone,
	actionLocale,
	actionDevice,
	actionAct,
}

type actRequest struct {
	Kind        string           `json:"kind,omitempty"`
	TargetID    string           `json:"targetId,omitempty"`
	Ref         string           `json:"ref,omitempty"`
	DoubleClick *bool            `json:"doubleClick,omitempty"`
	Button      string           `json:"button,omitempty"`
	Modifiers   []string         `json:"modifiers,omitempty"`
	Text        string           `json:"text,omitempty"`
	Submit      *bool            `json:"submit,omitempty"`
	Slowly      *bool            `json:"slowly,omitempty"`
	Key         string           `json:"key,omitempty"`
	DelayMs     *int             `json:"delayMs,omitempty"`
	StartRef    string           `json:"startRef,omitempty"`
	EndRef      string           `json:"endRef,omitempty"`
	Values      []string         `json:"values,omitempty"`
	Fields      []map[string]any `json:"fields,omitempty"`
	Width       *int             `json:"width,omitempty"`
	Height      *int             `json:"height,omitempty"`
	TimeMs      *int             `json:"timeMs,omitempty"`
	Selector    string           `json:"selector,omitempty"`
	URL         string           `json:"url,omitempty"`
	LoadState   string           `json:"loadState,omitempty"`
	TextGone    string           `json:"textGone,omitempty"`
	TimeoutMs   *int             `json:"timeoutMs,omitempty"`
	Fn          string           `json:"fn,omitempty"`
}

type input struct {
	Action         string            `json:"action"`
	Target         string            `json:"target,omitempty"`
	Node           string            `json:"node,omitempty"`
	Profile        string            `json:"profile,omitempty"`
	TargetURL      string            `json:"targetUrl,omitempty"`
	URL            string            `json:"url,omitempty"`
	TargetID       string            `json:"targetId,omitempty"`
	Operation      string            `json:"operation,omitempty"`
	Store          string            `json:"store,omitempty"`
	Value          string            `json:"value,omitempty"`
	Limit          *int              `json:"limit,omitempty"`
	MaxChars       *int              `json:"maxChars,omitempty"`
	Mode           string            `json:"mode,omitempty"`
	SnapshotFormat string            `json:"snapshotFormat,omitempty"`
	Refs           string            `json:"refs,omitempty"`
	Interactive    *bool             `json:"interactive,omitempty"`
	Compact        *bool             `json:"compact,omitempty"`
	Depth          *int              `json:"depth,omitempty"`
	Selector       string            `json:"selector,omitempty"`
	Frame          string            `json:"frame,omitempty"`
	Labels         *bool             `json:"labels,omitempty"`
	FullPage       *bool             `json:"fullPage,omitempty"`
	Ref            string            `json:"ref,omitempty"`
	Element        string            `json:"element,omitempty"`
	Type           string            `json:"type,omitempty"`
	Level          string            `json:"level,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Offline        *bool             `json:"offline,omitempty"`
	Username       string            `json:"username,omitempty"`
	Password       string            `json:"password,omitempty"`
	Cookie         map[string]any    `json:"cookie,omitempty"`
	Path           string            `json:"path,omitempty"`
	Paths          []string          `json:"paths,omitempty"`
	InputRef       string            `json:"inputRef,omitempty"`
	Filename       string            `json:"filename,omitempty"`
	TimeoutMs      *int              `json:"timeoutMs,omitempty"`
	Clear          *bool             `json:"clear,omitempty"`
	Accept         *bool             `json:"accept,omitempty"`
	PromptText     string            `json:"promptText,omitempty"`
	Latitude       *float64          `json:"latitude,omitempty"`
	Longitude      *float64          `json:"longitude,omitempty"`
	Accuracy       *float64          `json:"accuracy,omitempty"`
	Origin         string            `json:"origin,omitempty"`
	ColorScheme    string            `json:"colorScheme,omitempty"`
	TimezoneID     string            `json:"timezoneId,omitempty"`
	Locale         string            `json:"locale,omitempty"`
	Name           string            `json:"name,omitempty"`
	Request        *actRequest       `json:"request,omitempty"`
	Kind           string            `json:"kind,omitempty"`
	DoubleClick    *bool             `json:"doubleClick,omitempty"`
	Button         string            `json:"button,omitempty"`
	Modifiers      []string          `json:"modifiers,omitempty"`
	Text           string            `json:"text,omitempty"`
	Submit         *bool             `json:"submit,omitempty"`
	Slowly         *bool             `json:"slowly,omitempty"`
	Key            string            `json:"key,omitempty"`
	DelayMs        *int              `json:"delayMs,omitempty"`
	StartRef       string            `json:"startRef,omitempty"`
	EndRef         string            `json:"endRef,omitempty"`
	Values         []string          `json:"values,omitempty"`
	Fields         []map[string]any  `json:"fields,omitempty"`
	Width          *int              `json:"width,omitempty"`
	Height         *int              `json:"height,omitempty"`
	TimeMs         *int              `json:"timeMs,omitempty"`
	TextGone       string            `json:"textGone,omitempty"`
	LoadState      string            `json:"loadState,omitempty"`
	Fn             string            `json:"fn,omitempty"`
}

// Tool implements the first-class browser tool contract.
type Tool struct {
	defaultProfile  string
	evaluateEnabled bool
	navigation      navigationPolicy
	hostServer      *serverTargetConfig
	sandboxServer   *serverTargetConfig
	nodeTargets     map[string]serverTargetConfig
	profiles        map[string]ProfileConfig
	drivers         map[string]driver
	serverDriversMu sync.RWMutex
	serverDrivers   map[string]driver
}

// NewTool creates a native browser tool backed by MCP or browser-server.
func NewTool(cfg Config) (*Tool, error) { _ = "STUB: not implemented"; return nil, nil }

func newToolWithDrivers(
	defaultProfile string,
	evaluateEnabled bool,
	navigation navigationPolicy,
	hostServer *serverTargetConfig,
	sandboxServer *serverTargetConfig,
	nodeTargets map[string]serverTargetConfig,
	profiles map[string]ProfileConfig,
	drivers map[string]driver,
) *Tool {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *Tool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func browserSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func stringSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func numberSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func boolSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func stringArraySchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func validateTargetSelection(in input) error { _ = "STUB: not implemented"; return nil }

func (t *Tool) resolveDriver(
	in input,
) (string, driver, error) {
	_ = "STUB: not implemented"
	return "", *new(driver), nil
}

func (t *Tool) serverDriverForTarget(
	target *serverTargetConfig,
	profile string,
) (driver, bool) {
	_ = "STUB: not implemented"
	return *new(driver), false
}

func (t *Tool) handleProfiles(
	ctx context.Context,
) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

func (t *Tool) statusDriver(
	name string,
	cfg ProfileConfig,
) driver {
	_ = "STUB: not implemented"
	return *new(driver)
}

func (t *Tool) driverTypeForProfile(
	profile string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Tool) driverTypeForInput(
	profile string,
	in input,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Tool) handleStatus(
	ctx context.Context,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleStart(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleStop(
	profile string,
	driverType string,
	drv driver,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleTabs(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleOpen(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleFocus(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleClose(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleSnapshot(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleScreenshot(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleNavigate(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleConsole(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleCookies(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleStorage(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleOffline(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleHeaders(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleCredentials(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleGeolocation(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleMedia(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleTimezone(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleLocale(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleDevice(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handlePDF(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleDownload(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleWaitDownload(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleUpload(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleDialog(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (t *Tool) handleAct(
	ctx context.Context,
	profile string,
	driverType string,
	drv driver,
	in input,
) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func normalizeActRequest(in input) actRequest { _ = "STUB: not implemented"; return *new(actRequest) }

func (t *Tool) executeAct(
	ctx context.Context,
	drv driver,
	req actRequest,
	driverType string,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) executeFill(
	ctx context.Context,
	drv driver,
	req actRequest,
	driverType string,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) executeWait(
	ctx context.Context,
	drv driver,
	req actRequest,
	driverType string,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) textResult(
	action string,
	profile string,
	driverType string,
	maxChars *int,
	raw any,
) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

func (t *Tool) tabsResult(
	profile string,
	driverType string,
	raw any,
	limit *int,
) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

func selectTarget(
	ctx context.Context,
	drv driver,
	targetID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func boolValue(v *bool) bool { _ = "STUB: not implemented"; return false }

func intValue(v *int) int { _ = "STUB: not implemented"; return 0 }

func stateOperation(raw string) string { _ = "STUB: not implemented"; return "" }

func storageScope(raw string) string { _ = "STUB: not implemented"; return "" }

func addServerTimeoutArg(
	args map[string]any,
	driverType string,
	timeout *int,
) {
	_ = "STUB: not implemented"
	return
}

func hasServerSnapshotArgs(in input) bool { _ = "STUB: not implemented"; return false }

func addServerSnapshotArgs(args map[string]any, in input) { _ = "STUB: not implemented"; return }

func downloadOutputPath(in input) string { _ = "STUB: not implemented"; return "" }

func describeElement(ref string, fallback string) string { _ = "STUB: not implemented"; return "" }
