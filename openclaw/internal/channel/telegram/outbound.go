//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package telegram

import (
	"context"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	sessionDMPrefix     = channelID + ":dm:"
	sessionThreadPrefix = channelID + ":thread:"

	fileURLPrefix = "file://"

	mimePrefixImage = "image/"
	mimePrefixAudio = "audio/"
	mimePrefixVideo = "video/"
	mimePrefixText  = "text/"

	mimeVoiceOGG    = "audio/ogg"
	mimeOctetStream = "application/octet-stream"

	uploadModeDocument = "document"
	uploadModePhoto    = "photo"
	uploadModeAudio    = "audio"
	uploadModeVoice    = "voice"
	uploadModeVideo    = "video"

	workspaceFileFallback = "workspace-output"

	maxTelegramCaptionRunes = 1024
	maxTelegramExpandedDir  = 32
)

type outboundFilePayload struct {
	Name       string
	Data       []byte
	SourcePath string
}

// ResolveTextTargetFromSessionID converts a Telegram session id into the
// channel-specific outbound target used by SendText.
func ResolveTextTargetFromSessionID(sessionID string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SendText implements channel.TextSender for Telegram.
func (c *Channel) SendText(
	ctx context.Context,
	target string,
	text string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SendMessage implements channel.MessageSender for Telegram.
func (c *Channel) SendMessage(
	ctx context.Context,
	target string,
	msg channel.OutboundMessage,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) expandTelegramOutboundFiles(
	ctx context.Context,
	files []channel.OutboundFile,
) ([]channel.OutboundFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) expandTelegramOutboundFile(
	ctx context.Context,
	file channel.OutboundFile,
) ([]channel.OutboundFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isOpaqueOutboundRef(raw string) bool { _ = "STUB: not implemented"; return false }

func (c *Channel) resolveTelegramOutboundExistingPath(
	ctx context.Context,
	raw string,
) (string, os.FileInfo, bool) {
	_ = "STUB: not implemented"
	return "", *new(os.FileInfo), false
}

func outboundSessionUploadsRoot(
	ctx context.Context,
	stateRoot string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func parseTextTarget(target string) (int64, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func parseDMSessionTarget(raw string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func parseThreadSessionTarget(raw string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseLegacyDMSessionTarget(raw string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func leadingSessionToken(raw string) string { _ = "STUB: not implemented"; return "" }

func parseChatID(raw string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseTopicID(raw string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Channel) sendFile(
	ctx context.Context,
	chatID int64,
	threadID int,
	file channel.OutboundFile,
	caption string,
	plainCaption string,
	parseMode string,
	scope uploads.Scope,
	asVoice bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) recordSentFiles(
	ctx context.Context,
	chatID int64,
	threadID int,
	payload outboundFilePayload,
	savedPath string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *Channel) sendFileByMode(
	ctx context.Context,
	mode string,
	params tgapi.SendFileParams,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) persistDerivedOutboundFile(
	ctx context.Context,
	payload outboundFilePayload,
	scope uploads.Scope,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Channel) annotateExistingOutboundFile(
	ctx context.Context,
	store *uploads.Store,
	payload outboundFilePayload,
	scope uploads.Scope,
) string {
	_ = "STUB: not implemented"
	return ""
}

func outboundUploadScopeFromContext(
	ctx context.Context,
) uploads.Scope {
	_ = "STUB: not implemented"
	return *new(uploads.Scope)
}

func isValidUploadScope(scope uploads.Scope) bool { _ = "STUB: not implemented"; return false }

func telegramCaptionParts(
	text string,
	stateDir string,
) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func resolveOutboundFile(
	ctx context.Context,
	stateRoot string,
	file channel.OutboundFile,
) (outboundFilePayload, error) {
	_ = "STUB: not implemented"
	return *new(outboundFilePayload), nil
}

func resolveArtifactOutboundFile(
	ctx context.Context,
	raw string,
	nameHint string,
) (outboundFilePayload, error) {
	_ = "STUB: not implemented"
	return *new(outboundFilePayload), nil
}

func resolveWorkspaceOutboundFile(
	ctx context.Context,
	raw string,
	nameHint string,
) (outboundFilePayload, error) {
	_ = "STUB: not implemented"
	return *new(outboundFilePayload), nil
}

func resolveHostOutboundFile(
	ctx context.Context,
	stateRoot string,
	raw string,
	nameHint string,
) (outboundFilePayload, error) {
	_ = "STUB: not implemented"
	return *new(outboundFilePayload), nil
}

func resolveOutboundFilePath(
	ctx context.Context,
	stateRoot string,
	raw string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveSessionScopedOutboundPath(
	ctx context.Context,
	stateRoot string,
	raw string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func resolveSessionScopedJoinedPath(
	root string,
	raw string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func expandHomePath(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func withArtifactContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func detectUploadMode(
	name string,
	data []byte,
	asVoice bool,
) string {
	_ = "STUB: not implemented"
	return ""
}

func detectMediaType(name string, data []byte) string { _ = "STUB: not implemented"; return "" }

func typeFromExtension(ext string) string { _ = "STUB: not implemented"; return "" }

func isVoiceMedia(contentType string, name string) bool { _ = "STUB: not implemented"; return false }

func isVoiceCompatibleMedia(contentType string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func strconvParseInt(raw string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func strconvAtoi(raw string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func isDigitsOnly(raw string) bool { _ = "STUB: not implemented"; return false }
