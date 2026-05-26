//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"sync"
	"time"
)

// Cache-related string constants.
const (
	// CacheNamespacePrefix is the prefix for per-node cache namespaces.
	CacheNamespacePrefix = "__writes__"
)

// Cache is a minimal interface for storing and retrieving node results.
// Implementations must be concurrency-safe.
type Cache interface {
	// Get returns the cached value for the given namespace and key.
	// ok is true when a non-expired entry was found.
	Get(ns, key string) (val any, ok bool)
	// Set stores a value for the given namespace and key with the provided TTL.
	// ttl<=0 means no expiration.
	Set(ns, key string, val any, ttl time.Duration)
	// Clear removes all entries under the namespace.
	Clear(ns string)
}

// CachePolicy configures how cache keys are derived and how long entries live.
type CachePolicy struct {
	// KeyFunc derives a stable key bytes from the task input.
	// The implementation should ensure deterministic output for equivalent inputs.
	KeyFunc func(input any) ([]byte, error)
	// TTL controls entry lifetime. TTL<=0 means no expiration.
	TTL time.Duration
}

// DefaultCachePolicy returns a best-effort default policy using canonical JSON
// to produce a stable hash of the sanitized input.
func DefaultCachePolicy() *CachePolicy { _ = "STUB: not implemented"; return nil }

// InMemoryCache is a simple TTL cache backed by a nested map.
// It is intended for single-process use and testing.
type InMemoryCache struct {
	mu   sync.RWMutex
	data map[string]map[string]cacheEntry // ns -> key -> entry
}

type cacheEntry struct {
	v   any
	exp time.Time // zero means no expiration
}

// NewInMemoryCache creates a new in-memory cache instance.
func NewInMemoryCache() *InMemoryCache { _ = "STUB: not implemented"; return nil }

// Get implements Cache.
func (c *InMemoryCache) Get(ns, key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// lazy expire

// Return a deep copy to avoid shared references.

// Set implements Cache.
func (c *InMemoryCache) Set(ns, key string, val any, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Store a deep copy to isolate cache storage from caller mutations.

// Clear implements Cache.
func (c *InMemoryCache) Clear(ns string) { _ = "STUB: not implemented"; return }

// buildCacheNamespace builds a per-node namespace for cache entries.
// We scope by node ID only to align with the executor's node semantics.
func buildCacheNamespace(nodeID string) string { _ = "STUB: not implemented"; return "" }
