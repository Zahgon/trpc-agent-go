//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysql

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-agent-go/internal/session/sqldb"
)

// SQL templates for table creation (MySQL syntax)
const (
	sqlCreateSessionStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			state JSON DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	sqlCreateSessionEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			event JSON NOT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	sqlCreateSessionTrackEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			track VARCHAR(255) NOT NULL,
			event JSON NOT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	// Note: no created_at column because summaries are upsert (overwrite on duplicate key).
	sqlCreateSessionSummariesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			filter_key VARCHAR(255) NOT NULL DEFAULT '',
			summary JSON DEFAULT NULL,
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	sqlCreateAppStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			` + "`key`" + ` VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	sqlCreateUserStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			` + "`key`" + ` VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	// Index creation SQL (MySQL syntax)
	// Note: MySQL doesn't support IF NOT EXISTS for indexes until MySQL 8.0.13+
	// We'll handle duplicate index errors in the creation logic

	// session_states: unique index on (app_name, user_id, session_id, deleted_at)
	sqlCreateSessionStatesUniqueIndex = `
		CREATE UNIQUE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, deleted_at)`

	// session_states: TTL index on (expires_at)
	sqlCreateSessionStatesExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`

	// session_events: lookup index on (app_name, user_id, session_id, created_at)
	sqlCreateSessionEventsLookupIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, created_at)`

	// session_events: TTL index on (expires_at)
	sqlCreateSessionEventsExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`

	// session_track_events: lookup index on (app_name, user_id, session_id, created_at)
	sqlCreateSessionTracksIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, created_at)`

	// session_track_events: TTL index on (expires_at)
	sqlCreateSessionTracksExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`

	// session_summaries: TTL index on (expires_at)
	sqlCreateSessionSummariesExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`

	// app_states: unique index on (app_name, key, deleted_at)
	sqlCreateAppStatesUniqueIndex = `
		CREATE UNIQUE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, ` + "`key`" + `, deleted_at)`

	// app_states: TTL index on (expires_at)
	sqlCreateAppStatesExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`

	// user_states: unique index on (app_name, user_id, key, deleted_at)
	sqlCreateUserStatesUniqueIndex = `
		CREATE UNIQUE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, ` + "`key`" + `, deleted_at)`

	// user_states: TTL index on (expires_at)
	sqlCreateUserStatesExpiresIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at)`
)

// mysqlVarCharIndexPrefixLen is a safe prefix length for utf8mb4 indexes.
// InnoDB has a maximum index key length of 3072 bytes. For utf8mb4, each
// character can take up to 4 bytes. To avoid Error 1071 (Specified key was
// too long), we use 191 as the prefix length, which is the standard
// industry practice:
//   - 4 columns * 191 chars * 4 bytes/char = 3056 bytes < 3072 bytes limit
//   - Using 192 would be 4 * 192 * 4 = 3072 bytes, which is exactly on the
//     boundary and may cause issues in some MySQL versions.
const mysqlVarCharIndexPrefixLen = 191

// session_summaries: unique index on (app_name, user_id, session_id, filter_key).
// Note: This index does NOT include deleted_at because MySQL treats NULL != NULL,
// which would allow duplicate active records. To ensure uniqueness, we exclude
// deleted_at from the unique index. On subsequent writes, deleted_at is reset to
// NULL (via ON DUPLICATE KEY UPDATE), effectively "reviving" the record instead
// of preserving deleted historical versions.
//
// Note: We use prefix indexes to avoid Error 1071 (max key length is 3072 bytes).
var sqlCreateSessionSummariesUniqueIndex = fmt.Sprintf(
	`
		CREATE UNIQUE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name(%d), user_id(%d), session_id(%d), filter_key(%d))`,
	mysqlVarCharIndexPrefixLen,
	mysqlVarCharIndexPrefixLen,
	mysqlVarCharIndexPrefixLen,
	mysqlVarCharIndexPrefixLen,
)

// tableDefinition defines a table with its SQL template
type tableDefinition struct {
	name     string
	template string
}

// indexDefinition defines an index with its table, suffix and SQL template
type indexDefinition struct {
	table    string
	suffix   string
	template string
}

// tableColumn represents a table column definition for schema verification
type tableColumn struct {
	name     string
	dataType string
	nullable bool
}

// tableIndex represents a table index definition for schema verification
type tableIndex struct {
	table   string   // Base table name (without prefix) like "session_states"
	suffix  string   // Index suffix like "unique_active", "lookup", "expires"
	columns []string // Expected columns in the index
	unique  bool     // Whether this is a unique index
}

// tableSchema defines the expected schema for a table.
type tableSchema struct {
	columns []tableColumn
	indexes []tableIndex
}

