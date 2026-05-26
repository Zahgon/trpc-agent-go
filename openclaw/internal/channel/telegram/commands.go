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
	"sync"

	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	commandPrefix = "/"

	commandHelp   = "help"
	commandCancel = "cancel"

	commandReset     = "reset"
	commandNew       = "new"
	commandForget    = "forget"
	commandCron      = "cron"
	commandJobs      = "jobs"
	commandJobsClear = "jobs_clear"
	commandPersona   = "persona"
	commandPersonas  = "personas"
)

const (
	commandHelpDesc      = "Show help"
	commandCancelDesc    = "Cancel the current run"
	commandResetDesc     = "Start a new DM session"
	commandNewDesc       = "Alias of /reset"
	commandForgetDesc    = "Delete your saved data (DM only)"
	commandCronDesc      = "Manage scheduled jobs for this chat"
	commandJobsDesc      = "List scheduled jobs for this chat"
	commandJobsClearDesc = "Remove scheduled jobs for this chat"
	commandPersonaDesc   = "Show or set the active persona preset"
	commandPersonasDesc  = "List available persona presets"
)

const helpMessage = "Commands:\n" +
	"/help   " + commandHelpDesc + "\n" +
	"/cancel " + commandCancelDesc + "\n" +
	"/reset  " + commandResetDesc + "\n" +
	"/new    " + commandNewDesc + "\n" +
	"/forget " + commandForgetDesc + "\n" +
	"/cron   " + commandCronDesc + "\n" +
	"/persona " + commandPersonaDesc + "\n" +
	"/personas " + commandPersonasDesc

func defaultBotCommands() []tgapi.BotCommand { _ = "STUB: not implemented"; return nil }

type commandCall struct {
	Name string
	Args string
}

func parseCommand(text string, bot BotInfo) string { _ = "STUB: not implemented"; return "" }

func parseCommandCall(text string, bot BotInfo) commandCall {
	_ = "STUB: not implemented"
	return *new(commandCall)
}

type inflightRequests struct {
	mu sync.Mutex
	m  map[string]string
}

func newInflightRequests() *inflightRequests { _ = "STUB: not implemented"; return nil }

func (r *inflightRequests) Get(sessionID string) string { _ = "STUB: not implemented"; return "" }

func (r *inflightRequests) Set(sessionID, requestID string) { _ = "STUB: not implemented"; return }

func (r *inflightRequests) Clear(sessionID, requestID string) { _ = "STUB: not implemented"; return }

type laneLocker struct {
	mu    sync.Mutex
	lanes map[string]*laneEntry
}

type laneEntry struct {
	lock sync.Mutex
	refs int
}

func newLaneLocker() *laneLocker { _ = "STUB: not implemented"; return nil }

func (l *laneLocker) withLock(key string, fn func()) { _ = "STUB: not implemented"; return }

func (l *laneLocker) acquire(key string) *laneEntry { _ = "STUB: not implemented"; return nil }

func (l *laneLocker) release(key string, entry *laneEntry) { _ = "STUB: not implemented"; return }
