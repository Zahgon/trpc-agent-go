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
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/outbound"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolCron = "cron"

	actionStatus = "status"
	actionList   = "list"
	actionAdd    = "add"
	actionUpdate = "update"
	actionRemove = "remove"
	actionDelete = "delete"
	actionRun    = "run"
	actionClear  = "clear"
)

const (
	errJobIDRequired     = "job_id is required"
	errProfileIDRequired = "cron: runtime profile id is required"
	errAfterWithSchedule = "cron: after cannot be combined with " +
		"at, every, every_ms, or cron_expr"
	errAfterAliasConflict   = "cron: after and delay cannot both be set"
	errAfterMSDelayTooLarge = "cron: after_ms delay is too large"
	errHeadlessWithTarget   = "cron: headless jobs cannot set " +
		"channel or target"
	errDeliveryTargetUnavailable = "cron: current chat delivery " +
		"target is unavailable; pass channel/target explicitly " +
		"or set headless=true"
	cronTemplateHint = "For per-run counters or final-run text, " +
		"use Go text/template placeholders in message, " +
		"such as {{.Cron.RunIndex}}, {{.Cron.MaxRuns}}, " +
		"{{.Cron.RemainingRuns}}, and " +
		"{{if .Cron.IsFinalRun}}...{{end}}. Do not hardcode " +
		"counters like 1/5 into a repeating task."

	maxDelayMilliseconds = int64(1<<63-1) / int64(time.Millisecond)
)

// Tool exposes the scheduler to the model.
type Tool struct {
	svc *Service
}

// NewTool creates a cron tool for the scheduler service.
func NewTool(svc *Service) *Tool { _ = "STUB: not implemented"; return nil }

// SetService binds the scheduler service after tool construction.
func (t *Tool) SetService(svc *Service) { _ = "STUB: not implemented"; return }

func (t *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type toolInput struct {
	Action       string `json:"action"`
	JobID        string `json:"job_id,omitempty"`
	JobIDOld     string `json:"jobId,omitempty"`
	Name         string `json:"name,omitempty"`
	Message      string `json:"message,omitempty"`
	Prompt       string `json:"prompt,omitempty"`
	Task         string `json:"task,omitempty"`
	Enabled      *bool  `json:"enabled,omitempty"`
	ScheduleKind string `json:"schedule_kind,omitempty"`
	At           string `json:"at,omitempty"`
	RunAt        string `json:"run_at,omitempty"`
	RunAtOld     string `json:"runAt,omitempty"`
	Every        string `json:"every,omitempty"`
	Interval     string `json:"interval,omitempty"`
	Duration     string `json:"duration,omitempty"`
	EveryMS      *int64 `json:"every_ms,omitempty"`
	EveryMSOld   *int64 `json:"everyMs,omitempty"`
	After        string `json:"after,omitempty"`
	Delay        string `json:"delay,omitempty"`
	AfterMS      *int64 `json:"after_ms,omitempty"`
	DelayMS      *int64 `json:"delay_ms,omitempty"`
	AfterMSOld   *int64 `json:"afterMs,omitempty"`
	DelayMSOld   *int64 `json:"delayMs,omitempty"`
	CronExpr     string `json:"cron_expr,omitempty"`
	CronExprOld  string `json:"cronExpr,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	MaxRuns      *int   `json:"max_runs,omitempty"`
	MaxRunsOld   *int   `json:"maxRuns,omitempty"`
	EndsAt       string `json:"ends_at,omitempty"`
	EndsAtOld    string `json:"endsAt,omitempty"`
	Overlap      string `json:"overlap_policy,omitempty"`
	OverlapOld   string `json:"overlapPolicy,omitempty"`
	TimeoutSec   *int   `json:"timeout_sec,omitempty"`
	TimeoutOld   *int   `json:"timeoutSec,omitempty"`
	Channel      string `json:"channel,omitempty"`
	Target       string `json:"target,omitempty"`
	Headless     *bool  `json:"headless,omitempty"`
}

func (t *Tool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) list(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) add(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) update(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) remove(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) clear(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *Tool) runNow(
	ctx context.Context,
	in toolInput,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func scheduleFromInput(in toolInput, now time.Time) (Schedule, error) {
	_ = "STUB: not implemented"
	return *new(Schedule), nil
}

func policyFromInputWithError(
	in toolInput,
) (ExecutionPolicy, error) {
	_ = "STUB: not implemented"
	return *new(ExecutionPolicy), nil
}

func resolveScheduleKind(
	kind string,
	at string,
	every string,
	everyMS int64,
	cronExpr string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func resolveMessage(in toolInput) string { _ = "STUB: not implemented"; return "" }

func resolveAt(in toolInput, now time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveEvery(in toolInput) string { _ = "STUB: not implemented"; return "" }

func resolveDelay(in toolInput) (time.Duration, bool, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false, nil
}

func delayDurationInput(after string, delay string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func delayMillisecondsInput(values ...*int64) (int64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func resolveEndsAt(in toolInput) string { _ = "STUB: not implemented"; return "" }

func resolveJobID(in toolInput) string { _ = "STUB: not implemented"; return "" }

func currentUserID(ctx context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func currentOwnedJob(
	ctx context.Context,
	svc *Service,
	jobID string,
) (*Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveDelivery(
	ctx context.Context,
	channelID string,
	target string,
	headless bool,
) (outbound.DeliveryTarget, error) {
	_ = "STUB: not implemented"
	return *new(outbound.DeliveryTarget), nil
}

func optionalScopeDelivery(
	ctx context.Context,
	channelID string,
	target string,
) (outbound.DeliveryTarget, error) {
	_ = "STUB: not implemented"
	return *new(outbound.DeliveryTarget), nil
}

func isScheduledRunMutation(ctx context.Context, action string) bool {
	_ = "STUB: not implemented"
	return false
}

func hasScheduleInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func hasScheduleCoreInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func hasOnlyTimezoneScheduleInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func hasDelayInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func hasNonDelayScheduleInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func hasPolicyInput(in toolInput) bool { _ = "STUB: not implemented"; return false }

func firstIntValue(values ...*int) int { _ = "STUB: not implemented"; return 0 }

func firstInt64Value(values ...*int64) int64 { _ = "STUB: not implemented"; return 0 }

func firstString(values ...string) string { _ = "STUB: not implemented"; return "" }

func boolValue(value *bool) bool { _ = "STUB: not implemented"; return false }

func parseOptionalRFC3339(raw string) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ tool.CallableTool = (*Tool)(nil)
