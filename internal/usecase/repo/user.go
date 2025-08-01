package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/model"
	"github.com/jackc/pgx/v4"
)

func (r *Repo) GetUserByID(ctx context.Context, id uint64) (*model.User, error) {
	sql, args, err := r.Builder.
		Select("id", "username", "password_hash", "created_at").
		From("users").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("get user by id %d: build sql: %w", id, err)
	}

	var user model.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with id %d not found: %w", id, err)
		}
		return nil, fmt.Errorf("get user by id %d: scan: %w", id, err)
	}

	return &user, nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	sql, args, err := r.Builder.
		Select("id", "username", "password_hash", "created_at").
		From("users").
		Where("username = ?", username).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("get user by username %s: build sql: %w", username, err)
	}

	var user model.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with username %s not found: %w", username, err)
		}
		return nil, fmt.Errorf("get user by username %s: scan: %w", username, err)
	}

	return &user, nil
}

func (r *Repo) CreateUser(ctx context.Context, username string, passwordHash string) (*model.User, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username", "password_hash", "created_at").
		Values(username, passwordHash, "NOW()").
		Suffix("RETURNING id, username, password_hash, created_at").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("create user: build sql: %w", err)
	}

	var user model.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("create user: scan: %w", err)
	}

	return &user, nil
}
