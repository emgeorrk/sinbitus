package habit

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

	HabitUseCase interface {
		CreateHabit(ctx context.Context, userID uint64, habit entity.Habit) (*entity.Habit, error)
		UpdateHabit(ctx context.Context, userID uint64, habit entity.Habit) (*entity.Habit, error)
		DeleteHabit(ctx context.Context, userID, habitID uint64) error
		GetHabitsByUserID(ctx context.Context, userID uint64) ([]entity.Habit, error)
	}

	TimeProvider interface {
		Now() time.Time
	}
)
