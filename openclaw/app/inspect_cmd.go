//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

const subcmdInspect = "inspect"

const (
	inspectCmdPlugins    = "plugins"
	inspectCmdConfigKeys = "config-keys"
	inspectCmdDeps       = "deps"
)

const (
	registryKindChannel        = "channel"
	registryKindSessionBackend = "session backend"
	registryKindMemoryBackend  = "memory backend"
	registryKindToolProvider   = "tool provider"
	registryKindToolSet        = "toolset provider"
	registryKindModel          = "model"
)

func runInspect(args []string) int { _ = "STUB: not implemented"; return 0 }

func runInspectPlugins(args []string) int { _ = "STUB: not implemented"; return 0 }

func runInspectConfigKeys(args []string) int { _ = "STUB: not implemented"; return 0 }

func printInspectList(title string, items []string) { _ = "STUB: not implemented"; return }

func printInspectUsage() { _ = "STUB: not implemented"; return }
