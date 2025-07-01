package repo

import (
	"github.com/emgeorrk/sinbitus/internal/infra/postgres"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRepo),
)

type Repo struct {
	log *logger.Logger
	*postgres.Postgres
}

func NewRepo(log *logger.Logger, postgres *postgres.Postgres) *Repo {
	return &Repo{
		log:      log,
		Postgres: postgres,
	}
}
