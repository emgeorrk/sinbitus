package event

import (
	"strconv"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type UpdateEventRequest struct {
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type UpdateEventResponse CreateEventResponse

func (s *Controller) UpdateEvent(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	eventIDStr := c.Params("event_id")
	if eventIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event ID is required",
		})
	}

	eventID, err := strconv.ParseUint(eventIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid event ID format",
		})
	}

	var req UpdateEventRequest
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

	event := entity.Event{
		ID:          eventID,
		Description: utils.SafeDeref(req.Description),
	}

	res, err := s.events.UpdateEvent(ctx, claims.UserID, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update event",
		})
	}

	response := UpdateEventResponse{
		ID:          res.ID,
		HabitID:     res.HabitID,
		Description: res.Description,
		OccurredAt:  res.OccurredAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
