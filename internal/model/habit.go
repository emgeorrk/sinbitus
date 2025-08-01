package model

import (
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

type Habit struct {
	ID          uint64
	UserID      uint64
	Name        string
	Description string
	CreatedAt   time.Time

	LastActionAt *time.Time
}

func (h *Habit) ToEntity() *entity.Habit {
	return &entity.Habit{
		ID:           h.ID,
		Name:         h.Name,
		Description:  h.Description,
		CreatedAt:    h.CreatedAt,
		LastActionAt: h.LastActionAt,
	}
}
