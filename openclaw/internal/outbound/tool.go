//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package outbound

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolMessage = "message"

	maxExpandedFileCount = 32
)

// Tool sends plain text messages through OpenClaw channels.
type Tool struct {
	router *Router
}

// NewTool creates a message tool backed by the outbound router.
func NewTool(router *Router) *Tool { _ = "STUB: not implemented"; return nil }

func (t *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type toolInput struct {
	Text         string   `json:"text"`
	File         string   `json:"file,omitempty"`
	Files        []string `json:"files,omitempty"`
	Media        []string `json:"media,omitempty"`
	Channel      string   `json:"channel,omitempty"`
	Target       string   `json:"target,omitempty"`
	AsVoice      bool     `json:"as_voice,omitempty"`
	AudioAsVoice bool     `json:"audio_as_voice,omitempty"`
}

func (t *Tool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func recordSentText(
	ctx context.Context,
	target DeliveryTarget,
	msg channel.OutboundMessage,
) {
	_ = "STUB: not implemented"
	return
}

func buildOutboundMessage(in toolInput) (channel.OutboundMessage, error) {
	_ = "STUB: not implemented"
	return *new(channel.OutboundMessage), nil
}

func expandOutboundFiles(paths []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandOutboundPath(raw string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func isOpaqueRef(raw string) bool { _ = "STUB: not implemented"; return false }

func expandOutboundGlob(raw string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func expandOutboundDirectory(raw string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func absPaths(paths []string) []string { _ = "STUB: not implemented"; return nil }

func collectPaths(groups ...any) []string { _ = "STUB: not implemented"; return nil }

func appendPath(
	out []string,
	seen map[string]struct{},
	value string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

var _ tool.CallableTool = (*Tool)(nil)
