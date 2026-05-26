//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package octool

import (
	"context"

	pdfpkg "github.com/ledongthuc/pdf"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
)

const (
	toolReadDocument    = "read_document"
	toolReadSpreadsheet = "read_spreadsheet"

	errReadPathRequired       = "path is required when no upload is present"
	errDocumentUnsupported    = "unsupported document type"
	errSpreadsheetUnsupported = "unsupported spreadsheet type"
	errSpreadsheetSheetEmpty  = "spreadsheet has no sheets"

	docKindPDF  = "pdf"
	docKindDOCX = "docx"
	docKindText = "text"

	sheetKindXLSX = "xlsx"
	sheetKindCSV  = "csv"

	defaultReadDocumentChars = 6_000
	defaultReadSheetChars    = 4_000
	defaultSheetPreviewRows  = 20

	schemaTypeObject = "object"
	schemaTypeString = "string"
	schemaTypeNumber = "number"
	schemaTypeArray  = "array"
)

type readDocumentTool struct {
	uploads *uploads.Store
}

type readSpreadsheetTool struct {
	uploads *uploads.Store
}

type readDocumentInput struct {
	Path     string `json:"path,omitempty"`
	Page     *int   `json:"page,omitempty"`
	MaxChars *int   `json:"max_chars,omitempty"`
}

type readDocumentResult struct {
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	Title     string `json:"title"`
	PageCount int    `json:"page_count,omitempty"`
	Page      *int   `json:"page,omitempty"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated,omitempty"`
}

type readSpreadsheetInput struct {
	Path     string `json:"path,omitempty"`
	Sheet    string `json:"sheet,omitempty"`
	Row      *int   `json:"row,omitempty"`
	StartRow *int   `json:"start_row,omitempty"`
	EndRow   *int   `json:"end_row,omitempty"`
	MaxChars *int   `json:"max_chars,omitempty"`
}

type spreadsheetRow struct {
	Index  int      `json:"index"`
	Values []string `json:"values,omitempty"`
}

type readSpreadsheetResult struct {
	Path      string           `json:"path"`
	Kind      string           `json:"kind"`
	Title     string           `json:"title"`
	Sheet     string           `json:"sheet,omitempty"`
	StartRow  int              `json:"start_row,omitempty"`
	EndRow    int              `json:"end_row,omitempty"`
	RowCount  int              `json:"row_count,omitempty"`
	Rows      []spreadsheetRow `json:"rows,omitempty"`
	Text      string           `json:"text"`
	Truncated bool             `json:"truncated,omitempty"`
}

func NewReadDocumentTool(
	stores ...*uploads.Store,
) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func NewReadSpreadsheetTool(
	stores ...*uploads.Store,
) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func firstUploadStore(stores []*uploads.Store) *uploads.Store {
	_ = "STUB: not implemented"
	return nil
}

func (t *readDocumentTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func (t *readDocumentTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (t *readSpreadsheetTool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"
	return nil
}

func (t *readSpreadsheetTool) Call(
	ctx context.Context,
	args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveDocumentPath(
	rawPath string,
	env map[string]string,
) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func resolveSpreadsheetPath(
	rawPath string,
	env map[string]string,
) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func resolveInputPath(
	rawPath string,
	env map[string]string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveUploadContextPath(
	path string,
	env map[string]string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func matchRecentUploadPath(
	rawPath string,
	env map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func uploadMatchesReference(
	item execUploadMeta,
	rawPath string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func uploadReferenceCandidates(item execUploadMeta) []string { _ = "STUB: not implemented"; return nil }

func appendUniqueTrimmed(values []string, value string) []string {
	_ = "STUB: not implemented"
	return nil
}

func documentKindFromPath(path string) string { _ = "STUB: not implemented"; return "" }

func spreadsheetKindFromPath(path string) string { _ = "STUB: not implemented"; return "" }

func readDocumentText(
	path string,
	kind string,
	page *int,
) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func readPDFText(
	path string,
	page *int,
) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func pdfPageText(
	reader *pdfpkg.Reader,
	pageIndex int,
) string {
	_ = "STUB: not implemented"
	return ""
}

func readDOCXText(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func readTextFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func readSpreadsheetRows(
	path string,
	kind string,
	sheet string,
) ([][]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func readCSVRows(path string) ([][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func readWorkbookRows(
	path string,
	sheet string,
) ([][]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func selectSpreadsheetRows(
	rows [][]string,
	in readSpreadsheetInput,
) ([]spreadsheetRow, int, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func spreadsheetRange(
	totalRows int,
	in readSpreadsheetInput,
) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func sanitizeSpreadsheetCells(values []string) []string { _ = "STUB: not implemented"; return nil }

func formatSpreadsheetRows(rows []spreadsheetRow) string { _ = "STUB: not implemented"; return "" }

func formatSpreadsheetRow(row spreadsheetRow) string { _ = "STUB: not implemented"; return "" }

func truncateText(text string, maxChars int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func normalizedPositive(value *int) *int { _ = "STUB: not implemented"; return nil }

func resolvedMaxChars(value *int, fallback int) int { _ = "STUB: not implemented"; return 0 }

func minInt(a int, b int) int { _ = "STUB: not implemented"; return 0 }

var _ tool.CallableTool = (*readDocumentTool)(nil)
var _ tool.CallableTool = (*readSpreadsheetTool)(nil)
