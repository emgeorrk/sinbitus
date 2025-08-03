package event

import (
	"context"
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthUseCase interface {
		GenerateToken(ctx context.Context, userID uint64, username string) (string, error)
		ParseToken(ctx context.Context, tokenStr string) (*jwt.Token, error)
		ExtractClaims(ctx context.Context, token jwt.Token) (*entity.UserClaims, error)
	}

	UserUseCase interface {
		GetUserByID(ctx context.Context, id uint64) (*entity.User, error)
		CreateUser(ctx context.Context, username, password string) (*entity.User, error)
		Authenticate(ctx context.Context, username, password string) (*entity.User, error)
	}

	EventsUseCase interface {
		CreateEvent(ctx context.Context, userID uint64, event entity.Event) (*entity.Event, error)
		GetEventsByHabitID(ctx context.Context, userID, habitID uint64) ([]entity.Event, error)
		UpdateEvent(ctx context.Context, userID uint64, event entity.Event) (*entity.Event, error)
		DeleteEvent(ctx context.Context, userID, eventID uint64) error
	}

	TimeProvider interface {
		Now() time.Time
	}
)
