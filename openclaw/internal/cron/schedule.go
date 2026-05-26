//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package cron

import (
	"time"

	robcron "github.com/robfig/cron/v3"
)

const (
	cronFieldsWithMinutes = 5
	cronFieldsWithSeconds = 6
)

func computeInitialNextRun(
	schedule Schedule,
	now time.Time,
) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func computeNextAfterRun(
	schedule Schedule,
	base time.Time,
	now time.Time,
) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func computeNextRun(
	schedule Schedule,
	now time.Time,
) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseAtTime(raw string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseEveryInterval(schedule Schedule) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func parseCronSpec(schedule Schedule) (robcron.Schedule, error) {
	_ = "STUB: not implemented"
	return *new(robcron.Schedule), nil
}
