//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package toolsearch

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type knowledgeSearcher struct {
	model         model.Model
	systemPrompt  string
	toolKnowledge *ToolKnowledge
}

const defaultSystemPromptWithToolKnowledge = "Your goal is to identify the most relevant tools for answering the user's query. " +
	"Provide a natural-language description of the kind of tool needed (e.g., 'weather information', 'currency conversion', 'stock prices')."

func newKnowledgeSearcher(m model.Model, systemPrompt string, toolKnowledge *ToolKnowledge) *knowledgeSearcher {
	_ = "STUB: not implemented"
	return nil
}

func (s *knowledgeSearcher) Search(ctx context.Context, candidates map[string]tool.Tool, query string, topK int) (newCtx context.Context, selectedTools []string, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func (s *knowledgeSearcher) rewriteQuery(ctx context.Context, query string) (context.Context, string, *model.Usage, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), "", nil, nil
}
