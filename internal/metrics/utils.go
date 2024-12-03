package metrics

import (
	"fmt"
	"strings"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func NewCounter(cfg *config.App, name string, protocol string) prometheus.Counter {
	promCounter := prometheus.CounterOpts{}

	promCounter.Name = fmt.Sprintf("%s_%s_%s_requests_total", cfg.Name, name, protocol)
	promCounter.Help = fmt.Sprintf("The total number of %s %s requests", strings.ReplaceAll(name, "_", " "), protocol)

	return promauto.NewCounter(promCounter)
}
