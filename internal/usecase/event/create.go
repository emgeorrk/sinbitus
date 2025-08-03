package event

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/model"
)

func (u *UseCase) CreateEvent(ctx context.Context, userID uint64, event entity.Event) (*entity.Event, error) {
	ok, err := u.repo.IsHabitOwnedByUser(ctx, event.HabitID, userID)
	if err != nil {
		u.log.Error("CreateEvent error checking habit ownership", u.log.Err(err))
		return nil, err
	} else if !ok {
		return nil, constants.ErrHabitNotFound
	}

	eventModel := model.Event{
		HabitID:     event.HabitID,
		Description: event.Description,
	}

	createdEvent, err := u.repo.CreateEvent(ctx, eventModel)
	if err != nil {
		u.log.Error("CreateEvent error", u.log.Err(err))
		return nil, err
	}

	return createdEvent.ToEntity(), nil
}
