package user

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) Authenticate(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		u.log.Error("failed to get user by username", u.log.Err(err))
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		u.log.Error("failed to compare password", u.log.Err(err))
		return nil, fmt.Errorf("invalid password")
	}

	return user.ToEntity(), nil
}
