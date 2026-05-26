//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/gwclient"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/cron"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/gateway"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/memoryfile"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/outbound"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/persona"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/runtimeprofile"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	gwclientStatusErrFmt = "gwclient: status %d"

	gwclientStatusAPIErrorFmt = "gwclient: status %d: %s: %s"

	errNilGatewayServer = "gateway client: nil server"
	errNilCronService   = "gateway client: cron service unavailable"
	errNilPersonaStore  = "gateway client: persona store unavailable"
	errUnknownJob       = "gateway client: unknown scheduled job"

	debugTraceMetaFile = "meta.json"

	errEmptyForgetChannel = "gateway client: empty forget channel"
	errEmptyForgetUserID  = "gateway client: empty forget user id"
)

type inProcGatewayClient struct {
	srv      *gateway.Server
	appName  string
	sessions session.Service
	memories memory.Service
	cronSvc  *cron.Service

	debugDir        string
	uploads         *uploads.Store
	personas        *persona.Store
	memoryFileStore *memoryfile.Store
	profileCatalog  runtimeprofile.Catalog
	profileAppNames []string
}

func newInProcGatewayClient(
	srv *gateway.Server,
	appName string,
	sessions session.Service,
	memories memory.Service,
	debugDir string,
	uploadStores ...*uploads.Store,
) *inProcGatewayClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *inProcGatewayClient) SetCronService(svc *cron.Service) { _ = "STUB: not implemented"; return }

func (c *inProcGatewayClient) SetPersonaStore(store *persona.Store) {
	_ = "STUB: not implemented"
	return
}

func (c *inProcGatewayClient) SetMemoryFileStore(store *memoryfile.Store) {
	_ = "STUB: not implemented"
	return
}

func (c *inProcGatewayClient) SetRuntimeProfileAppNames(appNames []string) {
	_ = "STUB: not implemented"
	return
}

func (c *inProcGatewayClient) SetRuntimeProfileCatalog(
	catalog runtimeprofile.Catalog,
) {
	_ = "STUB: not implemented"
	return
}

func (c *inProcGatewayClient) SendMessage(
	ctx context.Context,
	req gwclient.MessageRequest,
) (gwclient.MessageResponse, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.MessageResponse), nil
}

func (c *inProcGatewayClient) StreamMessage(
	ctx context.Context,
	req gwclient.MessageRequest,
) (<-chan gwclient.StreamEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *inProcGatewayClient) StreamMessageWithOptions(
	ctx context.Context,
	req gwclient.MessageRequest,
	opts *gwclient.MessageStreamOptions,
) (<-chan gwclient.StreamEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *inProcGatewayClient) Cancel(
	ctx context.Context,
	requestID string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *inProcGatewayClient) ForgetUser(
	ctx context.Context,
	channel string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *inProcGatewayClient) forgetAppNames(
	ctx context.Context,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendUniqueUserIDs(
	base []string,
	extra ...string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *inProcGatewayClient) ListPresetPersonas() []persona.Preset {
	_ = "STUB: not implemented"
	return nil
}

func (c *inProcGatewayClient) GetPresetPersona(
	_ context.Context,
	scopeKey string,
) (persona.Preset, error) {
	_ = "STUB: not implemented"
	return *new(persona.Preset), nil
}

func (c *inProcGatewayClient) SetPresetPersona(
	ctx context.Context,
	scopeKey string,
	presetID string,
) (persona.Preset, error) {
	_ = "STUB: not implemented"
	return *new(persona.Preset), nil
}

func (c *inProcGatewayClient) ListScheduledJobs(
	_ context.Context,
	channel string,
	userID string,
	target string,
) ([]gwclient.ScheduledJobSummary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *inProcGatewayClient) ClearScheduledJobs(
	_ context.Context,
	channel string,
	userID string,
	target string,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *inProcGatewayClient) SetScheduledJobEnabled(
	_ context.Context,
	channel string,
	userID string,
	target string,
	jobID string,
	enabled bool,
) (gwclient.ScheduledJobSummary, error) {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary), nil
}

func (c *inProcGatewayClient) RemoveScheduledJob(
	_ context.Context,
	channel string,
	userID string,
	target string,
	jobID string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func errorForGWStatus(status int, apiErr *gwclient.APIError) error {
	_ = "STUB: not implemented"
	return nil
}

func cronDeliveryTarget(
	channel string,
	target string,
) outbound.DeliveryTarget {
	_ = "STUB: not implemented"
	return *new(outbound.DeliveryTarget)
}

func summarizeScheduledJob(job *cron.Job) gwclient.ScheduledJobSummary {
	_ = "STUB: not implemented"
	return *new(gwclient.ScheduledJobSummary)
}

func (c *inProcGatewayClient) scopedScheduledJob(
	channel string,
	userID string,
	target string,
	jobID string,
) (*cron.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func jobScheduleSummary(schedule cron.Schedule) string { _ = "STUB: not implemented"; return "" }

func cloneTime(src *time.Time) *time.Time { _ = "STUB: not implemented"; return nil }

type traceMeta struct {
	Start debugrecorder.TraceStart `json:"start"`
}

type traceRuntimeProfileEvent struct {
	Kind    string         `json:"kind"`
	Payload map[string]any `json:"payload,omitempty"`
}

func debugTraceMatchesApp(
	traceDir string,
	metaAppName string,
	appName string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func traceRuntimeProfileAppName(traceDir string) string { _ = "STUB: not implemented"; return "" }

func deleteDebugTraces(
	ctx context.Context,
	debugDir string,
	channel string,
	appName string,
	userID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func compactStrings(in []string) []string { _ = "STUB: not implemented"; return nil }
