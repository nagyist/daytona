// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"github.com/daytonaio/daytona/internal/util"
	"github.com/spf13/cobra"
)

var RunnerCmd = &cobra.Command{
	Use:     "runner",
	Short:   "Manage runners",
	GroupID: util.RUNNER_GROUP,
}

func init() {
	RunnerCmd.AddGroup(&cobra.Group{ID: util.RUNNER_GROUP, Title: "Runner"})
	RunnerCmd.AddCommand(runnerListCmd)
	RunnerCmd.AddCommand(runnerRegisterCmd)
	RunnerCmd.AddCommand(runnerUnregisterCmd)
}
