//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mem0

import (
	"net/url"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	metadataKeyTRPCTopics       = "trpc_topics"
	metadataKeyTRPCAppName      = "trpc_app_name"
	metadataKeyTRPCKind         = "trpc_kind"
	metadataKeyTRPCEventTime    = "trpc_event_time"
	metadataKeyTRPCParticipants = "trpc_participants"
	metadataKeyTRPCLocation     = "trpc_location"

	pathV1Memories = "/v1/memories/"
	pathV2Search   = "/v2/memories/search/"

	queryKeyUserID   = "user_id"
	queryKeyAppID    = "app_id"
	queryKeyPage     = "page"
	queryKeyPageSize = "page_size"

	memoryUserRole = "user"

	defaultListPageSize = 100
	defaultSearchTopK   = 20
)

func addOrgProjectQuery(q url.Values, opts serviceOpts) { _ = "STUB: not implemented"; return }

func addOrgProjectFilter(filters map[string]any, opts serviceOpts) {
	_ = "STUB: not implemented"
	return
}

func parseMem0Times(rec *memoryRecord) parsedTimes {
	_ = "STUB: not implemented"
	return *new(parsedTimes)
}

func parseMem0Time(s string) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func toEntry(appName, userID string, rec *memoryRecord) *memory.Entry {
	_ = "STUB: not implemented"
	return nil
}

func readTopicsFromMetadata(meta map[string]any) []string { _ = "STUB: not implemented"; return nil }

func readKindFromMetadata(meta map[string]any) memory.Kind {
	_ = "STUB: not implemented"
	return *new(memory.Kind)
}

func readEventTimeFromMetadata(meta map[string]any) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

func readParticipantsFromMetadata(meta map[string]any) []string {
	_ = "STUB: not implemented"
	return nil
}

func readLocationFromMetadata(meta map[string]any) string { _ = "STUB: not implemented"; return "" }

func messageText(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func matchesSearchFilters(entry *memory.Entry, opts memory.SearchOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func sortSearchResults(results []*memory.Entry, opts memory.SearchOptions) {
	_ = "STUB: not implemented"
	return
}

func searchCandidateLimit(opts memory.SearchOptions, maxResults int) int {
	_ = "STUB: not implemented"
	return 0
}

func isInvalidPageError(err error) bool { _ = "STUB: not implemented"; return false }

// cloneMetadata returns a deep clone of meta with no aliased nested state.
//
// Ingestion runs asynchronously on a worker goroutine, so the outer map and
// any nested containers must be independent of the caller's memory: otherwise
// a caller mutating its metadata map after IngestSession has returned could
// race with, or change, the payload the worker eventually marshals. The clone
// round-trips through JSON because the metadata is ultimately transmitted to
// mem0 as JSON — so the canonicalization is lossless with respect to what the
// backend actually receives.
func cloneMetadata(meta map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// Metadata that cannot be serialized would also fail downstream when
// the worker builds the createMemoryRequest payload; drop it here so
// the ingest payload simply omits metadata rather than aliasing the
// caller's map.
