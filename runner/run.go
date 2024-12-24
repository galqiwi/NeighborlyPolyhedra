package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func run(executable string, cpu int64, seed int64, topology int64) (string, error) {
	cmd := exec.Command("taskset", "--cpu-list", fmt.Sprint(cpu), executable)

	cmd.Stdin = strings.NewReader(fmt.Sprintf("%d\n%d\n", seed, topology))

	stdoutBuilder := strings.Builder{}
	stderrBuilder := strings.Builder{}

	cmd.Stdout = &stdoutBuilder
	cmd.Stderr = &stderrBuilder

	err := cmd.Run()
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("%s\n%s", stdoutBuilder.String(), stderrBuilder.String()))
	}

	stdout := stdoutBuilder.String()

	return stdout, nil
}

func runAndSaveStdout(logsFile string, executable string, cpu int64, seed int64, topology int64) error {
	stdout, err := run(executable, cpu, seed, topology)
	if err != nil {
		return err
	}

	log, err := NewSqliteLog(logsFile)
	if err != nil {
		return err
	}
	defer log.Close()

	return log.Write(stdout)
}
