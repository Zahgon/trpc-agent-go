//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telegram

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwclient"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/persona"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

func buildRequestID(
	chatID int64,
	messageThreadID int,
	messageID int,
) string {
	_ = "STUB: not implemented"
	return ""
}

func buildLaneKey(fromID string, thread string) string { _ = "STUB: not implemented"; return "" }

func parseDMPolicy(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseGroupPolicy(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseDMBlockCleanup(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func splitRunes(text string, maxRunes int) []string { _ = "STUB: not implemented"; return nil }

func splitIndex(segment []rune, maxRunes int) int { _ = "STUB: not implemented"; return 0 }

func resolveStateDir(stateDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func newOffsetStore(
	stateDir string,
	bot BotInfo,
) (*tgapi.FileOffsetStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func offsetKey(bot BotInfo) string { _ = "STUB: not implemented"; return "" }

func sanitizeFileToken(value string) string { _ = "STUB: not implemented"; return "" }

// PairingStorePath returns the path used for storing DM pairing state.
func PairingStorePath(stateDir string, bot BotInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *laneLocker) withLockErr(key string, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) reply(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	text string,
) {
	_ = "STUB: not implemented"
	return
}

const (
	cancelNoopMessage   = "No running request to cancel."
	cancelFailedMessage = "Cancel failed."
	cancelOKMessage     = "Canceled."

	resetOKMessage     = "Started a new session."
	resetFailedMessage = "Failed to start a new session."

	forgetOKMessage          = "Forgot your data."
	forgetFailedMessage      = "Failed to forget your data."
	forgetUnsupportedMessage = "Forget is not supported."

	jobsUnsupportedMessage = "Scheduled job management is not supported."
	jobsUsageMessage       = "Usage:\n" +
		"/cron list\n" +
		"/cron status <index|id>\n" +
		"/cron stop <index|id>\n" +
		"/cron resume <index|id>\n" +
		"/cron remove <index|id>\n" +
		"/cron clear"
	jobsListFailedMessage   = "Failed to list scheduled jobs."
	jobsClearFailedMessage  = "Failed to clear scheduled jobs."
	jobsUpdateFailedMessage = "Failed to update the scheduled job."
	jobsRemoveFailedMessage = "Failed to remove the scheduled job."
	jobsEmptyMessage        = "No scheduled jobs for this chat."
	jobsClearNoopMessage    = "No scheduled jobs to clear for this chat."
	jobsMessageHeader       = "Scheduled jobs for this chat:"
	jobsClearOKFmt          = "Cleared %d scheduled job(s) for this chat."
	jobsStopOKFmt           = "Stopped scheduled job %s."
	jobsResumeOKFmt         = "Resumed scheduled job %s."
	jobsRemoveOKFmt         = "Removed scheduled job %s."
	jobsStatusHeader        = "Scheduled job details:"
	jobsSelectorHint        = "Use the list index or a unique job id prefix."
	jobTimeLayout           = "2006-01-02 15:04:05 MST"

	personaUnsupportedMessage = "Preset personas are not supported."
	personaListFailedMessage  = "Failed to load persona presets."
	personaSetFailedMessage   = "Failed to update the persona preset."
	personaUnknownMessage     = "Unknown persona preset. " +
		"Use the personas list command."
	personaResetOKMessage = "Persona reset to default."
	personaSetOKFmt       = "Persona set to %s."
	personaMessageHeader  = "Persona presets for this chat:"
	personaCurrentPrefix  = "Current: "
	personaUsageMessage   = "Tap a button below to switch instantly. " +
		"Use /persona <id> if you prefer typing. Use default " +
		"to reset."

	personaCallbackPrefix      = "persona:set:"
	personaButtonActivePrefix  = "> "
	personaKeyboardColumns     = 2
	personaSelectionFailedHint = "Could not update the preset."
)

func (c *Channel) handleCancelCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	laneKey string,
) error {
	_ = "STUB: not implemented"
	return nil
}

type userForgetter interface {
	ForgetUser(ctx context.Context, channel, userID string) error
}

type scheduledJobManager interface {
	ListScheduledJobs(
		ctx context.Context,
		channel string,
		userID string,
		target string,
	) ([]gwclient.ScheduledJobSummary, error)
	ClearScheduledJobs(
		ctx context.Context,
		channel string,
		userID string,
		target string,
	) (int, error)
	SetScheduledJobEnabled(
		ctx context.Context,
		channel string,
		userID string,
		target string,
		jobID string,
		enabled bool,
	) (gwclient.ScheduledJobSummary, error)
	RemoveScheduledJob(
		ctx context.Context,
		channel string,
		userID string,
		target string,
		jobID string,
	) (bool, error)
}

type personaManager interface {
	ListPresetPersonas() []persona.Preset
	GetPresetPersona(
		ctx context.Context,
		scopeKey string,
	) (persona.Preset, error)
	SetPresetPersona(
		ctx context.Context,
		scopeKey string,
		presetID string,
	) (persona.Preset, error)
}

func (c *Channel) cancelInflight(
	ctx context.Context,
	laneKey string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Channel) handleResetCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	laneKey string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleForgetCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	laneKey string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleCronCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	userID string,
	args string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleJobsCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleJobsClearCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) replySelectedJob(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) setSelectedJobEnabled(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	manager scheduledJobManager,
	userID string,
	target string,
	jobs []gwclient.ScheduledJobSummary,
	action string,
	selector string,
	enabled bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) removeSelectedJob(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	manager scheduledJobManager,
	userID string,
	target string,
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handlePersonaCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	scopeKey string,
	args string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handlePersonasCommand(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	scopeKey string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) replyPersonaSummary(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	scopeKey string,
	manager personaManager,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handlePersonaCallbackQuery(
	ctx context.Context,
	q tgapi.CallbackQuery,
	scopeKey string,
	messageThreadID int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func currentChatTarget(chatID int64, messageThreadID int) string {
	_ = "STUB: not implemented"
	return ""
}

func formatScheduledJobsMessage(
	jobs []gwclient.ScheduledJobSummary,
) string {
	_ = "STUB: not implemented"
	return ""
}

func formatScheduledJobDetails(
	job gwclient.ScheduledJobSummary,
) string {
	_ = "STUB: not implemented"
	return ""
}

func formatScheduledJobLine(job gwclient.ScheduledJobSummary) string {
	_ = "STUB: not implemented"
	return ""
}

func scheduledJobDisplayName(job gwclient.ScheduledJobSummary) string {
	_ = "STUB: not implemented"
	return ""
}

func formatJobEnabled(enabled bool) string { _ = "STUB: not implemented"; return "" }

func valueOrDash(text string) string { _ = "STUB: not implemented"; return "" }

func formatScheduledJobRunCount(job gwclient.ScheduledJobSummary) string {
	_ = "STUB: not implemented"
	return ""
}

func formatScheduledJobOverlap(policy string) string { _ = "STUB: not implemented"; return "" }

func trimJobText(text string) string { _ = "STUB: not implemented"; return "" }

func firstCommandArg(args string) string { _ = "STUB: not implemented"; return "" }

func formatPersonaMessage(
	current persona.Preset,
	presets []persona.Preset,
) string {
	_ = "STUB: not implemented"
	return ""
}

func formatPersonaLine(
	preset persona.Preset,
	currentID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func personaSummaryData(
	ctx context.Context,
	scopeKey string,
	manager personaManager,
) (persona.Preset, []persona.Preset, error) {
	_ = "STUB: not implemented"
	return *new(persona.Preset), nil, nil
}

func personaReplyMarkup(
	current persona.Preset,
	presets []persona.Preset,
) *tgapi.InlineKeyboardMarkup {
	_ = "STUB: not implemented"
	return nil
}

func personaButtonText(
	preset persona.Preset,
	currentID string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func isPersonaCallbackData(data string) bool { _ = "STUB: not implemented"; return false }

func personaPresetIDFromCallback(data string) string { _ = "STUB: not implemented"; return "" }

func personaSelectionText(preset persona.Preset) string { _ = "STUB: not implemented"; return "" }
