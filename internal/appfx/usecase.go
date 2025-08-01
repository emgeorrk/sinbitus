package appfx

import (
	habitCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	userCtrl "github.com/emgeorrk/sinbitus/internal/controller/http/user"
	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"github.com/emgeorrk/sinbitus/internal/usecase/habit"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var UserUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(u *user.UseCase) *user.UseCase {
			return u
		},
			fx.As(new(userCtrl.UserUseCase)),
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
		),
	),
)

var HabitUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(h *habit.UseCase) *habit.UseCase {
			return h
		},
			fx.As(new(habitCtrl.HabitUseCase)),
		),
	),
)
