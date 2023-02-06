package web

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/go-project-skeleton/src/adapters/httpserver"
)

type appConfig struct {
	ServerConfig httpserver.ServerConfig
}

func loadAppConfig() (appConfig, error) {
	var config appConfig

	if err := envconfig.Process("", &config); err != nil {
		return appConfig{}, fmt.Errorf("failed to load config - %w", err)
	}

	return config, nil
}
