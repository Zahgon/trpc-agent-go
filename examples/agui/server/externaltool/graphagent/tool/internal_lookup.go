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

// InternalLookupName is the graph-executed internal lookup tool name.
const InternalLookupName = "internal_lookup"

func newInternalLookupTool() agenttool.Tool { _ = "STUB: not implemented"; return *new(agenttool.Tool) }

// InternalLookup returns deterministic internal context for the requested query.
func InternalLookup(_ context.Context, args InternalLookupArgs) (InternalLookupResult, error) {
	_ = "STUB: not implemented"
	return *new(InternalLookupResult), nil
}

// InternalLookupArgs is the argument schema for internal_lookup.
type InternalLookupArgs struct {
	Query string `json:"query" description:"The internal lookup query."`
}

// InternalLookupResult is the result schema for internal_lookup.
type InternalLookupResult struct {
	Result string `json:"result" description:"The internal lookup result."`
}
