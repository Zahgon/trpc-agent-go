//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main exposes the local trpc-agent-go code_search capability as an
// MCP server, so external MCP clients (e.g. Augment, Cursor, or another
// trpc-agent-go runner) can consume exactly the same AST-backed code search
// pipeline that comparison/local_agent.go uses.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	util "trpc.group/trpc-go/trpc-agent-go/examples/knowledge"
	"trpc.group/trpc-go/trpc-agent-go/knowledge"
	agenttool "trpc.group/trpc-go/trpc-agent-go/tool"
	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

const (
	serverName                = "trpc-agent-code-search-mcp"
	serverVersion             = "0.1.0"
	defaultServerAddr         = "127.0.0.1:3001"
	defaultServerPath         = "/mcp"
	repoURL                   = "https://github.com/trpc-group/trpc-agent-go"
	repoName                  = "trpc-agent-go"
	repoBranch                = "main"
	maxResults                = 5
	embeddingModelNameEnvKey  = "EMBEDDING_MODEL_NAME"
	defaultEmbeddingModelName = "text-embedding-3-small"
	vectorStoreTypeEnvKey     = "VECTOR_STORE_TYPE"
	filterSchemaDepth         = 2
)

const (
	codeSearchQueryDescription          = "The search query to find relevant information in the knowledge base. Can be empty when using only filters."
	codeSearchFilterDescription         = "Filter conditions to apply to the search query. Use lowercase operators eq ne gt gte lt lte in not in like not like between and or."
	codeSearchFilterFieldDescription    = "The metadata field to filter on. Use it for comparison operators and ignore it for logical operators and or."
	codeSearchFilterOperatorDescription = "The operator to use. Valid values are eq ne gt gte lt lte in not in like not like between and or."
	codeSearchFilterValueDescription    = "Comparison value for eq ne gt gte lt lte. Use an array for in not in between. Use an array of nested filter conditions for and or."
)

var (
	flagAddr        = flag.String("addr", defaultServerAddr, "HTTP listen address for the MCP server")
	flagPath        = flag.String("path", defaultServerPath, "HTTP path prefix for the MCP endpoint")
	flagSkipLoad    = flag.Bool("skip-load", false, "Skip repository ingestion and reuse the existing vector-store data as-is")
	flagTruncateOld = flag.Bool("truncate-old", false, "Recreate the vector store before ingestion (deletes all existing documents); implies -skip-load=false")
	flagStoreType   = flag.String("store", util.GetEnvOrDefault(vectorStoreTypeEnvKey, string(util.VectorStoreInMemory)), "Vector store type (inmemory|pgvector|sqlitevec|tcvector|elasticsearch|milvus)")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("code-search MCP server exited with error: %v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func buildKnowledge(
	storeType util.VectorStoreType,
	embeddingModelName string,
) (*knowledge.BuiltinKnowledge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCodeSearchMCPTool(decl *agenttool.Declaration) *mcp.Tool {
	_ = "STUB: not implemented"
	return nil
}

func withInputSchema(schema *openapi3.Schema) mcp.ToolOption {
	_ = "STUB: not implemented"
	return *new(mcp.ToolOption)
}

func codeSearchInputSchema() *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func codeSearchFilterConditionSchema(depth int) *openapi3.Schema {
	_ = "STUB: not implemented"
	return nil
}

func codeSearchFilterValueSchema(depth int) *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func codeSearchFilterValueAnyOf(depth int) openapi3.SchemaRefs {
	_ = "STUB: not implemented"
	return *new(openapi3.SchemaRefs)
}

func codeSearchFilterArrayItemSchema(depth int) *openapi3.Schema {
	_ = "STUB: not implemented"
	return nil
}

func codeSearchOperatorEnum() []any { _ = "STUB: not implemented"; return nil }

func schemaRef(schema *openapi3.Schema) *openapi3.SchemaRef { _ = "STUB: not implemented"; return nil }

func stringSchema(description string) *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func numberSchema(schemaType string) *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func booleanSchema() *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func objectSchema() *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func arraySchema(items *openapi3.Schema) *openapi3.Schema { _ = "STUB: not implemented"; return nil }

func disallowAdditionalProperties() openapi3.AdditionalProperties {
	_ = "STUB: not implemented"
	return *new(openapi3.AdditionalProperties)
}

// newCodeSearchHandler bridges an MCP CallToolRequest to the underlying
// trpc-agent-go CallableTool. The MCP arguments are re-encoded to JSON and
// forwarded as-is, so the behavior matches what comparison/local_agent.go
// sees when the local LLMAgent calls code_search directly.
func newCodeSearchHandler(callable agenttool.CallableTool) func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	return nil
}

func renderToolResult(result any) (string, error) { _ = "STUB: not implemented"; return "", nil }
