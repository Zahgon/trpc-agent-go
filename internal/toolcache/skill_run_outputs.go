//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package toolcache provides per-invocation caches for tools.
//
// It is used to share skill_run output_files (inline content) across tools
// without relying on host filesystem paths.
package toolcache

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

const stateKeySkillRunOutputFiles = "tool:skill_run:output_files"

type cachedSkillRunFile struct {
	Content  string
	MIMEType string
}

// SkillRunOutputFile is a read-only view of an exported skill_run output.
// It is safe to pass across tools because it contains inline content.
type SkillRunOutputFile struct {
	Name     string
	Content  string
	MIMEType string
}

// StoreSkillRunOutputFilesFromContext stores skill_run output_files into the
// invocation state carried by ctx. It is a no-op when ctx has no invocation.
func StoreSkillRunOutputFilesFromContext(
	ctx context.Context,
	files []codeexecutor.File,
) {
	_ = "STUB: not implemented"
	return
}

// DeleteSkillRunOutputFilesFromContext deletes skill_run output_files from
// the invocation state carried by ctx. It is a no-op when ctx has no
// invocation.
func DeleteSkillRunOutputFilesFromContext(
	ctx context.Context,
	names []string,
) {
	_ = "STUB: not implemented"
	return
}

// StoreSkillRunOutputFiles stores skill_run output_files into inv so other
// tools can look them up by name later.
func StoreSkillRunOutputFiles(
	inv *agent.Invocation,
	files []codeexecutor.File,
) {
	_ = "STUB: not implemented"
	return
}

// DeleteSkillRunOutputFiles deletes skill_run output_files from inv so other
// tools no longer look them up by name later.
func DeleteSkillRunOutputFiles(
	inv *agent.Invocation,
	names []string,
) {
	_ = "STUB: not implemented"
	return
}

// LookupSkillRunOutputFileFromContext looks up an exported skill_run output
// file by name from the invocation in ctx.
func LookupSkillRunOutputFileFromContext(
	ctx context.Context,
	name string,
) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// LookupSkillRunOutputFile looks up an exported skill_run output file by name
// from inv. It returns (content, mime, ok).
func LookupSkillRunOutputFile(
	inv *agent.Invocation,
	name string,
) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// SkillRunOutputFilesFromContext returns a stable list of exported skill_run
// output files from the invocation in ctx.
func SkillRunOutputFilesFromContext(
	ctx context.Context,
) []SkillRunOutputFile {
	_ = "STUB: not implemented"
	return nil
}

// SkillRunOutputFiles returns a stable list of exported skill_run output
// files from inv.
func SkillRunOutputFiles(
	inv *agent.Invocation,
) []SkillRunOutputFile {
	_ = "STUB: not implemented"
	return nil
}
