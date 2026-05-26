//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
)

const skillsLegacyMigrationStateKey = "processor:skills:legacy_migrated"

func maybeMigrateLegacySkillState(
	ctx context.Context,
	inv *agent.Invocation,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

type scopedKeyBuilder func(agentName string, skillName string) string

func migrateLegacyStateKey(
	inv *agent.Invocation,
	delta map[string][]byte,
	legacyKey string,
	legacyVal []byte,
	skillName string,
	owners map[string]string,
	buildKey scopedKeyBuilder,
) {
	_ = "STUB: not implemented"
	return
}

func legacySkillOwners(events []event.Event) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func addOwnersFromEvent(
	ev event.Event,
	owners map[string]string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
