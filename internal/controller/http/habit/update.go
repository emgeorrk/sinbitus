package habit

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type UpdateHabitRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type UpdateHabitResponse CreateHabitResponse

func (s *Controller) UpdateHabit(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	idStr := c.Params("habit_id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Habit ID is required",
		})
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit ID format",
		})
	}

	var req UpdateHabitRequest
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

	habit := entity.Habit{
		ID:          id,
		Name:        utils.SafeDeref(req.Name),
		Description: utils.SafeDeref(req.Description),
	}

	res, err := s.habits.UpdateHabit(ctx, claims.UserID, habit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update habit",
		})
	}

	response := UpdateHabitResponse{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
