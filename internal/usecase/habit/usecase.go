package habit

import "github.com/emgeorrk/sinbitus/internal/pkg/logger"

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
