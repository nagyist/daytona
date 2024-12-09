// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/daytonaio/daytona/cmd/daytona/config"
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/logs"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Id           string              `json:"id" validate:"required"`
	ServerApiKey string              `json:"serverApiKey" validate:"required"`
	ServerApiUrl string              `json:"serverApiUrl" validate:"required"`
	ProvidersDir string              `json:"providersDir" validate:"required"`
	LogFile      *logs.LogFileConfig `json:"logFile" validate:"required"`
} // @name RunnerConfig

func GetConfig() (*Config, error) {
	configFilePath, err := configFilePath()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		c, err := getDefaultConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get default config: %w", err)
		}

		err = Save(*c)
		if err != nil {
			return nil, fmt.Errorf("failed to save default config file: %w", err)
		}

		return c, nil
	}

	if err != nil {
		return nil, err
	}

	var c Config
	configContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configContent, &c)
	if err != nil {
		return nil, err
	}

	if c.Id == "" {
		c.Id = uuid.NewString()
	}

	if c.LogFile == nil {
		logFilePath, err := getDefaultLogFilePath()
		if err != nil {
			log.Error("failed to get default log file path")
		}

		c.LogFile = logs.GetDefaultLogFileConfig(logFilePath)
	}

	err = Save(c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetConfigDir() (string, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "runner"), nil
}

func Save(c Config) error {
	if err := util.DirectoryValidator(&c.ProvidersDir); err != nil {
		return err
	}

	configFilePath, err := configFilePath()
	if err != nil {
		return err
	}

	configContent, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0700)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, configContent, 0600)
	if err != nil {
		return err
	}

	return nil
}

func configFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

func getDefaultLogFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "log"), nil
}

func getDefaultConfig() (*Config, error) {
	providersDir, err := getDefaultProvidersDir()
	if err != nil {
		return nil, errors.New("failed to get default providers dir")
	}

	logFilePath, err := getDefaultLogFilePath()
	if err != nil {
		log.Error("failed to get default log file path")
	}

	c := Config{
		Id:           uuid.NewString(),
		ServerApiKey: "",
		ServerApiUrl: "",
		ProvidersDir: providersDir,
		LogFile:      logs.GetDefaultLogFileConfig(logFilePath),
	}

	return &c, nil
}

func getDefaultProvidersDir() (string, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "providers"), nil
}
