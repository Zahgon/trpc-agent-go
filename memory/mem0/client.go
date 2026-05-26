//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mem0

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

const (
	httpHeaderAuthorization = "Authorization"
	httpHeaderAccept        = "Accept"
	httpHeaderContentType   = "Content-Type"

	httpContentTypeJSON = "application/json"

	httpMethodGet  = "GET"
	httpMethodPost = "POST"

	maxResponseBodySize = 10 << 20

	maxRetries       = 3
	retryBaseBackoff = 200 * time.Millisecond
	retryMaxBackoff  = 2 * time.Second

	maxErrorBodyPreview = 512
)

type apiError struct {
	StatusCode int
	Body       string
}

func (e *apiError) Error() string { _ = "STUB: not implemented"; return "" }

type client struct {
	host      string
	apiKey    string
	orgID     string
	projectID string
	hc        *http.Client
	timeout   time.Duration
}

func newClient(opts serviceOpts) (*client, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *client) doJSON(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	in any,
	out any,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not retry non-idempotent write requests.

func (c *client) doJSONOnce(
	ctx context.Context,
	method string,
	urlStr string,
	payload []byte,
	out any,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldRetry(err error) bool { _ = "STUB: not implemented"; return false }

func retrySleep(attempt int, jitterFn func(max int64) int64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func cryptoJitter(max int64) int64 { _ = "STUB: not implemented"; return 0 }

func itoa(v int) string { _ = "STUB: not implemented"; return "" }
