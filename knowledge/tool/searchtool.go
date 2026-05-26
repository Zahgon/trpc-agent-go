//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package tool provides knowledge search tools for agents.
package tool

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/knowledge"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/searchfilter"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultMaxResults        = 10
	defaultMinScore          = 0.0
	agenticFilterPromptIntro = "You are a helpful assistant that can search for relevant information in the knowledge base."

	// metadataFilterHint is appended to the tool description when no explicit
	// agenticFilterInfo is provided. It teaches the LLM that metadata fields
	// from search results can be used to construct filter conditions.
	metadataFilterHint = `Each search result contains two types of filterable fields:

  1. Document text — shown as "text" in results, but use field name
     "content" when constructing filters. Use the "like" operator for
     substring matching against the original document content.

  2. "metadata.*" — structured metadata fields. Keys in results appear
     without prefix (e.g. "source", "topic"), but when constructing
     filters, prefix them with "metadata." (e.g. "metadata.source").

You can use these field names and values from previous search results
to construct filter conditions for more precise subsequent searches.

Example — given a search result like:
  {
    "text": "Graph state is the core mechanism...",
    "metadata": {"source": "docs/graph.md", "topic": "architecture"},
    "score": 0.87
  }

  Filter by document text (use "content" as field name, not "text";
  use %% wildcards for substring matching):
    {"field": "content", "operator": "like", "value": "%graph state%"}

  Filter by metadata (note the "metadata." prefix):
    {"field": "metadata.topic", "operator": "eq", "value": "architecture"}

  Filter by metadata (multiple values):
    {"field": "metadata.topic", "operator": "in", "value": ["architecture", "api"]}

  Combine content and metadata:
    {"operator": "and", "value": [
      {"field": "content", "operator": "like", "value": "%graph state%"},
      {"field": "metadata.source", "operator": "eq", "value": "docs/graph.md"}
    ]}

Available operators:
  eq, ne, gt, gte, lt, lte, in, not in, like, not like, between, and, or

Usage modes:
  - query + filter: use together for semantic search narrowed by metadata constraints.
  - filter only: omit query and use filter alone for pure metadata-based retrieval.
  - query only: omit filter for standard semantic search.
  At least one of query or filter must be provided.

Note:
  For logical operators (and/or), use "value" to specify an array of sub-conditions.`
)

// KnowledgeSearchRequest represents the input for the knowledge search tool.
type KnowledgeSearchRequest struct {
	Query string `json:"query" jsonschema:"description=The search query to find relevant information in the knowledge base"`
}

// KnowledgeSearchResponse represents the response from the knowledge search tool.
type KnowledgeSearchResponse struct {
	Documents []*DocumentResult `json:"documents"`
	Message   string            `json:"message,omitempty"`
}

