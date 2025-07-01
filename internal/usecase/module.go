package usecase

import (
	"github.com/emgeorrk/sinbitus/internal/usecase/auth"
	"github.com/emgeorrk/sinbitus/internal/usecase/repo"
	"github.com/emgeorrk/sinbitus/internal/usecase/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
	repo.Module,
	user.Module,
)
