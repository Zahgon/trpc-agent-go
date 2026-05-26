//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package arxiv

// GetShortID get short id from entry id
func (r *Result) GetShortID() string { _ = "STUB: not implemented"; return "" }

// GetDefaultFilename get default filename from result
func (r *Result) GetDefaultFilename(extension string) string { _ = "STUB: not implemented"; return "" }

// GetSourceURL get source url from result
func (r *Result) GetSourceURL() string { _ = "STUB: not implemented"; return "" }

// ArxivError arxiv api error
type ArxivError struct {
	URL     string
	Retry   int
	Message string
}

// Error error
func (e *ArxivError) Error() string { _ = "STUB: not implemented"; return "" }

// UnexpectedEmptyPageError unexpected empty page error
type UnexpectedEmptyPageError struct {
	ArxivError
}

// HTTPError http error
type HTTPError struct {
	ArxivError
	Status int
}

// NewSearch create new search instance
func NewSearch(query string, options ...SearchOption) Search {
	_ = "STUB: not implemented"
	return *new(Search)
}

// SearchOption search option
type SearchOption func(*Search)

// WithIDList set id list
func WithIDList(ids ...string) SearchOption { _ = "STUB: not implemented"; return *new(SearchOption) }

// WithMaxResults set max results
func WithMaxResults(max int) SearchOption { _ = "STUB: not implemented"; return *new(SearchOption) }

// WithSortBy set sort by
func WithSortBy(sortBy SortCriterion) SearchOption {
	_ = "STUB: not implemented"
	return *new(SearchOption)
}

// WithSortOrder set sort order
func WithSortOrder(sortOrder SortOrder) SearchOption {
	_ = "STUB: not implemented"
	return *new(SearchOption)
}

// DefaultClient return default client
func DefaultClient() *Client { _ = "STUB: not implemented"; return nil }
