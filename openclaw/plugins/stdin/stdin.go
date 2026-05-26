//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package stdin registers a simple "stdin channel" plugin.
//
// It is intended as a reference implementation for writing custom
// channels. The channel reads one line per message from STDIN, sends it
// to the OpenClaw gateway, and prints the reply to STDOUT.
package stdin

import (
	"context"

	occhannel "trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

const (
	pluginType = "stdin"

	defaultFrom           = "local"
	defaultUserLabel      = "User"
	defaultAssistantLabel = "Assistant"

	exitCmd1 = "/exit"
	exitCmd2 = "/quit"

	defaultScannerBufBytes = 64 * 1024
	defaultScannerMaxBytes = 1 << 20
)

func init() {
	if err := registry.RegisterChannel(pluginType, newChannel); err != nil {
		panic(err)
	}
}

type channelCfg struct {
	From           string `yaml:"from"`
	Thread         string `yaml:"thread"`
	MaxLineBytes   int    `yaml:"max_line_bytes"`
	ShowPrompt     bool   `yaml:"show_prompt,omitempty"`
	ShowRoleLabels bool   `yaml:"show_role_labels,omitempty"`
	UserLabel      string `yaml:"user_label,omitempty"`
	AssistantLabel string `yaml:"assistant_label,omitempty"`
}

func newChannel(
	deps registry.ChannelDeps,
	spec registry.PluginSpec,
) (occhannel.Channel, error) {
	_ = "STUB: not implemented"
	return *new(occhannel.Channel), nil
}

type channel struct {
	id     string
	gw     registry.GatewayClient
	from   string
	thread string

	showPrompt     bool
	showRoleLabels bool
	userLabel      string
	assistantLabel string

	bufBytes     int
	maxLineBytes int
}

func (c *channel) ID() string { _ = "STUB: not implemented"; return "" }

func (c *channel) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func defaultLabel(raw string, fallback string) string { _ = "STUB: not implemented"; return "" }

func (c *channel) printPrompt() { _ = "STUB: not implemented"; return }

func (c *channel) printPromptTerminator() { _ = "STUB: not implemented"; return }

func (c *channel) printReply(reply string) { _ = "STUB: not implemented"; return }
