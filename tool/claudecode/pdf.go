//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package claudecode

type pdfPageRange struct {
	FirstPage int
	LastPage  int
	Count     int
}

func pdfPageCount(raw []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func resolvePDFPageRange(raw string, totalPages int) (pdfPageRange, error) {
	_ = "STUB: not implemented"
	return *new(pdfPageRange), nil
}

func validatedPDFPageRange(raw string, firstPage int, lastPage int) (pdfPageRange, error) {
	_ = "STUB: not implemented"
	return *new(pdfPageRange), nil
}

func pdftoppmBinary() (string, error) { _ = "STUB: not implemented"; return "", nil }

func extractPDFPages(
	filePath string,
	pageRange pdfPageRange,
) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
