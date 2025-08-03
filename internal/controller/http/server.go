package http

import (
	"fmt"
	"time"

	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/controller/http/event"
	"github.com/emgeorrk/sinbitus/internal/controller/http/habit"
	"github.com/emgeorrk/sinbitus/internal/controller/http/metrics"
	"github.com/emgeorrk/sinbitus/internal/controller/http/user"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

type Server struct {
	address string
	port    uint16

	userCtrl    *user.Controller
	habitCtrl   *habit.Controller
	eventCtrl   *event.Controller
	metricsCtrl *metrics.Controller

	log *logger.Logger
	app *fiber.App

	time TimeProvider
}

func NewServer(
	metrics *metrics.Controller,
	user *user.Controller,
	habit *habit.Controller,
	event *event.Controller,
	log *logger.Logger,
	cfg *config.Config,
	clock TimeProvider,
) *Server {
	app := fiber.New(fiber.Config{
		AppName:      constants.ProjectName,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	s := &Server{
		log:         log,
		app:         app,
		time:        clock,
		address:     cfg.HTTP.Host,
		port:        cfg.HTTP.Port,
		metricsCtrl: metrics,
		userCtrl:    user,
		habitCtrl:   habit,
		eventCtrl:   event,
	}

	s.setupMiddleware()

	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.app.Use(
		recover.New(),
		requestid.New(),
		s.loggerMiddleware,
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		}),
	)
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.address, s.port)

	s.log.Info("HTTP server is listening on " + addr)

	return s.app.Listen(addr)
}

func (s *Server) Stop() error {
	return s.app.Shutdown()
}
