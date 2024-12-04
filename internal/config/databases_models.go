package config

type Databases struct {
	MySQL *MySQL `mapstructure:"mysql"`
	Redis *Redis `mapstructure:"redis"`
}
