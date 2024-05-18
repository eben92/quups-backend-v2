// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: company.sql

package repository

import (
	"context"
	"database/sql"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO companies (
        id,
        name,
        email,
        msisdn,
        about,
        image_url,
        banner_url,
        tin,
        brand_type,
        owner_id,
        is_active,
        currency_code,
        invitation_code,
        slug
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
    )
RETURNING id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, created_at, updated_at
`

type CreateCompanyParams struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Msisdn         string         `json:"msisdn"`
	About          sql.NullString `json:"about"`
	ImageUrl       sql.NullString `json:"image_url"`
	BannerUrl      sql.NullString `json:"banner_url"`
	Tin            sql.NullString `json:"tin"`
	BrandType      string         `json:"brand_type"`
	OwnerID        string         `json:"owner_id"`
	IsActive       bool           `json:"is_active"`
	CurrencyCode   string         `json:"currency_code"`
	InvitationCode sql.NullString `json:"invitation_code"`
	Slug           string         `json:"slug"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, createCompany,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Msisdn,
		arg.About,
		arg.ImageUrl,
		arg.BannerUrl,
		arg.Tin,
		arg.BrandType,
		arg.OwnerID,
		arg.IsActive,
		arg.CurrencyCode,
		arg.InvitationCode,
		arg.Slug,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.About,
		&i.Msisdn,
		&i.Email,
		&i.Tin,
		&i.ImageUrl,
		&i.BannerUrl,
		&i.BrandType,
		&i.OwnerID,
		&i.TotalSales,
		&i.IsActive,
		&i.CurrencyCode,
		&i.InvitationCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM companies WHERE id = $1
`

// WARNING: this will not work because of foreign key constraints
func (q *Queries) DeleteCompany(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, id)
	return err
}

const getCompanyByID = `-- name: GetCompanyByID :one
SELECT id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, created_at, updated_at FROM companies
    WHERE id = $1
    LIMIT 1
`

func (q *Queries) GetCompanyByID(ctx context.Context, id string) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByID, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.About,
		&i.Msisdn,
		&i.Email,
		&i.Tin,
		&i.ImageUrl,
		&i.BannerUrl,
		&i.BrandType,
		&i.OwnerID,
		&i.TotalSales,
		&i.IsActive,
		&i.CurrencyCode,
		&i.InvitationCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :one
UPDATE companies SET 
        name = $2, 
        email = $3,
        image_url = $4,
        about = $5,
        tin = $6,
        is_active = $7,
        msisdn = $8
    WHERE id = $1
RETURNING id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, created_at, updated_at
`

type UpdateCompanyParams struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	ImageUrl sql.NullString `json:"image_url"`
	About    sql.NullString `json:"about"`
	Tin      sql.NullString `json:"tin"`
	IsActive bool           `json:"is_active"`
	Msisdn   string         `json:"msisdn"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, updateCompany,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.ImageUrl,
		arg.About,
		arg.Tin,
		arg.IsActive,
		arg.Msisdn,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.About,
		&i.Msisdn,
		&i.Email,
		&i.Tin,
		&i.ImageUrl,
		&i.BannerUrl,
		&i.BrandType,
		&i.OwnerID,
		&i.TotalSales,
		&i.IsActive,
		&i.CurrencyCode,
		&i.InvitationCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
