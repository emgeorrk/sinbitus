package metrics

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Controller struct {
	metrics fiber.Handler
	time    TimeProvider
}

func NewController(time TimeProvider) *Controller {
	return &Controller{
		metrics: adaptor.HTTPHandler(promhttp.Handler()),
		time:    time,
	}
}

func (s *Controller) PrometheusMetrics(c fiber.Ctx) error {
	return s.metrics(c)
}
