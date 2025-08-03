package user

import (
	"github.com/gofiber/fiber/v3"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (s *Controller) SignUp(c fiber.Ctx) error {
	ctx := c.Context()

	var req SignupRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := s.users.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	token, err := s.auth.GenerateToken(ctx, user.ID, user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	resp := SignupResponse{
		Username: user.Username,
		Token:    token,
	}

	return c.JSON(resp)
}
