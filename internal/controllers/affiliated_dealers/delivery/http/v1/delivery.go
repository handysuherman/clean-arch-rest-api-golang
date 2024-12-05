package v1handler

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-playground/validator/v10"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/metrics"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
)

type Handler struct {
	log       logger.Logger
	cfg       *config.Config
	usecase   domain.Usecase
	server    *server.Hertz
	metrics   *metrics.Metrics
	validator *validator.Validate
}

func New(
	log logger.Logger,
	cfg *config.Config,
	usecase domain.Usecase,
	server *server.Hertz,
	metrics *metrics.Metrics,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		log:       log.WithPrefix(fmt.Sprintf("%s-%s", "affiliated-dealers", constants.Handler)),
		cfg:       cfg,
		usecase:   usecase,
		server:    server,
		metrics:   metrics,
		validator: validator,
	}
}
