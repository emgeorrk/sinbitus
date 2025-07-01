package http

import "github.com/gofiber/fiber/v3"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (s *Server) login(c fiber.Ctx) error {
	ctx := c.Context()

	var req LoginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := s.user.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid username or password",
		})
	}

	token, err := s.auth.GenerateToken(ctx, user.ID, user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	resp := LoginResponse{
		Token:    token,
		Username: user.Username,
	}

	return c.JSON(resp)
}
