package user

import (
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type GetProfileResponse struct {
	Profile entity.User `json:"profile"`
}

func (s *Controller) GetProfile(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	user, err := s.user.GetUserByID(ctx, claims.UserID)
	if err != nil || user == nil {
		s.log.Error("GetUserByID err", s.log.Err(err), "user", user)
		return c.JSON(fiber.Map{
			"error": "User not found",
		})
	}

	response := GetProfileResponse{
		Profile: *user,
	}

	return c.JSON(response)
}
