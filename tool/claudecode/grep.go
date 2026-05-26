//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package claudecode

import (
	"context"
	"regexp"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func newGrepTool(runtime *runtime) (tool.Tool, error) {
	_ = "STUB: not implemented"
	return *new(tool.Tool), nil
}

func grepDescription() string { _ = "STUB: not implemented"; return "" }

func runFallbackGrep(runtime *runtime, baseDir string, in grepInput) (grepOutput, error) {
	_ = "STUB: not implemented"
	return *new(grepOutput), nil
}

type fallbackGrepCollector struct {
	mode         string
	contentLines []string
	countLines   []string
	fileMatches  []string
}

func newFallbackGrepCollector(mode string) *fallbackGrepCollector {
	_ = "STUB: not implemented"
	return nil
}

func compileGrepPattern(in grepInput) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectFallbackGrepMatch(
	runtime *runtime,
	baseDir string,
	absPath string,
	re *regexp.Regexp,
	in grepInput,
	collector *fallbackGrepCollector,
) error {
	_ = "STUB: not implemented"
	return nil
}

func collectFallbackMultilineMatch(
	content string,
	relPath string,
	re *regexp.Regexp,
	in grepInput,
	collector *fallbackGrepCollector,
) error {
	_ = "STUB: not implemented"
	return nil
}

func collectFallbackLineMatch(
	content string,
	relPath string,
	re *regexp.Regexp,
	in grepInput,
	collector *fallbackGrepCollector,
) error {
	_ = "STUB: not implemented"
	return nil
}

func appendGrepContentLines(out *[]string, relPath string, lines []string, indexes []int, showLineNumbers bool) {
	_ = "STUB: not implemented"
	return
}

func showGrepLineNumbers(in grepInput) bool { _ = "STUB: not implemented"; return false }

func finalizeFallbackGrepOutput(baseDir string, in grepInput, collector *fallbackGrepCollector) grepOutput {
	_ = "STUB: not implemented"
	return *new(grepOutput)
}

func finalizeFallbackCountOutput(offset int, limit int, countLines []string) grepOutput {
	_ = "STUB: not implemented"
	return *new(grepOutput)
}

func multilineMatchLineIndexes(content string, matches [][]int, totalLines int) []int {
	_ = "STUB: not implemented"
	return nil
}

func lineIndexForOffset(lineStarts []int, offset int) int { _ = "STUB: not implemented"; return 0 }

func collectGrepCandidates(baseDir string, pathValue string, globValue string, typeValue string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func matchesAnyPattern(value string, patterns []string) bool {
	_ = "STUB: not implemented"
	return false
}

func expandContextLines(matches []int, total int, in grepInput) []int {
	_ = "STUB: not implemented"
	return nil
}

func typePatterns(typeValue string) []string { _ = "STUB: not implemented"; return nil }

func sliceStrings(items []string, offset int, limit int) ([]string, *int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runLocalRipgrep(
	ctx context.Context,
	baseDir string,
	in grepInput,
) (grepOutput, bool, error) {
	_ = "STUB: not implemented"
	return *new(grepOutput), false, nil
}

func runRipgrepCommand(
	ctx context.Context,
	baseAbs string,
	targetPath string,
	in grepInput,
) (grepOutput, bool, error) {
	_ = "STUB: not implemented"
	return *new(grepOutput), false, nil
}

func buildRipgrepArgs(mode string, targetPath string, in grepInput) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendRipgrepExcludes(args []string) []string { _ = "STUB: not implemented"; return nil }

func appendRipgrepMode(args []string, mode string, in grepInput) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendRipgrepContext(args []string, in grepInput) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendRipgrepPattern(args []string, pattern string) []string {
	_ = "STUB: not implemented"
	return nil
}

func formatRipgrepOutput(baseAbs string, mode string, in grepInput, lines []string) grepOutput {
	_ = "STUB: not implemented"
	return *new(grepOutput)
}

func formatRipgrepCountOutput(offset int, limit int, lines []string) grepOutput {
	_ = "STUB: not implemented"
	return *new(grepOutput)
}

func grepOffset(in grepInput) int { _ = "STUB: not implemented"; return 0 }

func grepLimit(in grepInput) int { _ = "STUB: not implemented"; return 0 }

func execRipgrep(
	ctx context.Context,
	baseAbs string,
	args ...string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitRipgrepLines(raw string) []string { _ = "STUB: not implemented"; return nil }

func splitGlobPatterns(raw string) []string { _ = "STUB: not implemented"; return nil }

func sortGrepPathsByMtime(baseAbs string, paths []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func parseGrepCountLine(line string) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func ripgrepCommand() string { _ = "STUB: not implemented"; return "" }
