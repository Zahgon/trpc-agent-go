//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqlitevec

import (
	"context"
)

const (
	defaultTableName = "memories"
	schemaBackupName = "__schema_backup"

	sqlCheckVecVersion = `SELECT vec_version();`

	sqlCreateMemoriesTable = `
CREATE VIRTUAL TABLE IF NOT EXISTS {{TABLE_NAME}} USING vec0(
  memory_id text primary key,
  embedding float[{{DIMENSION}}] distance_metric=cosine,
  app_name text,
  user_id text,
  created_at integer,
  updated_at integer,
  deleted_at integer,
  +memory_content text,
  +topics text,
  +memory_kind text,
  +event_time integer,
  +participants text,
  +location text
);`

	sqlCreateSchemaBackupTable = `
CREATE TABLE %s (
  memory_id text,
  embedding blob,
  app_name text,
  user_id text,
  created_at integer,
  updated_at integer,
  deleted_at integer,
  memory_content text,
  topics text,
  memory_kind text,
  event_time integer,
  participants text,
  location text
);`
)

var requiredSchemaColumns = []string{
	"memory_id",
	"embedding",
	"app_name",
	"user_id",
	"created_at",
	"updated_at",
	"deleted_at",
	"memory_content",
	"topics",
	"memory_kind",
	"event_time",
	"participants",
	"location",
}

var legacySchemaColumns = []string{
	"memory_id",
	"embedding",
	"app_name",
	"user_id",
	"created_at",
	"updated_at",
	"deleted_at",
	"memory_content",
	"topics",
}

func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Service) ensureSchemaColumns(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) recoverSchemaBackup(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) migrateLegacySchema(
	ctx context.Context,
	found map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) createSchemaBackup(
	ctx context.Context,
	backupTable string,
	found map[string]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) restoreSchemaBackup(
	ctx context.Context,
	backupTable string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) tableExists(
	ctx context.Context,
	tableName string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Service) schemaMissingColumns(
	ctx context.Context,
	tableName string,
) ([]string, map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *Service) createMemoriesTableSQL() string { _ = "STUB: not implemented"; return "" }

func (s *Service) schemaBackupTableName() string { _ = "STUB: not implemented"; return "" }

func (s *Service) outdatedSchemaError(
	missing []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) dropTable(
	ctx context.Context,
	tableName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func hasSchemaColumns(
	found map[string]struct{},
	columns []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func missingColumns(
	found map[string]struct{},
	required []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func optionalColumnExpr(
	found map[string]struct{},
	column string,
) string {
	_ = "STUB: not implemented"
	return ""
}
