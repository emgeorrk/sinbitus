package habit

import (
	"fmt"
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/gofiber/fiber/v3"
)

type CreateHabitRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=500"`
}

type CreateHabitResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *Controller) CreateHabit(c fiber.Ctx, claims entity.UserClaims) error {
	ctx := c.Context()

	var req CreateHabitRequest
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

	habit := entity.Habit{
		Name:        req.Name,
		Description: req.Description,
	}

	res, err := s.habit.CreateHabit(ctx, claims.UserID, habit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create habit",
		})
	}

	response := CreateHabitResponse{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
