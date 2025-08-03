package appfx

import (
	"github.com/emgeorrk/sinbitus/internal/usecase/event"
	"github.com/emgeorrk/sinbitus/internal/usecase/habit"
	"github.com/emgeorrk/sinbitus/internal/usecase/repo"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var Repo = fx.Options(
	fx.Provide(
		fx.Annotate(func(r *repo.Repo) *repo.Repo {
			return r
		},
			fx.As(new(user.Repository)),
			fx.As(new(habit.Repository)),
			fx.As(new(event.Repository)),
		),
	),
)
