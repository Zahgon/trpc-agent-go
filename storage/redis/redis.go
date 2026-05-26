//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package redis provides the redis instance info management.
package redis

import (
	"github.com/redis/go-redis/v9"
)

func init() {
	redisRegistry = make(map[string][]ClientBuilderOpt)
}

var redisRegistry map[string][]ClientBuilderOpt

type clientBuilder func(builderOpts ...ClientBuilderOpt) (redis.UniversalClient, error)

var globalBuilder clientBuilder = defaultClientBuilder

// SetClientBuilder sets the redis client builder.
func SetClientBuilder(builder clientBuilder) { _ = "STUB: not implemented"; return }

// GetClientBuilder gets the redis client builder.
func GetClientBuilder() clientBuilder {
	_ = "STUB: not implemented"
	return *

	// defaultClientBuilder is the default redis client builder.
	new(clientBuilder)
}

func defaultClientBuilder(builderOpts ...ClientBuilderOpt) (redis.UniversalClient, error) {
	_ = "STUB: not implemented"
	return *new(redis.UniversalClient), nil
}

// ClientBuilderOpt is the option for the redis client.
type ClientBuilderOpt func(*ClientBuilderOpts)

// ClientBuilderOpts is the options for the redis client.
type ClientBuilderOpts struct {
	// URL is the redis client url for clientBuilder.
	URL string

	// ExtraOptions is the extra options for the redis client.
	ExtraOptions []any
}

// WithClientBuilderURL sets the redis client url for clientBuilder.
func WithClientBuilderURL(url string) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithExtraOptions sets the redis client extra options for clientBuilder.
// this option mainly used for the customized redis client builder, it will be passed to the builder.
func WithExtraOptions(extraOptions ...any) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// RegisterRedisInstance registers a redis instance options.
func RegisterRedisInstance(name string, opts ...ClientBuilderOpt) {
	_ = "STUB: not implemented"
	return
}

// GetRedisInstance gets the redis instance options.
func GetRedisInstance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
