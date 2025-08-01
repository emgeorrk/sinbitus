package habit

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
)

func (u *UseCase) CreateHabit(ctx context.Context, userID uint64, name, description string) (*entity.Habit, error) {
	habit, err := u.repo.CreateHabit(ctx, userID, name, description)
	if err != nil {
		u.log.Error("CreateHabit error", u.log.Err(err))
		return nil, fmt.Errorf("failed to create habit: %w", err)
	}

	return habit.ToEntity(), nil
}
