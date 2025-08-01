package http

import (
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func (s *Server) loggerMiddleware(c fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	s.log.Info("REQUEST INFO",
		"method", c.Method(),
		"path", c.Path(),
		"status", c.Response().StatusCode(),
		"latency", time.Since(start).String(),
		"ip", c.IP(),
		"request_id", requestid.FromContext(c),
	)

	return err
}

func loginWrap(fn func(f fiber.Ctx, c entity.UserClaims) error) fiber.Handler {
	return func(c fiber.Ctx) error {
		userClaims, ok := c.Locals("user").(entity.UserClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to retrieve user claims",
			})
		}

		return fn(c, userClaims)
	}
}
