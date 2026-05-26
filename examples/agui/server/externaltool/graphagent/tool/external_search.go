//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import (
	"context"

	agenttool "trpc.group/trpc-go/trpc-agent-go/tool"
)

// ExternalSearchName is the caller-executed external search tool name.
const ExternalSearchName = "external_search"

func newExternalSearchTool() agenttool.Tool { _ = "STUB: not implemented"; return *new(agenttool.Tool) }

func externalSearchNotImplemented(context.Context, externalSearchArgs) (externalSearchResult, error) {
	_ = "STUB: not implemented"
	return *new(externalSearchResult), nil
}

type externalSearchArgs struct {
	Query string `json:"query" description:"The search query."`
}

type externalSearchResult struct {
	Result string `json:"result" description:"The tool result content."`
}
