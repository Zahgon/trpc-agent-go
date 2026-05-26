//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package model

import (
	"context"
)

// RecoverCallbackPanic converts provider callback panics into logged errors so
// user-defined hooks cannot crash the framework's streaming goroutines.
func RecoverCallbackPanic(ctx context.Context, stage string) { _ = "STUB: not implemented"; return }
