//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sessionopt provides internal pagination helpers for session options.
package sessionopt

import (
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// SortByUpdatedDesc sorts sessions by UpdatedAt descending, using ID descending
// as a tiebreaker for deterministic pagination results.
func SortByUpdatedDesc(sessions []*session.Session) { _ = "STUB: not implemented"; return }

// ApplyListPage applies offset/limit pagination based on Options.
// If ListSessionPage is nil, sessions are returned as-is.
func ApplyListPage(sessions []*session.Session, opt *session.Options) []*session.Session {
	_ = "STUB: not implemented"
	return nil
}
