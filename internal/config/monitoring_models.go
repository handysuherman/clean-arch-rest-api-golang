package config

type Monitoring struct {
	Probes *Probes `mapstructure:"probes"`
	Jaeger *Jaeger `mapstructure:"jaeger"`
}

type Probes struct {
	ReadinessPath string      `mapstructure:"readiness_path"`
	LivenessPath  string      `mapstructure:"liveness_path"`
	CheckInterval int         `mapstructure:"check_interval"`
	Port          string      `mapstructure:"port"`
	Prof          string      `mapstructure:"pprof"`
	Prometheus    *Prometheus `mapstructure:"prometheus"`
}

type Prometheus struct {
	Port string `mapstructure:"port"`
	Path string `mapstructure:"path"`
}

type Jaeger struct {
	HostPort string `mapstructure:"host_port"`
	Enable   bool   `mapstructure:"enable"`
	Logspan  bool   `mapstructure:"log_span"`
}
