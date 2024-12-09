// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"context"
	"os"
	"os/signal"

	"github.com/daytonaio/daytona/pkg/jobs"
	"github.com/daytonaio/daytona/pkg/jobs/build"
	"github.com/daytonaio/daytona/pkg/jobs/target"
	"github.com/daytonaio/daytona/pkg/jobs/workspace"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/runner/providermanager"
	"github.com/daytonaio/daytona/pkg/scheduler"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
)

// TODO: add lock when running interval func
// 1 second interval
const DEFAULT_JOB_POLL_INTERVAL = "*/1 * * * * *"

type IRunner interface {
	StartRunner(ctx context.Context) error
	CheckAndRunJobs(ctx context.Context) error
}

type RunnerConfig struct {
	models.Runner
	ProviderManager providermanager.IProviderManager
	ProvidersDir    string

	ListPendingJobs func(ctx context.Context) ([]*models.Job, error)
	UpdateJobState  func(ctx context.Context, job *models.Job, state models.JobState, err *error) error

	WorkspaceJobFactory workspace.IWorkspaceJobFactory
	TargetJobFactory    target.ITargetJobFactory
	BuildJobFactory     build.IBuildJobFactory
}

func NewRunner(config RunnerConfig) IRunner {
	return &Runner{
		Runner:          config.Runner,
		providerManager: config.ProviderManager,
		providersDir:    config.ProvidersDir,

		listPendingJobs: config.ListPendingJobs,
		updateJobState:  config.UpdateJobState,

		workspaceJobFactory: config.WorkspaceJobFactory,
		targetJobFactory:    config.TargetJobFactory,
		buildJobFactory:     config.BuildJobFactory,
	}
}

type Runner struct {
	models.Runner

	providerManager providermanager.IProviderManager
	providersDir    string

	listPendingJobs func(ctx context.Context) ([]*models.Job, error)
	updateJobState  func(ctx context.Context, job *models.Job, state models.JobState, err *error) error

	workspaceJobFactory workspace.IWorkspaceJobFactory
	targetJobFactory    target.ITargetJobFactory
	buildJobFactory     build.IBuildJobFactory
}

func (s *Runner) StartRunner(ctx context.Context) error {
	log.Info("Starting runner")
	go func() {
		interruptChannel := make(chan os.Signal, 1)
		signal.Notify(interruptChannel, os.Interrupt)

		for range interruptChannel {
			plugin.CleanupClients()
		}
	}()

	// Terminate orphaned provider processes
	err := s.providerManager.TerminateProviderProcesses(s.providersDir)
	if err != nil {
		log.Errorf("Failed to terminate orphaned provider processes: %s", err)
	}

	err = s.downloadDefaultProviders()
	if err != nil {
		return err
	}

	err = s.registerProviders()
	if err != nil {
		return err
	}

	scheduler := scheduler.NewCronScheduler()

	err = scheduler.AddFunc(DEFAULT_JOB_POLL_INTERVAL, func() {
		err := s.CheckAndRunJobs(ctx)
		if err != nil {
			log.Error(err)
		}
	})
	if err != nil {
		return err
	}

	scheduler.Start()

	log.Info("Runner started")
	return nil
}

func (s *Runner) CheckAndRunJobs(ctx context.Context) error {
	jobs, err := s.listPendingJobs(ctx)
	if err != nil {
		return err
	}

	// goroutines, sync group
	for _, job := range jobs {
		err = s.runJob(ctx, job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Runner) runJob(ctx context.Context, j *models.Job) error {
	var job jobs.IJob

	err := s.updateJobState(ctx, j, models.JobStateRunning, nil)
	if err != nil {
		return err
	}

	switch j.ResourceType {
	case models.ResourceTypeWorkspace:
		job = s.workspaceJobFactory.Create(*j)
	case models.ResourceTypeTarget:
		job = s.targetJobFactory.Create(*j)
	case models.ResourceTypeBuild:
		job = s.buildJobFactory.Create(*j)
	}

	err = job.Execute(ctx)
	if err != nil {
		return s.updateJobState(ctx, j, models.JobStateError, &err)
	}

	return s.updateJobState(ctx, j, models.JobStateSuccess, nil)
}
