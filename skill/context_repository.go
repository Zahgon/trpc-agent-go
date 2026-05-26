//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skill

import (
	"context"
)

// VisibilityFilter decides whether a skill summary is visible for the current
// run context.
type VisibilityFilter func(ctx context.Context, summary Summary) bool

// ContextRepository is an optional repository extension that can resolve a
// per-context view over the underlying skill set.
type ContextRepository interface {
	Repository
	SummariesForContext(ctx context.Context) []Summary
	GetForContext(ctx context.Context, name string) (*Skill, error)
	PathForContext(ctx context.Context, name string) (string, error)
}

type filteredRepository struct {
	base   Repository
	filter VisibilityFilter
}

type skillRunEnvProvider interface {
	SkillRunEnv(
		ctx context.Context,
		skillName string,
	) (map[string]string, error)
}

// NewFilteredRepository wraps a repository with an additional per-context
// visibility filter.
func NewFilteredRepository(
	base Repository,
	filter VisibilityFilter,
) ContextRepository {
	_ = "STUB: not implemented"
	return *new(ContextRepository)
}

// IsContextAwareRepository reports whether repo supports context-aware access.
func IsContextAwareRepository(repo Repository) bool { _ = "STUB: not implemented"; return false }

// SummariesForContext resolves skill summaries for the given context.
func SummariesForContext(
	ctx context.Context,
	repo Repository,
) []Summary {
	_ = "STUB: not implemented"
	return nil
}

// GetForContext resolves a skill for the given context.
func GetForContext(
	ctx context.Context,
	repo Repository,
	name string,
) (*Skill, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PathForContext resolves a skill path for the given context.
func PathForContext(
	ctx context.Context,
	repo Repository,
	name string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *filteredRepository) Summaries() []Summary { _ = "STUB: not implemented"; return nil }

func (r *filteredRepository) Get(name string) (*Skill, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *filteredRepository) Path(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Roots exposes the underlying repository roots when available.
func (r *filteredRepository) Roots() []string { _ = "STUB: not implemented"; return nil }

func (r *filteredRepository) SummariesForContext(
	ctx context.Context,
) []Summary {
	_ = "STUB: not implemented"
	return nil
}

func (r *filteredRepository) GetForContext(
	ctx context.Context,
	name string,
) (*Skill, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *filteredRepository) PathForContext(
	ctx context.Context,
	name string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *filteredRepository) SkillRunEnv(
	ctx context.Context,
	skillName string,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterSummaries(
	ctx context.Context,
	summaries []Summary,
	filter VisibilityFilter,
) []Summary {
	_ = "STUB: not implemented"
	return nil
}

func skillVisibleByName(name string, summaries []Summary) bool {
	_ = "STUB: not implemented"
	return false
}

func skillNotFoundError(name string) error { _ = "STUB: not implemented"; return nil }
