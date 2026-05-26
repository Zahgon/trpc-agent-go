//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package sqlitevec

import (
	"context"
)

const (
	sqlCheckVecVersion = `SELECT vec_version();`
)

// initDB checks sqlite-vec availability and creates the schema.
func (s *Store) initDB(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Check that sqlite-vec extension is available.
	return nil
}

// Create the main vec0 table.

// Create the metadata index table and its indices.

// buildCreateVecTableSQL returns the CREATE VIRTUAL TABLE statement for the
// vec0 main table.
func (s *Store) buildCreateVecTableSQL() string { _ = "STUB: not implemented"; return "" }

// buildCreateMetadataTableSQL returns the SQL statements for creating the
// metadata index table and its indices.
func (s *Store) buildCreateMetadataTableSQL() []string { _ = "STUB: not implemented"; return nil }
