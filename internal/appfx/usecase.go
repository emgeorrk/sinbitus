package appfx

import (
	"go.uber.org/fx"

	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"github.com/emgeorrk/sinbitus/internal/usecase/event"
	"github.com/emgeorrk/sinbitus/internal/usecase/habit"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"

	eventCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/event"
	habitCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	userCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/user"
)

var UserUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(u *user.UseCase) *user.UseCase {
			return u
		},
			fx.As(new(userCtrl.UsersUseCase)),
		),
	),
)

var AuthUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(a *auth.UseCase) *auth.UseCase {
			return a
		},
			fx.As(new(userCtrl.AuthUseCase)),
			fx.As(new(habitCtrl.AuthUseCase)),
			fx.As(new(eventCtrl.AuthUseCase)),
		),
	),
)

var HabitUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(h *habit.UseCase) *habit.UseCase {
			return h
		},
			fx.As(new(habitCtrl.HabitsUseCase)),
		),
	),
)

var EventUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(e *event.UseCase) *event.UseCase {
			return e
		},
			fx.As(new(eventCtrl.EventsUseCase)),
		),
	),
)
