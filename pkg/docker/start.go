// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package docker

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/daytonaio/daytona/pkg/build/detect"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/provider/util"
	"github.com/docker/docker/api/types/container"
)

func (d *DockerClient) StartWorkspace(opts *CreateWorkspaceOptions, daytonaDownloadUrl string) error {
	var err error
	containerUser := opts.Workspace.User

	builderType, err := detect.DetectWorkspaceBuilderType(opts.Workspace.BuildConfig, opts.WorkspaceDir, opts.SshClient)
	if err != nil {
		return err
	}

	switch builderType {
	case detect.BuilderTypeDevcontainer:
		var remoteUser RemoteUser
		remoteUser, err = d.startDevcontainerWorkspace(opts)
		containerUser = string(remoteUser)
	case detect.BuilderTypeImage:
		err = d.startImageWorkspace(opts)
	default:
		return fmt.Errorf("unknown builder type: %s", builderType)
	}

	if err != nil {
		return err
	}

	if len(opts.ContainerRegistries) == 0 {
		return d.startDaytonaAgent(opts.Workspace, containerUser, daytonaDownloadUrl, opts.LogWriter)
	}

	containerRegistriesConfigContent := make(map[string]string)
	for _, cr := range opts.ContainerRegistries {
		auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", cr.Username, cr.Password)))
		containerRegistriesConfigContent[cr.Server] = fmt.Sprintf(`{"auth":"%s"}`, auth)
	}

	jsonContent, err := json.Marshal(containerRegistriesConfigContent)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("~/.docker/config.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(jsonContent)
	if err != nil {
		return err
	}

	return d.startDaytonaAgent(opts.Workspace, containerUser, daytonaDownloadUrl, opts.LogWriter)
}

func (d *DockerClient) startDaytonaAgent(w *models.Workspace, containerUser, daytonaDownloadUrl string, logWriter io.Writer) error {
	errChan := make(chan error)

	go func() {
		result, err := d.ExecSync(d.GetWorkspaceContainerName(w), container.ExecOptions{
			Cmd:          []string{"bash", "-c", util.GetWorkspaceStartScript(daytonaDownloadUrl, w.ApiKey)},
			AttachStdout: true,
			AttachStderr: true,
			User:         containerUser,
		}, logWriter)
		if err != nil {
			errChan <- err
		}

		if result.ExitCode != 0 {
			errChan <- errors.New(result.StdErr)
		}
	}()

	go func() {
		// TODO: Figure out how to check if the agent is running here
		time.Sleep(5 * time.Second)
		errChan <- nil
	}()

	return <-errChan
}
