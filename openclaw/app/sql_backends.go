//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	openaiembedder "trpc.group/trpc-go/trpc-agent-go/knowledge/embedder/openai"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/session"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

type sqlSessionConfig struct {
	DSN        string `yaml:"dsn,omitempty"`
	Instance   string `yaml:"instance,omitempty"`
	SkipDBInit bool   `yaml:"skip_db_init,omitempty"`
	TablePref  string `yaml:"table_prefix,omitempty"`
}

func newMySQLSessionBackend(
	deps registry.SessionDeps,
	spec registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newPostgresSessionBackend(
	deps registry.SessionDeps,
	spec registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newClickHouseSessionBackend(
	deps registry.SessionDeps,
	spec registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

type sqlMemoryConfig struct {
	DSN        string `yaml:"dsn,omitempty"`
	Instance   string `yaml:"instance,omitempty"`
	TableName  string `yaml:"table_name,omitempty"`
	Schema     string `yaml:"schema,omitempty"`
	SkipDBInit bool   `yaml:"skip_db_init,omitempty"`
	SoftDelete *bool  `yaml:"soft_delete,omitempty"`
}

func newMySQLMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func newPostgresMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

type pgvectorMemoryConfig struct {
	sqlMemoryConfig `yaml:",inline"`

	IndexDimension int `yaml:"index_dimension,omitempty"`
	MaxResults     int `yaml:"max_results,omitempty"`

	Embedder *openAIEmbedderConfig `yaml:"embedder,omitempty"`
}

type openAIEmbedderConfig struct {
	Type         string `yaml:"type,omitempty"`
	Model        string `yaml:"model,omitempty"`
	Dimensions   int    `yaml:"dimensions,omitempty"`
	APIKey       string `yaml:"api_key,omitempty"`
	Organization string `yaml:"organization,omitempty"`
	BaseURL      string `yaml:"base_url,omitempty"`
}

func newPGVectorMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func newOpenAIEmbedder(
	cfg *openAIEmbedderConfig,
) (*openaiembedder.Embedder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func safeOption[O any](mk func(string) O, value string) (opt O, err error) {
	_ = "STUB: not implemented"
	return *new(O), nil
}
