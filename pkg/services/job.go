// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package services

import (
	"errors"

	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/stores"
)

type IJobService interface {
	Create(job *models.Job) error
	Update(job *models.Job) error
	Find(filter *stores.JobFilter) (*models.Job, error)
	List(filter *stores.JobFilter) ([]*models.Job, error)
	Delete(job *models.Job) error
}

var (
	ErrInvalidResourceJobAction = errors.New("invalid job action for resource")
)

func IsInvalidResourceJobAction(err error) bool {
	return err.Error() == ErrInvalidResourceJobAction.Error()
}