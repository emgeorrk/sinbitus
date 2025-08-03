package http

func (s *Server) setupRoutes() {
	{
		s.app.Get("/health", s.metricsCtrl.Health)
		s.app.Get("/metrics", s.metricsCtrl.PrometheusMetrics)
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
		protectedV1.Patch("/habits/:habit_id", loginWrap(s.habitCtrl.UpdateHabit))
		protectedV1.Delete("/habits/:habit_id", loginWrap(s.habitCtrl.DeleteHabit))

		protectedV1.Get("/habits/:habit_id/events", loginWrap(s.eventCtrl.GetEvents))
		protectedV1.Post("/habits/:habit_id/events", loginWrap(s.eventCtrl.CreateEvent))
		protectedV1.Patch("/habit/:habit_id/events/:event_id", loginWrap(s.eventCtrl.UpdateEvent))
		protectedV1.Delete("/habit/:habit_id/events/:event_id", loginWrap(s.eventCtrl.DeleteEvent))
	}
}
