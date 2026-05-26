//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package milvus

import (
	"fmt"
	"time"

	"github.com/milvus-io/milvus/client/v2/column"
	"github.com/milvus-io/milvus/client/v2/entity"
	client "github.com/milvus-io/milvus/client/v2/milvusclient"
	"google.golang.org/grpc"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
)

const (
	// Default collection name
	defaultCollectionName = "trpc_agent_documents"
	// Default vector dimension
	defaultVectorDimension = 1536
	// Batch processing constants
	metadataBatchSize = 5000
	// Field names
	idFieldName        = "id"
	nameFieldName      = "name"
	contentFieldName   = "content"
	vectorFieldName    = "vector"
	metadataFieldName  = "metadata"
	createdAtFieldName = "created_at"
	updatedAtFieldName = "updated_at"
)

// DocBuilderFunc is the document builder function.
type DocBuilderFunc func(tcDoc []column.Column) (*document.Document, []float64, error)

// Option represents a functional option for configuring the Milvus vector store.
type Option func(*options)

// options holds the configuration for the Milvus vector store.
type options struct {
	// Collection configuration
	collectionName string
	dimension      int

	// Connection configuration
	address  string
	username string
	password string
	dbName   string
	apiKey   string
	dialOpts []grpc.DialOption

	// Field names
	idField            string
	nameField          string
	contentField       string
	contentSparseField string
	vectorField        string
	metadataField      string
	createdAtField     string
	updatedAtField     string
	allFields          []string

	// Search configuration
	maxResults         int
	enableHNSW         bool
	hnswM              int
	hnswEfConstruction int
	metricType         entity.MetricType

	// docBuilder is the function to build document
	docBuilder DocBuilderFunc

	// reranker
	reranker client.Reranker
}

// defaultOptions returns the default configuration.
var defaultOptions = options{
	collectionName:     defaultCollectionName,
	dimension:          defaultVectorDimension,
	dialOpts:           []grpc.DialOption{grpc.WithTimeout(5 * time.Second)},
	maxResults:         10,
	enableHNSW:         true,
	hnswM:              16,
	hnswEfConstruction: 128,
	metricType:         entity.IP,
	idField:            idFieldName,
	nameField:          nameFieldName,
	contentField:       contentFieldName,
	contentSparseField: fmt.Sprintf("%s_sparse", contentFieldName),
	vectorField:        vectorFieldName,
	metadataField:      metadataFieldName,
	createdAtField:     createdAtFieldName,
	updatedAtField:     updatedAtFieldName,
}

// WithCollectionName sets the collection name.
func WithCollectionName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimension sets the vector dimension.
func WithDimension(dim int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAddress sets the Milvus server address.
func WithAddress(address string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUsername sets the username for authentication.
func WithUsername(username string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets the password for authentication.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDBName sets the database name.
func WithDBName(dbName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDialOptions sets the gRPC dial options.
func WithDialOptions(opts ...grpc.DialOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMaxResults sets the maximum number of search results.
func WithMaxResults(max int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHNSWParams sets HNSW index parameters.
func WithHNSWParams(m, efConstruction int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetricType sets the metric type for vector similarity search.
// Supported types: entity.IP (inner product), entity.L2 (Euclidean distance), entity.COSINE.
// Default is entity.IP which returns higher scores for more similar vectors.
func WithMetricType(metricType entity.MetricType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithIDField sets the ID field name.
func WithIDField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNameField sets the name field name.
func WithNameField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentField sets the content field name.
func WithContentField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVectorField sets the vector field name.
func WithVectorField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataField sets the metadata field name.
func WithMetadataField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCreatedAtField sets the createdAt field name.
func WithCreatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdatedAtField sets the updatedAt field name.
func WithUpdatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReranker sets the reranker.
func WithReranker(reranker client.Reranker) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDocBuilder sets the doc builder.
func WithDocBuilder(builder DocBuilderFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
