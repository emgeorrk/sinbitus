package repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/emgeorrk/sinbitus/internal/constants"
	"github.com/emgeorrk/sinbitus/internal/model"
)

func (r *Repo) CreateEvent(ctx context.Context, event model.Event) (*model.Event, error) {
	query, args, err := r.Builder.Insert("habit_events").
		Columns("habit_id", "description", "occured_at").
		Values(event.HabitID, event.Description, "NOW()").
		Suffix("RETURNING id, habit_id, description, occured_at").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("create event: build query: %w", err))
	}

	var res model.Event
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&res.ID,
		&res.HabitID,
		&res.Description,
		&res.OccurredAt,
	)

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("create event: scan: %w", err))
	}

	return &res, nil
}

func (r *Repo) GetEventsByHabitID(ctx context.Context, habitID uint64) ([]model.Event, error) {
	query, args, err := r.Builder.
		Select("id", "habit_id", "description", "occurred_at").
		From("habit_events").
		Where(sq.Eq{"habit_id": habitID}).
		OrderBy("occurred_at DESC").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get events by habit ID: build query: %w", err))
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get events by habit ID: query: %w", err))
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.ID, &event.HabitID, &event.Description, &event.OccurredAt); err != nil {
			return nil, constants.WrapAsErrDB(fmt.Errorf("get events by habit ID: scan: %w", err))
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("get events by habit ID: rows error: %w", err))
	}

	return events, nil
}

func (r *Repo) UpdateEvent(ctx context.Context, event model.Event) (*model.Event, error) {
	values := map[string]interface{}{
		"description": event.Description,
	}

	query, args, err := r.Builder.
		Update("habit_events").
		SetMap(values).
		Where(sq.Eq{"id": event.ID}).
		Suffix("RETURNING id, habit_id, description, occurred_at").
		ToSql()

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("update event: build query: %w", err))
	}

	var res model.Event
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&res.ID,
		&res.HabitID,
		&res.Description,
		&res.OccurredAt,
	)

	if err != nil {
		return nil, constants.WrapAsErrDB(fmt.Errorf("update event: scan: %w", err))
	}

	return &res, nil
}

func (r *Repo) DeleteEvent(ctx context.Context, eventID uint64) error {
	query, args, err := r.Builder.
		Delete("habit_events").
		Where(sq.Eq{"id": eventID}).
		ToSql()

	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("delete event: build query: %w", err))
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return constants.WrapAsErrDB(fmt.Errorf("delete event: exec: %w", err))
	}

	if result.RowsAffected() == 0 {
		return constants.ErrEventNotFound
	}

	return nil
}

func (r *Repo) IsHabitOwnedByUser(ctx context.Context, habitID, userID uint64) (bool, error) {
	query, args, err := r.Builder.
		Select("COUNT(*)").
		From("habits").
		Where(sq.Eq{"id": habitID, "user_id": userID}).
		ToSql()

	if err != nil {
		return false, constants.WrapAsErrDB(fmt.Errorf("is habit owned by user: build query: %w", err))
	}

	var count int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, constants.WrapAsErrDB(fmt.Errorf("is habit owned by user: scan: %w", err))
	}

	return count > 0, nil
}

func (r *Repo) IsEventOwnedByUser(ctx context.Context, eventID, userID uint64) (bool, error) {
	query, args, err := r.Builder.
		Select("COUNT(*)").
		From("habit_events he").
		Join("habits h ON he.habit_id = h.id").
		Where(sq.Eq{"he.id": eventID, "h.user_id": userID}).
		ToSql()

	if err != nil {
		return false, constants.WrapAsErrDB(fmt.Errorf("is event owned by user: build query: %w", err))
	}

	var count int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, constants.WrapAsErrDB(fmt.Errorf("is event owned by user: scan: %w", err))
	}

	return count > 0, nil
}
