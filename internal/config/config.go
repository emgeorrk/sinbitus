package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/sinbitus/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
