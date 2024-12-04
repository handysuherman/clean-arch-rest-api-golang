package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func New(path string) (*Config, error) {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)

	viper.SetConfigName(strings.TrimSuffix(fileName, ext))
	viper.SetConfigType(strings.TrimPrefix(ext, "."))
	viper.AddConfigPath(filepath.Dir(fileName))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig.err: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal.err: %v", err)
	}

	return &config, nil
}
