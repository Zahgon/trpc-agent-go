//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package gateway

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	whisperBinaryName = "whisper"

	audioTranscriptionTempPrefix = "openclaw-audio-transcription-"
	audioTranscriptionInputStem  = "audio"
	audioTranscriptionOutputExt  = ".txt"
	audioTranscriptionTask       = "transcribe"
	audioTranscriptionFP16False  = "False"

	audioTranscriptionFileMode = 0o600

	defaultAudioTranscriptionTimeout = 45 * time.Second
	minAudioTranscriptionBytes       = 1024
)

type audioTranscriber interface {
	Transcribe(
		ctx context.Context,
		audio *model.Audio,
	) (string, error)
}

type whisperCLITranscriber struct {
	bin string
}

func newDefaultAudioTranscriber() audioTranscriber {
	_ = "STUB: not implemented"
	return *new(audioTranscriber)
}

func (t *whisperCLITranscriber) Transcribe(
	ctx context.Context,
	audio *model.Audio,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func audioTranscriptionExt(format string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
