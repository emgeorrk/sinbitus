package model

import (
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

type Event struct {
	ID          uint64
	HabitID     uint64
	Description string
	OccurredAt  time.Time
}

func (e *Event) ToEntity() *entity.Event {
	return &entity.Event{
		Description: e.Description,
		OccurredAt:  e.OccurredAt,
	}
}
