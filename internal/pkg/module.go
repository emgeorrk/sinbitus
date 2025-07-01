package pkg

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/emgeorrk/sinbitus/internal/pkg/time"
	"go.uber.org/fx"
)

var Module = fx.Options(
	logger.Module,
	time.Module,
)
