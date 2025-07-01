package http

import (
	"fmt"
	"time"

	"github.com/emgeorrk/sinbitus/internal/config"
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewServer),
)

type Server struct {
	log *logger.Logger
	app *fiber.App

	auth AuthUseCase
	time TimeProvider
	user UserUseCase

	address string
	port    uint16
}

func NewServer(
	log *logger.Logger,
	cfg *config.Config,
	auth AuthUseCase,
	clock TimeProvider,
	user UserUseCase,
) *Server {
	app := fiber.New(fiber.Config{
		AppName:      constants.ProjectName,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	s := &Server{
		log:     log,
		app:     app,
		auth:    auth,
		time:    clock,
		user:    user,
		address: cfg.HTTP.Host,
		port:    cfg.HTTP.Port,
	}

	s.setupMiddleware()

	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.app.Use(recover.New())

	s.app.Use(requestid.New())

	s.app.Use(s.loggerMiddleware)

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
}

func (s *Server) setupRoutes() {
	s.app.Get("/health", s.health)

	s.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	v1 := s.app.Group("/api/v1")
	{
		v1.Post("/signup", s.signup)
		v1.Post("/login", s.login)
	}

	// protectedV1 := s.app.Group("/api/v1").Use(s.loginMiddleware)
	// {
	// 	protectedV1.Get("/profile", s.profileHandler)
	//
	// 	protectedV1.Post("/habits", s.createHabit)
	// 	protectedV1.Get("/habits", s.getHabits)
	//
	// 	protectedV1.Get("/habits/:id", s.getHabitHandler)
	// 	protectedV1.Put("/habits/:id", s.updateHabitHandler)
	// 	protectedV1.Delete("/habits/:id", s.deleteHabitHandler)
	//
	// 	protectedV1.Post("/habits/:id/track", s.trackHabitHandler)
	// 	protectedV1.Get("/habits/:id/track", s.getHabitTrackHandler)
	// }
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.address, s.port)

	s.log.Info("HTTP server is listening on " + addr)

	return s.app.Listen(addr)
}

func (s *Server) Stop() error {
	return s.app.Shutdown()
}

func (s *Server) health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "ok",
		"timestamp": s.time.Now().UTC(),
		"service":   constants.ProjectName,
	})
}
