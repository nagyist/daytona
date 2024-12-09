// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetRunner 			godoc
//
//	@Tags			runner
//	@Summary		Get a runner
//	@Description	Get a runner
//	@Param			runnerId	path	string	true	"Runner ID"
//	@Produce		json
//	@Success		200	{object}	RunnerDTO
//	@Router			/runner/{runnerId} [get]
//
//	@id				GetRunner
func GetRunner(ctx *gin.Context) {
	runnerId := ctx.Param("runnerId")

	server := server.GetInstance(nil)

	r, err := server.RunnerService.GetRunner(ctx.Request.Context(), runnerId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to remove runner: %w", err))
		return
	}

	ctx.JSON(200, r)
}
