package habit

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/entity"
)

func (u *UseCase) UpdateHabit(ctx context.Context, userID, habitID uint64, name, description *string) (*entity.Habit, error) {
	habit, err := u.repo.GetHabitByID(ctx, habitID)
	if err != nil {
		u.log.Error("UpdateHabit error", u.log.Err(err))
		return nil, err
	}

	if habit.UserID != userID {
		u.log.Warn("UpdateHabit unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("habitID", habitID))
		return nil, constants.ErrHabitNotFound
	}

	if name != nil {
		habit.Name = *name
	}

	if description != nil {
		habit.Description = *description
	}

	if err := u.repo.UpdateHabit(ctx, habit); err != nil {
		u.log.Error("UpdateHabit error", u.log.Err(err))
		return nil, err
	}

	return habit.ToEntity(), nil
}
