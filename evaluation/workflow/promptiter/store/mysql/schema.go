//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mysql

import (
	"context"

	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

const (
	tableNameRuns         = "promptiter_runs"
	appRunUniqueIndexName = "uniq_promptiter_runs_app_run"
	sqlCreateRunsTable    = `
		CREATE TABLE IF NOT EXISTS {{TABLE_NAME}} (
			id BIGINT NOT NULL AUTO_INCREMENT,
			app_name VARCHAR(255) NOT NULL,
			run_id VARCHAR(255) NOT NULL,
			status VARCHAR(32) NOT NULL DEFAULT '',
			run_result JSON NOT NULL,
			created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
			updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`
	sqlCreateAppRunUniqueIndex = `
		CREATE UNIQUE INDEX {{INDEX_NAME}} ON {{TABLE_NAME}}(app_name, run_id)`
)

func ensureSchema(ctx context.Context, db storage.Client, tableName string) error {
	_ = "STUB: not implemented"
	return nil
}
