package middleware

import (
	"strings"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
)

type middlewareManager struct {
	log logger.Logger
}

func NewMiddlewareManager(log logger.Logger) *middlewareManager {
	return &middlewareManager{log: log}
}

func (mw *middlewareManager) checkIgnoredURI(requestURI string, uriList []string) bool {
	for _, s := range uriList {
		if strings.Contains(requestURI, s) {
			return true
		}
	}

	return false
}
