//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package uploads

import (
	"context"
	"time"
)

const (
	defaultUploadsDir = "uploads"

	hostRefPrefix = "host://"

	metadataSuffix = ".meta.json"

	defaultChannelDir = "unknown-channel"
	defaultUserDir    = "unknown-user"
	defaultSessionDir = "unknown-session"
	defaultFileName   = "attachment"

	maxFileNameRunes = 96
	hashPrefixBytes  = 12
	hashPrefixHexLen = hashPrefixBytes * 2
	hashPrefixSep    = '-'

	fileMode = 0o600
	dirMode  = 0o755
)

const MetadataSuffix = metadataSuffix

const (
	KindImage = "image"
	KindAudio = "audio"
	KindVideo = "video"
	KindPDF   = "pdf"
	KindFile  = "file"
)

const (
	displayPhotoName      = "photo"
	displayAudioName      = "audio"
	displayVideoName      = "video"
	displayAnimationName  = "animation"
	displayDocumentName   = "document"
	displayAttachmentName = "attachment"
)

const (
	SourceInbound = "inbound"
	SourceDerived = "derived"
)

// FileMetadata describes optional metadata stored with one persisted file.
type FileMetadata struct {
	MimeType string `json:"mime_type,omitempty"`
	Source   string `json:"source,omitempty"`
}

// ListedFile describes one persisted upload entry.
type ListedFile struct {
	Scope        Scope
	Name         string
	Path         string
	HostRef      string
	RelativePath string
	MimeType     string
	Source       string
	SizeBytes    int64
	ModifiedAt   time.Time
}

// Scope identifies who owns a persisted upload.
type Scope struct {
	Channel   string
	UserID    string
	SessionID string
}

// SavedFile describes a persisted upload.
type SavedFile struct {
	Name    string
	Path    string
	HostRef string
}

// Store persists uploaded files under the OpenClaw state directory.
type Store struct {
	root string
}

// NewStore creates a new upload store rooted at stateDir/uploads.
func NewStore(stateDir string) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

// Root returns the uploads root directory.
func (s *Store) Root() string { _ = "STUB: not implemented"; return "" }

// ScopeDir returns the stable host directory for one upload scope.
func (s *Store) ScopeDir(scope Scope) string { _ = "STUB: not implemented"; return "" }

// Save persists data for the given scope and returns a stable host ref.
func (s *Store) Save(
	ctx context.Context,
	scope Scope,
	name string,
	data []byte,
) (SavedFile, error) {
	_ = "STUB: not implemented"
	return *new(SavedFile), nil
}

// SaveWithMetadata persists data together with optional metadata.
func (s *Store) SaveWithMetadata(
	ctx context.Context,
	scope Scope,
	name string,
	mimeType string,
	data []byte,
) (SavedFile, error) {
	_ = "STUB: not implemented"
	return *new(SavedFile), nil
}

// SaveWithInfo persists data together with optional metadata.
func (s *Store) SaveWithInfo(
	_ context.Context,
	scope Scope,
	name string,
	meta FileMetadata,
	data []byte,
) (SavedFile, error) {
	_ = "STUB: not implemented"
	return *new(SavedFile), nil
}

// DeleteUser removes all uploads for the given channel/user pair.
func (s *Store) DeleteUser(
	_ context.Context,
	channel string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListScope returns persisted uploads for one session scope, newest first.
func (s *Store) ListScope(scope Scope, limit int) ([]ListedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListAll returns persisted uploads across all users, newest first.
func (s *Store) ListAll(limit int) ([]ListedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Annotate writes metadata for one existing file already stored inside the
// uploads root.
func (s *Store) Annotate(path string, meta FileMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

// HostRef converts an absolute path into a host:// ref.
func HostRef(path string) string { _ = "STUB: not implemented"; return "" }

// PathFromHostRef returns the absolute host path for ref when possible.
func PathFromHostRef(ref string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (s *Store) scopeDirPath(scope Scope) string { _ = "STUB: not implemented"; return "" }

func sanitizeDirToken(raw string, fallback string) string { _ = "STUB: not implemented"; return "" }

func sanitizeFileName(raw string) string { _ = "STUB: not implemented"; return "" }

func writeFileIfMissing(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

func writeMetadataIfNeeded(path string, meta FileMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

func metadataPath(path string) string { _ = "STUB: not implemented"; return "" }

func isMetadataFileName(name string) bool { _ = "STUB: not implemented"; return false }

// IsMetadataPath reports whether path points to an uploads sidecar file.
func IsMetadataPath(path string) bool { _ = "STUB: not implemented"; return false }

func listStoredFiles(
	root string,
	walkRoot string,
	scope Scope,
	limit int,
) ([]ListedFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readMetadata(path string) (FileMetadata, error) {
	_ = "STUB: not implemented"
	return *new(FileMetadata), nil
}

func displayUploadName(name string) string { _ = "STUB: not implemented"; return "" }

func scopeFromRelativePath(rel string) Scope { _ = "STUB: not implemented"; return *new(Scope) }

// KindFromMeta returns a stable media kind from filename and MIME.
func KindFromMeta(name string, mimeType string) string { _ = "STUB: not implemented"; return "" }

// PreferredName returns a user-facing filename for one stored upload.
// It preserves meaningful names and rewrites generated Telegram placeholder
// names like "file_10.mp4" into stable names such as "video.mp4".
func PreferredName(name string, mimeType string) string { _ = "STUB: not implemented"; return "" }

func preferredNameBase(name string, mimeType string) string { _ = "STUB: not implemented"; return "" }

// StoredDisplayName strips one or more internal upload hash prefixes from a
// persisted filename while leaving normal filenames untouched.
func StoredDisplayName(name string) string { _ = "STUB: not implemented"; return "" }

func stripStoredHashPrefix(name string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func pathInsideRoot(path string, root string) bool { _ = "STUB: not implemented"; return false }

func preferredNameExt(mimeType string) string { _ = "STUB: not implemented"; return "" }

func isGeneratedPlaceholderName(name string) bool { _ = "STUB: not implemented"; return false }

func digitsOnly(raw string) bool { _ = "STUB: not implemented"; return false }

func sanitizeFileMetadata(meta FileMetadata) FileMetadata {
	_ = "STUB: not implemented"
	return *new(FileMetadata)
}

func sanitizeMetadataSource(raw string) string { _ = "STUB: not implemented"; return "" }
