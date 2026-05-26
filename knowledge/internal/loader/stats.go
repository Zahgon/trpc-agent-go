//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package loader contains internal helper utilities for concurrent knowledge
// base loading. These helpers are internal-only and not part of the public
// API.
package loader

// This file provides Stats, a small helper to collect document-size statistics.

// Stats tracks statistics for document sizes during a load run.
// It is shared by both source-level and load-level reporting paths.
type Stats struct {
	TotalDocs  int
	TotalSize  int
	MinSize    int
	MaxSize    int
	bucketCnts []int
}

// NewStats returns a Stats initialised for the provided buckets.
func NewStats(buckets []int) *Stats {
	_ = "STUB: not implemented"
	// Initialise with max-int for MinSize.
	return nil
}

// Add records the size of a document.
func (s *Stats) Add(size int, buckets []int) { _ = "STUB: not implemented"; return }

// Avg returns the average document size.
func (s *Stats) Avg() float64 { _ = "STUB: not implemented"; return 0 }

// Log outputs the collected statistics.
func (s *Stats) Log(buckets []int) { _ = "STUB: not implemented"; return }
