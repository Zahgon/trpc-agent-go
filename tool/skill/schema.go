//
// Tencent is pleased to support the open source community by making trpc-agent-go
// available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skill

import (
	skills "trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// maxSkillEnumValues bounds JSON schema size for large repositories.
const maxSkillEnumValues = 256

func skillNameSchema(repo skills.Repository, desc string) *tool.Schema {
	_ = "STUB: not implemented"
	return nil
}

func skillNameEnum(repo skills.Repository) []any { _ = "STUB: not implemented"; return nil }
