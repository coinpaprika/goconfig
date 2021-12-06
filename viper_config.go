package cpcfg

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

const (
	configName = "config"
)

type ViperConfig struct {
	*viper.Viper
}

func NewViperConfig(configDirVar string) (*ViperConfig, error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath("./.config/") // personal local config for project
	if !strings.HasPrefix(configDirVar, "$") {
		configDirVar = fmt.Sprintf("$%s", configDirVar)
	}
	viper.AddConfigPath(configDirVar)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error while reading config via viper, set %s environment variable or create local config in .config/ directory: %w", configDirVar, err)
	}

	return &ViperConfig{
		viper.GetViper(),
	}, nil
}

func (c *ViperConfig) Dir() string {
	return path.Dir(c.Viper.ConfigFileUsed())
}
