package user

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUseCase),
)

type UseCase struct {
	log  *logger.Logger
	repo Repository
}

func NewUseCase(log *logger.Logger, repo Repository) *UseCase {
	return &UseCase{
		log:  log,
		repo: repo,
	}
}
