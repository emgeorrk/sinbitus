package time

import (
	"time"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProvider),
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (t *Provider) Now() time.Time {
	return time.Now()
}
