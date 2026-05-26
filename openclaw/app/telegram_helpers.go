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
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/pairing"
	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	telegramChannelType = "telegram"

	defaultTelegramMaxRetries = 3
)

type telegramChannelConfig struct {
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
}

func resolveTelegramChannelSpecs(specs []pluginSpec) []pluginSpec {
	_ = "STUB: not implemented"
	return nil
}

func pairingStorePath(stateDir string, me tgapi.User) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func telegramClientNetOptions(
	cfg telegramChannelConfig,
) (tgapi.ClientNetOptions, error) {
	_ = "STUB: not implemented"
	return *new(tgapi.ClientNetOptions), nil
}

func pairingStoreOptions(rawPairingTTL string) ([]pairing.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
