// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"github.com/gin-gonic/gin"
)

type ExecuteCommandRequest struct {
	Command string `json:"command" validate:"required"`
} // @name ExecuteCommandRequest

type ExecuteCommandResponse struct {
	Code     int32  `json:"code"`
	Response string `json:"command"`
} // @name ExecuteCommandResponse

// ExecuteCommand 			godoc
//
//	@Tags			workspace
//	@Summary		Execute command inside workspace project
//	@Description	Execute command inside workspace project
//	@Produce		json
//	@Param			workspaceId	path	string	true	"Workspace ID or Name"
//	@Param			projectId	path	string	true	"Project ID"
//	@Param			command	body	ExecuteCommandDTO	true	"Execute command"
//	@Success		200			{object}	ExecuteCommandResponseDTO
//	@Router			/workspace/{workspaceId}/{projectId}/toolbox/execute [post]
//
//	@id				GetWorkspace
func ExecuteCommand(ctx *gin.Context) {
	// workspaceId := ctx.Param("workspaceId")
	// projectId := ctx.Param("projectId")

	// var executeCommandRequest ExecuteCommandRequest
	// err := ctx.BindJSON(&executeCommandRequest)
	// if err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
	// 	return
	// }

	// server := server.GetInstance(nil)

	// //	TODO: Implement execute command logic here

	// ctx.Status(200)
}
