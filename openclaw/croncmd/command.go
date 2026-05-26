//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package croncmd

import (
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwclient"
)

const (
	ActionHelp   = "help"
	ActionList   = "list"
	ActionStatus = "status"
	ActionStop   = "stop"
	ActionResume = "resume"
	ActionRemove = "remove"
	ActionClear  = "clear"
)

const shortJobIDSize = 8

var (
	ErrUnknownAction = errors.New("cron command: unknown action")
	ErrSelectorEmpty = errors.New("cron command: selector required")
	ErrSelectorMiss  = errors.New("cron command: job not found")
	ErrSelectorMany  = errors.New("cron command: job selector is ambiguous")
)

type Command struct {
	Action   string
	Selector string
}

func Parse(raw string) (Command, error) { _ = "STUB: not implemented"; return *new(Command), nil }

func NeedsSelector(action string) bool { _ = "STUB: not implemented"; return false }

func ResolveSelector(
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) (gwclient.ScheduledJobSummary, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary), nil
}

func ShortID(jobID string) string { _ = "STUB: not implemented"; return "" }

func isKnownAction(action string) bool { _ = "STUB: not implemented"; return false }

func parseIndex(raw string, size int) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func exactIDMatch(
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) (gwclient.ScheduledJobSummary, bool) {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary), false
}

func uniqueIDPrefixMatch(
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) (gwclient.ScheduledJobSummary, bool, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary), false, nil
}

func uniqueNameMatch(
	jobs []gwclient.ScheduledJobSummary,
	selector string,
) (gwclient.ScheduledJobSummary, bool, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary), false, nil
}
