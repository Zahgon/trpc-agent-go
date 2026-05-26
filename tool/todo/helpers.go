//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package todo

import (
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// stateKey builds the session.State key used to persist the checklist
// for the given branch. An empty branch yields just the prefix, which
// is the main (parent) agent's list.
//
// Each branch of an agent tree owns an independent list automatically,
// so parent and child agents never clobber each other's checklists.
//
// The key layout is a package-internal detail. External readers should
// use GetTodos; they do not need to know how keys are assembled.
func stateKey(prefix, branch string) string { _ = "STUB: not implemented"; return "" }

// GetTodos loads the current checklist for the given branch from the
// session. The returned slice may be nil (no list written yet, empty
// session, or the list was cleared after all items completed) or a
// populated slice decoded from state; callers should rely on len()
// instead of a nil check. An error is only returned if stored data
// is present but corrupt.
//
// Use the empty string for branch to read the main agent's list.
func GetTodos(sess *session.Session, branch string) ([]Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTodosWithPrefix is the advanced form of GetTodos that honours a
// custom state key prefix (see WithStateKeyPrefix). Return semantics
// match GetTodos.
//
// An empty prefix is treated as a request for the default layout and
// falls back to DefaultStateKeyPrefix — it is never read as a literal
// empty-prefix state key. Pass a concrete value if you need a custom
// layout, or use GetTodos if you do not.
func GetTodosWithPrefix(sess *session.Session, prefix, branch string) ([]Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readTodos is the package-internal variant of GetTodosWithPrefix that
// accepts a pre-computed state key. A poisoned state entry surfaces as
// a decode error and is the caller's responsibility to discard —
// Call() does this with `oldTodos, _ = readTodos(...)` so that a
// single bad write cannot stop the next one from landing.
func readTodos(sess *session.Session, key string) ([]Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
