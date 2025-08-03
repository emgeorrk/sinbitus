package event

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/emgeorrk/sinbitus/internal/model"
)

func (u *UseCase) UpdateEvent(ctx context.Context, userID uint64, event entity.Event) (*entity.Event, error) {
	ok, err := u.repo.IsEventOwnedByUser(ctx, event.ID, userID)
	if err != nil {
		u.log.Error("UpdateEvent error checking event ownership", u.log.Err(err))
		return nil, err
	} else if !ok {
		u.log.Warn("UpdateEvent unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("eventID", event.ID))
		return nil, constants.ErrEventNotFound
	}

	eventModel := model.Event{
		Description: event.Description,
	}

	updatedEvent, err := u.repo.UpdateEvent(ctx, eventModel)
	if err != nil {
		u.log.Error("UpdateEvent error", u.log.Err(err))
		return nil, err
	}

	return updatedEvent.ToEntity(), nil
}
