package main

import (
	"errors"
	"flag"
)

type args struct {
	minCPU      int64
	maxCPU      int64
	execPath    string
	startSeed   int64
	nTopologies int64
	logsPath    string
}

func parseArgs() (*args, error) {
	var result args

	flag.Int64Var(&result.minCPU, "min-cpu", 1, "Minimum CPU count")
	flag.Int64Var(&result.maxCPU, "max-cpu", 1, "Maximum CPU count")
	flag.StringVar(&result.execPath, "exec-path", "", "Executable path")
	flag.Int64Var(&result.startSeed, "start-seed", 0, "From which seed to start")
	flag.Int64Var(&result.nTopologies, "n-topologies", 59, "Number of topologies")
	flag.StringVar(&result.logsPath, "logs-path", "./logs.sqlite", "Sqlite logs path")

	flag.Parse()

	if result.execPath == "" {
		return nil, errors.New("exec-path is required")
	}

	return &result, nil
}
