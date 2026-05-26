//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package recall provides agent-facing tools for on-demand
// session history search and loading.
package recall

import (
	"context"
	"errors"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	// SearchToolName searches session history.
	SearchToolName = "session_search"
	// LoadToolName loads a small history window around one anchor result.
	LoadToolName = "session_load"

	// ScopeCurrentHidden searches summarized-away history in the current
	// session.
	ScopeCurrentHidden = "current_hidden"
	// ScopeCurrentSession searches the current session regardless of summary
	// cutoff. Use this when current-session details may have been compacted
	// out of the projected request.
	ScopeCurrentSession = "current_session"
	// ScopeOtherSessions searches other sessions owned by the same user.
	ScopeOtherSessions = "other_sessions"
	// ScopeAllSessions searches both current hidden history and other
	// sessions.
	ScopeAllSessions = "all_sessions"

	defaultSearchTopK    = 5
	maxSearchTopK        = 10
	defaultWindowBefore  = 1
	defaultWindowAfter   = 1
	maxWindowSpan        = 4
	maxSessionScanEvents = 10000
	searchExpandedHits   = 2
	searchSnippetBefore  = 2
	searchSnippetAfter   = 2
)

var (
	errInvocationContextRequired = errors.New("no invocation context found")
	errSessionRequired           = errors.New("invocation exists but no session available")
	errSearchUnavailable         = errors.New("session search is not available for this session service")
	errWindowUnavailable         = errors.New("session window loading is not available for this session service")
)

const summaryLastIncludedTsKey = "summary:last_included_ts"

// SearchSessionRequest is the input for session_search.
type SearchSessionRequest struct {
	Query      string             `json:"query" jsonschema:"description=Search query for prior conversation details. Prefer short keyword-style queries."`
	Scope      string             `json:"scope,omitempty" jsonschema:"description=Search scope: current_hidden, current_session, other_sessions, or all_sessions."`
	TopK       int                `json:"top_k,omitempty" jsonschema:"description=Maximum number of results to return. Defaults to 5 and is capped."`
	MinScore   float64            `json:"min_score,omitempty" jsonschema:"description=Optional minimum relevance score threshold between 0 and 1."`
	SearchMode session.SearchMode `json:"search_mode,omitempty" jsonschema:"description=Retrieval mode: dense or hybrid. Defaults to hybrid."`
}

// SearchSessionHit is one session_search result.
type SearchSessionHit struct {
	Scope     string                 `json:"scope"`
	SessionID string                 `json:"session_id"`
	EventID   string                 `json:"event_id"`
	Created   time.Time              `json:"created"`
	Role      model.Role             `json:"role,omitempty"`
	Score     float64                `json:"score"`
	Snippet   string                 `json:"snippet"`
	Context   []LoadedSessionMessage `json:"context,omitempty"`
}

// SearchSessionResponse is the response from session_search.
type SearchSessionResponse struct {
	Query   string             `json:"query"`
	Scope   string             `json:"scope"`
	Results []SearchSessionHit `json:"results"`
	Count   int                `json:"count"`
}

// LoadSessionRequest is the input for session_load.
type LoadSessionRequest struct {
	SessionID string `json:"session_id,omitempty" jsonschema:"description=Target session ID returned by session_search. Defaults to the current session when omitted."`
	EventID   string `json:"event_id" jsonschema:"description=Anchor event ID returned by session_search."`
	Before    int    `json:"before,omitempty" jsonschema:"description=How many messages before the anchor to include. Defaults to 1."`
	After     int    `json:"after,omitempty" jsonschema:"description=How many messages after the anchor to include. Defaults to 1."`
}

// LoadedSessionMessage is one historical message returned by session_load.
type LoadedSessionMessage struct {
	EventID string     `json:"event_id"`
	Role    model.Role `json:"role"`
	Created time.Time  `json:"created"`
	Content string     `json:"content"`
}

// LoadSessionResponse is the response from session_load.
type LoadSessionResponse struct {
	SessionID string                 `json:"session_id"`
	EventID   string                 `json:"event_id"`
	Before    int                    `json:"before"`
	After     int                    `json:"after"`
	Note      string                 `json:"note"`
	Messages  []LoadedSessionMessage `json:"messages"`
	Count     int                    `json:"count"`
}

// SupportsSearch reports whether session_search can be offered for this invocation.
func SupportsSearch(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

// SupportsLoad reports whether session_load can be offered for this invocation.
func SupportsLoad(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

// SupportsOnDemandSession reports whether both search and load are available.
func SupportsOnDemandSession(inv *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

func invocationFromContext(
	ctx context.Context,
) (*agent.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func searchableServiceFromContext(
	ctx context.Context,
) (session.SearchableService, *agent.Invocation, error) {
	_ = "STUB: not implemented"
	return *new(session.SearchableService), nil, nil
}

func windowServiceFromContext(
	ctx context.Context,
) (session.WindowService, *agent.Invocation, error) {
	_ = "STUB: not implemented"
	return *new(session.WindowService), nil, nil
}

func optionalWindowServiceFromInvocation(
	inv *agent.Invocation,
) session.WindowService {
	_ = "STUB: not implemented"
	return *new(session.WindowService)
}

func currentUserKey(
	inv *agent.Invocation,
) (session.UserKey, error) {
	_ = "STUB: not implemented"
	return *new(session.UserKey), nil
}

func currentSessionKey(
	inv *agent.Invocation,
	sessionID string,
) (session.Key, error) {
	_ = "STUB: not implemented"
	return *new(session.Key), nil
}

func normalizeScope(scope string) string { _ = "STUB: not implemented"; return "" }

func normalizeSearchMode(mode session.SearchMode) session.SearchMode {
	_ = "STUB: not implemented"
	return *new(session.SearchMode)
}

func normalizeTopK(topK int) int { _ = "STUB: not implemented"; return 0 }

func normalizeWindowSize(
	before, after int,
) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func currentSummaryCutoff(
	inv *agent.Invocation,
) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func extractSessionMessageText(
	evt event.Event,
) (string, model.Role, bool) {
	_ = "STUB: not implemented"
	return "", *new(model.Role), false
}

func loadedMessagesFromWindow(
	window *session.EventWindow,
) []LoadedSessionMessage {
	_ = "STUB: not implemented"
	return nil
}
