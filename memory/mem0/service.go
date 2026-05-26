//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mem0

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Service provides an ingest-first integration with mem0.
type Service struct {
	opts serviceOpts
	c    *client

	ingestWorker *ingestWorker

	precomputedTools []tool.Tool
}

// NewService creates a new mem0 service.
func NewService(options ...ServiceOpt) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tools returns the mem0 read-only tools exposed to the agent.
func (s *Service) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// IngestSession enqueues session transcript ingestion into mem0.
//
// Per-request settings are configured via session.IngestOption helpers
// (e.g. session.WithIngestMetadata, session.WithIngestAgentID,
// session.WithIngestRunID). The resolved snapshot is forwarded to mem0's
// POST /v1/memories/ payload as metadata, agent_id and run_id respectively.
//
// An invalid session scope (empty AppName / UserID) is surfaced as an error
// rather than silently dropped, so caller misconfiguration is distinguishable
// from a successful no-op, matching ReadMemories / SearchMemories behaviour.
func (s *Service) IngestSession(
	ctx context.Context,
	sess *session.Session,
	opts ...session.IngestOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadMemories reads memories for a user.
func (s *Service) ReadMemories(ctx context.Context, userKey memory.UserKey, limit int) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchMemories searches memories for a user.
func (s *Service) SearchMemories(ctx context.Context, userKey memory.UserKey, query string, opts ...memory.SearchOption) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close stops background workers and releases resources.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }
