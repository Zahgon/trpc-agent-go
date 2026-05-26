//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package responseusage provides helpers for attaching response usage metadata.
package responseusage

import "trpc.group/trpc-go/trpc-agent-go/model"

// PartialState stores reusable usage state for partial responses.
type PartialState struct {
	usage      *model.Usage
	timingInfo *model.TimingInfo
}

// TimingAttachment records a temporary TimingInfo attachment.
type TimingAttachment struct {
	response           *model.Response
	usage              *model.Usage
	timingInfo         *model.TimingInfo
	attachedUsage      *model.Usage
	attachedTimingInfo *model.TimingInfo
	createdUsage       bool
	reusedUsage        bool
}

// AttachTimingForCallback attaches TimingInfo before callbacks and returns a restorer.
func AttachTimingForCallback(
	response *model.Response,
	timingInfo *model.TimingInfo,
	partialState *PartialState,
) TimingAttachment {
	_ = "STUB: not implemented"
	return *new(TimingAttachment)
}

// RestoreIfTimingInfoChanged restores the temporary attachment if the target TimingInfo changed.
func (a TimingAttachment) RestoreIfTimingInfoChanged(timingInfo *model.TimingInfo) {
	_ = "STUB: not implemented"
	return
}

// Restore removes the temporary attachment if it is still unchanged.
func (a TimingAttachment) Restore() { _ = "STUB: not implemented"; return }

func usageOnlyHasTimingInfo(usage *model.Usage) bool { _ = "STUB: not implemented"; return false }

// AttachTiming attaches TimingInfo to response usage.
func AttachTiming(
	response *model.Response,
	timingInfo *model.TimingInfo,
	partialState *PartialState,
) {
	_ = "STUB: not implemented"
	return
}
