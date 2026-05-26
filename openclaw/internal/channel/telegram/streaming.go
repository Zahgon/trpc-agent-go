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
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"

	tgapi "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/telegram"
)

const (
	streamingOff      = "off"
	streamingBlock    = "block"
	streamingProgress = "progress"

	defaultStreamingMode = streamingOff
)

const (
	chatActionTyping = "typing"

	processingMessage = "Processing..."

	progressInterval = 2 * time.Second

	progressEditIntervalFast     = progressInterval
	progressEditIntervalMedium   = 10 * time.Second
	progressEditIntervalSlow     = 30 * time.Second
	progressEditIntervalVerySlow = time.Minute

	progressEditAfterMedium   = time.Minute
	progressEditAfterSlow     = 10 * time.Minute
	progressEditAfterVerySlow = 30 * time.Minute
)

type telegramMessageSummary struct {
	ChatID          int64  `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id,omitempty"`
	ReplyTo         int    `json:"reply_to,omitempty"`
	IncomingReplyTo int    `json:"incoming_reply_to,omitempty"`
	FromID          string `json:"from_id,omitempty"`
	Thread          string `json:"thread,omitempty"`
	RequestID       string `json:"request_id,omitempty"`
	SessionID       string `json:"session_id,omitempty"`
	Text            string `json:"text,omitempty"`
	Caption         string `json:"caption,omitempty"`

	HasPhoto     bool `json:"has_photo,omitempty"`
	HasDocument  bool `json:"has_document,omitempty"`
	HasAudio     bool `json:"has_audio,omitempty"`
	HasVoice     bool `json:"has_voice,omitempty"`
	HasVideo     bool `json:"has_video,omitempty"`
	HasAnimation bool `json:"has_animation,omitempty"`
	HasVideoNote bool `json:"has_video_note,omitempty"`

	ReplyHasPhoto     bool `json:"reply_has_photo,omitempty"`
	ReplyHasDocument  bool `json:"reply_has_document,omitempty"`
	ReplyHasAudio     bool `json:"reply_has_audio,omitempty"`
	ReplyHasVoice     bool `json:"reply_has_voice,omitempty"`
	ReplyHasVideo     bool `json:"reply_has_video,omitempty"`
	ReplyHasAnimation bool `json:"reply_has_animation,omitempty"`
	ReplyHasVideoNote bool `json:"reply_has_video_note,omitempty"`
}

func parseStreamingMode(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Channel) callGatewayAndReply(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	fromID string,
	thread string,
	sessionID string,
	requestID string,
	msg tgapi.Message,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func markReplyFilesAsVoice(
	files []channel.OutboundFile,
) []channel.OutboundFile {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) filterAlreadySentReplyFiles(
	requestID string,
	chatID int64,
	threadID int,
	files []channel.OutboundFile,
) []channel.OutboundFile {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel) sendPreviewMessage(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	mode string,
) (tgapi.Message, bool) {
	_ = "STUB: not implemented"
	return *new(tgapi.Message), false
}

func (c *Channel) startProgressLoop(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	preview tgapi.Message,
	hasPreview bool,
	mode string,
) (context.CancelFunc, *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc), nil
}

func (c *Channel) progressLoop(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	messageID int,
) {
	_ = "STUB: not implemented"
	return
}

func progressEditInterval(elapsed time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Channel) editPreview(
	ctx context.Context,
	chatID int64,
	messageID int,
	text string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Channel) sendReplyParts(
	ctx context.Context,
	chatID int64,
	messageThreadID int,
	replyTo int,
	parts []string,
) {
	_ = "STUB: not implemented"
	return
}
