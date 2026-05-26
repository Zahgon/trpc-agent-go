//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package cron

import "time"

// Patch represents optional job updates.
type Patch struct {
	Name             *string
	Message          *string
	Enabled          *bool
	Schedule         *Schedule
	ScheduleTimezone *string
	Policy           *ExecutionPolicy
	TimeoutSec       *int
	Channel          *string
	Target           *string
}

func applyPatch(job *Job, patch Patch, now time.Time) error { _ = "STUB: not implemented"; return nil }
