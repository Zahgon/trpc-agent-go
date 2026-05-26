//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package cron

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/outbound"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/runtimeprofile"
)

const (
	ScheduleKindAt    = "at"
	ScheduleKindAfter = "after"
	ScheduleKindEvery = "every"
	ScheduleKindCron  = "cron"
)

const (
	OverlapPolicySkip    = "skip"
	OverlapPolicyReplace = "replace"
)

const (
	runtimeStateScheduledRun = "openclaw.cron.scheduled_run"
	runtimeStateJobID        = "openclaw.cron.job_id"
	runtimeStateRunIndex     = "openclaw.cron.run_index"
	runtimeStateHasMaxRuns   = "openclaw.cron.has_max_runs"
	runtimeStateMaxRuns      = "openclaw.cron.max_runs"
	runtimeStateRemaining    = "openclaw.cron.remaining_runs"
	runtimeStateIsFinalRun   = "openclaw.cron.is_final_run"
)

const cronDeliverySkipMessageToolTarget = "duplicate_message_tool_text"

const cronSessionPrefix = "cron:"

const (
	StatusIdle           = "idle"
	StatusRunning        = "running"
	StatusSucceeded      = "succeeded"
	StatusFailed         = "failed"
	StatusDeliveryFailed = "delivery_failed"
)

const (
	defaultTickInterval = time.Second

	defaultCronDir  = "cron"
	defaultJobsFile = "jobs.json"

	maxStoredOutputRunes = 2_000
)

// Schedule describes when a cron job should run.
type Schedule struct {
	Kind     string `json:"kind"`
	At       string `json:"at,omitempty"`
	Every    string `json:"every,omitempty"`
	EveryMS  int64  `json:"every_ms,omitempty"`
	CronExpr string `json:"cron_expr,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

// ExecutionPolicy describes run bounds and overlap handling.
type ExecutionPolicy struct {
	MaxRuns       int        `json:"max_runs,omitempty"`
	EndsAt        *time.Time `json:"ends_at,omitempty"`
	OverlapPolicy string     `json:"overlap_policy,omitempty"`
}

// ExecutionStats tracks scheduler-visible run counters.
type ExecutionStats struct {
	RunCount             int `json:"run_count,omitempty"`
	SuccessCount         int `json:"success_count,omitempty"`
	FailureCount         int `json:"failure_count,omitempty"`
	DeliveryFailureCount int `json:"delivery_failure_count,omitempty"`
}

// RuntimeProfileRef stores the profile identity captured when a job is
// created. Scheduled jobs persist selector metadata, not secrets or full
// policy blobs, so profile IDs must remain resolvable by the configured
// resolver when the job runs later.
type RuntimeProfileRef struct {
	ID        string `json:"id,omitempty"`
	Version   string `json:"version,omitempty"`
	AppName   string `json:"app_name,omitempty"`
	Channel   string `json:"channel,omitempty"`
	TenantID  string `json:"tenant_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

type scheduledRunTemplateData struct {
	Cron cronRunTemplateData
}

type cronRunTemplateData struct {
	RunIndex      int
	HasMaxRuns    bool
	MaxRuns       int
	RemainingRuns int
	IsFinalRun    bool
}

// Job is one persisted scheduled agent turn.
type Job struct {
	ID         string                  `json:"id"`
	Name       string                  `json:"name,omitempty"`
	Enabled    bool                    `json:"enabled"`
	Schedule   Schedule                `json:"schedule"`
	Policy     ExecutionPolicy         `json:"policy,omitempty"`
	Stats      ExecutionStats          `json:"stats,omitempty"`
	Message    string                  `json:"message"`
	UserID     string                  `json:"user_id"`
	TimeoutSec int                     `json:"timeout_sec,omitempty"`
	Delivery   outbound.DeliveryTarget `json:"delivery,omitempty"`
	Profile    *RuntimeProfileRef      `json:"profile,omitempty"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastRunAt *time.Time `json:"last_run_at,omitempty"`
	NextRunAt *time.Time `json:"next_run_at,omitempty"`

	LastStatus string `json:"last_status,omitempty"`
	LastError  string `json:"last_error,omitempty"`
	LastOutput string `json:"last_output,omitempty"`
}

func runtimeProfileRefFromProfile(
	profile runtimeprofile.Profile,
) RuntimeProfileRef {
	_ = "STUB: not implemented"
	return *new(RuntimeProfileRef)
}

func (r *RuntimeProfileRef) profile() runtimeprofile.Profile {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Profile)
}

func (r *RuntimeProfileRef) hasProfile() bool { _ = "STUB: not implemented"; return false }

func (j *Job) clone() *Job { _ = "STUB: not implemented"; return nil }

func sanitizeStoredOutput(text string) string { _ = "STUB: not implemented"; return "" }

// IsRunSessionID reports whether a session id belongs to cron execution.
func IsRunSessionID(sessionID string) bool { _ = "STUB: not implemented"; return false }

func normalizeScheduleKind(kind string) string { _ = "STUB: not implemented"; return "" }

// ScheduleSummary returns a stable human-readable schedule summary.
func ScheduleSummary(schedule Schedule) string { _ = "STUB: not implemented"; return "" }

func freshRunSessionID(jobID string, now time.Time) string { _ = "STUB: not implemented"; return "" }

func freshRequestID(jobID string, now time.Time) string { _ = "STUB: not implemented"; return "" }

func normalizeOverlapPolicy(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func cloneTimePtr(src *time.Time) *time.Time { _ = "STUB: not implemented"; return nil }

func scheduledRunContext(job *Job) cronRunTemplateData {
	_ = "STUB: not implemented"
	return *new(cronRunTemplateData)
}
