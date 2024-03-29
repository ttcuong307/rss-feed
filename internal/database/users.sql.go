// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES (?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ApiKey    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ApiKey,
	)
	return err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
SELECT id, name, created_at, updated_at, api_key FROM users WHERE api_key = ? LIMIT 1
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApiKey,
	)
	return i, err
}
