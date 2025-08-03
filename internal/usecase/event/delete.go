package event

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/constants"
)

func (u *UseCase) DeleteEvent(ctx context.Context, userID, eventID uint64) error {
	ok, err := u.repo.IsEventOwnedByUser(ctx, eventID, userID)
	if err != nil {
		u.log.Error("DeleteEvent error checking event ownership", u.log.Err(err))
		return err
	} else if !ok {
		u.log.Warn("DeleteEvent unauthorized access attempt", u.log.Uint64("userID", userID), u.log.Uint64("eventID", eventID))
		return constants.ErrEventNotFound
	}

	if err := u.repo.DeleteEvent(ctx, eventID); err != nil {
		u.log.Error("DeleteEvent error", u.log.Err(err))
		return err
	}

	return nil
}
