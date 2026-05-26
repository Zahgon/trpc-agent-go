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

// ExternalApprovalName is the caller-executed approval tool name.
const ExternalApprovalName = "external_approval"

func newExternalApprovalTool() agenttool.Tool {
	_ = "STUB: not implemented"
	return *new(agenttool.Tool)
}

func externalApprovalNotImplemented(_ context.Context, args externalApprovalArgs) (externalApprovalResult, error) {
	_ = "STUB: not implemented"
	return *new(externalApprovalResult), nil
}

type externalApprovalArgs struct {
	Item string `json:"item" description:"The item that needs caller approval."`
}

type externalApprovalResult struct {
	Decision string `json:"decision" description:"The approval decision returned by the caller."`
}
