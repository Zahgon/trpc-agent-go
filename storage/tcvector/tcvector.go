//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package tcvector provides the tcvectordb instance info management.
package tcvector

import (
	"github.com/tencent/vectordatabase-sdk-go/tcvectordb"
)

func init() {
	tcvectorRegistry = make(map[string][]ClientBuilderOpt)
}

var tcvectorRegistry map[string][]ClientBuilderOpt

// ClientInterface is the interface for the tcvectordb client.
type ClientInterface interface {
	// DatabaseInterface is the interface for database operations.
	// Such as create database, collection, index, etc.
	tcvectordb.DatabaseInterface

	// FlatInterface is the interface for data operations.
	// Such as insert, update, delete, query, etc.
	tcvectordb.FlatInterface
}

type clientBuilder func(builderOpts ...ClientBuilderOpt) (ClientInterface, error)

// clientBuilder is the function to build the global tcvectordb client.
var globalBuilder clientBuilder = defaultClientBuilder

// SetClientBuilder sets the client builder for tcvectordb.
func SetClientBuilder(builder clientBuilder) { _ = "STUB: not implemented"; return }

// GetClientBuilder gets the tcvectordb client builder.
func GetClientBuilder() clientBuilder {
	_ = "STUB: not implemented"
	return *

	// defaultClientBuilder is the default client builder for tcvectordb.
	new(clientBuilder)
}

func defaultClientBuilder(builderOpts ...ClientBuilderOpt) (ClientInterface, error) {
	_ = "STUB: not implemented"
	return *new(ClientInterface), nil
}

// Validate required parameters

// ClientBuilderOpt is the option for the tcvectordb client.
type ClientBuilderOpt func(*ClientBuilderOpts)

// ClientBuilderOpts is the options for the tcvectordb client.
type ClientBuilderOpts struct {
	// HTTPURL is the http url for the tcvectordb client.
	HTTPURL string
	// UserName is the username for the tcvectordb client.
	UserName string
	// Key is the key for the tcvectordb client.
	Key string

	// ExtraOptions is reserved for custom client builders to accept
	// additional parameters. The default builder ignores it.
	ExtraOptions []any
}

// WithClientBuilderHTTPURL sets the http url for the tcvectordb client.
func WithClientBuilderHTTPURL(httpURL string) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithClientBuilderUserName sets the username for the tcvectordb client.
func WithClientBuilderUserName(userName string) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithClientBuilderKey sets the key for the tcvectordb client.
func WithClientBuilderKey(key string) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// WithExtraOptions sets extra options for customized client builders.
// The default builder does not consume these options.
func WithExtraOptions(extraOptions ...any) ClientBuilderOpt {
	_ = "STUB: not implemented"
	return *new(ClientBuilderOpt)
}

// RegisterTcVectorInstance registers a tcvectordb instance options.
// If the instance already exists, it will be overwritten.
func RegisterTcVectorInstance(name string, opts ...ClientBuilderOpt) {
	_ = "STUB: not implemented"
	return
}

// GetTcVectorInstance gets the tcvectordb instance options.
func GetTcVectorInstance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
