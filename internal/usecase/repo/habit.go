package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/model"
)

func (r *Repo) CreateHabit(ctx context.Context, habit model.Habit) (*model.Habit, error) {
	query, args, err := r.Builder.
		Insert("habits").
		Columns("user_id", "name", "description", "created_at").
		Values(habit.UserID, habit.Name, habit.Description, "NOW()").
		Suffix("RETURNING id, user_id, name, description, created_at").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("create habit: build query: %w", err))
	}

	var res model.Habit
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&res.ID,
		&res.UserID,
		&res.Name,
		&res.Description,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("create habit: scan: %w", err))
	}

	return &res, nil
}

func (r *Repo) GetHabitsByUserID(ctx context.Context, userID uint64) ([]model.Habit, error) {
	query, args, err := r.Builder.
		Select("h.id", "h.user_id", "h.name", "h.description", "h.created_at", "MAX(e.occurred_at) AS last_activity_at").
		From("habits AS h").
		LeftJoin("habit_events AS e ON e.habit_id = h.id").
		Where("h.user_id = ?", userID).
		GroupBy("h.id", "h.user_id", "h.name", "h.description", "h.created_at").
		OrderBy("h.created_at DESC").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get habits by user ID: build query: %w", err))
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get habits by user ID: query: %w", err))
	}
	defer rows.Close()

	var habits []model.Habit
	for rows.Next() {
		var habit model.Habit
		err := rows.Scan(&habit.ID, &habit.UserID, &habit.Name, &habit.Description, &habit.CreatedAt, &habit.LastActionAt)
		if err != nil {
			return nil, constants.WrapAsErrDB(fmt.Errorf("get habits by user ID: scan: %w", err))
		}
		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get habits by user ID: rows error: %w", err))
	}

	return habits, nil
}

func (r *Repo) GetHabitByID(ctx context.Context, habitID uint64) (*model.Habit, error) {
	query, args, err := r.Builder.
		Select("h.id", "h.user_id", "h.name", "h.description", "h.created_at", "MAX(e.occurred_at) last_activity_at").
		From("habits h").
		LeftJoin("habit_events e ON e.habit_id = h.id").
		Where(sq.Eq{"h.id": habitID}).
		GroupBy("h.id", "h.user_id", "h.name", "h.description", "h.created_at").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get habit by ID: build query: %w", err))
	}

	var habit model.Habit
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&habit.ID,
		&habit.UserID,
		&habit.Name,
		&habit.Description,
		&habit.CreatedAt,
		&habit.LastActionAt,
	)

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get habit by ID: scan: %w", err))
	}

	return &habit, nil
}

func (r *Repo) UpdateHabit(ctx context.Context, habit model.Habit) error {
	query, args, err := r.Builder.
		Update("habits").
		Set("name", habit.Name).
		Set("description", habit.Description).
		Where(sq.Eq{"id": habit.ID}).
		Suffix("RETURNING id, user_id, name, description, created_at").
		ToSql()

	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("update habit: build query: %w", err))
	}

	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&habit.ID,
		&habit.UserID,
		&habit.Name,
		&habit.Description,
		&habit.CreatedAt,
	)

	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("update habit: scan: %w", err))
	}

	return nil
}

func (r *Repo) DeleteHabit(ctx context.Context, habitID uint64) error {
	query, args, err := r.Builder.
		Delete("habits").
		Where(sq.Eq{"id": habitID}).
		ToSql()

	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("delete habit: build query: %w", err))
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("delete habit: exec: %w", err))
	}

	if result.RowsAffected() == 0 {
		return constants.ErrHabitNotFound
	}

	return nil
}
