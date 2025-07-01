package auth

import (
	"time"

	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUseCase),
)

type UseCase struct {
	log  *logger.Logger
	time TimeProvider

	secretKey []byte
	ttl       time.Duration
}

func NewUseCase(log *logger.Logger, cfg *config.Config, time TimeProvider) *UseCase {
	return &UseCase{
		log:       log,
		time:      time,
		secretKey: []byte(cfg.JWT.SecretKey),
		ttl:       cfg.JWT.TTL,
	}
}
