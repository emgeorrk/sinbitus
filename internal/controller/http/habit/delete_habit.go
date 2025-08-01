package habit

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

func (s *Controller) DeleteHabit(c fiber.Ctx, claims entity.UserClaims) error {
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

	err = s.habit.DeleteHabit(ctx, claims.UserID, id)
	if err != nil {
		s.log.Error("failed to delete habit", s.log.Err(err), "user_id", claims.UserID, "habit_id", id)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete habit",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
