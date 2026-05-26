//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package elasticsearch contains option definitions for the Elasticsearch vector store.
package elasticsearch

import (
	"encoding/json"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/storage/elasticsearch"
)

const (
	// defaultIndexName is the default index name for documents.
	defaultIndexName = "trpc_agent_documents"
	// defaultScoreThreshold is the default minimum similarity score.
	defaultScoreThreshold = 0.7
	// defaultVectorDimension is the default dimension for embedding vectors.
	defaultVectorDimension = 1536
	// defaultMaxResults is the default maximum number of search results.
	defaultMaxResults = 10
)

// DocBuilderFunc is a function that builds a document from hit source.
type DocBuilderFunc func(hitSource json.RawMessage) (*document.Document, []float64, error)

// options holds Elasticsearch vectorstore configuration.
type options struct {
	// addresses is a list of Elasticsearch node addresses.
	addresses []string
	// username for authentication.
	username string
	// password for authentication.
	password string
	// apiKey for API key authentication.
	apiKey string
	// certificateFingerprint for certificate-based authentication.
	certificateFingerprint string
	// compressRequestBody enables request compression.
	compressRequestBody bool
	// enableMetrics enables metrics collection.
	enableMetrics bool
	// enableDebugLogger enables debug logging.
	enableDebugLogger bool
	// retryOnStatus specifies HTTP status codes to retry on.
	retryOnStatus []int
	// maxRetries is the maximum number of retries.
	maxRetries int
	// indexName is the name of the Elasticsearch index.
	indexName string
	// scoreThreshold is the minimum similarity score threshold.
	scoreThreshold float64
	// maxResults is the maximum number of search results.
	maxResults int
	// vectorDimension is the dimension of embedding vectors.
	vectorDimension int
	// enableTSVector enables text search vector capabilities.
	enableTSVector bool
	// version is the Elasticsearch version to use (v7, v8, v9).
	version elasticsearch.ESVersion
	// idFieldName is the Elasticsearch field name for ID.
	idFieldName string
	// nameFieldName is the Elasticsearch field name for name/title.
	nameFieldName string
	// contentFieldName is the Elasticsearch field name for content.
	contentFieldName string
	// embeddingFieldName is the Elasticsearch field name for embedding.
	embeddingFieldName string
	// metadataFieldName is the Elasticsearch field name for metadata.
	metadataFieldName string
	// createdAtFieldName is the Elasticsearch field name for createdAt.
	createdAtFieldName string
	// updatedAtFieldName is the Elasticsearch field name for updatedAt.
	updatedAtFieldName string
	// extraOptions allows passing builder-specific extras to the storage client.
	extraOptions []any
	// docBuilder is the function to build document from hit source.
	docBuilder DocBuilderFunc
}

// defaultOptions returns default configuration.
var defaultOptions = options{
	addresses:           []string{"http://localhost:9200"},
	maxRetries:          3,
	compressRequestBody: true,
	enableMetrics:       false,
	enableDebugLogger:   false,
	retryOnStatus: []int{http.StatusInternalServerError, http.StatusBadGateway,
		http.StatusServiceUnavailable, http.StatusTooManyRequests},
	indexName:          defaultIndexName,
	scoreThreshold:     defaultScoreThreshold,
	maxResults:         defaultMaxResults,
	vectorDimension:    defaultVectorDimension,
	enableTSVector:     true,
	version:            elasticsearch.ESVersionV9, // Default to latest version.
	idFieldName:        "id",
	nameFieldName:      "name",
	contentFieldName:   "content",
	embeddingFieldName: "embedding",
	metadataFieldName:  "metadata",
	createdAtFieldName: "created_at",
	updatedAtFieldName: "updated_at",
}

// Option represents a functional option for configuring VectorStore.
type Option func(*options)

// WithAddresses sets the Elasticsearch node addresses.
func WithAddresses(addresses []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUsername sets the username for authentication.
func WithUsername(username string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets the password for authentication.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCertificateFingerprint sets the certificate fingerprint.
func WithCertificateFingerprint(fingerprint string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompressRequestBody enables request compression.
func WithCompressRequestBody(compress bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableMetrics enables metrics collection.
func WithEnableMetrics(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableDebugLogger enables debug logging.
func WithEnableDebugLogger(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetryOnStatus sets HTTP status codes to retry on.
func WithRetryOnStatus(statusCodes []int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxRetries sets the maximum number of retries.
func WithMaxRetries(maxRetries int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIndexName sets the Elasticsearch index name.
func WithIndexName(indexName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithScoreThreshold sets the minimum similarity score threshold.
func WithScoreThreshold(threshold float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxResults sets the maximum number of search results.
func WithMaxResults(maxResults int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVectorDimension sets the dimension of embedding vectors.
func WithVectorDimension(dimension int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTSVector enables text search vector capabilities.
func WithEnableTSVector(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVersion sets the Elasticsearch version to use (v7, v8, v9).
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIDField sets the Elasticsearch field name for ID.
func WithIDField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNameField sets the Elasticsearch field name for name/title.
func WithNameField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentField sets the Elasticsearch field name for content.
func WithContentField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEmbeddingField sets the Elasticsearch field name for embedding.
func WithEmbeddingField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataField sets the Elasticsearch field name for metadata.
func WithMetadataField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCreatedAtField sets the Elasticsearch field name for createdAt.
func WithCreatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdatedAtField sets the Elasticsearch field name for updatedAt.
func WithUpdatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraOptions sets extra builder-specific options for the storage client.
func WithExtraOptions(extraOptions ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDocBuilder sets the document builder function.
func WithDocBuilder(builder DocBuilderFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
