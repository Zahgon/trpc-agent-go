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
	"github.com/jackc/pgx/v5"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

// defaultMaxResults is the default maximum number of search results.
const defaultMaxResults = 10

const (
	// Default HNSW index parameters.
	defaultHNSWM              = 16 // Default: 16 connections per layer
	defaultHNSWEfConstruction = 64 // Default: 64 for construction

	// Default IVFFlat index parameters.
	defaultIVFFlatLists = 100 // Default: 100 lists
)

// VectorIndexType represents the type of vector index to use.
type VectorIndexType string

const (
	// VectorIndexHNSW uses Hierarchical Navigable Small World (HNSW) index.
	VectorIndexHNSW VectorIndexType = "hnsw"

	// VectorIndexIVFFlat uses Inverted File with Flat compression.
	VectorIndexIVFFlat VectorIndexType = "ivfflat"
)

// HNSWIndexParams contains parameters for HNSW index.
type HNSWIndexParams struct {
	// M is the maximum number of connections per layer (default: 16, range: 2-100).
	M int

	// EfConstruction is the size of dynamic candidate list for construction (default: 64, range: 4-1000).
	EfConstruction int
}

// IVFFlatIndexParams contains parameters for IVFFlat index.
type IVFFlatIndexParams struct {
	// Lists is the number of inverted lists (default: 100).
	Lists int
}

// HybridFusionMode represents the fusion mode for hybrid search.
type HybridFusionMode int

const (
	// HybridFusionWeighted uses weighted fusion (default).
	// Formula: score = vector_score * vectorWeight + text_score * textWeight
	HybridFusionWeighted HybridFusionMode = iota

	// HybridFusionRRF uses Reciprocal Rank Fusion.
	// Formula: score = sum(1 / (k + rank_i)) for each ranking list
	HybridFusionRRF
)

// RRFParams contains parameters for Reciprocal Rank Fusion.
type RRFParams struct {
	// K is the RRF constant (default: 60).
	// Smaller values give more weight to top-ranked results.
	// Must be > 0. Typical range: [1, 100].
	K int

	// CandidateRatio controls how many candidates to fetch from each sub-search.
	// For RRF, we fetch (limit * CandidateRatio) candidates from each search.
	// Larger values improve fusion quality but increase query cost.
	// Must be > 0. Default: 3 (fetch 3x candidates).
	CandidateRatio int
}

// DocBuilderFunc is the document builder function.
type DocBuilderFunc func(row pgx.Row) (*vectorstore.ScoredDocument, []float64, error)

