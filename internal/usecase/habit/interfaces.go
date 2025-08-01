package habit

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/model"
)

type Repository interface {
	CreateHabit(ctx context.Context, userID uint64, name, description string) (*model.Habit, error)
	UpdateHabit(ctx context.Context, habit *model.Habit) error
	DeleteHabit(ctx context.Context, habitID uint64) error
	GetHabitByID(ctx context.Context, habitID uint64) (*model.Habit, error)
	GetHabitsByUserID(ctx context.Context, userID uint64) ([]model.Habit, error)
}
