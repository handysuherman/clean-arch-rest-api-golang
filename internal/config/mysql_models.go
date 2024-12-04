package config

type MySQL struct {
	Driver       string `mapstructure:"driver"`
	Source       string `mapstructure:"source"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	EnableTLS    bool   `mapstructure:"enable_tls"`
	MigrationURL string `mapstructure:"migration_url"`
}
