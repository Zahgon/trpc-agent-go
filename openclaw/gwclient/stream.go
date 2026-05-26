//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package gwclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
)

// StreamMessage sends one message to the streaming gateway handler.
func (c *Client) StreamMessage(
	ctx context.Context,
	req MessageRequest,
) (<-chan StreamEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StreamMessageWithOptions sends one streaming message with opt-in
// streaming behavior controls.
func (c *Client) StreamMessageWithOptions(
	ctx context.Context,
	req MessageRequest,
	opts *MessageStreamOptions,
) (<-chan StreamEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) streamMessage(
	ctx context.Context,
	req MessageRequest,
	opts *MessageStreamOptions,
) (<-chan StreamEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type streamMessageRequest struct {
	MessageRequest
	StreamOptions *MessageStreamOptions `json:"stream_options,omitempty"`
}

func streamStatusError(status int, body []byte) error { _ = "STUB: not implemented"; return nil }

func parseSSEStream(
	ctx context.Context,
	reader io.Reader,
	out chan<- StreamEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

type streamResponseRecorder struct {
	header http.Header

	mu       sync.Mutex
	code     int
	wroteHdr bool

	headerReady chan struct{}
	pipeReader  *io.PipeReader
	pipeWriter  *io.PipeWriter
	body        bytes.Buffer
}

func newStreamResponseRecorder() *streamResponseRecorder { _ = "STUB: not implemented"; return nil }

func (r *streamResponseRecorder) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (r *streamResponseRecorder) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (r *streamResponseRecorder) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *streamResponseRecorder) Flush() { _ = "STUB: not implemented"; return }

func (r *streamResponseRecorder) Code() int { _ = "STUB: not implemented"; return 0 }

func (r *streamResponseRecorder) BodyBytes() []byte { _ = "STUB: not implemented"; return nil }

func (r *streamResponseRecorder) reader() io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (r *streamResponseRecorder) closeReader() { _ = "STUB: not implemented"; return }

func (r *streamResponseRecorder) waitHeader(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *streamResponseRecorder) finish(err error) { _ = "STUB: not implemented"; return }
