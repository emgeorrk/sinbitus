package event

import (
	"fmt"
	"strconv"
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type CreateEventRequest struct {
	Description string `json:"description" validate:"required,max=500"`
}

type CreateEventResponse struct {
	ID          uint64    `json:"id"`
	HabitID     uint64    `json:"habit_id"`
	Description string    `json:"description"`
	OccurredAt  time.Time `json:"occurred_at"`
}

func (s *Controller) CreateEvent(c fiber.Ctx, claims entity.UserClaims) error {
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

	var req CreateEventRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
	}

	event := entity.Event{
		HabitID:     habitID,
		Description: req.Description,
	}

	res, err := s.events.CreateEvent(ctx, claims.UserID, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create event",
		})
	}

	response := CreateEventResponse{
		ID:          res.ID,
		HabitID:     res.HabitID,
		Description: res.Description,
		OccurredAt:  res.OccurredAt,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
