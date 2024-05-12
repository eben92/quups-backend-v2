// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
)

const create = `-- name: Create :one
INSERT INTO users (
        email
    )
VALUES ($1)
RETURNING id, email, created_at, updated_at
`

func (q *Queries) Create(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, create, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findMany = `-- name: FindMany :many
SELECT id, email, created_at, updated_at FROM users
`

func (q *Queries) FindMany(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findMany)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
