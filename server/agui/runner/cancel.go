//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package runner

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
)

// Canceler cancels a running run identified by the request payload.
type Canceler interface {
	// Cancel cancels a running run and returns an error when the run key cannot be found.
	Cancel(ctx context.Context, input *adapter.RunAgentInput) error
}

// Cancel cancels a running run identified by appName, userID, and sessionID.
func (r *runner) Cancel(ctx context.Context, runAgentInput *adapter.RunAgentInput) error {
	_ = "STUB: not implemented"
	return nil
}
