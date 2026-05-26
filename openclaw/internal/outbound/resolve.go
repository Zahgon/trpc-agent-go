//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package outbound

import (
	"context"
)

const (
	runtimeStateDeliveryChannel = "openclaw.delivery.channel"
	runtimeStateDeliveryTarget  = "openclaw.delivery.target"

	wecomChannelName = "wecom"

	wecomThreadPrefix = wecomChannelName + ":thread:"
	wecomChatPrefix   = wecomChannelName + ":chat:"
	wecomDMPrefix     = wecomChannelName + ":dm:"

	wecomScopedUserSeparator = ":user:"

	wecomGroupTargetPrefix  = "group:"
	wecomSingleTargetPrefix = "single:"
)

// ResolveTarget chooses an outbound target from explicit args, runtime
// state, or the current session id.
func ResolveTarget(
	ctx context.Context,
	explicit DeliveryTarget,
) (DeliveryTarget, error) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), nil
}

func RuntimeStateForTarget(target DeliveryTarget) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func sanitizeTarget(target DeliveryTarget) DeliveryTarget {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget)
}

func fillTargetFromOpaqueValue(target DeliveryTarget) DeliveryTarget {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget)
}

func fillTargetFromRuntime(
	ctx context.Context,
	target DeliveryTarget,
) DeliveryTarget {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget)
}

func fillTargetFromSession(
	ctx context.Context,
	target DeliveryTarget,
) DeliveryTarget {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget)
}

// ResolveTargetFromSessionID infers an outbound target from a chat
// session id.
func ResolveTargetFromSessionID(sessionID string) (DeliveryTarget, bool) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), false
}

func resolveFromOpaqueTarget(value string) (DeliveryTarget, bool) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), false
}

func resolveTargetValue(
	channelID string,
	value string,
) (DeliveryTarget, bool) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), false
}

func resolveTelegramTargetValue(
	channelID string,
	value string,
) (DeliveryTarget, bool) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), false
}

func resolveWeComTargetValue(
	channelID string,
	value string,
) (DeliveryTarget, bool) {
	_ = "STUB: not implemented"
	return *new(DeliveryTarget), false
}

func validateTarget(target DeliveryTarget) error { _ = "STUB: not implemented"; return nil }

func normalizeWeComTarget(value string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func normalizeWeComChatTarget(value string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func normalizeWeComDMTarget(value string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseWeComPushTarget(value string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func parseWeComPushTargetWithPrefix(
	value string,
	prefix string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
