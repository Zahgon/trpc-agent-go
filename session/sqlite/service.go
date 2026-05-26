//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sqlite provides the sqlite session service.
package sqlite

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
	isummary "trpc.group/trpc-go/trpc-agent-go/session/internal/summary"
)

var _ session.Service = (*Service)(nil)
var _ session.TrackService = (*Service)(nil)

// SessionState is the state of a session.
type SessionState struct {
	ID        string           `json:"id"`
	State     session.StateMap `json:"state"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

// Service is the sqlite session service.
type Service struct {
	opts ServiceOpts
	db   *sql.DB

	stateWriteMu sync.Mutex

	eventPairChans  []chan *sessionEventPair
	trackEventChans []chan *trackEventPair
	persistWg       sync.WaitGroup

	asyncWorker *isummary.AsyncSummaryWorker

	cleanupTicker *time.Ticker
	cleanupDone   chan struct{}
	cleanupOnce   sync.Once

	once sync.Once

	tableSessionStates    string
	tableSessionEvents    string
	tableSessionTracks    string
	tableSessionSummaries string
	tableAppStates        string
	tableUserStates       string
}

type sessionEventPair struct {
	key   session.Key
	event *event.Event
}

type trackEventPair struct {
	key   session.Key
	event *session.TrackEvent
}

// NewService creates a new sqlite session service.
//
// The service owns the passed-in db and will close it in Close().
func NewService(db *sql.DB, options ...ServiceOpt) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the service and releases resources.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Service) fullTableName(base string) string { _ = "STUB: not implemented"; return "" }

func applyOptions(opts ...session.Option) *session.Options { _ = "STUB: not implemented"; return nil }

func calculateExpiresAt(now time.Time, ttl time.Duration) *int64 {
	_ = "STUB: not implemented"
	return nil
}

func unixNanoToTime(ns int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// CreateSession creates a new session.
func (s *Service) CreateSession(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
	opts ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) checkSessionExists(
	ctx context.Context,
	key session.Key,
) (bool, sql.NullInt64, error) {
	_ = "STUB: not implemented"
	return false, *new(sql.NullInt64), nil
}

func isExpired(expiresAt sql.NullInt64, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Service) upsertSessionState(
	ctx context.Context,
	key session.Key,
	stateBytes []byte,
	now time.Time,
	expiresAt *int64,
	exists bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSession gets a session.
func (s *Service) GetSession(
	ctx context.Context,
	key session.Key,
	opts ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListSessions lists sessions for a user.
func (s *Service) ListSessions(
	ctx context.Context,
	userKey session.UserKey,
	opts ...session.Option,
) ([]*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSession deletes a session.
func (s *Service) DeleteSession(
	ctx context.Context,
	key session.Key,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAppState updates the app state.
func (s *Service) UpdateAppState(
	ctx context.Context,
	appName string,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) upsertAppState(
	ctx context.Context,
	appName string,
	key string,
	value []byte,
	now time.Time,
	expiresAt *int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListAppStates lists all app states.
func (s *Service) ListAppStates(
	ctx context.Context,
	appName string,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// DeleteAppState deletes a single app state key.
func (s *Service) DeleteAppState(
	ctx context.Context,
	appName string,
	key string,
) error {
	_ = "STUB: not implemented"
	return nil
}
