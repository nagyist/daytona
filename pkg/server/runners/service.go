// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runners

import (
	"context"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/services"
	"github.com/daytonaio/daytona/pkg/stores"
	"github.com/daytonaio/daytona/pkg/telemetry"
)

type RunnerServiceConfig struct {
	RunnerStore         stores.RunnerStore
	RunnerMetadataStore stores.RunnerMetadataStore

	GenerateApiKey      func(ctx context.Context, name string) (string, error)
	RevokeApiKey        func(ctx context.Context, name string) error
	TrackTelemetryEvent func(event telemetry.ServerEvent, clientId string, props map[string]interface{}) error
}

func NewRunnerService(config RunnerServiceConfig) services.IRunnerService {
	return &RunnerService{
		runnerStore:         config.RunnerStore,
		runnerMetadataStore: config.RunnerMetadataStore,

		generateApiKey:      config.GenerateApiKey,
		revokeApiKey:        config.RevokeApiKey,
		trackTelemetryEvent: config.TrackTelemetryEvent,
	}
}

type RunnerService struct {
	runnerStore         stores.RunnerStore
	runnerMetadataStore stores.RunnerMetadataStore

	generateApiKey      func(ctx context.Context, name string) (string, error)
	revokeApiKey        func(ctx context.Context, name string) error
	trackTelemetryEvent func(event telemetry.ServerEvent, clientId string, props map[string]interface{}) error
}

func (s *RunnerService) GetRunner(ctx context.Context, runnerId string) (*services.RunnerDTO, error) {
	runner, err := s.runnerStore.Find(ctx, runnerId)
	if err != nil {
		return nil, stores.ErrRunnerNotFound
	}

	return &services.RunnerDTO{
		Runner: *runner,
		State:  runner.GetState(),
	}, nil
}

func (s *RunnerService) ListRunners(ctx context.Context) ([]*services.RunnerDTO, error) {
	runners, err := s.runnerStore.List(ctx)
	if err != nil {
		return nil, err
	}

	return util.ArrayMap(runners, func(runner *models.Runner) *services.RunnerDTO {
		return &services.RunnerDTO{
			Runner: *runner,
			State:  runner.GetState(),
		}
	}), nil
}

func (s *RunnerService) ListProviders(ctx context.Context) ([]models.ProviderInfo, error) {
	metadatas, err := s.runnerMetadataStore.List(ctx)
	if err != nil {
		return nil, err
	}

	providers := []models.ProviderInfo{}
	for _, metadata := range metadatas {
		providers = append(providers, metadata.Providers...)
	}

	return providers, nil
}

func (s *RunnerService) SetRunnerMetadata(ctx context.Context, runnerId string, metadata *models.RunnerMetadata) error {
	m, err := s.runnerMetadataStore.Find(ctx, runnerId)
	if err != nil {
		return stores.ErrRunnerMetadataNotFound
	}

	m.Uptime = metadata.Uptime
	m.RunningJobs = metadata.RunningJobs
	m.Providers = metadata.Providers
	m.UpdatedAt = metadata.UpdatedAt
	return s.runnerMetadataStore.Save(ctx, m)
}
