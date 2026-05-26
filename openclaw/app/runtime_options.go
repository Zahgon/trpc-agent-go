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
	"trpc.group/trpc-go/trpc-agent-go/openclaw/runtimeprofile"
)

// RuntimeOption customizes an embedded OpenClaw runtime.
type RuntimeOption func(*runtimeOptions)

type runtimeOptions struct {
	runtimeProfileResolver runtimeprofile.Resolver
	runtimeProfileCatalog  runtimeprofile.Catalog
	runtimeProfileRequired bool
}

// WithRuntimeProfileResolver injects per-request runtime profile resolution.
func WithRuntimeProfileResolver(
	resolver runtimeprofile.Resolver,
	required bool,
) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// WithRuntimeProfileCatalog injects profile metadata for cleanup/catalog use.
func WithRuntimeProfileCatalog(
	catalog runtimeprofile.Catalog,
) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// WithRuntimeProfileStore injects a reloadable runtime profile store.
//
// Callers that need Reload or Invalidate control can create a
// runtimeprofile.CachedResolver and pass WithRuntimeProfileResolver.
func WithRuntimeProfileStore(
	store runtimeprofile.Store,
	required bool,
) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

func buildRuntimeOptions(options []RuntimeOption) runtimeOptions {
	_ = "STUB: not implemented"
	return *new(runtimeOptions)
}
