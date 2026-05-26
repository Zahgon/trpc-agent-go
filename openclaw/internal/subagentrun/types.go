//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package subagentrun

import (
	"time"

	coretaskrun "trpc.group/trpc-go/trpc-agent-go/agent/taskrun"
	openclawsubagent "trpc.group/trpc-go/trpc-agent-go/openclaw/subagent"
)

const (
	subagentDirName      = "subagents"
	subagentRunsFileName = "runs.json"
	subagentIDPrefix     = "subagent:"

	metadataDeliveryChannel = "openclaw.delivery.channel"
	metadataDeliveryTarget  = "openclaw.delivery.target"

	defaultNotifyTimeout = 15 * time.Second

	notificationPrefixCompleted = "✅ subagent 已完成"
	notificationPrefixFailed    = "⚠️ subagent 失败"

	subagentRunPrompt = "You are running as an OpenClaw background " +
		"subagent. Complete the delegated task once. The parent " +
		"chat will receive your final result automatically. Keep " +
		"the result concise and action-oriented. Do not return " +
		"only a statement of what you will do; complete the " +
		"task and report the result or exact blocker. Do not " +
		"spawn more subagents from inside this subagent."
)

type deliveryTarget struct {
	Channel string `json:"channel,omitempty"`
	Target  string `json:"target,omitempty"`
}

type SpawnRequest struct {
	OwnerUserID                    string
	ParentSessionID                string
	Task                           string
	TimeoutSeconds                 int
	Delivery                       deliveryTarget
	SuppressCompletionNotification bool
}

func subagentStorePath(stateDir string) string { _ = "STUB: not implemented"; return "" }

func metadataForDelivery(target deliveryTarget) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func deliveryFromRun(run coretaskrun.Run) deliveryTarget {
	_ = "STUB: not implemented"
	return *new(deliveryTarget)
}

func timeoutDuration(seconds int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func newSubagentID() string { _ = "STUB: not implemented"; return "" }

func subagentRuntimeStateKeys() coretaskrun.RuntimeStateKeys {
	_ = "STUB: not implemented"
	return *new(coretaskrun.RuntimeStateKeys)
}

func projectRun(run coretaskrun.Run) openclawsubagent.Run {
	_ = "STUB: not implemented"
	return *new(openclawsubagent.Run)
}

func projectRunPtr(run *coretaskrun.Run) *openclawsubagent.Run {
	_ = "STUB: not implemented"
	return nil
}

func projectRuns(runs []coretaskrun.Run) []openclawsubagent.Run {
	_ = "STUB: not implemented"
	return nil
}

func cloneTimePtr(value *time.Time) *time.Time { _ = "STUB: not implemented"; return nil }
