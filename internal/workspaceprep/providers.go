//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package workspaceprep

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	rootskill "trpc.group/trpc-go/trpc-agent-go/skill"
)

// BootstrapSpec is the declarative workspace bootstrap description
// exposed to business code through llmagent.WithWorkspaceBootstrap.
//
// Files are staged first, then Commands run in declaration order. All
// entries are converted to Requirements by newBootstrapProvider.
type BootstrapSpec struct {
	// Files are static inputs (artifact://, host://, workspace://,
	// inline bytes) that must exist before commands run.
	Files []FileSpec
	// Commands are one-shot initialization commands such as
	// "python3 -m venv .venv" or "pip install -r requirements.txt".
	Commands []CommandSpec
}

// NewBootstrapProvider builds a Provider that emits the same set of
// static Requirements on every reconcile. Business code typically
// constructs one BootstrapSpec at agent construction time and passes
// it to llmagent.WithWorkspaceBootstrap.
func NewBootstrapProvider(
	spec BootstrapSpec,
) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

type staticProvider struct {
	name string
	reqs []Requirement
}

func (p *staticProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *staticProvider) Requirements(
	_ context.Context, _ *agent.Invocation,
) ([]Requirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLoadedSkillsProvider returns a Provider that walks the active
// invocation's session state (via skill.LoadedPrefix /
// StateKeyLoadedPrefix) and emits one SkillRequirement per skill the
// model has loaded for the current agent. Repository is used to
// resolve skill source paths; it should be the same repository that
// skill_load validates against.
func NewLoadedSkillsProvider(
	repo rootskill.Repository,
) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

type loadedSkillsProvider struct {
	repo rootskill.Repository
}

func (p *loadedSkillsProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *loadedSkillsProvider) Requirements(
	ctx context.Context, inv *agent.Invocation,
) ([]Requirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadedSkillsFromInvocation reads the session state for skills loaded
// by the active agent. It uses Session.SnapshotState() so that
// concurrent writers (parallel tool calls, async event handlers) do
// not race with the reconcile read. It checks both the agent-scoped
// and legacy unscoped prefixes so that older sessions keep working.
func loadedSkillsFromInvocation(inv *agent.Invocation) []string {
	_ = "STUB: not implemented"
	return nil
}

// Also include legacy unscoped keys so any older skill_load
// calls that happened before scoping was introduced still
// materialize the skill.

// NewConversationFilesProvider returns a Provider that emits a single
// Requirement wrapping the existing StageConversationFiles helper.
// The wrapping requirement fingerprints every file (file_id or
// sha256 over inline bytes) so that new uploads between reconciles
// force a re-stage, while repeated invocations with the same files
// are a fast no-op.
//
// Design note: v1 intentionally exposes one batch Requirement instead
// of one-per-file. The reasons are:
//
//   - StageConversationFiles already deduplicates per
//     file_id/sha256 inside the workspace metadata; reproducing
//     that logic at requirement granularity would duplicate the
//     source of truth.
//   - Conversation files are a tightly coupled set; they are staged
//     together, observed together, and their on-disk paths are
//     chosen by the helper rather than the requirement.
//
// If finer-grained Prepared/sentinel control becomes necessary we
// can split this into per-file requirements without changing the
// public surface (this provider stays an internal API).
func NewConversationFilesProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type conversationFilesProvider struct{}

func (p *conversationFilesProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *conversationFilesProvider) Requirements(
	_ context.Context, inv *agent.Invocation,
) ([]Requirement, error) {
	_ = "STUB: not implemented"
	// allConversationFiles already handles inv.Session == nil safely
	// by falling back to the current invocation's message parts, so we
	// only require the invocation itself to exist. Gating on Session !=
	// nil would regress the pre-refactor behavior where a user message
	// carrying a file part but no session history still got staged.
	return nil, nil
}

type conversationFilesRequirement struct{}

func (r *conversationFilesRequirement) Key() string { _ = "STUB: not implemented"; return "" }

func (r *conversationFilesRequirement) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (r *conversationFilesRequirement) Phase() Phase { _ = "STUB: not implemented"; return *new(Phase) }

func (r *conversationFilesRequirement) Required() bool { _ = "STUB: not implemented"; return false }

func (r *conversationFilesRequirement) Target() string { _ = "STUB: not implemented"; return "" }

// Fingerprint hashes the list of conversation files so that the same
// set of uploads produces a stable digest. We re-run StageConversationFiles
// whenever the fingerprint changes; the helper itself is idempotent
// for unchanged files because it tracks InputRecord entries in the
// workspace metadata.
func (r *conversationFilesRequirement) Fingerprint(
	ctx context.Context, rctx ApplyContext,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SentinelExists always returns true: StageConversationFiles is
// self-verifying via workspace metadata's existingByKey map, so a
// manual deletion of work/inputs/<name> on disk will not be
// re-materialized automatically. That matches the behavior users
// have today and keeps the provider a drop-in replacement for the
// current auto-staging path.
func (r *conversationFilesRequirement) SentinelExists(
	_ context.Context, _ ApplyContext,
) (bool, error) {
	_ = "STUB: not implemented"

	// Apply delegates to StageConversationFiles. Because this requirement
	// is Optional, the reconciler converts a non-nil error into a
	// warning rather than aborting; that matches the legacy auto-staging
	// behavior which never blocked execution on a partial stage.
	return false, nil
}

func (r *conversationFilesRequirement) Apply(
	ctx context.Context, rctx ApplyContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

// allConversationFiles collects a stable, sorted list of file
// identifiers from the current invocation. The traversal mirrors
// workspaceinput.StageConversationFiles exactly so that the
// fingerprint and the eventual stage operation agree on which files
// are part of the workspace. In particular:
//
//   - Session events are only considered when the message Role is
//     model.RoleUser. Tool/assistant messages may carry file parts
//     too, but they are not user-supplied inputs.
//   - The current invocation message contributes its file parts
//     unconditionally because the helper treats it as the active
//     user turn.
//   - Empty identifiers are dropped so that files that cannot be
//     fingerprinted (for example provider-side ids with no bytes)
//     do not perturb the digest.
//
// Events are read through a short locked snapshot to avoid racing
// with concurrent appenders (parallel tool calls, async event
// handlers).
func allConversationFiles(inv *agent.Invocation) []string { _ = "STUB: not implemented"; return nil }

func fileDigest(fileID string, data []byte) string { _ = "STUB: not implemented"; return "" }
