//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package content

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/tool"
)

var knowledgeTools = map[string]struct{}{
	"knowledge_search":                     {},
	"knowledge_search_with_agentic_filter": {},
}

// ExtractKnowledgeRecall builds a human-readable summary of knowledge tool responses.
func ExtractKnowledgeRecall(tools []*evalset.Tool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parseKnowledgeSearchResponse converts a function response payload into a typed knowledge search response.
func parseKnowledgeSearchResponse(t *evalset.Tool) (*tool.KnowledgeSearchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// var res tool.KnowledgeSearchResponse
