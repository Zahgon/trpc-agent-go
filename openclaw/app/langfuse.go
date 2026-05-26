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

	"go.opentelemetry.io/otel/baggage"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/gateway"
	langfuseobs "trpc.group/trpc-go/trpc-agent-go/telemetry/langfuse"
)

const (
	langfuseHostEnv        = "LANGFUSE_HOST"
	langfuseInsecureEnv    = "LANGFUSE_INSECURE"
	langfuseInitProjectEnv = "LANGFUSE_INIT_PROJECT_ID"

	langfuseTraceIDPlaceholder = "{{trace_id}}"

	langfuseTraceNameKey           = "langfuse.trace.name"
	langfuseUserIDKey              = "langfuse.user.id"
	langfuseSessionIDKey           = "langfuse.session.id"
	langfuseMetadataPrefix         = "langfuse.trace.metadata."
	langfuseMetadataAppName        = langfuseMetadataPrefix + "app_name"
	langfuseMetadataChannel        = langfuseMetadataPrefix + "channel"
	langfuseMetadataRequestID      = langfuseMetadataPrefix + "request_id"
	langfuseMetadataMessageID      = langfuseMetadataPrefix + "message_id"
	langfuseMetadataProfileID      = langfuseMetadataPrefix + "profile_id"
	langfuseMetadataProfileVersion = langfuseMetadataPrefix +
		"profile_version"

	langfuseTraceDefaultName = "request"
)

var langfuseStart = langfuseobs.Start

type langfuseRuntime struct {
	adminStatus       admin.LangfuseStatus
	runOptionResolver gateway.RunOptionResolver
	shutdown          func(context.Context) error
}

func maybeEnableLangfuse(
	ctx context.Context,
	opts runOptions,
) (*langfuseRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func langfuseStartOptions(
	opts runOptions,
) []langfuseobs.Option {
	_ = "STUB: not implemented"
	return nil
}

func buildLangfuseAdminStatus(
	opts runOptions,
) admin.LangfuseStatus {
	_ = "STUB: not implemented"
	return *new(admin.LangfuseStatus)
}

func resolvedLangfuseUIBaseURL(opts runOptions) string { _ = "STUB: not implemented"; return "" }

func resolvedLangfuseTraceURLTemplate(
	opts runOptions,
	uiBaseURL string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func buildLangfuseRunOptionResolver(
	opts runOptions,
) gateway.RunOptionResolver {
	_ = "STUB: not implemented"
	return *new(gateway.RunOptionResolver)
}

func withLangfuseBaggage(
	ctx context.Context,
	appName string,
	input gateway.RunOptionInput,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func setLangfuseBaggageMember(
	bag baggage.Baggage,
	key string,
	value string,
) baggage.Baggage {
	_ = "STUB: not implemented"
	return *new(baggage.Baggage)
}

func buildLangfuseTraceName(
	fallbackAppName string,
	input gateway.RunOptionInput,
) string {
	_ = "STUB: not implemented"
	return ""
}
