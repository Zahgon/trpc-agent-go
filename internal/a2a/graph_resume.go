//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package a2a

import (
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

// DecodeStateDeltaString decodes a JSON-encoded string value from state delta.
func DecodeStateDeltaString(stateDelta map[string][]byte, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// DecodeStateDeltaAny decodes a JSON-encoded value from state delta.
func DecodeStateDeltaAny(stateDelta map[string][]byte, key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// DecodeStateDeltaAnyMap decodes a JSON-encoded map value from state delta.
func DecodeStateDeltaAnyMap(stateDelta map[string][]byte, key string) (map[string]any, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// DecodePregelMetadata decodes PregelStepMetadata from the _pregel_metadata
// key in state delta.
func DecodePregelMetadata(stateDelta map[string][]byte) (graph.PregelStepMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(graph.PregelStepMetadata), false
}

// CloneAnyMap returns a shallow copy of a map[string]any.
func CloneAnyMap(src map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// GraphResumeStateFromStateDelta extracts graph resume state (lineage,
// checkpoint, resume command) from an encoded state_delta metadata payload.
func GraphResumeStateFromStateDelta(stateDeltaRaw any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Also extract checkpoint info from _pregel_metadata if the flat keys
// above are absent. The interrupt event already encodes lineage/checkpoint
// inside PregelStepMetadata via the normal state_delta path.

// GraphResumeStateFromMetadata extracts graph resume state from A2A message
// metadata. It prefers state_delta for checkpoint and resume data, then fills
// in any missing pieces from flattened metadata for backward compatibility.
func GraphResumeStateFromMetadata(metadata map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Backward compatibility for flattened fields.

// Fallback: extract resume info from a serialized Command struct that
// arrived via transferStateKey("*"). After JSON round-trip the
// *graph.Command becomes map[string]any with capitalized field names.

// extractResumeFromCommandMetadata tries to recover resume fields from a
// JSON-deserialized Command struct stored under StateKeyCommand in metadata.
func extractResumeFromCommandMetadata(metadata map[string]any, cmd *graph.ResumeCommand) bool {
	_ = "STUB: not implemented"
	return false
}
