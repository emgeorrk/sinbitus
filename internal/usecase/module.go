package usecase

import (
	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"github.com/emgeorrk/sinbitus/internal/usecase/event"
	"github.com/emgeorrk/sinbitus/internal/usecase/habit"
	"github.com/emgeorrk/sinbitus/internal/usecase/repo"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
	habit.Module,
	repo.Module,
	user.Module,
	event.Module,
)
