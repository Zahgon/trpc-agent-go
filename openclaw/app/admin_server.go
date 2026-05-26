//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"net"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/session"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/cron"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/octool"
	ocskills "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/skills"
)

const adminAutoPortSearchSpan = 32

type adminBinding struct {
	listener  net.Listener
	addr      string
	url       string
	relocated bool
}

func openAdminBinding(
	addr string,
	autoPort bool,
) (*adminBinding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isAddressInUse(err error) bool { _ = "STUB: not implemented"; return false }

func buildAdminConfig(
	opts runOptions,
	agentType string,
	instanceID string,
	langfuse admin.LangfuseStatus,
	stateDir string,
	debugDir string,
	startedAt time.Time,
	channels []channel.Channel,
	routes admin.Routes,
	cronSvc *cron.Service,
	execMgr *octool.Manager,
	promptController *RuntimePromptController,
	browserManaged admin.BrowserManagedStatusProvider,
	adminAddr string,
	adminURL string,
	skillsRepo *ocskills.Repository,
	skillsWatch *ocskills.WatchService,
	memoryFiles admin.MemoryFileStore,
	sessionSvc session.Service,
) admin.Config {
	_ = "STUB: not implemented"
	return *new(admin.Config)
}

func buildBrowserAdminConfig(
	specs []pluginSpec,
	managed admin.BrowserManagedStatusProvider,
) admin.BrowserConfig {
	_ = "STUB: not implemented"
	return *new(admin.BrowserConfig)
}

func runtimeHostname() string { _ = "STUB: not implemented"; return "" }

func adminModelName(
	opts runOptions,
	agentType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

type adminSkillsProvider struct {
	mu           sync.RWMutex
	configPath   string
	repo         *ocskills.Repository
	watch        *ocskills.WatchService
	roots        []string
	bundledRoot  string
	configKeys   []string
	allowBundled []string
	skillConfigs map[string]ocskills.SkillConfig
}

func buildAdminSkillsProvider(
	opts runOptions,
	stateDir string,
	repo *ocskills.Repository,
	watch *ocskills.WatchService,
) admin.SkillsStatusProvider {
	_ = "STUB: not implemented"
	return *new(admin.SkillsStatusProvider)
}
