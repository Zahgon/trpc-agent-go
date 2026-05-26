//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package openai

import (
	"bufio"
	"io"

	"github.com/openai/openai-go/packages/ssestream"
)

func init() {
	ssestream.RegisterDecoder("text/event-stream", newTolerantEventStreamDecoder)
	ssestream.RegisterDecoder("text/event-stream; charset=utf-8", newTolerantEventStreamDecoder)
	ssestream.RegisterDecoder("text/event-stream;charset=utf-8", newTolerantEventStreamDecoder)
}

func newTolerantEventStreamDecoder(rc io.ReadCloser) ssestream.Decoder {
	_ = "STUB: not implemented"
	return *new(ssestream.Decoder)
}

// tolerantEventStreamDecoder implements SSE parsing for OpenAI-compatible
// streams while ignoring provider-specific keep-alive or progress events that
// do not carry JSON payloads. Some OpenAI-compatible vendors emit events such
// as "data: : keep-alive" or "data: : OPENROUTER PROCESSING"; those events are
// not chat chunks and should not terminate the stream.
type tolerantEventStreamDecoder struct {
	evt ssestream.Event
	rc  io.ReadCloser
	scn *bufio.Scanner
	err error
}

func (s *tolerantEventStreamDecoder) Next() bool { _ = "STUB: not implemented"; return false }

// Dispatch event on an empty line.

// Split a string like "event: bar" into name="event" and value=" bar".

// Consume an optional space after the colon if it exists.

// A line starting with ":" is an SSE comment.

func (s *tolerantEventStreamDecoder) Event() ssestream.Event {
	_ = "STUB: not implemented"
	return *new(ssestream.Event)
}

func (s *tolerantEventStreamDecoder) Close() error { _ = "STUB: not implemented"; return nil }

func (s *tolerantEventStreamDecoder) Err() error { _ = "STUB: not implemented"; return nil }

func shouldSkipSSEPayload(data []byte) bool { _ = "STUB: not implemented"; return false }
