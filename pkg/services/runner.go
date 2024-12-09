// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package services

import (
	"context"
	"errors"

	"github.com/daytonaio/daytona/pkg/models"
)

type IRunnerService interface {
	RegisterRunner(ctx context.Context, req RegisterRunnerDTO) (*RunnerDTO, error)
	GetRunner(ctx context.Context, runnerId string) (*RunnerDTO, error)
	ListRunners(ctx context.Context) ([]*RunnerDTO, error)
	SetRunnerMetadata(ctx context.Context, runnerId string, metadata *models.RunnerMetadata) error
	RemoveRunner(ctx context.Context, runnerId string) error
}

type RunnerDTO struct {
	models.Runner
	State models.RunnerState `json:"state" validate:"required"`
} //	@name	RunnerDTO

type RegisterRunnerDTO struct {
	Id    string `json:"id" validate:"required"`
	Alias string `json:"alias" validate:"required"`
} // @name RegisterRunnerDTO

var (
	ErrRunnerAlreadyExists = errors.New("runner already exists")
)
