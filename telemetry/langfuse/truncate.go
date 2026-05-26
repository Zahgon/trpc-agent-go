//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package langfuse

import (
	"sync/atomic"
)

const defaultTruncateMarker = "…[truncated]…"

// observationMaxBytes stores the configured truncation threshold.
//
// Semantics:
// - < 0: truncation disabled
// - = 0: truncate everything
// - > 0: max byte length for a JSON leaf node (or a plain string value)
var observationMaxBytes atomic.Int64

func init() {
	// Default: truncation disabled unless configured.
	observationMaxBytes.Store(-1)
}

func setObservationMaxBytes(maxBytes *int) { _ = "STUB: not implemented"; return }

// getObservationMaxBytes returns the max byte length for each observation JSON leaf node.
func getObservationMaxBytes() int { _ = "STUB: not implemented"; return 0 }

// truncateObservationValue limits the size of Langfuse observation input/output.
//
// It is intentionally simple (scheme-agnostic): apply to the final string value.
// Truncation is disabled by default unless configured.
func truncateObservationValue(s string) string { _ = "STUB: not implemented"; return "" }

func truncateStringBytes(s string, maxBytes int) string { _ = "STUB: not implemented"; return "" }

// Prefer keeping more head than tail, but always keep both sides.

// Best-effort validity: should already be valid due to boundary trimming.

func safeUTF8Prefix(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

// Back up from a UTF-8 continuation byte.

func safeUTF8Suffix(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

// Advance to a rune boundary (skip UTF-8 continuation bytes).
