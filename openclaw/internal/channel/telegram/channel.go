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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwclient"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	channelID = "telegram"

	requestIDPrefix = "telegram:"

	threadTopicSep = ":topic:"

	maxReplyRunes = 4000

	tgChatTypePrivate = "private"

	tgChatMemberStatusKicked = "kicked"
	tgChatMemberStatusLeft   = "left"

	defaultStateRootDir = ".trpc-agent-go-github"
	defaultStateAppName = "openclaw"

	mentionPrefix = "@"

	offsetStoreDir = "telegram"

	offsetStoreFilePrefix = "update-offset-"
	offsetStoreFileSuffix = ".json"

	defaultOffsetKey = "default"

	pairingStoreFilePrefix = "pairing-"
	pairingStoreFileSuffix = ".json"

	dmPolicyDisabled  = "disabled"
	dmPolicyOpen      = "open"
	dmPolicyAllowlist = "allowlist"
	dmPolicyPairing   = "pairing"

	groupPolicyDisabled  = "disabled"
	groupPolicyOpen      = "open"
	groupPolicyAllowlist = "allowlist"

	defaultDMPolicy    = dmPolicyPairing
	defaultGroupPolicy = groupPolicyDisabled

	dmBlockCleanupNone   = "none"
	dmBlockCleanupReset  = "reset"
	dmBlockCleanupForget = "forget"

	defaultDMBlockCleanup = dmBlockCleanupReset

	defaultPairingTTL = time.Hour

	defaultRegisterCommands = true

	defaultMaxDownloadMiB         = 20
	defaultMaxDownloadBytes int64 = defaultMaxDownloadMiB << 20
)

// ChannelName is the stable channel identifier used across OpenClaw.
const ChannelName = channelID

const (
	notAllowedMessage = "You are not allowed to use this bot."

	pairingMessageTemplate = `Pairing required.

Code: %s

Ask the operator to approve:
openclaw pairing approve %s -config <CONFIG>`

	errNonPositiveMaxDownloadBytes = "telegram: non-positive max download bytes"
)

type gatewayClient interface {
	SendMessage(
		ctx context.Context,
		req gwclient.MessageRequest,
	) (gwclient.MessageResponse, error)

	Cancel(ctx context.Context, requestID string) (bool, error)
}

type botAPI interface {
	GetUpdates(
		ctx context.Context,
		offset int,
		timeout time.Duration,
	) ([]tgapi.Update, error)

	SendMessage(
		ctx context.Context,
		params tgapi.SendMessageParams,
	) (tgapi.Message, error)

	AnswerCallbackQuery(
		ctx context.Context,
		params tgapi.AnswerCallbackQueryParams,
	) error

	SendDocument(
		ctx context.Context,
		params tgapi.SendFileParams,
	) (tgapi.Message, error)

	SendPhoto(
		ctx context.Context,
		params tgapi.SendFileParams,
	) (tgapi.Message, error)

	SendAudio(
		ctx context.Context,
		params tgapi.SendFileParams,
	) (tgapi.Message, error)

	SendVoice(
		ctx context.Context,
		params tgapi.SendFileParams,
	) (tgapi.Message, error)

	SendVideo(
		ctx context.Context,
		params tgapi.SendFileParams,
	) (tgapi.Message, error)

	EditMessageText(
		ctx context.Context,
		params tgapi.EditMessageTextParams,
	) (tgapi.Message, error)

	SendChatAction(
		ctx context.Context,
		params tgapi.SendChatActionParams,
	) error

	SetMyCommands(
		ctx context.Context,
		params tgapi.SetMyCommandsParams,
	) error

	DownloadFileByID(
		ctx context.Context,
		fileID string,
		maxBytes int64,
	) (tgapi.File, []byte, error)
}

// BotInfo represents Telegram bot metadata used by the channel.
type BotInfo struct {
	ID       int64
	Username string
	Mention  string
}

// ProbeBotInfo fetches bot metadata via getMe.
func ProbeBotInfo(
	ctx context.Context,
	token string,
	opts ...tgapi.Option,
) (BotInfo, error) {
	_ = "STUB: not implemented"
	return *new(BotInfo), nil
}

func mentionFromUsername(username string) string { _ = "STUB: not implemented"; return "" }

