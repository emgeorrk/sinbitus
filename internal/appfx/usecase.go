package appfx

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var UserUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(u *user.UseCase) *user.UseCase {
			return u
		},
			fx.As(new(http.UserUseCase)),
		),
	),
)

var AuthUseCase = fx.Options(
	fx.Provide(
		fx.Annotate(func(a *auth.UseCase) *auth.UseCase {
			return a
		},
			fx.As(new(http.AuthUseCase)),
		),
	),
)
