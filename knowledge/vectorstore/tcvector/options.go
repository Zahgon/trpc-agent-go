//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tcvector

import (
	"github.com/tencent/vectordatabase-sdk-go/tcvectordb"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/source"
)

// defaultMaxResults is the default maximum number of search results.
const defaultMaxResults = 10

// DocBuilderFunc is the document builder function.
type DocBuilderFunc func(tcDoc tcvectordb.Document) (*document.Document, []float64, error)

// options contains the options for tcvectordb.
type options struct {
	username       string
	password       string
	url            string
	database       string
	collection     string
	indexDimension uint32
	replicas       uint32
	sharding       uint32
	enableTSVector bool
	sparseEncoder  TCSparseEncoder
	instanceName   string
	extraOptions   []any

	// Remote embedding configuration.
	// When embeddingModel is set, remote embedding is automatically enabled.
	embeddingModel string // Embedding model name for remote computation

	// Hybrid search scoring weights.
	vectorWeight float64 // Default: Vector similarity weight 70%
	textWeight   float64 // Default: Text relevance weight 30%
	language     string  // Default: zh, options: zh, en

	// Filter index configuration.
	filterAll bool // Enable filterAll to skip scalar field index creation and validation

	//field
	// idFieldName is the tcvectordb field name for ID.
	idFieldName string
	// nameFieldName is the tcvectordb field name for name/title.
	nameFieldName string
	// contentFieldName is the tcvectordb field name for content.
	contentFieldName string
	// embeddingFieldName is the tcvectordb field name for embedding.
	embeddingFieldName string
	// metadataFieldName is the tcvectordb field name for metadata.
	metadataFieldName string
	// createdAtFieldName is the tcvectordb field name for created at timestamp.
	createdAtFieldName string
	// updatedAtFieldName is the tcvectordb field name for updated at timestamp.
	updatedAtFieldName string
	// sparseVectorFieldName is the tcvectordb field name for sparse vector.
	sparseVectorFieldName string

	// filterField is the field name to filter the document.
	filterFields  []string
	filterIndexes []tcvectordb.FilterIndex

	docBuilder DocBuilderFunc

	// maxResults is the maximum number of search results.
	maxResults int
}

var defaultOptions = options{
	indexDimension:        1536,
	database:              "trpc-agent-go",
	collection:            "documents",
	replicas:              0,
	sharding:              1,
	enableTSVector:        true,
	embeddingModel:        "",
	vectorWeight:          0.7,
	textWeight:            0.3,
	language:              "en",
	filterFields:          []string{source.MetaURI, source.MetaSourceName},
	maxResults:            defaultMaxResults,
	idFieldName:           "id",
	nameFieldName:         "name",
	contentFieldName:      "content",
	embeddingFieldName:    "vector",
	metadataFieldName:     "metadata",
	createdAtFieldName:    "created_at",
	updatedAtFieldName:    "updated_at",
	sparseVectorFieldName: "sparse_vector",
	filterIndexes: []tcvectordb.FilterIndex{
		{
			FieldName: source.MetaURI,
			IndexType: tcvectordb.FILTER,
			FieldType: tcvectordb.String,
		},
		{
			FieldName: source.MetaSourceName,
			IndexType: tcvectordb.FILTER,
			FieldType: tcvectordb.String,
		},
	},
}

// Option is the option for tcvectordb.
type Option func(*options)

// WithURL sets the vector database URL.
func WithURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUsername sets the username for authentication.
func WithUsername(username string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets the password for authentication.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDatabase sets the database name.
func WithDatabase(database string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCollection sets the collection name.
func WithCollection(collection string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIndexDimension sets the vector dimension for the index.
func WithIndexDimension(dimension uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReplicas sets the number of replicas.
func WithReplicas(replicas uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSharding sets the number of shards.
func WithSharding(sharding uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTSVector sets the enableTSVector for the vector database.
func WithEnableTSVector(enableTSVector bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHybridSearchWeights sets the weights for hybrid search scoring.
// vectorWeight: Weight for vector similarity (0.0-1.0)
// textWeight: Weight for text relevance (0.0-1.0)
// Note: weights will be normalized to sum to 1.0
func WithHybridSearchWeights(vectorWeight, textWeight float64) Option {
	_ = "STUB: not implemented"
	return *

	// Normalize weights to sum to 1.0.
	new(Option)
}

// Fallback to defaults if invalid weights.

// WithLanguage sets the language for the vector database.
func WithLanguage(language string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTCVectorInstance uses a tcvectordb instance from storage.
// Note: WithURL, WithUserName, WithPassword has higher priority than WithTCVectorInstance.
// If both are specified, WithURL, WithUserName, WithPassword will be used.
func WithTCVectorInstance(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraOptions passes through extra client builder options to storage tcvector.
// It is mainly for customized client builders; the default builder ignores them.
func WithExtraOptions(extraOptions ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFilterIndexFields creates dedicated indexes for specified metadata fields.
// This is optional and provides better query performance for frequently queried fields.
// Other metadata fields can still be queried via the default JSON index.
//
// It will build additional indexes for the specified filter fields.
func WithFilterIndexFields(fields []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDocBuilder sets the document builder function.
func WithDocBuilder(builder DocBuilderFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxResults sets the maximum number of search results.
func WithMaxResults(maxResults int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIDField sets the tcvectordb field name for ID.
func WithIDField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNameField sets the tcvectordb field name for name/title.
func WithNameField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentField sets the tcvectordb field name for content.
func WithContentField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEmbeddingField sets the tcvectordb field name for embedding.
// This field value type is []float64
func WithEmbeddingField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataField sets the tcvectordb field name for metadata.
func WithMetadataField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCreatedAtField sets the tcvectordb field name for created_at.
// This field value type is uint64, so the value is converted to time.Time
func WithCreatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdatedAtField sets the tcvectordb field name for updated_at.
// This field value type is uint64, so the value is converted to time.Time
func WithUpdatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSparseVectorField sets the tcvectordb field name for sparse vector.
func WithSparseVectorField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTCSparseEncoder sets the sparse encoder for keyword and hybrid search.
func WithTCSparseEncoder(sparseEncoder TCSparseEncoder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRemoteEmbeddingModel sets the embedding model name for remote computation.
// When set, remote embedding is automatically enabled, and text queries will be sent
// directly to tcvectordb for embedding computation.
// Common models: bge-base-zh, bge-large-zh, m3e-base, text2vec-large-chinese, etc.
// Set to empty string to disable remote embedding.
func WithRemoteEmbeddingModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFilterAll enables filterAll mode for filter index configuration.
// When enabled, all scalar fields can be used for filtering without creating indexes,
// which skips index creation and validation for scalar fields.
// This is useful when you want to filter on many fields without the overhead of maintaining indexes.
func WithFilterAll(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }
