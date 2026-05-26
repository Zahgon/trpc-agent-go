//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package team

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

// Export exports the static structure of the team.
func (t *Team) Export(
	ctx context.Context,
	exportChild structure.ChildExporter,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportCoordinatorTeam(
	ctx context.Context,
	exportChild structure.ChildExporter,
	snapshot *structure.Snapshot,
	rootNodeID string,
	coordinator agent.Agent,
	members []agent.Agent,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportSwarmTeam(
	ctx context.Context,
	exportChild structure.ChildExporter,
	snapshot *structure.Snapshot,
	rootNodeID string,
	entryName string,
	members []agent.Agent,
) (*structure.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
