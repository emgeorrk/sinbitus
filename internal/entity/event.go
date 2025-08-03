package entity

import "time"

type Event struct {
	ID          uint64    `json:"id"`
	HabitID     uint64    `json:"habit_id"`
	Description string    `json:"description"`
	OccurredAt  time.Time `json:"occurred_at"`
}
