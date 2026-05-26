//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	"context"
	"time"

	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const subcmdDoctor = "doctor"

const (
	telegramLongPollTimeout = 25 * time.Second
	telegramTimeoutSlack    = 5 * time.Second
)

func runDoctor(args []string) int { _ = "STUB: not implemented"; return 0 }

func printBot(me tgapi.User) { _ = "STUB: not implemented"; return }

func checkTimeout(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

func checkWebhook(ctx context.Context, c *tgapi.Client) bool {
	_ = "STUB: not implemented"
	return false
}

func checkPolicies(
	dmPolicy string,
	groupPolicy string,
	allowUsers []string,
	allowThreads []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func checkPairingStore(
	ctx context.Context,
	rawStateDir string,
	dmPolicy string,
	rawPairingTTL string,
	me tgapi.User,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isPolicy(raw string, want string) bool { _ = "STUB: not implemented"; return false }
