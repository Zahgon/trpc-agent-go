//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package promptcore provides the internal text rendering engine behind the
// public prompt package.
package promptcore

// SyntaxMode controls which placeholder delimiters are recognized.
type SyntaxMode int

const (
	// SyntaxModeMixedBrace recognizes both single-brace and double-brace
	// placeholders. In both forms, name itself matches the regexp
	// `[^\s{}'"`?]+`. A trailing '?' marks the placeholder optional and is not
	// part of name. Double-brace placeholders still ignore outer whitespace.
	SyntaxModeMixedBrace SyntaxMode = iota
	// SyntaxModeSingleBrace recognizes placeholders such as {name} or {name?}.
	//
	// Here name matches the regexp `[^\s{}'"`?]+`. A trailing '?' marks the
	// placeholder optional and is not part of name.
	SyntaxModeSingleBrace
	// SyntaxModeDoubleBrace recognizes placeholders such as {{name}},
	// {{name?}}, or {{user:name}}.
	//
	// Here name matches the regexp `[^\s{}'"`?]+`. A trailing '?' marks the
	// placeholder optional and is not part of name. Outer whitespace is ignored.
	SyntaxModeDoubleBrace
)

// UnknownBehavior controls how unresolved placeholders are handled.
type UnknownBehavior int

const (
	// PreserveUnknown keeps unresolved placeholders in the rendered output.
	PreserveUnknown UnknownBehavior = iota
	// ErrorOnUnknown returns an error when a non-optional placeholder cannot be resolved.
	ErrorOnUnknown
)

// ResolveFunc resolves a placeholder name to a value.
type ResolveFunc func(name string) (string, bool, error)

// Env contains the runtime values available during rendering.
type Env struct {
	Vars    map[string]string
	Resolve ResolveFunc
}

// Option customizes promptcore parsing or rendering behavior.
type Option func(*config)

type config struct {
	acceptName func(string) bool
}

// WithAcceptName filters which parsed placeholder names are treated as real placeholders.
//
// Returning false keeps the placeholder unresolved in its preserved output form,
// even for syntactically valid optional placeholders.
func WithAcceptName(fn func(string) bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type textPart struct {
	literal     string
	placeholder *placeholderToken
}

type placeholderToken struct {
	raw      string
	name     string
	optional bool
	accepted bool
}

// Render replaces placeholders with values from env.
func Render(
	template string,
	syntax SyntaxMode,
	env Env,
	unknown UnknownBehavior,
	opts ...Option,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PlaceholderNames returns the unique placeholder names in template.
func PlaceholderNames(template string, syntax SyntaxMode, opts ...Option) []string {
	_ = "STUB: not implemented"
	return nil
}

func buildConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

func analyzeText(template string, syntax SyntaxMode, cfg config) []textPart {
	_ = "STUB: not implemented"
	return nil
}

func scanPlaceholder(
	template string,
	start int,
	syntax SyntaxMode,
	cfg config,
) (int, *placeholderToken) {
	_ = "STUB: not implemented"
	return 0, nil
}

func literalDoubleCurlySpan(template string, start int) int { _ = "STUB: not implemented"; return 0 }

func parseSingleBraceAt(
	template string,
	start int,
	cfg config,
) (int, *placeholderToken) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseDoubleCurlyAt(
	template string,
	start int,
	cfg config,
) (int, *placeholderToken) {
	_ = "STUB: not implemented"
	return 0, nil
}

type delimiterKind int

const (
	singleBraceDelimiter delimiterKind = iota
	doubleBraceDelimiter
)

func parsePlaceholder(
	raw, inner string,
	delimiter delimiterKind,
	cfg config,
) (placeholderToken, bool) {
	_ = "STUB: not implemented"
	return *new(placeholderToken), false
}

func parseName(inner string, delimiter delimiterKind) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func parseSingleBraceName(inner string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func parseDoubleBraceName(inner string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func isValidName(name string) bool { _ = "STUB: not implemented"; return false }

func renderPlaceholder(token placeholderToken, env Env) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func uniqueSortedStrings(values []string) []string { _ = "STUB: not implemented"; return nil }
