//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

//go:build cgo && sqliteveccgo

package util

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

const (
	sqliteVecDSNEnvKey           = "SQLITEVEC_DSN"
	defaultSQLiteVecDSN          = "file:knowledge.sqlite?_busy_timeout=5000"
	sqliteVecTableEnvKey         = "SQLITEVEC_TABLE"
	defaultSQLiteVecTable        = "knowledge_documents"
	sqliteVecMetadataTableEnvKey = "SQLITEVEC_METADATA_TABLE"
	defaultSQLiteVecMetadata     = "knowledge_document_meta"
)

func newSQLiteVecStore() (vectorstore.VectorStore, error) {
	_ = "STUB: not implemented"
	return *new(vectorstore.VectorStore), nil
}
