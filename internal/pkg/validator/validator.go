package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewValidator),
)

func NewValidator() *validator.Validate {
	v := validator.New()

	return v
}
