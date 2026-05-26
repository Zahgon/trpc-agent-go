//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/evaluation"
)

var (
	mysqlDSN    = flag.String("dsn", "user:password@tcp(localhost:3306)/db?parseTime=true&charset=utf8mb4", "MySQL DSN used by evaluation managers")
	tablePrefix = flag.String("table-prefix", "evaluation_example", "Table prefix applied to evaluation tables")
	skipDBInit  = flag.Bool("skip-db-init", false, "Skip table creation during manager initialization")
	modelName   = flag.String("model", "deepseek-v4-flash", "Model to use for evaluation runs")
	streaming   = flag.Bool("streaming", false, "Enable streaming responses from the agent")
	numRuns     = flag.Int("runs", 1, "Number of times to repeat the evaluation loop per case")
	evalSetID   = flag.String("eval-set", "math-basic", "Evaluation set identifier to execute")
)

const appName = "math-eval-app"

func main() {
	flag.Parse()
	if err := run(context.Background()); err != nil {
		log.Fatalf("run evaluation: %v", err)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func printSummary(result *evaluation.EvaluationResult) { _ = "STUB: not implemented"; return }
