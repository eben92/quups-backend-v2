// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
        email,
        username,
        first_name,
        last_name,
        msisdn,
        full_name,
        image_url,
        gender,
        dob,
        otp,
        password
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
    )
RETURNING id, email, email_verified, msisdn, username, first_name, last_name, full_name, image_url, tin_number, gender, dob, otp, app_push_token, web_push_token, password, created_at, updated_at
`

type CreateUserParams struct {
	Email     string         `json:"email"`
	Username  sql.NullString `json:"username"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Msisdn    sql.NullString `json:"msisdn"`
	FullName  sql.NullString `json:"full_name"`
	ImageUrl  sql.NullString `json:"image_url"`
	Gender    sql.NullString `json:"gender"`
	Dob       sql.NullTime   `json:"dob"`
	Otp       sql.NullString `json:"otp"`
	Password  sql.NullString `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.Msisdn,
		arg.FullName,
		arg.ImageUrl,
		arg.Gender,
		arg.Dob,
		arg.Otp,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.Msisdn,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.FullName,
		&i.ImageUrl,
		&i.TinNumber,
		&i.Gender,
		&i.Dob,
		&i.Otp,
		&i.AppPushToken,
		&i.WebPushToken,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, email_verified, msisdn, username, first_name, last_name, full_name, image_url, tin_number, gender, dob, otp, app_push_token, web_push_token, password, created_at, updated_at FROM users
    WHERE email = $1
    LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.Msisdn,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.FullName,
		&i.ImageUrl,
		&i.TinNumber,
		&i.Gender,
		&i.Dob,
		&i.Otp,
		&i.AppPushToken,
		&i.WebPushToken,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, email_verified, msisdn, username, first_name, last_name, full_name, image_url, tin_number, gender, dob, otp, app_push_token, web_push_token, password, created_at, updated_at FROM users
    WHERE id = $1
    LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.Msisdn,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.FullName,
		&i.ImageUrl,
		&i.TinNumber,
		&i.Gender,
		&i.Dob,
		&i.Otp,
		&i.AppPushToken,
		&i.WebPushToken,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByMsisdn = `-- name: GetUserByMsisdn :one
SELECT id, email, email_verified, msisdn, username, first_name, last_name, full_name, image_url, tin_number, gender, dob, otp, app_push_token, web_push_token, password, created_at, updated_at FROM users
    WHERE msisdn = $1
    LIMIT 1
`

func (q *Queries) GetUserByMsisdn(ctx context.Context, msisdn sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByMsisdn, msisdn)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.Msisdn,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.FullName,
		&i.ImageUrl,
		&i.TinNumber,
		&i.Gender,
		&i.Dob,
		&i.Otp,
		&i.AppPushToken,
		&i.WebPushToken,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, email_verified, msisdn, username, first_name, last_name, full_name, image_url, tin_number, gender, dob, otp, app_push_token, web_push_token, password, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
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
			&i.EmailVerified,
			&i.Msisdn,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.FullName,
			&i.ImageUrl,
			&i.TinNumber,
			&i.Gender,
			&i.Dob,
			&i.Otp,
			&i.AppPushToken,
			&i.WebPushToken,
			&i.Password,
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
