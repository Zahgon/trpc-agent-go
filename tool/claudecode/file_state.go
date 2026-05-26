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
	"os"
)

func readLocalFileSnapshot(absPath string, maxFileSize int64) (localFileSnapshot, error) {
	_ = "STUB: not implemented"
	return *new(localFileSnapshot), nil
}

func ensureWriteAllowed(
	absPath string,
	snapshot localFileSnapshot,
	state *fileState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func writeLocalFile(
	absPath string,
	content string,
	mode os.FileMode,
	encoding string,
	lineEnding string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func storeReadView(
	state *fileState,
	absPath string,
	content string,
	timestamp int64,
	offset *int,
	limit *int,
	pages string,
	isPartial bool,
	fromRead bool,
) {
	_ = "STUB: not implemented"
	return
}

func matchesReadView(
	view fileView,
	offset *int,
	limit *int,
	pages string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func intPtrsEqual(left *int, right *int) bool { _ = "STUB: not implemented"; return false }

func normalizeQuotes(raw string) string { _ = "STUB: not implemented"; return "" }

func findActualString(fileContent string, searchString string) string {
	_ = "STUB: not implemented"
	return ""
}

func preserveQuoteStyle(oldString string, actualOldString string, newString string) string {
	_ = "STUB: not implemented"
	return ""
}

func applyCurlyDoubleQuotes(raw string) string { _ = "STUB: not implemented"; return "" }

func applyCurlySingleQuotes(raw string) string { _ = "STUB: not implemented"; return "" }

func isOpeningQuote(chars []rune, idx int) bool { _ = "STUB: not implemented"; return false }

func editLocalFile(
	absPath string,
	in editInput,
	runtime *runtime,
) (editOutput, error) {
	_ = "STUB: not implemented"
	return *new(editOutput), nil
}

func writeOutputToEditOutput(absPath string, in editInput, oldContent *string, newContent string) editOutput {
	_ = "STUB: not implemented"
	return *new(editOutput)
}
