//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telegram

import (
	"net/http"
	"time"
)

const (
	BaseURLEnvName         = "OPENCLAW_TELEGRAM_BASE_URL"
	telegramBaseURLEnvName = BaseURLEnvName

	errDefaultTransportType = "telegram: default transport is not http.Transport"
)

type ClientNetOptions struct {
	ProxyURL   string
	Timeout    time.Duration
	MaxRetries int
}

func BuildClientOptionsFromEnv(cfg ClientNetOptions) ([]Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BuildHTTPClient(
	rawProxyURL string,
	httpTimeout time.Duration,
) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
