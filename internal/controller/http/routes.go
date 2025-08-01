package http

import (
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) setupRoutes() {
	{
		s.app.Get("/health", s.metricsCtrl.Health)
		s.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	}

	v1 := s.app.Group("/api/v1")
	{
		v1.Post("/signup", s.userCtrl.SignUp)
		v1.Post("/login", s.userCtrl.Login)
	}

	protectedV1 := v1.Use(s.userCtrl.LoginMiddleware)
	{
		protectedV1.Get("/profile", loginWrap(s.userCtrl.GetProfile))

		protectedV1.Get("/habits", loginWrap(s.habitCtrl.GetHabits))
		protectedV1.Post("/habits", loginWrap(s.habitCtrl.CreateHabit))
		protectedV1.Patch("/habits/:id", loginWrap(s.habitCtrl.UpdateHabit))
		protectedV1.Delete("/habits/:id", loginWrap(s.habitCtrl.DeleteHabit))

		// protectedV1.Get("/habits/:habit_id/events", loginWrap(s.habitCtrl.GetHabitEvents))
		// protectedV1.Post("/habits/:habit_id/events", loginWrap(s.habitCtrl.CreateHabitEvent))
		// protectedV1.Patch("/habit/:habit_id/events/:event_id", loginWrap(s.habitCtrl.UpdateHabitEvent))
		// protectedV1.Delete("/habit/:habit_id/events/:event_id", loginWrap(s.habitCtrl.DeleteHabitEvent))
	}
}
