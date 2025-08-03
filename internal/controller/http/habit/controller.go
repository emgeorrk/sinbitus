package habit

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	log       *logger.Logger
	validator *validator.Validate

	auth   AuthUseCase
	time   TimeProvider
	habits HabitsUseCase
}

func NewController(
	log *logger.Logger,
	validator *validator.Validate,
	auth AuthUseCase,
	clock TimeProvider,
	habit HabitsUseCase,
) *Controller {
	return &Controller{
		log:       log,
		validator: validator,
		auth:      auth,
		time:      clock,
		habits:    habit,
	}
}
