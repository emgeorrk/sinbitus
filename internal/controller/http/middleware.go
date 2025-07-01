package http

import (
	"strings"
	"time"

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

func (s *Server) loginMiddleware(c fiber.Ctx) error {
	ctx := c.Context()

	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing or invalid Authorization header",
		})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := s.auth.ParseToken(ctx, tokenStr)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	claims, err := s.auth.ExtractClaims(ctx, token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "failed to extract claims",
		})
	}

	if s.auth.IsExpired(ctx, token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token expired",
		})
	}

	c.Locals("user", claims)

	return c.Next()
}
