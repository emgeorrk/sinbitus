package metrics

import (
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/gofiber/fiber/v3"
)

func (s *Controller) Health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "ok",
		"timestamp": s.time.Now().UTC(),
		"service":   constants.ProjectName,
	})
}
