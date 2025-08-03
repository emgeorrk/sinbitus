package event

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type GetEventsRequest struct {
	Limit   int    `json:"limit" validate:"gte=1,lte=100"`
	Offset  int    `json:"offset" validate:"gte=0,lte=100"`
	OrderBy string `json:"order_by" validate:"required,oneof=id name created_at"`
}

type GetEventsResponse struct {
	Events []entity.Event `json:"events"`
}

func (s *Controller) GetEvents(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	habitIDStr := c.Params("habit_id")
	if habitIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Habit ID is required",
		})
	}

	habitID, err := strconv.ParseUint(habitIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit ID format",
		})
	}

	var req GetEventsRequest
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

	events, err := s.events.GetEventsByHabitID(ctx, claims.UserID, habitID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get events",
		})
	}

	response := GetEventsResponse{
		Events: events,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
