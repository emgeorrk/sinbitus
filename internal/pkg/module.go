package pkg

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/emgeorrk/sinbitus/internal/pkg/time"
	"github.com/emgeorrk/sinbitus/internal/pkg/validator"
	"go.uber.org/fx"
)

var Module = fx.Options(
	logger.Module,
	time.Module,
	validator.Module,
)
