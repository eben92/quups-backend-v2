// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company.sql

package repository

import (
	"context"
	"database/sql"
)

const addAddress = `-- name: AddAddress :one
INSERT INTO addresses (
        company_id,
        user_id,
        msisdn,
        is_default,
        latitude,
        longitude,
        description,
        formatted_address,
        country_code,
        region,
        street,
        city,
        country,
        postal_code
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
    )
RETURNING id, street, city, region, country, country_code, formatted_address, description, postal_code, latitude, longitude, msisdn, is_default, user_id, company_id, created_at, updated_at
`

type AddAddressParams struct {
	CompanyID        sql.NullString  `json:"company_id"`
	UserID           sql.NullString  `json:"user_id"`
	Msisdn           sql.NullString  `json:"msisdn"`
	IsDefault        bool            `json:"is_default"`
	Latitude         sql.NullFloat64 `json:"latitude"`
	Longitude        sql.NullFloat64 `json:"longitude"`
	Description      sql.NullString  `json:"description"`
	FormattedAddress sql.NullString  `json:"formatted_address"`
	CountryCode      string          `json:"country_code"`
	Region           string          `json:"region"`
	Street           string          `json:"street"`
	City             string          `json:"city"`
	Country          string          `json:"country"`
	PostalCode       sql.NullString  `json:"postal_code"`
}

func (q *Queries) AddAddress(ctx context.Context, arg AddAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, addAddress,
		arg.CompanyID,
		arg.UserID,
		arg.Msisdn,
		arg.IsDefault,
		arg.Latitude,
		arg.Longitude,
		arg.Description,
		arg.FormattedAddress,
		arg.CountryCode,
		arg.Region,
		arg.Street,
		arg.City,
		arg.Country,
		arg.PostalCode,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Street,
		&i.City,
		&i.Region,
		&i.Country,
		&i.CountryCode,
		&i.FormattedAddress,
		&i.Description,
		&i.PostalCode,
		&i.Latitude,
		&i.Longitude,
		&i.Msisdn,
		&i.IsDefault,
		&i.UserID,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

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
RETURNING id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, has_onboarded, is_deleted, created_at, updated_at
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
		&i.HasOnboarded,
		&i.IsDeleted,
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

const getAllCompanies = `-- name: GetAllCompanies :many
SELECT id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, has_onboarded, is_deleted, created_at, updated_at 
    FROM companies
    LIMIT 10
`

func (q *Queries) GetAllCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getAllCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Company{}
	for rows.Next() {
		var i Company
		if err := rows.Scan(
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
			&i.HasOnboarded,
			&i.IsDeleted,
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

const getCompanyByID = `-- name: GetCompanyByID :one
SELECT id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, has_onboarded, is_deleted, created_at, updated_at FROM companies
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
		&i.HasOnboarded,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCompanyByName = `-- name: GetCompanyByName :one
SELECT id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, has_onboarded, is_deleted, created_at, updated_at FROM companies
    WHERE name = $1
    LIMIT 1
`

func (q *Queries) GetCompanyByName(ctx context.Context, name string) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByName, name)
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
		&i.HasOnboarded,
		&i.IsDeleted,
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
        msisdn = $8,
        has_onboarded = $9
    WHERE id = $1
RETURNING id, name, slug, about, msisdn, email, tin, image_url, banner_url, brand_type, owner_id, total_sales, is_active, currency_code, invitation_code, has_onboarded, is_deleted, created_at, updated_at
`

type UpdateCompanyParams struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	ImageUrl     sql.NullString `json:"image_url"`
	About        sql.NullString `json:"about"`
	Tin          sql.NullString `json:"tin"`
	IsActive     bool           `json:"is_active"`
	Msisdn       string         `json:"msisdn"`
	HasOnboarded bool           `json:"has_onboarded"`
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
		arg.HasOnboarded,
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
		&i.HasOnboarded,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
