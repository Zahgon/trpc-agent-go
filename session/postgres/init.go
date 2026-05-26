//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package postgres

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/internal/session/sqldb"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/postgres"
)

// SQL templates for table creation
const (
	sqlCreateSessionStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			state JSONB DEFAULT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	sqlCreateSessionEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			event JSONB NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	sqlCreateSessionTrackEventsTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			track VARCHAR(255) NOT NULL,
			event JSONB NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	sqlCreateSessionSummariesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			session_id VARCHAR(255) NOT NULL,
			filter_key VARCHAR(255) NOT NULL DEFAULT '',
			summary JSONB DEFAULT NULL,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	sqlCreateAppStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			key VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	sqlCreateUserStatesTable = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGSERIAL PRIMARY KEY,
			app_name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			key VARCHAR(255) NOT NULL,
			value TEXT DEFAULT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP DEFAULT NULL,
			deleted_at TIMESTAMP DEFAULT NULL
		)`

	// Index creation SQL
	// session_states: partial unique index on (app_name, user_id, session_id) - only for non-deleted records
	sqlCreateSessionStatesUniqueIndex = `
		CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id)
		WHERE deleted_at IS NULL`

	// session_states: TTL index on (expires_at) - partial index for non-null values
	sqlCreateSessionStatesExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`

	// session_events: lookup index on (app_name, user_id, session_id, created_at)
	sqlCreateSessionEventsIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, created_at)`

	// session_events: TTL index on (expires_at) - partial index for non-null values
	sqlCreateSessionEventsExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`

	// session_track_events: lookup index on (app_name, user_id, session_id, track, created_at).
	sqlCreateSessionTracksIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, track, created_at)`

	// session_track_events: TTL index on (expires_at).
	sqlCreateSessionTracksExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`

	// session_summaries: partial unique index on (app_name, user_id, session_id, filter_key) - only for non-deleted records
	sqlCreateSessionSummariesUniqueIndex = `
		CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, session_id, filter_key)
		WHERE deleted_at IS NULL`

	// session_summaries: TTL index on (expires_at) - partial index for non-null values
	sqlCreateSessionSummariesExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`

	// app_states: partial unique index on (app_name, key) - only for non-deleted records
	sqlCreateAppStatesUniqueIndex = `
		CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, key)
		WHERE deleted_at IS NULL`

	// app_states: TTL index on (expires_at) - partial index for non-null values
	sqlCreateAppStatesExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`

	// user_states: partial unique index on (app_name, user_id, key) - only for non-deleted records
	sqlCreateUserStatesUniqueIndex = `
		CREATE UNIQUE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(app_name, user_id, key)
		WHERE deleted_at IS NULL`

	// user_states: TTL index on (expires_at) - partial index for non-null values
	sqlCreateUserStatesExpiresIndex = `
		CREATE INDEX IF NOT EXISTS {{INDEX_NAME}}
		ON {{TABLE_NAME}}(expires_at) WHERE expires_at IS NOT NULL`
)

// tableColumn represents a table column definition
type tableColumn struct {
	name     string
	dataType string
	nullable bool
}

// tableIndex represents a table index definition
type tableIndex struct {
	table   string // Base table name (without prefix/schema) like "session_states"
	suffix  string // Index suffix like "unique_active", "lookup", "expires"
	columns []string
}

// expectedSchema defines the expected schema for each table
var expectedSchema = map[string]struct {
	columns []tableColumn
	indexes []tableIndex
}{
	sqldb.TableNameSessionStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"user_id", "character varying", false},
			{"session_id", "character varying", false},
			{"state", "jsonb", true},
			{"created_at", "timestamp without time zone", false},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "session_id"}},
			{sqldb.TableNameSessionStates, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
	sqldb.TableNameSessionEvents: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"user_id", "character varying", false},
			{"session_id", "character varying", false},
			{"event", "jsonb", false},
			{"created_at", "timestamp without time zone", false},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionEvents, sqldb.IndexSuffixLookup, []string{"app_name", "user_id", "session_id", "created_at"}},
			{sqldb.TableNameSessionEvents, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
	sqldb.TableNameSessionTrackEvents: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"user_id", "character varying", false},
			{"session_id", "character varying", false},
			{"track", "character varying", false},
			{"event", "jsonb", false},
			{"created_at", "timestamp without time zone", false},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixLookup, []string{"app_name", "user_id", "session_id", "track", "created_at"}},
			{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
	sqldb.TableNameSessionSummaries: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"user_id", "character varying", false},
			{"session_id", "character varying", false},
			{"filter_key", "character varying", false},
			{"summary", "jsonb", true},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "session_id", "filter_key"}},
			{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
	sqldb.TableNameAppStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"key", "character varying", false},
			{"value", "text", true},
			{"created_at", "timestamp without time zone", false},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameAppStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "key"}},
			{sqldb.TableNameAppStates, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
	sqldb.TableNameUserStates: {
		columns: []tableColumn{
			{"id", "bigint", false},
			{"app_name", "character varying", false},
			{"user_id", "character varying", false},
			{"key", "character varying", false},
			{"value", "text", true},
			{"created_at", "timestamp without time zone", false},
			{"updated_at", "timestamp without time zone", false},
			{"expires_at", "timestamp without time zone", true},
			{"deleted_at", "timestamp without time zone", true},
		},
		indexes: []tableIndex{
			{sqldb.TableNameUserStates, sqldb.IndexSuffixUniqueActive, []string{"app_name", "user_id", "key"}},
			{sqldb.TableNameUserStates, sqldb.IndexSuffixExpires, []string{"expires_at"}},
		},
	},
}

