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
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
)

// DeliveryTarget identifies a channel-specific outbound destination.
type DeliveryTarget struct {
	Channel string `json:"channel,omitempty"`
	Target  string `json:"target,omitempty"`
}

// Router dispatches outbound messages to registered channels.
type Router struct {
	mu             sync.RWMutex
	textSenders    map[string]channel.TextSender
	messageSenders map[string]channel.MessageSender
}

// NewRouter creates an empty outbound router.
func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// Register adds a channel sender when the channel implements TextSender.
func (r *Router) Register(ch channel.Channel) { _ = "STUB: not implemented"; return }

// RegisterSender adds or replaces a sender for its channel id.
func (r *Router) RegisterSender(sender channel.TextSender) { _ = "STUB: not implemented"; return }

// RegisterMessageSender adds or replaces a media-capable sender.
func (r *Router) RegisterMessageSender(sender channel.MessageSender) {
	_ = "STUB: not implemented"
	return
}

// Channels returns the sorted list of registered channel ids.
func (r *Router) Channels() []string { _ = "STUB: not implemented"; return nil }

// SendText delivers plain text through the selected channel.
func (r *Router) SendText(
	ctx context.Context,
	target DeliveryTarget,
	text string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SendMessage delivers text and optional local files through the selected
// channel.
func (r *Router) SendMessage(
	ctx context.Context,
	target DeliveryTarget,
	msg channel.OutboundMessage,
) error {
	_ = "STUB: not implemented"
	return nil
}
