//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversationscope

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

const storageUserStatePrefix = "openclaw.conversation_storage_user:"
const storageScopeStatePrefix = "openclaw.conversation_storage_scope:"

var storageUserStateValue = []byte("1")

// RememberIndexedStorageUser records one storage user scope seen for the
// canonical user so cleanup flows can enumerate persisted conversation scopes.
func RememberIndexedStorageUser(
	ctx context.Context,
	svc session.Service,
	appName string,
	canonicalUserID string,
	storageUserID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListIndexedStorageUsers lists extra persisted storage scopes remembered for
// the canonical user.
func ListIndexedStorageUsers(
	ctx context.Context,
	svc session.Service,
	appName string,
	canonicalUserID string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RememberIndexedStorageScope records one persisted conversation scope at
// app level so admin surfaces can enumerate known chats.
func RememberIndexedStorageScope(
	ctx context.Context,
	svc session.Service,
	appName string,
	storageUserID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListIndexedStorageScopes lists persisted conversation scopes remembered
// for the application.
func ListIndexedStorageScopes(
	ctx context.Context,
	svc session.Service,
	appName string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteIndexedStorageUser removes one remembered storage scope index.
func DeleteIndexedStorageUser(
	ctx context.Context,
	svc session.Service,
	appName string,
	canonicalUserID string,
	storageUserID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func indexedStorageUsersFromState(state session.StateMap) []string {
	_ = "STUB: not implemented"
	return nil
}

func indexedStorageScopesFromState(state session.StateMap) []string {
	_ = "STUB: not implemented"
	return nil
}

func storageUserStateKey(storageUserID string) string { _ = "STUB: not implemented"; return "" }

func storageScopeStateKey(storageUserID string) string { _ = "STUB: not implemented"; return "" }
