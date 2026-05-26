//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package gemini provides Gemini embedder implementation.
package gemini

import (
	"context"

	"google.golang.org/genai"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
)

// Verify that Embedder implements the embedder.Embedder interface.
var _ embedder.Embedder = (*Embedder)(nil)

const (
	// DefaultModel is the default Gemini embedding model.
	DefaultModel = ModelGeminiEmbedding001
	// DefaultDimensions is the default embedding dimension.
	DefaultDimensions = 1536
	// DefaultTaskType is the default task type.
	DefaultTaskType = TaskTypeRetrievalQuery
	// DefaultRole is the default role.
	DefaultRole = genai.RoleUser

	// ModelGeminiEmbeddingExp0307 represents the gemini-embedding-exp-03-07 model.
	ModelGeminiEmbeddingExp0307 = "gemini-embedding-exp-03-07"
	// ModelGeminiEmbedding001 represents the gemini-embedding-001 model.
	ModelGeminiEmbedding001 = "gemini-embedding-001"

	// TaskTypeSemanticSimilarity is a task type for assessing text similarity.
	// Applicable scenarios examples: recommendation systems, duplicate detection.
	TaskTypeSemanticSimilarity = "SEMANTIC_SIMILARITY"
	// TaskTypeClassification is a task type for classifying texts according to preset labels.
	// Applicable scenarios examples: sentiment analysis, spam detection.
	TaskTypeClassification = "CLASSIFICATION"
	// TaskTypeClustering is a task type for clustering texts based on their similarities.
	// Applicable scenarios examples: document organization, market research, anomaly detection.
	TaskTypeClustering = "CLUSTERING"
	// TaskTypeRetrievalDocument is a task type for document search.
	// Applicable scenarios examples: indexing articles, books, web pages for search.
	TaskTypeRetrievalDocument = "RETRIEVAL_DOCUMENT"
	// TaskTypeRetrievalQuery is a task type for general search queries.
	// Use RETRIEVAL_QUERY for queries and RETRIEVAL_DOCUMENT for documents to be retrieved.
	// Applicable scenarios examples: custom search.
	TaskTypeRetrievalQuery = "RETRIEVAL_QUERY"
	// TaskTypeCodeRetrievalQuery is a task type for retrieval of code blocks based on natural language queries.
	// Use CODE_RETRIEVAL_QUERY for queries and RETRIEVAL_DOCUMENT for code blocks to be retrieved.
	// Applicable scenarios examples: code suggestions, code search.
	TaskTypeCodeRetrievalQuery = "CODE_RETRIEVAL_QUERY"
	// TaskTypeQuestionAnswering is a task type for questions in a question-answering system.
	// Use QUESTION_ANSWERING for questions and RETRIEVAL_DOCUMENT for documents to be retrieved.
	// Applicable scenarios examples: chatbot.
	TaskTypeQuestionAnswering = "QUESTION_ANSWERING"
	// TaskTypeFactVerification is a task type for statements that need to be verified.
	// Use FACT_VERIFICATION for the target text and RETRIEVAL_DOCUMENT for documents to be retrieved.
	// Applicable scenarios examples: automated fact-checking systems.
	TaskTypeFactVerification = "FACT_VERIFICATION"

	// GoogleAPIKeyEnv is the environment variable name for the Google API key.
	// nolint:gosec // This is just the environment variable name, not a hardcoded credential.
	GoogleAPIKeyEnv = "GOOGLE_API_KEY"
)

// Embedder implements the embedder.Embedder interface for Gemini API.
type Embedder struct {
	client         *genai.Client
	model          string
	dimensions     int
	taskType       string
	title          string
	apiKey         string
	role           genai.Role
	clientOptions  *genai.ClientConfig
	requestOptions *genai.EmbedContentConfig
}

// Option represents a functional option for configuring the Embedder.
type Option func(*Embedder)

// WithModel sets the embedding model to use.
func WithModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimensions sets the number of dimensions for the embedding.
func WithDimensions(dimensions int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTaskType sets the task type to optimize embedding results.
// Choosing the appropriate task type can improve accuracy and efficiency.
func WithTaskType(taskType string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTitle sets the title for the text.
// Only applicable when TaskType is RETRIEVAL_DOCUMENT.
func WithTitle(title string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the Google API key.
// If not provided, will use GOOGLE_API_KEY environment variable.
// APIKey priority: WithClientOptions > WithAPIKey > GOOGLE_API_KEY environment variable.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRole sets the role when generating embeddings content.
func WithRole(role genai.Role) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClientOptions sets additional options for the Gemini client config.
// APIKey priority: WithClientOptions > WithAPIKey > GOOGLE_API_KEY environment variable.
func WithClientOptions(clientOptions *genai.ClientConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRequestOptions sets additional options for the Gemini client requests.
func WithRequestOptions(requestOptions *genai.EmbedContentConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// New creates a new Gemini embedder with the given options.
func New(ctx context.Context, opts ...Option) (*Embedder, error) {
	_ = "STUB: not implemented"
	// Create embedder with defaults.
	return nil, nil
}

// Apply functional options.

// Build client options.

// Create Gemini client.

// GetEmbedding implements the embedder.Embedder interface.
// It generates an embedding vector for the given text.
func (e *Embedder) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract embedding from response.

// GetEmbeddingWithUsage implements the embedder.Embedder interface.
// It generates an embedding vector for the given text and returns usage information.
func (e *Embedder) GetEmbeddingWithUsage(ctx context.Context, text string) ([]float64, map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Extract embedding from response.

// GetDimensions implements the embedder.Embedder interface.
// It returns the number of dimensions in the embedding vectors.
func (e *Embedder) GetDimensions() int { _ = "STUB: not implemented"; return 0 }

func (e *Embedder) response(ctx context.Context, text string) (rsp *genai.EmbedContentResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove the `models/` prefix from the model id if it exists.

// Create content from text.

// Create request.

// Check for integer overflow before conversion.

// Call Gemini embeddings API.
