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
	"database/sql"
)

// metadataValueType enumerates the typed columns in the metadata index table.
const (
	metadataValueTypeText = "text"
	metadataValueTypeNum  = "num"
	metadataValueTypeBool = "bool"
	metadataValueTypeJSON = "json"
)

// metadataRow represents a single row in the metadata index table.
type metadataRow struct {
	docID     string
	key       string
	ordinal   int
	valueType string
	valueText sql.NullString
	valueNum  sql.NullFloat64
	valueBool sql.NullInt64
	valueJSON sql.NullString
}

// insertMetadataRows inserts expanded metadata rows in the given transaction.
func (s *Store) insertMetadataRows(ctx context.Context, tx *sql.Tx, docID string, metadata map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteMetadataRows deletes all metadata rows for the given document id.
func (s *Store) deleteMetadataRows(ctx context.Context, tx *sql.Tx, docID string) error {
	_ = "STUB: not implemented"
	return nil
}

// loadStoredMetadata loads metadata for a document from the metadata index table.
// It reconstructs the original stored metadata shape using value_json and ordinal.
func (s *Store) loadStoredMetadata(ctx context.Context, docID string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// classifyMetadataValues determines the typed representation of a metadata value.
// Scalar values produce a single row with ordinal 0. Arrays/slices produce one
// row per element.
func classifyMetadataValues(docID, key string, value any) []metadataRow {
	_ = "STUB: not implemented"
	return nil
}

func classifyMetadataScalar(docID, key string, ordinal int, value any) metadataRow {
	_ = "STUB: not implemented"
	return *new(metadataRow)
}

// Serialize original JSON for round-trip.

// Complex types (slices, maps, etc.) are stored as JSON only.
