//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package util provides utility functions for memory examples.
package util

import (
	"time"

	openaiembedder "trpc.group/trpc-go/trpc-agent-go/knowledge/embedder/openai"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// MemoryType defines the type of memory service.
type MemoryType string

// Memory type constants.
const (
	MemoryInMemory  MemoryType = "inmemory"
	MemorySQLite    MemoryType = "sqlite"
	MemorySQLiteVec MemoryType = "sqlitevec"
	MemoryRedis     MemoryType = "redis"
	MemoryPostgres  MemoryType = "postgres"
	MemoryPGVector  MemoryType = "pgvector"
	MemoryMySQL     MemoryType = "mysql"
	MemoryMySQLVec  MemoryType = "mysqlvec"
)

// MemoryServiceConfig holds configuration for creating a memory service.
type MemoryServiceConfig struct {
	// Soft delete configuration.
	SoftDelete bool
	// Extractor configuration for auto memory mode.
	Extractor extractor.MemoryExtractor
	// Async memory worker configuration.
	AsyncMemoryNum   int
	MemoryQueueSize  int
	MemoryJobTimeout time.Duration
}

// RunnerConfig holds configuration for creating a runner.
type RunnerConfig struct {
	AppName     string
	AgentName   string
	ModelName   string
	Instruction string
	MaxTokens   int
	Temperature float64
	Streaming   bool
}

// DefaultRunnerConfig returns a default runner configuration.
func DefaultRunnerConfig() RunnerConfig { _ = "STUB: not implemented"; return *new(RunnerConfig) }

// NewMemoryServiceByType creates a memory service based on the specified type.
//
// This function supports both manual memory mode and auto memory mode:
// - Manual mode: cfg.Extractor == nil, uses explicit memory tool calls
// - Auto mode: cfg.Extractor != nil, automatically extracts memories from conversations
//
// Parameters:
//   - memoryType: one of inmemory, sqlite, sqlitevec, redis, postgres,
//     pgvector, mysql, mysqlvec
//   - cfg: memory service configuration
//   - SoftDelete: enable soft delete for SQL backends
//   - Extractor: memory extractor for auto mode (nil = manual mode)
//   - AsyncMemoryNum: number of async workers for auto mode (default 1)
//   - MemoryQueueSize: queue size for memory jobs in auto mode (default 10)
//   - MemoryJobTimeout: timeout for each memory job in auto mode (default 30s)
//
// Environment variables by memory type:
//
//	sqlite:     SQLITE_MEMORY_DSN (default: file:memories.db?_busy_timeout=5000)
//	sqlitevec:  SQLITEVEC_MEMORY_DSN (default: file:memories_vec.db?_busy_timeout=5000)
//	redis:      REDIS_ADDR (default: localhost:6379)
//	postgres:   PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DATABASE
//	pgvector:   PGVECTOR_HOST, PGVECTOR_PORT, PGVECTOR_USER, PGVECTOR_PASSWORD, PGVECTOR_DATABASE, PGVECTOR_EMBEDDER_MODEL
//	mysql:      MYSQL_HOST, MYSQL_PORT, MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE
//	mysqlvec:   MYSQLVEC_HOST, MYSQLVEC_PORT, MYSQLVEC_USER, MYSQLVEC_PASSWORD, MYSQLVEC_DATABASE, MYSQLVEC_EMBEDDER_MODEL
func NewMemoryServiceByType(memoryType MemoryType, cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

const (
	sqliteMemoryDSNEnvKey     = "SQLITE_MEMORY_DSN"
	defaultSQLiteMemoryDBDSN  = "file:memories.db?_busy_timeout=5000"
	sqliteDriverName          = "sqlite3"
	defaultSQLiteMaxOpenConns = 1
	defaultSQLiteMaxIdleConns = 1
)

func newSQLiteMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

const (
	sqliteVecMemoryDSNEnvKey    = "SQLITEVEC_MEMORY_DSN"
	defaultSQLiteVecMemoryDBDSN = "file:memories_vec.db?_busy_timeout=5000"

	sqliteVecEmbedderModelEnvKey = "SQLITEVEC_EMBEDDER_MODEL"

	openAIEmbeddingAPIKeyEnvKey  = "OPENAI_EMBEDDING_API_KEY"
	openAIEmbeddingBaseURLEnvKey = "OPENAI_EMBEDDING_BASE_URL"
	openAIEmbeddingModelEnvKey   = "OPENAI_EMBEDDING_MODEL"
)

func getEmbeddingModel(defaultModel string) string { _ = "STUB: not implemented"; return "" }

func newOpenAIEmbedder(defaultModel string) *openaiembedder.Embedder {
	_ = "STUB: not implemented"
	return nil
}

// newInMemoryMemoryService creates an in-memory memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
func newInMemoryMemoryService(cfg MemoryServiceConfig) memory.Service {
	_ = "STUB: not implemented"
	return *new(memory.Service)
}

// Configure extractor for auto memory mode if provided.

// newRedisMemoryService creates a Redis memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
// Environment variables:
//   - REDIS_ADDR: Redis server address (default: localhost:6379)
func newRedisMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

// Configure extractor for auto memory mode if provided.

// newPostgresMemoryService creates a PostgreSQL memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
// Environment variables:
//   - PG_HOST: PostgreSQL host (default: localhost)
//   - PG_PORT: PostgreSQL port (default: 5432)
//   - PG_USER: PostgreSQL user (default: postgres)
//   - PG_PASSWORD: PostgreSQL password (default: empty)
//   - PG_DATABASE: PostgreSQL database (default: trpc-agent-go-pgmemory)
func newPostgresMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

// Configure extractor for auto memory mode if provided.

// newPGVectorMemoryService creates a pgvector memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
// Environment variables:
//   - PGVECTOR_HOST: PostgreSQL host (default: localhost)
//   - PGVECTOR_PORT: PostgreSQL port (default: 5432)
//   - PGVECTOR_USER: PostgreSQL user (default: postgres)
//   - PGVECTOR_PASSWORD: PostgreSQL password (default: empty)
//   - PGVECTOR_DATABASE: PostgreSQL database (default: trpc-agent-go-pgmemory)
//   - PGVECTOR_EMBEDDER_MODEL: Embedder model name (default: text-embedding-3-small)
func newPGVectorMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

// Create embedder - for simplicity, we'll use OpenAI embedder

// Configure extractor for auto memory mode if provided.

// newMySQLMemoryService creates a MySQL memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
// Environment variables:
//   - MYSQL_HOST: MySQL host (default: localhost)
//   - MYSQL_PORT: MySQL port (default: 3306)
//   - MYSQL_USER: MySQL user (default: root)
//   - MYSQL_PASSWORD: MySQL password (default: empty)
//   - MYSQL_DATABASE: MySQL database (default: trpc_agent_go)
func newMySQLMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

// Configure extractor for auto memory mode if provided.

// newMySQLVecMemoryService creates a MySQL vector memory service.
// Supports both manual mode (cfg.Extractor == nil) and auto mode (cfg.Extractor != nil).
// Environment variables:
//   - MYSQLVEC_HOST: MySQL host (default: localhost)
//   - MYSQLVEC_PORT: MySQL port (default: 3306)
//   - MYSQLVEC_USER: MySQL user (default: root)
//   - MYSQLVEC_PASSWORD: MySQL password (default: empty)
//   - MYSQLVEC_DATABASE: MySQL database (default: trpc_agent_go)
//   - MYSQLVEC_EMBEDDER_MODEL: Embedder model name (default: text-embedding-3-small)
func newMySQLVecMemoryService(cfg MemoryServiceConfig) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

// Configure extractor for auto memory mode if provided.

// NewRunner creates a runner with the given memory service and configuration.
func NewRunner(memoryService memory.Service, cfg RunnerConfig) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

// GetEnvOrDefault retrieves the value of an environment variable or returns a default value if not set.
func GetEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

// IntPtr returns a pointer to an int.
func IntPtr(v int) *int {
	_ = "STUB: not implemented"

	// FloatPtr returns a pointer to a float64.
	return nil
}

func FloatPtr(v float64) *float64 {
	_ = "STUB: not implemented"

	// PrintMemoryInfo prints memory service information based on type.
	return nil
}

func PrintMemoryInfo(memoryType MemoryType, softDelete bool) { _ = "STUB: not implemented"; return }

// GetAvailableToolsString returns a string describing available memory tools.
func GetAvailableToolsString() string { _ = "STUB: not implemented"; return "" }

// FormatToolCalls formats tool calls for display.
func FormatToolCalls(toolCalls []model.ToolCall) string { _ = "STUB: not implemented"; return "" }

// FormatToolResponses formats tool responses for display.
func FormatToolResponses(choices []model.Choice) string { _ = "STUB: not implemented"; return "" }
