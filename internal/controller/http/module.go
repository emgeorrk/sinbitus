package http

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	"github.com/emgeorrk/sinbitus/internal/controller/http/metrics"
	"github.com/emgeorrk/sinbitus/internal/controller/http/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Provide(metrics.NewController),
	fx.Provide(user.NewController),
	fx.Provide(habit.NewController),
)
