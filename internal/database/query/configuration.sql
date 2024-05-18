-- name: CreateConfiguration :one
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
RETURNING *;

-- name: UpdateConfiguationByCompanyID :one
UPDATE configurations SET 
        pickup = $2,
        cash_on_delivery = $3,
        digital_payments = $4,
        delivery = $5
    WHERE company_id = $1
RETURNING *;

-- name: GetConfigurationByCompanyID :one
SELECT * FROM configurations
 WHERE company_id = $1
 LIMIT 1;

 -- name: RemoveEmployee :exec
DELETE FROM configurations WHERE id = $1;


-- WORKING HOURS

-- name: AddWorkingHour :one
INSERT INTO working_hours (
        company_id,
        day,
        opens_at,
        closes_at
    )
VALUES (
    $1, $2, $3, $4
    )
RETURNING *;

-- name: UpdateWorkingHourByCompanyID :one
UPDATE working_hours SET 
        day = $2,
        closes_at = $3,
        opens_at = $4
    WHERE company_id = $1
RETURNING *;


-- name: GetWorkingHoursByCompanyID :many
SELECT * FROM working_hours
 WHERE company_id = $1
 LIMIT 7;

 -- name: RemoveWorkingHour :exec
DELETE FROM working_hours WHERE id = $1;