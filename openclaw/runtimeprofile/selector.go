//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runtimeprofile

import (
	"context"
)

const selectorConfigPath = "runtime_profiles.selectors"

// Selector maps request attributes to a runtime profile.
type Selector struct {
	ProfileID string   `yaml:"profile_id,omitempty"`
	Profile   string   `yaml:"profile,omitempty"`
	Channels  []string `yaml:"channels,omitempty"`
	Tenants   []string `yaml:"tenants,omitempty"`
	Users     []string `yaml:"users,omitempty"`
	Sessions  []string `yaml:"sessions,omitempty"`
}

// SelectorResolver chooses a selector profile before delegating to base.
type SelectorResolver struct {
	base      Resolver
	selectors []Selector
	aliases   map[string]string
}

// NewResolver creates a resolver from config, including selectors.
func NewResolver(cfg Config) (Resolver, error) {
	_ = "STUB: not implemented"
	return *new(Resolver), nil
}

// NewSelectorResolver wraps base with selector-based profile selection.
func NewSelectorResolver(
	base Resolver,
	selectors []Selector,
) (Resolver, error) {
	_ = "STUB: not implemented"
	return *new(Resolver), nil
}

func newSelectorResolver(
	base Resolver,
	selectors []Selector,
	aliases map[string]string,
) (Resolver, error) {
	_ = "STUB: not implemented"
	return *new(Resolver), nil
}

// Resolve implements Resolver.
func (r *SelectorResolver) Resolve(
	ctx context.Context,
	req Request,
) (Profile, error) {
	_ = "STUB: not implemented"
	return *new(Profile), nil
}

func (r *SelectorResolver) canonicalProfileID(profileID string) string {
	_ = "STUB: not implemented"
	return ""
}

func resolverProfileAliases(base Resolver) map[string]string { _ = "STUB: not implemented"; return nil }

// ProfileIDs implements Catalog when the base resolver also implements it.
func (r *SelectorResolver) ProfileIDs(
	ctx context.Context,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AppNames implements Catalog when the base resolver also implements it.
func (r *SelectorResolver) AppNames(
	ctx context.Context,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectProfileID returns the first selector profile that matches req.
func SelectProfileID(
	selectors []Selector,
	req Request,
) string {
	_ = "STUB: not implemented"
	return ""
}

func validateSelectors(
	selectors []Selector,
	known map[string]string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanSelectors(selectors []Selector) ([]Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectorError(index int, err error) error { _ = "STUB: not implemented"; return nil }

func (s Selector) runtimeProfileID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s Selector) hasCriteria() bool { _ = "STUB: not implemented"; return false }

func (s Selector) matches(req Request) bool { _ = "STUB: not implemented"; return false }

func matchesSelectorValue(values []string, got string) bool {
	_ = "STUB: not implemented"
	return false
}
