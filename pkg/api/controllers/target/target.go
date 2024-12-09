// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/api/util"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/services"
	"github.com/daytonaio/daytona/pkg/stores"
	"github.com/gin-gonic/gin"
)

// GetTarget 			godoc
//
//	@Tags			target
//	@Summary		Get target info
//	@Description	Get target info
//	@Produce		json
//	@Param			targetId	path		string	true	"Target ID or Name"
//	@Param			showOptions	query		bool	true	"Show target config options"
//	@Success		200			{object}	TargetDTO
//	@Router			/target/{targetId} [get]
//
//	@id				GetTarget
func GetTarget(ctx *gin.Context) {
	targetId := ctx.Param("targetId")
	showTargetConfigOptionsQuery := ctx.Query("showOptions")
	var showTargetConfigOptions bool
	if showTargetConfigOptionsQuery == "true" {
		showTargetConfigOptions = true
	}

	server := server.GetInstance(nil)

	t, err := server.TargetService.GetTarget(ctx.Request.Context(), &stores.TargetFilter{IdOrName: &targetId}, services.TargetRetrievalParams{})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if stores.IsTargetNotFound(err) || services.IsTargetDeleted(err) {
			statusCode = http.StatusNotFound
		}
		ctx.AbortWithError(statusCode, fmt.Errorf("failed to find target: %w", err))
		return
	}

	if !showTargetConfigOptions {
		t.TargetConfig.Options = ""
	}

	for i := range t.Workspaces {
		util.HideDaytonaEnvVars(&t.Workspaces[i].EnvVars)
	}

	util.HideDaytonaEnvVars(&t.EnvVars)

	ctx.JSON(200, t)
}

// ListTargets 			godoc
//
//	@Tags			target
//	@Summary		List targets
//	@Description	List targets
//	@Param			showOptions	query	bool	true	"Show target config options"
//	@Produce		json
//	@Success		200	{array}	TargetDTO
//	@Router			/target [get]
//
//	@id				ListTargets
func ListTargets(ctx *gin.Context) {
	server := server.GetInstance(nil)
	showTargetConfigOptionsQuery := ctx.Query("showOptions")
	var showTargetConfigOptions bool
	if showTargetConfigOptionsQuery == "true" {
		showTargetConfigOptions = true
	}

	targetList, err := server.TargetService.ListTargets(ctx.Request.Context(), nil, services.TargetRetrievalParams{})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to list targets: %w", err))
		return
	}

	for i := range targetList {
		if !showTargetConfigOptions {
			targetList[i].TargetConfig.Options = ""
		}

		for j := range targetList[i].Workspaces {
			util.HideDaytonaEnvVars(&targetList[i].Workspaces[j].EnvVars)
		}

		util.HideDaytonaEnvVars(&targetList[i].EnvVars)
	}

	ctx.JSON(200, targetList)
}
