//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package sqlitevec provides a sqlite-vec-backed implementation of the
// knowledge vector store.
package sqlitevec

const (
	defaultDriverName        = "sqlite3"
	defaultDSN               = ":memory:"
	defaultTableName         = "knowledge_documents"
	defaultMetadataTableName = "knowledge_document_meta"
	defaultIndexDimension    = 1536
	defaultMaxResults        = 10
)

type options struct {
	dsn               string
	driverName        string
	tableName         string
	metadataTableName string
	indexDimension    int
	maxResults        int
	skipDBInit        bool
}

var defaultOptions = options{
	dsn:               defaultDSN,
	driverName:        defaultDriverName,
	tableName:         defaultTableName,
	metadataTableName: defaultMetadataTableName,
	indexDimension:    defaultIndexDimension,
	maxResults:        defaultMaxResults,
}

// Option configures the sqlitevec vector store.
type Option func(*options)

// WithDSN sets the SQLite DSN used when opening a database internally.
// Common mattn/go-sqlite3 DSN examples:
//   - ":memory:" for an in-memory database
//   - "file:/tmp/knowledge.db?_busy_timeout=5000" for a local file
//   - "file::memory:?cache=shared" for a shared in-memory database
func WithDSN(dsn string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDriverName sets the SQL driver name used with WithDSN.
func WithDriverName(driverName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTableName sets the vec0 table name.
func WithTableName(tableName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataTableName sets the metadata index table name.
func WithMetadataTableName(tableName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIndexDimension sets the embedding dimension.
func WithIndexDimension(dimension int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxResults sets the default search result limit.
func WithMaxResults(maxResults int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipDBInit skips schema initialization.
func WithSkipDBInit(skip bool) Option { _ = "STUB: not implemented"; return *new(Option) }
