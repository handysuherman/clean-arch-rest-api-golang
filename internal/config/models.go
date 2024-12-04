package config

type Config struct {
	App        *App        `mapstructure:"app"`
	Monitoring *Monitoring `mapstructure:"monitoring"`
}
