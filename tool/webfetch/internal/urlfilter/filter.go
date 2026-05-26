//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package urlfilter provides URL filtering functions.
package urlfilter

import (
	"net/url"
)

// urlFilter is a function that determines if a URL should be allowed.
// It returns true if the URL is allowed, false otherwise.
type urlFilter func(url string) bool

// URLValidator combines a filter with an error message.
type URLValidator struct {
	Filter urlFilter
	ErrMsg string
}

// CheckURL checks the URL against the configured validators.
func CheckURL(validators []URLValidator, urlStr string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewBlockPatternFilter creates a filter that blocks URLs matching the pattern.
func NewBlockPatternFilter(pattern string) urlFilter {
	_ = "STUB: not implemented"
	return *new(urlFilter)
}

// Fail-safe: treat unparsable URLs as blocked

// NewAllowPatternsFilter creates a filter that allows URLs matching any of the patterns.
func NewAllowPatternsFilter(patterns []string) urlFilter {
	_ = "STUB: not implemented"
	return *new(urlFilter)
}

// matchPattern checks if the URL matches the given pattern (host + path prefix).
func matchPattern(u *url.URL, pattern string) bool {
	_ = "STUB: not implemented"
	// Split pattern into host and path
	// Pattern is expected to be like "example.com" or "example.com/foo"
	return false
}

// 1. Host match (case-insensitive)

// 2. Path match

// Normalize URL path

// Ensure absolute path comparison if pattern starts with / (which it does from split)

// Boundary check to avoid "/doc" matching "/docserver"
// Match if:
// - lengths are equal (exact match)
// - pattern ends with '/' (explicit directory match)
// - next char in uPath is '/' (sub-path match)

// matchHost checks if hostname matches target domain (exact or suffix).
// e.g., matchHost("www.example.com", "example.com") -> true
func matchHost(hostname, target string) bool { _ = "STUB: not implemented"; return false }
