//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package recorder

import (
	"sync"
)

type keyedLocker struct {
	mu    sync.Mutex
	locks map[string]*keyedLock
}

type keyedLock struct {
	mu   sync.Mutex
	refs int
}

func newKeyedLocker() *keyedLocker { _ = "STUB: not implemented"; return nil }

func (l *keyedLocker) lock(key string) { _ = "STUB: not implemented"; return }

func (l *keyedLocker) unlock(key string) { _ = "STUB: not implemented"; return }
