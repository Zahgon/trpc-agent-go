//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package cos

import (
	"context"
	"io"
	"net/http"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// client interface is unstable and may change in the future.
type client interface {
	GetBucket(ctx context.Context, prefix string) (*cos.BucketGetResult, error)
	PutObject(ctx context.Context, name string, content io.Reader, opt cos.ObjectPutOptions) error
	GetObject(ctx context.Context, name string) (body io.ReadCloser, header http.Header, err error)
	DeleteObject(ctx context.Context, name string) error
}

type cosClient struct {
	*cos.Client
}

func newCosClient(client *cos.Client) client { _ = "STUB: not implemented"; return *new(client) }

func (c *cosClient) GetBucket(ctx context.Context, prefix string) (*cos.BucketGetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cosClient) PutObject(ctx context.Context, name string, content io.Reader, opt cos.ObjectPutOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cosClient) GetObject(ctx context.Context, name string) (body io.ReadCloser, header http.Header, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(http.Header), nil
}

func (c *cosClient) DeleteObject(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}
