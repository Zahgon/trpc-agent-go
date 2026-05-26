//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package elasticsearch provides Elasticsearch-based vector storage implementation.
package elasticsearch

import (
	"encoding/json"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
)

type esDocument map[string]json.RawMessage

func (es esDocument) stringField(key string) string { _ = "STUB: not implemented"; return "" }

func (es esDocument) mapField(key string) map[string]any { _ = "STUB: not implemented"; return nil }

func (es esDocument) sliceField(key string) []float64 { _ = "STUB: not implemented"; return nil }

func (es esDocument) timeField(key string) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (vs *VectorStore) docBuilder(hitSource json.RawMessage) (*document.Document, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Parse the _source field using our unified esDocument struct.

// Create document.
