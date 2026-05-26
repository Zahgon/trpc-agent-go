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
	"context"
	"time"
)

const (
	defaultPollTimeout  = 25 * time.Second
	defaultErrorBackoff = 1 * time.Second
)

// UpdatesClient fetches updates from Telegram.
type UpdatesClient interface {
	GetUpdates(
		ctx context.Context,
		offset int,
		timeout time.Duration,
	) ([]Update, error)
}

// MessageHandler handles one inbound Telegram message.
type MessageHandler func(ctx context.Context, msg Message) error

// MyChatMemberHandler handles one my_chat_member update.
type MyChatMemberHandler func(ctx context.Context, ev ChatMemberEvent) error

// CallbackQueryHandler handles one callback_query update.
type CallbackQueryHandler func(ctx context.Context, q CallbackQuery) error

// Poller consumes updates via getUpdates and calls the handler for
// each message with user content.
type Poller struct {
	client          UpdatesClient
	timeout         time.Duration
	backoff         time.Duration
	startFromLatest bool
	offsetStore     OffsetStore
	onError         func(error)
	handler         MessageHandler
	callbackQuery   CallbackQueryHandler
	myChatMember    MyChatMemberHandler
}

// PollerOption configures a Poller.
type PollerOption func(*Poller)

// WithPollTimeout sets the long-poll timeout.
func WithPollTimeout(timeout time.Duration) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithErrorBackoff sets the backoff delay after polling/handler errors.
func WithErrorBackoff(backoff time.Duration) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithStartFromLatest controls whether the poller drains pending
// updates on startup.
func WithStartFromLatest(enabled bool) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithOffsetStore enables persisting polling offsets.
func WithOffsetStore(store OffsetStore) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithOnError registers a callback for non-fatal errors.
func WithOnError(onError func(error)) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithMessageHandler sets the message handler.
func WithMessageHandler(h MessageHandler) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithMyChatMemberHandler sets the my_chat_member handler.
func WithMyChatMemberHandler(h MyChatMemberHandler) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// WithCallbackQueryHandler sets the callback_query handler.
func WithCallbackQueryHandler(h CallbackQueryHandler) PollerOption {
	_ = "STUB: not implemented"
	return *new(PollerOption)
}

// NewPoller creates a poller.
func NewPoller(client UpdatesClient, opts ...PollerOption) (*Poller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run starts the polling loop and blocks until ctx is done.
func (p *Poller) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func hasUserContent(msg *Message) bool { _ = "STUB: not implemented"; return false }

func (p *Poller) bootstrapOffset(
	ctx context.Context,
	offset int,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Poller) persistOffset(ctx context.Context, offset int) { _ = "STUB: not implemented"; return }

func sleepWithContext(ctx context.Context, d time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}
