package habit

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/utils"
)

func (u *UseCase) UpdateHabit(ctx context.Context, userID uint64, habit entity.Habit) (*entity.Habit, error) {
	res, err := u.repo.GetHabitByID(ctx, habit.ID)
	if err != nil {
		u.log.Error("UpdateHabit error", u.log.Err(err))
		return nil, err
	}

	if res.UserID != userID {
		u.log.Warn("UpdateHabit unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("habitID", habit.ID))
		return nil, constants.ErrHabitNotFound
	}

	{
		res.Name = utils.FirstNonZero(habit.Name, res.Name)
		res.Description = utils.FirstNonZero(habit.Description, res.Description)
	}

	if err := u.repo.UpdateHabit(ctx, *res); err != nil {
		u.log.Error("UpdateHabit error", u.log.Err(err))
		return nil, err
	}

	return res.ToEntity(), nil
}
