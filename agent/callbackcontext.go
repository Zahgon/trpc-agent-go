//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// CallbackContext provides a typed wrapper around context with agent-specific operations.
// Similar to ADK Python's callback_context, this provides access to session-scoped operations
// like artifact management.
type CallbackContext struct {
	context.Context
	invocation *Invocation
	// State is the delta-aware state of the current session.
	State session.StateMap
}

// NewCallbackContext creates a CallbackContext from a standard context.
// Returns an error if no invocation is found in the context.
func NewCallbackContext(ctx context.Context) (*CallbackContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveArtifact saves an artifact and records it for the current session.
//
// Args:
//   - filename: The filename of the artifact
//   - artifact: The artifact to save
//
// Returns:
//   - The version of the artifact
func (cc *CallbackContext) SaveArtifact(filename string, artifact *artifact.Artifact) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LoadArtifact loads an artifact attached to the current session.
//
// Args:
//   - filename: The filename of the artifact
//   - version: The version of the artifact. If nil, the latest version will be returned.
//
// Returns:
//   - The artifact, or nil if not found
func (cc *CallbackContext) LoadArtifact(filename string, version *int) (*artifact.Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListArtifacts lists the filenames of the artifacts attached to the current session.
//
// Returns:
//   - A list of artifact filenames
func (cc *CallbackContext) ListArtifacts() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteArtifact deletes an artifact from the current session.
//
// Args:
//   - filename: The filename of the artifact to delete
//
// Returns:
//   - An error if the operation fails
func (cc *CallbackContext) DeleteArtifact(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListArtifactVersions lists all versions of an artifact.
//
// Args:
//   - filename: The filename of the artifact
//
// Returns:
//   - A list of all available versions of the artifact
func (cc *CallbackContext) ListArtifactVersions(filename string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getArtifactServiceAndSessionInfo extracts common logic for getting artifact service and session information.
func (cc *CallbackContext) getArtifactServiceAndSessionInfo() (s artifact.Service, sessionInfo artifact.SessionInfo, err error) {
	_ = "STUB: not implemented"
	return *new(artifact.Service), *new(artifact.SessionInfo), nil
}

// appUserSession extracts app name, user ID, and session ID from the invocation.
func (cc *CallbackContext) appUserSession() (appName, userID, sessionID string, err error) {
	_ = "STUB: not implemented"
	// Try to get from session.
	return "", "", "", nil
}

// Session has AppName and UserID fields.

// Return error if session exists but missing required fields.
