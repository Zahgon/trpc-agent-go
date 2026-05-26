//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telemetry

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// EmbeddingAttributes represents the attributes of an embedding call.
type EmbeddingAttributes struct {
	RequestEncodingFormat *string
	RequestModel          string
	Dimensions            int
	Error                 error
	InputToken            *int64
	Request               *string
	Response              *string
	ServerAddress         *string
	ServerPort            *int
}

// TraceEmbedding traces the invocation of an embedding call.
func TraceEmbedding(span trace.Span, embeddingAttributes *EmbeddingAttributes) {
	_ = "STUB: not implemented"
	return
}

func buildEmbeddingAttributes(embeddingAttributes *EmbeddingAttributes) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}
