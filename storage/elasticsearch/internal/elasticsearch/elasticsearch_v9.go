//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package elasticsearch

import (
	"context"

	esv9 "github.com/elastic/go-elasticsearch/v9"
)

var _ Client = (*clientV9)(nil)

// NewClientV9 creates a new clientV9.
func NewClientV9(esClient *esv9.Client) Client { _ = "STUB: not implemented"; return *new(Client) }

// client implements the ielasticsearch.Client interface.
type clientV9 struct {
	esClient *esv9.Client
}

// Ping checks if Elasticsearch is available.
func (c *clientV9) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CreateIndex creates an index with the provided body.
func (c *clientV9) CreateIndex(ctx context.Context, indexName string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteIndex deletes an index.
func (c *clientV9) DeleteIndex(ctx context.Context, indexName string) error {
	_ = "STUB: not implemented"
	return nil
}

// IndexExists checks if an index exists.
func (c *clientV9) IndexExists(ctx context.Context, indexName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IndexDocument indexes a document.
func (c *clientV9) IndexDoc(ctx context.Context, indexName, id string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDocument retrieves a document by ID.
func (c *clientV9) GetDoc(ctx context.Context, indexName, id string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateDoc updates a document.
func (c *clientV9) UpdateDoc(ctx context.Context, indexName, id string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteDoc deletes a document.
func (c *clientV9) DeleteDoc(ctx context.Context, indexName, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// Search performs a search query.
func (c *clientV9) Search(ctx context.Context, indexName string, body []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Count executes a count query.
func (c *clientV9) Count(ctx context.Context, indexName string, body []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Parse count response

// DeleteByQuery deletes documents matching the query.
func (c *clientV9) DeleteByQuery(ctx context.Context, indexName string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Refresh refreshes an index.
func (c *clientV9) Refresh(ctx context.Context, indexName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check response status
