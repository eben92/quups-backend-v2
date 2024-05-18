-- name: CreateCompany :one
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
RETURNING *;


-- name: UpdateCompany :one
UPDATE companies SET 
        name = $2, 
        email = $3,
        image_url = $4,
        about = $5,
        tin = $6,
        is_active = $7,
        msisdn = $8
    WHERE id = $1
RETURNING *;

-- name: GetCompanyByID :one
SELECT * FROM companies
    WHERE id = $1
    LIMIT 1;

-- WARNING: this will not work because of foreign key constraints
-- name: DeleteCompany :exec
DELETE FROM companies WHERE id = $1;

