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
	"encoding/json"
)

type storageUserIDContextKey struct{}

// WithStorageUserID records the storage user scope for the current request.
func WithStorageUserID(ctx context.Context, userID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// StorageUserIDFromContext resolves the persisted conversation user scope.
func StorageUserIDFromContext(ctx context.Context, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

// ResolveStorageUserID extracts an explicit storage user override from request
// extensions and falls back to the canonical request user when absent.
func ResolveStorageUserID(
	extensions map[string]json.RawMessage,
	fallback string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