// indexDefinition defines an index with its table, suffix and SQL template
type indexDefinition struct {
	table    string
	suffix   string
	template string
}

// tableDefinition defines a table with its SQL template
type tableDefinition struct {
	name     string
	template string
}

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
	// Partial unique indexes (only for non-deleted records)
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixUniqueActive, sqlCreateSessionStatesUniqueIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixUniqueActive, sqlCreateSessionSummariesUniqueIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixUniqueActive, sqlCreateAppStatesUniqueIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixUniqueActive, sqlCreateUserStatesUniqueIndex},
	// Lookup indexes (only session_events needs a separate lookup index)
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixLookup, sqlCreateSessionEventsIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixLookup, sqlCreateSessionTracksIndex},
	// TTL indexes
	{sqldb.TableNameSessionStates, sqldb.IndexSuffixExpires, sqlCreateSessionStatesExpiresIndex},
	{sqldb.TableNameSessionEvents, sqldb.IndexSuffixExpires, sqlCreateSessionEventsExpiresIndex},
	{sqldb.TableNameSessionTrackEvents, sqldb.IndexSuffixExpires, sqlCreateSessionTracksExpiresIndex},
	{sqldb.TableNameSessionSummaries, sqldb.IndexSuffixExpires, sqlCreateSessionSummariesExpiresIndex},
	{sqldb.TableNameAppStates, sqldb.IndexSuffixExpires, sqlCreateAppStatesExpiresIndex},
	{sqldb.TableNameUserStates, sqldb.IndexSuffixExpires, sqlCreateUserStatesExpiresIndex},
}

