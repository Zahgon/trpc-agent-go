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
	defaultTableName = "memories"

	sqlCreateMemoriesTable = `
CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
  memory_id TEXT PRIMARY KEY,
  app_name TEXT NOT NULL,
  user_id TEXT NOT NULL,
  memory_data BLOB NOT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL,
  deleted_at INTEGER DEFAULT NULL
);`

	sqlCreateMemoriesAppUserIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(app_name, user_id);`

	sqlCreateMemoriesUpdatedAtIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(updated_at DESC);`

	sqlCreateMemoriesDeletedAtIndex = `
CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
ON {{TABLE_NAME}}(deleted_at)
WHERE deleted_at IS NOT NULL;`
)

const (
	indexSuffixAppUser   = "app_user"
	indexSuffixUpdatedAt = "updated_at"
	indexSuffixDeletedAt = "deleted_at"
)

func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
