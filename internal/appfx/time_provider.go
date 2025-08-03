package appfx

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	"github.com/emgeorrk/sinbitus/internal/pkg/time"
	"go.uber.org/fx"

	eventsCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/event"
	habitCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	metricsCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/metrics"
	userCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/user"
	authUC "github.com/emgeorrk/sinbitus/internal/usecase/auth"
	eventUC "github.com/emgeorrk/sinbitus/internal/usecase/event"
	habitUC "github.com/emgeorrk/sinbitus/internal/usecase/habit"
	userUC "github.com/emgeorrk/sinbitus/internal/usecase/user"
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
			fx.As(new(eventsCtrl.TimeProvider)),
			fx.As(new(authUC.TimeProvider)),
			fx.As(new(userUC.TimeProvider)),
			fx.As(new(habitUC.TimeProvider)),
			fx.As(new(eventUC.TimeProvider)),
		),
	),
)
