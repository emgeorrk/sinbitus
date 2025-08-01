package habit

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

func (u *UseCase) GetHabitsByUserID(ctx context.Context, userID uint64) ([]entity.Habit, error) {
	habits, err := u.repo.GetHabitsByUserID(ctx, userID)
	if err != nil {
		u.log.Error("GetHabitsByUserID error", u.log.Err(err))
		return nil, fmt.Errorf("get habits by user id %d: %w", userID, err)
	}

	var entities []entity.Habit
	for _, habit := range habits {
		entities = append(entities, *habit.ToEntity())
	}

	return entities, nil
}