// DocumentResult represents a single document result with metadata and score.
type DocumentResult struct {
	Text     string         `json:"text"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Score    float64        `json:"score"`
}

// Option is a function that configures the knowledge search tool.
type Option func(*options)

// ResultPostProcessor post-processes the search response before it is returned
// to the LLM. It receives the current invocation context and the raw response,
// and may mutate the response (e.g. filter out duplicate documents) or rewrite
// the Message field. Implementations must be safe for concurrent use.
type ResultPostProcessor func(ctx context.Context, resp *KnowledgeSearchResponse) *KnowledgeSearchResponse

type options struct {
	toolName            string
	toolDescription     string
	staticFilter        map[string]any
	conditionedFilter   *searchfilter.UniversalFilterCondition
	maxResults          int
	minScore            float64
	excludeMetadataKeys map[string]struct{}
	postProcessor       ResultPostProcessor
}

// WithToolName sets the name of the knowledge search tool.
func WithToolName(toolName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithToolDescription sets the description of the knowledge search tool.
func WithToolDescription(toolDescription string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFilter sets a static metadata filter (simple AND logic).
// Multiple key-value pairs are combined with AND.
// For OR/nested conditions, use WithConditionedFilter.
func WithFilter(filter map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConditionedFilter sets a static complex filter with OR/AND/nested logic.
// Supports operators: eq, ne, gt, gte, lt, lte, in, not in, like, not like, between, and, or.
// For simple AND-only filters, use WithFilter instead.
func WithConditionedFilter(filterCondition *searchfilter.UniversalFilterCondition) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMaxResults sets the maximum number of documents to return.
// Default is 10 if not specified.
func WithMaxResults(maxResults int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMinScore sets the minimum relevance score threshold (0.0 to 1.0).
// Documents with scores below this threshold will be filtered out.
// Default is 0.0 if not specified.
func WithMinScore(minScore float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExcludeMetadataKeys excludes the specified metadata keys from each
// returned document before the result is handed to the LLM. This is useful
// when certain metadata fields are verbose or low-value for the model
// (e.g. import lists, raw language tag, chunk index).
//
// Semantics: this option is additive. Multiple calls to WithExcludeMetadataKeys
// accumulate into a single exclusion set on the underlying options, rather
// than replacing the previously-configured keys. Empty or duplicate keys are
// ignored. This makes it safe for higher-level helpers (for example, the code
// search tool constructor) to pass a default exclusion list and still let
// callers extend it.
func WithExcludeMetadataKeys(keys ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithResultPostProcessor registers a post-processor that is invoked right
// before the knowledge search response is returned to the LLM. Typical use
// cases include cross-call deduplication and result rewriting. Passing nil is
// a no-op.
func WithResultPostProcessor(p ResultPostProcessor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewKnowledgeSearchTool creates a function tool for knowledge search using
// the Knowledge interface.
// This tool allows agents to search for relevant information in the knowledge base.
func NewKnowledgeSearchTool(kb knowledge.Knowledge, opts ...Option) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

// Create search request - for tools, we don't have conversation history yet.
// This could be enhanced in the future to extract context from the agent's session.

// History, UserID, SessionID could be filled from agent context in the future.

// KnowledgeSearchRequestWithFilter represents the input with filter for the knowledge search tool.
type KnowledgeSearchRequestWithFilter struct {
	Query  string                                 `json:"query,omitempty" jsonschema:"description=The search query to find relevant information in the knowledge base. Can be empty when using only filters."`
	Filter *searchfilter.UniversalFilterCondition `json:"filter,omitempty" jsonschema:"description=Filter conditions to apply to the search query. Use lowercase operators eq ne gt gte lt lte in not in like not like between and or."`
}

// NewAgenticFilterSearchTool creates a knowledge search tool with dynamic agent-controlled filtering.
// The agent can analyze user queries and construct filters dynamically.
//
// Parameters:
//   - kb: The knowledge base to search
//   - agenticFilterInfo: Available metadata fields and values, e.g., {"category": ["doc", "tutorial"]}
//   - opts: Optional static filters (WithFilter/WithConditionedFilter) always applied
func NewAgenticFilterSearchTool(
	kb knowledge.Knowledge,
	agenticFilterInfo map[string][]any,
	opts ...Option,
) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

// Query can be empty when using only filters for metadata-based retrieval

// Set search mode based on whether query is provided
// When query is empty, use filter-only search mode

func composeAgenticToolDescription(toolDescription, filterInfo string) string {
	_ = "STUB: not implemented"
	return ""
}

// applyPostProcessor runs the optional ResultPostProcessor on resp. Passing a
// nil processor returns resp unchanged, so callers can always write
// `return applyPostProcessor(ctx, resp, opt.postProcessor), nil` without an
// extra nil check.
func applyPostProcessor(ctx context.Context, resp *KnowledgeSearchResponse, p ResultPostProcessor) *KnowledgeSearchResponse {
	_ = "STUB: not implemented"
	return nil
}

// convertSearchResults converts knowledge.SearchResult to KnowledgeSearchResponse.
func convertSearchResults(
	result *knowledge.SearchResult,
	excludeKeys map[string]struct{},
) (*KnowledgeSearchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertMetadataToFilterCondition converts a metadata map to UniversalFilterCondition.
func convertMetadataToFilterCondition(metadata map[string]any) *searchfilter.UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// mergeFilterConditions merges multiple filter conditions using AND logic.
// All non-nil conditions are combined with AND operator.
// Returns nil if all conditions are nil.
func mergeFilterConditions(
	conditions ...*searchfilter.UniversalFilterCondition,
) *searchfilter.UniversalFilterCondition {
	_ = "STUB: not implemented"
	return nil
}

// filterMetadata removes internal metadata keys with MetaPrefix from the metadata map.
// Keys present in excludeKeys are additionally stripped, regardless of prefix.
func filterMetadata(metadata map[string]any, excludeKeys map[string]struct{}) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func generateAgenticFilterPrompt(agenticFilterInfo map[string][]any) string {
	_ = "STUB: not implemented"
	return ""
}

// Build list of valid filter keys

// Separate keys with and without predefined values

// Print keys with predefined values first

// Print keys without predefined values
