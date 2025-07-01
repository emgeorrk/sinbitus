package appfx

import "go.uber.org/fx"

var Module = fx.Options(
	TimeProvider,
	Repo,
	UserUseCase,
	AuthUseCase,
)
