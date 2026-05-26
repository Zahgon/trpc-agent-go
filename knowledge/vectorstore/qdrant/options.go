//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package qdrant

import (
	"time"

	"github.com/qdrant/go-client/qdrant"

	"trpc.group/trpc-go/trpc-agent-go/log"
	qdrantstorage "trpc.group/trpc-go/trpc-agent-go/storage/qdrant"
)

const (
	defaultCollectionName  = "trpc_agent_documents"
	defaultDimension       = 1536
	defaultHNSWM           = 16
	defaultHNSWEfConstruct = 128
	defaultMaxResults      = 10
	defaultMaxRetries      = 3
	defaultBaseRetryDelay  = 100 * time.Millisecond
	defaultMaxRetryDelay   = 5 * time.Second

	// BM25 model for server-side sparse vector inference
	defaultBM25Model = "Qdrant/bm25"

	// Hybrid search prefetch settings
	defaultPrefetchMultiplier = 2    // Multiply limit by this for prefetch
	minPrefetchLimit          = 20   // Minimum prefetch to ensure quality fusion
	maxPrefetchLimit          = 1000 // Cap prefetch to avoid excessive memory usage

	// Batch size for pagination operations
	defaultBatchSize = 100
)

// Payload field names for stored documents.
const (
	fieldID        = "original_id"
	fieldName      = "name"
	fieldContent   = "content"
	fieldMetadata  = "metadata"
	fieldCreatedAt = "created_at"
	fieldUpdatedAt = "updated_at"
)

// Vector names for named vectors configuration
const (
	vectorNameDense  = "dense"  // Dense vector from embeddings
	vectorNameSparse = "sparse" // Sparse vector for BM25
)

// Distance represents the distance metric used for vector similarity search.
type Distance = qdrant.Distance

const (
	// DistanceCosine measures the cosine of the angle between two vectors.
	DistanceCosine = qdrant.Distance_Cosine

	// DistanceEuclid measures the straight-line distance between two points.
	DistanceEuclid = qdrant.Distance_Euclid

	// DistanceDot computes the dot product between two vectors.
	DistanceDot = qdrant.Distance_Dot

	// DistanceManhattan measures the sum of absolute differences (L1 norm).
	DistanceManhattan = qdrant.Distance_Manhattan
)

type options struct {
	client             qdrantstorage.Client
	clientBuilderOpts  []qdrantstorage.ClientBuilderOpt
	collectionName     string
	dimension          int
	distance           Distance
	onDiskVectors      bool
	onDiskPayload      bool
	hnswM              int
	hnswEfConstruct    int
	maxResults         int
	maxRetries         int
	baseRetryDelay     time.Duration
	maxRetryDelay      time.Duration
	bm25Enabled        bool
	prefetchMultiplier int
	logger             log.Logger
}

var defaultOptions = options{
	collectionName:     defaultCollectionName,
	dimension:          defaultDimension,
	distance:           DistanceCosine,
	hnswM:              defaultHNSWM,
	hnswEfConstruct:    defaultHNSWEfConstruct,
	maxResults:         defaultMaxResults,
	maxRetries:         defaultMaxRetries,
	baseRetryDelay:     defaultBaseRetryDelay,
	maxRetryDelay:      defaultMaxRetryDelay,
	prefetchMultiplier: defaultPrefetchMultiplier,
}

// Option is a functional option for configuring the VectorStore.
type Option func(*options)

// WithHost sets the Qdrant server host.
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPort sets the Qdrant server gRPC port.
func WithPort(port int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the API key for Qdrant Cloud authentication.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTLS enables TLS for secure connections (required for Qdrant Cloud).
func WithTLS(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCollectionName sets the collection name.
func WithCollectionName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimension sets the vector dimension. Must be positive.
func WithDimension(dim int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDistance sets the distance metric for similarity search.
func WithDistance(d Distance) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHNSWConfig sets HNSW index parameters.
func WithHNSWConfig(m, efConstruct int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOnDiskVectors enables on-disk vector storage for large datasets.
func WithOnDiskVectors(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOnDiskPayload enables on-disk payload storage.
func WithOnDiskPayload(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxResults sets the default maximum number of search results.
func WithMaxResults(max int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxRetries sets the maximum retry attempts for transient errors.
func WithMaxRetries(retries int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseRetryDelay sets the initial delay before the first retry.
func WithBaseRetryDelay(delay time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxRetryDelay sets the maximum delay between retries.
func WithMaxRetryDelay(delay time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClient sets a pre-created Qdrant client.
// When provided, connection options (WithHost, WithPort, WithAPIKey, WithTLS) are ignored.
// The caller retains ownership and must close the client separately.
func WithClient(client qdrantstorage.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBM25 enables BM25 sparse vectors for keyword and hybrid search.
// This creates a sparse vector index using Qdrant's native BM25 implementation
// with server-side inference using the "Qdrant/bm25" model.
//
// When enabled:
//   - Collection will have both dense and sparse vector configurations
//   - Documents are indexed with both embedding vectors and BM25 sparse vectors
//   - SearchModeKeyword uses BM25 sparse vector search
//   - SearchModeHybrid combines dense vector + BM25 with RRF fusion
//
// Note: Requires Qdrant Cloud or Qdrant with inference enabled.
func WithBM25(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPrefetchMultiplier sets the multiplier for hybrid search prefetch limit.
// In hybrid search, each sub-query (dense + BM25) prefetches limit × multiplier
// results before RRF fusion. Higher values improve fusion quality at the cost
// of increased latency and memory usage. Default is 2.
func WithPrefetchMultiplier(multiplier int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger sets the logger for operational messages.
// Used for logging warnings such as hybrid search fallback when BM25 is not enabled.
func WithLogger(logger log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }
