//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// defaultMaxDedupKeysPerInvocation caps the number of unique code chunk keys
// tracked per invocation. When the cap is exceeded, the oldest-inserted keys
// are evicted. This is a safety net against pathological prompts that loop
// through a code search tool many times within a single user turn.
const defaultMaxDedupKeysPerInvocation = 1024

const codeDedupRuntimeStateKey = "tool:code_search:dedup_state"

// codeDedupStore keeps a per-invocation set of already-seen code chunk keys.
// It stores the per-turn state on the invocation itself, so the dedup cache
// naturally disappears when the invocation finishes instead of accumulating on
// the long-lived tool instance.
type codeDedupStore struct {
	mu sync.Mutex
	// maxKeys caps the number of unique keys tracked per invocation. When
	// <= 0, defaultMaxDedupKeysPerInvocation is used.
	maxKeys int
}

type dedupEntry struct {
	mu   sync.Mutex
	keys map[string]struct{}
	// order tracks insertion order so that we can evict the oldest keys when
	// the per-invocation cap is reached.
	order []string
}

func newCodeDedupStore() *codeDedupStore { _ = "STUB: not implemented"; return nil }

// newCodeDedupStoreWithCap creates a dedup store with a custom per-invocation
// key cap. Non-positive values fall back to defaultMaxDedupKeysPerInvocation.
func newCodeDedupStoreWithCap(maxKeys int) *codeDedupStore { _ = "STUB: not implemented"; return nil }

// filter removes documents whose dedup key was already returned in previous
// calls under the same invocation.
//
// Return contract:
//   - When all documents are new: resp.Documents is kept as-is and Message is
//     rewritten to "Found N relevant document(s)" style text.
//   - When some documents are duplicates: only the new documents are kept in
//     resp.Documents and Message reports how many duplicates were omitted.
//   - When every top result was already returned in previous calls: the
//     response is NOT converted to an error. Instead, resp.Documents is set to
//     an empty (non-nil) slice and Message is rewritten to explicitly tell the
//     LLM that all top results were duplicates and that it should vary the
//     query/filter/repo before searching again. Callers must therefore handle
//     the empty-documents case without treating it as a failure.
func (s *codeDedupStore) filter(ctx context.Context, resp *KnowledgeSearchResponse) *KnowledgeSearchResponse {
	_ = "STUB: not implemented"
	return nil
}

// No invocation context means we cannot safely scope the dedup set;
// fall back to a no-op rather than leaking state across requests.

// Evict the oldest key to bound memory usage.

func (s *codeDedupStore) loadOrCreate(invocation *agent.Invocation) *dedupEntry {
	_ = "STUB: not implemented"
	return nil
}

// codeDedupKey returns a stable key identifying the underlying code chunk.
//
// The repository name (trpc_ast_repo_name) is always prepended to every key
// when present, so that the same full_name / file_path in different
// repositories is never mistakenly deduplicated against each other.
//
// Preference order (after the optional repo prefix):
//  1. AST fully-qualified symbol name (most specific for code entities).
//  2. file_path + line range (covers chunks that have no full_name but still
//     map to a unique span, e.g. markdown or raw text chunks).
//  3. file_path alone (coarse but non-empty for file-level documents).
//
// Returns an empty string when none of the keys are present, in which case
// the caller should keep the document without deduplication.
func codeDedupKey(doc *DocumentResult) string { _ = "STUB: not implemented"; return "" }

func stringFromMeta(md map[string]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// scalarFromMeta stringifies simple scalar values (int/float/string) so that
// they can be composed into a deterministic dedup key. Non-scalar values fall
// back to fmt.Sprint which is still deterministic for our purposes.
func scalarFromMeta(md map[string]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// JSON numbers decode as float64; keep int form when possible.

// Use the shortest representation that preserves full float64 precision
// so that two distinct values never collapse onto the same dedup key.
