package controller

import (
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	"go.uber.org/fx"
)

var Module = fx.Options(
	http.Module,
)
