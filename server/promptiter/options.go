//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package promptiter

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	promptitermanager "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/manager"
)

const (
	defaultBasePath      = "/promptiter/v1/apps"
	defaultStructurePath = "/structure"
	defaultRunsPath      = "/runs"
	defaultAsyncRunsPath = "/async-runs"
)

// Option configures the PromptIter server.
type Option func(*options)

type options struct {
	appName                string
	basePath               string
	structurePath          string
	runsPath               string
	asyncRunsPath          string
	timeout                time.Duration
	engine                 engine.Engine
	manager                promptitermanager.Manager
	responseResultSlimming engine.RunResultSlimming
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithAppName sets the app name exposed by the PromptIter server.
func WithAppName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBasePath sets the base collection path used by the PromptIter server.
func WithBasePath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStructurePath sets the structure endpoint path relative to BasePath/appName.
func WithStructurePath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunsPath sets the runs endpoint path relative to BasePath/appName.
func WithRunsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAsyncRunsPath sets the asynchronous runs endpoint path relative to BasePath/appName.
func WithAsyncRunsPath(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout sets the maximum execution time for a PromptIter run request.
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEngine sets the PromptIter engine used by the server.
func WithEngine(promptIterEngine engine.Engine) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithManager sets the PromptIter manager used by the server.
func WithManager(promptIterManager promptitermanager.Manager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithResponseResultSlimming omits selected fields from run response payloads.
func WithResponseResultSlimming(slimming engine.RunResultSlimming) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
