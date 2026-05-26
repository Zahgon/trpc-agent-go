//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package knowledge provides internal knowledge utilities.
package knowledge

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

// ApplyPreprocess applies a chain of transformers' Preprocess to documents.
// Returns the original documents if no transformers are provided.
func ApplyPreprocess(docs []*document.Document, transformers ...transform.Transformer) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyPostprocess applies a chain of transformers' Postprocess to documents.
// Returns the original documents if no transformers are provided.
func ApplyPostprocess(docs []*document.Document, transformers ...transform.Transformer) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
