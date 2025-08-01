package habit

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/model"
)

func (u *UseCase) CreateHabit(ctx context.Context, userID uint64, habit entity.Habit) (*entity.Habit, error) {
	m := model.Habit{
		UserID:      userID,
		Name:        habit.Name,
		Description: habit.Description,
	}

	res, err := u.repo.CreateHabit(ctx, m)
	if err != nil {
		u.log.Error("CreateHabit error", u.log.Err(err))
		return nil, fmt.Errorf("failed to create habit: %w", err)
	}

	return res.ToEntity(), nil
}
