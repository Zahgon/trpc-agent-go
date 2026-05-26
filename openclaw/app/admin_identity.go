//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	adminIdentityFileName = "IDENTITY.md"

	adminIdentityFilePerm = 0o600
	adminIdentityDirPerm  = 0o700

	adminAssistantNameMaxRunes = 32

	adminIdentityFallbackRuntime = "runtime product"

	adminDefaultNameSourceFile = "Default name from IDENTITY.md"
	adminDefaultNameSourceApp  = "Default name from runtime product"

	adminChatsHelpText = "" +
		"This runtime uses one default name across chats. " +
		"The Chats page shows tracked conversation scopes and " +
		"their recent history, but it does not keep a separate " +
		"current name for each chat."

	adminIdentityTrimCutset = "" +
		"\"'“”‘’<>《》「」『』【】()（）[]"

	adminChatKindTracked = "tracked"
	adminChatKindLabel   = "Tracked chat"

	adminChatHistorySessionLimit    = 12
	adminChatHistoryVisibleCount    = 5
	adminChatTranscriptSessionLimit = 8
	adminChatTranscriptVisibleCount = 2
	adminChatTranscriptTurnLimit    = 18
	adminChatTranscriptTurnVisible  = 6
	adminChatTranscriptTextLimit    = 1500
)

type adminIdentityProvider struct {
	filePath       string
	runtimeProduct string
}

type adminChatsProvider struct {
	identity *adminIdentityProvider
	appName  string
	session  session.Service
}

func buildAdminIdentityProvider(
	stateDir string,
	runtimeProduct string,
) *adminIdentityProvider {
	_ = "STUB: not implemented"
	return nil
}

func buildAdminChatsProvider(
	identity *adminIdentityProvider,
	appName string,
	sessionSvc session.Service,
) *adminChatsProvider {
	_ = "STUB: not implemented"
	return nil
}

func defaultAdminRuntimeProduct(raw string) string { _ = "STUB: not implemented"; return "" }

func normalizeAdminAssistantName(raw string) string { _ = "STUB: not implemented"; return "" }

func readAdminAssistantName(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func writeAdminAssistantName(
	path string,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func identityDefaultNameSource(
	status admin.IdentityStatus,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *adminIdentityProvider) IdentityStatus() (
	admin.IdentityStatus,
	error,
) {
	_ = "STUB: not implemented"
	return *new(admin.IdentityStatus), nil
}

func (p *adminIdentityProvider) SaveAssistantName(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminChatsProvider) ChatsStatus() (
	admin.ChatsStatus,
	error,
) {
	_ = "STUB: not implemented"
	return *new(admin.ChatsStatus), nil
}

func (p *adminChatsProvider) ChatDetail(
	baseSessionID string,
) (admin.ChatView, error) {
	_ = "STUB: not implemented"
	return *new(admin.ChatView), nil
}

func (p *adminChatsProvider) chatViews(
	defaultName string,
	defaultSource string,
) ([]admin.ChatView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *adminChatsProvider) chatSessions(
	baseSessionID string,
) ([]*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildAdminTrackedChatView(
	baseSessionID string,
	defaultName string,
	defaultSource string,
	sessions []*session.Session,
) admin.ChatView {
	_ = "STUB: not implemented"
	return *new(admin.ChatView)
}

func buildAdminChatTranscript(
	appName string,
	sessionSvc session.Service,
	baseSessionID string,
	sessions []*session.Session,
) ([]admin.ChatTranscriptView, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func buildAdminChatTranscriptView(
	appName string,
	sessionSvc session.Service,
	baseSessionID string,
	currentSessionID string,
	sessionMeta *session.Session,
) (admin.ChatTranscriptView, bool, error) {
	_ = "STUB: not implemented"
	return *new(admin.ChatTranscriptView), false, nil
}

func trimAdminChatTranscriptText(text string) string { _ = "STUB: not implemented"; return "" }

func sessionActivityTime(sess *session.Session) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func timeZero() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
