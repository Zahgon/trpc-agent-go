//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package parallelagent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

// Export exports the static structure of the parallel agent.
func (a *ParallelAgent) Export(
	ctx context.Context,
	exportChild structure.ChildExporter,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
