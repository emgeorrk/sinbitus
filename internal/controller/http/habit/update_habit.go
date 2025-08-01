package habit

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type UpdateHabitRequest struct {
	Name        *string `json:"name,omitempty" validate:",min=1,max=100"`
	Description *string `json:"description,omitempty" validate:"max=500"`
}

type UpdateHabitResponse CreateHabitResponse

func (s *Controller) UpdateHabit(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	idStr := c.Params("id")
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
		s.log.Error("failed to bind create habit request", s.log.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := s.validator.Struct(req); err != nil {
		s.log.Error("validation failed for create habit request", s.log.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	habit, err := s.habit.UpdateHabit(ctx, claims.UserID, id, req.Name, req.Description)
	if err != nil {
		s.log.Error("failed to update habit", s.log.Err(err), "user_id", claims.UserID, "habit_id", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update habit",
		})
	}

	response := UpdateHabitResponse{
		ID:          habit.ID,
		Name:        habit.Name,
		Description: habit.Description,
		CreatedAt:   habit.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
