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
)

const (
	errCommandPolicyRejected = "exec_command blocked by safety " +
		"policy: %s"

	reasonSensitivePath = "reading or modifying shell or " +
		"credential files is not allowed in chat"

	sensitivePathBoundaryChars = " \t\r\n\"'`=:/\\|&;()[]{}<>"

	envTRPCClawEnvFile  = "TRPC_CLAW_ENV_FILE"
	envTRPCClawStateDir = "TRPC_CLAW_STATE_DIR"

	protectedRuntimeEnvRelPath = "runtime/env.sh"
	protectedGitCredentialFile = "git-credentials"

	shellQuoteChars = `"'`
)

var protectedPathFragments = []string{
	".aws/credentials",
	".bash_profile",
	".bashrc",
	".config/gcloud",
	".docker/config.json",
	".git-credentials",
	".kube/config",
	".netrc",
	".npmrc",
	".profile",
	".pypirc",
	".ssh/",
	".zprofile",
	".zshenv",
	".zshrc",
}

// CommandRequest is the normalized command metadata checked by policies.
type CommandRequest struct {
	Command    string
	Workdir    string
	Env        map[string]string
	Pty        bool
	Background bool
	YieldMs    *int
	TimeoutS   *int
}

// CommandPolicy decides whether one exec_command call is allowed.
type CommandPolicy func(context.Context, CommandRequest) error

// NewChatCommandSafetyPolicy blocks direct access to protected shell and
// credential paths in chat contexts.
func NewChatCommandSafetyPolicy() CommandPolicy {
	_ = "STUB: not implemented"
	return *new(CommandPolicy)
}

func newCommandRequest(params execParams) CommandRequest {
	_ = "STUB: not implemented"
	return *new(CommandRequest)
}

func blocksSensitivePath(command string) bool { _ = "STUB: not implemented"; return false }

func blocksSensitivePathRequest(req CommandRequest) bool { _ = "STUB: not implemented"; return false }

func blocksSensitivePathValue(
	raw string,
	protectedFragments []string,
	dynamicFragments []string,
	env map[string]string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizePolicyCommand(command string) string { _ = "STUB: not implemented"; return "" }

func blocksSensitivePathWithFragments(
	command string,
	fragments []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func matchesProtectedPathFragments(
	command string,
	fragments []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func dynamicProtectedPathFragments(env map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func protectedWorkdirFragments() []string { _ = "STUB: not implemented"; return nil }

func dynamicProtectedWorkdirFragments(env map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func expandProtectedEnvReferences(
	command string,
	env map[string]string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func expandProtectedEnvReference(
	command string,
	envName string,
	value string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func appendProtectedPathFragment(out []string, raw string) []string {
	_ = "STUB: not implemented"
	return nil
}

func appendProtectedPathDir(out []string, raw string) []string {
	_ = "STUB: not implemented"
	return nil
}

func normalizePathFragment(raw string) string { _ = "STUB: not implemented"; return "" }

func stripShellQuotes(command string) string { _ = "STUB: not implemented"; return "" }

func containsSensitivePathFragment(command, fragment string) bool {
	_ = "STUB: not implemented"
	return false
}

func hasSensitivePathBoundaryBefore(command string, idx int) bool {
	_ = "STUB: not implemented"
	return false
}

func hasSensitivePathBoundaryAfter(
	command string,
	idx int,
	fragment string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isSensitivePathBoundary(ch byte) bool { _ = "STUB: not implemented"; return false }
