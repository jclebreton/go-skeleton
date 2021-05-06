package endpoints

import (
	"go.uber.org/zap"
	"skeleton-go/domain/usecases"
)

type Endpoints struct {
	logger   *zap.Logger
	usecases *usecases.Usecases
}

func NewEndpoints(logger *zap.Logger, usecases *usecases.Usecases) *Endpoints {
	return &Endpoints{
		logger:   logger.Named("HTTP"),
		usecases: usecases,
	}
}
