//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runtimeprofile

import (
	"context"
	"sync"
)

// Store loads runtime profile configuration from any backing system.
//
// Implementations can read from files, databases, config centers, or remote
// control planes. CachedResolver stays small so those stores do not need to
// understand runner internals.
type Store interface {
	Load(ctx context.Context) (Config, error)
}

// Catalog lists profile metadata without exposing full profile definitions.
type Catalog interface {
	ProfileIDs(ctx context.Context) ([]string, error)
	AppNames(ctx context.Context) ([]string, error)
}

// StoreFunc adapts a function to Store.
type StoreFunc func(ctx context.Context) (Config, error)

// Load implements Store.
func (f StoreFunc) Load(ctx context.Context) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// StaticStore stores one immutable profile config.
type StaticStore struct {
	Config Config
}

// Load implements Store.
func (s StaticStore) Load(context.Context) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// ProfileIDs implements Catalog.
func (s StaticStore) ProfileIDs(context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppNames implements Catalog.
func (s StaticStore) AppNames(context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CachedResolver lazily loads profiles and keeps a version-keyed resolver
// snapshot until Reload or Invalidate is called.
type CachedResolver struct {
	mu       sync.RWMutex
	store    Store
	resolver Resolver
	loaded   bool
}

// NewCachedResolver creates a resolver backed by a reloadable store.
func NewCachedResolver(store Store) *CachedResolver { _ = "STUB: not implemented"; return nil }

// Resolve implements Resolver.
func (r *CachedResolver) Resolve(
	ctx context.Context,
	req Request,
) (Profile, error) {
	_ = "STUB: not implemented"
	return *new(Profile), nil
}

// Reload refreshes the resolver snapshot from the store.
func (r *CachedResolver) Reload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Invalidate clears the current snapshot. The next Resolve will reload it.
func (r *CachedResolver) Invalidate() { _ = "STUB: not implemented"; return }

// ProfileIDs implements Catalog.
func (r *CachedResolver) ProfileIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppNames implements Catalog.
func (r *CachedResolver) AppNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CachedResolver) resolverForRead(
	ctx context.Context,
) (Resolver, error) {
	_ = "STUB: not implemented"
	return *new(Resolver), nil
}

func (r *CachedResolver) configForCatalog(
	ctx context.Context,
) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

func cloneConfig(cfg Config) Config { _ = "STUB: not implemented"; return *new(Config) }

func profileIDs(cfg Config) []string { _ = "STUB: not implemented"; return nil }

func appNames(cfg Config) []string { _ = "STUB: not implemented"; return nil }
