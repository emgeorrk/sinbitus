package model

import (
	"time"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

type User struct {
	ID           uint64
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Username: u.Username,
	}
}
