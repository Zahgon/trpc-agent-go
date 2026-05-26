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
	"net/http"
	"time"
)

const (
	routeRuntimeControlPage   = "/runtime-control"
	routeRuntimeControlAction = "/api/runtime/control/action"

	queryRuntimeVersion = "version"

	formRuntimeActionKind    = "kind"
	formRuntimeActionMode    = "mode"
	formRuntimeTargetVersion = "target_version"

	runtimeActionRestart = "restart"
	runtimeActionUpgrade = "upgrade"

	runtimeModeGraceful = "graceful"
	runtimeModeForce    = "force"
)

const pageSummaryRuntimeControl = "" +
	"Request graceful or forced restarts, switch versions, " +
	"and inspect release notes for the current runtime."

const (
	viewRuntimeControl adminView = "runtime_control"
)

type RuntimeLifecycleProvider interface {
	RuntimeLifecycleStatus() (RuntimeLifecycleStatus, error)
	RuntimeLifecycleVersions() (RuntimeLifecycleVersionIndex, error)
	RuntimeLifecycleChangelog(string) (
		RuntimeLifecycleChangelog,
		error,
	)
	RequestRuntimeLifecycleAction(
		RuntimeLifecycleActionRequest,
	) (RuntimeLifecycleActionResult, error)
}

type RuntimeLifecycleStatus struct {
	State           string                         `json:"state,omitempty"`
	CurrentVersion  string                         `json:"current_version,omitempty"`
	ActiveRequests  int                            `json:"active_requests"`
	RunningRequests int                            `json:"running_requests"`
	QueuedRequests  int                            `json:"queued_requests"`
	Pending         *RuntimeLifecyclePendingAction `json:"pending,omitempty"`
	UpdatedAt       time.Time                      `json:"updated_at,omitempty"`
	ExitCode        int                            `json:"exit_code"`
}

type RuntimeLifecyclePendingAction struct {
	ID             string    `json:"id,omitempty"`
	Kind           string    `json:"kind,omitempty"`
	Mode           string    `json:"mode,omitempty"`
	TargetVersion  string    `json:"target_version,omitempty"`
	Actor          string    `json:"actor,omitempty"`
	Source         string    `json:"source,omitempty"`
	RequestedAt    time.Time `json:"requested_at,omitempty"`
	CurrentVersion string    `json:"current_version,omitempty"`
	Summary        []string  `json:"summary,omitempty"`
}

type RuntimeLifecycleVersionIndex struct {
	LatestVersion      string                    `json:"latest_version,omitempty"`
	MinSupportedTarget string                    `json:"min_supported_target,omitempty"`
	Versions           []RuntimeLifecycleVersion `json:"versions,omitempty"`
}

type RuntimeLifecycleVersion struct {
	Version      string    `json:"version,omitempty"`
	PublishedAt  time.Time `json:"published_at,omitempty"`
	InstallURL   string    `json:"install_url,omitempty"`
	ChangelogURL string    `json:"changelog_url,omitempty"`
	Notes        []string  `json:"notes,omitempty"`
}

type RuntimeLifecycleChangelog struct {
	Version   string   `json:"version,omitempty"`
	Summary   []string `json:"summary,omitempty"`
	Changelog string   `json:"changelog,omitempty"`
}

type RuntimeLifecycleActionRequest struct {
	Kind          string `json:"kind,omitempty"`
	Mode          string `json:"mode,omitempty"`
	TargetVersion string `json:"target_version,omitempty"`
}

type RuntimeLifecycleActionResult struct {
	Status  RuntimeLifecycleStatus `json:"status"`
	Started bool                   `json:"started"`
}

type RuntimeLifecyclePageStatus struct {
	Enabled         bool                         `json:"enabled"`
	Error           string                       `json:"error,omitempty"`
	Status          RuntimeLifecycleStatus       `json:"status"`
	Index           RuntimeLifecycleVersionIndex `json:"index"`
	SelectedVersion string                       `json:"selected_version,omitempty"`
	Changelog       RuntimeLifecycleChangelog    `json:"changelog"`
}

func (s *Service) runtimeLifecycleProvider() RuntimeLifecycleProvider {
	_ = "STUB: not implemented"
	return *new(RuntimeLifecycleProvider)
}

func (s *Service) hasRuntimeLifecycleProvider() bool { _ = "STUB: not implemented"; return false }

func (s *Service) runtimeLifecycleStatus(
	r *http.Request,
) RuntimeLifecyclePageStatus {
	_ = "STUB: not implemented"
	return *new(RuntimeLifecyclePageStatus)
}

func (s *Service) runtimeLifecycleRefreshStatus(
	r *http.Request,
) RuntimeLifecyclePageStatus {
	_ = "STUB: not implemented"
	return *new(RuntimeLifecyclePageStatus)
}

func resolveRuntimeLifecycleSelectedVersion(
	r *http.Request,
	state RuntimeLifecyclePageStatus,
) string {
	_ = "STUB: not implemented"
	return ""
}

func resolveRuntimeLifecycleRefreshVersion(
	r *http.Request,
	status RuntimeLifecycleStatus,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *Service) handleRuntimeControlPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleRuntimeControlAction(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) redirectRuntimeControlWithMessage(
	w http.ResponseWriter,
	r *http.Request,
	key string,
	message string,
	version string,
	fragment string,
) {
	_ = "STUB: not implemented"
	return
}

func runtimeLifecycleActionRequestFromForm(
	r *http.Request,
) (RuntimeLifecycleActionRequest, error) {
	_ = "STUB: not implemented"
	return *new(RuntimeLifecycleActionRequest), nil
}

func errRuntimeActionRequired(message string) error { _ = "STUB: not implemented"; return nil }

type runtimeLifecycleFormError struct {
	Message string
}

func (e *runtimeLifecycleFormError) Error() string { _ = "STUB: not implemented"; return "" }

func normalizeRuntimeLifecycleAction(raw string) string { _ = "STUB: not implemented"; return "" }

func normalizeRuntimeLifecycleMode(raw string) string { _ = "STUB: not implemented"; return "" }

func runtimeLifecycleActionNotice(
	req RuntimeLifecycleActionRequest,
) string {
	_ = "STUB: not implemented"
	return ""
}
