//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package telegram

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwproto"
)

const (
	ffmpegBinaryName = "ffmpeg"

	convertedAudioPrefix = "openclaw-telegram-audio-"
	convertedAudioName   = "converted.wav"
	convertedAudioRate   = "16000"

	convertedAudioFileMode = 0o600
)

type audioInputSource struct {
	Name     string
	Path     string
	MimeType string
	Data     []byte
}

type convertedAudio struct {
	Data   []byte
	Format string
}

type audioInputConverter func(
	ctx context.Context,
	src audioInputSource,
) (*convertedAudio, error)

func defaultAudioInputConverter(
	ctx context.Context,
	src audioInputSource,
) (*convertedAudio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func looksConvertibleAudio(
	name string,
	filePath string,
	mimeType string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func convertAudioInputWithFFmpeg(
	ctx context.Context,
	src audioInputSource,
) (*convertedAudio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel) audioModelPart(
	ctx context.Context,
	name string,
	filePath string,
	mimeType string,
	data []byte,
) *gwproto.ContentPart {
	_ = "STUB: not implemented"
	return nil
}
