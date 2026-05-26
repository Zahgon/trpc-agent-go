//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package util provides utility functions.
package util

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

// VectorStoreType defines the type of vector store.
type VectorStoreType string

// Vector store type constants.
const (
	VectorStoreInMemory      VectorStoreType = "inmemory"
	VectorStorePGVector      VectorStoreType = "pgvector"
	VectorStoreSQLiteVec     VectorStoreType = "sqlitevec"
	VectorStoreTCVector      VectorStoreType = "tcvector"
	VectorStoreElasticsearch VectorStoreType = "elasticsearch"
	VectorStoreMilvus        VectorStoreType = "milvus"
)

// NewVectorStoreByType creates a vector store based on the specified type.
func NewVectorStoreByType(storeType VectorStoreType) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

// NewVectorStoreByTypeWithDimension creates a vector store based on the
// specified type and, when relevant, configures it for the embedder dimension
// used by the caller.
func NewVectorStoreByTypeWithDimension(
	storeType VectorStoreType,
	indexDimension int,
) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func newPGVectorStore(indexDimension int) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func newTCVectorStore() (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func newElasticsearchStore() (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func newMilvusStore() (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

// WaitForIndexRefresh waits for Elasticsearch index refresh.
// Elasticsearch need a short time to refresh index after index creation or data insertion.
// Milvus also needs a short time to load collection after data insertion.
func WaitForIndexRefresh(storeType VectorStoreType) { _ = "STUB: not implemented"; return }

// PrintEventWithToolCalls prints the event with tool calls.
func PrintEventWithToolCalls(evt *event.Event) { _ = "STUB: not implemented"; return }

// Print tool calls

// Print tool responses

// printToolResult pretty prints tool responses, focusing on document results.
func printToolResult(toolResult map[string]any) { _ = "STUB: not implemented"; return }

// printDocumentResults renders document arrays in a compact, readable form.
func printDocumentResults(toolResult map[string]any) bool { _ = "STUB: not implemented"; return false }

// normalizeText trims, flattens, and truncates text for display.
func normalizeText(value any, limit int) string { _ = "STUB: not implemented"; return "" }

// normalizeScore converts score-like values to float64.
func normalizeScore(value any) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

// formatMetadata prints metadata as sorted key-value pairs.
func formatMetadata(value any) string { _ = "STUB: not implemented"; return "" }

// GetEnvOrDefault retrieves the value of an environment variable or returns a default value if not set.
func GetEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

// ExampleDataPath returns the absolute path to example data files.
// relativePath is relative to the exampledata directory.
func ExampleDataPath(relativePath string) string {
	_ = "STUB: not implemented"
	// Get the directory of this source file (util.go)
	return ""
}

// ScoredDoc represents a document with its score for ranking comparison.
type ScoredDoc struct {
	Index int
	Score float64
	Text  string
}

// CalculateEmbeddingScores calculates cosine similarity scores between query and documents.
func CalculateEmbeddingScores(ctx context.Context, query string, documents []string, emb embedder.Embedder) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// NewOpenAIEmbedder creates a new OpenAI embedder with the specified model.
func NewOpenAIEmbedder(model string) embedder.Embedder {
	_ = "STUB: not implemented"
	return *new(embedder.Embedder)
}

// CosineSimilarity calculates cosine similarity between two vectors.
func CosineSimilarity(a, b []float64) float64 { _ = "STUB: not implemented"; return 0 }

// PrintEmbeddingResults prints embedding scores sorted by score descending.
func PrintEmbeddingResults(scores []float64, documents []string) { _ = "STUB: not implemented"; return }