// expectedSchema defines the expected schema for each table.
var expectedSchema = map[string]tableSchema{
	sqldb.TableNameSessionStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"user_id", "varchar", false},
			{"session_id", "varchar", false},
			{"state", "json", true},
			{"created_at", "timestamp", false},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "session_id", "deleted_at"}, true},
			{sqldb.TableNameSessionStates, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
	sqldb.TableNameSessionEvents: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"user_id", "varchar", false},
			{"session_id", "varchar", false},
			{"event", "json", false},
			{"created_at", "timestamp", false},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionEvents, sqldb.IndexSuffixLookup, []string{"app_name", "user_id", "session_id", "created_at"}, false},
			{sqldb.TableNameSessionEvents, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
	sqldb.TableNameSessionTrackEvents: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"user_id", "varchar", false},
			{"session_id", "varchar", false},
			{"track", "varchar", false},
			{"event", "json", false},
			{"created_at", "timestamp", false},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixLookup, []string{"app_name", "user_id", "session_id", "created_at"}, false},
			{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
	sqldb.TableNameSessionSummaries: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"user_id", "varchar", false},
			{"session_id", "varchar", false},
			{"filter_key", "varchar", false},
			{"summary", "json", true},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			// Unique index on business key only (no deleted_at) to prevent duplicate active records.
			{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "session_id", "filter_key"}, true},
			{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
	sqldb.TableNameAppStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"key", "varchar", false},
			{"value", "text", true},
			{"created_at", "timestamp", false},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameAppStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "key", "deleted_at"}, true},
			{sqldb.TableNameAppStates, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
	sqldb.TableNameUserStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "varchar", false},
			{"user_id", "varchar", false},
			{"key", "varchar", false},
			{"value", "text", true},
			{"created_at", "timestamp", false},
			{"updated_at", "timestamp", false},
			{"expires_at", "timestamp", true},
			{"deleted_at", "timestamp", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameUserStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "key", "deleted_at"}, true},
			{sqldb.TableNameUserStates, sqldb.IndexSuffixExpires, []string{"expires_at"}, false},
		},
	},
}

// tdsqlExpectedSchema extends expectedSchema with TDSQL-specific indexes.
var tdsqlExpectedSchema = func() map[string]tableSchema {
	s := make(map[string]tableSchema, len(expectedSchema))
	for k, v := range expectedSchema {
		s[k] = v
	}
	ss := s[sqldb.TableNameSessionStates]
	newIndexes := make([]tableIndex, len(ss.indexes)+1)
	copy(newIndexes, ss.indexes)
	newIndexes[len(ss.indexes)] = tableIndex{
		table:   sqldb.TableNameSessionStates,
		suffix:  "list",
		columns: []string{"app_name", "user_id", "updated_at"},
	}
	ss.indexes = newIndexes
	s[sqldb.TableNameSessionStates] = ss
	return s
}()

// Global table definitions
var tableDefs = []tableDefinition{
	{sqldb.TableNameSessionStates, sqlCreateSessionStatesTable},
	{sqldb.TableNameSessionEvents, sqlCreateSessionEventsTable},
	{sqldb.TableNameSessionTrackEvents, sqlCreateSessionTrackEventsTable},
	{sqldb.TableNameSessionSummaries, sqlCreateSessionSummariesTable},
	{sqldb.TableNameAppStates, sqlCreateAppStatesTable},
	{sqldb.TableNameUserStates, sqlCreateUserStatesTable},
}

// Global index definitions
var indexDefs = []indexDefinition{
	// Unique indexes
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixUniqueActive, sqlCreateSessionStatesUniqueIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixUniqueActive, sqlCreateSessionSummariesUniqueIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixUniqueActive, sqlCreateAppStatesUniqueIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixUniqueActive, sqlCreateUserStatesUniqueIndex},

	// Lookup indexes
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixLookup, sqlCreateSessionEventsLookupIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixLookup, sqlCreateSessionTracksIndex},

	// TTL indexes
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixExpires, sqlCreateSessionStatesExpiresIndex},
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixExpires, sqlCreateSessionEventsExpiresIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixExpires, sqlCreateSessionTracksExpiresIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixExpires, sqlCreateSessionSummariesExpiresIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixExpires, sqlCreateAppStatesExpiresIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixExpires, sqlCreateUserStatesExpiresIndex},
}

