package config

type Databases struct {
	Redis *Redis `mapstructure:"redis"`
}
