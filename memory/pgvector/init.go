//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package pgvector

import (
	"context"
)

// SQL templates for table creation (PostgreSQL with pgvector syntax).
const (
	sqlCreateExtension    = "CREATE EXTENSION IF NOT EXISTS vector"
	sqlCheckDDLPrivilege  = "SELECT has_schema_privilege($1, 'CREATE')"
	sqlCreateTablePattern = "CREATE TABLE IF NOT EXISTS %s (" +
		"memory_id TEXT PRIMARY KEY," +
		"app_name TEXT NOT NULL," +
		"user_id TEXT NOT NULL," +
		"memory_content TEXT NOT NULL," +
		"topics TEXT[]," +
		"embedding vector(%d)," +
		"memory_kind TEXT NOT NULL DEFAULT 'fact'," +
		"event_time TIMESTAMP NULL," +
		"participants TEXT[]," +
		"location TEXT NULL," +
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"deleted_at TIMESTAMP NULL DEFAULT NULL" +
		")"

	sqlCreateAppUserIndexPattern      = "CREATE INDEX IF NOT EXISTS %s ON %s(app_name, user_id)"
	sqlCreateUpdatedAtIndexPattern    = "CREATE INDEX IF NOT EXISTS %s ON %s(updated_at DESC)"
	sqlCreateDeletedAtIndexPattern    = "CREATE INDEX IF NOT EXISTS %s ON %s(deleted_at)"
	sqlCreateEventTimeIndexPattern    = "CREATE INDEX IF NOT EXISTS %s ON %s(event_time DESC) WHERE event_time IS NOT NULL"
	sqlCreateKindIndexPattern         = "CREATE INDEX IF NOT EXISTS %s ON %s(app_name, user_id, memory_kind)"
	sqlCreateParticipantsIndexPattern = "CREATE INDEX IF NOT EXISTS %s ON %s USING gin(participants) WHERE participants IS NOT NULL"

	sqlAddSearchVectorColumn         = "ALTER TABLE %s ADD COLUMN IF NOT EXISTS search_vector tsvector"
	sqlCreateSearchVectorIndex       = "CREATE INDEX IF NOT EXISTS %s ON %s USING gin(search_vector)"
	sqlBackfillSearchVector          = "UPDATE %s SET search_vector = to_tsvector('english', coalesce(memory_content, '')) WHERE search_vector IS NULL"
	sqlCreateSearchVectorTriggerFunc = `CREATE OR REPLACE FUNCTION %s_search_vector_update() RETURNS trigger AS $$
BEGIN
  NEW.search_vector := to_tsvector('english', coalesce(NEW.memory_content, ''));
  RETURN NEW;
END
$$ LANGUAGE plpgsql`
	sqlAttachSearchVectorTrigger = "DROP TRIGGER IF EXISTS tsvector_update ON %s; " +
		"CREATE TRIGGER tsvector_update BEFORE INSERT OR UPDATE ON %s " +
		"FOR EACH ROW EXECUTE FUNCTION %s_search_vector_update()"

	sqlCreateHNSWIndexPattern = "CREATE INDEX IF NOT EXISTS %s ON %s USING hnsw " +
		"(embedding vector_cosine_ops) WITH (m = %d, ef_construction = %d)"
)

// buildFullTableName builds the full table name with optional schema prefix.
func buildFullTableName(schema, tableName string) string { _ = "STUB: not implemented"; return "" }

// buildIndexName builds the index name from table name and suffix.
func buildIndexName(tableName, suffix string) string { _ = "STUB: not implemented"; return "" }

// buildCreateTableSQL builds the CREATE TABLE SQL.
func buildCreateTableSQL(schema, tableName string, dimension int) string {
	_ = "STUB: not implemented"
	return ""
}

// buildCreateIndexSQL builds the CREATE INDEX SQL.
func buildCreateIndexSQL(schema, tableName, indexSuffix, template string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildCreateHNSWIndexSQL builds the CREATE HNSW INDEX SQL.
func buildCreateHNSWIndexSQL(schema, tableName string, params *HNSWIndexParams) string {
	_ = "STUB: not implemented"
	return ""
}

// checkDDLPrivilege checks if the current user has DDL (CREATE) privilege.
func (s *Service) checkDDLPrivilege(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If no rows returned, hasPrivilege remains false (default).

// initDB initializes the database schema.
func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Enable pgvector extension.

// Check DDL privilege before proceeding.

// Use base table name from opts (before schema prefix is applied).

// Create table.

// Migrate existing tables: add episodic columns if they don't exist.
// This is safe to run on both new and existing tables because
// ADD COLUMN IF NOT EXISTS is a no-op when the column already exists.

// Index suffix constants.

// Create regular indexes.

// Create HNSW vector index.

// Add search_vector column for full-text search (hybrid search support).

// Create trigger function to auto-populate search_vector on insert/update.

// Attach trigger to table.

// Create GIN index on search_vector for fast full-text search.

// Backfill search_vector for existing rows that lack it.
