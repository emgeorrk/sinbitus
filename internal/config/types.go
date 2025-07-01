package config

import "time"

type (
	Config struct {
		Log      Log      `yaml:"log"`
		HTTP     HTTP     `yaml:"http"`
		Postgres Postgres `yaml:"postgres"`
		JWT      JWT      `yaml:"jwt"`
	}

	Log struct {
		Level        string `yaml:"level" env:"LOG_LEVEL"`
		Timestamp    string `yaml:"timestamp"`
		TSFormat     string `yaml:"tsformat"`
		Prefix       string `yaml:"prefix"`
		Label        string `yaml:"label"`
		Source       string `yaml:"source"`
		Formatter    string `yaml:"formatter"`
		SourceFormat string `yaml:"source_format"`
		Color        string `yaml:"color"`
		EnableFxLogs string `yaml:"enable_fx_logs"`
	}

	HTTP struct {
		Host string `yaml:"host" env:"HTTP_HOST"`
		Port uint16 `yaml:"port" env:"HTTP_PORT" env-required:"true"`
	}

	Postgres struct {
		URL          string        `yaml:"url" env:"POSTGRES_URL" env-required:"true"`
		MaxPoolSize  int32         `yaml:"max_pool_size" env:"POSTGRES_POOL_SIZE"`
		ConnAttempts int           `yaml:"conn_attempts" env:"POSTGRES_CONN_ATTEMPTS"`
		ConnTimeout  time.Duration `yaml:"conn_timeout" env:"POSTGRES_CONN_TIMEOUT"`
	}

	JWT struct {
		SecretKey string        `yaml:"secret_key" env:"AUTH_SECRET_KEY" env-required:"true"`
		TTL       time.Duration `yaml:"ttl" env:"AUTH_TTL" env-required:"true"`
	}
)
