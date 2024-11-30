// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"os/exec"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
)

type ExecuteRequest struct {
	Command string `json:"command" binding:"required"`
}

type ExecuteResponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func executeCommand(c *gin.Context) {
	var request ExecuteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ExecuteResponse{
			Code:   400,
			Result: "Invalid request: command is required",
		})
		return
	}

	cmdParts := strings.Fields(request.Command)
	if len(cmdParts) == 0 {
		c.JSON(400, ExecuteResponse{
			Code:   400,
			Result: "Empty command",
		})
		return
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	output, err := cmd.CombinedOutput()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
		}
	}

	c.JSON(200, ExecuteResponse{
		Code:   exitCode,
		Result: string(output),
	})
}
