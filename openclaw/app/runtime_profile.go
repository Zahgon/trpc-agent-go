//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"context"
	"reflect"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/cron"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/gateway"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/runtimeprofile"
)

const unsupportedRuntimeProfileAgentFmt = "runtime_profiles.profiles.%s." +
	"agent_name: unsupported agent name %q; OpenClaw currently supports %q"

const runtimeProfileRequiredFmt = "runtime profile resolution failed: %w"

const unknownRuntimeProfileID = "unknown"

func buildRuntimeProfileRunOptionResolver(
	resolver runtimeprofile.Resolver,
	required bool,
) gateway.RunOptionResolver {
	_ = "STUB: not implemented"
	return *new(gateway.RunOptionResolver)
}

func appendRuntimeProfileGatewayOption(
	opts []gateway.Option,
	resolver runtimeprofile.Resolver,
	required bool,
) []gateway.Option {
	_ = "STUB: not implemented"
	return nil
}

func newRuntimeProfileResolver(
	cfg *runtimeprofile.Config,
) runtimeprofile.Resolver {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Resolver)
}

func runtimeProfileCronOptions(
	resolver runtimeprofile.Resolver,
) []cron.Option {
	_ = "STUB: not implemented"
	return nil
}

func runtimeProfileResolverFromOptions(
	cfg *runtimeprofile.Config,
	runtimeOpts runtimeOptions,
) (runtimeprofile.Resolver, runtimeprofile.Catalog, bool) {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Resolver), *new(runtimeprofile.Catalog), false
}

type runtimeProfileCatalogs []runtimeprofile.Catalog

func (c runtimeProfileCatalogs) ProfileIDs(
	ctx context.Context,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c runtimeProfileCatalogs) AppNames(
	ctx context.Context,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runtimeProfileCatalogFromOptions(
	resolver runtimeprofile.Resolver,
	injected runtimeprofile.Catalog,
) runtimeprofile.Catalog {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Catalog)
}

func appendRuntimeProfileCatalog(
	catalogs []runtimeprofile.Catalog,
	value any,
) []runtimeprofile.Catalog {
	_ = "STUB: not implemented"
	return nil
}

type runtimeProfileCatalogID struct {
	typ reflect.Type
	ptr uintptr
}

func runtimeProfileCatalogIdentity(
	catalog runtimeprofile.Catalog,
) (runtimeProfileCatalogID, bool) {
	_ = "STUB: not implemented"
	return *new(runtimeProfileCatalogID), false
}

func appendUniqueRuntimeProfileCatalogValues(
	base []string,
	extra ...string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func runtimeProfileRequired(cfg *runtimeprofile.Config) bool {
	_ = "STUB: not implemented"
	return false
}

func validateRuntimeProfiles(cfg *runtimeprofile.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func runtimeProfileAppNames(cfg *runtimeprofile.Config) []string {
	_ = "STUB: not implemented"
	return nil
}

func runtimeProfileIDForError(
	key string,
	profile runtimeprofile.Profile,
) string {
	_ = "STUB: not implemented"
	return ""
}

func appendUniqueAppNames(base []string, extra ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

func runtimeProfileRequest(
	input gateway.RunOptionInput,
) (runtimeprofile.Request, bool, error) {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Request), false, nil
}

func runtimeProfileBaseRequest(
	input gateway.RunOptionInput,
) runtimeprofile.Request {
	_ = "STUB: not implemented"
	return *new(runtimeprofile.Request)
}
