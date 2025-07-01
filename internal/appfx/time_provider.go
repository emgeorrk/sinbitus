package appfx

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	"github.com/emgeorrk/sinbitus/internal/pkg/time"
	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"go.uber.org/fx"
)

var TimeProvider = fx.Options(
	fx.Provide(
		fx.Annotate(func(t *time.Provider) *time.Provider {
			return t
		},
			fx.As(new(http.TimeProvider)),
			fx.As(new(auth.TimeProvider)),
		),
	),
)
