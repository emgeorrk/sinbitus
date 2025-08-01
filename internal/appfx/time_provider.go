package appfx

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	habitCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	metricsCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/metrics"
	userCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/user"
	"github.com/emgeorrk/sinbitus/internal/pkg/time"
	authUC "github.com/emgeorrk/sinbitus/internal/usecase/auth"
	userUC "github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var TimeProvider = fx.Options(
	fx.Provide(
		fx.Annotate(func(t *time.Provider) *time.Provider {
			return t
		},
			fx.As(new(http.TimeProvider)),
			fx.As(new(userCtrl.TimeProvider)),
			fx.As(new(habitCtrl.TimeProvider)),
			fx.As(new(metricsCtrl.TimeProvider)),
			fx.As(new(authUC.TimeProvider)),
			fx.As(new(userUC.TimeProvider)),
		),
	),
)
