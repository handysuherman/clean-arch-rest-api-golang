package config

type App struct {
	ID          string `mapstructure:"id"`
	Name        string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
	Environment string `mapstructure:"environment"`
	BasePath    string `mapstructure:"base_path"`
}
