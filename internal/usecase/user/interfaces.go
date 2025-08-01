package user

import (
	"context"
	"time"

	"github.com/emgeorrk/sinbitus/internal/model"
)

type (
	Repository interface {
		GetUserByID(ctx context.Context, id uint64) (*model.User, error)
		GetUserByUsername(ctx context.Context, username string) (*model.User, error)
		CreateUser(ctx context.Context, username, passwordHash string) (*model.User, error)
	}

	TimeProvider interface {
		Now() time.Time
	}
)
