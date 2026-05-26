//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sse provides SSE service implementation.
package sse

import (
	"context"
	"io"
	"net/http"
	"time"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	aguisse "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/encoding/sse"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	aguirunner "trpc.group/trpc-go/trpc-agent-go/server/agui/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/service"
)

// sse is a SSE service implementation.
type sse struct {
	path                    string
	messagesSnapshotPath    string
	cancelPath              string
	writer                  *aguisse.SSEWriter
	runner                  aguirunner.Runner
	handler                 http.Handler
	messagesSnapshotEnabled bool
	cancelEnabled           bool
	heartbeatInterval       time.Duration
}

// New creates a new SSE service.
func New(runner aguirunner.Runner, opt ...service.Option) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

// Handler returns an http.Handler that exposes the AG-UI SSE endpoint.
func (s *sse) Handler() http.Handler {
	_ = "STUB: not implemented"

	// handle handles an AG-UI run request.
	return *new(http.Handler)
}

func (s *sse) handle(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// handleMessagesSnapshot streams a synthetic snapshot run to the client.
func (s *sse) handleMessagesSnapshot(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *sse) handleEvents(
	ctx context.Context,
	w http.ResponseWriter,
	events <-chan aguievents.Event,
	drain bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHeartbeat(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

// handleCancel cancels a running run identified by the request payload.
func (s *sse) handleCancel(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// runAgentInputFromReader parses an AG-UI run request payload from a reader.
func runAgentInputFromReader(r io.Reader) (*adapter.RunAgentInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drainEvents(events <-chan aguievents.Event) { _ = "STUB: not implemented"; return }
