package config

type TLSKey string

const (
	TLS_APP   TLSKey = "app"
	TLS_MYSQL TLSKey = "mysql"
	TLS_REDIS TLSKey = "redis"
)

type TLS struct {
	App   *Certs `mapstructure:"app"`
	MySQL *Certs `mapstructure:"mysql"`
	Redis *Certs `mapstructure:"redis"`
}

type Certs struct {
	CaPath   string `mapstructure:"ca_path"`
	CertPath string `mapstructure:"cert_path"`
	KeyPath  string `mapstructure:"key_path"`
}