// buildCreateTableSQL builds CREATE TABLE SQL with table prefix.
func buildCreateTableSQL(schema, prefix, tableName, template string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildCreateIndexSQL builds CREATE INDEX SQL with table and index names.
func buildCreateIndexSQL(schema, prefix, tableName, suffix, template string) string {
	_ = "STUB: not implemented"
	return ""
}

// parseTableName parses a full table name into schema and table components.
// Examples:
// - "session_states" -> ("public", "session_states")
// - "myschema.session_states" -> ("myschema", "session_states")
func parseTableName(fullTableName string) (schema, tableName string) {
	_ = "STUB: not implemented"
	return "", ""
}

// initDB initializes the database schema
func (s *Service) initDB(ctx context.Context) {
	_ = "STUB: not implemented"
	// Create tables
	return
}

// Create indexes

// Verify schema

// createTables creates all required tables with the given prefix.
// This function can be used by both Service and standalone InitDB.
func createTables(ctx context.Context, client storage.Client, schema, prefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// createIndexes creates all required indexes with the given prefix.
// This function can be used by both Service and standalone InitDB.
func createIndexes(ctx context.Context, client storage.Client, schema, prefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// verifySchema verifies that the database schema matches expectations
func (s *Service) verifySchema(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Check if table exists

// Verify columns

// Verify indexes

// tableExists checks if a table exists
func (s *Service) tableExists(ctx context.Context, fullTableName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// verifyColumns verifies that table columns match expectations
func (s *Service) verifyColumns(ctx context.Context, fullTableName string, expectedColumns []tableColumn) error {
	_ = "STUB: not implemented"
	return nil
}

// Get actual columns from database

// Check each expected column

// Check data type

// Check nullable

// verifyIndexes verifies that table indexes exist
func (s *Service) verifyIndexes(ctx context.Context, fullTableName string, expectedIndexes []tableIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Get actual indexes from database

// Check each expected index

// Use sqldb.BuildIndexNameWithSchema to construct the expected index.

// InitDBConfig contains configuration for standalone database initialization.
type InitDBConfig struct {
	host         string
	port         int
	user         string
	password     string
	database     string
	sslMode      string
	tablePrefix  string
	schema       string
	instanceName string
	extraOptions []any
}

// InitDBOpt is the option for InitDB.
type InitDBOpt func(*InitDBConfig)

// WithInitDBHost sets the PostgreSQL host.
func WithInitDBHost(host string) InitDBOpt { _ = "STUB: not implemented"; return *new(InitDBOpt) }

// WithInitDBPort sets the PostgreSQL port.
func WithInitDBPort(port int) InitDBOpt { _ = "STUB: not implemented"; return *new(InitDBOpt) }

// WithInitDBUser sets the database user.
func WithInitDBUser(user string) InitDBOpt { _ = "STUB: not implemented"; return *new(InitDBOpt) }

// WithInitDBPassword sets the database password.
func WithInitDBPassword(password string) InitDBOpt {
	_ = "STUB: not implemented"
	return *new(InitDBOpt)
}

// WithInitDBDatabase sets the database name.
func WithInitDBDatabase(database string) InitDBOpt {
	_ = "STUB: not implemented"
	return *new(InitDBOpt)
}

// WithInitDBSSLMode sets the SSL mode.
func WithInitDBSSLMode(sslMode string) InitDBOpt { _ = "STUB: not implemented"; return *new(InitDBOpt) }

// WithInitDBTablePrefix sets the table name prefix.
// Note: An underscore will be automatically added if not present.
// "trpc" and "trpc_" both result in "trpc_" prefix.
//
// Security: Uses internal/session/sqldb.ValidateTablePrefix to prevent SQL injection.
func WithInitDBTablePrefix(prefix string) InitDBOpt {
	_ = "STUB: not implemented"
	return *new(InitDBOpt)
}

// Use internal/session/sqldb validation

// Automatically add underscore if not present

// WithInitDBSchema sets the PostgreSQL schema name where tables will be created.
// Note: The schema must already exist in the database before calling InitDB.
// Security: Uses internal/session/sqldb.ValidateTableName to prevent SQL injection.
func WithInitDBSchema(schema string) InitDBOpt { _ = "STUB: not implemented"; return *new(InitDBOpt) }

// Use internal/session/sqldb validation

// WithInitDBInstanceName uses a postgres instance from storage.
// Note: Direct connection settings (WithInitDBHost, WithInitDBPort, etc.) have higher priority.
// If both are specified, direct connection settings will be used.
func WithInitDBInstanceName(instanceName string) InitDBOpt {
	_ = "STUB: not implemented"
	return *new(InitDBOpt)
}

// WithInitDBExtraOptions sets extra options for the postgres client builder.
// This option is mainly used for customized postgres client builders.
func WithInitDBExtraOptions(extraOptions ...any) InitDBOpt {
	_ = "STUB: not implemented"
	return *new(InitDBOpt)
}

// InitDB initializes the database schema with tables and indexes.
// This is a standalone function that can be used independently of the Service.
// It's useful for:
// - Manual database setup/migration
// - CI/CD pipelines
// - Initial deployment setup
//
// Note: You must import a PostgreSQL driver before calling this function:
//
//	import _ "github.com/lib/pq"
//
// Example usage:
//
//	err := postgres.InitDB(context.Background(),
//	    postgres.WithInitDBHost("localhost"),
//	    postgres.WithInitDBPort(5432),
//	    postgres.WithInitDBUser("admin"),
//	    postgres.WithInitDBPassword("secret"),
//	    postgres.WithInitDBDatabase("sessions"),
//	    postgres.WithInitDBSSLMode("disable"),
//	    postgres.WithInitDBTablePrefix("trpc_"),
//	)
//	if err != nil {
//	    panic(err)
//	}
//
// Or use registered instance:
//
//	err := postgres.InitDB(context.Background(),
//	    postgres.WithInitDBInstanceName("my-postgres"),
//	    postgres.WithInitDBTablePrefix("trpc_"),
//	)
func InitDB(ctx context.Context, opts ...InitDBOpt) error { _ = "STUB: not implemented"; return nil }

// Get postgres client builder

// Priority: direct connection settings > instance name
// If direct connection settings are provided, use them

// Otherwise, use instance name if provided

// Append extra options if provided

// Create tables using shared function

// Create indexes using shared function
