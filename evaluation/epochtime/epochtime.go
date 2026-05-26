//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package epochtime provides EpochTime type.
package epochtime

import (
	"time"
)

const (
	// zeroEpochLiteral is the literal for zero epoch.
	zeroEpochLiteral = "0"
	// nanosecondsPerSecond is the number of nanoseconds per second.
	nanosecondsPerSecond = float64(time.Second)
)

// EpochTime wraps time.Time to (un)marshal as unix seconds (float).
type EpochTime struct{ time.Time }

// MarshalJSON implements json.Marshaler to encode time as unix seconds (float).
func (t EpochTime) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler to decode unix seconds (float).
func (t *EpochTime) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
