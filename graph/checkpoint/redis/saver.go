//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package redis provides Redis-based checkpoint storage implementation
// for graph execution state persistence and recovery.
package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	keyPrefixCheckpoint   = "ckpt:"
	keyPrefixCheckpointTS = "ckpt_ts:"
	keyPrefixWrites       = "writes:"
	keyPrefixLineageNS    = "lineage_ns:"
)

const (
	lingeageIDKey         = "lineage_id"
	checkpointIDKey       = "checkpoint_id"
	checkpointNSKey       = "checkpoint_ns"
	parentCheckpointIDKey = "parent_checkpoint_id"
	tsKey                 = "ts"
	checkpointJSONKey     = "checkpoint_json"
	metadataJSONKey       = "metadata_json"
)

var luaPutCheckpoint = redis.NewScript(`
-- Put checkpoint core data and indexes atomically.
--
-- KEYS[1] = checkpointKey
-- KEYS[2] = checkpointTSKey
-- KEYS[3] = lineageNSKey
--
-- ARGV[1] = ttl_ms
-- ARGV[2] = lineage_id
-- ARGV[3] = checkpoint_ns
-- ARGV[4] = checkpoint_id
-- ARGV[5] = parent_checkpoint_id
-- ARGV[6] = ts (unix nano)
-- ARGV[7] = checkpoint_json
-- ARGV[8] = metadata_json
local ttl = tonumber(ARGV[1])

redis.call('HSET', KEYS[1],
  'lineage_id', ARGV[2],
  'checkpoint_ns', ARGV[3],
  'checkpoint_id', ARGV[4],
  'parent_checkpoint_id', ARGV[5],
  'ts', ARGV[6],
  'checkpoint_json', ARGV[7],
  'metadata_json', ARGV[8]
)

redis.call('ZADD', KEYS[2], ARGV[6], ARGV[4])
redis.call('SADD', KEYS[3], ARGV[3])

if ttl and ttl > 0 then
  redis.call('PEXPIRE', KEYS[1], ttl)
  redis.call('PEXPIRE', KEYS[2], ttl)
  redis.call('PEXPIRE', KEYS[3], ttl)
end

return 1
`)

var luaPutFullCheckpoint = redis.NewScript(`
-- Put checkpoint core data, indexes and writes atomically.
--
-- KEYS[1] = checkpointKey
-- KEYS[2] = checkpointTSKey
-- KEYS[3] = lineageNSKey
-- KEYS[4] = writesKey
--
-- ARGV[1] = ttl_ms
-- ARGV[2] = lineage_id
-- ARGV[3] = checkpoint_ns
-- ARGV[4] = checkpoint_id
-- ARGV[5] = parent_checkpoint_id
-- ARGV[6] = ts (unix nano)
-- ARGV[7] = checkpoint_json
-- ARGV[8] = metadata_json
-- ARGV[9...] = write_field_1, write_json_1, write_field_2, write_json_2, ...
local ttl = tonumber(ARGV[1])

redis.call('HSET', KEYS[1],
  'lineage_id', ARGV[2],
  'checkpoint_ns', ARGV[3],
  'checkpoint_id', ARGV[4],
  'parent_checkpoint_id', ARGV[5],
  'ts', ARGV[6],
  'checkpoint_json', ARGV[7],
  'metadata_json', ARGV[8]
)

redis.call('ZADD', KEYS[2], ARGV[6], ARGV[4])
redis.call('SADD', KEYS[3], ARGV[3])

for i = 9, #ARGV, 2 do
  redis.call('HSET', KEYS[4], ARGV[i], ARGV[i + 1])
end

if ttl and ttl > 0 then
  redis.call('PEXPIRE', KEYS[1], ttl)
  redis.call('PEXPIRE', KEYS[2], ttl)
  redis.call('PEXPIRE', KEYS[3], ttl)
  redis.call('PEXPIRE', KEYS[4], ttl)
end

return 1
`)

func ttlMilliseconds(ttl time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func checkpointKey(lineageID, checkpointNS, checkpointID string) string {
	_ = "STUB: not implemented"
	return ""
}

func checkpointTSKey(lineageID, checkpointNS string) string { _ = "STUB: not implemented"; return "" }

func writesKey(lineageID, checkpointNS, checkpointID string) string {
	_ = "STUB: not implemented"
	return ""
}

func lineageNSKey(lineageID string) string { _ = "STUB: not implemented"; return "" }

type writeData struct {
	TaskID    string `json:"task_id"`
	Idx       int    `json:"idx"`
	Channel   string `json:"channel"`
	ValueJSON []byte `json:"value_json"`
	TaskPath  string `json:"task_path"`
	Seq       int64  `json:"seq"`
}

// Saver is the redis checkpoint service.
type Saver struct {
	opts   Options
	client redis.UniversalClient
	once   sync.Once // ensure Close is called only once
}

// NewSaver creates a new saver.
func NewSaver(options ...Option) (*Saver, error) { _ = "STUB: not implemented"; return nil, nil }

// if instance name set, and url not set, use instance name to create redis client

// Get returns the checkpoint for the given config.
func (s *Saver) Get(ctx context.Context, config map[string]any) (*graph.Checkpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTuple returns the checkpoint tuple for the given config.
func (s *Saver) GetTuple(ctx context.Context, config map[string]any) (*graph.CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Saver) findCheckpointID(ctx context.Context, lineageID, checkpointNS, checkpointID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Find a latest checkpoint in the namespace.

// List returns checkpoints for the lineage/namespace, with optional filters.
func (s *Saver) List(ctx context.Context, config map[string]any, filter *graph.CheckpointFilter) ([]*graph.CheckpointTuple, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Saver) getCheckpointIDs(ctx context.Context, lineageID, checkpointNS string, filter *graph.CheckpointFilter) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Saver) getCheckpointScore(ctx context.Context, lineageID, checkpointNS, checkpointID string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Put stores the checkpoint and returns the updated config with checkpoint ID.
func (s *Saver) Put(ctx context.Context, req graph.PutRequest) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Saver) PutWrites(ctx context.Context, req graph.PutWritesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// PutFull atomically stores a checkpoint with its pending writes in a single transaction.
func (s *Saver) PutFull(ctx context.Context, req graph.PutFullRequest) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteLineage deletes all checkpoints and writes for the lineage.
func (s *Saver) DeleteLineage(ctx context.Context, lineageID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Saver) loadWrites(ctx context.Context, lineageID, checkpointNS, checkpointID string) ([]graph.PendingWrite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Saver) findCheckpointNamespace(ctx context.Context, lineageID, checkpointID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Close closes the service.
func (s *Saver) Close() error {
	_ = "STUB: not implemented"

	// Close redis connection.
	return nil
}
