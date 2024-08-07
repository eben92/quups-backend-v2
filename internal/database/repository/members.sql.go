// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: members.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const addMember = `-- name: AddMember :one
INSERT INTO members (
        company_id,
        name,
        email,
        msisdn,
        role,
        user_id,
        status
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7
    )
RETURNING id, name, msisdn, email, role, status, company_id, user_id, created_at, updated_at
`

type AddMemberParams struct {
	CompanyID string         `json:"company_id"`
	Name      string         `json:"name"`
	Email     sql.NullString `json:"email"`
	Msisdn    string         `json:"msisdn"`
	Role      string         `json:"role"`
	UserID    sql.NullString `json:"user_id"`
	Status    string         `json:"status"`
}

func (q *Queries) AddMember(ctx context.Context, arg AddMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, addMember,
		arg.CompanyID,
		arg.Name,
		arg.Email,
		arg.Msisdn,
		arg.Role,
		arg.UserID,
		arg.Status,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Msisdn,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.CompanyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMembersByCompanyID = `-- name: GetMembersByCompanyID :many
SELECT id, name, msisdn, email, role, status, company_id, user_id, created_at, updated_at FROM members
 WHERE company_id = $1
 ORDER BY created_at desc
 LIMIT $2 OFFSET $3
`

type GetMembersByCompanyIDParams struct {
	CompanyID string `json:"company_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) GetMembersByCompanyID(ctx context.Context, arg GetMembersByCompanyIDParams) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, getMembersByCompanyID, arg.CompanyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Member{}
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Msisdn,
			&i.Email,
			&i.Role,
			&i.Status,
			&i.CompanyID,
			&i.UserID,
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

const getUserTeam = `-- name: GetUserTeam :one
SELECT members.id, members.name, members.msisdn, members.email, members.role, members.status, members.company_id, members.user_id, members.created_at, members.updated_at,  
    companies.email as company_email,
    companies.name as company_name,
    companies.slug as company_slug,
    companies.banner_url as company_banner_url,
    companies.image_url as company_image_url,
    companies.about as company_about,
    companies.is_active as company_is_active,
    companies.has_onboarded as company_has_onboarded,
    companies.msisdn as company_msisdn
FROM members
JOIN companies ON members.company_id = companies.id
WHERE members.company_id = $1 AND members.user_id = $2
`

type GetUserTeamParams struct {
	CompanyID string         `json:"company_id"`
	UserID    sql.NullString `json:"user_id"`
}

type GetUserTeamRow struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Msisdn              string         `json:"msisdn"`
	Email               sql.NullString `json:"email"`
	Role                string         `json:"role"`
	Status              string         `json:"status"`
	CompanyID           string         `json:"company_id"`
	UserID              sql.NullString `json:"user_id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	CompanyEmail        string         `json:"company_email"`
	CompanyName         string         `json:"company_name"`
	CompanySlug         string         `json:"company_slug"`
	CompanyBannerUrl    sql.NullString `json:"company_banner_url"`
	CompanyImageUrl     sql.NullString `json:"company_image_url"`
	CompanyAbout        sql.NullString `json:"company_about"`
	CompanyIsActive     bool           `json:"company_is_active"`
	CompanyHasOnboarded bool           `json:"company_has_onboarded"`
	CompanyMsisdn       string         `json:"company_msisdn"`
}

func (q *Queries) GetUserTeam(ctx context.Context, arg GetUserTeamParams) (GetUserTeamRow, error) {
	row := q.db.QueryRowContext(ctx, getUserTeam, arg.CompanyID, arg.UserID)
	var i GetUserTeamRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Msisdn,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.CompanyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CompanyEmail,
		&i.CompanyName,
		&i.CompanySlug,
		&i.CompanyBannerUrl,
		&i.CompanyImageUrl,
		&i.CompanyAbout,
		&i.CompanyIsActive,
		&i.CompanyHasOnboarded,
		&i.CompanyMsisdn,
	)
	return i, err
}

const getUserTeams = `-- name: GetUserTeams :many
SELECT members.id, members.name, members.msisdn, members.email, members.role, members.status, members.company_id, members.user_id, members.created_at, members.updated_at,  
    companies.email as company_email,
    companies.name as company_name,
    companies.slug as company_slug,
    companies.banner_url as company_banner_url,
    companies.image_url as company_image_url,
    companies.about as company_about,
    companies.is_active as company_is_active,
    companies.has_onboarded as company_has_onboarded,
    companies.msisdn as company_msisdn
FROM members
JOIN companies ON members.company_id = companies.id
WHERE members.user_id = $1
`

type GetUserTeamsRow struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Msisdn              string         `json:"msisdn"`
	Email               sql.NullString `json:"email"`
	Role                string         `json:"role"`
	Status              string         `json:"status"`
	CompanyID           string         `json:"company_id"`
	UserID              sql.NullString `json:"user_id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	CompanyEmail        string         `json:"company_email"`
	CompanyName         string         `json:"company_name"`
	CompanySlug         string         `json:"company_slug"`
	CompanyBannerUrl    sql.NullString `json:"company_banner_url"`
	CompanyImageUrl     sql.NullString `json:"company_image_url"`
	CompanyAbout        sql.NullString `json:"company_about"`
	CompanyIsActive     bool           `json:"company_is_active"`
	CompanyHasOnboarded bool           `json:"company_has_onboarded"`
	CompanyMsisdn       string         `json:"company_msisdn"`
}

func (q *Queries) GetUserTeams(ctx context.Context, userID sql.NullString) ([]GetUserTeamsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTeams, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserTeamsRow{}
	for rows.Next() {
		var i GetUserTeamsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Msisdn,
			&i.Email,
			&i.Role,
			&i.Status,
			&i.CompanyID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CompanyEmail,
			&i.CompanyName,
			&i.CompanySlug,
			&i.CompanyBannerUrl,
			&i.CompanyImageUrl,
			&i.CompanyAbout,
			&i.CompanyIsActive,
			&i.CompanyHasOnboarded,
			&i.CompanyMsisdn,
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

const updateMember = `-- name: UpdateMember :one

UPDATE members SET 
        name = $2, 
        email = $3,
        role = $4,
        status = $5
    WHERE id = $1
RETURNING id, name, msisdn, email, role, status, company_id, user_id, created_at, updated_at
`

type UpdateMemberParams struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Email  sql.NullString `json:"email"`
	Role   string         `json:"role"`
	Status string         `json:"status"`
}

// SELECT members.*, companies.*
// FROM members
// JOIN companies ON members.company_id = companies.id
// WHERE members.user_id = sqlc.arg(user_id);
func (q *Queries) UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMember,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.Status,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Msisdn,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.CompanyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateMemberInvitationStatus = `-- name: UpdateMemberInvitationStatus :one
UPDATE members SET 
        status = $2
    WHERE id = $1
RETURNING id, name, msisdn, email, role, status, company_id, user_id, created_at, updated_at
`

type UpdateMemberInvitationStatusParams struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateMemberInvitationStatus(ctx context.Context, arg UpdateMemberInvitationStatusParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMemberInvitationStatus, arg.ID, arg.Status)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Msisdn,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.CompanyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
