package user

import (
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	log       *logger.Logger
	validator *validator.Validate

	auth AuthUseCase
	time TimeProvider
	user UserUseCase
}

func NewController(
	log *logger.Logger,
	auth AuthUseCase,
	clock TimeProvider,
	user UserUseCase,
) *Controller {
	return &Controller{
		log:       log,
		validator: validator.New(),
		auth:      auth,
		time:      clock,
		user:      user,
	}
}
