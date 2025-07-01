package infra

import (
	"github.com/emgeorrk/sinbitus/internal/infra/postgres"
	"go.uber.org/fx"
)

var Module = fx.Options(
	postgres.Module,
)
