//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
)

const (
	stateKeySkillRunOutputFiles = "tool:skill_run:output_files"
)

// skillRunOutputFile represents a file exported from skill_run output_files.
type skillRunOutputFile struct {
	Content  string
	MIMEType string
}

// lookupSkillRunOutputFileFromContext looks up a file from the
// skill_run output_files stored in the invocation state within ctx.
func lookupSkillRunOutputFileFromContext(
	ctx context.Context,
	relPath string,
) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// skillRunOutputFilesFromContext returns all files exported from
// skill_run output_files in the invocation state within ctx.
func skillRunOutputFilesFromContext(
	ctx context.Context,
) []skillRunOutputFile {
	_ = "STUB: not implemented"
	return nil
}

const (
	schemeSep = "://"

	schemeArtifact  = "artifact"
	schemeWorkspace = "workspace"

	artifactPrefix  = schemeArtifact + schemeSep
	workspacePrefix = schemeWorkspace + schemeSep
)

const errArtifactNameEmpty = "artifact name is empty"

// fileRef is a parsed file reference.
type fileRef struct {
	Scheme          string
	Path            string
	ArtifactName    string
	ArtifactVersion *int
	Raw             string
}

// parseFileRef parses raw into a fileRef.
func parseFileRef(raw string) (fileRef, error) {
	_ = "STUB: not implemented"
	return *new(fileRef), nil
}

func cleanRelPath(p string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// tryReadFileRef reads raw if it is a supported file reference.
func tryReadFileRef(
	ctx context.Context,
	raw string,
) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func loadArtifactFromContext(
	ctx context.Context,
	name string,
	version *int,
) ([]byte, string, int, error) {
	_ = "STUB: not implemented"
	return nil, "", 0, nil
}

func withArtifactContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
