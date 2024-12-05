package usecases

import (
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
)

type usecaseImpl struct {
	log        logger.Logger
	cfg        *config.Config
	repository repository.Repository
}

func New(
	log logger.Logger,
	cfg *config.Config,
	repository repository.Repository,
) domain.Usecase {
	return &usecaseImpl{
		log:        log.WithPrefix(fmt.Sprintf("%s-%s", "consumers", constants.Usecase)),
		cfg:        cfg,
		repository: repository,
	}
}

func (u *usecaseImpl) errorResponse(span opentracing.Span, details string, err error) error {
	errfmt := fmt.Errorf("%s: %v", details, err)
	u.log.Warn(errfmt)
	tracing.TraceWithError(span, errfmt)

	return err
}
