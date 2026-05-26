//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agent

import (
	"context"
	"errors"
	"sync"
)

const defaultStreamBufferSize = 256

const (
	errStreamNameEmpty        = "stream name is empty"
	errStreamWriterAlreadySet = "stream writer already opened"
	errStreamReaderAlreadySet = "stream reader already opened"
)

var (
	// ErrStreamNameEmpty indicates the stream name is empty.
	ErrStreamNameEmpty = errors.New(errStreamNameEmpty)
	// ErrStreamWriterAlreadySet indicates a writer is already opened.
	ErrStreamWriterAlreadySet = errors.New(errStreamWriterAlreadySet)
	// ErrStreamReaderAlreadySet indicates a reader is already opened.
	ErrStreamReaderAlreadySet = errors.New(errStreamReaderAlreadySet)
)

// StreamHub is an invocation-scoped registry for ephemeral streams.
//
// A StreamHub is designed for in-graph, node-to-node streaming consumption.
// It is not checkpointed and should not be used for durable data.
type StreamHub struct {
	mu      sync.Mutex
	streams map[string]*stream
}

func newStreamHub() *StreamHub { _ = "STUB: not implemented"; return nil }

// GetOrCreateStreamHub returns the invocation's StreamHub.
//
// The hub is stored in the invocation state and is intentionally preserved
// across invocation.Clone() calls so that different nodes in the same graph
// run can share streams.
func GetOrCreateStreamHub(inv *Invocation) *StreamHub { _ = "STUB: not implemented"; return nil }

// StreamHubFromContext returns the StreamHub stored on the invocation in ctx.
func StreamHubFromContext(ctx context.Context) (*StreamHub, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// OpenStreamWriter opens a stream writer from the invocation in ctx.
func OpenStreamWriter(
	ctx context.Context,
	streamName string,
) (*StreamWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenStreamReader opens a stream reader from the invocation in ctx.
func OpenStreamReader(
	ctx context.Context,
	streamName string,
) (*StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CloseAll closes all streams in the hub.
func (h *StreamHub) CloseAll(err error) { _ = "STUB: not implemented"; return }

// OpenWriter opens the named stream's writer.
//
// Only one writer may be opened per stream name.
func (h *StreamHub) OpenWriter(
	ctx context.Context,
	streamName string,
) (*StreamWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenReader opens the named stream's reader.
//
// Only one reader may be opened per stream name.
func (h *StreamHub) OpenReader(
	ctx context.Context,
	streamName string,
) (*StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *StreamHub) getOrCreate(streamName string) *stream { _ = "STUB: not implemented"; return nil }

type stream struct {
	name string

	ch   chan []byte
	done chan struct{}

	mu         sync.Mutex
	writerOpen bool
	readerOpen bool
	closeOnce  sync.Once
	closeErr   error
}

func newStream(name string, bufSize int) *stream { _ = "STUB: not implemented"; return nil }

func (s *stream) markWriterOpen() error { _ = "STUB: not implemented"; return nil }

func (s *stream) markReaderOpen() error { _ = "STUB: not implemented"; return nil }

func (s *stream) closeWithError(err error) { _ = "STUB: not implemented"; return }

func (s *stream) writerErr() error { _ = "STUB: not implemented"; return nil }

func (s *stream) readerErr() error { _ = "STUB: not implemented"; return nil }

func (s *stream) doneClosed() bool { _ = "STUB: not implemented"; return false }

// StreamWriter writes bytes into a named StreamHub stream.
//
// It is safe for concurrent use.
type StreamWriter struct {
	ctx context.Context
	s   *stream
}

func (w *StreamWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteString writes s into the stream.
func (w *StreamWriter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close closes the stream for writing.
func (w *StreamWriter) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError closes the stream for writing with err.
func (w *StreamWriter) CloseWithError(err error) error { _ = "STUB: not implemented"; return nil }

// StreamReader reads bytes from a named StreamHub stream.
//
// It is not safe for concurrent use.
type StreamReader struct {
	ctx context.Context
	s   *stream

	buf []byte
	off int
}

func (r *StreamReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close closes the reader and stops the writer.
func (r *StreamReader) Close() error { _ = "STUB: not implemented"; return nil }
