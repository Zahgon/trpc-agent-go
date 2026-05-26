//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqlite

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// UpdateSessionState updates the session-level state without appending an
// event. Keys with app: or user: prefixes are not allowed.
func (s *Service) UpdateSessionState(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeState(
	appState session.StateMap,
	userState session.StateMap,
	sess *session.Session,
) *session.Session {
	_ = "STUB: not implemented"
	return nil
}
