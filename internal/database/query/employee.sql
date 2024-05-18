-- name: AddEmployee :one
INSERT INTO company_employees (
        company_id,
        name,
        email,
        msisdn,
        role,
        user_id
    )
VALUES (
    $1, $2, $3, $4, $5, $6
    )
RETURNING *;

-- name: UpdateEmployee :one
UPDATE company_employees SET 
        name = $2, 
        email = $3,
        role = $4,
        status = $5
    WHERE id = $1
RETURNING *;

-- name: UpdateEmployeeInvitationStatus :one
UPDATE company_employees SET 
        status = $2
    WHERE id = $1
RETURNING *;

-- name: GetEmployeesByCompanyID :many
SELECT * FROM company_employees
 WHERE company_id = $1
 ORDER BY created_at desc
 LIMIT $2 OFFSET $3;

 -- name: RemoveEmployee :exec
DELETE FROM company_employees WHERE id = $1;