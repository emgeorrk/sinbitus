package habit

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
)

func (u *UseCase) DeleteHabit(ctx context.Context, userID, habitID uint64) error {
	habit, err := u.repo.GetHabitByID(ctx, habitID)
	if err != nil {
		u.log.Error("DeleteHabit error", u.log.Err(err))
		return err
	}

	if habit.UserID != userID {
		u.log.Warn("DeleteHabit unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("habitID", habitID))
		return constants.ErrHabitNotFound
	}

	if err := u.repo.DeleteHabit(ctx, habitID); err != nil {
		u.log.Error("DeleteHabit error", u.log.Err(err))
		return err
	}

	return nil
}
