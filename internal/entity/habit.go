package entity

import "time"

type Habit struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`

	LastActionAt *time.Time `json:"last_action_at"`
	Streak       uint64     `json:"streak"`
}
