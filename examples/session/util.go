//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package util provides utility functions for session examples.
package util

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"trpc.group/trpc-go/trpc-agent-go/event"
	openaiembedder "trpc.group/trpc-go/trpc-agent-go/knowledge/embedder/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// SessionType defines the type of session service.
type SessionType string

// Session type constants.
const (
	SessionInMemory   SessionType = "inmemory"
	SessionNoop       SessionType = "noop"
	SessionSQLite     SessionType = "sqlite"
	SessionRedis      SessionType = "redis"
	SessionPostgres   SessionType = "postgres"
	SessionPGVector   SessionType = "pgvector"
	SessionMySQL      SessionType = "mysql"
	SessionTDSQL      SessionType = "tdsql"
	SessionClickHouse SessionType = "clickhouse"
)

// SessionServiceConfig holds configuration for creating a session service.
type SessionServiceConfig struct {
	EventLimit       int
	TTL              time.Duration
	AppendEventHooks []session.AppendEventHook
	GetSessionHooks  []session.GetSessionHook
	EnableTracing    bool // enable OpenTelemetry tracing (redis only)
}

// NewSessionServiceByType creates a session service based on the specified
// type.
//
// Parameters:
//   - sessionType: one of inmemory, noop, sqlite, redis, postgres, pgvector,
//     mysql, clickhouse
//   - cfg: session service configuration (eventLimit, ttl, hooks)
//
// Environment variables by session type:
//
//	sqlite:     SQLITE_SESSION_DSN (default:
//	  file:sessions.db?_busy_timeout=5000)
//	redis:      REDIS_ADDR (default: localhost:6379)
//	postgres:   PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DATABASE
//	pgvector:   PGVECTOR_HOST, PGVECTOR_PORT, PGVECTOR_USER,
//	  PGVECTOR_PASSWORD, PGVECTOR_DATABASE, PGVECTOR_EMBEDDER_MODEL
//	mysql:      MYSQL_HOST, MYSQL_PORT, MYSQL_USER, MYSQL_PASSWORD,
//	  MYSQL_DATABASE
//	tdsql:      TDSQL_HOST, TDSQL_PORT, TDSQL_USER, TDSQL_PASSWORD,
//	  TDSQL_DATABASE
//	clickhouse: CLICKHOUSE_HOST, CLICKHOUSE_PORT, CLICKHOUSE_USER,
//	  CLICKHOUSE_PASSWORD, CLICKHOUSE_DATABASE
func NewSessionServiceByType(
	sessionType SessionType,
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

const (
	sqliteSessionDSNEnvKey    = "SQLITE_SESSION_DSN"
	defaultSQLiteSessionDBDSN = "file:sessions.db?_busy_timeout=5000"
	sqliteDriverName          = "sqlite3"
	defaultSQLiteMaxOpenConns = 1
	defaultSQLiteMaxIdleConns = 1

	openAIEmbeddingAPIKeyEnvKey  = "OPENAI_EMBEDDING_API_KEY"
	openAIEmbeddingBaseURLEnvKey = "OPENAI_EMBEDDING_BASE_URL"
	openAIEmbeddingModelEnvKey   = "OPENAI_EMBEDDING_MODEL"
)

func newSQLiteSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func newInMemorySessionService(cfg SessionServiceConfig) session.Service {
	_ = "STUB: not implemented"
	return *new(session.Service)
}

// newRedisSessionService creates a Redis session service.
// Environment variables:
//   - REDIS_ADDR: Redis server address (default: localhost:6379)
func newRedisSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

// newPostgresSessionService creates a PostgreSQL session service.
// Environment variables:
//   - PG_HOST: PostgreSQL host (default: localhost)
//   - PG_PORT: PostgreSQL port (default: 5432)
//   - PG_USER: PostgreSQL user (default: root)
//   - PG_PASSWORD: PostgreSQL password (default: empty)
//   - PG_DATABASE: PostgreSQL database (default: trpc_agent_go)
func newPostgresSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func getEmbeddingModel(defaultModel string) string { _ = "STUB: not implemented"; return "" }

func newOpenAIEmbedder(defaultModel string) *openaiembedder.Embedder {
	_ = "STUB: not implemented"
	return nil
}

// newPGVectorSessionService creates a PostgreSQL + pgvector
// backed session service.
// Environment variables:
//   - PGVECTOR_HOST: PostgreSQL host (default: localhost)
//   - PGVECTOR_PORT: PostgreSQL port (default: 5432)
//   - PGVECTOR_USER: PostgreSQL user (default: postgres)
//   - PGVECTOR_PASSWORD: PostgreSQL password (default: empty)
//   - PGVECTOR_DATABASE: PostgreSQL database (default:
//     trpc-agent-go-pgsession)
//   - PGVECTOR_EMBEDDER_MODEL: Embedder model name (default:
//     text-embedding-3-small)
func newPGVectorSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

// newMySQLSessionService creates a MySQL session service.
// Environment variables:
//   - MYSQL_HOST: MySQL host (default: localhost)
//   - MYSQL_PORT: MySQL port (default: 3306)
//   - MYSQL_USER: MySQL user (default: root)
//   - MYSQL_PASSWORD: MySQL password (default: empty)
//   - MYSQL_DATABASE: MySQL database (default: trpc_agent_go)
func newMySQLSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

// newTDSQLSessionService creates a TDSQL (distributed MySQL) session service.
// TDSQL requires shardkey-aware DDL/DML; WithTDSQLSharding enables the
// necessary table definitions and query routing.
// Environment variables:
//   - TDSQL_HOST: TDSQL proxy host (default: localhost)
//   - TDSQL_PORT: TDSQL proxy port (default: 3306)
//   - TDSQL_USER: TDSQL user (default: root)
//   - TDSQL_PASSWORD: TDSQL password (default: empty)
//   - TDSQL_DATABASE: TDSQL database (default: trpc_agent_go)
func newTDSQLSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

// newClickHouseSessionService creates a ClickHouse session service.
// Environment variables:
//   - CLICKHOUSE_HOST: ClickHouse host (default: localhost)
//   - CLICKHOUSE_PORT: ClickHouse native port (default: 9000)
//   - CLICKHOUSE_USER: ClickHouse user (default: default)
//   - CLICKHOUSE_PASSWORD: ClickHouse password (default: empty)
//   - CLICKHOUSE_DATABASE: ClickHouse database (default: trpc_agent_go)
func newClickHouseSessionService(
	cfg SessionServiceConfig,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

// RunnerConfig holds configuration for creating a runner.
type RunnerConfig struct {
	AppName     string
	AgentName   string
	ModelName   string
	Instruction string
	Tools       []tool.Tool
	MaxTokens   int
	Temperature *float64
	Streaming   bool
}

// DefaultRunnerConfig returns a default runner configuration.
func DefaultRunnerConfig() RunnerConfig { _ = "STUB: not implemented"; return *new(RunnerConfig) }

// NewRunner creates a runner with the given session service and
// configuration.
func NewRunner(
	sessionService session.Service,
	cfg RunnerConfig,
) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

// RunAgent runs the agent with the given message and optionally prints
// the conversation.
func RunAgent(
	ctx context.Context,
	r runner.Runner,
	userID string,
	sessionID string,
	message string,
	printConversation bool,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractResponse extracts the response content from an event.
func ExtractResponse(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

// GetEnvOrDefault retrieves the value of an environment variable or returns a
// default value if not set.
func GetEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

// Truncate truncates a string to maxLen characters.
func Truncate(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }

// IntPtr returns a pointer to an int.
func IntPtr(v int) *int {
	_ = "STUB: not implemented"

	// FloatPtr returns a pointer to a float64.
	return nil
}

func FloatPtr(v float64) *float64 {
	_ = "STUB: not implemented"

	// PrintSessionEvents prints all events for a session in debug mode.
	// It retrieves the session from the service and prints each event's
	// role and content.
	return nil
}

func PrintSessionEvents(
	ctx context.Context,
	svc session.Service,
	appName string,
	userID string,
	sessionID string,
) error {
	_ = "STUB: not implemented"
	return nil
}
