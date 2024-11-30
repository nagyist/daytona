// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

const PORT = 2280

func (t *Server) Start() error {
	r := gin.Default()

	r.GET("/files", listFiles)
	r.POST("/files", uploadFile)
	r.GET("/files/download", downloadFile)
	r.DELETE("/files", deleteFile)
	r.GET("/files/search", searchFiles)
	r.GET("/files/info", getFileDetails)
	r.PATCH("/files/permissions", setFilePermissions)
	r.GET("/files/find", findInFiles)
	r.POST("/files/replace", replaceInFiles)

	r.POST("/execute", executeCommand)

	r.POST("/git/clone", cloneRepository)
	r.GET("/git/status", getStatus)
	r.POST("/git/commit", commitChanges)
	r.POST("/git/push", pushChanges)
	r.GET("/git/branches", listBranches)
	r.POST("/git/branches", createBranch)
	r.GET("/git/history", getCommitHistory)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: r,
	}

	listener, err := net.Listen("tcp", httpServer.Addr)
	if err != nil {
		return err
	}

	return httpServer.Serve(listener)
}
