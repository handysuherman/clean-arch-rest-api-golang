package config

type Config struct {
	Services   *Services   `mapstructure:"services"`
	Databases  *Databases  `mapstructure:"databases"`
	TLS        *TLS        `mapstructure:"tls"`
	Monitoring *Monitoring `mapstructure:"monitoring"`
}