type config struct {
	stateDir        string
	startFromLatest bool
	pollTimeout     time.Duration
	errorBackoff    time.Duration

	dmPolicy    string
	groupPolicy string

	allowUsers   map[string]struct{}
	allowThreads map[string]struct{}

	pairingTTL time.Duration

	apiOptions []tgapi.Option

	maxDownloadBytes int64

	streamingMode string

	dmResetPolicy dmSessionResetPolicy

	dmBlockCleanup string

	registerCommands bool
}

// Option configures the Telegram channel.
type Option func(*config)

// WithStateDir sets the state directory for offsets.
func WithStateDir(dir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStartFromLatest controls whether the poller drains pending
// updates when no stored offset exists yet.
func WithStartFromLatest(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPollTimeout sets the long-poll timeout.
func WithPollTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithErrorBackoff sets the delay after polling/handler errors.
func WithErrorBackoff(backoff time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDMPolicy sets the policy for direct messages.
func WithDMPolicy(policy string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGroupPolicy sets the policy for group and thread messages.
func WithGroupPolicy(policy string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAllowUsers sets a per-channel allowlist.
func WithAllowUsers(users ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAllowThreads sets an allowlist for group chats and topics.
//
// Values should match the `thread` field derived by this channel:
//   - Group chat: "<chat_id>"
//   - Forum topic: "<chat_id>:topic:<message_thread_id>"
func WithAllowThreads(threads ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPairingTTL sets how long pairing codes stay valid.
func WithPairingTTL(ttl time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIOptions passes options to the underlying Telegram API client.
func WithAPIOptions(opts ...tgapi.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxDownloadBytes sets the per-file download limit for Telegram
// attachments.
func WithMaxDownloadBytes(maxBytes int64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamingMode controls how replies are delivered to Telegram.
func WithStreamingMode(mode string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDMSessionIdleReset configures an automatic reset when a DM
// stays idle longer than the given duration.
func WithDMSessionIdleReset(idle time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDMSessionDailyReset configures an automatic reset when the date
// changes (local time).
func WithDMSessionDailyReset(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDMBlockCleanup configures what happens when the bot is blocked.
//
// Supported values:
//   - "none":  keep server state intact
//   - "reset": rotate to a new active session
//   - "forget": delete sessions, memories, and debug traces
func WithDMBlockCleanup(action string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRegisterCommands controls whether the bot registers slash commands
// with Telegram on startup.
func WithRegisterCommands(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type pairingStore interface {
	IsApproved(ctx context.Context, userID string) (bool, error)
	Request(
		ctx context.Context,
		userID string,
	) (string, bool, error)
}

// Channel implements a Telegram long-polling chat surface.
type Channel struct {
	bot   botAPI
	info  BotInfo
	gw    gatewayClient
	store tgapi.OffsetStore
	state string

	sentFiles *sentFileTracker

	audioInputConverter audioInputConverter

	dmSessions     *dmSessionStore
	dmResetPolicy  dmSessionResetPolicy
	dmBlockCleanup string

	startFromLatest bool
	pollTimeout     time.Duration
	errorBackoff    time.Duration

	dmPolicy    string
	groupPolicy string

	allowUsers   map[string]struct{}
	allowThreads map[string]struct{}

	pairing pairingStore

	maxDownloadBytes int64

	streamingMode string

	registerCommands bool

	lanes    *laneLocker
	inflight *inflightRequests
}

// New creates a Telegram channel. It persists polling offsets under
// the configured state directory.
func New(
	token string,
	bot BotInfo,
	gw gatewayClient,
	opts ...Option,
) (*Channel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ID returns the channel identifier used by the gateway.
func (c *Channel) ID() string {
	_ = "STUB: not implemented"

	// Run starts polling Telegram and blocks until ctx is done.
	return ""
}

func (c *Channel) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Channel) registerBotCommands(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleMessage(
	ctx context.Context,
	msg tgapi.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleMyChatMember(
	ctx context.Context,
	ev tgapi.ChatMemberEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) handleCallbackQuery(
	ctx context.Context,
	q tgapi.CallbackQuery,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) sendDM(
	ctx context.Context,
	chatID int64,
	text string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *Channel) answerCallbackQuery(
	ctx context.Context,
	callbackID string,
	text string,
	showAlert bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) isUserAllowed(userID string) bool { _ = "STUB: not implemented"; return false }

func (c *Channel) isChatAllowed(isGroup bool, thread string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Channel) isDMAllowed(
	ctx context.Context,
	chatID int64,
	fromID string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
