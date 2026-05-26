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
	"net/http"
)

func normalizePath(baseDir string, raw string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (r *runtime) currentBaseDir() string { _ = "STUB: not implemented"; return "" }

func (r *runtime) setBaseDir(baseDir string) { _ = "STUB: not implemented"; return }

func relativePath(baseDir string, absPath string) string { _ = "STUB: not implemented"; return "" }

func readHTTPBody(
	resp *http.Response,
	maxContentLength int,
	maxTotalContentLength int,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func countLines(content string) int { _ = "STUB: not implemented"; return 0 }

func splitTextLines(content string) []string { _ = "STUB: not implemented"; return nil }

func sliceLines(content string, offset int, limit *int) (string, int, int) {
	_ = "STUB: not implemented"
	return "", 0, 0
}

func normalizeNewlines(content string) string { _ = "STUB: not implemented"; return "" }

func detectLineEnding(raw []byte) string { _ = "STUB: not implemented"; return "" }

func applyLineEnding(content string, lineEnding string) string {
	_ = "STUB: not implemented"
	return ""
}

func decodeTextBytes(raw []byte) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func encodeTextBytes(content string, encoding string, lineEnding string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileBase64(raw []byte) string { _ = "STUB: not implemented"; return "" }

func isProbablyBinary(raw []byte) bool { _ = "STUB: not implemented"; return false }

func buildStructuredPatch(oldContent string, newContent string) []patchHunk {
	_ = "STUB: not implemented"
	return nil
}

func matchSearchDomainFilters(
	rawURL string,
	allowed []string,
	blocked []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func searchURLHost(rawURL string) string { _ = "STUB: not implemented"; return "" }

func matchDomainRule(host string, rule string) bool { _ = "STUB: not implemented"; return false }

func extractHTMLText(raw []byte) string { _ = "STUB: not implemented"; return "" }

func collapseWhitespace(raw string) string { _ = "STUB: not implemented"; return "" }

func joinOutput(stdout string, stderr string) string { _ = "STUB: not implemented"; return "" }

func sortedCopy(items []string) []string { _ = "STUB: not implemented"; return nil }