// TDSQL table templates: PK includes shardkey (user_id for session tables, broadcast for app_states).
const (
	tdsqlCreateSessionStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			state JSON DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL,
			PRIMARY KEY (id, user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=user_id`

	tdsqlCreateSessionEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			event JSON NOT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL,
			PRIMARY KEY (id, user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=user_id`

	tdsqlCreateSessionTrackEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			track VARCHAR(255) NOT NULL,
			event JSON NOT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL,
			PRIMARY KEY (id, user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=user_id`

	// TDSQL session_summaries: app_name/session_id/filter_key use VARCHAR(128)
	// so the UNIQUE KEY fits within InnoDB's 3072-byte limit at full length:
	// 128*4*3 + 255*4 = 1536 + 1020 = 2556 < 3072. No prefix index needed.
	tdsqlCreateSessionSummariesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(128) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(128) NOT NULL,
			filter_key VARCHAR(128) NOT NULL DEFAULT '',
			summary JSON DEFAULT NULL,
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL,
			PRIMARY KEY (id, user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=user_id`

	tdsqlCreateAppStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			` + "`key`" + ` VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=noshardkey_allset`

	tdsqlCreateUserStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			` + "`key`" + ` VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			expires_at TIMESTAMP(6) NULL DEFAULT NULL,
			deleted_at TIMESTAMP(6) NULL DEFAULT NULL,
			PRIMARY KEY (id, user_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci shardkey=user_id`

	tdsqlCreateSessionStatesListIndex = `
		CREATE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, updated_at)`

	// TDSQL session_summaries uses VARCHAR(128) for app_name/session_id/filter_key,
	// so UNIQUE KEY fits at full column length without prefix indexes.
	tdsqlCreateSessionSummariesUniqueIndex = `
		CREATE UNIQUE INDEX {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, filter_key)`
)

var tdsqlTableDefs = []tableDefinition{
	{sqldb.TableNameSessionStates, tdsqlCreateSessionStatesTable},
	{sqldb.TableNameSessionEvents, tdsqlCreateSessionEventsTable},
	{sqldb.TableNameSessionTrackEvents, tdsqlCreateSessionTrackEventsTable},
	{sqldb.TableNameSessionSummaries, tdsqlCreateSessionSummariesTable},
	{sqldb.TableNameAppStates, tdsqlCreateAppStatesTable},
	{sqldb.TableNameUserStates, tdsqlCreateUserStatesTable},
}

var tdsqlIndexDefs = []indexDefinition{
	// Unique indexes (same as MySQL, shardkey already in UNIQUE KEYs)
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixUniqueActive, sqlCreateSessionStatesUniqueIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixUniqueActive, tdsqlCreateSessionSummariesUniqueIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixUniqueActive, sqlCreateAppStatesUniqueIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixUniqueActive, sqlCreateUserStatesUniqueIndex},

	// Lookup indexes (same as MySQL)
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixLookup, sqlCreateSessionEventsLookupIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixLookup, sqlCreateSessionTracksIndex},

	// TDSQL-specific: ListSessions sort index
	{sqldb.TableNameSessionStates, "list", tdsqlCreateSessionStatesListIndex},

	// TTL indexes (same as MySQL)
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixExpires, sqlCreateSessionStatesExpiresIndex},
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixExpires, sqlCreateSessionEventsExpiresIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixExpires, sqlCreateSessionTracksExpiresIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixExpires, sqlCreateSessionSummariesExpiresIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixExpires, sqlCreateAppStatesExpiresIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixExpires, sqlCreateUserStatesExpiresIndex},
}

// initDB initializes the database schema.
func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Select table and index definitions based on TDSQL mode.

// Create tables

// Create indexes

// MySQL doesn't have "IF NOT EXISTS" for indexes in older versions
// We'll use a different approach: try to create and ignore duplicate key errors

// Check if it's a duplicate index name error (error code 1061).
// This means the index already exists, which is safe to skip.

// Index already exists, log and continue.

// Verify schema

// verifySchema verifies that the database schema matches expectations.
func (s *Service) verifySchema(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Check if table exists

// Verify columns

// Verify indexes (non-fatal, just log warnings)

// tableExists checks if a table exists in the database.
func (s *Service) tableExists(ctx context.Context, tableName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// verifyColumns verifies that table columns match expectations.
func (s *Service) verifyColumns(ctx context.Context, tableName string, expectedColumns []tableColumn) error {
	_ = "STUB: not implemented"
	// Get actual columns from database
	return nil
}

// Check each expected column

// Check data type

// Check nullable

// verifyIndexes verifies that table indexes exist.
func (s *Service) verifyIndexes(ctx context.Context, fullTableName string, expectedIndexes []tableIndex) error {
	_ = "STUB: not implemented"
	// Build map of expected index names
	return nil
}

// Get actual indexes from database

// Check each expected index

// Build CREATE INDEX statement for user reference.

// Build DROP and CREATE INDEX statements for user reference.

// Check for extra/unexpected indexes

func stringSlicesEqual(a, b []string) bool { _ = "STUB: not implemented"; return false }

// buildIndexColumnsStr builds a comma-separated column list with appropriate
// prefix lengths for indexes that require them.
func buildIndexColumnsStr(table, suffix string, columns []string, tdsqlSharding bool) string {
	_ = "STUB: not implemented"
	// MySQL mode: session_summaries unique_active index requires prefix lengths
	// to avoid Error 1071. TDSQL mode uses VARCHAR(128) so no prefix needed.
	return ""
}

// For all other indexes, use columns as-is.

// buildCreateIndexSQL builds a CREATE INDEX SQL statement.
func buildCreateIndexSQL(indexName, tableName, columns string, unique bool) string {
	_ = "STUB: not implemented"
	return ""
}

// isDuplicateIndexNameError checks if the error is a MySQL duplicate index name error (1061).
// This is used when creating indexes - if the index name already exists, we can safely skip.
// Note: This should NOT match error 1062 (duplicate entry), which indicates a data constraint
// violation and should not be silently ignored.
func isDuplicateIndexNameError(err error) bool { _ = "STUB: not implemented"; return false }
