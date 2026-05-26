//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqlite

import (
	"context"
)

const (
	sqlCreateSessionStatesTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  session_id TEXT NOT NULL,
  state BLOB DEFAULT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateSessionEventsTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  session_id TEXT NOT NULL,
  event BLOB NOT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateSessionTrackEventsTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  session_id TEXT NOT NULL,
  track TEXT NOT NULL,
  event BLOB NOT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateSessionSummariesTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  session_id TEXT NOT NULL,
  filter_key TEXT NOT NULL DEFAULT '',
  summary BLOB DEFAULT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateAppStatesTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  key TEXT NOT NULL,
  value BLOB DEFAULT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateUserStatesTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  key TEXT NOT NULL,
  value BLOB DEFAULT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  expires_at INTEGER DEFAULT NULL,
  deleted_at INTEGER DEFAULT NULL
);`
)

const (
	sqlCreateSessionStatesUniqueIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id, session_id)
WHERE deleted_at IS NULL;`

	sqlCreateSessionStatesExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`

	sqlCreateSessionEventsLookupIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id, session_id, created_at);`

	sqlCreateSessionEventsExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`

	sqlCreateSessionTracksLookupIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id, session_id, track, created_at);`

	sqlCreateSessionTracksExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`

	sqlCreateSessionSummariesUniqueIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id, session_id, filter_key);`

	sqlCreateSessionSummariesExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`

	sqlCreateAppStatesUniqueIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, key)
WHERE deleted_at IS NULL;`

	sqlCreateAppStatesExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`

	sqlCreateUserStatesUniqueIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id, key)
WHERE deleted_at IS NULL;`

	sqlCreateUserStatesExpiresIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(expires_at)
WHERE expires_at IS NOT NULL;`
)

type tableDefinition struct {
	name     string
	template string
}

type indexDefinition struct {
	table    string
	suffix   string
	template string
}

func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
