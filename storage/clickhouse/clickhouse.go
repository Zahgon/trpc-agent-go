//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package clickhouse provides the ClickHouse instance info management.
package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// BatchFn is the processing function after the user initiates the batch insert request, and it is required.
// The batch insert call can be simplified, it is aborted when error != nil,
// otherwise the batch insert is automatically sent.
type BatchFn func(batch driver.Batch) error

// defaultClientBuilder is the default ClickHouse client builder.
// It creates a native ClickHouse client using the official Go driver.
func defaultClientBuilder(builderOpts ...ClientBuilderOpt) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Client defines the interface for ClickHouse operations.
// This is a subset of the native ClickHouse driver interface,
// containing only the methods needed by the session layer.
type Client interface {
	// Exec executes the ClickHouse insert/delete/update command.
	Exec(ctx context.Context, query string, args ...any) error

	// Query executes the ClickHouse select command and returns driver.Rows.
	Query(ctx context.Context, query string, args ...any) (driver.Rows, error)

	// QueryRow executes the ClickHouse QueryRow command.
	QueryRow(ctx context.Context, dest []any, query string, args ...any) error

	// QueryToStruct executes the ClickHouse select command and scans the results into the dest struct.
	QueryToStruct(ctx context.Context, dest any, query string, args ...any) error

	// QueryToStructs executes the ClickHouse select command and scans the results into the dest slice.
	QueryToStructs(ctx context.Context, dest any, query string, args ...any) error

	// BatchInsert executes batch insert with callback function.
	// fn is a callback function that receives driver.Batch as a parameter.
	// When error != nil is returned in fn, the batch insert is aborted, otherwise automatically sent.
	BatchInsert(ctx context.Context, query string, fn BatchFn, opts ...driver.PrepareBatchOption) error

	// AsyncInsert executes asynchronous insert.
	// https://clickhouse.com/docs/en/optimize/asynchronous-inserts#enabling-asynchronous-inserts
	AsyncInsert(ctx context.Context, query string, wait bool, args ...any) error

	// Close closes the ClickHouse connection.
	Close() error
}

// defaultClient wraps driver.Conn to implement the Client interface.
type defaultClient struct {
	conn driver.Conn
}

// newDefaultClient creates a new defaultClient with the given driver.Conn.
func newDefaultClient(conn driver.Conn) *defaultClient { _ = "STUB: not implemented"; return nil }

// Exec implements Client.Exec.
func (c *defaultClient) Exec(ctx context.Context, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Query implements Client.Query.
func (c *defaultClient) Query(ctx context.Context, query string, args ...any) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

// QueryRow implements Client.QueryRow.
func (c *defaultClient) QueryRow(ctx context.Context, dest []any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// QueryToStruct implements Client.QueryToStruct.
func (c *defaultClient) QueryToStruct(ctx context.Context, dest any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// QueryToStructs implements Client.QueryToStructs.
func (c *defaultClient) QueryToStructs(ctx context.Context, dest any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// BatchInsert implements Client.BatchInsert.
func (c *defaultClient) BatchInsert(ctx context.Context, query string, fn BatchFn, opts ...driver.PrepareBatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

// AsyncInsert implements Client.AsyncInsert.
func (c *defaultClient) AsyncInsert(ctx context.Context, query string, wait bool, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Close implements Client.Close.
func (c *defaultClient) Close() error { _ = "STUB: not implemented"; return nil }
