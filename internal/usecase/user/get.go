package user

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

func (u *UseCase) GetUserByID(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}

	return user.ToEntity(), nil
}
