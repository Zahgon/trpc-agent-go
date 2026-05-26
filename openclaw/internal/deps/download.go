//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package deps

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

const downloadHTTPTimeout = 30 * time.Minute

var downloadHTTPClient = &http.Client{
	Timeout: downloadHTTPTimeout,
}

func downloadInstallStep(
	toolchain Toolchain,
	sourceName string,
	action InstallAction,
) (Step, error) {
	_ = "STUB: not implemented"
	return *new(Step), nil
}

func executeDownloadStep(
	ctx context.Context,
	step Step,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeDownloadFile(
	targetPath string,
	reader io.Reader,
) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadTargetPath(
	stateDir string,
	sourceName string,
	action InstallAction,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func skillToolsRoot(
	stateDir string,
	sourceName string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func safeJoin(root string, rel string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func downloadCommandLine(
	rawURL string,
	targetPath string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func downloadFileName(rawURL string) string { _ = "STUB: not implemented"; return "" }

func openDownloadReader(
	ctx context.Context,
	rawURL string,
) (io.ReadCloser, func(), error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil, nil
}

func extractArchive(
	reader io.Reader,
	archiveKind string,
	targetPath string,
	stripComponents int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func normalizeArchiveKind(raw string) string { _ = "STUB: not implemented"; return "" }

func extractTar(
	reader io.Reader,
	targetPath string,
	stripComponents int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func extractZip(
	reader io.Reader,
	targetPath string,
	stripComponents int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func archiveTargetName(
	name string,
	stripComponents int,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func writeArchiveEntry(
	dst string,
	mode os.FileMode,
	reader io.Reader,
) error {
	_ = "STUB: not implemented"
	return nil
}

func archiveEntryPerm(mode os.FileMode) os.FileMode {
	_ = "STUB: not implemented"
	return *new(os.FileMode)
}
