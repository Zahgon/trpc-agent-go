//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package telegram registers the Telegram channel plugin.
package telegram

import (
	occhannel "trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	tgch "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/channel/telegram"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

const (
	pluginType = "telegram"

	defaultStreamingMode = "progress"
	defaultMaxRetries    = 3

	errMissingToken = "telegram channel: missing config.token"
)

func init() {
	if err := registry.RegisterChannel(pluginType, newChannel); err != nil {
		panic(err)
	}
}

type channelCfg struct {
	Token string `yaml:"token"`

	StartFromLatest *bool `yaml:"start_from_latest"`

	Proxy       string `yaml:"proxy"`
	HTTPTimeout string `yaml:"http_timeout"`
	MaxRetries  *int   `yaml:"max_retries"`

	Streaming    string   `yaml:"streaming"`
	DMPolicy     string   `yaml:"dm_policy"`
	GroupPolicy  string   `yaml:"group_policy"`
	AllowThreads []string `yaml:"allow_threads"`
	PairingTTL   string   `yaml:"pairing_ttl"`

	MaxDownloadBytes *int64 `yaml:"max_download_bytes"`

	SessionResetIdle  string `yaml:"session_reset_idle"`
	SessionResetDaily *bool  `yaml:"session_reset_daily"`

	OnBlock string `yaml:"on_block"`
}

func newChannel(
	deps registry.ChannelDeps,
	spec registry.PluginSpec,
) (occhannel.Channel, error) {
	_ = "STUB: not implemented"
	return *new(occhannel.Channel), nil
}

func resolveStreamingMode(raw string) string { _ = "STUB: not implemented"; return "" }

func logTelegramBot(bot tgch.BotInfo) { _ = "STUB: not implemented"; return }

func makeTelegramAPIOptions(cfg channelCfg) ([]tgapi.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
