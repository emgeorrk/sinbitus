package event

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/utils"
)

func (u *UseCase) GetEventsByHabitID(ctx context.Context, userID, habitID uint64) ([]entity.Event, error) {
	ok, err := u.repo.IsHabitOwnedByUser(ctx, habitID, userID)
	if err != nil {
		u.log.Error("GetEventsByHabitID error checking habit ownership", u.log.Err(err))
		return nil, err
	} else if !ok {
		u.log.Warn("GetEventsByHabitID unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("habitID", habitID))
		return nil, constants.ErrHabitNotFound
	}

	events, err := u.repo.GetEventsByHabitID(ctx, habitID)
	if err != nil {
		u.log.Error("GetEventsByHabitID error", u.log.Err(err))
		return nil, err
	}

	var result []entity.Event
	for _, event := range events {
		result = append(result, utils.SafeDeref(event.ToEntity()))
	}

	return result, nil
}
