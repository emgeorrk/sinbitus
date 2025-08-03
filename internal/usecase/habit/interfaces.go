package habit

import (
	"context"
	"time"

	"github.com/emgeorrk/sinbitus/internal/model"
)

type (
	Repository interface {
		CreateHabit(ctx context.Context, habit model.Habit) (*model.Habit, error)
		UpdateHabit(ctx context.Context, habit model.Habit) error
		DeleteHabit(ctx context.Context, habitID uint64) error
		GetHabitByID(ctx context.Context, habitID uint64) (*model.Habit, error)
		GetHabitsByUserID(ctx context.Context, userID uint64) ([]model.Habit, error)
	}

	TimeProvider interface {
		Now() time.Time
	}
)
