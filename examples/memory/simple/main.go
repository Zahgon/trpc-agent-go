//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates simple memory management using the Runner with
// streaming output, session management, and manual memory tool calling.
//
// Usage:
//
//	go run main.go -memory=inmemory
//	go run main.go -memory=sqlite
//	go run main.go -memory=sqlitevec
//	go run main.go -memory=redis
//	go run main.go -memory=mysql
//	go run main.go -memory=mysqlvec
//	go run main.go -memory=postgres
//	go run main.go -memory=pgvector
//
// Environment variables by memory type (example usage):
//
//	sqlite:
//		export SQLITE_MEMORY_DSN="file:memories.db?_busy_timeout=5000"
//
//	sqlitevec:
//		export SQLITEVEC_MEMORY_DSN="file:memories_vec.db?_busy_timeout=5000"
//		export SQLITEVEC_EMBEDDER_MODEL="text-embedding-3-small"
//
//	redis:
//		export REDIS_ADDR="localhost:6379"
//
//	mysql:
//		export MYSQL_HOST="localhost"
//		export MYSQL_PORT="3306"
//		export MYSQL_USER="root"
//		export MYSQL_PASSWORD=""
//		export MYSQL_DATABASE="trpc_agent_go"
//
//	mysqlvec:
//		export MYSQLVEC_HOST="localhost"
//		export MYSQLVEC_PORT="3306"
//		export MYSQLVEC_USER="root"
//		export MYSQLVEC_PASSWORD=""
//		export MYSQLVEC_DATABASE="trpc_agent_go"
//		export MYSQLVEC_EMBEDDER_MODEL="text-embedding-3-small"
//
//	postgres:
//		export PG_HOST="localhost"
//		export PG_PORT="5432"
//		export PG_USER="postgres"
//		export PG_PASSWORD=""
//		export PG_DATABASE="trpc_agent_go"
//
//	pgvector:
//		export PGVECTOR_HOST="localhost"
//		export PGVECTOR_PORT="5432"
//		export PGVECTOR_USER="postgres"
//		export PGVECTOR_PASSWORD=""
//		export PGVECTOR_DATABASE="trpc_agent_go"
//		export PGVECTOR_EMBEDDER_MODEL="text-embedding-3-small"
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/runner"

	util "trpc.group/trpc-go/trpc-agent-go/examples/memory"
)

var (
	modelName = flag.String(
		"model",
		"deepseek-v4-flash",
		"Name of the model to use",
	)
	memServiceName = flag.String(
		"memory",
		"inmemory",
		"Name of the memory service to use, "+
			"inmemory / sqlite / sqlitevec / redis / "+
			"mysql / mysqlvec / postgres / pgvector",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for responses",
	)
	softDelete = flag.Bool(
		"soft-delete",
		false,
		"Enable soft delete for SQLite/SQLiteVec/"+
			"MySQL/PostgreSQL/pgvector memory service",
	)
)

func main() {
	flag.Parse()

	fmt.Printf("🧠 Simple Memory Chat\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Memory Service: %s\n", *memServiceName)

	memoryType := util.MemoryType(*memServiceName)
	util.PrintMemoryInfo(memoryType, *softDelete)

	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Available tools: %s\n", util.GetAvailableToolsString())
	fmt.Println(strings.Repeat("=", 50))

	chat := &memoryChat{
		modelName:      *modelName,
		memServiceName: *memServiceName,
		streaming:      *streaming,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

type memoryChat struct {
	modelName      string
	memServiceName string
	streaming      bool
	runner         runner.Runner
	memoryService  memory.Service
	userID         string
	sessionID      string
}

func (c *memoryChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *memoryChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *memoryChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *memoryChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *memoryChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *memoryChat) hasToolCalls(event *event.Event) bool { _ = "STUB: not implemented"; return false }

func (c *memoryChat) hasToolResponses(event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *memoryChat) handleToolCalls(event *event.Event, assistantStarted bool) {
	_ = "STUB: not implemented"
	return
}

func (c *memoryChat) handleToolResponses(event *event.Event) { _ = "STUB: not implemented"; return }

func (c *memoryChat) extractContent(event *event.Event) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *memoryChat) startNewSession() { _ = "STUB: not implemented"; return }
