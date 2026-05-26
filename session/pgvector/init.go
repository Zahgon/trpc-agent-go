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

	storage "trpc.group/trpc-go/trpc-agent-go/storage/postgres"
)

// initDB initializes the database schema including pgvector
// extension, tables, indexes, and HNSW vector index.
func (s *Service) initDB(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// HNSW index creation may fail if pgvector is not
// installed; log a warning instead of panic.

// enablePgvectorExtension enables the pgvector
// extension in PostgreSQL.
func enablePgvectorExtension(
	ctx context.Context,
	client storage.Client,
) error {
	_ = "STUB: not implemented"
	return nil
}

// createTables creates all required session tables.
// Reuses the same table definitions as session/postgres.
func createTables(
	ctx context.Context,
	client storage.Client,
	schema, prefix string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// createIndexes creates all required table indexes.
func createIndexes(
	ctx context.Context,
	client storage.Client,
	schema, prefix string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// addVectorColumns adds content_text, role, embedding, and
// search_vector columns to session_events if they do not
// already exist.
func (s *Service) addVectorColumns(
	ctx context.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) validateEmbeddingColumnDimension(
	ctx context.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func parseVectorColumnDimension(typeName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// createTextSearchIndex creates a GIN index on the
// generated search_vector column for keyword search.
func (s *Service) createTextSearchIndex(
	ctx context.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

// createHNSWIndex creates an HNSW vector index on the
// embedding column for cosine similarity search.
func (s *Service) createHNSWIndex(
	ctx context.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}
