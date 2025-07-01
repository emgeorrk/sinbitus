package user

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/dto"
)

type RepoProvider interface {
	GetUserByID(ctx context.Context, id uint64) (*dto.User, error)
	GetUserByUsername(ctx context.Context, username string) (*dto.User, error)
	CreateUser(ctx context.Context, username, passwordHash string) (*dto.User, error)
}
