package habit

import (
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type GetHabitsRequest struct {
	Limit   int    `json:"limit" validate:"gte=1,lte=100"`
	Offset  int    `json:"offset" validate:"gte=0,lte=100"`
	OrderBy string `json:"order_by" validate:"required,oneof=id name created_at"`
}

type GetHabitsResponse struct {
	Habits []entity.Habit `json:"habits"`
}

func (s *Controller) GetHabits(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	var req GetHabitsRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	habits, err := s.habit.GetHabitsByUserID(ctx, claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get habits",
		})
	}

	response := GetHabitsResponse{
		Habits: habits,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
