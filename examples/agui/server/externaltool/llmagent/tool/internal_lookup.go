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

const internalLookupName = "internal_lookup"

func newInternalLookupTool() agenttool.Tool { _ = "STUB: not implemented"; return *new(agenttool.Tool) }

func internalLookup(_ context.Context, args internalLookupArgs) (internalLookupResult, error) {
	_ = "STUB: not implemented"
	return *new(internalLookupResult), nil
}

type internalLookupArgs struct {
	Subject string `json:"subject" description:"The subject to look up in server-side context."`
}

type internalLookupResult struct {
	Result string `json:"result" description:"The deterministic server-side lookup result."`
}
