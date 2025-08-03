package event

import (
	"context"
	"time"

	"github.com/emgeorrk/sinbitus/internal/model"
)

type (
	Repository interface {
		CreateEvent(ctx context.Context, event model.Event) (*model.Event, error)
		GetEventsByHabitID(ctx context.Context, habitID uint64) ([]model.Event, error)
		UpdateEvent(ctx context.Context, event model.Event) (*model.Event, error)
		DeleteEvent(ctx context.Context, eventID uint64) error

		IsHabitOwnedByUser(ctx context.Context, habitID, userID uint64) (bool, error)
		IsEventOwnedByUser(ctx context.Context, eventID, userID uint64) (bool, error)
	}

	TimeProvider interface {
		Now() time.Time
	}
)
