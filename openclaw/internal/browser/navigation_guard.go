//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package browser

import (
	"net/netip"
)

type navigationPolicy struct {
	AllowedDomains  []string
	BlockedDomains  []string
	AllowLoopback   bool
	AllowPrivateNet bool
	AllowFileURLs   bool
}

func (p navigationPolicy) Validate(raw string) error { _ = "STUB: not implemented"; return nil }

func normalizeDomains(input []string) []string { _ = "STUB: not implemented"; return nil }

func normalizeHost(raw string) string { _ = "STUB: not implemented"; return "" }

func hostMatchesDomain(host, domain string) bool { _ = "STUB: not implemented"; return false }

func isLoopbackHost(host string) bool { _ = "STUB: not implemented"; return false }

func isPrivateAddress(addr netip.Addr) bool { _ = "STUB: not implemented"; return false }
