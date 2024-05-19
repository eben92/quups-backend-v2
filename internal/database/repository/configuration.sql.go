// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: configuration.sql

package repository

import (
	"context"
	"time"
)

const addWorkingHour = `-- name: AddWorkingHour :one

INSERT INTO working_hours (
        company_id,
        day,
        opens_at,
        closes_at
    )
VALUES (
    $1, $2, $3, $4
    )
RETURNING id, day, opens_at, closes_at, company_id, created_at, updated_at
`

type AddWorkingHourParams struct {
	CompanyID string    `json:"company_id"`
	Day       string    `json:"day"`
	OpensAt   time.Time `json:"opens_at"`
	ClosesAt  time.Time `json:"closes_at"`
}

// WORKING HOURS
func (q *Queries) AddWorkingHour(ctx context.Context, arg AddWorkingHourParams) (WorkingHour, error) {
	row := q.db.QueryRowContext(ctx, addWorkingHour,
		arg.CompanyID,
		arg.Day,
		arg.OpensAt,
		arg.ClosesAt,
	)
	var i WorkingHour
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.OpensAt,
		&i.ClosesAt,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createConfiguration = `-- name: CreateConfiguration :one
INSERT INTO configurations (
        company_id,
        pickup,
        cash_on_delivery,
        digital_payments,
        delivery
    )
VALUES (
    $1, $2, $3, $4, $5
    )
RETURNING id, delivery, pickup, cash_on_delivery, digital_payments, company_id, created_at, updated_at
`

type CreateConfigurationParams struct {
	CompanyID       string `json:"company_id"`
	Pickup          bool   `json:"pickup"`
	CashOnDelivery  bool   `json:"cash_on_delivery"`
	DigitalPayments bool   `json:"digital_payments"`
	Delivery        bool   `json:"delivery"`
}

func (q *Queries) CreateConfiguration(ctx context.Context, arg CreateConfigurationParams) (Configuration, error) {
	row := q.db.QueryRowContext(ctx, createConfiguration,
		arg.CompanyID,
		arg.Pickup,
		arg.CashOnDelivery,
		arg.DigitalPayments,
		arg.Delivery,
	)
	var i Configuration
	err := row.Scan(
		&i.ID,
		&i.Delivery,
		&i.Pickup,
		&i.CashOnDelivery,
		&i.DigitalPayments,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getConfigurationByCompanyID = `-- name: GetConfigurationByCompanyID :one
SELECT id, delivery, pickup, cash_on_delivery, digital_payments, company_id, created_at, updated_at FROM configurations
 WHERE company_id = $1
 LIMIT 1
`

func (q *Queries) GetConfigurationByCompanyID(ctx context.Context, companyID string) (Configuration, error) {
	row := q.db.QueryRowContext(ctx, getConfigurationByCompanyID, companyID)
	var i Configuration
	err := row.Scan(
		&i.ID,
		&i.Delivery,
		&i.Pickup,
		&i.CashOnDelivery,
		&i.DigitalPayments,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getWorkingHoursByCompanyID = `-- name: GetWorkingHoursByCompanyID :many
SELECT id, day, opens_at, closes_at, company_id, created_at, updated_at FROM working_hours
 WHERE company_id = $1
 LIMIT 7
`

func (q *Queries) GetWorkingHoursByCompanyID(ctx context.Context, companyID string) ([]WorkingHour, error) {
	rows, err := q.db.QueryContext(ctx, getWorkingHoursByCompanyID, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WorkingHour{}
	for rows.Next() {
		var i WorkingHour
		if err := rows.Scan(
			&i.ID,
			&i.Day,
			&i.OpensAt,
			&i.ClosesAt,
			&i.CompanyID,
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

const updateConfiguationByCompanyID = `-- name: UpdateConfiguationByCompanyID :one
UPDATE configurations SET 
        pickup = $2,
        cash_on_delivery = $3,
        digital_payments = $4,
        delivery = $5
    WHERE company_id = $1
RETURNING id, delivery, pickup, cash_on_delivery, digital_payments, company_id, created_at, updated_at
`

type UpdateConfiguationByCompanyIDParams struct {
	CompanyID       string `json:"company_id"`
	Pickup          bool   `json:"pickup"`
	CashOnDelivery  bool   `json:"cash_on_delivery"`
	DigitalPayments bool   `json:"digital_payments"`
	Delivery        bool   `json:"delivery"`
}

func (q *Queries) UpdateConfiguationByCompanyID(ctx context.Context, arg UpdateConfiguationByCompanyIDParams) (Configuration, error) {
	row := q.db.QueryRowContext(ctx, updateConfiguationByCompanyID,
		arg.CompanyID,
		arg.Pickup,
		arg.CashOnDelivery,
		arg.DigitalPayments,
		arg.Delivery,
	)
	var i Configuration
	err := row.Scan(
		&i.ID,
		&i.Delivery,
		&i.Pickup,
		&i.CashOnDelivery,
		&i.DigitalPayments,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateWorkingHourByCompanyID = `-- name: UpdateWorkingHourByCompanyID :one
UPDATE working_hours SET 
        day = $2,
        closes_at = $3,
        opens_at = $4
    WHERE company_id = $1
RETURNING id, day, opens_at, closes_at, company_id, created_at, updated_at
`

type UpdateWorkingHourByCompanyIDParams struct {
	CompanyID string    `json:"company_id"`
	Day       string    `json:"day"`
	ClosesAt  time.Time `json:"closes_at"`
	OpensAt   time.Time `json:"opens_at"`
}

func (q *Queries) UpdateWorkingHourByCompanyID(ctx context.Context, arg UpdateWorkingHourByCompanyIDParams) (WorkingHour, error) {
	row := q.db.QueryRowContext(ctx, updateWorkingHourByCompanyID,
		arg.CompanyID,
		arg.Day,
		arg.ClosesAt,
		arg.OpensAt,
	)
	var i WorkingHour
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.OpensAt,
		&i.ClosesAt,
		&i.CompanyID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}