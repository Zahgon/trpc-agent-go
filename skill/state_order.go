//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skill

// LoadedOrderKey returns the session state key that stores the recent
// skill-touch order for a specific agent.
//
// The JSON payload is an array of unique skill names ordered from the
// oldest touch to the newest touch.
func LoadedOrderKey(agentName string) string { _ = "STUB: not implemented"; return "" }

// ParseLoadedOrder parses a stored skill-touch order.
func ParseLoadedOrder(raw []byte) []string { _ = "STUB: not implemented"; return nil }

// MarshalLoadedOrder serializes a normalized skill-touch order.
func MarshalLoadedOrder(names []string) []byte { _ = "STUB: not implemented"; return nil }

// TouchLoadedOrder moves the touched skills to the end of the order.
func TouchLoadedOrder(names []string, touched ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

func normalizeLoadedOrder(names []string) []string { _ = "STUB: not implemented"; return nil }

func removeLoadedOrderName(order []string, target string) []string {
	_ = "STUB: not implemented"
	return nil
}
