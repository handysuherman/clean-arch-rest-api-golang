package config

import "time"

type Internal struct {
	ID               string        `mapstructure:"id"`
	Name             string        `mapstructure:"name"`
	DNS              string        `mapstructure:"dns"`
	LogLevel         string        `mapstructure:"log_level"`
	ApiBasePath      string        `mapstructure:"api_base_path"`
	Environment      string        `mapstructure:"environment"`
	EnableTLS        bool          `mapstructure:"enable_tls"`
	Addr             string        `mapstructure:"addr"`
	Port             int           `mapstructure:"port"`
	XApiKey          string        `mapstructure:"x_api_key"`
	OperationTimeout time.Duration `mapstructure:"operation_timeout"`
}
