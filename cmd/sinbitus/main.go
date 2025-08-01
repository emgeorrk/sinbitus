package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emgeorrk/sinbitus/internal/appfx"
	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/controller"
	"github.com/emgeorrk/sinbitus/internal/controller/http"
	"github.com/emgeorrk/sinbitus/internal/infra"
	"github.com/emgeorrk/sinbitus/internal/pkg"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/emgeorrk/sinbitus/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	app := fx.New(
		config.Module,
		pkg.Module,
		appfx.Module,
		usecase.Module,
		infra.Module,
		controller.Module,
		fx.Invoke(Start),
	)

	if err := app.Err(); err != nil {
		os.Exit(1)
	}

	go func() {
		if err := app.Start(context.Background()); err != nil {
			fmt.Println("Error starting app:", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := app.Stop(ctx)
	if err != nil {
		fmt.Println("Error stopping app:", err)
	}
}

func Start(lc fx.Lifecycle, log *logger.Logger, s *http.Server) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Application started ☀️ ")

			go func() {
				err := s.Start()
				if err != nil {
					log.Error("Failed to start http server", log.Err(err))
					panic("failed to start http server")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Application is shutting down... ☄️ ")

			if err := s.Stop(); err != nil {
				log.Error("Failed to stop http server", log.Err(err))
				return fmt.Errorf("failed to stop http server: %w", err)
			}

			return nil
		},
	})

	return nil
}
