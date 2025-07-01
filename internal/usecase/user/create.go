package user

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) CreateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("failed to hash password", u.log.Err(err))
		return nil, fmt.Errorf("CreateUser: %w", err)
	}

	user, err := u.repo.CreateUser(ctx, username, string(hashedBytes))
	if err != nil {
		u.log.Error("failed to create user", u.log.Err(err))
		return nil, fmt.Errorf("CreateUser: %w", err)
	}

	return user.ToEntity(), nil
}
