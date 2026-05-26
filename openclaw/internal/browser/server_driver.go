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
	"context"
	"net/http"
	"time"
)

const browserServerTimeout = 60 * time.Second

type serverProfileDriver struct {
	baseURL string
	token   string
	profile string
	client  *http.Client
}

func newServerProfileDriver(
	baseURL string,
	token string,
	profile string,
) *serverProfileDriver {
	_ = "STUB: not implemented"
	return nil
}

func (d *serverProfileDriver) Start(
	ctx context.Context,
) (driverStatus, error) {
	_ = "STUB: not implemented"
	return *new(driverStatus), nil
}

func (d *serverProfileDriver) Status(
	ctx context.Context,
) (driverStatus, error) {
	_ = "STUB: not implemented"
	return *new(driverStatus), nil
}

func (d *serverProfileDriver) Stop() error { _ = "STUB: not implemented"; return nil }

func (d *serverProfileDriver) Call(
	ctx context.Context,
	toolName string,
	args map[string]any,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func mapActArgs(kind string, args map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func stringValue(value any) string { _ = "STUB: not implemented"; return "" }

func numberValue(value any) int { _ = "STUB: not implemented"; return 0 }

func (d *serverProfileDriver) request(
	ctx context.Context,
	method string,
	path string,
	body any,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func queryArgs(body any) map[string]string { _ = "STUB: not implemented"; return nil }
