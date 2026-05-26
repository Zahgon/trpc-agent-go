//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package session provides the core session functionality.
package session

import (
	"context"
	"errors"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// StateMap is a map of state key-value pairs.
type StateMap map[string][]byte

var (
	// ErrAppNameRequired is the error for app name required.
	ErrAppNameRequired = errors.New("appName is required")
	// ErrUserIDRequired is the error for user id required.
	ErrUserIDRequired = errors.New("userID is required")
	// ErrSessionIDRequired is the error for session id required.
	ErrSessionIDRequired = errors.New("sessionID is required")
	// ErrNilSession is the error for session is nil.
	ErrNilSession = errors.New("session is nil")
	// ErrEventPageOnlyForGetSession indicates event paging is not supported by ListSessions.
	ErrEventPageOnlyForGetSession = errors.New("event page is only supported by GetSession")
	// ErrEventPageUnsupported indicates the backend does not support event paging.
	ErrEventPageUnsupported = errors.New("event page is only supported by postgres/mysql GetSession")
	// ErrInvalidEventPage indicates event paging arguments are invalid.
	ErrInvalidEventPage = errors.New("event page requires offset >= 0 and limit > 0")
	// ErrEventPageConflictsWithEventFilters indicates paging cannot be mixed with context filters.
	ErrEventPageConflictsWithEventFilters = errors.New("event page cannot be combined with EventNum or EventTime")
	// ErrInvalidListSessionPage indicates list-session paging arguments are invalid.
	ErrInvalidListSessionPage = errors.New("list session page requires offset >= 0 and limit >= 0")
)

// SummaryFilterKeyAllContents is the filter key representing
// the full-session summary with no filtering applied.
const SummaryFilterKeyAllContents = ""

// Session is the interface that all sessions must implement.
type Session struct {
	ID       string                 `json:"id"`      // ID is the session id.
	AppName  string                 `json:"appName"` // AppName is the app name.
	UserID   string                 `json:"userID"`  // UserID is the user id.
	State    StateMap               `json:"state"`   // State is the session state with delta support.
	Events   []event.Event          `json:"events"`  // Events is the session events.
	EventMu  sync.RWMutex           `json:"-"`
	Tracks   map[Track]*TrackEvents `json:"tracks,omitempty"` // Tracks stores track events.
	TracksMu sync.RWMutex           `json:"-"`
	// Summaries holds filter-aware summaries. The key is the event filter key.
	SummariesMu sync.RWMutex        `json:"-"`                   // SummariesMu is the read-write mutex for Summaries.
	Summaries   map[string]*Summary `json:"summaries,omitempty"` // Summaries is the filter-aware summaries.
	UpdatedAt   time.Time           `json:"updatedAt"`           // UpdatedAt is the last update time.
	CreatedAt   time.Time           `json:"createdAt"`           // CreatedAt is the creation time.

	// Hash is the pre-computed slot hash value for asynchronous task dispatching.
	// It is calculated once during session creation by hashing
	// "appName:userID:sessionID" and remains immutable throughout the session's lifecycle.
	// This field is computed once during session creation and never modified.
	Hash int `json:"-"`

	// ServiceMeta stores service-layer metadata (memory only, not persisted).
	// Used internally by session service implementations for version routing, etc.
	// Users should not access or modify this field directly.
	ServiceMeta map[string]string `json:"-"`

	stateMu sync.RWMutex `json:"-"` // stateMu is the read-write mutex for State.
}

// Clone returns a copy of the session.
func (sess *Session) Clone() *Session { _ = "STUB: not implemented"; return nil }

// Create new state to avoid reference sharing.

// Add missing CreatedAt field.

// Copy events.

// Copy track events.

// Copy state.

// Copy summaries.

// Shallow copy is fine since Summary is immutable after write.

// Copy service metadata.

// SessionOptions is the options for a session.
type SessionOptions func(*Session)

// WithSessionEvents is the option for the session events.
func WithSessionEvents(events []event.Event) SessionOptions {
	_ = "STUB: not implemented"
	return *new(SessionOptions)
}

// WithSessionSummaries is the option for the session summaries.
func WithSessionSummaries(summaries map[string]*Summary) SessionOptions {
	_ = "STUB: not implemented"
	return *new(SessionOptions)
}

// WithSessionState is the option for the session state.
func WithSessionState(state StateMap) SessionOptions {
	_ = "STUB: not implemented"
	return *new(SessionOptions)
}

// WithSessionCreatedAt is the option for the session createdAt.
func WithSessionCreatedAt(createdAt time.Time) SessionOptions {
	_ = "STUB: not implemented"
	return *new(SessionOptions)
}

// WithSessionUpdatedAt is the option for the session updatedAt.
func WithSessionUpdatedAt(updatedAt time.Time) SessionOptions {
	_ = "STUB: not implemented"
	return *new(SessionOptions)
}

// HashString computes a non-negative deterministic hash for the given string.
// It is used for slot-based dispatching of sessions and track events.
// The result is always >= 0, safe for use as a slice index after modulus.
func HashString(s string) int { _ = "STUB: not implemented"; return 0 }

// NewSession creates a new session.
func NewSession(appName, userID, sessionID string, options ...SessionOptions) *Session {
	_ = "STUB: not implemented"
	return nil
}

// GetState returns a copy of the state value for the given key.
// The returned slice is copied to avoid callers mutating shared memory.
func (sess *Session) GetState(key string) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SetState sets the state value for the given key.
// The provided slice is copied to avoid retaining caller-owned memory.
func (sess *Session) SetState(key string, value []byte) { _ = "STUB: not implemented"; return }

// DeleteState deletes a key from the session state.
func (sess *Session) DeleteState(key string) { _ = "STUB: not implemented"; return }

// SnapshotState returns a deep copy of the current session state.
// This is safe to iterate without holding locks.
func (sess *Session) SnapshotState() StateMap { _ = "STUB: not implemented"; return *new(StateMap) }

// HasStateKeyWithPrefix reports whether the session state contains at least one
// key with the provided prefix and a non-empty value.
//
// Nil or empty values are treated as absent.
func (sess *Session) HasStateKeyWithPrefix(prefix string) bool {
	_ = "STUB: not implemented"
	return false
}

// SnapshotTracksState returns a copy of the tracks state value (State["tracks"]).
// Returns nil if no tracks are registered.
func (sess *Session) SnapshotTracksState() []byte { _ = "STUB: not implemented"; return nil }

// GetEvents returns the session events.
func (sess *Session) GetEvents() []event.Event { _ = "STUB: not implemented"; return nil }

// GetEventCount returns the session event count.
func (sess *Session) GetEventCount() int { _ = "STUB: not implemented"; return 0 }

// AppendTrackEvent appends a track event to the session.
func (sess *Session) AppendTrackEvent(event *TrackEvent, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Track index is stored in session state; protect it with the state mutex.

// GetTrackEvents returns the track events snapshot.
func (sess *Session) GetTrackEvents(track Track) (*TrackEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnsureEventStartWithUser filters events to ensure they start with RoleUser.
// It removes events from the beginning until it finds the first event from RoleUser.
func (sess *Session) EnsureEventStartWithUser() { _ = "STUB: not implemented"; return }

// Find the first event that starts with RoleUser

// If event has no response or choices, continue to next event

// If no user event found, clear all events

// Keep events starting from the first user event

// UpdateUserSession updates the user session with the given event and options.
func (sess *Session) UpdateUserSession(event *event.Event, opts ...Option) {
	_ = "STUB: not implemented"
	return
}

// Apply filtering options.

// ApplyEventFiltering applies event number and time filtering to session events
// It ensures that the filtered events still contain at least one user message.
func (sess *Session) ApplyEventFiltering(opts ...Option) { _ = "STUB: not implemented"; return }

// Apply event time filter - keep events after the specified time

// No events after the specified time, clear all events

// Apply event number limit

// check if has user message

// find the last user message from original events

// ApplyEventStateDelta merges the state delta of the event into the session state.
func (sess *Session) ApplyEventStateDelta(e *event.Event) { _ = "STUB: not implemented"; return }

// Copy to avoid retaining caller-owned memory.

// ApplyEventStateDeltaMap merges the state delta of the event into the session state.
func ApplyEventStateDeltaMap(state StateMap, e *event.Event) { _ = "STUB: not implemented"; return }

// Copy to avoid retaining caller-owned memory.

func applyOptions(opts ...Option) *Options { _ = "STUB: not implemented"; return nil }

// Summary represents a concise, structured summary of a conversation branch.
// It is stored on the session object rather than in the StateMap.
type Summary struct {
	Summary   string    `json:"summary"`          // Summary is the concise conversation summary.
	Topics    []string  `json:"topics,omitempty"` // Topics is the optional topics list.
	UpdatedAt time.Time `json:"updated_at"`       // UpdatedAt is the update timestamp in UTC.
}

// Options contains shared session-service options.
// Not every field applies to every service method.
type Options struct {
	EventNum            int              // EventNum is the number of recent events (context-window mode).
	EventTime           time.Time        // EventTime is the after time.
	EventPage           *EventPage       // EventPage enables GetSession-only offset pagination when non-nil.
	ListSessionOnlyMeta bool             // ListSessionOnlyMeta is only honored by ListSessions.
	ListSessionPage     *ListSessionPage // ListSessionPage enables ListSessions offset pagination when non-nil.
}

// ListSessionPage specifies offset-based pagination for ListSessions.
// When non-nil in Options, the backend applies LIMIT/OFFSET to the session list.
type ListSessionPage struct {
	Offset int // Offset is the number of sessions to skip.
	Limit  int // Limit is the maximum number of sessions to return per page.
}

// EventPage specifies GetSession-only offset-based pagination for session events.
// When set, the backend returns a strict event page and skips ApplyEventFiltering.
// Offset counts from the most recent event (0 = most recent page).
type EventPage struct {
	Offset int // Offset is the number of most-recent events to skip.
	Limit  int // Limit is the maximum number of events to return per page.
}

// Option is the option for a session.
type Option func(*Options)

// WithEventNum is the option for the number of recent events.
func WithEventNum(num int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEventTime is the option for the time of the recent events.
func WithEventTime(time time.Time) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithListSessionOnlyMeta requests ListSessions to return only session metadata
// without events or tracks. Callers should only use this option with ListSessions.
func WithListSessionOnlyMeta() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithListSessionPage enables offset/limit pagination for ListSessions.
// This option is orthogonal to WithEventNum/WithEventTime: those control event-level
// filtering within each session, while this controls session-level pagination.
func WithListSessionPage(offset, limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGetSessionEventPage enables strict offset/limit pagination for GetSession events.
// offset counts backwards from the most recent event (0 = most recent page).
// limit is the page size. This option is only supported by postgres/mysql GetSession.
func WithGetSessionEventPage(offset, limit int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ValidateGetSessionOptions validates GetSession-only option semantics.
func ValidateGetSessionOptions(opts *Options, supportsEventPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateListSessionsOptions validates ListSessions-only option semantics.
func ValidateListSessionsOptions(opts *Options) error { _ = "STUB: not implemented"; return nil }

// SummaryOption is the option for getting session summary.
type SummaryOption func(*SummaryOptions)

// SummaryOptions is the options for getting session summary.
type SummaryOptions struct {
	// FilterKey specifies which filter's summary to retrieve.
	// When empty (SummaryFilterKeyAllContents), retrieves the full-session summary.
	FilterKey string
}

// WithSummaryFilterKey sets the filter key for summary retrieval.
// When empty (SummaryFilterKeyAllContents), retrieves the full-session summary.
// Use this option to get summaries for specific event filters (e.g., "user-messages").
func WithSummaryFilterKey(filterKey string) SummaryOption {
	_ = "STUB: not implemented"
	return *new(SummaryOption)
}

// SearchMode selects the retrieval strategy for session
// event search.
type SearchMode string

const (
	// SearchModeDense uses embedding similarity only.
	SearchModeDense SearchMode = "dense"
	// SearchModeHybrid combines dense similarity and
	// keyword search using rank fusion.
	SearchModeHybrid SearchMode = "hybrid"
)

// EventSearchRequest describes a session event search.
type EventSearchRequest struct {
	// Query is the user query to search against.
	Query string
	// UserKey scopes the search namespace. It is
	// required for all searches.
	UserKey UserKey
	// SessionIDs optionally restricts the search to
	// specific sessions under UserKey. When empty, all
	// sessions for the user are considered.
	SessionIDs []string
	// ExcludeSessionIDs removes specific sessions from
	// the search scope.
	ExcludeSessionIDs []string
	// MaxResults overrides the backend default result
	// count when > 0.
	MaxResults int
	// MinScore filters out low-confidence matches.
	MinScore float64
	// FilterKey restricts events using hierarchical
	// session/event branch semantics.
	FilterKey string
	// Roles restricts matches to specific message roles.
	Roles []model.Role
	// CreatedAfter restricts matches to events created on
	// or after this time.
	CreatedAfter *time.Time
	// CreatedBefore restricts matches to events created on
	// or before this time.
	CreatedBefore *time.Time
	// SearchMode selects the retrieval strategy. Empty
	// means SearchModeDense.
	SearchMode SearchMode
	// HybridRRFK controls the Reciprocal Rank Fusion
	// constant when SearchModeHybrid is used. When <= 0,
	// the backend default is used.
	HybridRRFK int
	// HybridCandidateRatio controls how many candidates
	// each hybrid branch fetches before fusion. When <= 0,
	// the backend default is used.
	HybridCandidateRatio int
}

// EventSearchResult wraps a recalled session event with
// its search metadata.
type EventSearchResult struct {
	// SessionKey identifies the matched session.
	SessionKey Key
	// SessionCreatedAt is the matched session creation
	// time, when available.
	SessionCreatedAt time.Time
	// EventCreatedAt is the persisted event row creation
	// time.
	EventCreatedAt time.Time
	// Event is the matched event payload.
	Event event.Event
	// Role is the normalized message role used for
	// indexing/search.
	Role model.Role
	// Text is the indexed text returned for prompt
	// injection or debugging.
	Text string
	// Score is the backend relevance score. For dense
	// search this is cosine similarity; for hybrid search
	// this is the fused rank score.
	Score float64
	// DenseScore is the dense retrieval contribution when
	// available.
	DenseScore float64
	// SparseScore is the sparse retrieval contribution
	// when available.
	SparseScore float64
}

// SearchableService extends session.Service with
// vector-based semantic search over session events.
// Session backends that support embedding-based retrieval
// should implement this interface.
// Non-vector backends (sqlite, mysql, redis, etc.) do not
// need to implement this interface.
type SearchableService interface {
	// SearchEvents returns the most relevant events for
	// the given request. Results are ordered by
	// descending relevance score.
	// Only events with meaningful text content
	// (user/assistant messages and tool results) are
	// searchable; tool calls and partial events are
	// excluded.
	SearchEvents(
		ctx context.Context,
		req EventSearchRequest,
	) ([]EventSearchResult, error)
}

// EventWindowRequest describes a request to load a
// small event window around one anchor event.
type EventWindowRequest struct {
	// Key identifies the target session.
	Key Key
	// AnchorEventID identifies the center event.
	AnchorEventID string
	// Before controls how many events before the anchor
	// are included. Negative values are rejected.
	Before int
	// After controls how many events after the anchor are
	// included. Negative values are rejected.
	After int
	// Roles optionally restrict the window to specific
	// message roles.
	Roles []model.Role
}

// EventWindowEntry stores one event plus its persisted
// creation timestamp.
type EventWindowEntry struct {
	Event     event.Event
	CreatedAt time.Time
}

// EventWindow contains the events loaded around one
// anchor event.
type EventWindow struct {
	SessionKey    Key
	AnchorEventID string
	Entries       []EventWindowEntry
}

// WindowService extends session.Service with precise
// anchor-based window loading.
type WindowService interface {
	// GetEventWindow returns a small ordered event window
	// around one anchor event.
	GetEventWindow(
		ctx context.Context,
		req EventWindowRequest,
	) (*EventWindow, error)
}

// Service is the interface that all session services must implement.
type Service interface {
	// CreateSession creates a new session.
	CreateSession(ctx context.Context, key Key, state StateMap, options ...Option) (*Session, error)

	// GetSession gets a session.
	GetSession(ctx context.Context, key Key, options ...Option) (*Session, error)

	// ListSessions lists all sessions by user scope of session key.
	// When WithListSessionOnlyMeta is provided, supported implementations omit events
	// and tracks and only return session metadata plus merged state.
	ListSessions(ctx context.Context, userKey UserKey, options ...Option) ([]*Session, error)

	// DeleteSession deletes a session.
	DeleteSession(ctx context.Context, key Key, options ...Option) error

	// UpdateAppState updates the state by target scope and key.
	UpdateAppState(ctx context.Context, appName string, state StateMap) error

	// DeleteAppState deletes the state by target scope and key.
	DeleteAppState(ctx context.Context, appName string, key string) error

	// GetState gets the state by target scope and key.
	ListAppStates(ctx context.Context, appName string) (StateMap, error)

	// UpdateUserState updates the state by target scope and key.
	UpdateUserState(ctx context.Context, userKey UserKey, state StateMap) error

	// GetUserState gets the state by target scope and key.
	ListUserStates(ctx context.Context, userKey UserKey) (StateMap, error)

	// DeleteUserState deletes the state by target scope and key.
	DeleteUserState(ctx context.Context, userKey UserKey, key string) error

	// UpdateSessionState updates the session-level state directly without appending an event.
	// This is useful for state initialization, correction, or synchronization scenarios
	// where event history is not needed.
	// Keys with app: or user: prefixes are not allowed (use UpdateAppState/UpdateUserState instead).
	// Keys with temp: prefix are allowed as they represent session-scoped ephemeral state.
	UpdateSessionState(ctx context.Context, key Key, state StateMap) error

	// AppendEvent appends an event to a session.
	AppendEvent(ctx context.Context, session *Session, event *event.Event, options ...Option) error

	// CreateSessionSummary triggers summarization for the session.
	// When filterKey is non-empty, implementations should limit work to the
	// matching branch using hierarchical rules consistent with event.Filter.
	// Implementations should preserve original events and store summaries on
	// the session object. The operation should be non-blocking for the main
	// flow where possible. Implementations may group deltas by branch internally.
	CreateSessionSummary(ctx context.Context, sess *Session, filterKey string, force bool) error

	// EnqueueSummaryJob enqueues a summary job for asynchronous processing.
	// This method provides a non-blocking way to trigger summary generation.
	// When async processing is enabled, the job will be processed by background workers.
	// When async processing is disabled or unavailable, it falls back to synchronous processing.
	// The method validates session parameters before enqueueing and returns appropriate errors.
	EnqueueSummaryJob(ctx context.Context, sess *Session, filterKey string, force bool) error

	// GetSessionSummaryText returns the latest summary text for the session if any.
	// The boolean indicates whether a summary exists.
	// When no options are provided, returns the full-session summary (SummaryFilterKeyAllContents).
	// Use WithSummaryFilterKey to specify a different filter key.
	GetSessionSummaryText(ctx context.Context, sess *Session, opts ...SummaryOption) (string, bool)

	// Close closes the service.
	Close() error
}

// Key is the key for a session.
type Key struct {
	AppName   string // app name
	UserID    string // user id
	SessionID string // session id
}

// CheckSessionKey checks if a session key is valid.
func (s *Key) CheckSessionKey() error { _ = "STUB: not implemented"; return nil }

// CheckUserKey checks if a user key is valid.
func (s *Key) CheckUserKey() error { _ = "STUB: not implemented"; return nil }

// UserKey is the key for a user.
type UserKey struct {
	AppName string // app name
	UserID  string // user id
}

// CheckUserKey checks if a user key is valid.
func (s *UserKey) CheckUserKey() error { _ = "STUB: not implemented"; return nil }

func checkSessionKey(appName, userID, sessionID string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkUserKey(appName, userID string) error { _ = "STUB: not implemented"; return nil }
