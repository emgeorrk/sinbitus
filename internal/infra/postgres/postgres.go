package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
)

type Postgres struct {
	log *logger.Logger

	maxPoolSize  int32
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(lc fx.Lifecycle, log *logger.Logger, cfg *config.Config) (*Postgres, error) {
	pg := &Postgres{
		log:          log,
		maxPoolSize:  cfg.Postgres.MaxPoolSize,
		connAttempts: cfg.Postgres.ConnAttempts,
		connTimeout:  cfg.Postgres.ConnTimeout,
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(cfg.Postgres.URL)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = pg.maxPoolSize

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		pg.log.Info(fmt.Sprintf("Postgres is trying to connect, attempts left: %d", pg.connAttempts))

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - New - connAttempts == 0: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pg.log.Info("Closing Postgres pool")
			pg.Close()
			return nil
		},
	})

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
