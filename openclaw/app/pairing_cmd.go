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

	tgch "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/channel/telegram"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/pairing"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	subcmdPairing = "pairing"

	flagConfig   = "config"
	flagChannel  = "channel"
	flagStateDir = "state-dir"

	pairingCmdList    = "list"
	pairingCmdApprove = "approve"
)

var probeBotInfo = func(
	ctx context.Context,
	token string,
	opts ...tgapi.Option,
) (tgch.BotInfo, error) {
	return tgch.ProbeBotInfo(ctx, token, opts...)
}

func runPairing(args []string) int { _ = "STUB: not implemented"; return 0 }

func printPairingUsage() { _ = "STUB: not implemented"; return }

func runPairingList(
	ctx context.Context,
	store *pairing.FileStore,
) int {
	_ = "STUB: not implemented"
	return 0
}

func runPairingApprove(
	ctx context.Context,
	store *pairing.FileStore,
	code string,
) int {
	_ = "STUB: not implemented"
	return 0
}

func openPairingStore(
	ctx context.Context,
	opts runOptions,
	wantChannel string,
) (*pairing.FileStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveTelegramPairingChannel(
	opts runOptions,
	wantChannel string,
) (pluginSpec, error) {
	_ = "STUB: not implemented"
	return *new(pluginSpec), nil
}

func normalizePairingArgs(args []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func flagName(arg string) string { _ = "STUB: not implemented"; return "" }