func defaultDocBuilder(row pgx.Row) (*vectorstore.ScoredDocument, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// options contains the options for pgvector.
type options struct {
	host           string // PostgreSQL host
	port           int    // PostgreSQL port
	user           string // PostgreSQL user
	password       string // PostgreSQL password
	database       string // PostgreSQL database
	table          string // PostgreSQL table
	indexDimension int    // PostgreSQL index dimension
	sslMode        string // PostgreSQL SSL mode
	dsn            string // PostgreSQL DSN
	enableTSVector bool   // Enable text search vector
	instanceName   string // Registered postgres instance name from storage/postgres
	extraOptions   []any  // Extra options for storage/postgres

	// Vector index configuration
	vectorIndexType VectorIndexType     // Type of vector index (hnsw or ivfflat)
	hnswParams      *HNSWIndexParams    // HNSW index parameters
	ivfflatParams   *IVFFlatIndexParams // IVFFlat index parameters

	// Hybrid search scoring weights
	vectorWeight float64 // Weight for vector similarity (0.0-1.0)
	textWeight   float64 // Weight for text relevance (0.0-1.0)
	language     string  // Default: english, if you install zhparser or jieba, you can set it to your configuration

	// Hybrid search fusion mode
	fusionMode HybridFusionMode // Fusion mode for hybrid search (weighted or RRF)
	rrfParams  *RRFParams       // RRF parameters (K and CandidateRatio)

	// Sparse (text) score configuration (internal use only, not exposed as public options)
	sparseRankFunc     string  // PostgreSQL ranking function: ts_rank or ts_rank_cd (default: ts_rank)
	sparseQueryFunc    string  // PostgreSQL query parsing function: plainto_tsquery, phraseto_tsquery, or websearch_to_tsquery (default: websearch_to_tsquery)
	sparseNormConstant float64 // Normalization constant for sparse score formula x/(x+c), smaller c means more sensitive (default: 0.1)

	docBuilder DocBuilderFunc

	maxResults int // Maximum number of search results

	//field
	// idFieldName is the PostgreSQL field name for ID.
	idFieldName string
	// nameFieldName is the PostgreSQL field name for name/title.
	nameFieldName string
	// contentFieldName is the PostgreSQL field name for content.
	contentFieldName string
	// embeddingFieldName is the PostgreSQL field name for embedding.
	embeddingFieldName string
	// metadataFieldName is the PostgreSQL field name for metadata.
	metadataFieldName string
	// createdAtFieldName is the PostgreSQL field name for created at timestamp.
	createdAtFieldName string
	// updatedAtFieldName is the PostgreSQL field name for updated at timestamp.
	updatedAtFieldName string
}

// defaultOptions is the default options for pgvector.
var defaultOptions = options{
	database:           "trpc_agent_go",
	table:              "documents",
	enableTSVector:     true,
	indexDimension:     1536,
	sslMode:            "disable",
	vectorWeight:       0.7,                    // Default: Vector similarity weight 70%
	textWeight:         0.3,                    // Default: Text relevance weight 30%
	sparseRankFunc:     "ts_rank",              // Default: more generous ranking than ts_rank_cd
	sparseQueryFunc:    "websearch_to_tsquery", // Default: more flexible than plainto_tsquery
	sparseNormConstant: 0.1,                    // Default: normalization constant for text score
	language:           "english",
	maxResults:         defaultMaxResults,
	docBuilder:         defaultDocBuilder,

	// Hybrid fusion defaults
	fusionMode: HybridFusionWeighted, // Default: weighted fusion
	rrfParams: &RRFParams{
		K:              60, // Default: RRF constant k=60
		CandidateRatio: 3,  // Default: fetch 3x candidates for RRF
	},

	// Vector index defaults (HNSW)
	vectorIndexType: VectorIndexHNSW,
	hnswParams: &HNSWIndexParams{
		M:              defaultHNSWM,
		EfConstruction: defaultHNSWEfConstruction,
	},
	ivfflatParams: &IVFFlatIndexParams{
		Lists: defaultIVFFlatLists,
	},

	idFieldName:        "id",
	nameFieldName:      "name",
	contentFieldName:   "content",
	embeddingFieldName: "embedding",
	metadataFieldName:  "metadata",
	createdAtFieldName: "created_at",
	updatedAtFieldName: "updated_at",
}

// Option is the option for pgvector.
type Option func(*options)

// WithHost sets the PostgreSQL host.
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPort sets the PostgreSQL port.
func WithPort(port int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUser sets the username for authentication.
func WithUser(user string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets the password for authentication.
func WithPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDatabase sets the database name.
func WithDatabase(database string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTable sets the table name.
func WithTable(table string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIndexDimension sets the vector dimension for the index.
// This dimension is used when creating the table schema (vector column definition).
func WithIndexDimension(dimension int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSSLMode sets the SSL mode for connection.
func WithSSLMode(sslMode string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPGVectorClientDSN sets the DSN for connection.
func WithPGVectorClientDSN(dsn string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTSVector sets the enable text search vector.
func WithEnableTSVector(enableTSVector bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHybridSearchWeights sets the weights for hybrid search scoring.
// vectorWeight: Weight for vector similarity (0.0-1.0)
// textWeight: Weight for text relevance (0.0-1.0)
// Note: weights will be normalized to sum to 1.0
// Note: This option only applies when fusionMode is HybridFusionWeighted
func WithHybridSearchWeights(vectorWeight, textWeight float64) Option {
	_ = "STUB: not implemented"
	return *

	// Normalize weights to sum to 1.0
	new(Option)
}

// Fallback to defaults if invalid weights

// WithHybridFusionMode sets the fusion mode for hybrid search.
// Default is HybridFusionWeighted.
// Use HybridFusionRRF for Reciprocal Rank Fusion.
func WithHybridFusionMode(mode HybridFusionMode) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRRFParams sets the parameters for Reciprocal Rank Fusion.
// Values <= 0 are ignored (defaults are kept).
// Note: This option only applies when fusionMode is HybridFusionRRF.
func WithRRFParams(params *RRFParams) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLanguageExtension sets the language extension for the index.
func WithLanguageExtension(languageExtension string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMaxResults sets the maximum number of search results.
func WithMaxResults(maxResults int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIDField sets the PostgreSQL field name for ID.
func WithIDField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNameField sets the PostgreSQL field name for name/title.
func WithNameField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentField sets the PostgreSQL field name for content.
func WithContentField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEmbeddingField sets the PostgreSQL field name for embedding.
func WithEmbeddingField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataField sets the PostgreSQL field name for metadata.
func WithMetadataField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCreatedAtField sets the PostgreSQL field name for created_at.
func WithCreatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdatedAtField sets the PostgreSQL field name for updated_at.
func WithUpdatedAtField(field string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDocBuilder sets the document builder function.
func WithDocBuilder(builder DocBuilderFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPostgresInstance uses a postgres instance from storage/postgres.
// Note: Direct connection settings (WithHost, WithPort, etc.) have higher priority than WithPostgresInstance.
// If both are specified, direct connection settings will be used.
func WithPostgresInstance(instanceName string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtraOptions sets extra options for storage/postgres.
// This is mainly used for customized postgres client builders.
func WithExtraOptions(extraOptions ...any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVectorIndexType sets the type of vector index to use.
// Supported types: VectorIndexHNSW (default), VectorIndexIVFFlat.
func WithVectorIndexType(indexType VectorIndexType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHNSWIndexParams sets HNSW index parameters.
func WithHNSWIndexParams(params *HNSWIndexParams) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithIVFFlatIndexParams sets IVFFlat index parameters.
func WithIVFFlatIndexParams(params *IVFFlatIndexParams) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
