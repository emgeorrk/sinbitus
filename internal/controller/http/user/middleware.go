package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Controller) LoginMiddleware(c fiber.Ctx) error {
	ctx := c.Context()

	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing or invalid Authorization header",
		})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := s.auth.ParseToken(ctx, tokenStr)
	if err != nil {
		var errorMsg string
		switch {
		case errors.Is(err, jwt.ErrSignatureInvalid):
			errorMsg = "invalid token signature"
		case errors.Is(err, jwt.ErrTokenExpired):
			errorMsg = "token has expired"
		default:
			s.log.Error("Error parsing token", s.log.Err(err))
			errorMsg = fmt.Sprintf("failed to parse token: %v", err)
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": errorMsg,
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	claims, err := s.auth.ExtractClaims(ctx, *token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "failed to extract claims",
		})
	}

	c.Locals("user", *claims)

	return c.Next()
}
