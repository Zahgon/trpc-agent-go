//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skill

const stateKeyScopeDelimiter = "/"

func escapeScopeSegment(value string) string { _ = "STUB: not implemented"; return "" }

// LoadedKey returns the session state key used to mark a skill as loaded for
// a specific agent.
//
// When agentName is empty, it falls back to the legacy unscoped key.
func LoadedKey(agentName string, skillName string) string { _ = "STUB: not implemented"; return "" }

// DocsKey returns the session state key used to store doc selection for a
// specific agent.
//
// When agentName is empty, it falls back to the legacy unscoped key.
func DocsKey(agentName string, skillName string) string { _ = "STUB: not implemented"; return "" }

// LoadedPrefix returns the prefix used to scan loaded-skill keys for the
// provided agentName.
//
// When agentName is empty, it returns the legacy prefix.
func LoadedPrefix(agentName string) string { _ = "STUB: not implemented"; return "" }

// DocsPrefix returns the prefix used to scan doc-selection keys for the
// provided agentName.
//
// When agentName is empty, it returns the legacy prefix.
func DocsPrefix(agentName string) string { _ = "STUB: not implemented"; return "" }
