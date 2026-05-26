//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"context"

	"gopkg.in/yaml.v3"

	"trpc.group/trpc-go/trpc-agent-go/knowledge"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

const (
	genericKnowledgeToolName = "knowledge_search"
	knowledgeToolNameSuffix  = "_knowledge_search"
)

// knowledgeEntry is the parsed intermediate representation of a knowledge
// provider configuration.
type knowledgeEntry struct {
	Type        string
	Name        string
	Description string
	MaxResults  int
	Config      *yaml.Node
}

type knowledgeToolsBundle struct {
	tools []tool.Tool
}

type rawKnowledgeComponent struct {
	Node *yaml.Node
}

func (r *rawKnowledgeComponent) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil

	// builtinKnowledgeConfig is the config schema for the "builtin" knowledge
	// provider (embedder + vector_store).
}

type builtinKnowledgeConfig struct {
	Embedder    *rawKnowledgeComponent `yaml:"embedder,omitempty"`
	VectorStore *rawKnowledgeComponent `yaml:"vector_store,omitempty"`
}

func newBuiltinKnowledge(
	_ registry.KnowledgeProviderDeps,
	spec registry.PluginSpec,
) (knowledge.Knowledge, error) {
	_ = "STUB: not implemented"
	return *new(knowledge.Knowledge), nil
}

func buildKnowledgeTools(
	entries []knowledgeEntry,
) (*knowledgeToolsBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type knowledgeTypeConfig struct {
	Type string `yaml:"type,omitempty"`
}

type openAIKnowledgeEmbedderConfig struct {
	Type       string `yaml:"type,omitempty"`
	Model      string `yaml:"model,omitempty"`
	BaseURL    string `yaml:"base_url,omitempty"`
	APIKey     string `yaml:"api_key,omitempty"`
	Dimensions *int   `yaml:"dimensions,omitempty"`
}

type inmemoryKnowledgeVectorStoreConfig struct {
	Type       string `yaml:"type,omitempty"`
	MaxResults *int   `yaml:"max_results,omitempty"`
}

type pgvectorKnowledgeVectorStoreConfig struct {
	Type           string `yaml:"type,omitempty"`
	URL            string `yaml:"url,omitempty"`
	Table          string `yaml:"table,omitempty"`
	EnableTSVector *bool  `yaml:"enable_tsvector,omitempty"`
	IndexDimension *int   `yaml:"index_dimension,omitempty"`
	MaxResults     *int   `yaml:"max_results,omitempty"`
}

type elasticsearchKnowledgeVectorStoreConfig struct {
	Type            string   `yaml:"type,omitempty"`
	Addresses       []string `yaml:"addresses,omitempty"`
	Username        string   `yaml:"username,omitempty"`
	Password        string   `yaml:"password,omitempty"`
	APIKey          string   `yaml:"api_key,omitempty"`
	IndexName       string   `yaml:"index_name,omitempty"`
	VectorDimension *int     `yaml:"vector_dimension,omitempty"`
	MaxResults      *int     `yaml:"max_results,omitempty"`
}

type knowledgeVectorStoreBuildContext struct {
	embedder embedder.Embedder
}

type knowledgeEmbedderBuilder func(
	node *yaml.Node,
) (embedder.Embedder, error)

type knowledgeVectorStoreBuilder func(
	node *yaml.Node,
	ctx knowledgeVectorStoreBuildContext,
) (vectorstore.VectorStore, error)

var knowledgeEmbedderBuilders = map[string]knowledgeEmbedderBuilder{
	"":       buildOpenAIKnowledgeEmbedder,
	"openai": buildOpenAIKnowledgeEmbedder,
}

var knowledgeVectorStoreBuilders = map[string]knowledgeVectorStoreBuilder{
	"inmemory":      buildInMemoryKnowledgeVectorStore,
	"pgvector":      buildPGVectorKnowledgeVectorStore,
	"elasticsearch": buildElasticsearchKnowledgeVectorStore,
}

func buildKnowledgeEmbedder(
	cfg *rawKnowledgeComponent,
) (embedder.Embedder, error) {
	_ = "STUB: not implemented"
	return *new(embedder.Embedder), nil
}

func buildKnowledgeVectorStore(
	cfg *rawKnowledgeComponent,
	emb embedder.Embedder,
) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func knowledgeComponentType(node *yaml.Node) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func buildOpenAIKnowledgeEmbedder(
	node *yaml.Node,
) (embedder.Embedder, error) {
	_ = "STUB: not implemented"
	return *new(embedder.Embedder), nil
}

func buildInMemoryKnowledgeVectorStore(
	node *yaml.Node,
	_ knowledgeVectorStoreBuildContext,
) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func buildPGVectorKnowledgeVectorStore(
	node *yaml.Node,
	ctx knowledgeVectorStoreBuildContext,
) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func buildElasticsearchKnowledgeVectorStore(
	node *yaml.Node,
	ctx knowledgeVectorStoreBuildContext,
) (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}

func knowledgeEmbedderDimensions(e embedder.Embedder) int { _ = "STUB: not implemented"; return 0 }

func knowledgeToolName(name string) string { _ = "STUB: not implemented"; return "" }

type knowledgeIndexBaseTool struct {
	original  tool.Tool
	indexName string
}

func newKnowledgeIndexTool(indexName string, original tool.Tool) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func (t *knowledgeIndexBaseTool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"
	return nil
}

func (t *knowledgeIndexBaseTool) KnowledgeIndexName() string { _ = "STUB: not implemented"; return "" }

func (t *knowledgeIndexBaseTool) SkipSummarization() bool { _ = "STUB: not implemented"; return false }

type knowledgeIndexCallableTool struct {
	*knowledgeIndexBaseTool
}

func (t *knowledgeIndexCallableTool) Call(
	ctx context.Context,
	jsonArgs []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type knowledgeIndexStreamableTool struct {
	*knowledgeIndexBaseTool
}

func (t *knowledgeIndexStreamableTool) StreamableCall(
	ctx context.Context,
	jsonArgs []byte,
) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type knowledgeIndexCallableStreamableTool struct {
	*knowledgeIndexBaseTool
}

func (t *knowledgeIndexCallableStreamableTool) Call(
	ctx context.Context,
	jsonArgs []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *knowledgeIndexCallableStreamableTool) StreamableCall(
	ctx context.Context,
	jsonArgs []byte,
) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sanitizeKnowledgeToolSegment(name string) string { _ = "STUB: not implemented"; return "" }
