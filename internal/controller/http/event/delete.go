package event

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

func (s *Controller) DeleteEvent(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	eventIDStr := c.Params("event_id")
	if eventIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Habit ID is required",
		})
	}

	habitID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit ID format",
		})
	}

	if err := s.events.DeleteEvent(ctx, claims.UserID, habitID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete event",
		})
	}

	c.Status(fiber.StatusOK)

	return nil
}
