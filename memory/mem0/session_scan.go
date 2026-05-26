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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

func readLastExtractAt(sess *session.Session) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func writeLastExtractAt(sess *session.Session, ts time.Time) { _ = "STUB: not implemented"; return }

func scanDeltaSince(sess *session.Session, since time.Time) (time.Time, []model.Message) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
