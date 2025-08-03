package event

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	log       *logger.Logger
	validator *validator.Validate

	auth   AuthUseCase
	time   TimeProvider
	events EventsUseCase
}

func NewController(
	log *logger.Logger,
	validator *validator.Validate,
	auth AuthUseCase,
	clock TimeProvider,
	events EventsUseCase,
) *Controller {
	return &Controller{
		log:       log,
		validator: validator,
		auth:      auth,
		time:      clock,
		events:    events,
	}
}
