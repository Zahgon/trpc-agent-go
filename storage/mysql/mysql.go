//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package mysql provides the mysql instance info management and client interface.
package mysql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlRegistry = make(map[string][]ClientBuilderOpt)
}

var mysqlRegistry map[string][]ClientBuilderOpt

// Client defines the interface for database operations using callback pattern.
// This interface abstracts the common database operations needed by the
// memory service, making it easier to inject mock implementations for testing.
type Client interface {
	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)

	// Query executes a query that returns rows, calling the next function for each row.
	Query(ctx context.Context, next NextFunc, query string, args ...any) error

	// QueryRow executes a query that is expected to return at most one row and scans into dest.
	QueryRow(ctx context.Context, dest []any, query string, args ...any) error

	// Transaction executes a function within a transaction.
	Transaction(ctx context.Context, fn TxFunc, opts ...TxOption) error

	// Close closes the database connection.
	Close() error
}

// NextFunc is called for each row in a query result.
// Return ErrBreak to stop iteration early, or any other error to abort with error.
type NextFunc func(*sql.Rows) error

// TxFunc is a user transaction function.
// Return nil to commit, or any error to rollback.
type TxFunc func(*sql.Tx) error

// TxOption configures transaction options.
type TxOption func(*sql.TxOptions)

// ErrBreak can be returned from NextFunc to stop iteration early without error.
var ErrBreak = errors.New("mysql scan rows break")

// sqlDBClient wraps *sql.DB to implement the Client interface using callback pattern.
type sqlDBClient struct {
	db *sql.DB
}

// WrapSQLDB wraps a *sql.DB connection into a Client.
//
// WARNING: This function is for INTERNAL USE ONLY!
// Do NOT call this function directly from external packages.
// This is an internal implementation detail that may change without notice.
// Use the public API provided by the parent storage/mysql package instead.
//
// This function is only exported to allow access from other internal packages
// within the same module (memory/mysql, etc.).
func WrapSQLDB(db *sql.DB) Client { _ = "STUB: not implemented"; return *new(Client) }

// Exec implements Client.Exec.
func (c *sqlDBClient) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query implements Client.Query using callback pattern.
func (c *sqlDBClient) Query(ctx context.Context, next NextFunc, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterate all rows, calling the callback function.

// QueryRow implements Client.QueryRow.
func (c *sqlDBClient) QueryRow(ctx context.Context, dest []any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Transaction implements Client.Transaction using callback pattern.
func (c *sqlDBClient) Transaction(ctx context.Context, fn TxFunc, opts ...TxOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Execute user transaction function.

// Close implements Client.Close.
func (c *sqlDBClient) Close() error { _ = "STUB: not implemented"; return nil }

// clientBuilder is the function type for building Client instances.
type clientBuilder func(builderOpts ...ClientBuilderOpt) (Client, error)

var globalBuilder clientBuilder = defaultClientBuilder

// SetClientBuilder sets the mysql client builder.
func SetClientBuilder(builder clientBuilder) { _ = "STUB: not implemented"; return }

// GetClientBuilder gets the mysql client builder.
func GetClientBuilder() clientBuilder {
	_ = "STUB: not implemented"
	return *

	// defaultClientBuilder is the default mysql client builder.
	new(clientBuilder)
}

func defaultClientBuilder(builderOpts ...ClientBuilderOpt) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Set connection pool settings if provided.

// Test connection.

// ClientBuilderOpt is the option for the mysql client.
type ClientBuilderOpt func(*ClientBuilderOpts)

// ClientBuilderOpts is the options for the mysql client.
type ClientBuilderOpts struct {
	// DSN is the mysql data source name for clientBuilder.
	// Format: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// Example: user:password@tcp(localhost:3306)/dbname?parseTime=true
	DSN string

	// MaxOpenConns is the maximum number of open connections to the database.
	MaxOpenConns int

	// MaxIdleConns is the maximum number of connections in the idle connection pool.
	MaxIdleConns int

	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	ConnMaxLifetime time.Duration

	// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
	ConnMaxIdleTime time.Duration

	// ExtraOptions is the extra options for the mysql client.
	ExtraOptions []any
}

// WithClientBuilderDSN sets the mysql client DSN for clientBuilder.
func WithClientBuilderDSN(dsn string) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithMaxOpenConns sets the maximum number of open connections to the database.
func WithMaxOpenConns(n int) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithMaxIdleConns sets the maximum number of connections in the idle connection pool.
func WithMaxIdleConns(n int) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithConnMaxLifetime sets the maximum amount of time a connection may be reused.
func WithConnMaxLifetime(d time.Duration) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithConnMaxIdleTime sets the maximum amount of time a connection may be idle.
func WithConnMaxIdleTime(d time.Duration) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithExtraOptions sets the mysql client extra options for clientBuilder.
// this option mainly used for the customized mysql client builder, it will be passed to the builder.
func WithExtraOptions(extraOptions ...any) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// RegisterMySQLInstance registers a mysql instance options.
func RegisterMySQLInstance(name string, opts ...ClientBuilderOpt) {
	_ = "STUB: not implemented"
	return
}

// GetMySQLInstance gets the mysql instance options.
func GetMySQLInstance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
