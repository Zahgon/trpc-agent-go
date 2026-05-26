//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skill

import (
	"archive/tar"
	"archive/zip"
	"net/url"
	"os"
)

const (
	cacheAppDir    = "trpc-agent-go"
	cacheSkillsDir = "skills"

	cacheReadyFile    = ".ready"
	cacheDownloadFile = "download"
	cacheExtractDir   = "root"

	cacheTempPrefix = "tmp-skill-root-"

	dirPerm  = 0o755
	filePerm = 0o644

	bytesPerMiB = 1 << 20

	maxDownloadBytes = 64 * bytesPerMiB

	maxExtractFileBytes  = 64 * bytesPerMiB
	maxExtractTotalBytes = 256 * bytesPerMiB
)

// EnvSkillsCacheDir overrides where URL-based skills roots are cached.
// When empty, the user cache directory is used.
const EnvSkillsCacheDir = "SKILLS_CACHE_DIR"

func resolveSkillsRoot(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func fileURLPath(u *url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }

func cacheURLRoot(u *url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }

func skillsCacheDir() string { _ = "STUB: not implemented"; return "" }

func downloadURLToFile(u *url.URL, path string) error { _ = "STUB: not implemented"; return nil }

func extractURLPayload(u *url.URL, srcPath string, destDir string) error {
	_ = "STUB: not implemented"
	return nil
}

type archiveKind int

const (
	archiveKindUnknown archiveKind = iota
	archiveKindZip
	archiveKindTar
	archiveKindTarGZ
)

const (
	extZip   = ".zip"
	extTar   = ".tar"
	extTGZ   = ".tgz"
	extTarGZ = ".tar.gz"
)

func archiveKindFromName(name string) archiveKind {
	_ = "STUB: not implemented"
	return *new(archiveKind)
}

func detectArchiveKind(srcPath string) archiveKind {
	_ = "STUB: not implemented"
	return *new(archiveKind)
}

func extractZip(srcPath string, destDir string) error { _ = "STUB: not implemented"; return nil }

func extractZipFile(f *zip.File, destDir string, total *int64) error {
	_ = "STUB: not implemented"
	return nil
}

func extractTar(srcPath string, destDir string) error { _ = "STUB: not implemented"; return nil }

func extractTarGZ(srcPath string, destDir string) error { _ = "STUB: not implemented"; return nil }

func extractTarReader(tr *tar.Reader, destDir string) error { _ = "STUB: not implemented"; return nil }

func writeSingleSkillFile(srcPath string, destDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanArchivePath(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func sanitizePerm(m os.FileMode) os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

const (
	tarPermUserRead  = 0o400
	tarPermUserWrite = 0o200
	tarPermUserExec  = 0o100

	tarPermGroupRead  = 0o040
	tarPermGroupWrite = 0o020
	tarPermGroupExec  = 0o010

	tarPermOtherRead  = 0o004
	tarPermOtherWrite = 0o002
	tarPermOtherExec  = 0o001
)

func tarHeaderPerm(mode int64) os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

func fileExists(path string) bool { _ = "STUB: not implemented"; return false }

func sha256Hex(s string) string { _ = "STUB: not implemented"; return "" }

func validateTarSize(size int64) error { _ = "STUB: not implemented"; return nil }

func validateZipEntrySize(f *zip.File) error { _ = "STUB: not implemented"; return nil }

func addExtractedBytes(total *int64, n int64) error { _ = "STUB: not implemented"; return nil }
