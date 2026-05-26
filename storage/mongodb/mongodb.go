//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package mongodb provides the MongoDB instance info management.
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	mongodbRegistry = make(map[string][]ClientBuilderOpt)
}

var mongodbRegistry map[string][]ClientBuilderOpt

type clientBuilder func(ctx context.Context, builderOpts ...ClientBuilderOpt) (Client, error)

var globalBuilder clientBuilder = defaultClientBuilder

// SetClientBuilder sets the mongodb client builder.
func SetClientBuilder(builder clientBuilder) { _ = "STUB: not implemented"; return }

// GetClientBuilder gets the mongodb client builder.
func GetClientBuilder() clientBuilder {
	_ = "STUB: not implemented"
	return *

	// mongoConnector is the function used to connect to MongoDB.
	new(clientBuilder)
}

var mongoConnector = func(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	return mongo.Connect(ctx, opts...)
}

// defaultClientBuilder is the default mongodb client builder.
// It creates a native MongoDB client using the official Go driver.
func defaultClientBuilder(ctx context.Context, builderOpts ...ClientBuilderOpt) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Create MongoDB client options

// Connect to MongoDB

// Verify connection

// RegisterMongoDBInstance registers a mongodb instance with the given options.
func RegisterMongoDBInstance(name string, opts ...ClientBuilderOpt) {
	_ = "STUB: not implemented"
	return
}

// GetMongoDBInstance gets the mongodb instance options by name.
func GetMongoDBInstance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Client defines the interface for MongoDB operations.
// This is a subset of the internal mongodb.Client interface,
// containing only the methods needed by the session layer.
type Client interface {
	// InsertOne executes an insert command to insert a single document into the collection.
	InsertOne(ctx context.Context, database string, coll string, document any,
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	// UpdateOne executes an update command to update at most one document in the collection.
	UpdateOne(ctx context.Context, database string, coll string, filter any, update any,
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)

	// DeleteOne executes a delete command to delete at most one document from the collection.
	DeleteOne(ctx context.Context, database string, coll string, filter any,
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

	// DeleteMany executes a delete command to delete documents from the collection.
	DeleteMany(ctx context.Context, database string, coll string, filter any,
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

	// FindOne executes a find command and returns a SingleResult for one document in the collection.
	FindOne(ctx context.Context, database string, coll string, filter any,
		opts ...*options.FindOneOptions) *mongo.SingleResult

	// Find executes a find command and returns a Cursor over the matching documents in the collection.
	Find(ctx context.Context, database string, coll string, filter any,
		opts ...*options.FindOptions) (*mongo.Cursor, error)

	// CountDocuments returns the number of documents in the collection.
	CountDocuments(ctx context.Context, database string, coll string, filter any,
		opts ...*options.CountOptions) (int64, error)

	// Transaction executes a transaction.
	// The sf parameter is a function that receives a mongo.SessionContext for transaction operations.
	Transaction(ctx context.Context, sf func(sc mongo.SessionContext) error, tOpts []*options.TransactionOptions,
		opts ...*options.SessionOptions) error

	// Disconnect closes the mongo client.
	Disconnect(ctx context.Context) error
}

// session defines the interface for MongoDB session operations.
type session interface {
	EndSession(ctx context.Context)
	WithTransaction(ctx context.Context, fn func(sc mongo.SessionContext) (any, error),
		opts ...*options.TransactionOptions) (any, error)
}

// defaultClient wraps *mongo.Client to implement the Client interface.
type defaultClient struct {
	client       *mongo.Client
	startSession func(opts ...*options.SessionOptions) (session, error)
}

// newDefaultClient creates a new defaultClient with the given mongo.Client.
func newDefaultClient(client *mongo.Client) *defaultClient { _ = "STUB: not implemented"; return nil }

// InsertOne implements Client.InsertOne.
func (c *defaultClient) InsertOne(ctx context.Context, database string, coll string, document any,
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateOne implements Client.UpdateOne.
func (c *defaultClient) UpdateOne(ctx context.Context, database string, coll string, filter any,
	update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteOne implements Client.DeleteOne.
func (c *defaultClient) DeleteOne(ctx context.Context, database string, coll string, filter any,
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteMany implements Client.DeleteMany.
func (c *defaultClient) DeleteMany(ctx context.Context, database string, coll string, filter any,
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindOne implements Client.FindOne.
func (c *defaultClient) FindOne(ctx context.Context, database string, coll string, filter any,
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

// Find implements Client.Find.
func (c *defaultClient) Find(ctx context.Context, database string, coll string, filter any,
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountDocuments implements Client.CountDocuments.
func (c *defaultClient) CountDocuments(ctx context.Context, database string, coll string, filter any,
	opts ...*options.CountOptions) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Transaction implements Client.Transaction.
func (c *defaultClient) Transaction(ctx context.Context, sf func(sc mongo.SessionContext) error,
	tOpts []*options.TransactionOptions, opts ...*options.SessionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnect implements Client.Disconnect.
func (c *defaultClient) Disconnect(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
