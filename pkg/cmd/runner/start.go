// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"github.com/spf13/cobra"
)

var startRunnerCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the runner",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
