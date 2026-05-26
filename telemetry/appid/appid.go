//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package appid records and provides default app/agent names.
// It captures the first Runner's app and agent names as defaults and
// offers safe getters for fallback in observability or reporting paths.
package appid

import (
	"sync"
)

// Info holds a pair of application and agent names.
type Info struct {
	App   string
	Agent string
}

var (
	once  sync.Once
	first Info
	mu    sync.RWMutex
	seen  = make(map[string]struct{})
)

// pairSepStr is the string form of pairSepByte used when joining.
const pairSepStr string = "\x1f"

// RegisterRunner records a Runner's app and agent names.
// The first call defines the defaults used by DefaultApp/DefaultAgent.
func RegisterRunner(appName, agentName string) {
	_ = "STUB: not implemented"
	// Set the first runner info once.
	return
}

// Track the pair for optional diagnostics.

// DefaultApp returns the first Runner's app or the process name.
func DefaultApp() string { _ = "STUB: not implemented"; return "" }

// DefaultAgent returns the first Runner's agent or the process name.
func DefaultAgent() string { _ = "STUB: not implemented"; return "" }

// Runners returns a snapshot of all seen app/agent pairs.
func Runners() []Info { _ = "STUB: not implemented"; return nil }

// k is app pairSepStr agent

// procName returns the current process basename.
func procName() string { _ = "STUB: not implemented"; return "" }
