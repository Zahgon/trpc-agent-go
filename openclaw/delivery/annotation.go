//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package delivery

import (
	"encoding/json"
)

const extensionKey = "openclaw:delivery_target:v1"

// Target stores one channel-specific default outbound destination.
type Target struct {
	Channel string `json:"channel,omitempty"`
	Target  string `json:"target,omitempty"`
}

// MergeRequestExtension stores delivery metadata in request extensions.
func MergeRequestExtension(
	extensions map[string]json.RawMessage,
	target Target,
) (map[string]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TargetFromRequestExtensions decodes request delivery metadata.
func TargetFromRequestExtensions(
	extensions map[string]json.RawMessage,
) (Target, bool, error) {
	_ = "STUB: not implemented"
	return *new(Target), false, nil
}

func sanitizeTarget(target Target) Target { _ = "STUB: not implemented"; return *new(Target) }

func isZeroTarget(target Target) bool { _ = "STUB: not implemented"; return false }
